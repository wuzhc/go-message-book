package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbuser := beego.AppConfig.String("db.user")
	dbpwd := beego.AppConfig.String("db.password")
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	dbcharset := beego.AppConfig.String("db.charset")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	maxIdle := 30 // 最大空闲链接
	maxConn := 30 // 最大数据库链接
	dsn := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + dbcharset
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)
	orm.RegisterModelWithPrefix(
		"tb_",
		new(Topic),
		new(Comment),
	)

	// 开启orm调试模式
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
