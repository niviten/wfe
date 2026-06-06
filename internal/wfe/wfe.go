package wfe

type Workflow struct {
	actions []Action
}

func New() *Workflow {
	return &Workflow{actions: make([]Action, 0)}
}

func (wf *Workflow) Insert(action Action) *Workflow {
	wf.actions = append(wf.actions, action)
	return wf
}

func (wf *Workflow) Execute(input any) (any, error) {
	data := input
	for _, action := range wf.actions {
		output, err := action.RunAction(data)
		if err != nil {
			return output, err
		}
		data = output
	}
	return data, nil
}

type Action interface {
	RunAction(inputs any) (any, error)
}

type RunnableAction struct {
	fn func(inputs any) (any, error)
}

func NewRunnableAction(fn func(inputs any) (any, error)) *RunnableAction {
	return &RunnableAction{fn: fn}
}

func (action *RunnableAction) RunAction(inputs any) (any, error) {
	return action.fn(inputs)
}
