package main

import (
	"github.com/gin-gonic/gin"

	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"md5": MD5,
	})
	r.LoadHTMLFiles("index.html", "bookshelf.html")
	r.GET("/html", func(c *gin.Context) {
		novelurl := c.Query("novelurl")
		c.HTML(200, "index.html", novelurl)
	})
	r.GET("/bookshelf", func(c *gin.Context) {
		c.HTML(200, "bookshelf.html", "")
	})
	r.Run(":8080")
	// MD5(url1)
}

const (
	baseUrl string = "https://www.vbiquge.com/"
	url1    string = "https://www.vbiquge.com/8_8088/"     //神话版三国
	url2    string = "https://www.vbiquge.com/76_76099/"   //特拉福买家俱乐部
	url3    string = "https://www.vbiquge.com/102_102727/" //镇妖博物馆
)

func MD5(novel_url string) interface{} {
	last_novel_url := Parse1(novel_url)
	novel := Parse(last_novel_url)
	return template.HTML(novel)
}

//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}

//定义 Spider get的方法
func (keyword Spider) get_html_header() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)

}

//根据首页url获取最新章节title和url
func Parse(url string) string {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	//创建excel文件
	f, err := os.Create("./haha31.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	spider := &Spider{url, header}
	html := spider.get_html_header()
	f.WriteString(html)

	pattern := `<title>(.*?)- 新笔趣阁</title>`
	rp := regexp.MustCompile(pattern)
	find_txt2 := rp.FindAllStringSubmatch(html, -1)
	title := "<h1>" + find_txt2[0][1] + "</h1>"

	pattern = `<div id="content">(.*?)</div>`
	rp = regexp.MustCompile(pattern)
	find_txt2 = rp.FindAllStringSubmatch(html, -1)

	novel := find_txt2[0][0]
	f.WriteString(title + novel)

	return title + novel

}

//根据章节url获得章节内容
func Parse1(url string) string {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	spider := &Spider{url, header}
	html := spider.get_html_header()

	pattern := `p>最新章节：<a href="(.*?)" target="_blank">`
	rp := regexp.MustCompile(pattern)
	find_txt2 := rp.FindAllStringSubmatch(html, -1)
	last_novel_url := baseUrl + find_txt2[0][1]

	return last_novel_url

}
