package zocket

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func downloadFile(url string, c chan []byte) {
	resp, err := http.Get(url)
	if err != nil {
		c <- nil
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- nil
		return
	}

	c <- data
}

func main() {
	urls := []string{
		"https://example.com/file1.txt",
		"https://example.com/file2.txt",
		"https://example.com/file3.txt",
	}

	c := make(chan []byte)

	for _, url := range urls {
		go downloadFile(url, c)
	}

	for range urls {
		data := <-c
		if data == nil {
			fmt.Println("Error downloading file")
		} else {
			fmt.Println("Successfully downloaded file with length:", len(data))
		}
	}
}
