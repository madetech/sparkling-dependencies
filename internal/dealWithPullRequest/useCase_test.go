package dealWithPullRequest

import "testing"

type SpyPresenter struct {
	ExitCalled bool
}

func (s *SpyPresenter) Exit() {
	s.ExitCalled = true
}

func TestExitsIfNotPullRequestTargetEvent(t *testing.T) {
	t.Run("Can exit if not a pull request target event", func(t *testing.T) {
		dealWithPullRequest := New(Event{Name: "not_expected"})
		presenter := SpyPresenter{}
		dealWithPullRequest.Execute(&presenter)
		if !presenter.ExitCalled {
			t.Errorf("Expected Exit to Have been called.")
		}
	})
}
