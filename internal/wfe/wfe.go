package wfe

type Workflow struct {
	actions   []Action
	startNode *Node
	endNode   *Node
}

func New() *Workflow {
	return &Workflow{
		startNode: nil,
		endNode:   nil,
	}
}

func (wf *Workflow) Insert(action Action) *Workflow {
	node := &Node{action: action, next: nil}

	if wf.startNode == nil {
		wf.startNode = node
		wf.endNode = node
		return wf
	}

	wf.endNode.next = node
	wf.endNode = node
	return wf
}

func (wf *Workflow) Execute(input any) (any, error) {
	data := input
	node := wf.startNode
	for node != nil {
		output, err := node.action.RunAction(data)
		if err != nil {
			return output, err
		}
		data = output
		node = node.next
	}
	return data, nil
}

type Node struct {
	action Action
	next   *Node
}

type Action interface {
	RunAction(input any) (any, error)
}

type RunnableAction struct {
	fn func(input any) (any, error)
}

func NewRunnableAction(fn func(input any) (any, error)) *RunnableAction {
	return &RunnableAction{fn: fn}
}

func (action *RunnableAction) RunAction(input any) (any, error) {
	return action.fn(input)
}
