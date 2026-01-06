package main

import (
	"fmt"
	"os"
	"strings"
)

// 方案1
//
//	func main() {
//		var s, sep string
//		for i := 1; i < len(os.Args); i++ {
//			s += sep + os.Args[i]
//			sep = " "
//		}
//		fmt.Println(s)
//	}

// 方案2
// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// 方案3
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
