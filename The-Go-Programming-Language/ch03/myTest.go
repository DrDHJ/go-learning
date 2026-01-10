package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	fmt.Printf("\n====\n")

	ascill := 'a'
	unicode := '国'
	newLine := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascill)
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newLine)

	fmt.Printf("\n====\n")

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}
