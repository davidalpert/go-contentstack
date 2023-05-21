package schema

import "time"

type EnvironmentUpsertRequest struct {
	Environment Environment `json:"environment"`
}

type LocaleURLPair struct {
	Locale string `json:"locale"`
	Url    string `json:"url"`
}

type Environment struct {
	Name          string          `json:"name"`
	Urls          []LocaleURLPair `json:"urls"`
	UID           string          `json:"uid"`
	CreatedBy     string          `json:"created_by"`
	UpdatedBy     string          `json:"updated_by"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	ACL           interface{}     `json:"ACL"`
	Tags          interface{}     `json:"tags"`
	Version       int             `json:"_version"`
	DeployContent bool            `json:"deploy_content"`
}

type SingleEnvironmentWrapper struct {
	Environment Environment `json:"environment"`
}

type EnvironmentListWrapper struct {
	Environments []Environment `json:"environments"`
}

type UpsertEnvironmentResponse struct {
	SingleEnvironmentWrapper
	Notice string `json:"notice"`
}
