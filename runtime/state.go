package runtime

type StateCode int32

const (
	STATE_CODE_INIT StateCode = iota
	STATE_CODE_RUNNING
	STATE_CODE_STOPPED
) 

type State struct {
	NodeId string
	Code StateCode
}

func NewState(nodeId string) *State {
	return &State{
		NodeId: nodeId,
		Code: STATE_CODE_INIT,
	}
}

func (s *State) SetRunning() error {
	s.Code = STATE_CODE_RUNNING
	return nil
}

func (s *State) SetStop() error {
	s.Code = STATE_CODE_STOPPED
	return nil
}

func (s *State) IsRunning() bool {
	return s.Code == STATE_CODE_RUNNING
}