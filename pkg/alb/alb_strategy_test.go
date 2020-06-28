package alb

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequestBase64Encoded(t *testing.T) {
	s := NewALBStrategy()

	body := map[string]interface{}{
		"abc": 123,
		"123": "abc",
	}

	bodyJSON, _ := json.Marshal(&body)

	payload := map[string]interface{}{
		"httpMethod": "POST",
		"path":       "/this/is/a/path",
		"headers": map[string]string{
			"Content-Type": "application/json",
		},
		"isBase64Encoded": true,
		"body":            base64.StdEncoding.EncodeToString(bodyJSON),
	}

	payloadJSON, _ := json.Marshal(&payload)
	req, err := s.ParseRequest(context.Background(), payloadJSON)

	expectedReq, _ := http.NewRequest("POST", "/this/is/a/path", strings.NewReader(string(bodyJSON)))
	expectedReq.Header.Add("Content-Type", "application/json")

	assert.Nil(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, expectedReq.Method, req.Method)
	assert.Equal(t, expectedReq.Header, req.Header)
	assert.Equal(t, expectedReq.URL.Path, req.URL.Path)

	expectedPayload, _ := ioutil.ReadAll(expectedReq.Body)
	actualPayload, _ := ioutil.ReadAll(req.Body)

	assert.Equal(t, string(expectedPayload), string(actualPayload))

}
