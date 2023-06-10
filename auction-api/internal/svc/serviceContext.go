package svc

import (
	"auction_server/auction-api/internal/config"
	"auction_server/auction-rpc/auctioncenter"
	"auction_server/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	AuctionRpcClient auctioncenter.Auctioncenter
	UserRpcClient    usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		AuctionRpcClient: auctioncenter.NewAuctioncenter(zrpc.MustNewClient(c.AuctionRpcConf)),
		UserRpcClient:    usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
