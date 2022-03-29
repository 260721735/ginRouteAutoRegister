# ginRouteAutoRegister
making gin route auto registration as simple as irismvc
##使用方式
```
import (
	"github.com/260721735/ginRouteAutoRegister/Gin-mvc"
)
func Route(engine *gin.Engine) {
	//原生路由注册方式
	engine.GET("/hello", controller.GetTest)
	//自动注册方式
	Gin-mvc.GinAutoRoute(engine, "/api", NewHelloController())
}

- controller层
  参照hello_controller.go
  -新建子location，controller层新增方法 GetLoation或PostLocation等
  -将controller注册到route/route.go中Route方法里
```
## controller

```
- controller层
  参照hello_controller.go
  -新建子location，controller层新增方法 GetLoation或PostLocation等
  -将controller注册到route/route.go中Route方法里
```
## router
```
-目前支持AnyOptions


