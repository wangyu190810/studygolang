package controllers

import (
	"github.com/astaxie/beego"
	_ "beeHello/models"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.Ctx.WriteString(NotPV)
	//c.Ctx.
}
