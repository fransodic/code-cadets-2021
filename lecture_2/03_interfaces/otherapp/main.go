package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/03_interfaces/stacklibfer"
	"code-cadets-2021/lecture_2/03_interfaces/stacklibfoi"
)

func main() {
	stack1 := stacklibfer.New()

	pushPopPrint(stack1, []int{1, 2, 3, 4})

	stack2 := stacklibfoi.New()

	pushPopPrint(stack2, []int{1, 2, 3, 4})
}

func pushPopPrint(stack Stack, numbers []int) {
	for _, num := range numbers {
		stack.Push(num)
	}

	for {
		num, ok := stack.Pop()
		if !ok {
			fmt.Println("Done")
			break
		}

		fmt.Println("Removed from stack:", num)
	}
}

type Stack interface {
	Push(a int)
	Pop() (int, bool)
}
