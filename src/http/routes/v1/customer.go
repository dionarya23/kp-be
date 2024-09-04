package v1routes

import (
	customerv1controller "github.com/dionarya23/kredit-plus/src/http/controllers/customer"
)

func (i *V1Routes) MountCustomer() {
	g := i.Echo.Group("/customer")

	customerController := customerv1controller.New(&customerv1controller.V1Customer{
		DB: i.DB,
	})

	g.POST("/register", customerController.Register)
	g.POST("/login", customerController.Login)
}
