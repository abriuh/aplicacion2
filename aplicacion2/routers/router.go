// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/abriuh/aplicacion2/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/peliculas",
			beego.NSInclude(
				&controllers.PeliculasController{},
			),
		),

		beego.NSNamespace("/actores",
			beego.NSInclude(
				&controllers.ActoresController{},
			),
		),

		beego.NSNamespace("/cliente",
			beego.NSInclude(
				&controllers.ClienteController{},
			),
		),

		beego.NSNamespace("/factura",
			beego.NSInclude(
				&controllers.FacturaController{},
			),
		),

		beego.NSNamespace("/discos",
			beego.NSInclude(
				&controllers.DiscosController{},
			),
		),

		beego.NSNamespace("/contreras",
			beego.NSInclude(
				&controllers.ContrerasController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
