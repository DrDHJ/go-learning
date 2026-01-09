package main

import (
	"The-Go-Programming-Language/ch02/money"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	len := len(os.Args)
	if len < 2 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printConversion(input.Text())
		}
	} else {
		for _, s := range os.Args[1:] {
			printConversion(s)
		}
	}
}

func printConversion(s string) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf:%v\n", err)
		os.Exit(1)
	}
	d := money.Dollar(v)
	y := money.Yuan(v)
	fmt.Printf("%s = %s, %s = %s\n",
		d, money.DToY(d),
		y, money.YToD(y),
	)
}
