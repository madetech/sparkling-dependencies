package github

import (
	"context"
	"github.com/google/go-github/v48/github"
	"github.com/madetech/sparkling-dependencies/internal/dealWithPullRequest"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type GitHubPresenter struct {
	Client *github.Client
}

func (g GitHubPresenter) Exit() {
	os.Exit(0)
}

func (g GitHubPresenter) PostComment(comment dealWithPullRequest.Comment) {
	splitRepository := strings.Split(comment.Repository, "/")
	owner := splitRepository[0]
	repo := splitRepository[1]
	context.Background()
	_, _, err := g.Client.Issues.CreateComment(context.Background(), owner, repo, int(comment.Number), &github.IssueComment{
		Body: &comment.Body,
	})
	if err != nil {
		panic(err)
	}
}

func New(token string) GitHubPresenter {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	presenter := GitHubPresenter{Client: client}
	return presenter
}
