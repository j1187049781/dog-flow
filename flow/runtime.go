package flow


const MAX_EVENT_QUEUE_SIZE = 100

type Runtime struct {
	flowInstance map[string]*FlowInstance
	eventQueue chan *Event
}

func NewRuntime() *Runtime {
	return &Runtime{
		flowInstance: make(map[string]*FlowInstance),
		eventQueue: make(chan *Event, MAX_EVENT_QUEUE_SIZE),
	}
}

func (r *Runtime) InitFlowInstance(flowId string, dag *Dag) error {
	r.flowInstance[flowId] = NewFlowInstance(flowId, dag)

	return nil
}

func (r *Runtime) StartFlowInstance(flowId string) error {
	flowInstance := r.flowInstance[flowId]

	if flowInstance == nil {
		return nil
	}

	return flowInstance.Start(r.eventQueue)
}

func (r *Runtime) StopFlowInstance(flowId string) error {
	flowInstance := r.flowInstance[flowId]

	if flowInstance == nil {
		return nil
	}

	return flowInstance.Stop()
}

func (r *Runtime) ReceiveEvent(event *Event) error {
	r.eventQueue <- event

	return nil
}

func (r *Runtime) Run()  {
	for {
		event := <- r.eventQueue
		r.handleEvent(event)
	}
}

func (r *Runtime) handleEvent(event *Event) error {
	flowInstance := r.flowInstance[(*event).ToFlowInstanceId()]

	if flowInstance == nil {
		return nil
	}

	return flowInstance.HandleEvent(event)
}