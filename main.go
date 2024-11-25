package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/api"
	"github.com/kitamersion/kita-go-auth/domains/authentication"
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectedDb()
	initializers.MigrateDatabase()
}

func main() {
	r := gin.Default()
	r.Use(middleware.RateLimiter)

	v1 := r.Group("/v1")
	{
		v1.POST("/register", authentication.Register)
		v1.POST("/login", authentication.Login)
		v1.POST("/token/refresh", authentication.RefreshToken)

		// Inner group for user-specific actions
		user := v1.Group("/user", middleware.RequireAuth)
		{
			user.GET("/", api.User)
			user.POST("/activate", api.ActivateUser)     // TODO: middleware to only deactivate self or admin
			user.POST("/deactivate", api.DeactivateUser) // TODO: middleware to only activate self on login or admin
			user.DELETE("/delete", api.DeleteUser)       // TODO: middle ware to only delete self or admin

			user.POST("/role", api.AddUserRole)      // TODO: moddleware for only admins to modify
			user.DELETE("/role", api.RemoveUserRole) // TODO: moddleware for only admins to modify

			user.POST("/logout", authentication.Logout)
		}
	}

	r.Run()
}
