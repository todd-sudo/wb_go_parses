package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/wb_go/internal/parser"
)

func main() {
	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	t1 := time.Now().UnixNano()
	parser.CreateTasks()
	t2 := time.Now().UnixNano()
	dt := float64(t2-t1) / 1000000.0
	fmt.Println(dt)
}
