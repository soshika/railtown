package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"railtown/utills/responses"
	"strings"
)

type Request struct {
	URL     string          `json:"url"`
	Method  string          `json:"method"`
	Payload *strings.Reader `json:"payload"`
}

func (r *Request) Init(url string, method string, payload *strings.Reader) {
	r.Method = method
	r.URL = url
	r.Payload = payload
}

func (r *Request) AddToURL(query string) {
	r.URL = r.URL + fmt.Sprintf("?key=cbafa1ae884641f08f9134613230201&aqi=yes&q=%s", query)
}

func (r *Request) GET() ([]byte, *responses.Response) {
	url := r.URL
	method := r.Method

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, responses.NewBadRequestError("could not send request to server", "please try again later", http.StatusBadRequest)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, responses.NewBadRequestError("could not send request to server", "please try again later", http.StatusBadRequest)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, responses.NewBadRequestError("could not send request to server", "please try again later", http.StatusBadRequest)
	}
	return body, nil
}
