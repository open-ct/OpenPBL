// Copyright 2021 The OpenPBL Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
