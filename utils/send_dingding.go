package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SendDing struct {
	Url string
	Token string
	Body []byte
}

func NewSendDing(token string, body []byte) *SendDing {
	return &SendDing{
		Url:   "https://oapi.dingtalk.com/robot/send",
		Token: token,
		Body:  body,
	}
}

func (ding *SendDing) Send() ([]byte, error) {
	req, err := http.NewRequest("POST", ding.Url, bytes.NewReader(ding.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	query := req.URL.Query()
	query.Add("access_token", ding.Token)
	req.URL.RawQuery = query.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c := &http.Client{
		Transport: tr,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(respBody))
	return respBody, nil
}