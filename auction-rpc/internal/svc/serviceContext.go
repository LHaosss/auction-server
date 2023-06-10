package svc

import (
	"auction_server/auction-rpc/internal/config"
	"auction_server/auction-rpc/model/auction_model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AuctionModel auction_model.AuctionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AuctionModel: auction_model.NewAuctionModel(sqlx.NewMysql(c.DB.DataSource), c.CacheRedis),
	}
}
