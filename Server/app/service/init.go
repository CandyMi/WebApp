package service

import (
	"app/service/controllers"
	"app/service/models"
	"app/service/views"
	"github.com/gin-gonic/gin"
)

func Init() {
	models.InitModels()
	views.InitViews()
}

func Run() {
	controllers.InitRoute(gin.Default())
}
