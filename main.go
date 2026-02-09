package main

import "fmt"

const (
	clear = "\033[2J\033[3J\033[0H"
	dude  = " <( ._.)>"
)

func main() {
	fmt.Printf("%s%s\n\r", clear, dude)
}
