package postgre

import (
	"context"
	"time"
	"log"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sony/gobreaker"
	"github.com/nahwinrajan/resilient-go/products-service/internal/config"
)

// DB entry point for db handler structure
type DB struct {
	cfg config.Database
	connObj *sqlx.DB
	cb *gobreaker.CircuitBreaker
}

// New return pointer instance of DB
func New(conf config.Database) (*DB, error) {
	dbConnection, err := sqlx.Connect(conf.Driver, conf.Connection)
	if err != nil {
		return nil, err
	}
	dbConnection.SetMaxOpenConns(conf.MaxOpenConn)
	dbConnection.SetMaxIdleConns(conf.MaxIdleConn)
	dbConnection.SetConnMaxLifetime(time.Duration(conf.MaxConnLifetime) * time.Second)

	cbSetting := gobreaker.Settings{
		Name:        conf.CircuitBreaker.Name,
		MaxRequests: uint32(conf.CircuitBreaker.MaxRequest),
		Interval:    time.Duration(conf.CircuitBreaker.Interval) * time.Second,
		Timeout:     time.Duration(conf.CircuitBreaker.TimeOut) * time.Second,
		OnStateChange: func(name string, from, to gobreaker.State) {
			if to == gobreaker.StateOpen && from == gobreaker.StateClosed {
				// in real word sample we will do slack notification here
				log.Printf("[circuit-breaker][cb-%s] is open, database is unavailable", name)
			}

			if from == gobreaker.StateHalfOpen && to == gobreaker.StateClosed {
				// in real word sample we will do slack notification here
				log.Printf("[circuit-breaker][cb-%s] is closed, database back to normal", name)
			}
		},
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > uint32(conf.CircuitBreaker.ConsecutiveFailures)
		},
	}
	cbCli := gobreaker.NewCircuitBreaker(cbSetting)

	dbHandler := DB{
		cfg: conf,
		connObj: dbConnection,
		cb: cbCli,
	}

	return &dbHandler, nil
}


func (dbObj *DB) queryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	objInterface, err := dbObj.cb.Execute(func() (interface{}, error) {
		return dbObj.connObj.QueryxContext(ctx, query, args...)
	})

	if (err == ErrOpenState || err == ErrTooManyRequests) && objInterface == nil {
		// send metrics
		return nil, errors.New("database is unavailable")
	}

	if objInterface == nil {
		return nil, err
	}

	return objInterface.(*sqlx.Rows), err
}