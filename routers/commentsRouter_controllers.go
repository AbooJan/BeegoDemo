package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["BeegoDemo1/controllers:TestController"] = append(beego.GlobalControllerRouter["BeegoDemo1/controllers:TestController"],
		beego.ControllerComments{
			Method: "version",
			Router: `/version`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
