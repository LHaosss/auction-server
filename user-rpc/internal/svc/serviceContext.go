package svc

import (
	"auction_server/user-rpc/internal/config"
	"auction_server/user-rpc/model/user_model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user_model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user_model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
	}
}
