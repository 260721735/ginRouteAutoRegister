# ginRouteAutoRegister
making gin route auto registration as simple as irismvc
##使用方式
```
import (
	"github.com/260721735/ginRouteAutoRegister"
)
func Route(engine *gin.Engine) {
	//原生路由注册方式
	engine.GET("/hello", controller.GetTest)
	//自动注册方式
	ginRouteAutoRegister.GinAutoRoute(engine, "/api", NewHelloController())
}
```
##Demo
###将HelloController.GetTest注册到http://localhost:8888/api/test路径上
```
import (
	"github.com/260721735/ginRouteAutoRegister"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())
	Route(router)
	srv := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//自定义退出时prestop逻辑，这里简单用等待5s
	time.Sleep(1 * time.Second)
	log.Println("Server exiting")
}
func Route(engine *gin.Engine) {
	ginRouteAutoRegister.GinAutoRoute(engine, "/api", NewHelloController())
}

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}
func (controller *HelloController) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "GetTest",
		"data": nil,
	})
}
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


