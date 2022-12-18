package main

import "fmt"

type GameActions interface {
	showBoard(move Move) bool
	init()
	performMove(move Move)
}

type Game struct {
	chessBoard ChessBoard
	moves      []Move
}

func (game *Game) init() {
	game.chessBoard = [BOARD_X][BOARD_Y]*ChessPiece{}

	game.chessBoard[0][0] = &ChessPiece{Chariot, RED}
	game.chessBoard[0][1] = &ChessPiece{Horse, RED}
	game.chessBoard[0][2] = &ChessPiece{Elephant, RED}
	game.chessBoard[0][3] = &ChessPiece{Guard, RED}
	game.chessBoard[0][4] = &ChessPiece{General, RED}
	game.chessBoard[0][5] = &ChessPiece{Guard, RED}
	game.chessBoard[0][6] = &ChessPiece{Elephant, RED}
	game.chessBoard[0][7] = &ChessPiece{Horse, RED}
	game.chessBoard[0][8] = &ChessPiece{Chariot, RED}
	game.chessBoard[2][1] = &ChessPiece{Cannon, RED}
	game.chessBoard[2][7] = &ChessPiece{Cannon, RED}
	game.chessBoard[3][0] = &ChessPiece{Soldier, RED}
	game.chessBoard[3][2] = &ChessPiece{Soldier, RED}
	game.chessBoard[3][4] = &ChessPiece{Soldier, RED}
	game.chessBoard[3][6] = &ChessPiece{Soldier, RED}
	game.chessBoard[3][8] = &ChessPiece{Soldier, RED}

	game.chessBoard[9][0] = &ChessPiece{Chariot, Black}
	game.chessBoard[9][1] = &ChessPiece{Horse, Black}
	game.chessBoard[9][2] = &ChessPiece{Elephant, Black}
	game.chessBoard[9][3] = &ChessPiece{Guard, Black}
	game.chessBoard[9][4] = &ChessPiece{General, Black}
	game.chessBoard[9][5] = &ChessPiece{Guard, Black}
	game.chessBoard[9][6] = &ChessPiece{Elephant, Black}
	game.chessBoard[9][7] = &ChessPiece{Horse, Black}
	game.chessBoard[9][8] = &ChessPiece{Chariot, Black}
	game.chessBoard[7][1] = &ChessPiece{Cannon, Black}
	game.chessBoard[7][7] = &ChessPiece{Cannon, Black}
	game.chessBoard[6][0] = &ChessPiece{Soldier, Black}
	game.chessBoard[6][2] = &ChessPiece{Soldier, Black}
	game.chessBoard[6][4] = &ChessPiece{Soldier, Black}
	game.chessBoard[6][6] = &ChessPiece{Soldier, Black}
	game.chessBoard[6][8] = &ChessPiece{Soldier, Black}
}

func (game *Game) showBoard() {
	for x := BOARD_X - 1; x >= 0; x-- {
		for y := 0; y < BOARD_Y; y++ {
			if game.chessBoard[x][y] != nil {
				fmt.Printf(" %s", game.chessBoard[x][y].getPrintName())
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
}

func (game *Game) performMove(move Move) {
	game.chessBoard.performMove(move)
	game.moves = append(game.moves, move)
}

func main() {
	var game Game
	game.init()
	game.showBoard()
	game.performMove(Move{from: Position{x: 2, y: 1}, to: Position{x: 2, y: 4}})
	game.showBoard()
}
