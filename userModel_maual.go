package user_model

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"

// 	"github.com/zeromicro/go-zero/core/stores/sqlc"
// 	"github.com/zeromicro/go-zero/core/stores/sqlx"
// )

// func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username sql.NullString) (*User, error) {
// 	var resp User
// 	err := m.QueryRowCtx(ctx, resp, "", func(ctx context.Context, conn sqlx.SqlConn, v any) error { // QueryCtxFn func(ctx context.Context, conn sqlx.SqlConn, v any) error
// 		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
// 		return conn.QueryRowCtx(ctx, v, query, username)
// 	})

// 	switch err {
// 	case nil:
// 		return &resp, nil
// 	case sqlc.ErrNotFound:
// 		return nil, ErrNotFound
// 	default:
// 		return nil, err
// 	}
// }
