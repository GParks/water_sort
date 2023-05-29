package main

import (
	"container/list"
	"flag"
	"fmt"
	"solv_watr_sort/game"
)

var pfQuiet = flag.Bool("q", false, "quiet mode")
var bDebug bool = false

func pop(l *list.List) game.Board {
	f := l.Front()
	if f != nil {
		b := l.Remove(f).(game.Board)
		return b
	} else {
		return game.Board{}
	}
}

func maybeAdd(a, b *list.List, n game.Board) bool {
	if !*pfQuiet {
		fmt.Println("    maybeAdd: ")
	}
	// fmt.Println(n)
	idx := 0
	for e := a.Front(); e != nil; e = e.Next() {
		r := e.Value.(game.Board)
		if n.Equals(r) {
			if !*pfQuiet {
				fmt.Println("      already in open list at", idx)
			}
			if bDebug {
				fmt.Printf(" \t  \t (%s == %s) [a]\n", n.Encode(), r.Encode())
			}
			return false
		} else {
			if bDebug {
				fmt.Printf(" \t  \t (%s != %s) [a]\n", n.Encode(), r.Encode())
			}
		}
		idx++
	}
	if !*pfQuiet {
		fmt.Println("    after checking first list (open), checking second (closed)")
	}
	idx = 0
	/* what am I missing here ? */
	for e2 := b.Front(); e2 != nil; e2 = e2.Next() {
		r := e2.Value.(game.Board)
		if n.Equals(r) {
			if !*pfQuiet {
				fmt.Println("      already in closed list at", idx)
			}
			if bDebug {
				fmt.Printf(" \t  \t (%s == %s) [b]\n", n.Encode(), r.Encode())
			}
			return false
		} else {
			if bDebug {
				fmt.Printf(" \t  \t (%s != %s) [b]\n", n.Encode(), r.Encode())
			}
		}
		idx++
	}
	if !*pfQuiet {
		fmt.Println("    not in either list, adding to 1st (open)")
		fmt.Printf(" \t  \t adding %s \n", n.Encode())
	}
	a.PushFront(n)
	return true
}

func printList(l list.List, terse bool) {
	for e := l.Front(); e != nil; e = e.Next() {
		if terse {
			fmt.Printf(" \t  %s\n", e.Value.(game.Board).Encode())
		} else {
			fmt.Println(e.Value)
			fmt.Println()
		}
	}
}

func main() {
	// fmt.Println("starting; number of cmd line args: ", len(os.Args))
	// if len(os.Args) < 2 {
	// 	panic("specify at least one cmd. line arg, the filename")
	// var v game.Board
	// fmt.Println(v)
	// fmt.Println("  initializing...")
	//
	// tt := game.Tube{}
	// tt.Debug()
	// // tt.done = true
	// tt.Add(game.RED)
	// // fmt.Println("  num_slots now ", tt.num_used)
	// tt.Debug()
	//
	// t1 := game.Board{}
	// fmt.Println(t1)
	// fmt.Println("  adding RED to tube 0")
	// t1.Add(0, game.RED)
	// t1.Add(0, game.BLUE)
	// fmt.Println(t1)
	// return
	// }
	// var fn string
	// flag.StringVar(&fn, "f", "", "filename")

	var iter int
	var bVerbose bool
	flag.IntVar(&iter, "i", 5, "number of iterations")
	flag.BoolVar(&bDebug, "d", false, "debug mode")
	flag.BoolVar(&bVerbose, "v", false, "verbose mode")

	flag.Parse()

	// fmt.Println("\t 'second' = " + os.Args[1])
	fn := flag.Arg(0)
	fmt.Println("\t \"zeroth\" arg = " + fn)

	if *pfQuiet {
		fmt.Println("  quiet mode!")
	} else {
		fmt.Println("  not quiet mode!")
	}

	init, err := game.LoadBoard(fn)
	if err != nil {
		fmt.Println("Oops!  " + err.Error())
		return
	}

	open := list.List{}
	var closed list.List

	gate := 150

	open.PushBack(init)
	for next := pop(&open); !next.IsBlank(); next = pop(&open) {
		if !(*pfQuiet) {
			fmt.Println("\n\t after popping next, the open list contains", open.Len(), "items, and "+
				"the closed list contains", closed.Len(), "items")
			if bVerbose {
				fmt.Println("next = ")
				fmt.Println(next)
			}
		}
		ms := next.ValidMoves()
		// fmt.Printf("    %d (new) valid moves\n", ms.Len())
		cntMoves := 0
		for m := ms.Front(); m != nil; m = m.Next() {
			x := m.Value.(game.Move)
			if !*pfQuiet {
				fmt.Printf("    next move is %s from %d to %d\n", game.ColorName(x.Color),
					x.Src, x.Dst)
			}
			n := next.DoMove(x, *pfQuiet)
			// TO DO: Coming Soon!!!
			// if n.IsSolved() {
			// 	fmt.Println("Solved!")
			// 	return
			// }
			if maybeAdd(&open, &closed, n) {
				cntMoves++
			}

			if n.IsSolved() {
				fmt.Printf("\n\t Solved!\n\n")
				fmt.Println(n)
				n.PrintMoves()
			}
			if cntMoves == 0 {
				fmt.Println("	no moves from here")
				fmt.Println(n)
			}
		}
		closed.PushBack(next)

		fmt.Println("\n\t after `pushing back` next, the open list contains", open.Len(), "items, and",
			"the closed list contains", closed.Len(), "items")
		// if (closed.Len() % 2500) == 0 {
		if closed.Len() > gate {
			fmt.Println(" ** closed list contains", closed.Len(), "items")
			printList(closed, !bVerbose)
			fmt.Println("\t open list:")
			printList(open, !bVerbose)
			gate *= 5
		}
		if iter--; iter == 0 {
			fmt.Println(" ** closed list contains", closed.Len(), "items")
			printList(closed, !bVerbose)
			fmt.Println("\t open list:")
			printList(open, !bVerbose)
			break
		}
	}

	// // fmt.Println(init)
	// ms := init.ValidMoves()
	// fmt.Println("  num valid moves: ", ms.Len())
	// for m := ms.Front(); m != nil; m = m.Next() {
	// 	fmt.Println("    trying next move...")
	// 	n := init.DoMove(m.Value.(game.Move))
	// 	fmt.Println(n)
	// }
}
