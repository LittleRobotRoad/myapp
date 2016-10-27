package main

import (
	"org/myapp/controllers"
	"org/myapp/models"
	_ "org/myapp/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{}) //自动路由 自动截取XxxxController的Yyyyy函数变为xxxxx/yyyyy路由
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.CategoryController{}) //自动路由 自动截取XxxxController的Yyyyy函数变为xxxxx/yyyyy路由
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	////作为静态文件
	//beego.SetStaticPath("/attachment", "attachment")
	//作为单独一个控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})

	//i18n.SetMessage("en-US", "conf/locale_en-US.ini")
	//i18n.SetMessage("zh-CN", "conf/locale_zh-CN.ini")
	//beego.AddFuncMap("i18n",i18n.Tr)
	//beego.Router("/test",&controllers.MainController{})
	beego.Run()
}
