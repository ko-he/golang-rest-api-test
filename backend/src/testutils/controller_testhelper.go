package testutils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var (
	onceDoc    sync.Once
	openapiDoc *openapi3.T
)

func AssertRequestWithOpenAPISpec(t *testing.T, request *http.Request, recorder *httptest.ResponseRecorder) {
	t.Helper()

	// loader := openapi3.Loader{}

	onceDoc.Do(func() {})

	router := lo.Must(legacy.NewRouter(openapiDoc))

	route, params := lo.Must2(router.FindRoute(request))

	requestInput := &openapi3filter.RequestValidationInput{
		Request:    request,
		PathParams: params,
		Route:      route,
	}

	assert.NoError(t, openapi3filter.ValidateRequest(context.Background(), requestInput))

	responseInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestInput,
		Status:                 recorder.Code,
		Header:                 recorder.Header(),
	}

	responseInput.SetBodyBytes(lo.Must(io.ReadAll(recorder.Body)))

	assert.NoError(t, openapi3filter.ValidateResponse(context.Background(), responseInput))
}

type RequestBuilder struct {
	method  string
	target  string
	query   map[string][]string
	body    any
	headers map[string]string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		method:  "GET",
		target:  "/",
		query:   map[string][]string{},
		body:    map[string]any{},
		headers: map[string]string{},
	}
}

func (b *RequestBuilder) WithMethod(method string) *RequestBuilder {
	b.method = method

	return b
}

func (b *RequestBuilder) WithQuery(query map[string][]string) *RequestBuilder {
	b.query = query

	return b
}

func (b *RequestBuilder) WithTarget(target string) *RequestBuilder {
	b.target = target

	return b
}

func (b *RequestBuilder) WithBody(body any) *RequestBuilder {
	b.body = body

	return b
}

func (b *RequestBuilder) WithHeader(key string, value string) *RequestBuilder {
	b.headers[key] = value

	return b
}

func (b *RequestBuilder) Build() (*http.Request, error) {
	query := url.Values{}
	for key, values := range b.query {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	body, err := json.Marshal(b.body)
	if err != nil {
		return nil, err
	}

	request := httptest.NewRequest(
		b.method,
		fmt.Sprintf("%s?%s", b.target, query.Encode()),
		strings.NewReader(string(body)),
	)

	for key, value := range b.headers {
		request.Header.Set(key, value)
	}

	return request, nil
}

func (b *RequestBuilder) MustBuild() *http.Request {
	return lo.Must(b.Build())
}
