package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

type Env struct {
	Hostname string // GOPG_HOST
	Port     string // GOPG_PORT
	User     string // GOPG_USER
	Password string // GOPG_PASSWORD
	Dbname   string // GOPG_DBNAME
}

func NewClient() (*Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	e := &Env{
		Hostname: os.Getenv("GOPG_HOSTNAME"),
		Port:     os.Getenv("GOPG_PORT"),
		User:     os.Getenv("GOPG_USERNAME"),
		Password: os.Getenv("GOPG_PASSWORD"),
		Dbname:   os.Getenv("GOPG_DATABASE"),
	}
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
