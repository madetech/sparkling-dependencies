package dealWithPullRequest

type PullRequest struct {
	Sender string
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

	if d.Event.PullRequest.Sender != "dependabot[bot]" {
		presenter.Exit()
		return
	}

	presenter.PostComment(Comment{Repository: "madetech/wow", Body: "@dependabot merge", Number: 1})
}

func (d dependencies) hasNoValidPullRequestData() bool {
	return d.Event.PullRequest == nil
}

func (d dependencies) isNotPullRequestTarget() bool {
	return d.Event.Name != "pull_request_target"
}
