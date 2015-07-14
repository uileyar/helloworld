// exercise-web-crawler.go
/*
在这个练习中，将会使用 Go 的并发特性来并行执行 web 爬虫。
修改 Crawl 函数来并行的抓取 URLs，并且保证不重复。
*/

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
	craw bool
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		delete(f, url)
		return res.body, res.urls, nil
		/*
			if f[url].craw == false {
				f[url].craw = true
				return res.body, res.urls, nil
			} else {
				return "", nil, fmt.Errorf("again : %s", url)
			}
		*/
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
		false,
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
		false,
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
		false,
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
		false,
	},
}

func CrawlLoop(url string, depth int, fetcher Fetcher, ch chan string) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	//ch <- url
	for _, u := range urls {
		go CrawlLoop(u, depth-1, fetcher, ch)
	}

	return
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {

	CrawlLoop(url, depth, fetcher, ch)

	return
}

func main() {
	ch := make(chan string)

	Crawl("http://golang.org/", 4, fetcher, ch)
}
