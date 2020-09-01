package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

/*
 * 目标：获取并打印所有城市第一页用户的详细信息
 */
func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	/*
	 * bufio.Reader的Peek，的确是会影响下面的resp.Body的。
	 * 我们用bufio.Reader包装后，针对bufio.Reader的Peek不会影响bufio.Reader的Read
	 * 但是bufio.Reader肚子里的那个io.Reader，就没法保证了。
	 */
	bodyReader := bufio.NewReader(resp.Body)
	// 确认编码
	e := determineEncoding(bodyReader)
	// 转换编码
	// utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)

}

// 确认页面编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 拿到前 1024 个 bytes
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
