package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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

type UserData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	User     UserData    `json:"user"`
	Resource []Resources `json:"resource"`
}

type Resources struct {
	Id     int    `json:"id"`
	UserId int   `json:"user_id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	MiningSpeed float32   `json:"speed"`
}

var NewResource = Resources{MiningSpeed: 1.5, Amount: 0, Id: 0, UserId: 1, Name: "Coin"} 

func main() {
	r := gin.Default()
	r.GET("/", GetVars)
	r.POST("/config", PostConfig)

	go NewResource.genCoins()

	err := r.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

}
type a1 struct {
	Id       int    `json:"id"`
}

func (r *Resources) genCoins() {
	for {
		r.Amount += 1
		time.Sleep(time.Duration(1000000000 / r.MiningSpeed)) // 1 second / speed
	}
}
func (c *Resources) ChangeConfig(new Resources) {
	c.MiningSpeed = new.MiningSpeed
}
func PostConfig(c *gin.Context) {
	var new a1
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(new.Id)
	var n1 Resources = Resources{MiningSpeed: float32(new.Id)}
	NewResource.ChangeConfig(n1)
}

func GetVars(c *gin.Context) {
	c.JSON(http.StatusOK, NewResource)
}
