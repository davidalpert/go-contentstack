// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2023-05-16 23:28:44.659642 -0500 CDT m=+0.034579584
package version

import (
	"fmt"
	"os/user"
	"runtime"
	"sort"
	"strings"
)

// Detail provides an easy global way to
var Detail = NewVersionDetail()

// NewVersionDetail builds a new version DetailStruct
func NewVersionDetail() DetailStruct {
	s := DetailStruct{
		AppName:              "cs",
		BuildDate:            "2023-05-16 23:28:44.659642 -0500 CDT m=+0.034579584",
		CoreVersion:          "0.2.1",
		GitBranch:            "main",
		GitCommit:            "5b5933f",
		GitCommitSummary:     "release notes for v0.2.1",
		GitDirty:             true,
		GitDirtyHasModified:  true,
		GitDirtyHasStaged:    true,
		GitDirtyHasUntracked: false,
		Version:              "0.2.1+5b5933f.dirty",
	}
	s.UserAgentString = s.ToUserAgentString()
	if s.GitDirty {
		s.GitWorkingState = "dirty"
	}
	return s
}
// DetailStruct provides an easy way to grab all the govvv version details together
type DetailStruct struct {
	AppName              string `json:"app_name"`
	BuildDate            string `json:"build_date"`
	CoreVersion          string `json:"core_version"`
	GitBranch            string `json:"branch"`
	GitCommit            string `json:"commit"`
	GitCommitSummary     string `json:"commit_summary"`
	GitDirty             bool `json:"dirty"`
	GitDirtyHasModified  bool `json:"dirty_modified"`
	GitDirtyHasStaged    bool `json:"dirty_staged"`
	GitDirtyHasUntracked bool `json:"dirty_untracked"`
	GitWorkingState      string `json:"working_state"`
	GitSummary           string `json:"summary"`
	UserAgentString      string `json:"user_agent"`
	Version              string `json:"version"`
}

// String implements Stringer
func (d *DetailStruct) String() string {
	if d == nil {
		return "n/a"
	}
	return fmt.Sprintf("%s %s", d.AppName, d.Version)
}

// ToUserAgentString formats a DetailStruct as a User-Agent string
func (s DetailStruct) ToUserAgentString() string {
	productName := s.AppName
	productVersion := s.Version

	productDetails := map[string]string{ }

	user, err := user.Current()
	if err == nil {
		username := user.Username
		if username == "" {
			username = "unknown"
		}
	}

	detailParts := []string{}
	for k, v := range productDetails {
		detailParts = append(detailParts, fmt.Sprintf("%s: %s", k, v))
	}
	sort.Slice(detailParts, func(i, j int) bool {
		return detailParts[i] < detailParts[j]
	})
	productDetail := strings.Join(detailParts, ", ")

	return fmt.Sprintf("%s/%s (%s) %s (%s)", productName, productVersion, productDetail, runtime.GOOS, runtime.GOARCH)
}
