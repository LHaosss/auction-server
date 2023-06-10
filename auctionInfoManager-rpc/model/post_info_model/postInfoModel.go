package post_info_model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostInfoModel = (*customPostInfoModel)(nil)

type (
	// PostInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostInfoModel.
	PostInfoModel interface {
		postInfoModel
	}

	customPostInfoModel struct {
		*defaultPostInfoModel
	}
)

// NewPostInfoModel returns a model for the database table.
func NewPostInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PostInfoModel {
	return &customPostInfoModel{
		defaultPostInfoModel: newPostInfoModel(conn, c, opts...),
	}
}
