from flask import Flask, render_template
from spider.spider import *
import threading


app = Flask(__name__)



baseurl = "http://www.31xiaoshuo.com/"
chapter_dict = dict()
urls = [
	"http://www.31xiaoshuo.com/0/196/",
	"http://www.31xiaoshuo.com/4/4542/",
	"http://www.31xiaoshuo.com/176/176372/",
]


@app.route('/', methods=['GET'])
def bookshelf():
    AsyncGetChapter()
    print(chapter_dict)
    return render_template('bookshelf.html',novels=chapter_dict)

@app.route('/novel/<path:novelurl>', methods=['GET'])
def novel(novelurl):
    # return novelurl
    return render_template('index.html',novel=GetNovel(baseurl+novelurl))

def getchapter(pos):
    chapter_dict[pos] = GetLatestChapter(urls[pos])

def cost_time(fn):
	def _wrapper(*args,**kwargs):
		start = time.perf_counter()
		a = fn(*args, **kwargs)
		print("%s cost %s  second"%(fn.__name__,time.perf_counter() - start))
		return a 
	return _wrapper

@cost_time
def AsyncGetChapter(): 
    # for pos in range(3):
    #     chapter_dict[pos] = GetLatestChapter(urls[pos])
    threads = []
    for i in range(3):
        t = threading.Thread(target=getchapter,args=(i,))
        threads.append(t)

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()




if __name__ == '__main__':
    app.run(debug=True)