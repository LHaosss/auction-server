package user

import (
	"context"
	"database/sql"

	"auction_server/user-api/internal/svc"
	"auction_server/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, sql.NullString{String: req.UserName, Valid: true})
	if err != nil {
		return &types.UserLoginResp{
			Flag:        false,
			Description: "未找到该用户",
		}, nil
	}

	if req.Password != user.Password.String {
		return &types.UserLoginResp{
			Flag:        false,
			Description: "密码错误",
		}, nil
	}

	return &types.UserLoginResp{
		Flag:        true,
		Description: "登陆成功",
	}, nil
}
