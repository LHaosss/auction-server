// Code generated by goctl. DO NOT EDIT.

package auction_model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	auctionFieldNames          = builder.RawFieldNames(&Auction{})
	auctionRows                = strings.Join(auctionFieldNames, ",")
	auctionRowsExpectAutoSet   = strings.Join(stringx.Remove(auctionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	auctionRowsWithPlaceHolder = strings.Join(stringx.Remove(auctionFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAuctionServerAuctionIdPrefix  = "cache:auctionServer:auction:id:"
	cacheAuctionServerAuctionXidPrefix = "cache:auctionServer:auction:xid:"
)

type (
	auctionModel interface {
		Insert(ctx context.Context, data *Auction) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Auction, error)
		FindOneByXid(ctx context.Context, xid sql.NullString) (*Auction, error)
		Update(ctx context.Context, data *Auction) error
		Delete(ctx context.Context, id int64) error
		// manual functions
		FindMajority(ctx context.Context) (*[]*Auction, error) 
	}

	defaultAuctionModel struct {
		sqlc.CachedConn
		table string
	}

	Auction struct {
		Id            int64          `db:"id"`
		Xid           sql.NullString `db:"xid"`
		AuctionName   string         `db:"auction_name"`
		OfferUserXid  sql.NullString `db:"offer_user_xid"`
		CurrentPrice  sql.NullInt64  `db:"current_price"`
		PostUserXid   string         `db:"post_user_xid"`
		Type          string         `db:"type"`
		BasePrice     sql.NullInt64  `db:"base_price"`
		AuctionimgUrl string         `db:"auctionimg_url"`
	}
)

func newAuctionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAuctionModel {
	return &defaultAuctionModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`auction`",
	}
}

func (m *defaultAuctionModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	auctionServerAuctionIdKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionIdPrefix, id)
	auctionServerAuctionXidKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionXidPrefix, data.Xid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, auctionServerAuctionIdKey, auctionServerAuctionXidKey)
	return err
}

func (m *defaultAuctionModel) FindOne(ctx context.Context, id int64) (*Auction, error) {
	auctionServerAuctionIdKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionIdPrefix, id)
	var resp Auction
	err := m.QueryRowCtx(ctx, &resp, auctionServerAuctionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", auctionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAuctionModel) FindOneByXid(ctx context.Context, xid sql.NullString) (*Auction, error) {
	auctionServerAuctionXidKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionXidPrefix, xid)
	var resp Auction
	err := m.QueryRowIndexCtx(ctx, &resp, auctionServerAuctionXidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `xid` = ? limit 1", auctionRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, xid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAuctionModel) Insert(ctx context.Context, data *Auction) (sql.Result, error) {
	auctionServerAuctionIdKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionIdPrefix, data.Id)
	auctionServerAuctionXidKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionXidPrefix, data.Xid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, auctionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Xid, data.AuctionName, data.OfferUserXid, data.CurrentPrice, data.PostUserXid, data.Type, data.BasePrice, data.AuctionimgUrl)
	}, auctionServerAuctionIdKey, auctionServerAuctionXidKey)
	return ret, err
}

func (m *defaultAuctionModel) Update(ctx context.Context, newData *Auction) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	auctionServerAuctionIdKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionIdPrefix, data.Id)
	auctionServerAuctionXidKey := fmt.Sprintf("%s%v", cacheAuctionServerAuctionXidPrefix, data.Xid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, auctionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Xid, newData.AuctionName, newData.OfferUserXid, newData.CurrentPrice, newData.PostUserXid, newData.Type, newData.BasePrice, newData.AuctionimgUrl, newData.Id)
	}, auctionServerAuctionIdKey, auctionServerAuctionXidKey)
	return err
}

func (m *defaultAuctionModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAuctionServerAuctionIdPrefix, primary)
}

func (m *defaultAuctionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", auctionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAuctionModel) tableName() string {
	return m.table
}
