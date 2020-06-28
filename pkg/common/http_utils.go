package common

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// FlattenHeaders flattens the multi-value http.Header into a map[string][]
// multi-header values are separated by ', '
func FlattenHeaders(h *http.Header) map[string]string {
	r := map[string]string{}
	for key, values := range *h {
		r[key] = strings.Join(values, ", ")
	}
	return r
}

// ShouldDecodePayload returns true when the ALB/APIGateway event has a payload that needs to be decoded before use.
func ShouldDecodePayload(event *events.ALBTargetGroupRequest) bool {
	return (event.HTTPMethod == "POST" || event.HTTPMethod == "PUT" ||
		event.HTTPMethod == "PATCH") && event.IsBase64Encoded
}
