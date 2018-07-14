package routers

import (
	"management/controllers"
	"github.com/astaxie/beego"

)
func init(){
	beego.Router("/", &controllers.UserController{} ,"get:First")
	beego.Router("/user", &controllers.UserController{},"get:Info") 
	beego.Router("/review", &controllers.UserController{},"get:Review") 	
	beego.Router("/v1/register", &controllers.UserController{}, "get:Register")
}
