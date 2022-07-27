package quizzes

import (
	"fmt"
	"time"
)

func GoFuncLiteralsInLoop() {
	fmt.Println("Go func literals in loop with counter.")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
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
//	1
//	0
//	2
