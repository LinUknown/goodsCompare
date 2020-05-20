package datasource

import (
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB{
	return db
}

func init(){
	//path := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8&parseTime=true"}, "")
	path := "root:710069741@(47.112.216.17:3306)/bs?charset=utf8&parseTime=true"
	var err error
	db,err = gorm.Open("mysql",path)
	if err != nil{
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)	//表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	//Createtable()
}