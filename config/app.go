package config

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"rpg-go/internal/repository/sql"
)

type AppConfig struct {
	DB            *pgxpool.Pool
	Querier       sql.Querier
	InProduction  bool
	Domain        string
	Version       string
	WebServerPort string
}
