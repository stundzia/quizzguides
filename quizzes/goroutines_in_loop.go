package quizzes

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printIt(i *int) {
	fmt.Println(*i)
	wg.Done()
}

func GoroutinesInLoop() {
	fmt.Println("Goroutine scheduling in loop with counter.")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printIt(&i)
	}
	wg.Wait()
}

// Guess the output:
// A)
//	0
//	1
//	2
// B)
//	2
//	2
//	2
// C)
//	3
//	3
//	3
// D)
//	0xc000018030
//	0xc000018030
//	0xc000018030
//	(or similar addresses)
