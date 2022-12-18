package main

type ChessPieceType int

const (
	Soldier ChessPieceType = iota
	Elephant
	Horse
	Chariot
	Cannon
	Guard
	General
)

type Owner int

const (
	RED Owner = iota
	Black
)

type ChessPiece struct {
	pieceType ChessPieceType
	owner     Owner
}

type ChessPieceActions interface {
	validateMove(chessBoard ChessBoard, move Move) bool
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (chessPiece ChessPiece) validateMove(chessBoard ChessBoard, move Move) bool {
	switch chessPiece.pieceType {
	case Soldier:
		if chessPiece.owner == RED {
			if move.from.x > 5 {
				movedDistance := abs(move.to.x-move.from.x) +
					abs(move.to.y-move.from.y)
				return move.to.x >= move.from.x && movedDistance == 1
			} else {
				return move.to.x == move.from.x+1 && move.to.y == move.from.y
			}
		} else {
			if move.from.x <= 5 {
				movedDistance := abs(move.to.x-move.from.x) +
					abs(move.to.y-move.from.y)
				return move.to.x <= move.from.x && movedDistance == 1
			} else {
				return move.to.x == move.from.x-1 && move.to.y == move.from.y
			}
		}
	default:
		return false
	}
}
