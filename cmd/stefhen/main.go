package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	r := gin.Default()

	r.GET("/summoner/:name", summoner)

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

	c.JSON(200, summoner[name])

}

func e(err error) {
	if err != nil {
		log.Fatal("%s", err)
	}
}
