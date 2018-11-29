package routers

import (
	"github.com/astaxie/beego"
	"CmsProject/controllers"
)

/**
 * 路由
 */
func init() {

	beego.Router("/", &controllers.MainController{})

	//============================管理员相关操作==========================
	beego.Router("/admin/login", &controllers.AdminController{}, "POST:AdminLogin")             //登陆
	beego.Router("/admin/info", &controllers.AdminController{}, "GET:GetAdminInfo")             //获取管理员信息
	beego.Router("/admin/singout", &controllers.AdminController{}, "GET:SignOut")               //管理员退出
	beego.Router("/admin/count", &controllers.AdminController{}, "GET:GetAdminCount")           //获取管理员总数
	beego.Router("/statis/admin/*/count", &controllers.AdminController{}, "GET:GetAdminStatis") ////获取某一日的管理员增长统计数据
	beego.Router("/admin/all", &controllers.AdminController{}, "GET:GetAdminList")              //查询整所有的用户列表

	//============================用户模块相关操作========================
	beego.Router("/statis/user/*/count", &controllers.UserController{}, "GET:UserStatisDaily") //获取某个日期的用户的增长数据
	beego.Router("/v1/users/count", &controllers.UserController{}, "GET:GetUserCount")
	beego.Router("/v1/users/list", &controllers.UserController{}, "GET:UserList")

	//============================订单模块相关操作========================
	beego.Router("/statis/order/*/count", &controllers.OrderController{}, "GET:GetOrderStatis")
	beego.Router("/bos/orders/count", &controllers.OrderController{}, "GET:GetOrderCount")

	//==================文件相关操作==================//未测试
	//beego.Router("/v1/addimg/:username", &controllers.FileController{}, "POST:UploadImg")
	beego.Router("/admin/update/avatar/:adminId", &controllers.FileController{}, "POST:UpdateAdminAvatar")
}
