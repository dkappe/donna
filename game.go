package lape

import (
        `bytes`
        `fmt`
	`math`
	`regexp`
)

const (
        MIN = float64(-math.MaxInt16)
        MAX = float64(math.MaxInt16)
        MATE = float64(math.MinInt32/10)
)

type Game struct {
	pieces	[64]Piece
	players	[2]*Player
        attacks *Attack
        current int
}

func (g *Game)Initialize() *Game {
        g.players[0] = new(Player).Initialize(g, WHITE)
        g.players[1] = new(Player).Initialize(g, BLACK)
        g.attacks = new(Attack).Initialize()
        g.current = WHITE

        return g
}

func (g *Game) Setup(white, black string) *Game {
	re := regexp.MustCompile(`\W+`)
	whitePieces, blackPieces := re.Split(white, -1), re.Split(black, -1)
	return g.SetupSide(whitePieces, 0).SetupSide(blackPieces, 1)
}

func (g *Game) SetupSide(moves []string, color int) *Game {
	re := regexp.MustCompile(`([KQRBN]?)([a-h])([1-8])`)

	for _, move := range moves {
		arr := re.FindStringSubmatch(move)
		if len(arr) == 0 {
			fmt.Printf("Invalid move %s for %c\n", move, C(color))
			return g
		}
		name, col, row := arr[1], arr[2][0]-'a', arr[3][0]-'1'

		var piece Piece
		switch name {
		case `K`:
			piece = King(color)
		case `Q`:
			piece = Queen(color)
		case `R`:
			piece = Rook(color)
		case `B`:
			piece = Bishop(color)
		case `N`:
			piece = Knight(color)
		default:
			piece = Pawn(color)
		}
		g.Set(int(row), int(col), piece)
	}
	return g
}

func (g *Game)Set(row, col int, piece Piece) *Game {
        g.pieces[Index(row, col)] = piece

        return g
}

func (g *Game)Get(row, col int) Piece {
        return g.pieces[Index(row, col)]
}

func (g *Game)SetInitialPosition() *Game {
        return g.Setup(`Kg1,Qh1,Bh8`, `Kg8,Rf8,f7,g6,h7`)
        return g.Setup(`Kh1,Ra7,Rc7,Ba8`, `Kh8`)
        return g.Setup(`Kh1,h2,g2,Qh4,Bf6,g5,g4,d4`, `Kg8,Rf8,f7,g6,h7,c8`)
        return g.Setup(`Kh1,g2,h2,Nh6,Qe6`, `Kh8,Rf8,g7,h7`)
        return g.Setup(`Kh1,Ra6,Rb5`, `Kh7`)
        return g.Setup(`Kh1,Ra1`, `Kg8,f7,g7,h7`)

        return g.Setup(`Kg1,f2,g2,h2`, `Kg8,Ra1`)
        return g.Setup(`Kg1,f3,e2,e3`, `Kh3,Ra1`)
        return g.Setup(`d2,f3,g2,Rf2,Kg1`, `Kg3,Ra1`)
        return g.Setup(`Ra1,Nb1,Bc1,Qd1,Ke1,Bf1,Ng1,Rh1,a2,b2,c2,d2,e2,f2,g2,h2`, `Ra8,Nb8,Bc8,Qd8,Ke8,Bf8,Ng8,Rh8,a7,b7,c7,d7,e7,f7,g7,h7`)
        return g.Setup(`a3,Bb4,a5,c3,e7,Kh2`, `a7,a5,b6,Bc7,Kg8`)
        return g.Setup(`a2,Ra3,b3,a7,Kg1`, `d4,Rc4,c3,c5,Bb6,Kg8`)
}

func (g *Game)Search(depth int) (best *Move) {
        position := new(Position).Initialize(g, g.pieces, g.current)
        moves := position.Moves(g.current)
        estimate := MIN

        if len(moves) > 0 {
                for i, move := range moves {
                        score := -position.MakeMove(g, move).Score(depth*2-1, g.current, MIN, MAX)
                        fmt.Printf("  %d/%d: %s for %s, score is %.2f\n", i+1, len(moves), move, C(g.current), score)
                        if score >= estimate {
                                if -score == MATE { // Just pick the move if it mates.
                                        return move
                                }
                                estimate = score
                                best = move
                                fmt.Printf("  New best move for %s is %s\n\n", C(g.current), best)
                        }
                }
        } else if position.IsCheck(g.current) {
                fmt.Printf("Checkmate for %s\n", C(g.current))
        } else {
                fmt.Printf("Stalemate\n") // TODO
        }
	return
}

func (g *Game)String() string {
	buffer := bytes.NewBufferString("  a b c d e f g h\n")
	for row := 7;  row >= 0;  row-- {
		buffer.WriteByte('1' + byte(row))
		for col := 0;  col <= 7; col++ {
			index := Index(row, col)
			buffer.WriteByte(' ')
			if piece := g.pieces[index]; piece != 0 {
				buffer.WriteString(piece.String())
			} else {
				buffer.WriteString("\u22C5")
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}
