package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var pieceTypeMap = map[PieceType]string{
	KingPiece:   " K ",
	QueenPiece:  " Q ",
	PawnPiece:   " P ",
	KnightPiece: " N ",
	BishopPiece: " B ",
	RookPiece:   " R ",
}

const (
	rank = "  a  b  c  d  e  f  g  h"
)

// PrintBoard prints an ascii representation of the GameBoard to stdout
func PrintBoard(g GameBoard) {
	buf := &bytes.Buffer{}
	black := color.New(color.FgBlack)
	white := color.New(color.FgWhite)

	fmt.Println(rank)

	blackPieces := g.Players[ColorBlack]
	whitePieces := g.Players[ColorWhite]

	shift := uint64(1)
	for i := 0; i < 8; i++ {
		// print from rank 1 to 8
		fmt.Fprint(buf, strconv.Itoa(i+1))
		for j := 0; j < 8; j++ {
			// print each piece on the rank
			// for black and for white
			pt := getPiece(g, shift)
			if shift&blackPieces > 0 {
				black.Fprint(buf, pieceTypeMap[pt])
			} else if shift&whitePieces > 0 {
				white.Fprint(buf, pieceTypeMap[pt])
			} else {
				buf.WriteString("   ")
			}

			shift <<= 1
		}
		fmt.Fprint(buf, "\n")
	}

	fmt.Println(string(buf.Bytes()))
}

func getPiece(g GameBoard, shift uint64) PieceType {
	for pt, pb := range g.Pieces {
		if pb&shift > 0 {
			return pt
		}
	}

	return -1
}
