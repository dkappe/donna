// Copyright (c) 2013-2014 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package donna

import(`github.com/michaeldv/donna/expect`; `testing`)

func TestPositionMoves010(t *testing.T) {
	p := NewGame(`Ke1,e2`, `Kg8,d7,f7`).Start(White)
	expect.Eq(t, p.enpassant, 0)

	p = p.MakeMove(NewMove(p, E2, E4))
	expect.Eq(t, p.enpassant, 0)

	p = p.MakeMove(NewMove(p, D7, D5))
	expect.Eq(t, p.enpassant, 0)

	p = p.MakeMove(NewMove(p, E4, E5))
	expect.Eq(t, p.enpassant, 0)

	p = p.MakeMove(NewEnpassant(p, F7, F5))
	expect.Eq(t, p.enpassant, F6)
}

// Castle tests.
func TestPositionMoves030(t *testing.T) { // Everything is OK.
	p := NewGame(`Ke1,Ra1,Rh1`, `Ke8`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.True(t, kingside)
	expect.True(t, queenside)

	p = NewGame(`Ke1`, `Ke8,Ra8,Rh8`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.True(t, kingside)
	expect.True(t, queenside)
}

func TestPositionMoves040(t *testing.T) { // King checked.
	p := NewGame(`Ke1,Ra1,Rh1`, `Ke8,Bg3`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke1,Bg6`, `Ke8,Ra8,Rh8`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves050(t *testing.T) { // Attacked square.
	p := NewGame(`Ke1,Ra1,Rh1`, `Ke8,Bb3,Bh3`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke1,Bb6,Bh6`, `Ke8,Ra8,Rh8`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves060(t *testing.T) { // Wrong square.
	p := NewGame(`Ke1,Ra8,Rh8`, `Ke5`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke2,Ra1,Rh1`, `Ke8`).Start(White)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke4`, `Ke8,Ra1,Rh1`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke4`, `Ke7,Ra8,Rh8`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves070(t *testing.T) { // Missing rooks.
	p := NewGame(`Ke1`, `Ke8`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke1`, `Ke8`).Start(Black)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves080(t *testing.T) { // Rooks on wrong squares.
	p := NewGame(`Ke1,Rb1`, `Ke8`).Start(White)
	kingside, queenside := p.canCastle(p.color)
	expect.False(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke1,Rb1,Rh1`, `Ke8`).Start(White)
	kingside, queenside = p.canCastle(p.color)
	expect.True(t, kingside)
	expect.False(t, queenside)

	p = NewGame(`Ke1,Ra1,Rf1`, `Ke8`).Start(White)
	kingside, queenside = p.canCastle(p.color)
	expect.False(t, kingside)
	expect.True(t, queenside)
}

func TestPositionMoves081(t *testing.T) { // Rook has moved.
	p := NewGame(`Ke1,Ra1,Rh1`, `Ke8`).Start(White)
	p = p.MakeMove(NewMove(p, A1, A2))
	p = p.MakeMove(NewMove(p, E8, E7))
	p = p.MakeMove(NewMove(p, A2, A1))

	kingside, queenside := p.canCastle(White)
	expect.True(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves082(t *testing.T) { // King has moved.
	p := NewGame(`Ke1`, `Ke8,Ra8,Rh8`).Start(Black)
	p = p.MakeMove(NewMove(p, E8, E7))
	p = p.MakeMove(NewMove(p, E1, E2))
	p = p.MakeMove(NewMove(p, E7, E8))

	kingside, queenside := p.canCastle(Black)
	expect.False(t, kingside)
	expect.False(t, queenside)
}

func TestPositionMoves083(t *testing.T) { // Rook is taken.
	p := NewGame(`Ke1,Nb6`, `Ke8,Ra8,Rh8`).Start(White)
	p = p.MakeMove(NewMove(p, B6, A8))

	kingside, queenside := p.canCastle(Black)
	expect.True(t, kingside)
	expect.False(t, queenside)
}

// Blocking kingside knight.
func TestPositionMoves084(t *testing.T) {
	p := NewGame(`Ke1`, `Ke8,Ra8,Rh8,Ng8`).Start(Black)

	kingside, queenside := p.canCastle(Black)
	expect.False(t, kingside)
	expect.True(t, queenside)
}

// Blocking queenside knight.
func TestPositionMoves085(t *testing.T) {
	p := NewGame(`Ke1`, `Ke8,Ra8,Rh8,Nb8`).Start(Black)

	kingside, queenside := p.canCastle(Black)
	expect.True(t, kingside)
	expect.False(t, queenside)
}

// Straight repetition.
func TestPositionMoves100(t *testing.T) {
	p := NewGame().Start() // Initial 1.
	p = p.MakeMove(NewMove(p, G1, F3))
	p = p.MakeMove(NewMove(p, G8, F6)) // 1.
	expect.False(t, p.repetition())
	expect.False(t, p.thirdRepetition())

	p = p.MakeMove(NewMove(p, F3, G1))
	p = p.MakeMove(NewMove(p, F6, G8)) // Initial 2.
	expect.True(t, p.repetition())
	expect.False(t, p.thirdRepetition())

	p = p.MakeMove(NewMove(p, G1, F3))
	p = p.MakeMove(NewMove(p, G8, F6)) // 2.
	expect.True(t, p.repetition())
	expect.False(t, p.thirdRepetition())

	p = p.MakeMove(NewMove(p, F3, G1))
	p = p.MakeMove(NewMove(p, F6, G8)) // Initial 3.
	expect.True(t, p.repetition())
	expect.True(t, p.thirdRepetition())

	p = p.MakeMove(NewMove(p, G1, F3))
	p = p.MakeMove(NewMove(p, G8, F6)) // 3.
	expect.True(t, p.repetition())
	expect.True(t, p.thirdRepetition())
}

// Repetition with some moves in between.
func TestPositionMoves110(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, E7, E5))
	p = p.MakeMove(NewMove(p, G1, F3))
	p = p.MakeMove(NewMove(p, G8, F6)) // 1.
	p = p.MakeMove(NewMove(p, B1, C3))
	p = p.MakeMove(NewMove(p, B8, C6))
	p = p.MakeMove(NewMove(p, F1, C4))
	p = p.MakeMove(NewMove(p, F8, C5))
	p = p.MakeMove(NewMove(p, C3, B1))
	p = p.MakeMove(NewMove(p, C6, B8))
	p = p.MakeMove(NewMove(p, C4, F1))
	p = p.MakeMove(NewMove(p, C5, F8)) // 2.
	expect.True(t, p.repetition())
	expect.False(t, p.thirdRepetition())

	p = p.MakeMove(NewMove(p, F1, C4))
	p = p.MakeMove(NewMove(p, F8, C5))
	p = p.MakeMove(NewMove(p, B1, C3))
	p = p.MakeMove(NewMove(p, B8, C6))
	p = p.MakeMove(NewMove(p, C4, F1))
	p = p.MakeMove(NewMove(p, C5, F8))
	p = p.MakeMove(NewMove(p, C3, B1))
	p = p.MakeMove(NewMove(p, C6, B8)) // 3.
	expect.True(t, p.repetition())
	expect.True(t, p.thirdRepetition())
}

// Irreversible 0-0.
func TestPositionMoves120(t *testing.T) {
	p := NewGame(`Ke1,Rh1,h2`, `Ke8,Ra8,a7`).Start(White)
	p = p.MakeMove(NewMove(p, H2, H4))
	p = p.MakeMove(NewMove(p, A7, A5)) // 1.
	p = p.MakeMove(NewMove(p, E1, E2))
	p = p.MakeMove(NewMove(p, E8, E7)) // King has moved.
	p = p.MakeMove(NewMove(p, E2, E1))
	p = p.MakeMove(NewMove(p, E7, E8)) // 2.
	p = p.MakeMove(NewMove(p, E1, E2))
	p = p.MakeMove(NewMove(p, E8, E7)) // King has moved again.
	p = p.MakeMove(NewMove(p, E2, E1))
	p = p.MakeMove(NewMove(p, E7, E8)) // 3.
	expect.True(t, p.repetition())
	expect.False(t, p.thirdRepetition()) // <-- Lost 0-0 right.

	p = p.MakeMove(NewMove(p, E1, E2))
	p = p.MakeMove(NewMove(p, E8, E7)) // King has moved again.
	p = p.MakeMove(NewMove(p, E2, E1))
	p = p.MakeMove(NewMove(p, E7, E8)) // 4.
	expect.True(t, p.repetition())
	expect.True(t, p.thirdRepetition()) // <-- 3 time repetioion with lost 0-0 right.
}

// Incremental hash recalculation tests (see book_test.go).
func TestPositionMoves200(t *testing.T) { // 1. e4
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x823C9B50FD114196))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0x0B2D6B38C0B92E91))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, uint8(0x0F))
}

func TestPositionMoves210(t *testing.T) { // 1. e4 d5
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, D7, D5))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x0756B94461C50FB0))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0x76916F86F34AE5BE))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, uint8(0x0F))
}

// 1. e4 d5 2. e5
func TestPositionMoves220(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, E5))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x662FAFB965DB29D4))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0xEF3E5FD1587346D3))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, uint8(0x0F))
}

// 1. e4 d5 2. e5 f5 <-- Enpassant
func TestPositionMoves230(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, E5))
	p = p.MakeMove(NewEnpassant(p, F7, F5))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x22A48B5A8E47FF78))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0x83871FE249DCEE04))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, F6)
	expect.Eq(t, p.castles, uint8(0x0F))
}

// 1. e4 d5 2. e5 f5 3. Ke2 <-- White Castle
func TestPositionMoves240(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, E5))
	p = p.MakeMove(NewMove(p, F7, F5))
	p = p.MakeMove(NewMove(p, E1, E2))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x652A607CA3F242C1))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0x83871FE249DCEE04))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, castleKingside[Black]|castleQueenside[Black])
}

// 1. e4 d5 2. e5 f5 3. Ke2 Kf7 <-- Black Castle
func TestPositionMoves250(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, E5))
	p = p.MakeMove(NewMove(p, F7, F5))
	p = p.MakeMove(NewMove(p, E1, E2))
	p = p.MakeMove(NewMove(p, E8, F7))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x00FDD303C946BDD9))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0x83871FE249DCEE04))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, uint8(0))
}

// 1. a2a4 b7b5 2. h2h4 b5b4 3. c2c4 <-- Enpassant
func TestPositionMoves260(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, A2, A4))
	p = p.MakeMove(NewMove(p, B7, B5))
	p = p.MakeMove(NewMove(p, H2, H4))
	p = p.MakeMove(NewMove(p, B5, B4))
	p = p.MakeMove(NewEnpassant(p, C2, C4))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x3C8123EA7B067637))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0xB5AA405AF42E7052))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1)
	expect.Eq(t, p.enpassant, C3)
	expect.Eq(t, p.castles, uint8(0x0F))
}

// 1. a2a4 b7b5 2. h2h4 b5b4 3. c2c4 b4xc3 4. Ra1a3 <-- Enpassant/Castle
func TestPositionMoves270(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, A2, A4))
	p = p.MakeMove(NewMove(p, B7, B5))
	p = p.MakeMove(NewMove(p, H2, H4))
	p = p.MakeMove(NewMove(p, B5, B4))
	p = p.MakeMove(NewEnpassant(p, C2, C4))
	p = p.MakeMove(NewMove(p, B4, C3))
	p = p.MakeMove(NewMove(p, A1, A3))
	hash, pawnHash := p.polyglot()

	expect.Eq(t, hash, uint64(0x5C3F9B829B279560))
	expect.Eq(t, hash, p.hash)
	expect.Eq(t, pawnHash, uint64(0xE214F040EAA135A0))
	expect.Eq(t, pawnHash, p.pawnHash)

	expect.Eq(t, p.balance, len(materialBase) - 1 - materialBalance[Pawn])
	expect.Eq(t, p.enpassant, 0)
	expect.Eq(t, p.castles, castleKingside[White] | castleKingside[Black] | castleQueenside[Black])
}

// Incremental material hash calculation.

// 1. e4 d5 2. e4xd5
func TestPositionMoves280(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4)); p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, D5))

	expect.Eq(t, p.balance, len(materialBase) - 1 - materialBalance[BlackPawn])
}

// 1. e4 d5 2. e4xd5 Ng8-f6 3. Nb1-c3 Nf6xd5
func TestPositionMoves281(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4)); p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, D5)); p = p.MakeMove(NewMove(p, G8, F6))
	p = p.MakeMove(NewMove(p, B1, C3)); p = p.MakeMove(NewMove(p, F6, D5))

	expect.Eq(t, p.balance, len(materialBase) - 1 - materialBalance[Pawn] - materialBalance[BlackPawn])
}

// 1. e4 d5 2. e4xd5 Ng8-f6 3. Nb1-c3 Nf6xd5 4. Nc3xd5 Qd8xd5
func TestPositionMoves282(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4)); p = p.MakeMove(NewMove(p, D7, D5))
	p = p.MakeMove(NewMove(p, E4, D5)); p = p.MakeMove(NewMove(p, G8, F6))
	p = p.MakeMove(NewMove(p, B1, C3)); p = p.MakeMove(NewMove(p, F6, D5))
	p = p.MakeMove(NewMove(p, C3, D5)); p = p.MakeMove(NewMove(p, D8, D5))

	expect.Eq(t, p.balance, len(materialBase) - 1 - materialBalance[Pawn] - materialBalance[Knight] - materialBalance[BlackPawn] - materialBalance[BlackKnight])
}

// Pawn promotion.
func TestPositionMoves283(t *testing.T) {
	p := NewGame(`Kh1`, `Ka8,a2,b7`).Start(Black)
	expect.Eq(t, p.balance, 2 * materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, A1).promote(Rook))
	expect.Eq(t, p.balance, materialBalance[BlackPawn] + materialBalance[BlackRook])
}

// Last pawn promotion.
func TestPositionMoves284(t *testing.T) {
	p := NewGame(`Kh1`, `Ka8,a2`).Start(Black)
	expect.Eq(t, p.balance, materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, A1).promote(Rook))
	expect.Eq(t, p.balance, materialBalance[BlackRook])
}

// Pawn promotion with capture.
func TestPositionMoves285(t *testing.T) {
	p := NewGame(`Kh1,Nb1,Ng1`, `Ka8,a2,b7`).Start(Black)
	expect.Eq(t, p.balance, 2 * materialBalance[Knight] + 2 * materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, B1).promote(Queen))
	expect.Eq(t, p.balance, materialBalance[Knight] + materialBalance[BlackPawn] + materialBalance[BlackQueen])
}

// Pawn promotion with last piece capture.
func TestPositionMoves286(t *testing.T) {
	p := NewGame(`Kh1,Nb1`, `Ka8,a2,b7`).Start(Black)
	expect.Eq(t, p.balance, materialBalance[Knight] + 2 * materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, B1).promote(Queen))
	expect.Eq(t, p.balance, materialBalance[BlackPawn] + materialBalance[BlackQueen])
}

// Last pawn promotion with capture.
func TestPositionMoves287(t *testing.T) {
	p := NewGame(`Kh1,Nb1,Ng1`, `Ka8,a2`).Start(Black)
	expect.Eq(t, p.balance, 2 * materialBalance[Knight] + materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, B1).promote(Queen))
	expect.Eq(t, p.balance, materialBalance[Knight] + materialBalance[BlackQueen])
}

// Last pawn promotion with last piece capture.
func TestPositionMoves288(t *testing.T) {
	p := NewGame(`Kh1,Nb1`, `Ka8,a2`).Start(Black)
	expect.Eq(t, p.balance, materialBalance[Knight] + materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, A2, B1).promote(Queen))
	expect.Eq(t, p.balance, materialBalance[BlackQueen])
}

// Capture.
func TestPositionMoves289(t *testing.T) {
	p := NewGame(`Kh1,Nc3,Nf3`, `Ka8,d4,e4`).Start(Black)
	expect.Eq(t, p.balance, 2 * materialBalance[Knight] + 2 * materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, D4, C3))
	expect.Eq(t, p.balance, materialBalance[Knight] + 2 * materialBalance[BlackPawn])
}

// Last piece capture.
func TestPositionMoves290(t *testing.T) {
	p := NewGame(`Kh1,Nc3`, `Ka8,d4,e4`).Start(Black)
	expect.Eq(t, p.balance, materialBalance[Knight] + 2 * materialBalance[BlackPawn])

	p = p.MakeMove(NewMove(p, D4, C3))
	expect.Eq(t, p.balance, 2 * materialBalance[BlackPawn])
}

// En-passant capture: 1. e2-e4 e7-e6 2. e4-e5 d7-d5 3. e4xd5
func TestPositionMoves291(t *testing.T) {
	p := NewGame().Start()
	expect.Eq(t, p.balance, len(materialBase) - 1)

	p = p.MakeMove(NewMove(p, E2, E4)); p = p.MakeMove(NewMove(p, E7, E6))
	p = p.MakeMove(NewMove(p, E4, E5)); p = p.MakeMove(NewEnpassant(p, D7, D5))
	p = p.MakeMove(NewMove(p, E5, D6))
	expect.Eq(t, p.balance, len(materialBase) - 1 - materialBalance[BlackPawn])
}

// Last pawn en-passant capture.
func TestPositionMoves292(t *testing.T) {
	p := NewGame(`Kh1,c2`, `Ka8,d4`).Start(White)
	expect.Eq(t, p.balance, materialBalance[Pawn] + materialBalance[BlackPawn])

	p = p.MakeMove(NewEnpassant(p, C2, C4)); p = p.MakeMove(NewMove(p, D4, C3))
	expect.Eq(t, p.balance, materialBalance[BlackPawn])
}

// Unobstructed pins.
func TestPositionMoves300(t *testing.T) {
	position := NewGame(`Ka1,Qe1,Ra8,Rh8,Bb5`, `Ke8,Re7,Bc8,Bf8,Nc6`).Start(White)
	pinned := position.pinnedMask(E8)

	expect.Eq(t, pinned, bit[C6]|bit[C8]|bit[E7]|bit[F8])
}

func TestPositionMoves310(t *testing.T) {
	position := NewGame(`Ke4,Qe5,Rd5,Nd4,Nf4`, `Ka7,Qe8,Ra4,Rh4,Ba8`).Start(Black)
	pinned := position.pinnedMask(E4)

	expect.Eq(t, pinned, bit[D5]|bit[E5]|bit[D4]|bit[F4])
}

// Not a pin (friendly blockers).
func TestPositionMoves320(t *testing.T) {
	position := NewGame(`Ka1,Qe1,Ra8,Rh8,Bb5,Nb8,Ng8,e4`, `Ke8,Re7,Bc8,Bf8,Nc6`).Start(White)
	pinned := position.pinnedMask(E8)

	expect.Eq(t, pinned, bit[C6])
}

func TestPositionMoves330(t *testing.T) {
	position := NewGame(`Ke4,Qe7,Rc6,Nb4,Ng4`, `Ka7,Qe8,Ra4,Rh4,Ba8,c4,e6,f4`).Start(Black)
	pinned := position.pinnedMask(E4)

	expect.Eq(t, pinned, bit[C6])
}

// Not a pin (enemy blockers).
func TestPositionMoves340(t *testing.T) {
	position := NewGame(`Ka1,Qe1,Ra8,Rh8,Bb5`, `Ke8,Re7,Rg8,Bc8,Bf8,Nc6,Nb8,e4`).Start(White)
	pinned := position.pinnedMask(E8)

	expect.Eq(t, pinned, bit[C6])
}

func TestPositionMoves350(t *testing.T) {
	position := NewGame(`Ke4,Qe7,Rc6,Nb4,Ng4,c4,e5,f4`, `Ka7,Qe8,Ra4,Rh4,Ba8`).Start(Black)
	pinned := position.pinnedMask(E4)

	expect.Eq(t, pinned, bit[C6])
}

// Position after null move.
func TestPositionMoves400(t *testing.T) {
	p := NewGame(`Ke1,Qd1,d2,e2`, `Kg8,Qf8,f7,g7`).Start(White)

	p = p.MakeNullMove()
	expect.True(t, p.isNull())

	p = p.UndoNullMove()
	p = p.MakeMove(NewMove(p, E2, E4))
	expect.False(t, p.isNull())
}

// isInCheck
func TestPositionMoves410(t *testing.T) {
	p := NewGame().Start()
	p = p.MakeMove(NewMove(p, E2, E4))
	p = p.MakeMove(NewMove(p, F7, F6))
	position := p.MakeMove(NewMove(p, D1, H5))

	expect.True(t, position.isInCheck(position.color))
	expect.True(t, position.isInCheck(p.color^1))
}
