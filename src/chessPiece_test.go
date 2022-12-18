package main

import (
	"testing"
)

func TestValidateMove(t *testing.T) {
	var chessBoard ChessBoard
	chessPiece := ChessPiece{
		pieceType: Soldier,
		owner:     RED,
	}

	testCases := []struct {
		in   Move
		want bool
	}{
		{
			Move{
				from: Position{x: 3, y: 0},
				to:   Position{x: 3, y: 1},
			}, false,
		},
		{
			Move{
				from: Position{x: 3, y: 0},
				to:   Position{x: 4, y: 0},
			}, true,
		},
		{
			Move{
				from: Position{x: 6, y: 1},
				to:   Position{x: 6, y: 2},
			}, true,
		},
		{
			Move{
				from: Position{x: 6, y: 1},
				to:   Position{x: 6, y: 1},
			}, false,
		},
		{
			Move{
				from: Position{x: 6, y: 1},
				to:   Position{x: 5, y: 1},
			}, false,
		},
		{
			Move{
				from: Position{x: 6, y: 1},
				to:   Position{x: 6, y: 0},
			}, true,
		},
	}

	for _, tc := range testCases {
		result := chessPiece.validateMove(chessBoard, tc.in)
		if result != tc.want {
			t.Errorf("validateMove: %d, %t, want %t", tc.in, result, tc.want)
		}
	}
}
