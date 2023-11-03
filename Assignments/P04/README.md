## P04 - Concurrent Image Downloader
### Stephanie Nagel
### Description:

Implement and compare concurrent downloads versus sequential downloads using Golang. In lieu of channels, I've opted for WaitGroups to ensure proper goroutine outputs, they seemed to fit the overall logic better.

Sequential downloads took 575.5325ms to complete.

Concurrent downloads took 59.3535ms to complete. 

### Files

|   #   | File            | Description                                        |
| :---: | --------------- | -------------------------------------------------- |
| 1 | [main.go](https://github.com/aelious/4143-PLC-Nagel/blob/main/Assignments/P04/main.go) | Main executable of our code. |

The packages used are all built in so there is no need to include a go.mod file for versioning control.

### Instructions

Ensure you have go installed before trying these steps >:-(

1. Download [main.go](https://github.com/aelious/4143-PLC-Nagel/raw/main/Assignments/P04/main.go) and save it to a desired directory (Alternatively you can just download and run main.exe)

1. Navigate to file directory

1. Run go run ./main.go

1. WATCH THE MAGIC HAPPEN 🤯


