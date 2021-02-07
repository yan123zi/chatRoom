package app

import (
	"chatRoom/server/config"
	"chatRoom/server/controllers/api"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func InitIris() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	app.AllowMethods(iris.MethodOptions)

	app.Any("/", func(ctx *context.Context) {
		ctx.HTML("<h1>chatRoom is powered!</h1>")
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/login").Handle(new(api.LoginController))
		m.Party("/captcha").Handle(new(api.CaptchaController))
	})

	app.Listen(":" + config.AppConfig.Port)
}
