package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Request struct {
	or string
}

func NewRequest() *Request {
	r := Request{}
	return &r
}

func (r *Request) GetMethod(url string) ([]byte, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	req, err := newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = do(client, req, buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func newClient() (*http.Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar:     jar,
		Timeout: 30 * time.Second,
	}
	return client, nil
}

func newRequest(method, path string, values url.Values) (*http.Request, error) {
	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func do(client *http.Client, req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			// give *bytes.Buffer to get raw bytes instead of json decoded string
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)

			// ignore the error caused by an empty response
			if err == io.EOF {
				err = nil
			}
		}
	}

	return resp, nil
}
