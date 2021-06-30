package spider

import (
	// "os"
	"regexp"
	"strings"
)

type Chapter struct {
	Title string
	URL   string
}

//根据小说首页url获得最新章节url和标题
func GetLatestChapter(url string) Chapter {
	spider := &Spider{url}
	html := spider.get_html_part()

	pattern := `<p>最&nbsp;&nbsp;&nbsp;&nbsp;新：<a href="(.*?)">(.*?)</a>`
	rp := regexp.MustCompile(pattern)
	find_txt2 := rp.FindAllStringSubmatch(html, -1)

	last_novel_url := find_txt2[0][1]
	chapter := Chapter{
		URL:   last_novel_url,
		Title: find_txt2[0][2],
	}
	return chapter
}

//根据最新章节url获取小说内容
func GetNovel(url string) string {
	spider := &Spider{url}
	html := spider.get_html_all()

	pattern := `<title>(.*?)_31小说网</title>`
	rp := regexp.MustCompile(pattern)
	find_txt2 := rp.FindAllStringSubmatch(html, -1)
	title := "<h1>" + find_txt2[0][1] + "</h1>"

	pattern = `(<div id="content">[\s\S]+?)<div` //多行匹配
	rp = regexp.MustCompile(pattern)
	find_txt2 = rp.FindAllStringSubmatch(html, -1)

	novelurl := find_txt2[0][1]

	//前后章
	pattern = `link1\(\);</script>([\s\S]+?)<script>link2`
	rp = regexp.MustCompile(pattern)
	find_txt2 = rp.FindAllStringSubmatch(html, -1)
	bottom := "</div> " + strings.ReplaceAll(find_txt2[0][1], `href="`, `href="/novel`)

	return title + novelurl + bottom

}
