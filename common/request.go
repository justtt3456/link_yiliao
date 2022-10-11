package common

import (
	"io/ioutil"
	"net/http"
)

func QuotesRequest(url string, token string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return bytes
}
