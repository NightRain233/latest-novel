package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"latest-novel/conf"
	"latest-novel/spider"
	"sync"
)

const baseUrl string = "http://www.31xiaoshuo.com"

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
		fmt.Printf("\nnovel:%+v\n\n", AsyncGetChapter())
		c.HTML(200, "bookshelf.html", AsyncGetChapter())
	})
	r.Run(":4040")
}

func AsyncGetChapter() []conf.Novel {
	var wg sync.WaitGroup
	novels := conf.GetNovels()
	for i := 0; i < len(novels); i++ {
		wg.Add(1)
		go func(pos int, novel []conf.Novel) {
			if chapter := spider.GetLatestChapter(novels[pos].URL); chapter != nil {
				novels[pos].Pos = pos + 1
				novel[pos].ChapterURL = chapter.URL
				novel[pos].Title = chapter.Title
			}
			wg.Done()
		}(i, novels)
	}
	wg.Wait()
	return novels
}

func GetNovel(novel_url string) interface{} {
	novel := spider.GetNovel(baseUrl + "/" + novel_url)
	return template.HTML(novel)
}
