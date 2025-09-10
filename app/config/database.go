package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func dsn() string {
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "abc123")
	host := getenv("DB_HOST", "db")
	port := getenv("DB_PORT", "3306")
	name := getenv("DB_NAME", "appdb")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, pass, host, port, name)
}

func OpenGormWithRetry(ctx context.Context, maxWait time.Duration) (*gorm.DB, error) {
	deadline := time.Now().Add(maxWait)
	for {

		sqlDB, err := sql.Open("mysql", dsn())
		if err == nil {
			pctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			err = sqlDB.PingContext(pctx)
			cancel()
			if err == nil {
				return gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
			}
			_ = sqlDB.Close()
		}
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("waiting MySQL timed out: %w", err)
		}
		time.Sleep(3 * time.Second)
	}
}
