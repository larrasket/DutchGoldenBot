package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"

	"github.com/dghubble/oauth1"
)

func PublishTW(ids ...string) (id string, err error) {
	config := oauth1.NewConfig(token.TWcon, token.TWconSEC)
	tkn := oauth1.NewToken(token.TW, token.TWsec)
	if len(ids) > 0 {
		tkn = oauth1.NewToken(token.TWbot, token.TWbotSEC)
	}
	httpClient := config.Client(oauth1.NoContext, tkn)
	b := &bytes.Buffer{}
	form := multipart.NewWriter(b)
	fw, err := form.CreateFormFile("media", fileName)
	if err != nil {
		panic(err)
	}
	opened, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fw, opened)
	if err != nil {
		panic(err)
	}
	_ = form.Close()
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

	values := url.Values{"status": {""},
		"media_ids": {mid},
	}
	if len(ids) > 0 {
		values = url.Values{"status": {"@" + token.TWid + " " + art},
			"media_ids":             {mid},
			"in_reply_to_status_id": {ids[0]},
		}
	}
	resp, err = httpClient.PostForm("https://api.twitter.com/1.1/statuses/update.json", values)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	ms := make(map[string]interface{})
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&ms)
	if err != nil {
		return
	}
	return fmt.Sprint(ms["id_str"]), nil
}
