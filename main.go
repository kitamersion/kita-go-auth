package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kitamersion/kita-go-auth/api"
	"github.com/kitamersion/kita-go-auth/domains/authentication"
	"github.com/kitamersion/kita-go-auth/events"
	"github.com/kitamersion/kita-go-auth/events/handlers"
	"github.com/kitamersion/kita-go-auth/initializers"
	"github.com/kitamersion/kita-go-auth/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectedDb()
	initializers.MigrateDatabase()
	initializers.SeedPermissionData(initializers.DB)
}

func main() {
	events.InitalizeEventBus()

	// TODO: remove somewhere else
	// register event handlers
	events.EventBusGo.Subscribe(events.RoleAssigned, handlers.RoleAssignedHandler{})
	events.EventBusGo.Subscribe(events.RoleRevoked, handlers.RoleRevokedHandler{})

	r := gin.Default()

	r.Use(middleware.CORS)
	r.Use(middleware.RateLimiter)

	v1 := r.Group("/v1")
	{
		// No Auth
		v1.POST("/register", authentication.Register)
		v1.POST("/login", authentication.Login)
		v1.POST("/token/refresh", authentication.RefreshToken)

		// Auth
		v1.GET("/whoami", middleware.RequireAuth, api.WhoAmI)
		v1.POST("/logout", middleware.RequireAuth, authentication.Logout)

		// Auth + protected by permissions
		user := v1.Group("/user")
		user.Use(middleware.RequireAuth, middleware.CanViewEditSelf)
		{
			user.GET("/:id", api.User)
			user.PUT("/:id/activate", api.ActivateUser)
			user.PUT("/:id/deactivate", api.DeactivateUser)
			user.DELETE("/:id/delete", api.DeleteUser)

			// Admin-restricted role management
			roles := user.Group("/:id/role")
			roles.Use(middleware.AdminOnly)
			{
				roles.PUT("/", api.AddUserRole)
				roles.DELETE("/", api.RemoveUserRole)
			}
		}
	}

	r.Run()
}
