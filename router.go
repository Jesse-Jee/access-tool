package main

import (
	"access-tool/api"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine){
	router.GET("/team", api.NewTeam().Get)
}