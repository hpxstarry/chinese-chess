package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	var game Game
	game.init()
	game.showBoard()
}
