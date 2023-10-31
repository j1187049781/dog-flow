package flow

import (
	"errors"
)

type Dag struct {
	// Nodes is a map of node IDs to nodes.
	Nodes map[string]Node

	// key: ndoeID, value: edge list. 
	// Edges is a map of node IDs to edges.
	Edges map[string][]Edge
}

// New returns a new, empty DAG.
func New() *Dag {
	return &Dag{
		Nodes: make(map[string]Node),
		Edges: make(map[string][]Edge),
	}
}

// AddNode adds a node to the DAG.
func (d *Dag) AddNode(n Node) error{
	id := n.ID()

	if _, ok := d.Nodes[id]; ok {
		return errors.New("node already exists");
	}

	d.Nodes[id] = n

	return nil
}

// AddEdge adds an edge to the DAG.
func (d *Dag) AddEdge(e Edge) error {
	from := e.From()
	to := e.To()

	if _, ok := d.Nodes[from.ID()]; !ok {
		return errors.New("from node not exists");
	} 

	if _, ok := d.Nodes[to.ID()]; !ok {
		return errors.New("to node not exists");
	}

	if _, ok := d.Edges[from.ID()]; !ok {
		d.Edges[from.ID()] = make([]Edge, 0)
	}

	d.Edges[from.ID()] = append(d.Edges[from.ID()], e)

	return nil
}