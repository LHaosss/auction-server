// Code generated by goctl. DO NOT EDIT.

package post_info_model

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
	postInfoFieldNames          = builder.RawFieldNames(&PostInfo{})
	postInfoRows                = strings.Join(postInfoFieldNames, ",")
	postInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(postInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	postInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(postInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAuctionServerPostInfoIdPrefix  = "cache:auctionServer:postInfo:id:"
	cacheAuctionServerPostInfoXidPrefix = "cache:auctionServer:postInfo:xid:"
)

type (
	postInfoModel interface {
		Insert(ctx context.Context, data *PostInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PostInfo, error)
		FindOneByXid(ctx context.Context, xid sql.NullString) (*PostInfo, error)
		Update(ctx context.Context, data *PostInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPostInfoModel struct {
		sqlc.CachedConn
		table string
	}

	PostInfo struct {
		Id         int64          `db:"id"`
		Xid        sql.NullString `db:"xid"`
		UserXid    sql.NullString `db:"user_xid"`
		AuctionXid sql.NullString `db:"auction_xid"`
	}
)

func newPostInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultPostInfoModel {
	return &defaultPostInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`post_info`",
	}
}

func (m *defaultPostInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	auctionServerPostInfoIdKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoIdPrefix, id)
	auctionServerPostInfoXidKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoXidPrefix, data.Xid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, auctionServerPostInfoIdKey, auctionServerPostInfoXidKey)
	return err
}

func (m *defaultPostInfoModel) FindOne(ctx context.Context, id int64) (*PostInfo, error) {
	auctionServerPostInfoIdKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoIdPrefix, id)
	var resp PostInfo
	err := m.QueryRowCtx(ctx, &resp, auctionServerPostInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postInfoRows, m.table)
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

func (m *defaultPostInfoModel) FindOneByXid(ctx context.Context, xid sql.NullString) (*PostInfo, error) {
	auctionServerPostInfoXidKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoXidPrefix, xid)
	var resp PostInfo
	err := m.QueryRowIndexCtx(ctx, &resp, auctionServerPostInfoXidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `xid` = ? limit 1", postInfoRows, m.table)
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

func (m *defaultPostInfoModel) Insert(ctx context.Context, data *PostInfo) (sql.Result, error) {
	auctionServerPostInfoIdKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoIdPrefix, data.Id)
	auctionServerPostInfoXidKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoXidPrefix, data.Xid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, postInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Xid, data.UserXid, data.AuctionXid)
	}, auctionServerPostInfoIdKey, auctionServerPostInfoXidKey)
	return ret, err
}

func (m *defaultPostInfoModel) Update(ctx context.Context, newData *PostInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	auctionServerPostInfoIdKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoIdPrefix, data.Id)
	auctionServerPostInfoXidKey := fmt.Sprintf("%s%v", cacheAuctionServerPostInfoXidPrefix, data.Xid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Xid, newData.UserXid, newData.AuctionXid, newData.Id)
	}, auctionServerPostInfoIdKey, auctionServerPostInfoXidKey)
	return err
}

func (m *defaultPostInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAuctionServerPostInfoIdPrefix, primary)
}

func (m *defaultPostInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultPostInfoModel) tableName() string {
	return m.table
}