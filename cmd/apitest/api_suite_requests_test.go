package apitest

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/steinfletcher/apitest"
)

type jsonCallOptions struct {
	responseData    any
	requestData     any
	queryParams     map[string]string
	queryCollection map[string][]string
	responseStatus  int
}

type JSONCallOption func(options *jsonCallOptions)

func WithResponseData(data any) JSONCallOption {
	if reflect.TypeOf(data).Kind() != reflect.Pointer {
		panic("JSONCall WithResponseData must be a pointer!")
	}
	return func(options *jsonCallOptions) {
		options.responseData = data
	}
}

func WithRequestData(data any) JSONCallOption {
	return func(options *jsonCallOptions) {
		options.requestData = data
	}
}

func WithQueryParams(params map[string]string) JSONCallOption {
	return func(options *jsonCallOptions) {
		options.queryParams = params
	}
}

func WithQueryCollection(collection map[string][]string) JSONCallOption {
	return func(options *jsonCallOptions) {
		options.queryCollection = collection
	}
}

func WithResponseStatus(httpStatus int) JSONCallOption {
	return func(options *jsonCallOptions) {
		options.responseStatus = httpStatus
	}
}

func (suite *ApiTestSuite) NewCall(t *testing.T, name, method, endpoint string, reqData func(request *apitest.Request) *apitest.Request) *apitest.Response {
	path, err := url.JoinPath(suite.baseURL, endpoint)
	if err != nil {
		panic(err)
	}
	request := apitest.New(name).
		EnableNetworking().
		Method(method)
	if reqData != nil {
		return reqData(request).URL(path).Expect(t)
	}
	return request.URL(path).Expect(t)
}

func (suite *ApiTestSuite) NewTextCall(t *testing.T, name, method, endpoint string, options ...JSONCallOption) {
	opts := jsonCallOptions{
		responseStatus: http.StatusOK,
	}
	for _, option := range options {
		option(&opts)
	}

	end := suite.NewCall(t, name, method, endpoint, func(request *apitest.Request) *apitest.Request {
		req := request.Header("Content-Type", "text/plain")
		if opts.requestData != nil {
			stringData, ok := opts.requestData.(string)
			if !ok {
				t.Fatal("payload is not a string")
			}
			req = req.Body(stringData)
		}
		if opts.queryParams != nil {
			req.QueryParams(opts.queryParams)
		}
		if opts.queryCollection != nil {
			req.QueryCollection(opts.queryCollection)
		}
		return req
	}).Status(opts.responseStatus).End()
	if opts.responseData != nil {
		end.JSON(opts.responseData)
	}
}
