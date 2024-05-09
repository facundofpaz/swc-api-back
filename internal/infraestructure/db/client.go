package db

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Client interface {
	Select(ctx context.Context, dest any, query string, args ...any) error
	Insert() (uint64, error)
}

type client struct {
	db *sqlx.DB
}

func New() (Client, error) {
	usuario := "root"
	pass := "root"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "agenda"
	newDb, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	newClient := client{
		db: newDb,
	}
	return newClient, nil
}

func (c client) Insert() (uint64, error) {
	panic("unimplemented")
}

func (c client) Select(ctx context.Context, dest any, query string, args ...any) error {
	if err := c.db.SelectContext(ctx, dest, query, args...); err != nil {
		return err
	}

	return nil
}
