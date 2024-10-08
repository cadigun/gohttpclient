package easyhttp

import (
	"net/http"

	"github.com/cadigun/goeasyclient/api"
)

type EasyHttpBuilder struct {
	easyHttpClient *EasyHttp
	reqBody        api.RequestBody
}

func Builder() *EasyHttpBuilder {
	return &EasyHttpBuilder{easyHttpClient: easyhttpDefault}
}

func (b *EasyHttpBuilder) WithRequestBody(url string, headers map[string]string, payload interface{}) *EasyHttpBuilder {
	b.reqBody = bindRequestBody(url, headers, payload)
	return b
}

func bindRequestBody(url string, headers map[string]string, payload interface{}) api.RequestBody {
	return api.RequestBody{URL: url, Headers: headers, Payload: payload}
}

func (b *EasyHttpBuilder) Post() (api.ResponseBody, error) {
	return b.easyHttpClient.Post(b.reqBody)
}

func (b *EasyHttpBuilder) Put() (api.ResponseBody, error) {
	return b.easyHttpClient.Put(b.reqBody)
}

func (b *EasyHttpBuilder) Delete() (api.ResponseBody, error) {
	return b.easyHttpClient.Delete(b.reqBody)
}

func (b *EasyHttpBuilder) Patch() (api.ResponseBody, error) {
	return b.easyHttpClient.Patch(b.reqBody)
}

func (b *EasyHttpBuilder) Get() (api.ResponseBody, error) {
	return b.easyHttpClient.Get(b.reqBody)
}

func (b *EasyHttpBuilder) Route(fn func() (*http.Response, error)) (api.ResponseBody, error) {
	response, err := fn()
	if err != nil {
		return api.EmptyResponseBody, err
	}
	return api.ResourceToResponseBody(response), nil
}
