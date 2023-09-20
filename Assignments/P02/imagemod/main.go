/*****************************************************************************
*
*  Editor:           Stephanie Nagel
*  Email:            aeliousx@gmail.com
*  Label:            P02
*  Title:            Baby Steps :-)
*  Course:           CMPS 4143
*  Semester:         Fall 2023
*
*  Description:
*       This program uses an external library, gg, provided by Fogleman via github. 
*	It generates and manipulates images in fun, funky ways. This small snippet
*	generates a blank png containing only a black rectangle. It also includes
*	a separate constructor that makes a copy of a pre-existing image.
*
*  Usage:
*       N/A
*
*  Files:
*       main.go     	: driver program
*       go.mod		: necessary modules dependencies to run our go code
*	go.sum 		: maintains checksum to keep packages from reinstalling
*	imagemod	: package containing our image manipulator tools
*       mustangs.png	: our created png image
*	mustangs.jpg	: a copy of our png image stored as a jpg
*
*****************************************************************************/

// imageManipulator/main.go

package main

import (
	"fmt"                       // fmt is only included for error handling
	"imagemod/imageManipulator" // adding package name and saving instantly fills path!
)

func main() {
	// Create an ImageManipulator instance with an existing image - using the path
	newIm, err := imageManipulator.NewImageManipulatorWithImage("mustangs.png")
	// Error handling in case the path couldn't be found.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newIm.DrawRectangle(150, 50, 560, 411)
	// Save the copied image to a file
	newIm.SaveToFile("mustangs.jpg")
}
