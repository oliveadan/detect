package routers

import (
	"detect/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/start", &controllers.DetectController{}, "post:DoNow")
}
