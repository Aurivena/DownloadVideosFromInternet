package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

// The directory where the videos will be saved
const dir = ""

// The path where the linked file is stored with urls
const pathFile = ""

func main() {
	urls := getUrls()

	wg := sync.WaitGroup{}
	wg.Add(len(urls))
	out := make(chan string)
	for _, url := range urls {
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

func getUrls() []string {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls
}

func download(out chan<- string, url string) {
	cmd := exec.Command("yt-dlp", "-f", "bv*+ba", "--merge-output-format", "mp4", url)
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	out <- string(output)
}
