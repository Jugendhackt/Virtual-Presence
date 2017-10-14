package main

import (
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	vals := []float64{0.0, 0.0}

	r := gin.Default()

	r.GET("/w/:u/:alpha", func(c *gin.Context) {
		i := parseInt(c, "u")
		vals[i] = parse(c, "alpha")
		c.JSON(200, vals[i])
	})

	r.GET("/r/:u", func(c *gin.Context) {
		c.JSON(200, vals[parseInt(c, "u")])
	})

	r.GET("/calc", func(c *gin.Context) {
		delta := vals[0] - vals[1]
		var led string
		if delta < -90 {
			led = "B"
		} else if delta < -80 && delta >= -90 {
			led = "C"
		} else if delta < -70 && delta >= -80 {
			led = "D"
		} else if delta < -60 && delta >= -70 {
			led = "E"
		} else if delta < -50 && delta >= -60 {
			led = "F"
		} else if delta < -40 && delta >= -50 {
			led = "G"
		} else if delta < -30 && delta >= -40 {
			led = "H"
		} else if delta < -20 && delta >= -30 {
			led = "I"
		} else if delta < -10 && delta >= -20 {
			led = "J"
		} else if delta < -5 && delta >= -10 {
			led = "K"
		} else if delta >= -5 && delta < 5 {
			led = "L"
		} else if delta >= 5 && delta <= 10 {
			led = "M"
		} else if delta > 10 && delta <= 20 {
			led = "N"
		} else if delta > 20 && delta <= 30 {
			led = "O"
		} else if delta > 30 && delta <= 40 {
			led = "P"
		} else if delta > 40 && delta <= 50 {
			led = "Q"
		} else if delta > 50 && delta <= 60 {
			led = "R"
		} else if delta > 60 && delta <= 70 {
			led = "S"
		} else if delta > 70 && delta <= 80 {
			led = "T"
		} else if delta > 80 && delta <= 90 {
			led = "U"
		} else if delta > 90 {
			led = "V"
		}
		ioutil.WriteFile("/dev/ttyUSB0", []byte(led), 0666)
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

func parseInt(c *gin.Context, v string) int {
	alpha, err := strconv.Atoi(c.Param(v))
	if err != nil {
		c.JSON(500, gin.H{"message": "Invalid parameter!"})
		return 0.0
	}
	return alpha
}
