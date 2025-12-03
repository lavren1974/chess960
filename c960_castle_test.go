package chess

import "testing"

func Test960CastleAvailableForWhite(t *testing.T) {
    f, err := FEN("qrbnkbrn/pppppppp/8/8/8/8/PPPPPPPP/NNRQKBBR w KQkq - 0 1")
    if err != nil { t.Fatal(err) }
    g := NewGame(UseNotation(UCINotation{}), f)
    seq := []string{"e2e4", "e7e5", "d1g4", "d8e6"}
    for _, m := range seq {
        if err := g.MoveStr(m); err != nil { t.Fatalf("move %s err %v", m, err) }
    }
    found := false
    for _, mv := range g.ValidMoves() {
        if mv.HasTag(KingSideCastle) || mv.HasTag(QueenSideCastle) {
            found = true
        }
    }
    if !found {
        t.Fatalf("expected a castle move, none found; fen %s", g.Position().String())
    }
}
