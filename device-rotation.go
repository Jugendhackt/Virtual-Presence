package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/:alpha/:beta/:gamma", func(c *gin.Context) {
		a := parse(c, "alpha")
		b := parse(c, "beta")
		g := parse(c, "gamma")
		fmt.Println(a, b, g)
	})
	r.Run(":3000")
}

func parse(c *gin.Context, v string) float64 {
	alpha, err := strconv.ParseFloat(c.Param(v), 64)
	if err != nil {
		c.JSON(500, gin.H{"message": "Invalid parameter!"})
		return 0.0
	}
	return alpha
}
