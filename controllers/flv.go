package controllers

import (
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

// CMSFLV API
type FlvController struct {
	beego.Controller
}

// @Title getFlv
// @Description get admin page
// @Param xxl_sso_token query string false "The tokenText"
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router /flvlist [get]
func (c *FlvController) GetFlvList() {
	var token string
	token = c.Ctx.GetCookie("token")
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	if token != "" {
		// beego.Info(c.Ctx.Request.URL.String())
		// beego.Info(c.Ctx.Input.URI())
		// beego.Info(c.Ctx.Request.RequestURI)
		// beego.Info(strings.Split(c.Ctx.Request.URL.String(), "?")[0])
		urlarray := strings.Split(c.Ctx.Request.URL.String(), "?")
		if len(urlarray) > 1 {
			c.Redirect(strings.Split(c.Ctx.Request.URL.String(), "?")[0], 302)
		} else {
			userid, username, usernickname, err := utils.LubanCheckToken(token)
			if err != nil {
				beego.Error(err)
			}
			c.Ctx.SetCookie("token", token, "3600", "/")
			c.SetSession("uname", username)
			c.SetSession("userid", userid)
			c.SetSession("usernickname", usernickname)
		}
	} else {
		token = c.Input().Get("xxl_sso_token")
		if token == "" {
			c.Redirect("https://www.54lby.com/sso/login?redirect_url="+site+c.Ctx.Request.URL.String(), 302)
		} else {
			userid, username, usernickname, err := utils.LubanCheckToken(token)
			if err != nil {
				beego.Error(err)
			}
			c.Ctx.SetCookie("token", token, "3600", "/")
			c.SetSession("uname", username)
			c.SetSession("userid", userid)
			c.SetSession("usernickname", usernickname)
			urlarray := strings.Split(c.Ctx.Request.URL.String(), "?")
			if len(urlarray) > 1 {
				c.Redirect(strings.Split(c.Ctx.Request.URL.String(), "?")[0], 302)
			}
		}
	}

	c.TplName = "flvlist.tpl"
}

// @Title getFlv
// @Description get admin page
// @Param filepath query string true "The mp4 file path"
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router / [get]
func (c *FlvController) Get() {
	mp4link := c.Input().Get("filepath")
	c.Data["Mp4Link"] = mp4link
	c.TplName = "flv.tpl"
}
