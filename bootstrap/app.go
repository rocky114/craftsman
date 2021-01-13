package bootstrap

import (
	"craftsman/config"
	"craftsman/route"
	"fmt"
	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	GlobalConfig config.Application
	MysqlConn    *gorm.DB
	Router       *gin.Engine
)

func init() {
	initConfig()
	initDatabase()
	initRouter()
	runServer()
}

func initConfig() {
	var configFile string
	flag.StringVarP(&configFile, "config", "c", "config.toml", "choose config file(shorthand)")

	flag.Parse()

	viper.SetConfigFile(configFile)

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		panic(fmt.Errorf("config parse error: %s \n", err))
	}
}

func initDatabase() {
	m := GlobalConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Options

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("mysql connect failed %s", err))
	}

	MysqlConn = dbConn
}

func initRouter() {
	Router = route.InitRouter(gin.Default())
}

func runServer() {
	endPoint := "0.0.0.0:8011"

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()

	if err != nil {
		fmt.Printf("server err: %s", err)
	}
}
