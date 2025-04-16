package main

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var args = cli.Command{
	Name:      "delivery",
	Usage:     "delivery service",
	UsageText: "starts delivery service",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "listen-host",
			Aliases:     []string{"l"},
			Value:       "",
			Destination: &cfg.Listen.Host,
			Usage:       "listen host",
			Sources:     cli.EnvVars("LISTEN_HOST"),
		},
		&cli.IntFlag{
			Name:        "listen-port",
			Aliases:     []string{"p"},
			Destination: &cfg.Listen.Port,
			Value:       8000,
			Usage:       "listen host",
			Sources:     cli.EnvVars("LISTEN_PORT"),
		},
		&cli.StringFlag{
			Name:        "db-name",
			Destination: &cfg.Database.Name,
			Usage:       "database name",
			Value:       "delivery",
			Sources:     cli.EnvVars("DB_NAME"),
		},
		&cli.IntFlag{
			Name:        "db-port",
			Destination: &cfg.Database.Port,
			Value:       5432,
			Usage:       "database name",
			Sources:     cli.EnvVars("DB_PORT"),
		},
		&cli.StringFlag{
			Name:        "db-host",
			Destination: &cfg.Database.Host,
			Usage:       "database name",
			Sources:     cli.EnvVars("DB_HOST"),
		},
		&cli.StringFlag{
			Name:        "db-user",
			Destination: &cfg.Database.User,
			Required:    true,
			Usage:       "database user",
			Sources:     cli.EnvVars("DB_USER"),
		},
		&cli.StringFlag{
			Name:        "db-password",
			Destination: &cfg.Database.Password,
			Required:    true,
			Usage:       "database password",
			Sources:     cli.EnvVars("DB_PASSWORD"),
		},
		&cli.IntFlag{
			Name:        "db-pool-size",
			Destination: &cfg.Database.PoolSize,
			Usage:       "database connections pool size",
			Value:       20,
			Sources:     cli.EnvVars("DB_POOL_SIZE"),
		},
		&cli.StringFlag{
			Name:        "redis-address",
			Destination: &cfg.Redis.Address,
			Required:    true,
			Usage:       "redis address",
			Sources:     cli.EnvVars("REDIS_ADDRESS"),
		},
		&cli.StringFlag{
			Name:        "redis-password",
			Destination: &cfg.Redis.Password,
			Value:       "",
			Usage:       "redis password",
			Sources:     cli.EnvVars("REDIS_PASSWORD"),
		},
		&cli.IntFlag{
			Name:        "redis-database",
			Destination: &cfg.Redis.Database,
			Value:       0,
			Usage:       "redis database",
			Sources:     cli.EnvVars("REDIS_DATABASE"),
		},
		&cli.IntFlag{
			Name:        "redis-pool-sizr",
			Destination: &cfg.Redis.PoolSize,
			Value:       1000,
			Usage:       "redis pool size",
			Sources:     cli.EnvVars("REDIS_POOL_SIZE"),
		},
		&cli.StringFlag{
			Name:        "log-level",
			Destination: &cfg.Logging.Level,
			Value:       "INFO",
			Usage:       "log level",
			Sources:     cli.EnvVars("LOG_LEVEL"),
			Validator: func(s string) error {
				if _, ok := LogLevels[s]; !ok {
					return fmt.Errorf("unexpected log level %s", s)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "jwt-key",
			Destination: &cfg.JWTKey,
			Value:       "", // NOTE: We allow empty because this is purely test app.
			Usage:       "jwt signing key",
			Sources:     cli.EnvVars("JWT_KEY"),
		},
	},
	Action: run,
}
