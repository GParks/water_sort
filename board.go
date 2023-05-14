package solvwatrsort

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

const NUM_SLOTS = 4
const NUM_TUBES = 14

const (
	NC     = iota // 0 --> no color
	RED    = iota
	GREEN  = iota
	BLUE   = iota
	YELLOW = iota
	PURPLE = iota
	ORANGE = iota
)

// corresponding array of "names" -- I'll figure out how to use this later
var color_names = []string{
	"Red",
	"Green",
	"Blue",
	"Yellow",
	"Purple",
	"Orange"}

func colorOfName(s string) int {
	for i, n := range color_names {
		if n == s {
			return i
		}
	}
	return NC
}

type Tube struct {
	slots    [NUM_SLOTS]int // array of ints representing colors
	num_used int
	done     bool
}

func (t Tube) top() int {
	// for i := NUM_SLOTS; i >= 0; i--  {
	// 	if t.slots[i] != NC {
	// 		return t.slots[i]
	// 	}
	// }
	// return NC
	return t.slots[t.num_used]
}

func (t Tube) empty() bool {
	// return t.top() == NC
	return t.num_used == 0
}

func (t Tube) full() bool {
	// return t.slots[NUM_SLOTS-1] != NC
	return t.num_used == NUM_SLOTS
}

func (t Tube) add(c int) {
	t.slots[t.num_used] = c
	t.num_used++
}

func (t Tube) pour() int {
	if t.num_used < NUM_SLOTS {
		t.slots[t.num_used] = NC
	}
	t.num_used--
	return t.slots[t.num_used+1]
}

func (t Tube) Equals(other Tube) bool {
	// if t.num_used != other.num_used {
	// 	return false
	// }
	// for i := 0; i < t.num_used; i++ {
	// 	if t.slots[i] != other.slots[i] {
	// 		return false
	// 	}
	// }
	return t.slots == other.slots
}

type Move struct {
	src int
	dst int
}

type Board struct {
	tubes [NUM_SLOTS]Tube
	moves list.List
}

func loadBoard(fn string) Board, err {
	f, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		err = "couldn't read file \"" + fn + "\""
		return Board{}, err
	}
	s := bufio.NewScanner(f)
	var b Board
	for i := 0; i < NUM_TUBES; i++ {
		s.Scan()
		for j := 0; j < NUM_SLOTS; j++ {
			l := s.Text()
			sls := strings.Split(l, ",")

			for k, sl := range sls {
				c := colorOfName(strings.Trim(sl, " \t"))
				if c != NC { // it should never be "NC" in this context
					b.tubes[i].add(c)
				} else {
					err = "not a color: \"" + sl + "\", line " + k
					return Board{}, err
			}
		}
	}
	f.Close()

	return b
}

func (b Board) valid_moves() list.List {
	var moves list.List
	for fr := 0; fr < NUM_TUBES; fr++ {
		for to := 0; to < NUM_TUBES; to++ {
			if fr != to && !b.tubes[fr].empty() && !b.tubes[to].full() {
				if fr.top() == to.top() {
					// var m Move
					// m.src = fr
					// m.dst = to
					moves.PushBack(m{fr, to})
				}
			}
		}
	}
	return moves
}

// this should *copy* the board, not modify the original
func (b Board) move(m Move) Board {
	r := b
	r.tubes[m.dst].add(r.tubes[m.src].pour())
	return r
}
