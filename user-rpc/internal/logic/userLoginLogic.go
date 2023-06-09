package logic

import (
	"context"
	"database/sql"

	"auction_server/user-rpc/internal/svc"
	"auction_server/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{String: in.Username, Valid: true})
	if err != nil {
		return &pb.UserLoginResp{
			Flag:        false,
			Description: "未找到该用户",
		}, nil
	}

	if in.Password != user.Password.String {
		return &pb.UserLoginResp{
			Flag:        false,
			Description: "密码错误",
		}, nil
	}

	return &pb.UserLoginResp{
		Flag:        true,
		Description: "登陆成功",
	}, nil
}
