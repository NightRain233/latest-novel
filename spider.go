package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "fmt"
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
	r.LoadHTMLFiles("index.html")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})
	r.Run(":8080")
}

func MD5() interface{} {
	url := "https://www.vbiquge.com/8_8088/9152347.html"
	novel := Parse(url)
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

	// url := "https://www.vbiquge.com/8_8088/9152347.html"
	spider := &Spider{url, header}
	html := spider.get_html_header()
	// fmt.Println(html)

	pattern := `<title>(.*?)- 新笔趣阁</title>`
	rp := regexp.MustCompile(pattern)
	find_txt2 := rp.FindAllStringSubmatch(html, -1)
	title := "<h1>" + find_txt2[0][1] + "</h1>"

	pattern = `<div id="content">(.*?)</div>`
	rp = regexp.MustCompile(pattern)
	find_txt2 = rp.FindAllStringSubmatch(html, -1)

	novel := find_txt2[0][0]
	// f.WriteString(html)

	return title + novel

}
