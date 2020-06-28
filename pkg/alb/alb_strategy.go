package alb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambdarouter "github.com/riglerr/go-aws-lambda-router/pkg"
	"github.com/riglerr/go-aws-lambda-router/pkg/common"
)

// Strategy handling incoming ALBTargetGroupRequest(s), and writing ALBTargetGroupResponse(s)
// Implements the LambdaStrategy interface
type Strategy struct {
	AwsEvent *events.ALBTargetGroupRequest
}

// NewALBStrategy creates a new ALB LambdaStrategy
func NewALBStrategy() *Strategy {
	return &Strategy{
		AwsEvent: &events.ALBTargetGroupRequest{},
	}
}

// ParseRequest parses the raw lambda payload into an http.Request
func (s *Strategy) ParseRequest(ctx context.Context, payload []byte) (*http.Request, error) {
	s.AwsEvent = &events.ALBTargetGroupRequest{}
	if err := json.Unmarshal(payload, s.AwsEvent); err != nil {
		log.Printf("%+v", fmt.Errorf("Unable to parse ALB request, %w", err))
	}
	// Todo context
	var body io.Reader = strings.NewReader(s.AwsEvent.Body)
	if common.ShouldDecodePayload(s.AwsEvent) {
		body = base64.NewDecoder(base64.StdEncoding, body)
	}

	req, err := http.NewRequestWithContext(ctx, s.AwsEvent.HTTPMethod, s.AwsEvent.Path, body)
	if err != nil {
		log.Printf("%+v", fmt.Errorf("Unable to create HTTP request, %w", err))
		return nil, err
	}

	// Todo multi-value headers
	for k, v := range s.AwsEvent.Headers {
		req.Header.Add(k, v)
	}

	return req, nil
}

// NewResponseWriter does something
func (s *Strategy) NewResponseWriter() lambdarouter.AWSResponseWriter {
	return &ResponseWriter{
		headers: map[string][]string{},
	}
}
