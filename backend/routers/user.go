package routers

import (
    "github.com/gin-gonic/gin"
)

// RegisterUserRoutes defines user-related routes
func RegisterUserRoutes(router *gin.Engine) {
    userGroup := router.Group("/user")
    {
        userGroup.GET("/info", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "User Info",
            })
        })

        userGroup.POST("/create", func(c *gin.Context) {
            c.JSON(201, gin.H{
                "message": "User Created",
            })
        })

        userGroup.PUT("/update", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "User Updated",
            })
        })

        userGroup.DELETE("/delete", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "message": "User Deleted",
            })
        })
    }
}