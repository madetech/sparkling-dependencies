package dealWithPullRequest

type PullRequest struct {
	Sender     string
	Number     uint32
	Repository string
}

type Event struct {
	Name        string
	PullRequest *PullRequest
}

type dependencies struct {
	Event Event
}

func New(event Event) *dependencies {
	return &dependencies{
		Event: event,
	}
}

type Comment struct {
	Body       string
	Repository string
	Number     uint32
}

type Presenter interface {
	Exit()
	PostComment(comment Comment)
}

func (d dependencies) Execute(presenter Presenter) {
	if d.isNotPullRequestTarget() || d.hasNoValidPullRequestData() {
		presenter.Exit()
		return
	}

	if d.author() != "dependabot[bot]" {
		presenter.Exit()
		return
	}

	presenter.PostComment(
		Comment{
			Repository: d.repository(),
			Body:       "@dependabot merge",
			Number:     d.prNumber(),
		},
	)
}

func (d dependencies) author() string {
	return d.Event.PullRequest.Sender
}

func (d dependencies) prNumber() uint32 {
	return d.Event.PullRequest.Number
}

func (d dependencies) repository() string {
	return d.Event.PullRequest.Repository
}

func (d dependencies) hasNoValidPullRequestData() bool {
	return d.Event.PullRequest == nil
}

func (d dependencies) isNotPullRequestTarget() bool {
	return d.Event.Name != "pull_request_target"
}
