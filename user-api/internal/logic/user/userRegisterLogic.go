package user

import (
	"context"
	"database/sql"
	"errors"

	"auction_server/user-api/internal/svc"
	"auction_server/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	m_user "auction_server/user-api/model/user_model"
)

const (
	ErrWithParamer = "something wrong with paramer"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// 验证请求
	err = validUserRegisterReq(req)
	if err != nil {
		return &types.UserRegisterResp{
			Flag:        false,
			Description: ErrWithParamer,
		}, nil
	}

	// 	// 查询数据库是否用户已存在
	_, err = l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{String: req.UserName, Valid: true})
	if err == nil {
		return &types.UserRegisterResp{
			Flag:        false,
			Description: "该用户已存在",
		}, nil
	}
	if err != nil && err != sqlx.ErrNotFound {
		return &types.UserRegisterResp{
			Flag:        false,
			Description: "注册出错，请重试1",
		}, nil
	}

	// 用户不存在，向数据库中添加用户信息
	data := &m_user.User{
		Username: sql.NullString{
			String: req.UserName,
			Valid:  true,
		},
		Password: sql.NullString{
			String: req.Password,
			Valid:  true,
		},
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, data)
	if err != nil {
		return &types.UserRegisterResp{
			Flag:        false,
			Description: "注册出错，请重试2",
		}, nil
	}

	return &types.UserRegisterResp{
		Flag:        true,
		Description: "用户创建成功",
	}, nil
}

func validUserRegisterReq(req *types.UserRegisterReq) error {
	if req.Password == "" || req.UserName == "" {
		return errors.New(ErrWithParamer)
	}

	return nil
}
