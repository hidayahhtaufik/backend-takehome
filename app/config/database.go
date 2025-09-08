package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func OpenGorm() (*gorm.DB, error) {
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "abc123")
	host := getenv("DB_HOST", "db") // service name di compose
	port := getenv("DB_PORT", "3306")
	name := getenv("DB_NAME", "appdb")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		user, pass, host, port, name)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
