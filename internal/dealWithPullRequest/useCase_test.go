package dealWithPullRequest

import "testing"

type SpyPresenter struct {
	ExitCalled bool
}

func (s *SpyPresenter) Exit() {
	s.ExitCalled = true
}

func (s *SpyPresenter) AssertExitNotCalled(t *testing.T) {
	if s.ExitCalled {
		t.Errorf("Expected Exit to not have been called.")
	}
}

func (s *SpyPresenter) AssertExitCalled(t *testing.T) {
	if !s.ExitCalled {
		t.Errorf("Expected Exit to Have been called.")
	}
}

func TestExitsIfNotPullRequestTargetEvent(t *testing.T) {
	t.Run("Can exit if not a pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "not_expected"})
		presenter := SpyPresenter{}
		dealWithPullRequest.Execute(&presenter)
		presenter.AssertExitCalled(t)
	})

	t.Run("Does not exit if pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "pull_request_target"})
		presenter := SpyPresenter{}
		dealWithPullRequest.Execute(&presenter)
		presenter.AssertExitNotCalled(t)
	})
}
