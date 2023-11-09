package sample

import (
	"dog-flow/flow"
	"fmt"
)

type FemaleHandlerNode struct {
	id string
	state State
}

func NewFemaleHandlerNode(id string) *FemaleHandlerNode {
	return &FemaleHandlerNode{
		id: id,
		state: INIT,
	}
}

// implement Node interface

func (n *FemaleHandlerNode) ID() string {
	return n.id
}

func (n *FemaleHandlerNode) Init(context flow.FlowContext) error {
	fmt.Printf("FemaleHandlerNode[%s] init\n", n.id)
	return nil
}

func (n *FemaleHandlerNode) Run(context flow.FlowContext) error {
	n.state = RUNNING
	fmt.Printf("FemaleHandlerNode[%s] run\n", n.id)
	n.state = DONE
	context.DoneNode(n.id)
	return nil
}

func (n * FemaleHandlerNode) IsDone(context flow.FlowContext) bool {
	return n.state == DONE
}

func (n *FemaleHandlerNode) IsRunning() bool {
	return n.state == RUNNING
}

// implement Handler interface

func (n *FemaleHandlerNode) OnEvent(context flow.FlowContext, event flow.Event) error {
	fmt.Printf("FemaleHandlerNode[%s] on event\n", n.id)
	return nil
}

func (n *FemaleHandlerNode) Enter(context flow.FlowContext) error {
	fmt.Printf("FemaleHandlerNode[%s] enter\n", n.id)
	return nil
}

func (n *FemaleHandlerNode) Leave(context flow.FlowContext) error {
	fmt.Printf("FemaleHandlerNode[%s] leave\n", n.id)
	return nil
}

type FemaleEdge struct {
	id string
	from flow.Node
	to flow.Node
}

func NewFemaleEdge(id string, from flow.Node, to flow.Node) *FemaleEdge {
	return &FemaleEdge{
		id: id,
		from: from,
		to: to,
	}
}

func (e *FemaleEdge) ID() string{
	return e.id
}

func (e *FemaleEdge) From() flow.Node {
	return e.from
}

func (e *FemaleEdge) To() flow.Node {
	return e.to
}

func (e *FemaleEdge) IsPassable(context flow.FlowContext) bool {
	formId := context.GetFlowId();

	data, err := FormServiceMockSingleton.LoadForm(formId)
	if err != nil {
		return false
	}

	age, ok := data["age"]
	sex, ok1 := data["sex"]
	if ok && ok1 {
		if age.(int) >= 18 && sex.(string) == "female" {
			fmt.Printf("FemaleEdge[%s] is passable\n", e.id)
			return true
		}
	}

	return false
}