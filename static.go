package main

type Token struct {
	FB       string `json:"facebook_token"`
	FBid     string `json:"facebook_page_id"`
	TW       string `json:"tw_token"`
	TWsec    string `json:"tw_sec"`
	TWbot    string `json:"tw_bot_token"`
	TWbotSEC string `json:"tw_bot_sec"`
	TWcon    string `json:"tw_consum"`
	TWconSEC string `json:"tw_consum_sec"`
	TWid     string `json:"twitter_id"`
}

const (
	secertFile        = "secret.json"
	ErrReadingJson    = "Error occurred during reading JSON:"
	ErrPostingFB      = "Error occurred during posting to Facebook:"
	ErrPostingTW      = "Error occurred during posting to Twitter:"
	ErrReadingPost    = "Error occurred during reading post content:"
	ErrPostingTWReply = "Error occurred during posting twitter reply:"
)
