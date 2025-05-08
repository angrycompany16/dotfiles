package networking

import elevalgo "sanntidslab/elevalgo"

// Contains all information needed to ensure consistency among peers, should always be
// sent periodically
type Heartbeat struct {
	Uptime          int64
	SenderId        string
	State           elevalgo.Elevator
	PendingRequests PendingRequestList
	WorldView       map[string]peer
}

func newHeartbeat(node node) Heartbeat {
	return Heartbeat{
		SenderId:        nodeID,
		Uptime:          uptime,
		State:           node.state,
		PendingRequests: node.pendingRequestList,
		WorldView:       node.peers,
	}
}
