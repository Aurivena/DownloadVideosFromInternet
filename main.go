package main

import (
	"fmt"
	"log"
	"sync"
)

// The directory where the videos will be saved
const dir = "/home/answer/Видео"

// The path where the linked file is stored with urls
const pathFile = "urls.txt"

func main() {
	wg := sync.WaitGroup{}
	out := make(chan string)
	for {

		url, err := getNextUrl()
		if err != nil {
			if err.Error() == "file is empty" {
				break
			}
			log.Fatal(err)

		}

		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			download(out, u)
		}(url)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	for val := range out {
		fmt.Println(val)
	}
}
