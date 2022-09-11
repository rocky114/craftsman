package bootstrap

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)

var DbConn *sqlx.DB

func initDb() {
	var err error
	dbConf := GlobalConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database, dbConf.Options)
	DbConn, err = sqlx.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("database initial err: %v", err)
	}

	DbConn.SetMaxIdleConns(10)
	DbConn.SetMaxOpenConns(100)
}

func initMigrate() {
	fmt.Println(DbConn.DriverName())
	migrations := &migrate.FileMigrationSource{
		Dir: GlobalConfig.Migrate.Path,
	}
	fmt.Println(DbConn)
	n, err := migrate.Exec(DbConn.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		logrus.Errorf("migrate sql err: %v", err)
	}

	logrus.Infof("migrate sql apply number: %d", n)
}
