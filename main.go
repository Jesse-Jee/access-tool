package main

import (
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	Router(router)

	if err := router.Run(":9080"); err != nil {
		_ = log4go.Error(err)
	}

}
