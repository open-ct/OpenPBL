package routers

import (
	"OpenPBL/controllers"
	"github.com/astaxie/beego/context"
	"net/http"
	"os"
	"strings"
)

func apiFilter(ctx *context.Context) {
	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/api/project-list") ||
		strings.HasPrefix(urlPath, "/api/project") ||
		strings.HasPrefix(urlPath, "/api/message") ||
		strings.HasPrefix(urlPath, "/api/student") {
		user := ctx.Input.Session("user")
		if user == nil {
			ctx.Output.JSON(controllers.Response{
				Code: 401,
				Msg:  "请先登录",
			}, true, false)
		}
	}
}

func TransparentStatic(ctx *context.Context) {
	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/api") {
		apiFilter(ctx)
		return
	}

	path := "web/build"
	if urlPath == "/" {
		path += "/index.html"
	} else {
		path += urlPath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, "web/build/index.html")
	} else {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	}
}
