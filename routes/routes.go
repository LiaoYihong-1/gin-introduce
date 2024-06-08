package routes

import (
	"gin/controller"
	"gin/service"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claims, _ := utils.ParseToken(token)
		privileges, err := service.GetUserPrivilegesInRedis(claims.Id)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("privileges", privileges)
		c.Next()
	}
}
func checkPrivilege(privilege utils.PrivilegeType) gin.HandlerFunc {
	return func(c *gin.Context) {
		privilegesInterface, err := c.Get("privileges")
		if !err {
			c.Status(http.StatusInternalServerError)
		}
		privileges, e := privilegesInterface.([]utils.PrivilegeType)
		if !e {
			c.Status(http.StatusInternalServerError)
		}
		hasPrivilege := utils.HasPrivilege(privilege, privileges)
		if !hasPrivilege {
			c.Status(http.StatusForbidden)
			c.Abort()
			return
		}
		c.Next()
	}
}
func controlFlow() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 1)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			c.Status(http.StatusTooManyRequests)
			c.Abort()
			return
		}
	}
}
func Routes(ginServer *gin.Engine) {
	ginServer.Use(controlFlow())
	ginServer.GET("/users/:id", TokenAuth(), checkPrivilege(utils.GET), controller.GetUserById)
	ginServer.GET("/users", TokenAuth(), checkPrivilege(utils.GET), controller.GetAllUsers)
	ginServer.GET("/users/login", controller.Login)
	ginServer.POST("/users", TokenAuth(), checkPrivilege(utils.CREATE), controller.CreateUser)
	ginServer.POST("/users/register", controller.Register)
	ginServer.DELETE("/users/:id", TokenAuth(), checkPrivilege(utils.DELETE), controller.DeleteUser)
	ginServer.PUT("/users/:id", TokenAuth(), checkPrivilege(utils.UPDATE), controller.UpdateUser)
}
