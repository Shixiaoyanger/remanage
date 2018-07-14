package routers

import (
	"management/controllers"
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)
func init(){
	beego.Router("/admin", &controllers.AdminController{},"get:Admin") 
	beego.Router("/admin/status", &controllers.AdminController{},"get:Status")
	beego.Router("/admin/manage", &controllers.AdminController{},"get:Manage")
}
