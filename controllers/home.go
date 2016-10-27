package controllers

import (
	"github.com/astaxie/beego"
	"org/myapp/models"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"),this.Input().Get("label"),true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories

}

func Log(log ...string) {
	fmt.Println("********************")
	for _, v := range log {
		fmt.Println(v)
	}
	fmt.Println("********************")
}