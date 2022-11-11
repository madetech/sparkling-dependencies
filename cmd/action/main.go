package main

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-github/v48/github"
	"github.com/madetech/sparkling-dependencies/internal/dealWithPullRequest"
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
	_, _, err := g.client.PullRequests.CreateComment(context.Background(), owner, repo, int(comment.Number), &github.PullRequestComment{
		Body: &comment.Body,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	eventPayloadPath, _ := os.LookupEnv("GITHUB_EVENT_PATH")
	//token, _ := os.LookupEnv("INPUT_GITHUB-TOKEN")
	//
	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: token},
	//)
	//tc := oauth2.NewClient(ctx, ts)
	//
	//client := github.NewClient(tc)

	file := readFile(eventPayloadPath)
	spew.Dump(file)
	//var data struct {
	//	EventName string `json:"eventName"`
	//	Payload   struct {
	//		sender struct {
	//			login string `json:"login"`
	//		} `json:"sender"`
	//	} `json:"payload"`
	//}
	//err := json.Unmarshal(file, data)
	//if err != nil {
	//	panic(err)
	//}
	//useCase := dealWithPullRequest.New(dealWithPullRequest.Event{})
	//useCase.Execute(GitHubPresenter{client: client})
}

func readFile(eventPayloadPath string) []byte {
	file, err := os.ReadFile(eventPayloadPath)
	if err != nil {
		panic(err)
	}
	return file
}
