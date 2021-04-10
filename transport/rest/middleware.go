package rest

import (
	"discount-service/core/auth"
	"discount-service/logger"
	"discount-service/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)



func (ac *RContext) RegisterMiddlewares(e *echo.Echo) {

	// TODO change logger to custom logger
	e.Use(middleware.Logger())
	// TODO change recover to custom logger
	e.Use(middleware.Recover())
	t := logger.LoggerUtils{SugaredLogger: ac.Logger}
	e.HTTPErrorHandler = t.CustomZapHttpErrorHandler

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

// Allow to check authentication of routes
func (ac *RContext) Allow(h echo.HandlerFunc, roles ...string) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var authContext = auth.AuthenticateContext{ApplicationContext: &ac.ApplicationContext}
			identity, err := authContext.Authenticate(&c, "", roles...)
			if err != nil {
				return utils.StandardHttpErrorResponse{Message: err.Error()}
			}
			c.Set("identity", identity)
			return h(c)
		}
	}
}
