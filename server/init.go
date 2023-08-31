package main

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.POST("/createSegment", CreateSegment)
	r.POST("/deleteSegment", DeleteSegment)
	r.POST("/addUserToSegment", AddUserToSegment)
	r.GET("/getUserSegments", GetUserSegments)
	r.GET("/getAllUsers", GetAllUsers)
}
