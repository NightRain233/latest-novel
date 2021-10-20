package spider

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
)

//定义新的数据类型
type Spider struct {
	url string
}

//获取全部内容
func (keyword Spider) get_html_all() string {
	resp := keyword.get_html_resp()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

//只获取目录页最新章节部分
func (keyword Spider) get_html_part() string {
	resp := keyword.get_html_resp()
	br := bufio.NewReader(resp.Body)
	var line []byte
	for i := 0; i < 78; i++ {
		line, _, _ = br.ReadLine()
	}
	return string(line)
}

func (keyword Spider) get_html_resp() *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
