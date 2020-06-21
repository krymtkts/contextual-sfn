package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// inputEvent is form of task 1 output.
type inputEvent struct {
	Key    int      `json:"k"`
	Text   string   `json:"t"`
	ArrayB []string `json:"a"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event inputEvent) {
	log.Printf("input event: %v\n", event)
}

func main() {
	lambda.Start(Handler)
}
