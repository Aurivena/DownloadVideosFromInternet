package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getNextUrl() (string, error) {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer file.Close()

	// The first url in the file
	var first string
	// Next urls for write to file
	var nextUrls []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if first == "" {
			first = scanner.Text()
			continue
		}
		nextUrls = append(nextUrls, scanner.Text())
	}

	if err = clearFile(nextUrls); err != nil {
		log.Fatal(err)
		return "", err
	}

	if first == "" {
		return "", fmt.Errorf("file is empty")
	}
	return first, nil
}

func clearFile(urls []string) error {
	// Clear file
	err := os.WriteFile(pathFile, []byte{}, 0644)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(pathFile, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write url in file
	for _, line := range urls {
		_, _ = f.WriteString(line + "\n")
	}

	return nil
}
