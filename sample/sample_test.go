package sample_test

import (
	"dog-flow/flow"
	"dog-flow/sample"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFemaleDag(t *testing.T) {

	// build FemalDag
	startNode := sample.NewStartNode("start")
	femaleHandlerNode := sample.NewFemaleHandlerNode("femaleHandler")
	edge := sample.NewFemaleEdge("famaleEdge",startNode, femaleHandlerNode)
	dag := flow.New()
	dag.AddNode(startNode)
	dag.AddNode(femaleHandlerNode)
	dag.AddEdge(edge)

	instance := flow.NewFlowInstance("form1", dag)
	instance.Start()

	data := map[string]interface{}{
		"sex": "female",
		"age": 18,
	}
	instance.HandleEvent(sample.NewStartNodeEvent("form1", "start", data))

	assert.Equal(t, femaleHandlerNode.IsDone(instance), true)
}