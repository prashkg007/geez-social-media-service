package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var adapter *gorillamux.GorillaMuxAdapterV2

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/post/create", createPostHandler).Methods("POST")
	adapter = gorillamux.NewV2(r)
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	res, err := adapter.ProxyWithContext(ctx, request)
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, err
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("damn son, that is something"))
}

func main() {
	lambda.Start(handler)
}
