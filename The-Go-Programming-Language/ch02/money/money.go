// Package money provides types and functions for handling
package money

import "fmt"

type Dollar float64
type Yuan float64

func (d Dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (y Yuan) String() string {
	return fmt.Sprintf("¥%.2f", y)
}

func DToY(d Dollar) Yuan {
	return Yuan(d * 7.0)
}

func YToD(y Yuan) Dollar {
	return Dollar(y / 7.0)
}
