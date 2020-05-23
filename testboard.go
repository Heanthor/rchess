package main

func StartingPos() GameBoard {
	// starting positions
	kings := uint64(
		// white king
		(1 << 4) |
			// black king
			(1<<4)<<(8*7))
	queens := uint64(
		(1 << 3) |
			(1<<3)<<(8*7))

	knights := uint64(
		(1 << 1) | (1 << 6) |
			((1<<1)|(1<<6))<<(8*7))

	rooks := uint64(
		(1 << 0) | (1 << 7) |
			((1<<0)|(1<<7))<<(8*7))

	bishops := uint64(
		(1 << 2) | (1 << 5) |
			((1<<2)|(1<<5))<<(8*7))

	pawns := uint64(
		0xFF<<8 |
			0xFF<<(8*6))

	g := GameBoard{
		Pieces: map[PieceType]uint64{
			KingPiece:   kings,
			QueenPiece:  queens,
			KnightPiece: knights,
			BishopPiece: bishops,
			RookPiece:   rooks,
			PawnPiece:   pawns,
		},
		Players: map[Color]uint64{
			// populate first two ranks
			ColorWhite: uint64(0xFF) | uint64(0xFF)<<8,
			ColorBlack: uint64(0xFF)<<(8*6) | uint64(0xFF)<<(8*7),
		},
	}

	PrintBoard(g)

	return g
}
