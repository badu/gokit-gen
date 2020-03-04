package factory

import (
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func createRepository(t *testing.T) (Repository, error) {
	var client interface{}
	// For MongoDB
	/**
	    opts := options.Client().ApplyURI("mongodb://root:foobar@localhost:27017")
		mongoClient, err := mongo.NewClient(opts)
		if err != nil {
			return nil, err
		}
		if err = mongoClient.Connect(context.Background()); err != nil {
			return nil, err
		}

		// cleanup previous data
		err = mongoClient.Database("factory").Drop(context.Background())
		if err != nil {
			t.Fatalf("failed to drop database %s: %s", "stock_transaction", err)
		}
	    **/
	// For MySQL _ "github.com/go-sql-driver/mysql"
	/**
	    conf := MysqlConfig{
	        Username:"root",
	        Password:"foobar",
	        Host:"localhost",
	        Port:"
	        ConnectionTimeout:5,
	        MaxConnLifetime:0,
	        MaxIdleConns:2,
	        MaxOpenConns:0,
	        ReadTimeout:360,
	    }
	    dbsrc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%ds&readTimeout=%ds&parseTime=true",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DatabaseName,
			conf.ConnectionTimeout,
			conf.ReadTimeout)
		mysqlClient, err := sql.Open("mysql", dbsrc)
		if err != nil {
			t.Fatalf("error : %v",err)
		}
		mysqlClient.SetMaxIdleConns(conf.MaxIdleConns)
		mysqlClient.SetMaxOpenConns(conf.MaxOpenConns)
		mysqlClient.SetConnMaxLifetime(time.Duration(conf.MaxConnLifetime) * time.Second)
		if err := mysqlClient.Ping(); err != nil {
			t.Fatalf("error : %v",err)
		}
	    **/
	// TODO : make connection
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.Caller(5))
	logger = level.NewFilter(logger, level.AllowDebug())

	return NewRepository(logger, client), nil
}

func TestMakeBox(t *testing.T) {
	repo, err := createRepository(t)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	_ = repo
	t.Skip("not implemented")
}
func TestGetBoxes(t *testing.T) {
	repo, err := createRepository(t)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	_ = repo
	t.Skip("not implemented")
}
func TestStatus(t *testing.T) {
	repo, err := createRepository(t)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	_ = repo
	t.Skip("not implemented")
}
