package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/Namchee/ethos/internal"
	"github.com/Namchee/ethos/internal/constants"
	"github.com/Namchee/ethos/internal/entity"
	"github.com/google/go-github/v32/github"
)

type commitValidator struct {
	client internal.GithubClient
	config *entity.Config
	meta   *entity.Meta
	Name   string
}

func NewCommitValidator(
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

func (v *commitValidator) IsValid(pullRequest *github.PullRequest) *entity.ValidationResult {
	if !v.config.Commits {
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

	pattern := regexp.MustCompile(v.config.Pattern)

	for _, commit := range commits {
		message := commit.Commit.GetMessage()

		if !pattern.Match([]byte(message)) {
			return &entity.ValidationResult{
				Name: v.Name,
				Result: fmt.Errorf(
					"commit %s does not have valid commit message", commit.Commit.GetSHA(),
				),
			}
		}
	}

	return &entity.ValidationResult{
		Name:   v.Name,
		Result: nil,
	}
}