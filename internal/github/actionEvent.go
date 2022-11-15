package github

import (
	"encoding/json"
	"github.com/madetech/sparkling-dependencies/internal/dealWithPullRequest"
	"os"
)

func GetEvent() dealWithPullRequest.Event {
	eventPayloadPath, _ := os.LookupEnv("GITHUB_EVENT_PATH")
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
	event := dealWithPullRequest.Event{
		Name: "pull_request_target",
		PullRequest: &dealWithPullRequest.PullRequest{
			Sender:     data.Sender.Login,
			Number:     data.PullRequest.Number,
			Repository: data.Repository.FullName,
			Title:      data.PullRequest.Title,
		},
	}
	return event
}

func readFile(eventPayloadPath string) []byte {
	file, err := os.ReadFile(eventPayloadPath)
	if err != nil {
		panic(err)
	}
	return file
}
