package main

import fb "github.com/huandu/facebook/v2"

func PublishFB() (err error) {
	tkn, id := token.FB, token.FBid
	_, err = fb.Post(id+"/photos", fb.Params{
		"caption":      caption,
		"message":      caption,
		"url":          imageUrl,
		"access_token": tkn,
	})
	return
}
