package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	_ = fmt.Print
	_ = log.Print
	_ = time.Sleep
	_ = rand.Intn
	_ = os.Args
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
	Resources []Resource `json:"resource"`
	// units []unit `json:"units"`
}

type Resource struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	Name        string  `json:"name"`
	Amount      float64 `json:"amount"`
	MiningSpeed float64 `json:"speed"`
}

var NewResource = []Resource{}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:9090"
		}}))

	r.GET("/", GetVars)
	r.POST("/config", PostConfig)

	// go generateMaterials()
	csvIntoStruct()
	err := r.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

}

type a1 struct {
	Id int `json:"id"`
}

func generateMaterials() {
	var sec = time.Second
	for {
		for i:=0 ; i<len(NewResource) ; i++ {
			NewResource[i].Amount += NewResource[i].MiningSpeed
			time.Sleep(sec) 
		}
	}
}
func (c *Resource) ChangeConfig(new Resource) {
	c.MiningSpeed += new.MiningSpeed
}
func PostConfig(c *gin.Context) {
	var new a1
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// var n1 = Resource{MiningSpeed: float64(new.Id)}
	// NewResource.ChangeConfig(n1)
}

func GetVars(c *gin.Context) {
	c.JSON(http.StatusOK, NewResource)
}

func csvIntoStruct() {
	source, err := os.Open("./MOCK_DATA.csv")
	if err != nil {
		log.Fatal(err)
	}
	var reader = csv.NewReader(source)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 100; i++ { 	// len(data)
		var id int
		var user_id int
		var amount float64
		var speed float64

		_, _ = fmt.Sscan(data[i][0], &id)
		_, _ = fmt.Sscan(data[i][1], &user_id)
		_, _ = fmt.Sscan(data[i][3], &amount)
		_, _ = fmt.Sscan(data[i][4], &speed)
		var n1 = Resource{
			Id:          id,
			UserId:      user_id,
			Name: data[i][1],
			Amount:      amount,
			MiningSpeed: speed,
		}
		NewResource = append(NewResource, n1)}
}