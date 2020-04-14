package main

import (
	"encoding/json"
	"git.seevo.online/jgerlach/Games.Kings.API/controller"
	"git.seevo.online/jgerlach/Games.Kings.API/model/memory"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

var games memory.Games

func main() {

	games = memory.Games{
		Games: []memory.GameEngine{},
	}

	rulesString := getRules()
	var rules memory.RulesTopLevel
	err := json.Unmarshal([]byte(rulesString), &rules)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/games/:room_id", controller.GetGame(&games))

		apiRoutes.POST("/games", controller.PostNewGame(&games))

		apiRoutes.POST("/games/:room_id/players", controller.PostAddPlayer(&games))
		apiRoutes.POST("/games/:room_id/players/draw", controller.PostPlayerDrawCard(&games, rules))

	}

	r.Run(":8080")

}

func getRules() string {
	return `
{
  "rules" :[
    {
      "code": "KH",
      "rule": "KING!!",
      "holdable" : false
    },
    {
      "code": "KD",
      "rule": "KING!!",
      "holdable" : false
    },
    {
      "code": "KS",
      "rule": "KING!!",
      "holdable" : false
    },
    {
      "code": "KC",
      "rule": "KING!!",
      "holdable" : false
    },

    {
      "code": "QH",
      "rule": "Thumbs",
      "holdable" : true
    },
    {
      "code": "QD",
      "rule": "Thumbs",
      "holdable" : true
    },
    {
      "code": "QS",
      "rule": "Thumbs",
      "holdable" : true
    },
    {
      "code": "QC",
      "rule": "Thumbs",
      "holdable" : true
    },

    {
      "code": "JH",
      "rule": "Make a Rule",
      "holdable" : false
    },
    {
      "code": "JD",
      "rule": "Make a Rule",
      "holdable" : false
    },
    {
      "code": "JS",
      "rule": "Make a Rule",
      "holdable" : false
    },
    {
      "code": "JC",
      "rule": "Make a Rule",
      "holdable" : false
    },

    {
      "code": "0H",
      "rule": "Categories",
      "holdable" : false
    },
    {
      "code": "0D",
      "rule": "Categories",
      "holdable" : false
    },
    {
      "code": "0S",
      "rule": "Categories",
      "holdable" : false
    },
    {
      "code": "0C",
      "rule": "Categories",
      "holdable" : false
    },

    {
      "code": "9H",
      "rule": "Rhymes",
      "holdable" : false
    },
    {
      "code": "9D",
      "rule": "Rhymes",
      "holdable" : false
    },
    {
      "code": "9S",
      "rule": "Rhymes",
      "holdable" : false
    },
    {
      "code": "9C",
      "rule": "Rhymes",
      "holdable" : false
    },

    {
      "code": "8H",
      "rule": "Me and a Mate",
      "holdable" : false
    },
    {
      "code": "8D",
      "rule": "Me and a Mate",
      "holdable" : false
    },
    {
      "code": "8S",
      "rule": "Me and a Mate",
      "holdable" : false
    },
    {
      "code": "8C",
      "rule": "Me and a Mate",
      "holdable" : false
    },

    {
      "code": "7H",
      "rule": "Heaven",
      "holdable" : true
    },
    {
      "code": "7D",
      "rule": "Heaven",
      "holdable" : true
    },
    {
      "code": "7S",
      "rule": "Heaven",
      "holdable" : true
    },
    {
      "code": "7C",
      "rule": "Heaven",
      "holdable" : true
    },

    {
      "code": "6H",
      "rule": "Boys Drink",
      "holdable" : false
    },
    {
      "code": "6D",
      "rule": "Boys Drink",
      "holdable" : false
    },
    {
      "code": "6S",
      "rule": "Boys Drink",
      "holdable" : false
    },
    {
      "code": "6C",
      "rule": "Boys Drink",
      "holdable" : false
    },

    {
      "code": "5H",
      "rule": "Social",
      "holdable" : false
    },
    {
      "code": "5D",
      "rule": "Social",
      "holdable" : false
    },
    {
      "code": "5S",
      "rule": "Social",
      "holdable" : false
    },
    {
      "code": "5C",
      "rule": "Social",
      "holdable" : false
    },

    {
      "code": "4H",
      "rule": "Girls Drink",
      "holdable" : false
    },
    {
      "code": "4D",
      "rule": "Girls Drink",
      "holdable" : false
    },
    {
      "code": "4S",
      "rule": "Girls Drink",
      "holdable" : false
    },
    {
      "code": "4C",
      "rule": "Girls Drink",
      "holdable" : false
    },

    {
      "code": "3H",
      "rule": "You Drink",
      "holdable" : false
    },
    {
      "code": "3D",
      "rule": "You Drink",
      "holdable" : false
    },
    {
      "code": "3S",
      "rule": "You Drink",
      "holdable" : false
    },
    {
      "code": "3C",
      "rule": "You Drink",
      "holdable" : false
    },

    {
      "code": "2H",
      "rule": "Nominate a Drink",
      "holdable" : false
    },
    {
      "code": "2D",
      "rule": "Nominate a Drink",
      "holdable" : false
    },
    {
      "code": "2S",
      "rule": "Nominate a Drink",
      "holdable" : false
    },
    {
      "code": "2C",
      "rule": "Nominate a Drink",
      "holdable" : false
    },

    {
      "code": "AH",
      "rule": "Waterfall",
      "holdable" : false
    },
    {
      "code": "AD",
      "rule": "Waterfall",
      "holdable" : false
    },
    {
      "code": "AS",
      "rule": "Waterfall",
      "holdable" : false
    },
    {
      "code": "AC",
      "rule": "Waterfall",
      "holdable" : false
    }
  ]
}
`
}
