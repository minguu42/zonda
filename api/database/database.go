package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/minguu42/zonda/api/config"
	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ErrModelNotFound = errors.New("model not found in database")

type Client struct {
	gormDB *gorm.DB
}

func NewClient(conf config.DB) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local&parseTime=True",
		conf.User,
		conf.Password,
		net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)),
		conf.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime)

	maxRetryCount := 20
	for i := range maxRetryCount {
		if err := db.Ping(); err != nil {
			break
		}
		if i == maxRetryCount {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		time.Sleep(1 * time.Second)
	}

	gormDB, err := gorm.Open(gormMySQL.New(gormMySQL.Config{Conn: db}), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm client: %w", err)
	}
	return &Client{gormDB: gormDB}, nil
}

func (c *Client) Close() error {
	db, err := c.gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get *sql.DB: %w", err)
	}
	return db.Close()
}

func (c *Client) db(ctx context.Context) *gorm.DB {
	return c.gormDB.WithContext(ctx)
}
