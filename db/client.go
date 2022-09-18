package db

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AI1411/go-psql_grpc_gql/config"
)

type Client struct {
	db *gorm.DB
}

func NewClient(c *config.Config, logger *zap.Logger) (*Client, error) {
	gormLogger := initGormLogger(logger)
	db, err := open(c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlPassword,
		c.Postgres.PostgresqlDbname,
		c.Postgres.PostgresqlPort,
	)

	db.Logger = db.Logger.LogMode(gormLogger.LogLevel)
	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}

func open(host, username, password, database, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, username, password, database, port,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (c *Client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}
