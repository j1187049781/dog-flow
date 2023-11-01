package flow

import "dog-flow/util"


type FlowInstance struct {
	Id string
	Dag *Dag

	nodeDoneQ util.Queue[string]

	FlowContext *FlowContext
}

func NewFlowInstance(id string, dag *Dag) *FlowInstance {
	f := &FlowInstance{
		Id: id,
		Dag: dag,
		nodeDoneQ: util.NewQueue[string](),
	}
	f.FlowContext = NewFlowContext(f, nil)
	return f
}

func (f *FlowInstance) Start() {
	// 初始化所有的节点状态
	for _, node := range f.Dag.Nodes {
		node.Done(nil)
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
		}	
	}
}