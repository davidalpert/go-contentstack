package management

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"net/http"
)

func (c *Client) GetAllPublishingEnvironments() ([]schema.Environment, error) {
	endpoint := fmt.Sprintf("/v3/environments")
	var r schema.EnvironmentListWrapper
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.Environments, nil
}

func (c *Client) GetOnePublishingEnvironment(name string) (*schema.Environment, error) {
	endpoint := fmt.Sprintf("/v3/environments/%s", name)
	var r schema.SingleEnvironmentWrapper
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &r.Environment, nil
}

func (c *Client) CreatePublishingEnvironment(g *schema.SingleEnvironmentWrapper) (*schema.SingleEnvironmentWrapper, error) {
	endpoint := fmt.Sprintf("/v3/environments")
	var r schema.SingleEnvironmentWrapper
	resp, err := c.client.R().SetBody(g).SetResult(&r).Post(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &r, nil
}

func (c *Client) UpdatePublishingEnvironment(e *schema.Environment) (*schema.Environment, error) {
	if e == nil {
		return nil, fmt.Errorf("cannot update a nil Environment")
	}
	endpoint := fmt.Sprintf("/v3/environments/%s", e.Name)
	var r schema.UpsertEnvironmentResponse
	resp, err := c.client.R().SetBody(schema.SingleEnvironmentWrapper{Environment: *e}).SetResult(&r).Put(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &r.Environment, nil
}

func (c *Client) DeletePublishingEnvironment(name string) error {
	if name == "" {
		return fmt.Errorf("cannot update a publishing environment without a name")
	}
	endpoint := fmt.Sprintf("/v3/environments/%s", name)
	var r UpsertContentTypeResponse
	resp, err := c.client.R().SetResult(&r).Delete(endpoint)
	if err != nil {
		return err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return nil
}
