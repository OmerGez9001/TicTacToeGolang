package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartGame(c *gin.Context) {
	GameOver = false
	XTurn = true
	OTurn = true
	GameGrid = make([]string, 9)
	for i := 0; i < 9; i++ {
		GameGrid[i] = strconv.Itoa(i + 1)
	}
	StartGameIntroduction := "\n----------------------------------------------------------\n" +
		"|  Now starting tic tac toe game!  |\n" +
		"----------------------------------------------------------\n\n" +
		"Game Grid:\n" + "------------------\n" +
		"1     2     3\n4     5     6 \n7     8     9\n\n" +
		"Please mark which spot (a number from 1 to 9) to fill with X or O with the following GET request:\n" +
		"For player X: URL/player-x?mark=spot\n" +
		"For player O: URL/player-o?mark=spot\n"
	c.Data(200, "application/json", []byte(StartGameIntroduction))
}

func CurrentGrid() string {
	currentGrid := "-------------------------\n"
	for i := 0; i < 9; i++ {
		currentGrid += GameGrid[i]
		if i != 2 && i != 5 && i != 8 {
			currentGrid += "     "
		} else {
			currentGrid += "\n"
		}
	}
	return (currentGrid + "\n")
}

func WinStatus(winStreak string) bool {
	if GameGrid[0]+GameGrid[1]+GameGrid[2] == winStreak ||
		GameGrid[3]+GameGrid[4]+GameGrid[5] == winStreak ||
		GameGrid[6]+GameGrid[7]+GameGrid[8] == winStreak ||
		GameGrid[0]+GameGrid[3]+GameGrid[6] == winStreak ||
		GameGrid[1]+GameGrid[4]+GameGrid[7] == winStreak ||
		GameGrid[2]+GameGrid[5]+GameGrid[8] == winStreak ||
		GameGrid[0]+GameGrid[4]+GameGrid[8] == winStreak ||
		GameGrid[2]+GameGrid[4]+GameGrid[6] == winStreak {
		return true
	}
	return false
}

func GameGridIsFull() bool {
	for i := 0; i < 9; i++ {
		if _, err := strconv.Atoi(GameGrid[i]); err == nil {
			return false
		}
	}
	return true
}
