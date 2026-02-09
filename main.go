package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	clear       = "\033[2J\033[3J\033[0H"
	defaultDude = "( ._.)"
)

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".dude"
	}
	return filepath.Join(home, ".dude")
}

func main() {
	config := getConfigPath()

	if _, err := os.Stat(config); err != nil {
		os.Create(config)
		os.WriteFile(config, []byte(defaultDude), 0644)
	}

	var userDude string
	var reset bool

	flag.StringVar(&userDude, "dude", "", "set you dude")
	flag.StringVar(&userDude, "d", "", "set you dude")

	flag.BoolVar(&reset, "reset", false, "reset you dude to default (use this flag on it's own)")

	flag.Parse()

	if reset {
		os.WriteFile(config, []byte(defaultDude), 0644)
		return
	}

	if userDude != "" {
		os.WriteFile(config, []byte(userDude), 0644)
	}

	dude, err := os.ReadFile(config)
	if err != nil {
		fmt.Printf("%s\n\r", clear)
		return
	}

	fmt.Printf("%s%s\n\r", clear, dude)
}
