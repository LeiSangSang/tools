package tools

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func CurlPost(url, params string, types int) ([]byte, error) {
	client := &http.Client{}
	urls := url
	req, err := http.NewRequest("POST", urls, strings.NewReader(params))
	if err != nil {
		return nil, err
	}
	switch types {
	case 1:
		req.Header.Set("Content-Type", "application/json")
	case 2:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func CurlGet(url string) ([]byte, error) {
	client := &http.Client{}
	var body []byte
	var req *http.Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ = ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func CurlGetFile(url string, filePath string) error {
	client := &http.Client{}
	var req *http.Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}
