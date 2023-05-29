package game

import (
	"testing"
)

func TestCompare(t *testing.T) {
	b, e := LoadBoard("../first.brd") // this makes certain assumptions about the existence of that file
	if e != nil {
		t.Errorf("couldn't load boardL %s", e)
		return
	}
	if b.Encode() != "VWBNONYYaGVBONbGPaoabNOBoOWPPoYbGbBPRVRRVYoGRaWW________" {
		t.Error("board didn't load correctly")
	}

	b12 := b.DoMove(Move{3, 12, GREEN}, true).DoMove(Move{10, 12, GREEN}, false)
	b2 := b.DoMove(Move{10, 12, GREEN}, true)
	b3 := b2.DoMove(Move{3, 12, GREEN}, false)
	if !b3.Equals(b12) {
		t.Error("Boards are not equal")
	}
	/**
	 * I wish this worked!
	 * if b3 == b12 {
	 * 	t.Log("b3 == b12")
	 * } else {
	 * 	t.Error("Boards are not equal")
	 * }
	 **/

}
