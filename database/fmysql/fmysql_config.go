package fmysql

import (
	"fmt"
	"time"

	"github.com/ZYallers/fine/os/fcfg"
)

const (
	DefaultPort             = "3306"
	DefaultCharset          = "utf8mb4"
	DefaultLoc              = "Local"
	DefaultParseTime        = "true"
	DefaultMaxAllowedPacket = "0"
	DefaultTimeout          = "15s"
	DefaultMaxIdleConns     = 5
	DefaultMaxOpenConns     = 50
	DefaultConnMaxIdleTime  = 5 * time.Minute
	DefaultConnMaxLifeTime  = 10 * time.Minute
)

type mysqlConfig struct {
	Host             string
	Port             string
	User             string
	Password         string
	Database         string
	Charset          string
	Loc              string
	ParseTime        string
	MaxAllowedPacket string
	Timeout          string
	MaxIdleConns     int
	MaxOpenConns     int
	ConnMaxIdleTime  time.Duration
	ConnMaxLifetime  time.Duration
}

func GetConfig(name string) *mysqlConfig {
	cfg := mysqlConfig{
		Host:             fcfg.GetEnvString(fmt.Sprintf("database.mysql.%s.host", name)),
		Port:             fcfg.GetEnvString(fmt.Sprintf("database.mysql.%s.port", name)),
		User:             fcfg.GetEnvString(fmt.Sprintf("database.mysql.%s.user", name)),
		Password:         fcfg.GetEnvString(fmt.Sprintf("database.mysql.%s.password", name)),
		Database:         fcfg.GetEnvString(fmt.Sprintf("database.mysql.%s.database", name)),
		Charset:          fcfg.GetString(fmt.Sprintf("database.mysql.%s.charset", name)),
		Loc:              fcfg.GetString(fmt.Sprintf("database.mysql.%s.loc", name)),
		ParseTime:        fcfg.GetString(fmt.Sprintf("database.mysql.%s.parseTime", name)),
		MaxAllowedPacket: fcfg.GetString(fmt.Sprintf("database.mysql.%s.maxAllowedPacket", name)),
		Timeout:          fcfg.GetString(fmt.Sprintf("database.mysql.%s.timeout", name)),
		MaxIdleConns:     fcfg.GetInt(fmt.Sprintf("database.mysql.%s.maxIdleConns", name)),
		MaxOpenConns:     fcfg.GetInt(fmt.Sprintf("database.mysql.%s.maxOpenConns", name)),
		ConnMaxIdleTime:  fcfg.GetDuration(fmt.Sprintf("database.mysql.%s.connMaxIdleTime", name)) * time.Second,
		ConnMaxLifetime:  fcfg.GetDuration(fmt.Sprintf("database.mysql.%s.connMaxLifetime", name)) * time.Second,
	}
	if cfg.Port == "" {
		cfg.Port = DefaultPort
	}
	if cfg.Charset == "" {
		cfg.Charset = DefaultCharset
	}
	if cfg.ParseTime == "" {
		cfg.ParseTime = DefaultParseTime
	}
	if cfg.Loc == "" {
		cfg.Loc = DefaultLoc
	}
	if cfg.MaxAllowedPacket == "" {
		cfg.MaxAllowedPacket = DefaultMaxAllowedPacket
	}
	if cfg.Timeout == "" {
		cfg.Timeout = DefaultTimeout
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = DefaultMaxIdleConns
	}
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = DefaultMaxOpenConns
	}
	if cfg.ConnMaxIdleTime <= 0 {
		cfg.ConnMaxIdleTime = DefaultConnMaxIdleTime
	}
	if cfg.ConnMaxLifetime <= 0 {
		cfg.ConnMaxLifetime = DefaultConnMaxLifeTime
	}
	return &cfg
}

func ConvertToDSN(cfg *mysqlConfig) string {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s&maxAllowedPacket=%s&timeout=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset, cfg.ParseTime, cfg.Loc,
		cfg.MaxAllowedPacket, cfg.Timeout)
	return dns
}
