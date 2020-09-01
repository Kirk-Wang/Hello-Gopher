package routes

import (
	"github.com/Kirk-Wang/Hello-Gopher/lottery/wheel/services"
)

func Configure(b *bootstrap.Bootstrapper) {
		userService := services.NewUserService()
		giftService := services.NewGiftService()
		codeService := services.NewCodeService()
		resultService := services.NewResultService()
		userdayService := services.NewUserdayService()
		blackipService := services.NewBlackipService()

		index := mvc.New(b.Party("/"))
		index.Register(
			userService,
			codeService,
			resultService,
			userdayService,
			blackipService
		)
		index.Handle(new(controllers.IndexController))
}