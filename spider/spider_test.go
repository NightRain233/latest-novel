package spider

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetHTML(t *testing.T) {
	url := "https://www.31xiaoshuo.com/0/196/"
	spider := &Spider{url}
	html := spider.get_html_all()
	ioutil.WriteFile("test.txt", []byte(html), 0644)
	chapter := GetLatestChapter(url)
	fmt.Printf("\nchapter:%+v\n\n", chapter)
	// GetNovel(url)
	// fmt.Println(html)
}
