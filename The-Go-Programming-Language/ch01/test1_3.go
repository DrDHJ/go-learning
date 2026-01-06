package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, seq := "", ""
	start := time.Now()
	for _, arg := range os.Args {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
	end := time.Since(start).Seconds()
	fmt.Println("循环：", end)
	start = time.Now()

	fmt.Println(strings.Join(os.Args, " "))
	end = time.Since(start).Seconds()
	fmt.Println("Join：", end)

}
