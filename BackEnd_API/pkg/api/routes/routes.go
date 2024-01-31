package routes

import (
	"github.com/gin-gonic/gin"
	"test/pkg/api/handlers"
)

func SetupRoutes(router *gin.Engine) {

	//TODO AUTHORIZATION MIDDLEWARE BEARER FOR ( INVITE,DELTE_ORG ,UPDATE_ORG,READ_ALLORG,READ_ONE_ORG,CREATE_ORG
	// Users Route
	router.GET("/users", handlers.GetUsersHandler)
	router.GET("/user/:id", handlers.GetUserHandler)
	router.DELETE("/user/:id", handlers.DeleteUserHandler)
	router.POST("/user/signup", handlers.AddUserHandler)
	router.POST("/user/signin", handlers.LoginUserHandler)
	router.PUT("/user/:id", handlers.UpdateUserHandler)

	// Organizations Route
	router.GET("/organizations", handlers.GetOrganizationsHandler)
	router.GET("/organization/:id", handlers.GetOrganizationHandler)
	router.POST("/organization", handlers.AddOrganizationHandler)
	router.POST("/organization/:organization_id/invite", handlers.InviteOrganizationHandler)
	router.PUT("/organization/:organization_id", handlers.UpdateOrganizationHandler)
	router.DELETE("/organization/:organization_id", handlers.DeleteOrganizationHandler)
	//Test
	router.GET("/test", handlers.TestHandler)
}
