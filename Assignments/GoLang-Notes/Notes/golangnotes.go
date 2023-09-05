// This folder is just going to have my GoLang notes, I'm uploading them to keep all of
// the Go stuff in one place and to reference changes in the future. If you're
// struggling with any of this stuff, feel free to reach out.

// Download Go: https://go.dev/doc/install
// Ensure proper installation by entering this in the terminal: go version
// Download proper plugins if using VS Code. Switch to VS Code if not using VS Code.

// Set up your environment:
// ~~~~~~~~~~~~~~~~~~~~~~~~
// Make sure you are in the correct directory in the terminal and run go mod url
// Note: url can be any link, I've used the associated repo for conciseness

// Okay yeah so go work and go mod are kinda weird, I'll need to read more about them :-)

// Running a program in Go:
// ~~~~~~~~~~~~~~~~~~~~~~~~
// You have 2 options for running go files. You can compile them and create an executable
// or you can compile them and run the program without creating and executable.

// Running without an executable:
// go run fileName.go

// Running with an executable:
// From the command line inside the desired directory:
// go build fileName.go
// After the executable is created:
// fileName.exe

// Packages are like directories or libraries in Go.
// package main tells the compiler that this is the file that will be executed!
package main

// Fmt is an I/O formatting package!
import "fmt"

// Function declaration in Go looks very similar to C++
func main() {
	// Here we are calling the function Println from the package fmt and passing it our greeting.
	fmt.Println("Hello World.")
}

// So to run this program without creating an executable, I navigate to the directory it is in
// on the command line and run: go run .\golangnotes.go (using tab to save time :-))
// To run this program while creating an executable, you do the same steps but run: go build .\golangnotes.go
// Then run the executable by typing .\golangnotes.exe in the terminal! Ezpz

// *** NOTE: You must have the project you're working on as the opened workspace in VS Code or you
// will get a plethora of errors. The go compiler is reading go.mod and when you have multiple conflicting
// go.mod files, it confuses the system. I believe this can be remedied with go work, but I'll need to
// read more about it
