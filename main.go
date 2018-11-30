package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type OEmbed struct {
	Status          int    `json:"-"`
	Type            string `json:"type,omitempty"`
	CacheAge        uint64 `json:"cache_age,omitempty"`
	URL             string `json:"url,omitempty"`
	ProviderURL     string `json:"provider_url,omitempty"`
	ProviderName    string `json:"provider_name,omitempty"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	Width           uint64 `json:"width,omitempty"`
	Height          uint64 `json:"height,omitempty"`
	ThumbnailURL    string `json:"thumbnail_url,omitempty"`
	ThumbnailWidth  uint64 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight uint64 `json:"thumbnail_height,omitempty"`
	AuthorName      string `json:"author_name,omitempty"`
	AuthorURL       string `json:"author_url,omitempty"`
	HTML            string `json:"html,omitempty"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	body := "Error"

	data := OEmbed{
		Type:            "rich",
		Title:           "oEmbed Banner",
		AuthorName:      "Duoc Nguyen",
		AuthorURL:       "https://12bit.vn",
		ProviderName:    "12bit",
		ProviderURL:     "https://12bit.vn",
		CacheAge:        3600,
		Width:           300,
		Height:          200,
		ThumbnailURL:    "https://ddc2txxlo9fx3.cloudfront.net/static/fwd_media_preview.png",
		ThumbnailHeight: 175,
		ThumbnailWidth:  280,
		HTML:            `<iframe src="https://oembed.fwdeveryone.com?thread-id=e8RFukWTS5Wo54fBNbZ2yQ" width="700" height="825" scrolling="yes" frameborder="0" allowfullscreen></iframe>`,
	}

	bData, err := json.Marshal(data)

	if err == nil {
		body = string(bData)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       body,
	}, nil
}

func main() {
	lambda.Start(handler)
}
