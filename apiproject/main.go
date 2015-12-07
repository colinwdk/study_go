package main

import (
	_ "apiproject/docs"
	_ "apiproject/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	
	beego.Run()
}
