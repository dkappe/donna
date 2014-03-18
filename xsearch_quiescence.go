// Copyright (c) 2013-2014 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package donna

import(`fmt`)

// Quiescence search.
func (p *Position) xSearchQuiescence(alpha, beta int, checks bool) int {
        if p.isRepetition() {
                return 0
        }

        bestScore := p.Evaluate()
        if bestScore > alpha {
                if bestScore >= beta {
                        return bestScore
                }
                alpha = bestScore
        }

        gen := p.StartMoveGen(Ply()).GenerateCaptures().rank()
        for move := gen.NextMove(); move != 0; move = gen.NextMove() {
                if position := p.MakeMove(move); position != nil {
                        fmt.Printf("%*squie/%s> ply: %d, move: %s\n", Ply()*2, ` `, C(p.color), Ply(), move)
                        moveScore := 0
                        if position.isInCheck(position.color) {
                                moveScore = -position.xSearchQuiescenceInCheck(-beta, -alpha)
                        } else {
                                moveScore = -position.xSearchQuiescence(-beta, -alpha, false)
                        }

                        position.TakeBack(move)
                        if moveScore > bestScore {
                                if moveScore > alpha {
                                        if moveScore >= beta {
                                                return moveScore
                                        }
                                        alpha = moveScore
                                }
                                beta = moveScore
                        }
                }
        }

        if checks {
                gen := p.StartMoveGen(Ply()).GenerateChecks().rank()
                for move := gen.NextMove(); move != 0; move = gen.NextMove() {
                        if position := p.MakeMove(move); position != nil {
                                fmt.Printf("%*squix/%s> ply: %d, move: %s\n", Ply()*2, ` `, C(p.color), Ply(), move)
                                moveScore := -position.xSearchQuiescenceInCheck(-beta, -alpha)

                                position.TakeBack(move)
                                if moveScore > bestScore {
                                        if moveScore > alpha {
                                                if moveScore >= beta {
                                                        return moveScore
                                                }
                                                alpha = moveScore
                                        }
                                        beta = moveScore
                                }
                        }
                }
        }
        return bestScore
}

// Quiescence search (in check).
func (p *Position) xSearchQuiescenceInCheck(alpha, beta int) int {
        if p.isRepetition() {
                return 0
        }

        bestScore := Ply() - Checkmate
        if bestScore >= beta {
                return bestScore
        }

        gen := p.StartMoveGen(Ply()).GenerateEvasions().rank()
        for move := gen.NextMove(); move != 0; move = gen.NextMove() {
                if position := p.MakeMove(move); position != nil {
                        fmt.Printf("%*squic/%s> ply: %d, move: %s\n", Ply()*2, ` `, C(p.color), Ply(), move)
                        moveScore := 0
                        if position.isInCheck(position.color) {
                                moveScore = -position.xSearchQuiescenceInCheck(-beta, -alpha)
                        } else {
                                moveScore = -position.xSearchQuiescence(-beta, -alpha, false)
                        }

                        position.TakeBack(move)
                        if moveScore > bestScore {
                                if moveScore > alpha {
                                        if moveScore >= beta {
                                                return moveScore
                                        }
                                        alpha = moveScore
                                }
                                beta = moveScore
                        }
                }
        }

        return bestScore
}