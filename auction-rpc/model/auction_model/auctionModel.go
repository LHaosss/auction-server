package auction_model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AuctionModel = (*customAuctionModel)(nil)

type (
	// AuctionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAuctionModel.
	AuctionModel interface {
		auctionModel
	}

	customAuctionModel struct {
		*defaultAuctionModel
	}
)

// NewAuctionModel returns a model for the database table.
func NewAuctionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AuctionModel {
	return &customAuctionModel{
		defaultAuctionModel: newAuctionModel(conn, c, opts...),
	}
}
