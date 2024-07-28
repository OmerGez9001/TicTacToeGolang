package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func XPlayerTurn(c *gin.Context) {
	if !XTurn {
		c.Data(200, "application/json", []byte("Now it's O player's turn. Current game grid:\n"+CurrentGrid()))
		return
	}
	response, hasMoved := PlayerMove("X", c.Request.URL.Query().Get("mark"))
	c.Data(200, "application/json", response)
	if hasMoved {
		XTurn = false
		OTurn = true
	}
}

func OPlayerTurn(c *gin.Context) {
	if !OTurn {
		c.Data(200, "application/json", []byte("Now it's X player's turn. Current game grid:\n"+CurrentGrid()))
		return
	}
	response, hasMoved := PlayerMove("O", c.Request.URL.Query().Get("mark"))
	c.Data(200, "application/json", response)
	if hasMoved {
		OTurn = false
		XTurn = true
	}
}

func PlayerMove(p, mark string) ([]byte, bool) {
	index, err := strconv.Atoi(mark)
	if err != nil || mark == "" {
		return []byte("Invalid value. Please make a GET request with the following param: ?mark=NUMBER(1,9)"), false
	}
	if index < 1 || index > 9 {
		return []byte("number chosen not in range. Please pick a number between 1-9"), false
	}
	if GameGrid[index-1] == "X" || GameGrid[index-1] == "O" {
		return []byte("Spot taken, try again. Current game grid:\n" + CurrentGrid()), false
	}
	GameGrid[index-1] = p
	if WinStatus(p + p + p) {
		GameOver = true
		return []byte("Player " + p + " Won!\nWinning grid:\n" + CurrentGrid()), false
	}
	if GameGridIsFull() {
		GameOver = true
		return []byte("Game over, it's a draw. Endgame grid:\n" + CurrentGrid()), false
	}
	return []byte("Player " + p + " has made a move. Current grid:\n" + CurrentGrid()), true
}
