package models

import (
	"ginDemo02/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DBEngine *xorm.Engine

func init() {
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType := sec.Key("TYPE").MustString("mysql")
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").MustString("root")
	password := sec.Key("PASSWORD").MustString("root")
	host := sec.Key("HOST").String()
	port := sec.Key("PORT").String()
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	dataSourceName := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	DBEngine, err = xorm.NewEngine(dbType, dataSourceName)
	if err != nil {
		log.Println(err)
	}

	// 表前缀
	tbPrefix := names.NewPrefixMapper(names.SnakeMapper{}, tablePrefix)
	DBEngine.SetTableMapper(tbPrefix)

	// 打印sql
	if setting.RunMode == "debug" {
		DBEngine.ShowSQL(true)
	}

	// 连接池的空闲数
	DBEngine.SetMaxIdleConns(10)
	// 最大打开连接数
	DBEngine.SetMaxOpenConns(100)
	// 连接的最大生存时间
	DBEngine.SetConnMaxLifetime(3)

	if err = DBEngine.Ping(); err != nil {
		log.Fatalf("ping to db fail! err:%+v", err)
	}
}