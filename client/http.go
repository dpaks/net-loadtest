package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	timeout = 30
)

type httpClient struct {
	Client  http.Client
	header  map[string]string
	qparams map[string]string
}

func HTTPClient() httpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   timeout * time.Second,
			KeepAlive: timeout * time.Second,
		}).DialContext,
	}
	cl := httpClient{
		Client:  http.Client{Transport: transport, Timeout: time.Second * timeout},
		header:  make(map[string]string),
		qparams: make(map[string]string),
	}

	return cl
}

func (cl *httpClient) SerializeData(d interface{}) []byte {
	body, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

func (cl *httpClient) AddQueryParam(k, v string) {
	cl.qparams[k] = v
}

func (cl *httpClient) AddHeader(k, v string) {
	cl.header[k] = v
}

func (cl *httpClient) Post(url string, body []byte) (res *http.Response, err error) {
	return cl.send(http.MethodPost, url, body)
}

func (cl *httpClient) Delete(url string) (res *http.Response, err error) {
	return cl.send(http.MethodDelete, url, []byte{})
}

func (cl *httpClient) Get(url string) (res *http.Response, err error) {
	return cl.send(http.MethodGet, url, []byte{})
}

func (cl *httpClient) send(method, url string, body []byte) (res *http.Response, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range cl.header {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for k, v := range cl.qparams {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()

	res, err = cl.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (cl *httpClient) ParseResponse(res *http.Response) (code int, resBody []byte, err error) {
	if res == nil {
		err = fmt.Errorf("Empty response body")
		return
	}

	code = res.StatusCode

	if res.Body != nil {
		resBody, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
	}

	return
}

func (cl *httpClient) DeserializeData(data []byte, out interface{}) (err error) {
	err = json.Unmarshal(data, out)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
