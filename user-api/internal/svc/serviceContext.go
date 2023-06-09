package svc

import (
	"auction_server/user-api/internal/config"
	"auction_server/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
