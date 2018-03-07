package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ActoresController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ClienteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:ContrerasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:DiscosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:FacturaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"] = append(beego.GlobalControllerRouter["github.com/abriuh/aplicacion2/controllers:PeliculasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
