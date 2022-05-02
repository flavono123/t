package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// func NewRequest(method, url string, body io.Reader) (*Request, error)
	req, err := http.NewRequest("GET", "http://flavono123.github.io", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "owner")

	// type Client struct {
	// 	 Transport RoundTripper
	//	 CheckRedirect func(req *Request, via []*Request) error
	//   Jar CookieJar
	//	 Timeout time.Duration
	// }
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// func ReadAll(r io.Reader) ([]byte, error)
	// - As of Go 1.16, this function simply calls io.ReadAll.
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) // bytes to string
	fmt.Println(str)
}
