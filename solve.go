package main

import (
	"container/list"
	"fmt"
	"os"
	"solv_watr_sort/game"
)

func pop(l *list.List) game.Board {
	f := l.Front()
	if f != nil {
		b := l.Remove(f).(game.Board)
		return b
	} else {
		return game.Board{}
	}
}

func maybeAdd(a, b *list.List, n game.Board) {
	fmt.Println("    maybeAdd: ")
	// fmt.Println(n)
	idx := 0
	for e := a.Front(); e != nil; e = e.Next() {
		if n == e.Value.(game.Board) {
			fmt.Println("      already in open list at", idx)
			return
		}
		idx++
	}
	fmt.Println("    after checking first list (open), checking second (closed)")
	idx = 0
	for e := b.Front(); e != nil; e = e.Next() {
		if n == e.Value.(game.Board) {
			fmt.Println("      already in open list at", idx)
			return
		}
		idx++
	}
	fmt.Println("    not in either list, adding to 1st (open)")
	a.PushFront(n)
}

func main() {
	fmt.Println("starting; number of cmd line args: ", len(os.Args))
	if len(os.Args) < 2 {
		panic("specify at least one cmd. line arg, the filename")
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
	}
	fmt.Println("\t 'second' = " + os.Args[1])
	init, err := game.LoadBoard(os.Args[1])
	if err != nil {
		fmt.Println("Oops!  " + err.Error())
		return
	}

	open := list.List{}
	var closed list.List

	open.PushBack(init)
	for next := pop(&open); !next.IsBlank(); next = pop(&open) {
		fmt.Println("\n\t after popping next, the open list contains", open.Len(), "items, and "+
			"the closed list contains", closed.Len(), "items")
		fmt.Println("next = ")
		fmt.Println(next)
		ms := next.ValidMoves()
		fmt.Printf("    %d (new) valid moves\n", ms.Len())
		for m := ms.Front(); m != nil; m = m.Next() {
			x := m.Value.(game.Move)
			fmt.Printf("    next move is %s from %d to %d\n", game.ColorName(x.Color),
				x.Src, x.Dst)
			n := next.DoMove(x)
			// TO DO: Coming Soon!!!
			// if n.IsSolved() {
			// 	fmt.Println("Solved!")
			// 	return
			// }
			maybeAdd(&open, &closed, n)
		}
		closed.PushBack(next)
		fmt.Println("\n\t after `pushing back` next, the open list contains", open.Len(), "items, and "+
			"the closed list contains", closed.Len(), "items")
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
