package main

import (
	"context"
	"encoding/json"
	_ "github.com/breml/rootcerts"
	"github.com/google/go-github/v48/github"
	"github.com/madetech/sparkling-dependencies/internal/dealWithPullRequest"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type GitHubPresenter struct {
	client *github.Client
}

func (g GitHubPresenter) Exit() {
	os.Exit(0)
}

func (g GitHubPresenter) PostComment(comment dealWithPullRequest.Comment) {
	splitRepository := strings.Split(comment.Repository, "/")
	owner := splitRepository[0]
	repo := splitRepository[1]
	context.Background()
	_, _, err := g.client.Issues.CreateComment(context.Background(), owner, repo, int(comment.Number), &github.IssueComment{
		Body: &comment.Body,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	eventPayloadPath, _ := os.LookupEnv("GITHUB_EVENT_PATH")
	token, _ := os.LookupEnv("INPUT_GITHUB-TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	file := readFile(eventPayloadPath)

	var data struct {
		Sender struct {
			Login string `json:"login"`
		} `json:"sender"`
		PullRequest struct {
			Number uint32 `json:"number"`
			Title  string `json:"title"`
		} `json:"pull_request"`
		Repository struct {
			FullName string `json:"full_name"`
		} `json:"repository"`
	}
	err := json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	useCase := dealWithPullRequest.New(dealWithPullRequest.Event{
		Name: "pull_request_target",
		PullRequest: &dealWithPullRequest.PullRequest{
			Sender:     data.Sender.Login,
			Number:     data.PullRequest.Number,
			Repository: data.Repository.FullName,
			Title:      data.PullRequest.Title,
		},
	})
	useCase.Execute(GitHubPresenter{client: client})
}

func readFile(eventPayloadPath string) []byte {
	file, err := os.ReadFile(eventPayloadPath)
	if err != nil {
		panic(err)
	}
	return file
}
