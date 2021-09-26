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

package controllers

import (
	"OpenPBL/util"
	"github.com/astaxie/beego"
	"github.com/casdoor/casdoor-go-sdk/auth"
)

type AuthController struct {
	beego.Controller
}

func InitCasdoor() {
	var CasdoorEndpoint = beego.AppConfig.String("casdoorEndpoint")
	var ClientId = beego.AppConfig.String("clientId")
	var ClientSecret = beego.AppConfig.String("clientSecret")
	var JwtSecret = beego.AppConfig.String("jwtSecret")
	var CasdoorOrganization = beego.AppConfig.String("casdoorOrganization")
	var ApplicationName = beego.AppConfig.String("applicationName")
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtSecret, CasdoorOrganization, ApplicationName)
}


func (c *AuthController) GetSessionUser() *auth.Claims {
	s := c.GetSession("user")
	if s == nil {
		return nil
	}
	claims := &auth.Claims{}
	err := util.JsonToStruct(s.(string), claims)
	if err != nil {
		panic(err)
	}

	return claims
}

func (c *AuthController) SetSessionUser(claims *auth.Claims) {
	if claims == nil {
		c.DelSession("user")
		return
	}

	s := util.StructToJson(claims)

	c.SetSession("user", s)
}

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// Login
// @Title Login
// @Description User login
// @Param code body string true	"code"
// @Param state body string true "state"
// @Success 200 {object} Response
// @router /login [post]
func (c *AuthController) Login() {
	code := c.Input().Get("code")
	state := c.Input().Get("state")
	token, err := auth.GetOAuthToken(code, state)
	if err != nil {
		c.Data["json"] = Response{
			Code: 403,
			Msg:  err.Error(),
		}
		c.ServeJSON()
		return
	}
	claims, err := auth.ParseJwtToken(token.AccessToken)
	if err != nil {
		c.Data["json"] = Response{
			Code: 403,
			Msg:  err.Error(),
		}
		c.ServeJSON()
		return
	}

	claims.AccessToken = token.AccessToken
	c.SetSessionUser(claims)

	resp := &Response{
		Code: 200,
		Msg: "登录成功",
		Data: claims,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// Logout
// @Title Logout
// @Description User logout
// @Success 200 {object} Response
// @router /logout [post]
func (c *AuthController) Logout() {
	var resp Response
	c.SetSessionUser(nil)
	resp = Response{
		Code: 200,
		Msg: "",
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// GetAccount
// @Title GetAccount
// @Description Get user account
// @Success 200 {object} Response
// @router /account [get]
func (c *AuthController) GetAccount() {
	var resp Response
	user := c.GetSessionUser()
	if user == nil {
		resp = Response{
			Code: 404,
			Msg: "账号不存在",
		}
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	resp = Response{
		Code: 200,
		Msg: "",
		Data: user,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
