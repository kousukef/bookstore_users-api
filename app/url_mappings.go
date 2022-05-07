package app

import (
	"github.com/kousukef/bookstore_users-api/controllers/ping"
	"github.com/kousukef/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/search", users.Search)
	router.GET("/users/:user_id", users.Get)
	router.POST("/users/login", users.Login)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
