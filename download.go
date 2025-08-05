package main

import (
	"fmt"
	"os/exec"
)

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
