// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	/*ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		
		beego.NSNamespace("/worker",
			//beego.NSNamespace("/querybagreq",
				beego.NSInclude(
					&controllers.WorkerController{},
				),
			//),
		),
	)
	beego.AddNamespace(ns)*/

	
	beego.Router("/v1/user", &controllers.UserController{})
	beego.Router("/v1/worker/querybagreq/:workerid/:token", &controllers.WorkerController{}, "get:Querybagreq")
	beego.Router("/v1/worker/getbagreq/:workerid/:token", &controllers.WorkerController{}, "get:Getbagreq")
	beego.Router("/v1/worker/getbagreqlist/:workerid/:token", &controllers.WorkerController{}, "get:GetbagreqList")
	beego.Router("/v1/worker/getworkerid", &controllers.WorkerController{}, "get:GetWorkerId")
	
	
	beego.Router("/v1/worker/postworkerinfo", &controllers.WorkerController{}, "*:PostWorkerInfo")
	beego.Router("/v1/worker/postworkerinfo1/:workerid/:token", &controllers.WorkerController{}, "post:PostWorkerInfo1")
	
	beego.InsertFilter("/v1/worker/getworkerid", beego.BeforeRouter, FilterUser)
}


	var FilterUser = func(ctx *context.Context) {
	    /*_, ok := ctx.Input.Session("uid").(int)
	    if !ok && ctx.Request.RequestURI != "/login" {
	        ctx.Redirect(302, "/login")
	    }*/
		url := ctx.Input.Url()
		workerId := ctx.Request.URL.Query().Get("workerid")
		token := ctx.Request.URL.Query().Get("token")
		
		beego.Debug("1234567890:"+ url, workerId, token)
}

