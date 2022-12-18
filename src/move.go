package main

type Position struct {
	x int
	y int
}

type Move struct {
	chessPiece ChessPiece
	from       Position
	to         Position
}
