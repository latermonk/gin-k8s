package main

import (
	"github.com/gin-gonic/gin"
)
func main() {

	r := gin.Default()
	//r.GET("/", xxx)
	//r.NoRoute("")



	r.Run(":9999")

}
