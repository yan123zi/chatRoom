package api

import (
	"chatRoom/server/common"
	"github.com/dchest/captcha"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"strings"
)

type CaptchaController struct {
	Ctx iris.Context
}

func (c *CaptchaController) GetRequest() *common.JsonResult {
	captchaId := captcha.NewLen(5)
	uuid,_:=uuid.NewUUID()
	captchaUrl := "/api/captcha/show?captchaId=" + captchaId+"&r="+strings.ReplaceAll(uuid.String(),"-","")
	return common.NewEmptyRespBuild().
		Put("captchaId", captchaId).
		Put("captchaUrl", captchaUrl).JsonResult()
}

func (c *CaptchaController) GetShow() {
	captchaId := c.Ctx.URLParam("captchaId")

	if !captcha.Reload(captchaId) {
		c.Ctx.StatusCode(404)
		return
	}

	c.Ctx.Header("Content-Type", "image/png")

	if err := captcha.WriteImage(c.Ctx.ResponseWriter(), captchaId, 200, 100); err != nil {
		logrus.Errorf("show code image err:%s", err.Error())
		return
	}
}
