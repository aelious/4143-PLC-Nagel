/*****************************************************************************
*
*  Author:           Stephanie Nagel
*  Email:            aeliousx@gmail.com
*  Label:            P01
*  Title:            Run a Go Program
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*       This program was created following a GoLang tutorial. It includes
* 		several of the basics of running go.
*
*  Usage:
*       N/A
*
*  Files:
*       main.go     	: driver program
*       go.mod			: necessary dependencies to run go
*		mascot.go		: package containing our BestMascot function
*       mascot_test.go  : tests output from BestMascot to ensure proper code
*****************************************************************************/

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
