package alb

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestCreatesCorrectResponse(t *testing.T) {
	w := NewResponseWriter()

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Customheader", "customheadervalue")
	w.Write([]byte("Test"))

	expected := &events.ALBTargetGroupResponse{
		StatusCode:        200,
		StatusDescription: http.StatusText(200),
		Headers: map[string]string{
			"Content-Type": "text/plain",
			"Customheader": "customheadervalue",
		},
		MultiValueHeaders: w.headers,
		Body:              "Test",
	}

	assert.Equal(t, expected, w.ToAwsResponse())
}
