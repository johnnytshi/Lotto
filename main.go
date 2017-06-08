package main

import (
	"fmt"
	"sync"

	"github.com/johnnytshi/lotto/processor"
)

func main() {
	input := []string{"569815571556", "4938532894754", "4938532894754", "1234567"}

	// Use Hashmap of (string, int array) to remove dups
	output := make(map[string][]int)

	// Remove dups
	for _, element := range input {
		output[element] = nil
	}

	// Utilize go routine to parrallelize processing
	var wg sync.WaitGroup // Use WaitGroup to make sure everything is done before exiting the program
	wg.Add(len(output))

	for k := range output {
		go func(k string) {
			defer wg.Done()
			output[k] = processor.Process(k)
			if len(output[k]) != 7 { // Remove unsuccessful ones
				delete(output, k)
			} else {
				fmt.Printf("%s -> %v\n", k, output[k])
			}
		}(k)
	}

	wg.Wait()
}
