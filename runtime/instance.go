package runtime

import (
	"dog-flow/util"
	"dog-flow/flow/dag"
	"dog-flow/runtime/event"
)

type FlowInstance struct {
	Id string
	Dag *dag.Dag 
	RunningNodes util.Set[string]
}

func NewFlowInstance(id string, dag *dag.Dag) *FlowInstance {
	return &FlowInstance{
		Id: id,
		Dag: dag,
		RunningNodes: util.NewSet[string](),
	}
}

func (f *FlowInstance) Start() error {
	return nil
}

func (f *FlowInstance) Stop() error {
	return nil
}

func (f *FlowInstance) HandleEvent(event *event.Event) error {
	return nil
}