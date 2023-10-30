package runtime

type Signal struct {
	Code SignalCode
	NodeId string
}

type SignalCode int32

const (
	// StartSignal is the signal that is fired when a node is started.
	StartSignal SignalCode = iota
	// DoneSignal is the signal that is fired when a node is done.
	DoneSignal
)