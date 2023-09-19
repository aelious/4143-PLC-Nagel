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
*       This program uses an external package, fogleman/gg via github, to
*		generate and manipulate images in fun, funky ways. This small snippet
*		generates a blank png containing only a black rectangle. It also includes
*		a separate constructor that can take a pre-existing image and make a copy
*		of the image.
*
*  Usage:
*       N/A
*
*  Files:
*       main.go     	: driver program
*       go.mod			: necessary modules dependencies to run our go code
*		go.sum 			: maintains checksum to keep packages from reinstalling
*		imagemod		: package containing our image manipulator tools
*       mustangs.png	: our created png image
*		mustangs.jpg	: a copy of our png image stored as a jpg
*
*****************************************************************************/

// imageManipulator/main.go

package main

import (
	"fmt"                                // fmt is only included for error handling
	"imagemod/imagemod/imageManipulator" // include imagemod and saving auto fills our path
)

func main() {
	// Create an ImageManipulator instance of 800px by 600px
	im := imageManipulator.NewImageManipulator(800, 600)

	// Draw a rectangle
	im.DrawRectangle(150, 50, 560, 411)

	// Save the modified image to a file
	im.SaveToFile("mustangs.png")

	// Create an ImageManipulator instance with an existing image - using the path
	newIm, err := imageManipulator.NewImageManipulatorWithImage("mustangs.png")
	// Error handling in case the path couldn't be found.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Save the copied image to a file
	newIm.SaveToFile("mustangs.jpg")
}
