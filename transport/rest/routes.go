package rest

import (
	"discount-service/transport/controllers"
	"github.com/labstack/echo"
)

func (ac *RContext) RegisterRoutes(e *echo.Echo) {
	cc := controllers.ControllerContext{ApplicationContext: &ac.ApplicationContext}
	v1 := e.Group("/v1")

	v1.POST("/discount/rules", cc.CreateDiscountRule)
	v1.GET("/discount/rules/:account_id", cc.GetDiscountRule)

	v1.GET("/hello", cc.Hello, ac.Allow(cc.Hello, ""))
}
