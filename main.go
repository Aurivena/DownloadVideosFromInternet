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

func getNextUrl() (string, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer file.Close()

	// The first url in the file
	var first string
	// next urls for write to file
	var nextUrls []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if first == "" {
			first = scanner.Text()
			continue
		}
		nextUrls = append(nextUrls, scanner.Text())
	}

	// clear file
	err = os.WriteFile(pathFile, []byte{}, 0644)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(pathFile, os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// write url in file
	for _, line := range nextUrls {
		_, _ = f.WriteString(line + "\n")
	}

	if first == "" {
		return "", fmt.Errorf("file is empty")
	}

	return first, nil
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
