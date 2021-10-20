package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"latest-novel/spider"
	"sync"
)

const baseUrl string = "http://www.31xiaoshuo.com"

var chapter_map = map[string]string{}
var urls = [3]string{
	"http://www.31xiaoshuo.com/0/196/",
	"http://www.31xiaoshuo.com/4/4542/",
	"http://www.31xiaoshuo.com/176/176372/",
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"getnovel": GetNovel,
	})
	r.LoadHTMLFiles("template/index.html", "template/bookshelf.html")
	r.GET("/novel/*novelurl", func(c *gin.Context) {
		novelurl := c.Param("novelurl")
		c.HTML(200, "index.html", novelurl)
	})
	r.GET("/bookshelf", func(c *gin.Context) {
		AsyncGetChapter()
		c.HTML(200, "bookshelf.html", chapter_map)
	})
	r.Run(":8080")
}

func AsyncGetChapter() {
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(pos int, url string) {
			chapter := spider.GetLatestChapter(url)
			if chapter != nil {

				chapter_map[fmt.Sprintf("url%d", pos)] = chapter.URL
				chapter_map[fmt.Sprintf("title%d", pos)] = chapter.Title
			}
			wg.Done()
		}(i, url)
	}
	wg.Wait()
}

func GetNovel(novel_url string) interface{} {
	novel := spider.GetNovel(baseUrl + "/" + novel_url)
	return template.HTML(novel)
}
