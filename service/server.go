package service

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	name := this.Ctx.Input.Param(":name")
	this.Ctx.WriteString("hello, " + name + "!")
}

func init() {
	beego.Router("/cloudgo/?:name", &MainController{})
}
