package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/config"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.PGConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)
	return sql.Open("postgres", dsn)
}
