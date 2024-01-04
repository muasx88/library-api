package database

import (
	"time"

	"github.com/muasx88/library-api/internal/config"
)

func getMaxOpenConns() int {
	var defaultMaxConns uint8
	defaultMaxConns = 3
	maxConnEnv := config.Config.DB.ConnectionPool.MaxOpenConnetcion

	if maxConnEnv > 0 {
		defaultMaxConns = maxConnEnv
	}

	return int(defaultMaxConns)
}

func getMaxIdleConns() int {
	var defaultMaxConns uint8
	defaultMaxConns = 3
	maxIdleConnEnv := config.Config.DB.ConnectionPool.MaxIdleConnection

	if maxIdleConnEnv > 0 {
		defaultMaxConns = maxIdleConnEnv
	}

	return int(defaultMaxConns)
}

func getConnMaxIdleTime() time.Duration {
	var defaultIdleTime uint8
	defaultIdleTime = 120
	idleTimeEnv := config.Config.DB.ConnectionPool.MaxIdletimeConnection

	if idleTimeEnv > 0 {
		defaultIdleTime = idleTimeEnv
	}

	return time.Duration(defaultIdleTime)
}
