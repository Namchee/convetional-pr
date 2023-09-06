package entity

import (
	"net/url"
	"os"
	"testing"

	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	baseUrl, _ := url.Parse("https://api.github.com/")

	tests := []struct {
		name    string
		mocks   map[string]string
		want    *Configuration
		wantErr error
	}{
		{
			name: "should read config correctly",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":    "foo_bar",
				"INPUT_DRAFT":           "false",
				"INPUT_CLOSE":           "true",
				"INPUT_ISSUE":           "true",
				"INPUT_BODY":            "true",
				"INPUT_BOT":             "false",
				"INPUT_MAXIMUM_CHANGES": "11",
				"INPUT_SIGNED":          "true",
				"INPUT_IGNORED_USERS":   "Namchee, snyk-bot",
				"INPUT_EDIT":            "true",
				"INPUT_LABEL":           "cpr:invalid",
				"INPUT_VERBOSE":         "true",
				"INPUT_MESSAGE":         "foo bar",
				"GITHUB_API_URL":        "https://api.github.com/",
			},
			want: &Configuration{
				Token:        "foo_bar",
				Draft:        false,
				Issue:        true,
				Close:        true,
				Body:         true,
				Bot:          false,
				FileChanges:  11,
				Signed:       true,
				Label:        "cpr:invalid",
				IgnoredUsers: []string{"Namchee", "snyk-bot"},
				Edit:         true,
				Verbose:      true,
				Message:      "foo bar",
				BaseURL:      baseUrl,
			},
			wantErr: nil,
		},
		{
			name:    "should throw an error when token is empty",
			mocks:   map[string]string{},
			want:    nil,
			wantErr: constants.ErrMissingToken,
		},
		{
			name: "should throw an error when fileChanges is negative",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":    "foo",
				"INPUT_MAXIMUM_CHANGES": "-1",
			},
			want:    nil,
			wantErr: constants.ErrNegativeFileChange,
		},
		{
			name: "should throw an error when title pattern is invalid",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":  "a",
				"INPUT_TITLE_PATTERN": "[",
			},
			want:    nil,
			wantErr: constants.ErrInvalidTitlePattern,
		},
		{
			name: "should throw an error when commit pattern is invalid",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":   "b",
				"INPUT_TITLE_PATTERN":  "a",
				"INPUT_COMMIT_PATTERN": "[",
			},
			want:    nil,
			wantErr: constants.ErrInvalidCommitPattern,
		},
		{
			name: "should throw an error when branch pattern is invalid",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN":   "token",
				"INPUT_TITLE_PATTERN":  "a",
				"INPUT_COMMIT_PATTERN": "a",
				"INPUT_BRANCH_PATTERN": "[",
			},
			want:    nil,
			wantErr: constants.ErrInvalidBranchPattern,
		},
		{
			name: "should throw an error when base URL is invalid",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN": "token",
				"GITHUB_API_URL":     " https://api.github.com",
			},
			want:    nil,
			wantErr: constants.ErrInvalidBaseURL,
		},
		{
			name: "should not append trailing slash on base URL when URL has trailing slash",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN": "token",
				"GITHUB_API_URL":     "https://api.github.com/",
			},
			want: &Configuration{
				Token:        "token",
				BaseURL:      baseUrl,
				IgnoredUsers: []string{},
			},
			wantErr: nil,
		},
		{
			name: "should append trailing slash on base URL",
			mocks: map[string]string{
				"INPUT_ACCESS_TOKEN": "token",
				"GITHUB_API_URL":     "https://api.github.com",
			},
			want: &Configuration{
				Token:        "token",
				BaseURL:      baseUrl,
				IgnoredUsers: []string{},
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for key, val := range tc.mocks {
				os.Setenv(key, val)
				defer os.Unsetenv(key)
			}

			got, err := ReadConfig()

			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
