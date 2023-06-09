package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 用户客户端句柄
	UserRpcConf zrpc.RpcClientConf
}
