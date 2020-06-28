package lambdarouter

import (
	"context"
	"encoding/json"
	"net/http"
)

// LambdaRouter is a struct representing the lambda configuration.
// Implements lambda.Handler
type LambdaRouter struct {
	Strategy LambdaStrategy
}

// Invoke is the function called by lambda.StartHandler at lambda runtime.
func (r *LambdaRouter) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	req, _ := r.Strategy.ParseRequest(ctx, payload)

	w := r.Strategy.NewResponseWriter()
	mux := http.DefaultServeMux
	mux.ServeHTTP(w, req)

	responseData, _ := json.Marshal(w.ToAwsResponse())
	return responseData, nil
}

// NewLambdaRouter creates a new Lambda Router
func NewLambdaRouter(s LambdaStrategy) *LambdaRouter {
	return &LambdaRouter{
		Strategy: s,
	}
}
