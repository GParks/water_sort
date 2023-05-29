package game

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"os"
	"strings"
)

const NUM_SLOTS = 4
const NUM_TUBES = 14

/**
 * there must be some way to create an enum, but I don't know how -- yet
 */
const (
	NC        = iota // 0 --> no color
	RED       = iota
	GREEN     = iota
	OLIVE     = iota
	LTBLUE    = iota
	BLUE      = iota
	YELLOW    = iota
	PURPLE    = iota
	ORANGE    = iota
	BROWN     = iota
	PINK      = iota
	NEONGREEN = iota
	GRAY      = iota
)

// corresponding array of "names" -- I'll figure out how to use this later
var color_names = []string{
	"*** No Color ***",
	"Red",
	"Green",
	"Olive",
	"LiteBlue",
	"Blue",
	"Yellow",
	"Purple",
	"Orange",
	"Brown",
	"Pink",
	"NeonGreen",
	"Gray"}

// string (i.e. "Green", "Purple") to int conversion
// only used internally, not exported
func colorOfName(s string) int {
	for i, n := range color_names {
		if n == s {
			return i
		}
	}
	return NC
}

// int to string
func ColorName(c int) string {
	return color_names[c]
}

type Tube struct {
	slots    [NUM_SLOTS]int // array of ints representing colors
	num_used int
	done     bool
}

func (t Tube) Top() int {
	retColor := NC
	if t.num_used > 0 {
		retColor = t.slots[t.num_used-1]
	}
	return retColor
}

func (t Tube) IsEmpty() bool {
	return t.num_used == 0
}

func (t Tube) IsFull() bool {
	return t.num_used == NUM_SLOTS
}

func (t Tube) IsDone() bool {
	return t.done
}

func (t *Tube) add(c int) {
	t.slots[t.num_used] = c
	t.num_used++
	// fmt.Printf("  \t  dbg: Added %s to this tube, now %d 'used'\n", color_names[c], t.num_used)
	if t.IsFull() {
		for i := 0; i < NUM_SLOTS-1; i++ {
			if t.slots[i] != c {
				return
			}
		}
		// IFF this tube is full and all the same color, then it's done
		t.done = true
	}
}

func (t Tube) debug() {
	fmt.Printf("  \t  dbg: (a) tube has used %d slot(s)\n", t.num_used)
	for i := NUM_SLOTS - 1; i >= 0; i-- {
		fmt.Printf("   %s, ", color_names[t.slots[i]])
	}
	if t.done {
		fmt.Println("  Done!")
	} else {
		fmt.Println("  not done.")
	}
}

func (t *Tube) pour() int {
	var retColor = NC
	if t.num_used > 0 {
		retColor = t.slots[t.num_used-1]
		t.num_used--
		t.slots[t.num_used] = NC
	} // else {
	// 	return NC
	//}
	return retColor
}

func (t Tube) Equals(other Tube) bool {
	// if t.num_used != other.num_used {
	//     return false
	// }
	// for i := 0; i < t.num_used; i++ {
	//     if t.slots[i] != other.slots[i] {
	//         return false
	//     }
	// }
	return t.slots == other.slots
}

type Move struct {
	Src   int
	Dst   int
	Color int
}

type Board struct {
	tubes [NUM_TUBES]Tube
	moves list.List
}

func LoadBoard(fn string) (Board, error) {
	f, e := os.Open(fn)
	if e != nil {
		fmt.Println("Error opening file: ", e)
		return Board{}, errors.New("couldn't read file \"" + fn + "\"")
	}
	s := bufio.NewScanner(f)
	// var b Board
	b := Board{}
	for i := 0; i < NUM_TUBES; i++ {
		if !s.Scan() {
			/**
			 *  It's okay if there are fewer lines than Tubes; the last (two) will be empty

			 * return Board{}, errors.New("not enough lines in th file; only " + strconv.Itoa(i))
			 */
			fmt.Println("    LoadBoard: fewer lines, only ", i, "; exiting here")
			break
		}
		l := s.Text()
		sls := strings.Split(l, ",")
		for j, sl := range sls {
			// fmt.Printf("\t dbg: sl = %s , ", sl)
			c := colorOfName(strings.Trim(sl, " \t"))
			if c != NC { // it should never be "NC" in this context
				b.Add(i, c)
			} else {
				var err = errors.New(fmt.Sprintf("not a color: \"%s\", line %d", sl, j+1))
				return Board{}, err
			}
			if j == NUM_SLOTS-1 {
				break
			}
		}
		// fmt.Println("  dbg: done with line ", i)
	}
	// fmt.Println("  dbg: done loading board from file")
	f.Close()

	return b, nil
}

func (b *Board) Add(t, c int) {
	// fmt.Printf("      dbg: (Board) adding %s to tube %d\n", color_names[c], t)
	b.tubes[t].add(c)
}

func (b Board) String() string {
	var s string
	s = ""
	for j := NUM_SLOTS - 1; j >= 0; j-- {
		s += "|"
		for i := 0; i < NUM_TUBES/2; i++ {
			if c := b.tubes[i].slots[j]; c == NC {
				s += fmt.Sprintf("     _     |")
			} else {
				s += fmt.Sprintf(" %9s |", color_names[c])
			}
		}
		s += "\n"
	}
	s += "\n"
	for j := NUM_SLOTS - 1; j >= 0; j-- {
		s += "|"
		for i := NUM_TUBES / 2; i < NUM_TUBES; i++ {
			if c := b.tubes[i].slots[j]; c == NC {
				s += fmt.Sprintf("     _     |")
			} else {
				s += fmt.Sprintf(" %9s |", color_names[c])
			}
		}
		s += "\n"
	}
	return s
}

func (b Board) Encode() string {
	var s string
	c := "_RGobBYVOWPNa"
	/*
	 *      Red,         Green,    (o)live,
	 *   lite(b)lue,      Blue,    Yellow,
	 * Violet [Purple],  Orange,   bro(w)n,
	 *     Pink,        NeonGreen, gr(a)y
	 */
	for i := 0; i < NUM_TUBES; i++ {
		for j := 0; j < NUM_SLOTS; j++ {
			n := b.tubes[i].slots[j]
			s += c[n : n+1]
		}
	}
	return s
}

func (b Board) IsBlank() bool {
	for i := 0; i < NUM_TUBES; i++ {
		if !b.tubes[i].IsEmpty() {
			return false
		}
	}
	return true
}

func (b Board) ValidMoves() list.List {
	var moves list.List
	for fr := 0; fr < NUM_TUBES; fr++ {
		for to := 0; to < NUM_TUBES; to++ {
			// fmt.Println("  dbg: checking ", fr, " to ", to)
			if fr != to && !b.tubes[fr].IsEmpty() &&
				!b.tubes[fr].IsDone() &&
				!b.tubes[to].IsFull() {
				if b.tubes[to].IsEmpty() ||
					b.tubes[fr].Top() == b.tubes[to].Top() {
					moves.PushBack(Move{fr, to, b.tubes[fr].Top()})
				}
			}
		}
	}
	/**
	 * I should filter out moves that are "regressing" ...
	 *  (this might be tricky)
	 */
	return moves
}

func (b *Board) SaveMove(m Move) {
	b.moves.PushBack(m)
}

func (b Board) Top(t int) int {
	return b.tubes[t].Top()
}

// this should *copy* the board, not modify the original
func (b Board) DoMove(m Move, quiet bool) Board {
	r := b
	/**
	 * I wished things worked as 'simply' as this (maybe I can do it later)
	 * for r.Add(m.Dst, r.tubes[m.Src].pour()); r.tubes[m.Dst].Top() == r.tubes[m.Src].Top() &&
	 * 	!r.tubes[m.Dst].IsFull(); r.Add(m.Dst, r.tubes[m.Src].pour()) {
	 * }
	 **/

	/**
	 * I regretted trying to add err return here;
	 * currently, all my code assumes that DoMove will succeed
	 * (i.e. the colors match, source and destination are 'valid,' etc.)
	 * if r.tubes[m.Src].Top() != m.Color {
	 * 	fmt.Printf("      DoMove: ERROR: `Top` of tubes[%d] = %s, move color is %s\n",
	 * 		m.Src, color_names[r.tubes[m.Src].Top()], color_names[m.Color])
	 * 	return Board{}, errors.New("DoMove: source tube Top is not the color of the move")
	 * }
	 **/

	r.Add(m.Dst, r.tubes[m.Src].pour())
	// fmt.Println("    DoMove: after pouring:")
	// fmt.Println(r)
	for !r.tubes[m.Src].IsEmpty() && r.tubes[m.Dst].Top() == r.tubes[m.Src].Top() &&
		!r.tubes[m.Dst].IsFull() {
		r.Add(m.Dst, r.tubes[m.Src].pour())
	}
	if !quiet {
		fmt.Printf("      DoMove: all poured:\n%s\n", r)
	}
	r.SaveMove(m)
	return r
}

func (b Board) IsSolved() bool {
	for i := 0; i < NUM_TUBES; i++ {
		if !b.tubes[i].IsDone() {
			return false
		}
	}
	return true
}

func (b Board) Equals(othr Board) bool {
	// fmt.Println("\t  \t  dbg: comparing boards")
	// NOTE:
	// right now (5/28/2023), I could, *theoretically*, just compare the
	// tubes array(s)  (i.e. `b.tubes != ohtr.tubes`)
	// however, (1) I wouldn't be able to debug which tube is different
	// and, (2) I might want to add more fields to the Tube struct later...
	for i := 0; i < NUM_TUBES; i++ {
		if !b.tubes[i].Equals(othr.tubes[i]) {
			// fmt.Printf("\t  \t  dbg: tubes at %d  are not equal\n", i)
			return false
		}
	}
	return true
}

func Equal(b1, b2 Board) bool {
	fmt.Println("\t  \t  dbg: comparing two boards")
	return b1.Equals(b2)
}

func (b Board) NumMoves() int {
	return b.moves.Len()
}
func (b Board) PrintMoves() {
	for e := b.moves.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
