package flow

import "dog-flow/util"

type FlowInstance struct {
	Id  string
	Dag *Dag
	nodeDoneQ util.Queue[string]
}

func NewFlowInstance(id string, dag *Dag) *FlowInstance {
	f := &FlowInstance{
		Id:        id,
		Dag:       dag,
		nodeDoneQ: util.NewQueue[string](),
	}
	return f
}

// implement FlowContext
func (f *FlowInstance) GetFlowId() string {
	return f.Id
}

func (f *FlowInstance) DoneNode(nodeId string) error {
	f.Dag.Nodes[nodeId].Leave(f)

	f.nodeDoneQ.Push(nodeId)
	return nil
}

func (f *FlowInstance) Start() {
	// 初始化所有的节点状态
	for _, node := range f.Dag.Nodes {
		node.Init(f)
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
			// 启动节点
			f.execRun(node)
		}
	}

	f.forward()
}

func (f *FlowInstance)HandleEvent(event Event) {
	nodeId := event.ToNodeId()

	f.Dag.Nodes[nodeId].OnEvent(f, event)
	
	f.forward()
}

func (f *FlowInstance) forward() {

	for !f.nodeDoneQ.Empty() {
		nodeId := f.nodeDoneQ.Pop()
		edges := f.Dag.Edges[nodeId]
		for _, edge := range edges {
			if edge.IsPassable(f) {
				f.execRun(edge.To())
			}
		}
	}
}

func (f *FlowInstance) execRun(node Node) {
	node.Enter(f)

	node.Run(f)
}
