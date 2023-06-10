package offer_info_model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OfferInfoModel = (*customOfferInfoModel)(nil)

type (
	// OfferInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOfferInfoModel.
	OfferInfoModel interface {
		offerInfoModel
	}

	customOfferInfoModel struct {
		*defaultOfferInfoModel
	}
)

// NewOfferInfoModel returns a model for the database table.
func NewOfferInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OfferInfoModel {
	return &customOfferInfoModel{
		defaultOfferInfoModel: newOfferInfoModel(conn, c, opts...),
	}
}
