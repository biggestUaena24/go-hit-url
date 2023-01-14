package main

import (
	"fmt"
	"net/http"
)

type result struct{
	url string
	status string
}

func hitURL (url string, channel chan result) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400{
		channel <- result{url:url, status:"failed"}
	}else{
		channel <- result{url:url, status:"OK"}
	}

}

func FakeMain() {

	var channel = make(chan result)

	urls := []string{"https://www.airbnb.com/",
	"https://www.google.com/",
	"https://www.amazon.com/",
	"https://www.reddit.com/",
	"https://www.google.com/",
	"https://soundcloud.com/",
	"https://www.facebook.com/",
	"https://www.instagram.com/",
	"https://nomadcoders.co/",}

	for _, url := range(urls){
		go hitURL(url, channel)
	}

	for i:=0;i<len(urls);i++{
		fmt.Println(<-channel)
	}

	
}