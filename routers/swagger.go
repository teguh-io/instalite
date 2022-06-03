package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type swaggerRouter struct {
	router *gin.Engine
}

func NewSwaggerRouter(r *gin.Engine) *swaggerRouter {
	return &swaggerRouter{
		router: r,
	}

}

func (sr *swaggerRouter) Start() {
	sr.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
