package spider

import (
	"fmt"
	"testing"
)

func TestGetHTML(t *testing.T) {
	url := "http://www.31xiaoshuo.com/0/196/"
	spider := &Spider{url}
	html := spider.get_html_part()
	fmt.Println(html)
}
