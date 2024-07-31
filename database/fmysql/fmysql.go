package fmysql

import (
	"fmt"

	"github.com/ZYallers/fine/frame/fmsg"
	"github.com/ZYallers/fine/internal/consts"
	"github.com/ZYallers/fine/internal/instance"
	"github.com/ZYallers/fine/internal/intlog"
	"github.com/ZYallers/fine/os/flog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBFunc(name string) func() *gorm.DB {
	return func() *gorm.DB { return DB(name) }
}

func DB(name string) *gorm.DB {
	if name == "" {
		return nil
	}
	instanceKey := fmt.Sprintf("%s.%s", consts.FrameComponentDatabaseMysql, name)
	result := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		flog.Use().Info(fmt.Sprintf("mysql.%s.init", name))
		cfg := GetConfig(name)
		dialect := mysql.Open(ConvertToDSN(cfg))
		gormConfig := &gorm.Config{DisableAutomaticPing: true, Logger: intlog.NewGormLogger(name)}
		db, err := gorm.Open(dialect, gormConfig)
		if err != nil {
			return fmt.Errorf("db.%s.open.error: %s", name, err)
		}
		sdb, err := db.DB()
		if err != nil {
			return fmt.Errorf("db.%s.db.error: %s", name, err)
		}
		if err := sdb.Ping(); err != nil {
			return fmt.Errorf("db.%s.ping.error: %s", name, err)
		}
		sdb.SetMaxIdleConns(cfg.MaxIdleConns)
		sdb.SetMaxOpenConns(cfg.MaxOpenConns)
		sdb.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
		sdb.SetConnMaxLifetime(cfg.ConnMaxLifetime)
		return db
	})
	switch v := result.(type) {
	case *gorm.DB:
		return v
	case error:
		instance.Remove(instanceKey)
		flog.Use().Error(v.Error())
		fmsg.Sender().Simple(v.Error(), true)
		return nil
	default:
		instance.Remove(instanceKey)
		msg := fmt.Sprintf("db.%s.error: %s", name, "unknown error")
		flog.Use().Error(msg)
		fmsg.Sender().Simple(msg, true)
		return nil
	}
}
