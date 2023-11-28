package rest

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Call(method string, url string, headerList []string) (*http.Response, int64) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	req.Header.Set("Content-Type", "application/json")
	for _, v := range headerList {
		headerAndValue := strings.Split(v, ":")
		req.Header.Set(headerAndValue[0], headerAndValue[1])

	}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	elapsed := time.Since(start).Milliseconds()

	return resp, elapsed
}
