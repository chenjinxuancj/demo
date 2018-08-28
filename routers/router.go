package routers

import (
	"hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/?:id", &controllers.MainController{})

	beego.Router("/insert", &controllers.MainController{},"*:Insert")
}
