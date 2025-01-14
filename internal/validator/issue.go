package validator

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

var (
	keywordPattern = regexp.MustCompile(`(?mi)\b(close|closes|closed|fix|fixes|fixed|resolve|resolves|resolved)\s+(.+/.+)?#(\d+)\b`)
)

type issueValidator struct {
	Name string

	client internal.GithubClient
	config *entity.Configuration
}

// NewIssueValidator creates a new validator that validates issue resolution
func NewIssueValidator(
	client internal.GithubClient,
	config *entity.Configuration,
) internal.Validator {
	return &issueValidator{
		Name:   constants.IssueValidatorName,
		client: client,
		config: config,
	}
}

func (v *issueValidator) IsValid(
	ctx context.Context,
	pullRequest *entity.PullRequest,
) *entity.ValidationResult {
	if !v.config.Issue {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: false,
			Result: nil,
		}
	}

	references, err := v.client.GetIssueReferences(
		ctx,
		&pullRequest.Repository,
		pullRequest.Number,
	)
	if err != nil {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: true,
			Result: nil,
		}
	}

	for _, reference := range references {
		repo := reference.Owner + "/" + reference.Name
		meta := pullRequest.Repository.Owner + "/" + pullRequest.Repository.Name

		if repo == meta {
			return &entity.ValidationResult{
				Name:   constants.IssueValidatorName,
				Active: true,
				Result: nil,
			}
		}
	}

	if v.hasIssueMagicString(ctx, pullRequest) {
		return &entity.ValidationResult{
			Name:   constants.IssueValidatorName,
			Active: true,
			Result: nil,
		}
	}

	return &entity.ValidationResult{
		Name:   constants.IssueValidatorName,
		Active: true,
		Result: constants.ErrNoIssue,
	}
}

func (v *issueValidator) hasIssueMagicString(
	_ context.Context,
	pullRequest *entity.PullRequest,
) bool {
	keywords := keywordPattern.FindAllStringSubmatch(pullRequest.Body, -1)

	for _, keyword := range keywords {
		org := &pullRequest.Repository
		num, _ := strconv.Atoi(keyword[3]) 

		if len(keyword[2]) > 0 {
			tokens := strings.Split(keyword[2], "/")

			org = &entity.Meta{
				Name: tokens[1],
				Owner: tokens[0],
			}
		}

		issue, _ := v.client.GetIssue(
			context.Background(),
			org,
			num,
		)

		if issue != nil {
			return true
		}
	}

	return false
}
