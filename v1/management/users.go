package management

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"net/http"
)

func (c *Client) GetUser() (*schema.User, error) {
	endpoint := "/v3/user"
	var r schema.UserResponse
	resp, err := c.client.R().SetBody(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &r.User, nil
}
