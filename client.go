package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type Site struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type GUID struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Title struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Content struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Excerpt struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}

type Post struct {
	ID            int     `json:"id,omitempty"`
	Date          string  `json:"date,omitempty"`
	DateGMT       string  `json:"date_gmt,omitempty"`
	GUID          GUID    `json:"guid,omitempty"`
	Link          string  `json:"link,omitempty"`
	Modified      string  `json:"modified,omitempty"`
	ModifiedGMT   string  `json:"modifiedGMT,omitempty"`
	Password      string  `json:"password,omitempty"`
	Slug          string  `json:"slug,omitempty"`
	Status        string  `json:"status,omitempty"`
	Type          string  `json:"type,omitempty"`
	Title         Title   `json:"title,omitempty"`
	Content       Content `json:"content,omitempty"`
	Author        int     `json:"author,omitempty"`
	Excerpt       Excerpt `json:"excerpt,omitempty"`
	FeaturedImage int     `json:"featured_image,omitempty"`
	CommentStatus string  `json:"comment_status,omitempty"`
	PingStatus    string  `json:"ping_status,omitempty"`
	Format        string  `json:"format,omitempty"`
	Sticky        bool    `json:"sticky,omitempty"`
}

type GeneralError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"` // Unsure if this is consistent
}

type ListOption struct {
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Order   string `url:"order,omitempty"`
}

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseUrl string) *Client {
	wpClient := &Client{
		baseURL: baseUrl,
	}
	wpClient.client = &http.Client{}

	return wpClient
}

func (c *Client) GetPosts(option *ListOption) ([]*Post, error) {
	response, _, err := c.Get("posts", option)
	if err != nil {
		return nil, err
	}

	var posts []*Post
	if err := json.Unmarshal(response, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (c *Client) GetSiteInfo() (*Site, error) {
	body, _, err := c.Get("", nil)
	if err != nil {
		return nil, err
	}
	var site Site
	if err := json.Unmarshal(body, &site); err != nil {
		return nil, err
	}
	return &site, nil
}

func (c *Client) Get(endpoint string, option interface{}) ([]byte, *http.Response, error) {
	queryString, _ := query.Values(option)
	var url string
	if endpoint != "" {
		url = fmt.Sprintf("%s/wp-json/wp/v2/%s?%s", c.baseURL, endpoint, queryString.Encode())
	} else {
		url = fmt.Sprintf("%s/wp-json", c.baseURL)
	}
	response, err := c.client.Get(url)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode != 200 {
		var data GeneralError
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, nil, err
		}
		return nil, nil, errors.New(data.Message)
	}

	return body, response, nil
}
