package lambdarouter

import (
	"context"
	"net/http"
)

// LambdaStrategy represents a way of handling Lambda requests and responses.
// Current Implementations: ALBStrategy or APIGatewayStrategy
type LambdaStrategy interface {
	ParseRequest(ctx context.Context, payload []byte) (*http.Request, error)
	NewResponseWriter() AWSResponseWriter
}
