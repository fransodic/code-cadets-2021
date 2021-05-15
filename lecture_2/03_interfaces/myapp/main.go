package main

import (
	"log"

	"code-cadets-2021/lecture_2/03_interfaces/stacklibfer"
)

func main() {
	stack := stacklibfer.New()

	pushing(stack)
}

func pushing(pusher Pusher) {
	pusher.Push(1)
	pusher.Push(2)
	pusher.Push(3)
	pusher.Push(4)

	// I cant access other methods here (pop)

	log.Println(pusher)
}

type Pusher interface {
	Push(a int)
}
