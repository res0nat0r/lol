package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	r.GET("/stats/:name", stats)

	r.Run(":" + port)
}

func summoner(c *gin.Context) {
	name := c.Param("name")
	var player Summoner
	player = getSummoner(name)

	c.JSON(http.StatusOK, player)
}

func stats(c *gin.Context) {
	name := c.Param("name")
	summoner := getSummoner(name)
	url := "https://na.api.pvp.net/api/lol/na/v1.3/stats/by-summoner/" +
		string(summoner.Id) + "/summary?season=SEASON2016&api_key=" + key

	resp, err := http.Get(url)
	e(err)

	body, err := ioutil.ReadAll(resp.Body)
	e(err)

	b := fmt.Sprintf("%s", body)
	c.String(http.StatusOK, b)
}

func getSummoner(name string) Summoner {
	var summoner map[string]Summoner
	url := "https://na.api.pvp.net/api/lol/na/v1.4/summoner/by-name/" + name + "?api_key=" + key

	resp, err := http.Get(url)
	e(err)

	body, err := ioutil.ReadAll(resp.Body)
	e(err)

	err = json.Unmarshal(body, &summoner)
	e(err)

	return summoner[name]
}

func e(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}
