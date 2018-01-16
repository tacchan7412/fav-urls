package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	api = newTwitterAPI()
}

// Handler is AWS lambda handler
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	v := url.Values{}
	v.Set("count", "20")
	tweets, err := api.GetFavorites(v)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	urls := extractURLsFromTweets(tweets)

	if len(urls) == 0 {
		return events.APIGatewayProxyResponse{}, errors.New("no new urls")
	}

	bytes, err := json.Marshal(urls)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Body:            string(bytes[:]),
		IsBase64Encoded: true,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
