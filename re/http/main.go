package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// resp, err := http.Get("https://lotteryjs.com")
	req, _ := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
