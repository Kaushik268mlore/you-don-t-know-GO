package main

import (
	"fmt"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f ", request.ID),
		Ok:      true,
	}, nil
}
func main() {
	lambda.start(handler)
}
