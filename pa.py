import requests        #导入requests包
from bs4 import BeautifulSoup
import re

url = 'https://www.meegoq.com/info31314.html'

# 爬取小说的最新章节，并保存到本地
def spider(url):
    strhtml = requests.get(url)        #Get方式获取网页数据
    text = strhtml.text
    # 获取小说最新章节
    pattern = re.compile('最新章节：<i><a href="//(.+?)"')
    str2 = pattern.search(text)
    newurl = 'https://'+str2.group(1)

    #爬取最新章节并保存到本地
    strhtml = requests.get(newurl)        #Get方式获取网页数据
    text = strhtml.text
    soup = BeautifulSoup(text, "html.parser")
    novel = soup.find_all(class_='content')
    file = open('1.txt','w')
    file.write(str(novel[0]))
    file.close()


spider(url)