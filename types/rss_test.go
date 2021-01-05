package types

import (
	"fmt"
	"testing"
)

func TestRegisterRSS(t *testing.T) {
	rss := RSSHandler{}
	items := rss.GetNewItems("https://housepetscomic.com/feed")
	for _, item := range items {
		fmt.Println(item.Image)
	}
}
