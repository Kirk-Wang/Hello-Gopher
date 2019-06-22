package	bootstrap

import (
	"time"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris"
)

type Configurator func(bootstrapper *bootstrap.Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName string
	AppOwner string
	AppSpawDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper{
	b := &Bootstrapper{
		Application: iris.New(),
		AppName: appName,
		AppOwner: appOwner,
		AppSpawDate: time.Now(),
	}

	for _, cfg  := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true) // 开发模式
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string{
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string{
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app": b.AppName,
			"status": ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) setupCron() {
	// TODO:
}

const (
	StaticAssets = "./public"
	Favicon = "favicon.ico"
)

func (b *Bootstrapper) Bootstrap() *Bootstrapper{
	b.SetupViews("./views")
	b.SetupErrorHandlers()

	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	b.setupCron()

	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}