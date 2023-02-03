package management

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"net/http"
)

type GetContentTypesResponse struct {
	ContentTypes []schema.ContentType `json:"content_types"`
}

func (c *Client) GetAllContentTypes(includeCount int64, includeGlobalFieldSchema bool) ([]schema.ContentType, error) {
	endpoint := fmt.Sprintf("/v3/content_types?include_count=%d&include_global_field_schema=%b&include_branch=false", includeCount, includeGlobalFieldSchema)
	var r GetContentTypesResponse
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.ContentTypes, nil
}

type GetOneContentTypeResponse struct {
	ContentType *schema.ContentType `json:"content_type"`
}

func (c *Client) GetOneContentType(uid string) (*schema.ContentType, error) {
	endpoint := fmt.Sprintf("/v3/content_types/%s", uid)
	var r GetOneContentTypeResponse
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.ContentType, nil
}

type UpsertContentTypeRequestBody struct {
	ContentType *schema.ContentType `json:"content_type"`
}

type UpsertContentTypeResponse struct {
	Notice      string              `json:"notice"`
	ContentType *schema.ContentType `json:"content_type"`
}

func (c *Client) CreateContentType(g *schema.ContentType) (*schema.ContentType, error) {
	endpoint := fmt.Sprintf("/v3/content_types?include_branch=false")
	requestBody := UpsertContentTypeRequestBody{ContentType: g}
	var r UpsertContentTypeResponse
	resp, err := c.client.R().SetBody(requestBody).SetResult(&r).Post(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.ContentType, nil
}

func (c *Client) UpdateContentType(g *schema.ContentType) (*schema.ContentType, error) {
	if g == nil {
		return nil, fmt.Errorf("cannot update a nil ContentType")
	}
	endpoint := fmt.Sprintf("/v3/content_types/%s", g.UID)
	requestBody := UpsertContentTypeRequestBody{ContentType: g}
	var r UpsertContentTypeResponse
	resp, err := c.client.R().SetBody(requestBody).SetResult(&r).Put(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.ContentType, nil
}

func (c *Client) DeleteContentType(uid string) error {
	if uid == "" {
		return fmt.Errorf("cannot update a ContentType without a uid")
	}
	endpoint := fmt.Sprintf("/v3/content_types/%s", uid)
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
