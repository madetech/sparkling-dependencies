package dealWithPullRequest

type Event struct {
	Name string
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
	if d.Event.Name != "pull_request_target" {
		presenter.Exit()
	}

}
