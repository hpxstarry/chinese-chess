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
	getPrintName() string
}

func (chessPiece ChessPiece) getPrintName() string {
	switch chessPiece.pieceType {
	case Soldier:
		if chessPiece.owner == RED {
			return "兵"
		}
		return "卒"
	case Elephant:
		if chessPiece.owner == RED {
			return "相"
		}
		return "象"
	case Horse:
		if chessPiece.owner == RED {
			return "傌"
		}
		return "馬"
	case Chariot:
		if chessPiece.owner == RED {
			return "车"
		}
		return "車"
	case Cannon:
		if chessPiece.owner == RED {
			return "炮"
		}
		return "砲"
	case Guard:
		if chessPiece.owner == RED {
			return "仕"
		}
		return "士"
	case General:
		if chessPiece.owner == RED {
			return "帅"
		}
		return "将"
	default:
		return ""
	}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
