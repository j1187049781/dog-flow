package flow


type FlowInstance struct {
	Id string
	Dag *Dag 
	States map[string]*State

	eventQueue chan<- *Event
}

func NewFlowInstance(id string, dag *Dag) *FlowInstance {
	return &FlowInstance{
		Id: id,
		Dag: dag,
		States: make(map[string]*State),
	}
}

func (f *FlowInstance) Start(eventQueue chan<- *Event) error {

	f.eventQueue = eventQueue

	// 初始化所有的节点状态
	for _, node := range f.Dag.Nodes {
		f.States[node.ID()] = NewState(node.ID())
	}
	
	// 启动所有的入度为0的节点
	InDegree := make(map[string]int)
	for _, node := range f.Dag.Nodes {
		InDegree[node.ID()] = 0
	}

	for _, edges := range f.Dag.Edges {
		for _, edge := range edges {
			toId := edge.To().ID()
			InDegree[toId] += 1
		}
	}

	for _, node := range f.Dag.Nodes {
		if InDegree[node.ID()] == 0 {
			f.runANode(node.ID())
		}	
	}

	return nil
}

func (f *FlowInstance) Stop() error {
	return nil
}

func (f *FlowInstance) HandleEvent(event *Event) error {

	if err := (*event).Handle(NewFlowContext(nil, f, nil)) ; err != nil{
		return err
	}

	
	return nil
}


func (f *FlowInstance) TriggerNextNodes(fromNodeId string) error {
	edges := f.Dag.Edges[fromNodeId]
	for _, edge := range edges {
		// TODO
		toId := edge.To().ID()
		f.runANode(toId)
	}
	return nil
}

func (f *FlowInstance) runANode(nodeId string) error {
	f.States[nodeId].SetRunning()
	return nil
}

func (f *FlowInstance) doneANode(nodeId string) error {
	f.States[nodeId].SetStop()
	return nil
}