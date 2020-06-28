package lambdarouter

import "net/http"

// AWSResponseWriter interface
type AWSResponseWriter interface {
	http.ResponseWriter
	ToAwsResponse() interface{}
}
