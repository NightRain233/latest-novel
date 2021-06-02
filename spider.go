package main

import (
	// "github.com/gin-gonic/gin"
	// "html/template"
	// "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html") // 解析模板，返回 Template
	t.Execute(w, "<h1>Hello World!</h1>")     // 执行模板，并将其传递给 w
}

func parseString(w http.ResponseWriter, r *http.Request) {
	tmpl1 := `<!DOCTYPE html> <html>
        <head>
            <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
            <title>Go Web Programming</title>
        </head>
        <body>`
	p := Parse("https://www.vbiquge.com/8_8088/9152335.html")
	tmpl2 := `</body> 
    </html>`
	s := tmpl1 + p + tmpl2
	t := template.New("index.html")
	t.Parse(s)
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/template", parseString)
	http.ListenAndServe(":8080", nil)
	// s := Parse("https://www.vbiquge.com/8_8088/9152335.html")
	// fmt.Println(s)
}

func MD5(in string) string {
	return in + "hello"
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
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
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

	// url := "https://www.vbiquge.com/8_8088/9152335.html"
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
