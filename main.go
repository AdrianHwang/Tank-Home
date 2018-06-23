package main

import (
	_ "Tank-Home/routers"
	"github.com/astaxie/beego"
	"net/http"
	"strings"
	"github.com/astaxie/beego/context"
)

func main() {
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath(){
	//过滤静态目录
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context){
	// 测试
	orPath:=ctx.Request.URL.Path
	beego.Debug("request url:", orPath)

	if strings.Index(orPath, "api")>=0{
		return
	}

	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+orPath)
}