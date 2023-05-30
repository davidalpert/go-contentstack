package management

import (
	"fmt"
	"github.com/davidalpert/go-contentstack/v1/schema"
	"net/http"
)

type LocaleListWrapper struct {
	Locales []schema.Locale `json:"locales"`
}

func (c *Client) GetAllLocales() ([]schema.Locale, error) {
	endpoint := fmt.Sprintf("/v3/locales")
	var r LocaleListWrapper
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return r.Locales, nil
}

type GetOneLocaleResult struct {
	Locale schema.Locale `json:"locale"`
}

func (c *Client) GetOneLocale(name string) (*schema.Locale, error) {
	endpoint := fmt.Sprintf("/v3/locales/%s", name)
	var r GetOneLocaleResult
	resp, err := c.client.R().SetResult(&r).Get(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &r.Locale, nil
}

type LocaleCreateRequest struct {
	Locale struct {
		Name           string `json:"name"`
		Code           string `json:"code"`
		FallbackLocale string `json:"fallback_locale"`
	} `json:"locale"`
}

type UpsertLocaleResult struct {
	Notice string        `json:"notice"`
	Locale schema.Locale `json:"locale"`
}

func (c *Client) CreateLocale(code, name, fallbackCode string) (*schema.Locale, error) {
	endpoint := fmt.Sprintf("/v3/locales")
	var input = LocaleCreateRequest{
		Locale: struct {
			Name           string `json:"name"`
			Code           string `json:"code"`
			FallbackLocale string `json:"fallback_locale"`
		}{
			Name:           name,
			Code:           code,
			FallbackLocale: fallbackCode,
		},
	}
	var output UpsertLocaleResult
	resp, err := c.client.R().SetBody(input).SetResult(&output).Post(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &output.Locale, nil
}

type LocaleUpdateRequest struct {
	Locale struct {
		Name           string `json:"name"`
		FallbackLocale string `json:"fallback_locale"`
	} `json:"locale"`
}

func (c *Client) UpdateLocale(e *schema.Locale) (*schema.Locale, error) {
	if e == nil {
		return nil, fmt.Errorf("cannot update a nil Locale")
	}
	endpoint := fmt.Sprintf("/v3/locales/%s", e.Code)
	var input = LocaleUpdateRequest{Locale: struct {
		Name           string `json:"name"`
		FallbackLocale string `json:"fallback_locale"`
	}{Name: e.Name, FallbackLocale: e.FallbackLocaleCode}}
	var output UpsertLocaleResult
	resp, err := c.client.R().SetBody(input).SetResult(&output).Put(endpoint)
	if err != nil {
		return nil, err
	}
	body := resp.Body()
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("calling %#v  %s: %s", endpoint, resp.Status(), string(body))
	}

	return &output.Locale, nil
}

func (c *Client) DeleteLocale(code string) error {
	if code == "" {
		return fmt.Errorf("cannot delete a Locale without a code")
	}
	endpoint := fmt.Sprintf("/v3/locales/%s", code)
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
