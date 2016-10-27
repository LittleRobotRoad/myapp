package controllers

import (
	"github.com/astaxie/beego"
	"org/myapp/models"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController)Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController)Delete() {
	id := this.Ctx.Input.Param("0")
	if len(id) == 0 {
		beego.Error("id is  null id:" + id)
		this.Redirect("/category", 301)
		return
	}
	err := models.DeleteCategory(id)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/category", 301)
}

func (this *CategoryController)Add() {
	name := this.Input().Get("name")
	Log(name)
	if len(name) == 0 {
		return
	}
	err := models.AddCategory(name)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/category", 301)
}