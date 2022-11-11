package dealWithPullRequest

import "testing"

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

func TestExitsIfNotPullRequestTargetEvent(t *testing.T) {
	t.Run("Can exit if not a pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "not_expected"})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitCalled(t)
	})

	t.Run("Does not exit if pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &PullRequest{Sender: "dependabot[bot]"}})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitNotCalled(t)
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
		}
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &pr})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertCommented(t, "madetech/wow", 1, "@dependabot merge")
	})

	t.Run("Tells dependabot to merge 2", func(t *testing.T) {
		pr := PullRequest{
			Sender:     "dependabot[bot]",
			Number:     2,
			Repository: "madetech/cool",
		}
		dealWithPullRequest := New(Event{Name: "pull_request_target", PullRequest: &pr})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertCommented(t, "madetech/cool", 2, "@dependabot merge")
	})
}
