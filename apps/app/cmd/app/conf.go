package main

import (
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Host     string
	Port     int64
	Name     string
	User     string
	Password string
	PoolSize int64
}

func (d Database) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)
}

func (d Database) GetConfig() *pgxpool.Config {
	config, err := pgxpool.ParseConfig(d.URL())
	if err != nil {
		panic(err)
	}
	config.MinConns = 1
	config.MaxConns = int32(d.PoolSize)
	slog.Info("pool size", "size", config.MaxConns)
	return config
}

type Logging struct {
	Level string
}

type Cache struct {
	URL string
}

type Listen struct {
	Host string
	Port int64
}

type Conf struct {
	Cache    Cache
	Database Database
	Logging  Logging
	Listen   Listen
	JWTKey   string
}

var cfg Conf
