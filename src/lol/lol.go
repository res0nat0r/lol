// This is the main lol package.
// It pulls some stats.
package lol

import (
	"encoding/json"
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

// Summoner holds various summoner data.
type Summoner struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int    `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

// Stats holds LoL game stats from the supplied summoner name.
type Stats struct {
	PlayerStatSummaries []struct {
		AggregatedStats struct {
			TotalAssists              int `json:"totalAssists"`
			TotalChampionKills        int `json:"totalChampionKills"`
			TotalMinionKills          int `json:"totalMinionKills"`
			TotalNeutralMinionsKilled int `json:"totalNeutralMinionsKilled"`
			TotalTurretsKilled        int `json:"totalTurretsKilled"`
		} `json:"aggregatedStats"`
		ModifyDate            int    `json:"modifyDate"`
		PlayerStatSummaryType string `json:"playerStatSummaryType"`
		Wins                  int    `json:"wins"`
	} `json:"playerStatSummaries"`
	SummonerID int `json:"summonerId"`
}

// New creates a new gin context and launches a server.
func New() {
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if key == "" {
		log.Fatal("$RIOT_API_KEY must be set")
	}

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/summoner/:name", summoner)
	r.GET("/stats/:season/:name", stats)

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
	season := c.Param("season")

	var stats Stats
	url := "https://na.api.pvp.net/api/lol/na/v1.3/stats/by-summoner/" +
		string(summoner.ID) + "/summary?season=SEASON" + season + "&api_key=" + key

	resp, err := http.Get(url)
	e(err)

	body, err := ioutil.ReadAll(resp.Body)
	e(err)

	err = json.Unmarshal(body, &stats)
	e(err)

	c.JSON(http.StatusOK, stats)
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
