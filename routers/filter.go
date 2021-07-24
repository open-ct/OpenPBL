package routers

import (
	"github.com/astaxie/beego/context"
	"net/http"
	"os"
	"strings"
)

func TransparentStatic(ctx *context.Context) {
	urlPath := ctx.Request.URL.Path
	if strings.HasPrefix(urlPath, "/api") {
		return
	}

	path := "openpbl-landing/build"
	if urlPath == "/" {
		path += "/index.html"
	} else {
		path += urlPath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, "openpbl-landing/build/index.html")
	} else {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, path)
	}
}
