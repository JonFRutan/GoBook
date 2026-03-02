package main

import (
	"fmt"
	"slices"
	"math/rand"
)

type gameBoard struct {
	Board [][]string
	UsedPositions []int
}

func (game *gameBoard) CheckWin () {
	for i, value := range game.Board {
		if value == "X" {
			return
		} else if value == "O" {
			return
		}
	}
}

//currently just picks a random number until an empty position is found
func (game *gameBoard) OpponentTurn () {
	for {
		tryPosition := rand.Intn(9)
		if game.PlayPiece(2, tryPosition) {
			fmt.Printf("Opponent Plays Position %d\n", tryPosition)
			return
		} 
	}
}

func (game *gameBoard) PlayPiece(player int, position int) bool {
	if slices.Contains(game.UsedPositions, position) {
		return false
	} 
	var character string
	switch player {
	case 1:
		character = "X"
	case 2:
		character = "O"
	}
	xpos := position / 3
	ypos := position % 3
	game.Board[xpos][ypos] = character
	game.UsedPositions = append(game.UsedPositions, position)
	return true
}

func main() {
	board := [][]string{ []string{"0", "1", "2"}, []string{"3", "4", "5"}, []string{"6", "7", "8"}}
	var positions []int
	game := gameBoard{Board: board, UsedPositions: positions}

	var playPos int
	fmt.Printf("%s\n%s\n%s\n", game.Board[0], game.Board[1], game.Board[2])
	//equivalent of while True loop
	for {
		for {
			fmt.Print("Choose your next position: ")
			fmt.Scan(&playPos)
			if game.PlayPiece(1, playPos) {
				break
			}
		}
		game.OpponentTurn()
		fmt.Printf("%s\n%s\n%s\n", game.Board[0], game.Board[1], game.Board[2])
	}
}