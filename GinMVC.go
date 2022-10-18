package ginRouteAutoRegister

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

//var methodsSync = struct {
//	sync.RWMutex
//	maps map[string]map[string]reflect.Value
//}{maps: make(map[string]map[string]reflect.Value)}
var methods = struct {
	sync.RWMutex
	maps map[string]map[string]map[string]reflect.Value
}{maps: make(map[string]map[string]map[string]reflect.Value)}

func GinAutoRoute(engine *gin.Engine, relativePath string, controller interface{}) {
	if strings.HasSuffix(relativePath, "/") {
		relativePath = relativePath[1 : len(relativePath)-1]
	}
	if !strings.HasPrefix(relativePath, "/") {
		return
	}
	relativePath = relativePath + "/:action"
	engine.Any(relativePath, ginAutoHand(controller))
}
func ginAutoHand(controller interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := strings.Title(strings.ToLower(c.Request.Method))
		realAction := strings.ToLower(c.Param("action"))
		controllerType := reflect.TypeOf(controller)
		pkgName := controllerType.String()
		//methods.RLock()
		_, hasMethod := methods.maps[method]
		//methods.RUnlock()
		if !hasMethod {
			//methods.Lock()
			methods.maps[method] = make(map[string]map[string]reflect.Value)
			//methods.Unlock()
		}
		//methods.RLock()
		value, hasRealAction := methods.maps[method][pkgName][realAction]
		//methods.RUnlock()
		if !hasRealAction {
			methodslen := len(methods.maps[pkgName])
			switch methodslen {
			case 0:
				controllerValue := reflect.ValueOf(controller)
				//methods.Lock()
				methods.maps[method][pkgName] = make(map[string]reflect.Value)
				//methods.Unlock()
				for i := 0; i < controllerValue.NumMethod(); i++ {
					subLocation := controllerType.Method(i).Name
					if strings.HasPrefix(subLocation, method) {
						//methods.Lock()
						if controllerValue.Method(i).Type().NumIn() == 1 && controllerValue.Method(i).Type().In(0) == reflect.TypeOf(c) {
							methods.maps[method][pkgName][strings.ToLower(strings.Replace(subLocation, method, "", 1))] = controllerValue.Method(i)
						}
						//methods.maps[method][pkgName][strings.ToLower(strings.Replace(subLocation, method, "", 1))] = controllerValue.Method(i)
						//methods.Unlock()
					}
				}
				//methods.RLock()
				v, hasRealAction := methods.maps[method][pkgName][realAction]
				//methods.RUnlock()
				value = v
				if !hasRealAction && (len(methods.maps[method]) > 0 || len(methods.maps[method][pkgName]) > 0) {
					http.NotFound(c.Writer, c.Request)
					c.Abort()
					return
				}
				break
			case 1:
				http.NotFound(c.Writer, c.Request)
				c.Abort()
				return
			}
		}
		value.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
