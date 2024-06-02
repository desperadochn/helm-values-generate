package router

import (
	"github.com/gin-gonic/gin"
	"helm-values-generate/pkg/controller"
)

func InitRouter() {
	router := gin.Default()
	api := router.Group("api")
	{
		helm := api.Group("helm")
		{
			helm.POST("CreateProject", controller.CreateHelmProject)
		}

	}
}
