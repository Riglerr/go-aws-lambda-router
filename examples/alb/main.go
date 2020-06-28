package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	lambdarouter "github.com/riglerr/go-aws-lambda-router/pkg"
	"github.com/riglerr/go-aws-lambda-router/pkg/alb"
)

// LoginHandler handles the HTTP event: GET /login
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received /login request, %+v", req)
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

// CallbackHandler handles the HTTP event: GET /callback
func CallbackHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received /callback request, %+v", req)
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("CALLBACK OK"))
}

func main() {

	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/callback", CallbackHandler)
	lambda.StartHandler(lambdarouter.NewLambdaRouter(alb.NewALBStrategy()))
}
