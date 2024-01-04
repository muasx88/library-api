package database

import (
	"context"
	"fmt"
	"log"

	cfg "github.com/muasx88/library-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBConn *sqlx.DB

func createConnectionDB(ctx context.Context) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Config.DB.User,
		cfg.Config.DB.Password,
		cfg.Config.DB.Host,
		cfg.Config.DB.Port,
		cfg.Config.DB.Name,
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// db.SetConnMaxLifetime(time.Duration(cfg.Config.DB.ConnectionPool.MaxLifetimeConnection) * time.Second)
	// db.SetConnMaxIdleTime(getConnMaxIdleTime())
	// db.SetMaxOpenConns(getMaxOpenConns())
	// db.SetMaxIdleConns(getMaxIdleConns())

	return db, nil
}

func ConnectDB(ctx context.Context) (*sqlx.DB, error) {
	if DBConn == nil {
		var err error
		DBConn, err = createConnectionDB(ctx)
		if err != nil {
			return nil, err
		}

		log.Println("Database Connected....")
	}

	return DBConn, nil
}
