package routers

import (
	"zeus/controllers/admin"
	"zeus/middleware"

	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/"+beego.AppConfig.String("adminPath"),
			//中间件:匹配路由前会执,可以用于权限验证
			//注意引入的包： github.com/astaxie/beego/context
			beego.NSBefore(middleware.AdminAuth),

			//后台管理
			beego.NSRouter("/", &admin.MainController{}),
			beego.NSRouter("/home", &admin.HomeController{}, "get:Home"),


			//--------资产
			beego.NSRouter("/asset/host", &admin.AssetController{}, "get:Host"),
			beego.NSRouter("/asset/hostAdd", &admin.AssetController{}, "post:HostAdd"),
			beego.NSRouter("/asset/hostEdit", &admin.AssetController{}, "post:HostEdit"),

			beego.NSRouter("/asset/phost", &admin.AssetController{}, "get:Phost"),
			beego.NSRouter("/asset/phostAdd", &admin.AssetController{}, "post:PhostAdd"),
			beego.NSRouter("/asset/phostEdit", &admin.AssetController{}, "post:PhostEdit"),

			beego.NSRouter("/asset/vhost", &admin.AssetController{}, "get:Vhost"),
			beego.NSRouter("/asset/vhostAdd", &admin.AssetController{}, "post:VhostAdd"),
			beego.NSRouter("/asset/vhostEdit", &admin.AssetController{}, "post:VhostEdit"),

			beego.NSRouter("/asset/networkDevice", &admin.AssetController{}, "get:NetworkDevice"),
			beego.NSRouter("/asset/networkDeviceAdd", &admin.AssetController{}, "post:NetworkDeviceAdd"),
			beego.NSRouter("/asset/networkDeviceEdit", &admin.AssetController{}, "post:NetworkDeviceEdit"),

			beego.NSRouter("/asset/domainName", &admin.AssetController{}, "get:DomainName"),
			beego.NSRouter("/asset/domainNameAdd", &admin.AssetController{}, "post:DomainNameAdd"),
			beego.NSRouter("/asset/domainNameEdit", &admin.AssetController{}, "post:DomainNameEdit"),

			beego.NSRouter("/asset/certificate", &admin.AssetController{}, "get:Certificate"),
			beego.NSRouter("/asset/certificateAdd", &admin.AssetController{}, "post:CertificateAdd"),
			beego.NSRouter("/asset/certificateEdit", &admin.AssetController{}, "post:CertificateEdit"),


			beego.NSRouter("/main/changeStatus", &admin.MainController{}, "get:ChangeStatus"),
			beego.NSRouter("/main/editNum", &admin.MainController{}, "get:EditNum"),

			//-------------系统设置
			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/login/doLogin", &admin.LoginController{}, "post:DoLogin"),
			beego.NSRouter("/login/loginOut", &admin.LoginController{}, "get:LoginOut"),
			//角色管理
			beego.NSRouter("/role", &admin.RoleController{}),
			beego.NSRouter("/role/add", &admin.RoleController{}, `get:Add`),
			beego.NSRouter("/role/edit", &admin.RoleController{}, `get:Edit`),
			beego.NSRouter("/role/doAdd", &admin.RoleController{}, `post:DoAdd`),
			beego.NSRouter("/role/doEdit", &admin.RoleController{}, `post:DoEdit`),
			beego.NSRouter("/role/delete", &admin.RoleController{}, `get:Delete`),
			beego.NSRouter("/role/auth", &admin.RoleController{}, `get:Auth`),
			beego.NSRouter("/role/doAuth", &admin.RoleController{}, `post:DoAuth`),
			//管理员管理
			beego.NSRouter("/manager", &admin.ManagerController{}),
			beego.NSRouter("/manager/add", &admin.ManagerController{}, "get:Add"),
			beego.NSRouter("/manager/edit", &admin.ManagerController{}, "get:Edit"),
			beego.NSRouter("/manager/doAdd", &admin.ManagerController{}, `post:DoAdd`),
			beego.NSRouter("/manager/doEdit", &admin.ManagerController{}, `post:DoEdit`),
			beego.NSRouter("/manager/delete", &admin.ManagerController{}, `get:Delete`),
			//权限管理
			beego.NSRouter("/access", &admin.AccessController{}),
			beego.NSRouter("/access/add", &admin.AccessController{}, "get:Add"),
			beego.NSRouter("/access/edit", &admin.AccessController{}, "get:Edit"),
			beego.NSRouter("/access/doAdd", &admin.AccessController{}, `post:DoAdd`),
			beego.NSRouter("/access/doEdit", &admin.AccessController{}, `post:DoEdit`),
			beego.NSRouter("/access/delete", &admin.AccessController{}, `get:Delete`),

		)
	//注册 namespace
	beego.AddNamespace(ns)
}
