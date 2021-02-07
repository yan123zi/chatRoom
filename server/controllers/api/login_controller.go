package api

import (
	"chatRoom/server/common"
	"chatRoom/server/services"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) PostSignup() *common.JsonResult {
	var (
		userName    = c.Ctx.PostValueTrim("username")
		passWord    = c.Ctx.PostValueTrim("password")
		rePassword  = c.Ctx.PostValueTrim("rePassword")
		captchaId   = c.Ctx.PostValueTrim("captchaId")
		captchaCode = c.Ctx.PostValueTrim("captchaCode")
	)
	if !captcha.VerifyString(captchaId, captchaCode) {
		return common.JsonError(common.CaptchaErr)
	}
	user, err := services.LoginService.SignUp(userName, passWord, rePassword)
	if err != nil {
		return common.JsonErrorMsg(err.Error())
	}
	return common.NewEmptyRespBuild().Put("user", user).JsonResult()
}
