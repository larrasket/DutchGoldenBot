package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"

	// "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	fb "github.com/huandu/facebook/v2"
)

var token Token

type MediaUpload struct {
	MediaId int `json:"media_id"`
}

func main() {
	err := ReadToken()
	if err != nil {
		fmt.Fprintf(os.Stderr, ErrReadingJson)
		return
	}
	PublishTW("file.png")
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
func PublishFB(image_url, caption string) (err error) {
	tkn, id := token.FB, token.FBid
	_, err = fb.Post(id+"/photos", fb.Params{
		"caption":      "..",
		"message":      caption,
		"url":          image_url,
		"access_token": tkn,
	})
	return
}
func PublishTW(file string) {
	config := oauth1.NewConfig(token.TWcon, token.TWconSEC)
	token := oauth1.NewToken(token.TW, token.TWsec)
	httpClient := config.Client(oauth1.NoContext, token)
	b := &bytes.Buffer{}
	form := multipart.NewWriter(b)
	fw, err := form.CreateFormFile("media", file)
	if err != nil {
		panic(err)
	}
	opened, err := os.Open("logo.png")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, opened)
	if err != nil {
		panic(err)
	}
	form.Close()
	resp, err :=
		httpClient.Post("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image",
			form.FormDataContentType(), bytes.NewReader(b.Bytes()))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer resp.Body.Close()
	m := &MediaUpload{}
	_ = json.NewDecoder(resp.Body).Decode(m)
	mid := strconv.Itoa(m.MediaId)
	resp, err = httpClient.PostForm("https://api.twitter.com/1.1/statuses/update.json",
		url.Values{"status": {"Post the status!"}, "media_ids": {mid}})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
