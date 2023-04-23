package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"rpg-go/config"
	"rpg-go/internal/repository/sql"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func setup() (*http.Server, error) {
	// -> loading envs
	conf, err := config.LoadEnv("./")
	if err != nil {
		return &http.Server{}, err
	}

	// -> config and connect database
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.DBUser, conf.DBPassword),
		Host:   fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort),
		Path:   conf.DBName,
	}
	q := dsn.Query()
	q.Add("sslmode", conf.SSLMode)
	dsn.RawQuery = q.Encode()
	pool, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return &http.Server{}, err
	}

	// -> database healthcheck
	if err := pool.Ping(context.Background()); err != nil {
		return &http.Server{}, err
	}

	// -> config querier sqlc
	queries := sql.New(pool)
	querier := sql.Querier(queries)

	// -> config app
	app = config.AppConfig{
		DB:            pool,
		Querier:       querier,
		InProduction:  conf.Environment == "production",
		Domain:        conf.Domain,
		Version:       conf.Version,
		WebServerPort: conf.WebServerPort,
	}

	// -> config server
	svc := &http.Server{
		Addr:              conf.WebServerPort,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	return svc, nil
}
