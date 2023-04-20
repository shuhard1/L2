package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type PassThru struct {
	total int64 // Общее количество переданных байтов
}

// Write переопределяет базовый метод Write io.Writer.
func (pt *PassThru) Write(p []byte) (n int, err error) {
	b := len(p)
	pt.total += int64(b)
	return b, nil
}

func main() {
	var url, fileName string

	fmt.Println("Enter url: ")
	fmt.Scan(&url)

	fileName = strings.Split(url, "/")[len(strings.Split(url, "/"))-1]
	fmt.Println("Downloading", url, "to", fileName)

	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	traf := PassThru{0}
	quit := make(chan bool)
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("downloaded", traf.total, "bytes")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	n, err := io.Copy(output, io.TeeReader(response.Body, &traf))
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	quit <- true

	fmt.Println(n, "bytes downloaded.")
}
