package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	o := orm.NewOrm()

	user := new(models.User)

	user.Id ,_ =   strconv.Atoi(c.Input().Get("id"))

	o.Read(user)

	c.Data["User"] = user


	c.TplExt = "html"

	c.TplName = "index.html"
}


func (c *MainController) Insert() {
	o := orm.NewOrm()

	user := new(models.User)
	user.Name = c.Input().Get("name")
	o.Insert(user)
	c.TplName = "index.html"
}
