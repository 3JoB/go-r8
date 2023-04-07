package db

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/3JoB/go-r8/internal/config"
	errs "github.com/3JoB/ulib/err"
)

var (
	db *gorm.DB
	kc = config.F()
)

func init() {
	var (
		dialector gorm.Dialector
		conf    gorm.Config
	)
	client := kc.String("database.client")
	switch client {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", kc.String("database.mysql.user"), kc.String("database.mysql.pass") , kc.String("database.mysql.addr") , kc.String("database.mysql.db"))
		dialector = mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		})
		conf = gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
	case "sqlite", "sqlite3":
		dialector = sqlite.Open(config.SqlitePath())
	case "pgsql":
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",kc.String("database.pgsql.addr"), kc.String("database.pgsql.user"), kc.String("database.pgsql.pass"), kc.String("database.pgsql.db"), kc.String("database.pgsql.port"), kc.String("database.pgsql.zone"))
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	default:
		err := &errs.Err{Op: "internal/db/db.Client", Err: "Invalid database client!!!"}
		panic(err)
	}

	conn, err := gorm.Open(dialector, &conf)
	if err != nil {
		err := &errs.Err{Op: "internal/db/db.Open", E: err}
		panic(err)
	}
	sqlDB, err := conn.DB()
	if err != nil {
		err := &errs.Err{Op: "internal/db/db.conn", E: err}
		panic(err)
	}
	sqlDB.SetMaxOpenConns(kc.Int("database.connect.maxopen"))
	sqlDB.SetMaxIdleConns(kc.Int("database.connect.maxidle"))
	sqlDB.SetConnMaxLifetime(time.Duration(kc.Int("database.connect.maxlife")) * time.Second)

	db = conn
}

func NewDB() *gorm.DB {
	sqlDB, err := db.DB()
	if err != nil {
		err := &errs.Err{Op: "internal/db/db.conn", E: err}
		panic(err)
	}
	db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})
	if err = sqlDB.Ping(); err != nil {
		err := &errs.Err{Op: "internal/db/db.Ping", E: err}
		panic(err)
	}
	return db
}
