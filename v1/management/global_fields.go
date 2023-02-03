package management

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"net/http"
)

type GetGlobalFieldsResponse struct {
	GlobalFields []schema.GlobalField `json:"global_fields"`
}

func (c *Client) GetAllGlobalFields() ([]schema.GlobalField, error) {
	endpoint := fmt.Sprintf("/v3/global_fields")
	var r GetGlobalFieldsResponse
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.GlobalFields, nil
}

type GetOneGlobalFieldResponse struct {
	GlobalField *schema.GlobalField `json:"global_field"`
}

func (c *Client) GetOneGlobalField(uid string) (*schema.GlobalField, error) {
	endpoint := fmt.Sprintf("/v3/global_fields/%s", uid)
	var r GetOneGlobalFieldResponse
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.GlobalField, nil
}

type UpsertGlobalFieldRequestBody struct {
	GlobalField *schema.GlobalField `json:"global_field"`
}

type UpsertGlobalFieldResponse struct {
	Notice      string              `json:"notice"`
	GlobalField *schema.GlobalField `json:"global_field"`
}

func (c *Client) CreateGlobalField(g *schema.GlobalField) (*schema.GlobalField, error) {
	endpoint := fmt.Sprintf("/v3/global_fields?include_branch=false")
	requestBody := UpsertGlobalFieldRequestBody{GlobalField: g}
	var r UpsertGlobalFieldResponse
	resp, err := c.client.R().SetBody(requestBody).SetResult(&r).Post(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.GlobalField, nil
}

func (c *Client) UpdateGlobalField(g *schema.GlobalField) (*schema.GlobalField, error) {
	if g == nil {
		return nil, fmt.Errorf("cannot update a nil GlobalField")
	}
	endpoint := fmt.Sprintf("/v3/global_fields/%s", g.UID)
	requestBody := UpsertGlobalFieldRequestBody{GlobalField: g}
	var r UpsertGlobalFieldResponse
	resp, err := c.client.R().SetBody(requestBody).SetResult(&r).Put(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.GlobalField, nil
}

func (c *Client) DeleteGlobalField(uid string) error {
	if uid == "" {
		return fmt.Errorf("cannot update a GlobalField without a uid")
	}
	endpoint := fmt.Sprintf("/v3/global_fields/%s", uid)
	var r UpsertGlobalFieldResponse
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
