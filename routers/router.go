// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"bee_http_game/controllers"
	"github.com/beego/beego/v2/server/web/filter/cors"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 允许跨域请求
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/compensation_items",
			beego.NSInclude(
				&controllers.CompensationItemsController{},
			),
		),

		beego.NSNamespace("/compensations",
			beego.NSInclude(
				&controllers.CompensationsController{},
			),
		),

		beego.NSNamespace("/server",
			beego.NSInclude(
				&controllers.ServerController{},
			),
		),

		beego.NSNamespace("/server_config",
			beego.NSInclude(
				&controllers.ServerConfigController{},
			),
		),

		beego.NSNamespace("/user_account",
			beego.NSInclude(
				&controllers.UserAccountController{},
			),
		),

		beego.NSNamespace("/user_server",
			beego.NSInclude(
				&controllers.UserServerController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.SetStaticPath("/swagger", "swagger")
	beego.SetStaticPath("/client", "client")
}
