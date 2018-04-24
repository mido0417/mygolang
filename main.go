package main

// import (
// 	_ "mygolang/routers"

// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// )

// func main() {
// 	beego.Run()
// 	orm.RunSyncdb("default", false, true)
// }


mport (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"myapp/controllers"
	"mygolang/models"
	_ "mygolang/routers"
)

func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 注册 beego 路由
	//beego.Router("/", &controllers.HomeController{})

	// 启动 beego
	beego.Run()
}

