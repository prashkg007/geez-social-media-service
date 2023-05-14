package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/post/create":
		return createPost(ctx, request)
	default:
		return events.APIGatewayProxyResponse{Body: "Not Found", StatusCode: 404}, nil
	}
}

func createPost(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		Body:       "Hoka is such a bitch!",
		StatusCode: 200,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
