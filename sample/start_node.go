package sample

import (
	"dog-flow/flow"
	"fmt"
)

type State int

const (
	INIT State = iota
	RUNNING
	DONE
)

// start node
type StartNode struct {
	id string
	state State
}

func NewStartNode(id string) *StartNode {
	return &StartNode{
		id: id,
		state: INIT,
	}
}

// implement Node interface

func (n *StartNode) ID() string {
	return n.id
}

func (n *StartNode) Init(context flow.FlowContext) error {
	fmt.Printf("StartNode[%s] init\n", n.id)
	return nil
}

func (n *StartNode) Run(context flow.FlowContext) error {
	n.state = RUNNING
	fmt.Printf("StartNode[%s] run\n", n.id)
	return nil
}

func (n *StartNode) IsDone(context flow.FlowContext) bool {
	return n.state == DONE
}

func (n *StartNode) IsRunning() bool {
	return n.state == RUNNING
}

// implement Handler interface

func (n *StartNode) OnEvent(context flow.FlowContext, event flow.Event) error {
	fmt.Printf("StartNode[%s] on event\n", n.id)

	if event.ToNodeId() == n.id {

		// save form
		startNodeEvent, ok := event.(*StartNodeEvent)
		if !ok {
			return nil
		}

		FormServiceMockSingleton.SaveForm(startNodeEvent.flowInstanceId, startNodeEvent.data)

		// done node
		n.state = DONE
		context.DoneNode(n.id)
	}
	return nil
}

func (n *StartNode) Enter(context flow.FlowContext) error {
	fmt.Printf("StartNode[%s] enter\n", n.id)
	return nil
}

func (n *StartNode) Leave(context flow.FlowContext) error {
	fmt.Printf("StartNode[%s] leave\n", n.id)
	return nil
}

type StartNodeEvent struct {
	id string
	flowInstanceId string
	nodeId string
	// {"name": "zhangsan", "age": 18,"sex": female}
	data map[string]any
}

func NewStartNodeEvent(flowInstanceId string, nodeId string, data map[string]any) *StartNodeEvent {
	return &StartNodeEvent{
		id: flowInstanceId + "." + nodeId,
		flowInstanceId: flowInstanceId,
		nodeId: nodeId,
		data: data,
	}
}

func (e *StartNodeEvent) ID() string {
	return e.id
}

func (e *StartNodeEvent) ToFlowInstanceId() string {
	return e.flowInstanceId
}

func (e *StartNodeEvent) ToNodeId() string {
	return e.nodeId
}

