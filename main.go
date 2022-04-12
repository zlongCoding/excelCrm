package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zlongCoding/excelCrm/controller"
)

func main() {
	r := gin.Default()
	// r.Get("/ping", )
	r.GET("/ping", controller.Ping)
	err := r.Run()
	if err != nil {
		panic(err)
	}

}
