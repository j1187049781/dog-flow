package edge

import (
	"dog-flow/flow/dag/node"
)

type Edge interface {
	// ID returns the ID of the edge.
	ID() string
	From() node.Node
	To() node.Node
	isPassable() bool
}