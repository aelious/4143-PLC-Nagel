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
*	includesa separate constructor that makes a copy of a pre-existing image. A
*	black rectangle is then drawn over the top of our image.
*
*  Usage:
*       go run .\main.go
*	> generates an edited image of a pre-existing image via its path
*
*  Files:
*       main.go     		: driver program
*       go.mod			: necessary modules dependencies to run our go code
*	go.sum 			: maintains checksum to keep packages from reinstalling
*	imagemManipulator	: contains our imagemod package and tools
*       mustangs.png		: our created png image, now including horses :-)
*	mustangs.jpg		: a copy of our png image stored as a jpg with a fun
*					black rectangle drawn on it :-)
*****************************************************************************/

// imagemod/main.go

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
