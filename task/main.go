package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// OutputEvent is form of task 1 output.
type OutputEvent struct {
	Text     string   `json:"t"`
	IntArray []int    `json:"ia"`
	StrArray []string `json:"sa"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (OutputEvent, error) {
	return OutputEvent{
		Text:     "nanigashi",
		IntArray: []int{1, 2, 4, 8, 16, 32, 64},
		StrArray: []string{"A", "B", "C", "D", "E"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
