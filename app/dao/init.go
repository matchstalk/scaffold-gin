package dao

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/matchstalk/scaffold-gin/common/gdb"
	"github.com/spf13/viper"
	"log"
)

const DriverMysql = "mysql"

func Setup() {
	host := viper.GetString("database.mysql.host")
	user := viper.GetString("database.mysql.user")
	password := viper.GetString("database.mysql.password")
	name := viper.GetString("database.mysql.name")
	charset := viper.GetString("database.mysql.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, name, charset)
	var db *gorm.DB
	var err error
	switch viper.Get("database.driver") {
	case DriverMysql:
		db, err = gorm.Open("mysql", dsn)
	}
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect mysql %s", err.Error()))
	} else {
		db.DB().SetMaxIdleConns(viper.GetInt("database.mysql.pool.min"))
		db.DB().SetMaxOpenConns(viper.GetInt("database.mysql.pool.max"))
		if gin.Mode() != gin.ReleaseMode {
			db.LogMode(true)
		}
	}
	gdb.SetDB(db)
}
