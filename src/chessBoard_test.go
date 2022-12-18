package main

import (
	"testing"
)

func TestMove(t *testing.T) {
	var chessBoard ChessBoard
	chessPiece := ChessPiece{
		pieceType: Soldier,
		owner:     Black,
	}
	chessBoard[0][1] = &chessPiece
	move := Move{
		from: Position{x: 1, y: 2},
		to:   Position{x: 0, y: 1},
	}

	result := chessBoard.validateMove(move)
	if result {
		t.Errorf("Validate move failed")
	}
}
