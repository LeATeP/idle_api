package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"net/http"
	"database/sql"
	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

)
var (
	_ = fmt.Print
	_ = log.Print
	_ = time.Sleep
	_ = rand.Intn
	_ = http.ListenAndServe
	_ = gin.Default
	_ = sql.Open
	_ = json.NewEncoder
)

type vars struct {
	Num int `json:"number"`
}


func main() {
	rou := gin.Default()
	rou.GET("/", GetVars)

	_ = rou.Run(":9090")
}

func GetVars(c *gin.Context) {
	a := vars{Num: rand.Int()}
	c.JSON(http.StatusOK, a)
}