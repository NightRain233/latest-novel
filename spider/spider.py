import requests        #导入requests包
from bs4 import BeautifulSoup
import re,time

# def cost_time(fn):
# 	def _wrapper(*args,**kwargs):
# 		start = time.perf_counter()
# 		a = fn(*args, **kwargs)
# 		print("%s cost %s  second"%(fn.__name__,time.perf_counter() - start))
# 		return a 
# 	return _wrapper

class Chapter:
	def __init__(self,url,title) :
		self.title = title
		self.url = url

#定义新的数据类型


def GetLatestChapter(url):
	# s = requests.session()
	# s.keep_alive = False # 关闭多余连接

	strhtml = requests.get(url)        #Get方式获取网页数据
	text = strhtml.text
	pattern = re.compile('<p>最&nbsp;&nbsp;&nbsp;&nbsp;新：<a href="(.*?)">(.*?)</a>')
	str2 = pattern.search(text)
	
	novel = Chapter(str2.group(1),str2.group(2))
	return novel

def GetNovel(url):
	# s = requests.session()
	# s.keep_alive = False # 关闭多余连接

	strhtml = requests.get(url)        #Get方式获取网页数据
	text = strhtml.text	

	pattern = re.compile('<title>(.*?)_31小说网</title>')
	str2 = pattern.search(text)
	title = "<h1>" + str2.group(1) + "</h1>"

	pattern = re.compile('(<div id="content">[\s\S]+?)<div')
	str2 = pattern.search(text)
	novel = str2.group(1)

	pattern = re.compile('link1\(\);</script>([\s\S]+?)<script>link2')
	str2 = pattern.search(text)
	bottom =  "</div> " + str2.group(1).replace('href="', 'href="/novel')

	return title+novel+bottom


		
url = "http://www.31xiaoshuo.com/0/196/"

if __name__ == '__main__':
	print(1)
	s1 = GetLatestChapter(url)
	print( s1.url,s1.title)

