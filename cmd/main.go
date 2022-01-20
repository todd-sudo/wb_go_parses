package main

import (
	"fmt"
	// "runtime"
	// "time"

	"github.com/wb_go/internal/parser"
)

// func main() {
// 	numcpu := runtime.NumCPU()
// 	fmt.Println("NumCPU", numcpu)
// 	t1 := time.Now().UnixNano()

// 	go parser.SaveProduct(1, 100)

// 	t2 := time.Now().UnixNano()
// 	dt := float64(t2-t1) / 1000000.0
// 	fmt.Println(dt)
// }

func main() {
	c := make(chan int) // Делает канал для связи
	for i := 0; i < 5; i++ {
		go parser.SleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // Получает значение от канала
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}
