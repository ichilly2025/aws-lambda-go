package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"strconv"
)

// RequestPayload defines request payload
type RequestPayload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler handles lambda requests
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse request body
	var payload RequestPayload
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid JSON payload",
		}, nil
	}

	// Read path parameter
	name, exists := request.PathParameters["name"]
	if exists {
		payload.Name = name
	}

	// Read query parameter
	age, exists := request.QueryStringParameters["age"]
	if exists {
		iAge, err := strconv.Atoi(age)
		if err == nil {
			payload.Age = iAge
		}
	}

	// Create response body
	responseBody := map[string]interface{}{
		"message": fmt.Sprintf("Hello, %s! You are %d years old.", payload.Name, payload.Age),
	}
	responseJSON, _ := json.Marshal(responseBody)

	// Return proxy response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseJSON),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
