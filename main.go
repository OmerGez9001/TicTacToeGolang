package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode) //comment this line only for the use of debugging
	var activePort string
	if len(os.Args) == 2 {
		activePort = os.Args[1]
	} else {
		activePort = "3333"
	}
	r := gin.Default()
	GameOver = true
	r.GET("/start", func(c *gin.Context) {
		StartGame(c)
	})
	r.GET("/player-x", func(c *gin.Context) {
		if !GameOver {
			XPlayerTurn(c)
		}
	})
	r.GET("/player-o", func(c *gin.Context) {
		if !GameOver {
			OPlayerTurn(c)
		}
	})
	r.Run(":" + activePort)
}
