package task

import (
	"fmt"
	"main/filter"
	"sync"
)

func ProcessImagesWaitGroup(files []string, imgFilter string, dstFolder string) {
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			err := filter.ApplyFilter(filename, imgFilter, dstFolder)
			if err != nil {
				fmt.Println(err)
			}
		}(file)
	}

	wg.Wait()
}

func ProcessImagesChannel(files []string, filter string, dstFolder string, poolSize int) {
	jobs := make(chan string, poolSize)
	results := make(chan string)

	for i := 0; i < poolSize; i++ {
		go worker(jobs, results, filter, dstFolder)
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	for i := 0; i < len(files); i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(jobs <-chan string, results chan<- string, imgFilter string, dstFolder string) {
	for filename := range jobs {
		err := filter.ApplyFilter(filename, imgFilter, dstFolder)
		if err != nil {
			results <- fmt.Sprintf("Error processing: %s - %s", filename, err)
		} else {
			results <- fmt.Sprintf("Processed: %s", filename)
		}
	}
}
