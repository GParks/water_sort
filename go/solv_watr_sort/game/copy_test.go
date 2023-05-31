package game

import (
	"testing"
)

func (ck Board) check_initial_conditions(t *testing.T) {
	/**
	 * I really ought not be using hard-coded tube numbers here
	 */
	if !ck.tubes[0].IsFull() {
		t.Error("tube 0 should be full")
	}
	// all the tubes, thru 11, should be full; I'm only checking a few
	if !ck.tubes[2].IsFull() {
		t.Error("tube 2 should be full")
	}
	if !ck.tubes[5].IsFull() {
		t.Error("tube 5 should be full")
	}
	if !ck.tubes[9].IsFull() {
		t.Error("tube 5 should be full")
	}
	if !ck.tubes[12].IsEmpty() {
		t.Error("tube 12 should be empty")
	}
	if !ck.tubes[NUM_TUBES-1].IsEmpty() {
		t.Errorf("tube %d should be empty", NUM_TUBES-1)
	}

	if ck.Top(3) != GREEN {
		t.Error("tube 3 should have green on top")
	}
}

func TestCopyCtor(t *testing.T) {
	b, e := LoadBoard("../first.brd") // this makes certain assumptions about the existence of that file
	if e != nil {
		t.Errorf("couldn't load boardL %s", e)
		return
	}
	t.Log("Loaded board:")
	t.Log(b)
	b.check_initial_conditions(t)
	t.Log("checked initial conditions")

	ms := b.ValidMoves()
	if ms.Len() != 24 {
		t.Error("there should be 24 valid moves")
	}

	t2 := b.DoMove(Move{3, 12, GREEN}, true)

	b.check_initial_conditions(t)

	if t2.tubes[3].IsFull() {
		t.Error("tube 3 should no longer be full")
	}
	if t2.tubes[12].IsEmpty() {
		t.Error("tube 12 shouldn't be empty (after the move)")
	}

	if t2.Top(12) != GREEN {
		t.Error("tube 12 should have green on top, after moving")
	}
}
