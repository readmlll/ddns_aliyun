package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Requests struct {
	COMMON_TIME_OUT float32
}

func (that *Requests) CommonReq(_type string, apiUrl string, data map[string]string, head map[string]string) (string, []byte) {

	switch _type {
	case "GET":
		// URL param
		getData := url.Values{}
		for key, val := range data {
			getData.Set(key, val)
		}

		u, err := url.ParseRequestURI(apiUrl)
		if err != nil {
			fmt.Printf("parse url requestUrl failed, err:%v\n", err)
		}

		u.RawQuery = getData.Encode() // URL encode

		fmt.Println(u.String())

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			fmt.Printf("NewRequest failed, err:%v\n", err)
		}

		// 添加请求头
		req.Header.Add("Content-type", "application/json;charset=utf-8")
		// 添加cookie
		cookie1 := &http.Cookie{
			Name:  "aaa",
			Value: "aaa-value",
		}
		req.AddCookie(cookie1)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("get resp failed, err:%v\n", err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("get resp ioutil failed, err:%v\n", err)
			return "", nil
		}
		fmt.Println(string(b))
		return string(b), b
	}
	return "", nil
}
