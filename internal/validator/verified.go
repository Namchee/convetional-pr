package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
	"github.com/google/go-github/v32/github"
)

type verifiedValidator struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

// NewVerifiedValidator creates a new validator that will validate all commit messages in a pull request
func NewVerifiedValidator(
	client internal.GithubClient,
	config *entity.Config,
	meta *entity.Meta,
) internal.Validator {
	return &commitValidator{
		Name:   constants.CommitValidatorName,
		client: client,
		config: config,
		meta:   meta,
	}
}

func (v *verifiedValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if v.config.CommitPattern == "" {
		return &entity.ValidationResult{
			Name:   v.Name,
			Result: nil,
		}
	}

	ctx := context.Background()

	commits, _ := v.client.GetCommits(
		ctx,
		v.meta.Owner,
		v.meta.Name,
		pullRequest.GetNumber(),
	)

	pattern := regexp.MustCompile(v.config.CommitPattern)

	for _, commit := range commits {
		message := commit.Commit.GetMessage()

		if !pattern.Match([]byte(message)) {
			return &entity.ValidationResult{
				Name: v.Name,
				Result: fmt.Errorf(
					"commit %s does not have valid commit message", commit.GetSHA(),
				),
			}
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: nil,
	}
}