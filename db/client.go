package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/go-pg-ci-example/env"
)

type Client struct {
	db *gorm.DB
}

func NewClient(e *env.Env) (*Client, error) {
	db, err := open(e.Hostname, e.User, e.Password, e.Port, e.Dbname)
	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}

func open(host, username, password, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, username, password, database, port,
	)
	log.Printf("dsn: %s", dsn)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (c *Client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}
