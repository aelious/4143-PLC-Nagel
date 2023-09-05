// Download Go: https://go.dev/doc/install
// Ensure proper installation by entering this in the terminal: go version

// Set up your environment:
// ~~~~~~~~~~~~~~~~~~~~~~~~
// Step 1: From the root directory of your environment, run: go work init
// This will create a go.work module
// Step 2: You will need to create a go.mod module inside of each directory that will require
// go. To do this, navigate one level up from the directory and run: go mod init directoryName
// Note: no .\ or "" will be used in Step 2 ALSO, ensure proper Go directory name, no spaces! >:-(
// Step 3: You will now want to run go work use .\fileName\ on any folder that will contain
// your go code.

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
