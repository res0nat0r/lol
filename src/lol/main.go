package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

var (
	port = os.Getenv("PORT")
	key  = os.Getenv("RIOT_API_KEY")
)

type Summoner struct {
	Name          string `json:"name"`
	Id            int64  `json:"id"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

type Random struct {
	Summoner string `json:"name"`
	UUID     string `json:"id"`
}

func main() {
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if key == "" {
		log.Fatal("$RIOT_API_KEY must be set")
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/summoner/:name", summoner)
	r.GET("/random/:name", random)
	r.Run(":" + port)
}

func summoner(c *gin.Context) {
	name := c.Param("name")
	var summoner map[string]Summoner
	url := "https://na.api.pvp.net/api/lol/na/v1.4/summoner/by-name/" + name + "?api_key=" + key

	resp, err := http.Get(url)
	e(err)

	body, err := ioutil.ReadAll(resp.Body)
	e(err)

	err = json.Unmarshal(body, &summoner)
	e(err)

	c.JSON(http.StatusOK, summoner[name])

}

func random(c *gin.Context) {
	var resp Random
	name := c.Param("name")
	id := uuid.NewV4()

	resp.Summoner = name
	resp.UUID = fmt.Sprintf("%s", id)

	c.JSON(http.StatusOK, resp)
}

func e(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}
