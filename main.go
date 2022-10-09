package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var token Token

type MediaUpload struct {
	MediaId int `json:"media_id"`
}

// var file string
var art string
var caption string
var (
	fileName string
	imageUrl string
)

func main() {
	err := ReadToken()
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrReadingPost)
		return
	}
	lines, err := ReadLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrPostingTW, err)
		return
	}
	imageUrl = lines[0]
	caption = strings.Join(lines[2:], "\n")
	art = lines[len(lines)-1]
	if strings.Contains(art, "Via.") {
		art = lines[len(lines)-2]
	}
	DownloadImage()
	err = PublishFB()
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrPostingFB, err)
		return
	}
	id, err := PublishTW()
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrPostingTW, err)
		return
	}
	os.Remove(fileName)
	imageUrl = lines[1]
	DownloadImage()
	_, err = PublishTW(id)
	os.Remove(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, ErrPostingTWReply, err)
		return
	}

}

func DownloadImage() {
	buildFileName()
	file := createFile()
	putFile(file, httpClient())
}
func ReadToken() (err error) {

	jsonFile, err := os.Open(secertFile)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	values, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(values, &token)
	return
}

func ReadLines() ([]string, error) {
	file, err := os.Open("post")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
