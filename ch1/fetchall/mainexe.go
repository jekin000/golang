//fetchall read many urls and calc their response time
package main

//exe 1.11, data from: http://alexa.chinaz.com/Global/
/*
[root@cb7dcbbbc335 fetchall]#go run main.go http://google.com http://youtube.com http://facebook.com http://baidu.com  http://yahoo.com http://amazon.com http://wikipedia.org http://qq.com http://google.co.in http://twitter.com http://live.com http://msn.com http://sina.com.cn http://yahoo.co.jp http://google.co.jp http://linkedin.com http://weibo.com http://bing.com
0.12s      81 http://baidu.com
0.68s   12495 http://google.com
0.92s   13212 http://google.co.in
0.96s  238308 http://qq.com
1.06s   11375 http://google.co.jp
1.47s    3592 http://live.com
1.52s   54499 http://msn.com
1.62s   18586 http://yahoo.co.jp
2.02s  398600 http://youtube.com
2.09s   48126 http://linkedin.com
2.59s  652391 http://facebook.com
2.87s   94341 http://weibo.com
3.35s    6400 http://amazon.com
3.94s  567970 http://sina.com.cn
6.10s  115238 http://bing.com
7.71s   75874 http://wikipedia.org
20.88s  520486 http://yahoo.com
Get http://twitter.com: EOF
30.55s elapsed
*/

//exe 1.10
/*
go run mainexe.go http://google.com http://youtube.com http://facebook.com http://baidu.com  http://yahoo.com http://amazon.com http://wikipedia.org http://qq.com http://google.co.in http://twitter.com http://live.com http://msn.com http://sina.com.cn http://yahoo.co.jp http://google.co.jp http://linkedin.com http://weibo.com http://bing.com
0.11s      81 http://baidu.com
0.11s      81 http://baidu.com
0.36s   11327 http://google.co.jp
0.36s   11368 http://google.co.jp
0.37s   13218 http://google.co.in
0.37s   12476 http://google.com
0.38s   12502 http://google.com
0.45s   13198 http://google.co.in
0.53s  115233 http://bing.com
0.53s    3592 http://live.com
0.54s  115233 http://bing.com
0.57s   18585 http://yahoo.co.jp
0.65s   18585 http://yahoo.co.jp
0.91s    3592 http://live.com
1.02s   54495 http://msn.com
1.09s   48233 http://linkedin.com
1.26s   48233 http://linkedin.com
1.46s   54523 http://msn.com
1.78s  369951 http://youtube.com
1.82s  376215 http://youtube.com
1.86s  652395 http://facebook.com
1.93s  651921 http://facebook.com
2.00s   93321 http://weibo.com
2.22s  238366 http://qq.com
2.35s    6400 http://amazon.com
2.35s    6400 http://amazon.com
2.42s   93321 http://weibo.com
2.46s  238366 http://qq.com
3.74s   75874 http://wikipedia.org
4.10s   75874 http://wikipedia.org
7.03s  567892 http://sina.com.cn
7.04s  567892 http://sina.com.cn
19.69s  524469 http://yahoo.com
Get http://twitter.com: EOF
Get http://twitter.com: EOF
31.03s  524470 http://yahoo.com
31.03s elapsed
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a goroutine,and comunicate it by ch
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s:%v", url, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
