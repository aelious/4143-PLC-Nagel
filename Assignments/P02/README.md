## P02 - Baby Steps
### Stephanie Nagel
### Description:

Further implements and utilizes various Go tools including importing foreign packages via Github.


### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
| 1 | [imagemod](https://github.com/aelious/4143-PLC-Nagel/tree/main/Assignments/P02/imagemod/)| Directory containing the entirety of our code. |
| 2 | [main.go](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P02/imagemod/main.go) | Contains the executable code for our program. |
| 3 | [go.mod](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P02/imagemod/go.mod) | Encapsulates our module and includes the necessary dependencies.  |
| 4 | [go.sum](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P02/imagemod/go.sum) | Provides a checksum to ensure there are no repeat installation of packages. |
| 5 | [imageManipulator](https://github.com/aelious/4143-PLC-Nagel/tree/main/Assignments/P02/imagemod/imageManipulator)| Directory containing our image manipulator package. |
| 6 | [mustangs.png](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P02/imagemod/mustangs.png) | Image created using the gg library. |
| 7 | [mustangs.jpg](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P02/imagemod/mustangs.jpg) | Image created using gg library to copy pre-existing png image. |

### Instructions:

Ensure you have the latest version of Go installed! Download all files, except mustangs.jpg. Ensure that the gg library is successfully acquired.

In the terminal from the imagemod directory, run the command: go run .\main.go

This will create a copy of your mustangs.png and save it as a jpg and then will draw a fun lil' black rectangle over the top.
