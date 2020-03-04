package example

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

/**
type MysqlConfig struct {
	Host              string `env:"DB_READ_HOST" envDefault:"local-example"`
	Port              string `env:"DB_READ_PORT" envDefault:"3306"`
	ConnectionTimeout int    `env:"DB_READ_CONNECTION_TIMEOUT" envDefault:"5"`
	MaxConnLifetime   int    `env:"DB_READ_MAX_CONN_LIFETIME" envDefault:"0"`
	MaxIdleConns      int    `env:"DB_READ_MAX_IDLE_CONNS" envDefault:"2"`
	MaxOpenConns      int    `env:"DB_READ_MAX_OPEN_CONNS" envDefault:"0"`
	ReadTimeout       int    `env:"DB_READ_READ_TIMEOUT" envDefault:"360"`
	Username          string `env:"DB_READ_USERNAME" envDefault:"root"`
	Password          string `env:"DB_READ_PASSWORD" envDefault:"foobar"`
	DatabaseName      string `env:"DB_READ_DATABASE" envDefault:"example"`
}
**/
/**

**/
type Repository interface {
	HalfDuplex(ctx context.Context, msg string) (string, error)
	FullDuplex(ctx context.Context, msg string) (string, error)
}

type repositoryImpl struct {
	log log.Logger
	db  interface{} // TODO : use your own kind, e.g. *mongo.Client
}

func NewRepository(logger log.Logger, db interface{}) Repository {
	return &repositoryImpl{log: logger, db: db}
}

func (r *repositoryImpl) HalfDuplex(ctx context.Context, msg string) (string, error) {
	level.Error(r.log).Log("repository", "repository not implemented")
	return "", errors.New("repository not implemented")
}

func (r *repositoryImpl) FullDuplex(ctx context.Context, msg string) (string, error) {
	level.Error(r.log).Log("repository", "repository not implemented")
	return "", errors.New("repository not implemented")
}
