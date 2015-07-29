package routers

import (
	"home/john/myworking/helloworld/btest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
