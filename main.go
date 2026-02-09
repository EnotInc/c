package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	clear       = "\033[2J\033[3J\033[0H"
	defaultDude = "( ._.)"
)

func main() {
	var userDude string
	var reset bool

	flag.StringVar(&userDude, "dude", "", "set you dude")
	flag.StringVar(&userDude, "d", "", "set you dude")

	flag.BoolVar(&reset, "reset", false, "reset you dude to default (use this flag on it's own)")

	flag.Parse()

	if reset {
		os.WriteFile("dude.txt", []byte(defaultDude), 0644)
		return
	}

	if userDude != "" {
		os.WriteFile("dude.txt", []byte(userDude), 0644)
	}

	dude, err := os.ReadFile("dude.txt")
	if err != nil {
		fmt.Printf("%s\n\r", clear)
		return
	}

	fmt.Printf("%s%s\n\r", clear, dude)
}
