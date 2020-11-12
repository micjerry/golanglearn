package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//go test -run=TestConcurrent -race -v gopl.io/ch9/memo
//go test -run=TestSeq -race -v gopl.io/ch9/memo

type M interface {
	Get(key string) (interface{}, error)
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://news.qq.com",
			"https://news.sina.com.cn",
			"http://www.xinhuanet.com",
			"http://www.xinhuanet.com",
			"http://news.ifeng.com",
			"https://news.163.com",
			"http://www.people.com.cn",
			"https://news.cctv.com",
			"https://www.huanqiu.com",
			"https://news.sina.com.cn",
			"https://news.qq.com",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

func sequential(t *testing.T, m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}


func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	concurrent(t, m)
}

func TestSeq(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	sequential(t, m)
}