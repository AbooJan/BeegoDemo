package routers

import (
	"BeegoDemo1/controllers"

	"github.com/astaxie/beego"
)

// router init func
func init() {
	beego.Router("/test/version", &controllers.TestController{}, "*:Version")
	beego.Router("/test/user", &controllers.TestController{}, "get:Get")
	beego.Router("/test/friends", &controllers.TestController{}, "get:Friends")
	beego.Router("/test/login", &controllers.TestController{}, "post:Post")
	beego.Router("/test/uploadAvatar", &controllers.TestController{}, "post:UploadAvatar")
	beego.Router("/test/avatar", &controllers.TestController{}, "get:Avatar")
}
