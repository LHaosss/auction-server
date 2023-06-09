package svc

import (
	"auction_server/user-api/internal/config"
	user "auction_server/user-api/model/user_model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
	}
}
