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
		Games: []memory.GameEngine{
		},
	}

	rulesString := getRules()
	var rules memory.RulesTopLevel
	err := json.Unmarshal([]byte(rulesString), &rules)
	if err != nil {
		log.Fatal(err)
	}


	//rules := memory.RulesTopLevel{}
	//err := gonfig.GetConf("./rules.json", &rules)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//
	//cardService, err := cardapiservice.NewCardApiService()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//card, err := cardService.DrawCard()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//kingsCard := memory.InitKingsCard(card, rules)
	//
	//log.Print(kingsCard)

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
      "rule": "KING!!"
    },
    {
      "code": "KD",
      "rule": "KING!!"
    },
    {
      "code": "KS",
      "rule": "KING!!"
    },
    {
      "code": "KC",
      "rule": "KING!!"
    },

    {
      "code": "QH",
      "rule": "Thumbs"
    },
    {
      "code": "QD",
      "rule": "Thumbs"
    },
    {
      "code": "QS",
      "rule": "Thumbs"
    },
    {
      "code": "QC",
      "rule": "Thumbs"
    },

    {
      "code": "JH",
      "rule": "Make a Rule"
    },
    {
      "code": "JD",
      "rule": "Make a Rule"
    },
    {
      "code": "JS",
      "rule": "Make a Rule"
    },
    {
      "code": "JC",
      "rule": "Make a Rule"
    },

    {
      "code": "0H",
      "rule": "Categories"
    },
    {
      "code": "0D",
      "rule": "Categories"
    },
    {
      "code": "0S",
      "rule": "Categories"
    },
    {
      "code": "0C",
      "rule": "Categories"
    },

    {
      "code": "9H",
      "rule": "Rhymes"
    },
    {
      "code": "9D",
      "rule": "Rhymes"
    },
    {
      "code": "9S",
      "rule": "Rhymes"
    },
    {
      "code": "9C",
      "rule": "Rhymes"
    },

    {
      "code": "8H",
      "rule": "Me and a Mate"
    },
    {
      "code": "8D",
      "rule": "Me and a Mate"
    },
    {
      "code": "8S",
      "rule": "Me and a Mate"
    },
    {
      "code": "8C",
      "rule": "Me and a Mate"
    },

    {
      "code": "7H",
      "rule": "Heaven"
    },
    {
      "code": "7D",
      "rule": "Heaven"
    },
    {
      "code": "7S",
      "rule": "Heaven"
    },
    {
      "code": "7C",
      "rule": "Heaven"
    },

    {
      "code": "6H",
      "rule": "Boys Drink"
    },
    {
      "code": "6D",
      "rule": "Boys Drink"
    },
    {
      "code": "6S",
      "rule": "Boys Drink"
    },
    {
      "code": "6C",
      "rule": "Boys Drink"
    },

    {
      "code": "5H",
      "rule": "Social"
    },
    {
      "code": "5D",
      "rule": "Social"
    },
    {
      "code": "5S",
      "rule": "Social"
    },
    {
      "code": "5C",
      "rule": "Social"
    },

    {
      "code": "4H",
      "rule": "Girls Drink"
    },
    {
      "code": "4D",
      "rule": "Girls Drink"
    },
    {
      "code": "4S",
      "rule": "Girls Drink"
    },
    {
      "code": "4C",
      "rule": "Girls Drink"
    },

    {
      "code": "3H",
      "rule": "You Drink"
    },
    {
      "code": "3D",
      "rule": "You Drink"
    },
    {
      "code": "3S",
      "rule": "You Drink"
    },
    {
      "code": "3C",
      "rule": "You Drink"
    },

    {
      "code": "2H",
      "rule": "Nominate a Drink"
    },
    {
      "code": "2D",
      "rule": "Nominate a Drink"
    },
    {
      "code": "2S",
      "rule": "Nominate a Drink"
    },
    {
      "code": "2C",
      "rule": "Nominate a Drink"
    },

    {
      "code": "AH",
      "rule": "Waterfall"
    },
    {
      "code": "AD",
      "rule": "Waterfall"
    },
    {
      "code": "AS",
      "rule": "Waterfall"
    },
    {
      "code": "AC",
      "rule": "Waterfall"
    }
  ]
}
`
}
