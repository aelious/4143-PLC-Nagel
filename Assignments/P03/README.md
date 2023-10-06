## P03 - Image Ascii Art 
### Stephanie Nagel
### Description:

Utilize go get to import a package :-)

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
| 1 | [main.go](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P03/main.go) | Main executable of our code. |
| 2 | [downloaded_image.png](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P03/downloaded_image.png) | Image downloaded from the internet. |
| 3 | [pixels_RGB.txt](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P03/pixels_RGB.txt) | Text file containing each of the image's pixels' RGB values.  |
| 4 | [gray_image.png](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P03/gray_image.png) | Image converted to grayscale. |
| 5 | [hello.png](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P03/hello.png)  | Image with "Hello, world!" written over the top.  |
| 6 | [imgMod](https://github.com/aelious/imgMod)   | Repo containing our imgMod Package. |

NOTE: I'm going to skip the inclusion of the go.mod and go.sum files in the directory from here on out as they don't seem essential to the logic of the structure.

### Instructions

Ensure you have go installed before trying these steps >:-(

1. Download [main.go](https://github.com/aelious/4143-PLC-Nagel/raw/main/Assignments/P03/main.go) and save it to a desired directory (Alternatively you can just download and run main.exe)

1. Navigate to the directory that contains main.go in the terminal

1. Run go mod init someName

1. Run go mod tidy and wait for updates and go.sum file creation

1. Run go run ./main.go

1. WATCH THE MAGIC HAPPEN 🤯

I heavily studied Fogleman's gg library to properly structure my files and better understand go.mod and how the main package is used.

You can have multiple main files in one directory using the main package. The only way to differentiate which file is run is by using the specific filename.

Also, go.mod doesn't seem necessary in the package declaration. Your package's files only need to be in the same directory and have the same package name. Note that my package [imgMod](https://github.com/aelious/imgMod) did not require any mod/sum files. Similarly, if you're using a package from the web and it's included somewhere in the code in an import statement, the go.mod file doesn't need to be included with your main files either. You can simply init one, name it whatever, and use go mod tidy and it will pull in the necessary dependencies and create a go.sum file to keep track of their versions.

An example of this use case can be seen in running this file via the instructions. You init your mod file after downloading the go code and then go mod tidy tells it which packages it needs to use go get on.
