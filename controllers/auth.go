package controllers

import (
	"OpenPBL/util"
	"fmt"
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
	auth.InitConfig(CasdoorEndpoint, ClientId, ClientSecret, JwtSecret, CasdoorOrganization)
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
// @Title
// @Description
// @Param code body string true	"code"
// @Param state body string true	"state"
// @Success 200 {int} models.Teacher.Id
// @Failure 403 body is empty
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

	fmt.Println(claims.StandardClaims)

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
// @Title
// @Description
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
// @Title
// @Description
// @Success 200 {object} Response
// @Failure 401 {object} Response
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
