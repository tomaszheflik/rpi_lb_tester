package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var numer string

func generuj() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}

func root(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("RPI containr ID: %s\n", numer))
}

func main() {
	numer = strconv.Itoa(generuj())
	r := gin.Default()
	r.GET("/", root)
	r.Run(":9090")
}
