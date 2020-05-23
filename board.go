package main

type Color int
type PieceType int

const (
	ColorBlack Color = iota
	ColorWhite

	KingPiece PieceType = iota
	QueenPiece
	PawnPiece
	KnightPiece
	BishopPiece
	RookPiece
)

// Bitboard is the representation of a single type of piece for a single player
// Rank and file
//
// The rank are the horizontal rows 1-8
// The files are the vertical columns a-h
//
// BLACK
//  a b c d e f g h
// 8                8
// 7                7
// 6                6
// 5                5
// 4                4
// 3                3
// 2                2
// 1                1
//  a b c d e f g h
// WHITE
//
type Bitboard uint64

// Game is the top level struct for a game
type GameBoard struct {
	// Piece bitboards
	// color-independent
	Pieces map[PieceType]uint64

	// Player bitboards
	// mark which squares have black or white pieces
	Players map[Color]uint64
}

func (g GameBoard) GetPawns(c Color) uint64 {
	// Return all pawns of the given color
	return g.Pieces[PawnPiece] & g.Players[c]
}

// PlayerState contains the board state for one player
// type PlayerState struct {
// 	Color     uint8
// 	Ownership Bitboard
// }
