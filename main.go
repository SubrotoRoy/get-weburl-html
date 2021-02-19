package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		contents1, err := download("http://example.com")
		if err != nil {
			log.Println("Error encountered while downloading :", err)
		}
		fileName1 := "example.com" + ".html"
		exampleFile, err := createFile(contents1, fileName1)

		defer exampleFile.Close()
		wg.Done()
	}()

	go func() {
		contents2, err := download("http://google.com")
		if err != nil {
			log.Println("Error encountered while downloading :", err)
		}
		fileName2 := "google.com" + ".html"
		googleFile, err := createFile(contents2, fileName2)

		defer googleFile.Close()
		wg.Done()
	}()

	wg.Wait()
}

func download(url string) ([]byte, error) {

	var contents []byte
	resp, err := http.Get(url)
	if err != nil {
		return contents, err
	}

	return ioutil.ReadAll(resp.Body)
}

func createFile(contents []byte, fileName string) (*os.File, error) {

	err := os.WriteFile(fileName, contents, 0666)
	if err != nil {
		return nil, err
	}
	return os.Open(fileName)
}
