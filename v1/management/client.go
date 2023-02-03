package management

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	client *resty.Client
	Configuration
}

func NewClient(cfg *Configuration) (*Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("configuration is required")
	}

	// Create a Resty Client
	client := resty.New()

	// Unique settings at Client level
	//--------------------------------
	// Enable debug mode
	client.SetDebug(cfg.Debug)

	// Assign Client TLSClientConfig
	// One can set custom root-certificate. Refer: http://golang.org/pkg/crypto/tls/#example_Dial
	//client.SetTLSClientConfig(&tls.Config{RootCAs: roots})

	// or One can disable security check (https)
	//client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// Set client timeout as per your need
	client.SetTimeout(1 * time.Minute)

	// You can override all below settings and options at request level if you want to
	//--------------------------------------------------------------------------------
	// Host URL for all request. So you can use relative URL in the request
	client.SetBaseURL(cfg.Host)

	// Headers for all request
	commonHeaders := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"api_key":       cfg.Key,
		"Authorization": cfg.Token,
	}
	if cfg.UserAgent != "" {
		commonHeaders["User-Agent"] = cfg.UserAgent
	}
	client.SetHeaders(commonHeaders)
	//client.SetAuthToken(cfg.Token)

	return &Client{
		client:        client,
		Configuration: *cfg,
	}, nil
}
