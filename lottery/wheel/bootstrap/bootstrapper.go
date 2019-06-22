package	bootstrap

import (
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
		AppSpawDate: time.Now()
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