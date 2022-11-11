package dealWithPullRequest

import (
	"testing"
)

type SpyPresenter struct {
	ExitCalled bool
	Comment    *Comment
}

func (s *SpyPresenter) PostComment(comment Comment) {
	s.Comment = &comment
}

func (s *SpyPresenter) Exit() {
	s.ExitCalled = true
}

func (s *SpyPresenter) AssertExitNotCalled(t *testing.T) {
	t.Helper()
	if s.ExitCalled {
		t.Errorf("Expected Exit to not have been called.")
	}
}

func (s *SpyPresenter) AssertExitCalled(t *testing.T) {
	t.Helper()
	if !s.ExitCalled {
		t.Errorf("Expected Exit to Have been called.")
	}

	if s.Comment != nil {
		t.Errorf("Expect no comment to have been made")
	}
}

func (s *SpyPresenter) AssertCommented(t *testing.T, repo string, number uint32, content string) {
	t.Helper()
	if s.Comment.Body != content {
		t.Errorf("Expected body to be %#v, got %#v", content, s.Comment.Body)
	}

	if s.Comment.Repository != repo {
		t.Errorf("Expected repository to be %#v, got %#v", repo, s.Comment.Repository)
	}

	if s.Comment.Number != number {
		t.Errorf("Expected issue number to be %#v, got %#v", number, s.Comment.Number)
	}
}

func (uc dependencies) ExecuteWithSpy() SpyPresenter {
	presenter := SpyPresenter{}
	uc.Execute(&presenter)
	return presenter
}

func TestDealWithPullRequest(t *testing.T) {
	t.Run("Can exit if not a pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "not_expected"})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitCalled(t)
	})

	t.Run("Exits if not dependabot", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &PullRequest{Sender: "craigjbass"}})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitCalled(t)
	})

	t.Run("Tells dependabot to merge", func(t *testing.T) {
		pr := PullRequest{
			Sender:     "dependabot[bot]",
			Number:     1,
			Repository: "madetech/wow",
			Title:      "Bump python from 3.10.7 to 3.11.0",
		}
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &pr})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertCommented(t, "madetech/wow", 1, "@dependabot merge")
		presenter.AssertExitNotCalled(t)
	})

	t.Run("Tells dependabot to merge 2", func(t *testing.T) {
		pr := PullRequest{
			Sender:     "dependabot[bot]",
			Number:     2,
			Repository: "madetech/cool",
			Title:      "Bump python from 3.10.7 to 3.11.0",
		}
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &pr})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertCommented(t, "madetech/cool", 2, "@dependabot merge")
	})

	t.Run("Tells dependabot to close", func(t *testing.T) {
		pr := PullRequest{
			Sender:     "dependabot[bot]",
			Number:     1,
			Repository: "madetech/wow",
			Title:      "Bump python from 3.10.7 to 3.11.0rc2",
		}
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &pr})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertCommented(t, "madetech/wow", 1, "@dependabot close")
	})
}

func assertVersionIsNil(t *testing.T, input string) {
	t.Helper()
	version := getVersionFromTitle(input)
	if version != nil {
		t.Errorf("Expected nil got %#v", version)
	}
}

func assertExtractsVersion(t *testing.T, input string, expected string) {
	t.Helper()
	version := *getVersionFromTitle(input)
	if version != expected {
		t.Errorf("Expected %#v got %#v", expected, version)
	}
}

func TestParseTitle(t *testing.T) {
	t.Run("can parse", func(t *testing.T) {
		assertVersionIsNil(t, "")
		assertVersionIsNil(t, "Some title with no version")
		assertExtractsVersion(t, "Bump dependency from 1.0.0 to 2.0.0", "2.0.0")
		assertExtractsVersion(t, "Bump dependency from 1.0.0 to 3.0.0", "3.0.0")
		assertExtractsVersion(t, "Bump to from 1.0.0 to 3.0.0", "3.0.0")
		assertExtractsVersion(t, "Bump python from 3.10.7 to 3.11.0rc2", "3.11.0rc2")
		assertExtractsVersion(t, "Build(deps-dev): Bump boto3 from 1.26.6 to 1.26.7 in /some/thing", "1.26.7")
		assertExtractsVersion(t, "Build(deps): Bump boto3 from 1.26.4 to 1.26.5 in /a-thing", "1.26.5")
		assertExtractsVersion(t, "Bump github.com/go-redis/redis/v9 from 9.0.0-beta.3 to 9.0.0-rc.1 in /one/two-three", "9.0.0-rc.1")
		assertExtractsVersion(t, "Bump envoyproxy/envoy from v1.23.1 to v1.24.0 in /four/five-six", "1.24.0")
		assertExtractsVersion(t, "Bump google.golang.org/grpc from 1.50.0 to 1.50.1 in /a/b/c/d/e/f/g", "1.50.1")
	})
}
