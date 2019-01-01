package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/bootstrap"
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/web/middleware/identity"
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Superstar database", "Kirk")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
