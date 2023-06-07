package repository

import (
	"book-store/models"
	"book-store/services"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	helperModel "git.innovasive.co.th/backend/models"
)

type psqlRepository struct {
	conn *sql.DB
}

func NewRepository(conn *sql.DB) services.RepositoryInterface {
	return &psqlRepository{
		conn: conn,
	}
}

func (p psqlRepository) FetchListBooks(ctx context.Context, args *sync.Map, paginator *helperModel.Paginator) ([]*models.Books, error) {
	var paginatorSql string
	var books []*models.Books
	if paginator != nil {
		var limit = int(paginator.PerPage)
		var skipItem = (int(paginator.Page) - 1) * int(paginator.PerPage)
		paginatorSql = fmt.Sprintf(`LIMIT %d
			OFFSET %d
			`,
			limit,
			skipItem,
		)
	}
	conds, valArgs := whereListBook(args)
	var where string
	if len(conds) > 0 {
		where = "WHERE " + strings.Join(conds, " AND ")
	}
	sql := fmt.Sprintf(`
		select
			*
		from
			books
		%s
		%s
	`,
		where,
		paginatorSql,
	)
	if err := p.conn.QueryRowContext(ctx, sql, valArgs...).Scan(&books); err != nil {
		return nil, err
	}
	return books, nil
}

func whereListBook(args *sync.Map) ([]string, []interface{}) {
	var conds []string
	var valArgs []interface{}

	if v, ok := args.Load("search_word"); ok {
		if v != nil && v != "" {
			sql := `books.title like CONCAT('%%',?,'%%')`
			conds = append(conds, sql)
			valArgs = append(valArgs, v.(string))
		}
	}
	if min, ok := args.Load("min_price"); ok {
		if max, ok := args.Load("max_price"); ok {
			sql := `books.price >= ? and books.price <= ?`
			conds = append(conds, sql)
			valArgs = append(valArgs, min.(string), max.(string))
		} else {
			sql := `books.price = ?`
			conds = append(conds, sql)
			valArgs = append(valArgs, min.(string))
		}
	}

	return conds, valArgs
}
