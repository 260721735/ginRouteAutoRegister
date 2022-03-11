package route

import (
	"ginRouteAutoRegister/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
func Route(engine *gin.Engine) {
	engine.Use(Options)
	engine.Use(Secure)
	//原生路由注册方式
	engine.GET("/hello", controller.GetTest)
	//自动注册方式
	AutoRoute(engine, "/api", controller.NewHelloController())
	AutoRoute(engine, "/api2", controller.NewHelloController())
	//兼容原先gin拦截方式，自己写好LoginRequired，在LoginRequired之后的AutoRoute都会被拦截
	//engine.Use(LoginRequired)
	//AutoRoute(engine, "/filter", controller.NewHelloController())

}
