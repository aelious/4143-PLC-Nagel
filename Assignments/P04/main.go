package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	// Sequentially loops through each url and runs one at a time.
	for i, imageUrl := range urls {
		downloadImage(imageUrl, fmt.Sprintf("seq_image%d.png", i+1))
	}
	return
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(url string, i int, wg *sync.WaitGroup) {
	// Decrements currently running concurrent function counter once the
	// function is finished running and returns.
	defer wg.Done()

	downloadImage(url, fmt.Sprintf("con_image%d.png", i+1))
	return
}

func main() {
	// WaitGroup allows for better concurrent processing to ensure each function
	// completes before the program ends.
	var wg sync.WaitGroup
	urls := []string{
		"https://images.unsplash.com/photo-1599474401061-465ef4b5bc05?auto=format&fit=crop&q=80&w=2070&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://images.pexels.com/photos/76957/tree-frog-frog-red-eyed-amphibian-76957.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://images.pexels.com/photos/1293261/pexels-photo-1293261.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://images.unsplash.com/photo-1692003996152-59e60ea85fe5?auto=format&fit=crop&q=80&w=1935&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		"https://cdn.pixabay.com/photo/2023/07/16/09/31/cat-8130334_960_720.jpg",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	for i, url := range urls {
		// Adds to counter of how many concurrent functions are currently running
		wg.Add(1)
		// Initiates goroutine for each url in the url list
		go downloadImagesConcurrent(url, i, &wg)
	}
	// Waits for concurrent functions to complete
	wg.Wait()
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return err
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return err
	}

	// Create a new file to save the image
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return err
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return err
	}

	fmt.Printf("Image successfully downloaded and saved as '%s'\n", filename)
	return nil
}
