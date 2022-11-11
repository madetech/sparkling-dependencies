package dealWithPullRequest

type Event struct {
	Name   string
	Sender *string
}

type dependencies struct {
	Event Event
}

func New(event Event) *dependencies {
	return &dependencies{
		Event: event,
	}
}

type Presenter interface {
	Exit()
}

func (d dependencies) Execute(presenter Presenter) {
	if d.Event.Sender == nil {
		return
	}

	if d.Event.Name != "pull_request_target" {
		presenter.Exit()
	}

	if *d.Event.Sender != "dependabot[bot]" {
		presenter.Exit()
	}
}
