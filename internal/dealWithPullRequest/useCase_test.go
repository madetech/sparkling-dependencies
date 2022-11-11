package dealWithPullRequest

import "testing"

type SpyPresenter struct {
	ExitCalled bool
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
}

func (uc dependencies) ExecuteWithSpy() SpyPresenter {
	presenter := SpyPresenter{}
	uc.Execute(&presenter)
	return presenter
}

func TestExitsIfNotPullRequestTargetEvent(t *testing.T) {
	t.Run("Can exit if not a pull request target event", func(t *testing.T) {
		sender := "dependabot[bot]"
		dealWithPullRequest := New(Event{Name: "not_expected", Sender: &sender})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitCalled(t)
	})

	t.Run("Does not exit if pull request target event", func(t *testing.T) {
		sender := "dependabot[bot]"
		dealWithPullRequest := New(Event{Name: "pull_request_target", Sender: &sender})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitNotCalled(t)
	})

	t.Run("Exits if not dependabot", func(t *testing.T) {
		sender := "craigjbass"
		dealWithPullRequest := New(Event{Name: "pull_request_target", Sender: &sender})
		presenter := dealWithPullRequest.ExecuteWithSpy()
		presenter.AssertExitCalled(t)
	})
}
