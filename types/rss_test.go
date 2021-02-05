package types

import (
	"fmt"
	"testing"
)

func TestRegisterRSS(t *testing.T) {
	rss := RSSHandler{}
	items := rss.GetNewItems("https://sandbox.facepunch.com/rss/news")
	for _, item := range items {
		fmt.Println(item.Image)
		fmt.Println(item.TimeID)
	}
}
