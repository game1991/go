package main

import (
	_ "firstAPI/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)//注册MySQL的driver
	orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/test2?charset=utf8")//本地数据库的账号。密码等
	orm.RunSyncdb("default", false, true)

}
func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"//静态文档
	}

	beego.Run()
}
