package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func (controller *HelloController) GetTestBAD1(c *gin.Context, a string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "GetTestbad1",
		"data": nil,
	})
}
func (controller *HelloController) GetTestBAD2() {

}
func (controller *HelloController) PostTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "PostTest",
		"data": nil,
	})
}
func (controller *HelloController) DeleteTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "DeleteTest",
		"data": nil,
	})
}
func (controller *HelloController) PutTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "PutTest",
		"data": nil,
	})
}

func (controller *HelloController) OptionsTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OptionsTest",
		"data": nil,
	})
}
func GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "GetTest",
		"data": nil,
	})
}
