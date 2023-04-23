package config

import (
	"rpg-go/internal/repository/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AppConfig struct {
	DB            *pgxpool.Pool
	Querier       sql.Querier
	InProduction  bool
	Domain        string
	Version       string
	WebServerPort string
}
