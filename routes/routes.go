package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/medical/abe"
	"github.com/medical/app/controllers"
	"github.com/medical/app/middlewares"
)

// SetupRouter sets up the application routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Setup authentication middleware
	authMiddleware := middlewares.AuthMiddleware()

	// Setup public routes
	public := r.Group("/public")
	{
		public.POST("/users", controllers.CreateUser)
		public.POST("/login", controllers.Login)
	}

	// Setup private routes
	private := r.Group("/api")
	private.Use(authMiddleware)
	{
		private.POST("/patients", controllers.CreatePatient)
		private.GET("/patients/:id", controllers.GetPatient)
		private.PUT("/patients/:id", controllers.UpdatePatient)
		private.DELETE("/patients/:id", controllers.DeletePatient)

		private.POST("/reports", controllers.CreateReport)
		private.GET("/reports/:id", controllers.GetReport)
		private.PUT("/reports/:id", controllers.UpdateReport)
		private.DELETE("/reports/:id", controllers.DeleteReport)
	}

	// Setup ABE key generation routes
	abeGroup := r.Group("/abe")
	abeGroup.Use(authMiddleware)
	{
		abeGroup.POST("/generate-keys", abe.GenerateKeys)
	}

	return r
}
