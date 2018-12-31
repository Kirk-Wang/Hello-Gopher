package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/bootstrap"
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/web/middleware/identity"
	"github.com/Kirk-Wang/Hello-Gopher/iris-1/web/routes"
)

func main() {
	app := bootstrap.New("Superstar database", "一凡Sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8081")
}
