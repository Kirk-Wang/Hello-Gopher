
package fetcher

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func Fetch(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}