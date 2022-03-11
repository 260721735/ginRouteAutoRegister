# ginRouteAutoRegister
making gin route auto registration as simple as irismvc
## controller

```
- controller层
  参照hello_controller.go
  -新建子location，controller层新增方法 GetLoation或PostLocation等
  -将controller注册到route/route.go中Route方法里
```
## router
```
-目前仅添加get，post，put，delete，options五个方法
-新增方式
 -在route/autoRoute.go下AutoRoute方法中新增engine.(所需options)(relativePath, AutoHand(controller))


