package fredis

import (
	"fmt"
	"time"

	"github.com/ZYallers/fine/os/fcfg"
	"github.com/go-redis/redis"
)

type redisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 int
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
}

func getRedisOptions(name string) *redis.Options {
	cfg := redisConfig{
		Host:               fcfg.GetEnvString(fmt.Sprintf("database.redis.%s.host", name)),
		Port:               fcfg.GetEnvString(fmt.Sprintf("database.redis.%s.port", name)),
		Password:           fcfg.GetEnvString(fmt.Sprintf("database.redis.%s.password", name)),
		Db:                 fcfg.GetInt(fmt.Sprintf("database.redis.%s.db", name)),
		DialTimeout:        fcfg.GetDuration(fmt.Sprintf("database.redis.%s.dialTimeout", name)) * time.Second,
		ReadTimeout:        fcfg.GetDuration(fmt.Sprintf("database.redis.%s.readTimeout", name)) * time.Second,
		WriteTimeout:       fcfg.GetDuration(fmt.Sprintf("database.redis.%s.writeTimeout", name)) * time.Second,
		PoolSize:           fcfg.GetInt(fmt.Sprintf("database.redis.%s.poolSize", name)),
		MinIdleConns:       fcfg.GetInt(fmt.Sprintf("database.redis.%s.minIdleConns", name)),
		MaxConnAge:         fcfg.GetDuration(fmt.Sprintf("database.redis.%s.maxConnAge", name)) * time.Second,
		PoolTimeout:        fcfg.GetDuration(fmt.Sprintf("database.redis.%s.poolTimeout", name)) * time.Second,
		IdleTimeout:        fcfg.GetDuration(fmt.Sprintf("database.redis.%s.idleTimeout", name)) * time.Second,
		IdleCheckFrequency: fcfg.GetDuration(fmt.Sprintf("database.redis.%s.idleCheckFrequency", name)) * time.Second,
	}
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Db,
	}
	if cfg.DialTimeout > 0 {
		options.DialTimeout = cfg.DialTimeout
	}
	if cfg.ReadTimeout > 0 {
		options.ReadTimeout = cfg.ReadTimeout
	}
	if cfg.WriteTimeout > 0 {
		options.WriteTimeout = cfg.WriteTimeout
	}
	if cfg.PoolSize > 0 {
		options.PoolSize = cfg.PoolSize
	}
	if cfg.MinIdleConns > 0 {
		options.MinIdleConns = cfg.MinIdleConns
	}
	if cfg.MaxConnAge > 0 {
		options.MaxConnAge = cfg.MaxConnAge
	}
	if cfg.PoolTimeout > 0 {
		options.PoolTimeout = cfg.PoolTimeout
	}
	if cfg.IdleTimeout > 0 {
		options.IdleTimeout = cfg.IdleTimeout
	}
	if cfg.IdleCheckFrequency > 0 {
		options.IdleCheckFrequency = cfg.IdleCheckFrequency
	}
	return options
}
