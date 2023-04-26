package main1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main2() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic:", err)
			}
		}()
	})
	r.GET("/test", A)
	r.GET("/testNotRecover", B)
	r.GET("/testRecover", D)
	r.Run(":8080")
}

func A(c *gin.Context) {
	c.String(200, "Test request handled by A()")
}

func D(c *gin.Context) {
	panic("Panic message from C()")
	c.String(200, "Test recover success by D")
}

func B(c *gin.Context) {
	go func() {
		C()
	}()
	c.String(200, "Test request Panic by C()")
}

func C() {
	panic("Panic message from C()")
}
