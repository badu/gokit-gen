package route_guide

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

/**
type MysqlConfig struct {
	Host              string `env:"DB_READ_HOST" envDefault:"local-route_guide"`
	Port              string `env:"DB_READ_PORT" envDefault:"3306"`
	ConnectionTimeout int    `env:"DB_READ_CONNECTION_TIMEOUT" envDefault:"5"`
	MaxConnLifetime   int    `env:"DB_READ_MAX_CONN_LIFETIME" envDefault:"0"`
	MaxIdleConns      int    `env:"DB_READ_MAX_IDLE_CONNS" envDefault:"2"`
	MaxOpenConns      int    `env:"DB_READ_MAX_OPEN_CONNS" envDefault:"0"`
	ReadTimeout       int    `env:"DB_READ_READ_TIMEOUT" envDefault:"360"`
	Username          string `env:"DB_READ_USERNAME" envDefault:"root"`
	Password          string `env:"DB_READ_PASSWORD" envDefault:"foobar"`
	DatabaseName      string `env:"DB_READ_DATABASE" envDefault:"route_guide"`
}
**/
/**

**/
type Repository interface {
	GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error)
	HalfDuplex(ctx context.Context, lo *Point, hi *Point) (string, *Point, error)
	ReverseHalfDuplex(ctx context.Context, latitude int32, longitude int32) (int32, int32, int32, int32, error)
	FullDuplex(ctx context.Context, location *Point) (*Point, error)
}

type repositoryImpl struct {
	log log.Logger
	db  interface{} // TODO : use your own kind, e.g. *mongo.Client
}

func NewRepository(logger log.Logger, db interface{}) Repository {
	return &repositoryImpl{log: logger, db: db}
}

func (r *repositoryImpl) GetFeature(ctx context.Context, latitude int32, longitude int32) (string, *Point, error) {
	level.Error(r.log).Log("repository", "not implemented")
	return "", nil, errors.New("not implemented")
}

func (r *repositoryImpl) HalfDuplex(ctx context.Context, lo *Point, hi *Point) (string, *Point, error) {
	level.Error(r.log).Log("repository", "not implemented")
	return "", nil, errors.New("not implemented")
}

func (r *repositoryImpl) ReverseHalfDuplex(ctx context.Context, latitude int32, longitude int32) (int32, int32, int32, int32, error) {
	level.Error(r.log).Log("repository", "not implemented")
	return 0, 0, 0, 0, errors.New("not implemented")
}

func (r *repositoryImpl) FullDuplex(ctx context.Context, location *Point) (*Point, error) {
	level.Error(r.log).Log("repository", "not implemented")
	return nil, errors.New("not implemented")
}
