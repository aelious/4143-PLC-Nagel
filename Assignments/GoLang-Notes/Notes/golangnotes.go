// Download Go: https://go.dev/doc/install
// Ensure proper installation by entering this in the terminal: go version
// Download proper plugins if using VS Code. Switch to VS Code if not using VS Code.

// Set up your environment:
// ~~~~~~~~~~~~~~~~~~~~~~~~
// Make sure you are in the correct directory in the terminal and run go mod url
// Note: url can be any link, I've used the associated repo for conciseness

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
