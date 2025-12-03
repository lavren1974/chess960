package chess

// A MoveTag represents a notable consequence of a move.
type MoveTag uint16

const (
	// KingSideCastle indicates that the move is a king side castle.
	KingSideCastle MoveTag = 1 << iota
	// QueenSideCastle indicates that the move is a queen side castle.
	QueenSideCastle
	// Capture indicates that the move captures a piece.
	Capture
	// EnPassant indicates that the move captures via en passant.
	EnPassant
	// Check indicates that the move puts the opposing player in check.
	Check
	// inCheck indicates that the move puts the moving player in check and
	// is therefore invalid.
	inCheck
)

// A Move is the movement of a piece from one square to another.
type Move struct {
	s1    Square
	s2    Square
	promo PieceType
	tags  MoveTag
}

// String returns a string useful for debugging.  String doesn't return
// algebraic notation.
func (m *Move) String() string {
	return m.s1.String() + m.s2.String() + m.promo.String()
}

// S1 returns the origin square of the move.
func (m *Move) S1() Square {
	return m.s1
}

// S2 returns the destination square of the move.
func (m *Move) S2() Square {
	return m.s2
}

// Promo returns promotion piece type of the move.
func (m *Move) Promo() PieceType {
	return m.promo
}

// HasTag returns true if the move contains the MoveTag given.
func (m *Move) HasTag(tag MoveTag) bool {
	return (tag & m.tags) > 0
}

func (m *Move) addTag(tag MoveTag) {
	m.tags = m.tags | tag
}

type moveSlice []*Move

// normalize960Castle ensures Chess960 castle moves carry the appropriate tags,
// even if the king remains on its starting square (king-to-rook UCI).
func normalize960Castle(pos *Position, m *Move) {
	if m == nil {
		return
	}
	p := pos.Board().Piece(m.s1)
	if p.Type() != King {
		return
	}
	rook := NewPiece(Rook, p.Color())
	if pos.Board().Piece(m.s2) != rook {
		return
	}
	if m.s2.File() > m.s1.File() {
		m.addTag(KingSideCastle)
	} else {
		m.addTag(QueenSideCastle)
	}
}

func (a moveSlice) find(m *Move) *Move {
	if m == nil {
		return nil
	}
	wantCastle := KingSideCastle
	wantCastleSet := false
	if m.HasTag(KingSideCastle) {
		wantCastle = KingSideCastle
		wantCastleSet = true
	} else if m.HasTag(QueenSideCastle) {
		wantCastle = QueenSideCastle
		wantCastleSet = true
	}
	var fallback *Move
	for _, move := range a {
		if move.String() == m.String() {
			if wantCastleSet && move.HasTag(wantCastle) {
				return move
			}
			if fallback == nil {
				fallback = move
			}
		}
	}
	return fallback
}
