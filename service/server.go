package service

import (
	"github.com/astaxie/beego" // 引入beego框架
)

type MainController struct {
	beego.Controller // beego控制器
}

func (this *MainController) Get() {
	name := this.Ctx.Input.Param(":name")          // 得到uri中的name
	this.Ctx.WriteString("hello, " + name + "!\n") // 写入信息
}

func init() {
	beego.Router("/cloudgo/?:name", &MainController{}) // 路由设置
}
