package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(imageUrl)
	checkError(err)
	defer resp.Body.Close()
	_, err = io.Copy(file, resp.Body)
	defer file.Close()
	checkError(err)
}

func buildFileName() {
	fileUrl, err := url.Parse(imageUrl)
	checkError(err)
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
