package main

import (
	_ "chartmiru/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"chartmiru/controllers"
)

func init() {
	// DB setup
	setupDB()	
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

func setupDB() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    mysqlUser := beego.AppConfig.String("mysqlUser")
    mysqlPass := beego.AppConfig.String("mysqlPass")
    // mysqlHost:= beego.AppConfig.String("mysqlHost")
	mysqlDb := beego.AppConfig.String("mysqlDb")

    orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPass+"@/"+mysqlDb+"?charset=utf8")
}
