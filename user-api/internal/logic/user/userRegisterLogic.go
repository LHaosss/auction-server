package user

import (
	"context"

	"auction_server/user-api/internal/svc"
	"auction_server/user-api/internal/types"
	"auction_server/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
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
	rpcResp, err := l.svcCtx.UserRpcClient.UserRegister(l.ctx, &usercenter.UserRegisterReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserRegisterResp{
		Flag:        rpcResp.GetFlag(),
		Description: rpcResp.GetDescription(),
	}, nil
}
