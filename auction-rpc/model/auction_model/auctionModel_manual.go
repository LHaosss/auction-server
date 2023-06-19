package auction_model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

func (m *defaultAuctionModel) FindMajority(ctx context.Context) (*[]*Auction, error) {
	var resp []*Auction
	query := fmt.Sprintf("select * from %s limit 16", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
