package mocks

import (
	"context"
	"errors"

	"github.com/Namchee/conventional-pr/internal"
	"github.com/Namchee/conventional-pr/internal/constants"
	"github.com/Namchee/conventional-pr/internal/entity"
)

// GitHub's client mock. Used in testing
type githubClientMock struct{}

func (m *githubClientMock) GetPullRequest(
	_ context.Context,
	_ *entity.Meta,
	prNumber int,
) (*entity.PullRequest, error) {
	if prNumber == 123 {
		return &entity.PullRequest{}, nil
	}

	return nil, errors.New("not found")
}

func (m *githubClientMock) GetSelf(ctx context.Context) (*entity.Actor, error) {
	return &entity.Actor{
		Type:  constants.NormalUser,
		Login: "Namchee",
	}, nil
}

func (m *githubClientMock) GetComments(
	_ context.Context,
	_ *entity.Meta,
	number int,
) ([]*entity.Comment, error) {
	if number == 1 {
		return nil, errors.New("error")
	}

	if number == 2 {
		return []*entity.Comment{}, nil
	}

	return []*entity.Comment{
		{
			ID:   3,
			Body: "foobar",
			Author: entity.Actor{
				Login: "Namchee",
			},
		},
	}, nil
}

func (m *githubClientMock) GetPermissions(
	_ context.Context,
	_ *entity.Meta,
	user string,
) ([]string, error) {
	admin := constants.AdminUser
	writeOnly := "write"

	switch user {
	case "foo":
		return []string{admin}, nil
	case "bar":
		return []string{writeOnly}, nil
	default:
		return []string{}, errors.New("mock error")
	}
}

func (m *githubClientMock) GetCommits(
	ctx context.Context,
	_ *entity.Meta,
	event int,
) ([]*entity.Commit, error) {
	good := "feat(test): test something"
	bad := "this is bad"

	if event == 123 {
		return []*entity.Commit{
			{
				Message: good,
			},
		}, nil
	}

	if event == 69 {
		return []*entity.Commit{
			{
				Hash:    "e21b424",
				Message: bad,
			},
		}, nil
	}

	if event == 1 {
		return []*entity.Commit{}, errors.New("mock error")
	}

	return []*entity.Commit{}, nil
}

func (m *githubClientMock) CreateComment(
	_ context.Context,
	_ *entity.Meta,
	event int,
	_ string,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) EditComment(
	_ context.Context,
	_ *entity.Meta,
	id int,
	_ string,
) error {
	if id == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) Label(
	_ context.Context,
	_ *entity.Meta,
	event int,
	_ string,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("Error")
}

func (m *githubClientMock) Close(
	_ context.Context,
	_ *entity.Meta,
	event int,
) error {
	if event == 123 {
		return nil
	}

	return errors.New("error")
}

func (m *githubClientMock) GetIssue(
	_ context.Context,
	meta *entity.Meta,
	issueNumber int,
) (*entity.IssueReference, error) {
	if meta.Owner == "vitejs" && meta.Name == "vite" {
		return &entity.IssueReference{
			Number: issueNumber,
			Meta: *meta,
		}, nil
	}

	if issueNumber == 3 {
		return &entity.IssueReference{
			Number: issueNumber,
			Meta: entity.Meta{
				Owner: "Namchee",
				Name:  "conventional-pr",
			},
		}, nil
	}

	return nil, errors.New("mock error")
}

func (m *githubClientMock) GetIssueReferences(
	_ context.Context,
	_ *entity.Meta,
	prNumber int,
) ([]*entity.IssueReference, error) {
	switch prNumber {
	case 1:
		return []*entity.IssueReference{
			{
				Meta: entity.Meta{
					Owner: "Namchee",
					Name:  "conventional-pr",
				},
			},
		}, nil
	case 2:
		return []*entity.IssueReference{}, nil
	case 3:
		return []*entity.IssueReference{}, nil
	default:
		return nil, errors.New("mock error")
	}
}

// NewGithubClientMock creates a GitHub client mock for testing purposes
func NewGithubClientMock() internal.GithubClient {
	return &githubClientMock{}
}
