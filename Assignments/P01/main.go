package main

import (
	"fmt"

	"github.com/aelious/4143-PLC-Nagel/tree/main/Assignments/P01/mascot"
)

func main() {
	// Calls BestMascot from the mascot package and returns the result to the
	// Println call from the fmt package.
	fmt.Println(mascot.BestMascot())
}
