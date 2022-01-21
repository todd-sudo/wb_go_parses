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
	countPage := parser.GetCountPage("/zhenshchinam/odezhda/bryuki-i-shorty")
	countGor := parser.CheckCountPage(countPage)

	for i := 0; i <= countGor; i++ {
		go parser.SaveProduct(1, 100)
	}

	// fmt.Println()
	// time.Sleep(2 * time.Second)
	t2 := time.Now().UnixNano()
	dt := float64(t2-t1) / 1000000.0
	fmt.Println(dt)
}
