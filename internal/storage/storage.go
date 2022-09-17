package storage

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rocky114/craftsman/internal/config"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

var dbConn *sqlx.DB

func InitDatabase() {
	var err error

	dbConf := config.GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database, dbConf.Options)
	dbConn, err = sqlx.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("database initial err: %v", err)
	}

	dbConn.SetMaxIdleConns(10)
	dbConn.SetMaxOpenConns(100)
}

func InitMigrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: config.GlobalConfig.Migrate.Path,
	}
	n, err := migrate.Exec(dbConn.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		logrus.Errorf("migrate sql err: %v", err)
	}

	logrus.Infof("migrate sql apply number: %d", n)
}

var Querier *Queries
var queriesOnce sync.Once

func GetQueries() *Queries {
	queriesOnce.Do(func() {
		Querier = New(dbConn.DB)
	})

	return Querier
}
