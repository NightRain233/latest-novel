package spider

import (
	"fmt"
	// "io/ioutil"
	"testing"
)

func TestGetHTML(t *testing.T) {
	url := "https://www.31xiaoshuo.com/0/196/69766211.html"
	// spider := &Spider{url}
	// html := spider.get_html_all()
	// ioutil.WriteFile("test.txt", []byte(html), 0644)
	html := GetNovel(url)
	fmt.Println(html)
}
