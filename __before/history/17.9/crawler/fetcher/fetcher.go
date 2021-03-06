package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Kirk-Wang/Hello-Gopher/history/17.9/crawler/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 反爬，控制下速度，100ms 一个 tick
var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	// resp, err := http.Get()
	<-rateLimiter // 100 个人会抢同一个
	log.Printf("Fetching url %s", url)
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 确认页面编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 拿到前 1024 个 bytes
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
