package main

import (
	"testing"

	"github.com/pastelnetwork/goris/ris"
)

func TestImgFromURL(t *testing.T) {
	imageurl := "https://github.com/pastelnetwork/goris/blob/master/myavatar.png?raw=true"
	webpages := false
	results, _ := ris.DefImg(webpages).ImgFromURL(imageurl)
	t.Log(results)
}
