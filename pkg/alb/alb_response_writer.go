package alb

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/riglerr/go-aws-lambda-router/pkg/common"
)

// ResponseWriter converts the http response to an ALBTargetGroupResponse
type ResponseWriter struct {
	body       []byte
	statusCode int
	headers    http.Header
}

// Header returns the headers
func (w *ResponseWriter) Header() http.Header {
	return w.headers
}

// Writes bytes to the body
func (w *ResponseWriter) Write(payload []byte) (int, error) {
	w.body = append(w.body, payload...)
	return len(payload), nil
}

// WriteHeader writes status code header
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

// ToAwsResponse wrties ALB response
func (w *ResponseWriter) ToAwsResponse() interface{} {
	return &events.ALBTargetGroupResponse{
		StatusCode:        w.statusCode,
		StatusDescription: http.StatusText(w.statusCode),
		Headers:           common.FlattenHeaders(&w.headers),
		MultiValueHeaders: w.headers,
		Body:              string(w.body),
	}
}

// NewResponseWriter creates a new ResponseWriter
func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		headers: map[string][]string{},
	}
}
