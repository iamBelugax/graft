package domain

// Term represents the logical time in Raft.
// It is a monotonically increasing number used to identify
// the "era" of leadership. Each new election increments the term.
type Term uint64

// NodeState represents the state of a node in the Raft consensus algorithm.
// A node can either be a follower, candidate, or leader depending on the election state.
type NodeState uint8

const (
	// The default state of a node.
	// Follows the leader and responds to requests from candidates/leaders.
	FOLLOWER NodeState = iota

	// A node becomes a candidate when it times out waiting for heartbeats.
	// It then starts a new election to try to become the leader.
	CANDIDATE

	// The node that wins the election becomes the leader.
	// Responsible for handling client requests and replicating log entries.
	LEADER
)

// String returns a human readable representation of the NodeState.
func (s NodeState) String() string {
	switch s {
	case FOLLOWER:
		return "Follower"
	case CANDIDATE:
		return "Candidate"
	case LEADER:
		return "Leader"
	default:
		return "Unknown"
	}
}
