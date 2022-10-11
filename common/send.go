package common

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	POST_METHOD_1 = "application/x-www-form-urlencoded" //"user=123&a=123"
	POST_METHOD_2 = "application/json"
)

//post请求
func PostJson(
	urlPath string,
	postParam []byte,
	header map[string]string,
	cookies map[string]string,
	postMethod int8,
) (body []byte, err error) {
	client := &http.Client{}
	//设置请求参数
	//postForm请求
	if postMethod == 3 {
		postData := url.Values{}
		var tmp map[string]string
		err = json.Unmarshal(postParam, &tmp)
		if err != nil {
			return nil, errors.New("postParam:解析出错")
		}
		for k, v := range tmp {
			postData.Set(k, v)
		}
		postParam = []byte(postData.Encode())

	}

	payload := strings.NewReader(string(postParam))
	//创建POST请求参数
	req, err := http.NewRequest(http.MethodPost, urlPath, payload)
	if err != nil {
		return nil, err
	}
	//urlencode请求
	if postMethod == 1 || postMethod == 3 {
		req.Header.Add("content-type", POST_METHOD_1)
	}
	//json流请求
	if postMethod == 2 {
		req.Header.Add("content-type", POST_METHOD_2)
	}

	//设置头部信息
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	//设置cookie
	if cookies != nil {
		for k, v := range cookies {
			c := &http.Cookie{
				Name:  k,
				Value: v,
			}
			req.AddCookie(c)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
