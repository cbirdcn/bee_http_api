package main

import (
	_ "bee_http_game/routers"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 错误日志记录
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

	// 注册数据库驱动（mysql已默认注册，可忽略）
	//_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// beego要求必须注册一个别名为default的数据库，并且会在项目启动时初始化default数据库连接
	user, _ := web.AppConfig.String("mysqluser")
	pass, _ := web.AppConfig.String("mysqlpass")
	url, _ := web.AppConfig.String("mysqlurls")
	db, _ := web.AppConfig.String("mysqldb")
	port, _ := web.AppConfig.Int("mysqlport")
	maxIdleCons, _ := web.AppConfig.Int("maxIdleCons")
	maxOpenCons, _ := web.AppConfig.Int("maxOpenCons")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pass, url, port, db)
	
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		panic("db connection error")
	} else {
		// 最大空闲连接
		orm.SetMaxIdleConns("default", maxIdleCons)
		// 最大数据库连接
		orm.SetMaxOpenConns("default", maxOpenCons)
	}
	// orm操作打印日志（dev环境，默认os.stderr）
	//orm.Debug = true

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	// 启动
	beego.Run()
}

