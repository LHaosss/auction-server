package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuctionRpcConf            zrpc.RpcClientConf
	UserRpcConf               zrpc.RpcClientConf
	AuctionInfoManagerRpcConf zrpc.RpcClientConf
}
