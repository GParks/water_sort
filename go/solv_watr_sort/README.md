# Water Sort
**Solving the Water Sort game**

### Executing the program
Here is the cmd line I executed "last" &ndash; this is the prog. that has been running for 36 hrs.

    ./solv_watr_sort -q -i 1000000 first.brd | tee test_run_q_1000000.txt | grep "pushing back"

The resulting file was over 200k.


## Introduction

This project had three goals &ndash; and a fourth to document the other three; 

1. to learn what [CoPilot](https://github.com/features/copilot) could do
1. to teach myself the `go` programming language (https://go.dev/)
1. to "solve" a game: [Water Sort Puzzle](https://apps.apple.com/ph/app/water-sort-puzzle/id1514542157)

I used the 3rd in order to accomplish something of the other two &mdash; I always find it 'difficult' to "learn" something new
if I don't have some problem or other to solve.  For me, personally, "learn x, y, or z" is not easy without a plan. <br>
So my plan, usually, as I've gotten wiser, is to "use x, y or z, to solve a (somewhat interesting) problem."  
My brain can stick to something longer if I am *engaged*, and that is more likely with something I 'care' about...

So, my first impression of CoPilot was positive.  As background, I've been a software engineer for over quarter of a century, 
professionally, and much longer if you could my educational career.  At this point, I've used several different development 
environments, up to and including simple text editor and command-line compilers with no *GUI* &ndash; I've used `vi`, 
I've used IDEs that no longer exist, I've used Eclipse, my "favorite" at the moment is still **`Emacs`**. <br>
CoPilot needed [`VS Code`](https://code.visualstudio.com/docs#vscode), so I adapted.  My personal dev work had been stalled, 
so I needed a jump start on that, and this seemed as good a chance as any.

## **Golang** (https://github.com/golang)
Learning <kbd>Go</kbd> was somewhat challenging, even with CoPilot, for two reasons &ndash; learning file/module organization, 
and one 'feature' I wasn't expecting to have trouble with: pointers (and/or "pass by value" vs. "pass by reference")
First, CoPilot was no help organizing my methods into separate files, and putting more than one in the same *module*.
I did various searches, even asking [`ChatGPT`](https://chat.openai.com/), and using <https://you.com/>. <br>
The best answer I found was on [StackOverflow](https://stackoverflow.com/questions/9985559/20188012#20188012). 

### Links/Refs
Among the tabs I still have open, or have used, there are these: 
* ["Return and handle an error"](https://go.dev/doc/tutorial/handle-errors)
* [Standard library](https://pkg.go.dev/std)
  * [container](https://pkg.go.dev/container@go1.20.4)
    * [list](https://pkg.go.dev/container/list@go1.20.4#List)
* ["How to check if a slice contains an element in Golang?"](https://www.tutorialspoint.com/how-to-check-if-a-slice-contains-an-element-in-golang)
* <https://go.dev/doc/code>
* **[Effective Go](https://go.dev/doc/effective_go)**
  * [errors](https://go.dev/doc/effective_go#errors)

* [Go Resources](https://www.golang-book.com/11)
  * ["An Introduction to Programming in Go"](https://www.golang-book.com/books/intro)
* <https://go.dev/learn/>
* [A Tour of Go](https://go.dev/tour/)
  * ["For continued"](https://go.dev/tour/flowcontrol/2)

* [Structures in Go](https://medium.com/rungo/structures-in-go-76377cc106a2)


## CoPilot
As for <kbd>CoPilot</kbd>, well...  it helped me, let me write methods like this:
```go
func (t Tube) add(c int) {
	t.slots[t.num_used] = c
	t.num_used++
}
```
and 
```go 
func (t Tube) pour() int {
	if t.num_used < NUM_SLOTS {
		t.slots[t.num_used] = NC
	}
	t.num_used--
	return t.slots[t.num_used+1]
}
```
<br/>
And I spent far too long figuring out why this didn't actually <em>change</em> anything; here's the working version, now: 

```go
func (t *Tube) add(c int) {
	t.slots[t.num_used] = c
	t.num_used++
	// fmt.Printf("  \t  dbg: Added %s to this tube, now %d 'used'\n", color_names[c], t.num_used)
}
```
and 
```go
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
```
&mdash;never mind the other bug(s) found/fixed, in order to modify the *`struct`* which was created, 
I needed to define these functions to work on a *pointer*!  (Someone who knows `golang` better than I 
could give me a more accurate description, or definition of what I'm really doing, and why)

And that's probably also on me.  I like that CoPilot helped &ndash; the "auto-complete" of functions by hitting tab 
has really been a time saver; it has given me code that is syntactically correct, 
and usually pretty close to the semantics I was after; but there are things that it does not understand. <br>
As always, I need to remember to continue to use the two and a half pound computer situated betweeen my ears.

Nonetheless, **<kbd>CoPilot</kbd>** is something that I can see myself using, and actually enjoying.

## Closer to Completion 
Sunday, 21st May, 2023 &ndash; I've gotten to this point, the program ran for six hours today, and I noticed a couple things to clean up.

Also, I'll need to add more than a little testing (and figure out how to "properly" do that, in <kbd>Go</kbd>. 

First, implementing `IsDone` for both `Tubes` and the `Board` itself; 
maybe also `DoNotMove` for those that are "complete" in and of themselves, 

Additionally, I need to be able to `ShowMoves` when I want to actually _figure out_ how I got 'here'

Later, I'll have to add heuristics, and determine what "closer to done" looks like, and what is a "good" path &ndash; 
then implement a "best-first" search, storing the open list in a "priority queue" or some "ordered," (sorted) data structure.

## "SofA" (State of Affairs) end of May, 2023
With version 1 of `Equals` (for a `Board`), I let this prog. run for over a day -- it took more than a day to get to 400,000 on the "closed list" (1 d. 2 1/2 hrs); 
due to how I structured dumping out the current closed list, I'm letting it run past 468,751 (150* 5^5 +1).

### What I think of <kbd><b>Go</b></kbd>
At this point, disappointed by two features that are missing -- *enums* and no operator-overloading (so I can't define 'b1 == b2' to be other than "all fields of the struct (or orray) match...) <br>
Those two features aside, I like that it compiles to a binary; learning cmd-line arg. processing wasn't hard (though there was a minor gotcha: processing stops at the first non-matched, so `prog -q first.brd` isn't the same as `prog first.brd -q`); <br>
I'd like to learn more about *interfaces*, but it seems to match <kbd>Clojure</kbd> in my mind, how `def-protocol` behaves&mdash;though without strict declaration, so that could make some debugging of design "tricky," if a developer were to miss (or *mis-type*) a func. name, for example
