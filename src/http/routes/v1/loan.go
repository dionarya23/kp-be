package v1routes

import (
	loanv1controller "github.com/dionarya23/kredit-plus/src/http/controllers/loan"
	"github.com/dionarya23/kredit-plus/src/http/middlewares"
)

func (i *V1Routes) MountLoan() {
	g := i.Echo.Group("/loan")

	customerController := loanv1controller.New(&loanv1controller.V1Loan{
		DB: i.DB,
	})

	g.POST("/", customerController.Create, middlewares.Authentication())
	g.GET("/", customerController.FindMany, middlewares.Authentication())
	g.GET("/:id", customerController.FindOne, middlewares.Authentication())
}
