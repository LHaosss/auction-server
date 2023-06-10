package user

import (
	"context"

	"auction_server/user-api/internal/svc"
	"auction_server/user-api/internal/types"
	"auction_server/user-rpc/usercenter"

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
	rpcResp, err := l.svcCtx.UserRpcClient.UserLogin(l.ctx, &usercenter.UserLoginReq{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		Id:       int(rpcResp.GetId()),
		Xid:      rpcResp.GetXid(),
		UserName: rpcResp.GetUsername(),
	}, nil
}
