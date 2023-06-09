package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"auction_server/user-rpc/internal/svc"
	"auction_server/user-rpc/model/user_model"
	"auction_server/user-rpc/pb"

	"github.com/rs/xid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	// 验证请求
	err := validUserRegisterReq(in)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// // 查询数据库是否用户已存在
	_, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{String: in.Username, Valid: true})
	if err == nil {
		fmt.Println(err)
		return nil, errors.New("该用户已存在")
	}

	if err != nil && err != sqlx.ErrNotFound {
		fmt.Println("####", err)
		return nil, errors.New("注册出错，请重试1")
	}

	// 用户不存在，向数据库中添加用户信息
	userXid := xid.New().String()
	data := &user_model.User{
		Username: sql.NullString{
			String: in.Username,
			Valid:  true,
		},
		Password: sql.NullString{
			String: in.Password,
			Valid:  true,
		},
		Xid: userXid,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, data)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("注册出错，请重试2")
	}

	return &pb.UserRegisterResp{
		Flag:        true,
		Description: "用户创建成功",
	}, nil
}

func validUserRegisterReq(pb *pb.UserRegisterReq) error {
	if pb.Password == "" || pb.Username == "" {
		return errors.New("请求参数出错")
	}

	return nil
}
