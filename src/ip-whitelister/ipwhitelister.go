package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	//"github.com/aws/aws-sdk-go"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
