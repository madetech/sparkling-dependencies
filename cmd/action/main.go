package main

import (
	_ "github.com/breml/rootcerts"
	"github.com/madetech/sparkling-dependencies/internal/dealWithPullRequest"
	"github.com/madetech/sparkling-dependencies/internal/github"
	"os"
)

func main() {
	useCase := dealWithPullRequest.New(github.GetEvent())

	token, _ := os.LookupEnv("INPUT_GITHUB-TOKEN")
	useCase.Execute(github.New(token))
}
