package main

const BOARD_X = 10 // max height
const BOARD_Y = 9  // max width

type ChessBoardActions interface {
	validateMove(move Move) bool
	performMove(move Move)
}

/*
  - - - - Y - - - -
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
    楚河 汉界          X
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
  - * * * * * * * *   |
*/
type ChessBoard [BOARD_X][BOARD_Y]*ChessPiece

func (chessBoard ChessBoard) performMove(move Move) {
	if chessBoard.validateMove(move) {
		chessBoard[move.to.x][move.to.y] = chessBoard[move.from.x][move.from.y]
		chessBoard[move.from.x][move.from.y] = nil
	}
}

func (chessBoard ChessBoard) validateMove(move Move) bool {
	fromChessPiece := chessBoard[move.from.x][move.from.y]
	toChessPiece := chessBoard[move.to.x][move.to.y]

	if fromChessPiece == nil {
		return false
	}
	// validate if "to" position has already been occupied by same owner
	if toChessPiece != nil && toChessPiece.owner == fromChessPiece.owner {
		return false
	}
	// validate based on the type of chess piece
	return true
}
