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
	defer resp.Body.Close()
	elapsed := time.Since(start).Milliseconds()

	return resp, elapsed
}

func RenderUrl(url string, isHttps bool) string {
	if isHttps {
		return "https://" + url
	} else {
		return url
	}
}

func MethodI2S(v int) string {
	switch v {
		case 0:
			return "GET"
		case 1:
			return "POST"
		case 2:
			return "PUT"
		case 3:
			return "PATCH"
		case 5:
			return "DELETE"
		default:
			return "GET"
	}
}