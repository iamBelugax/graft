package domain

import graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"

// Term represents the logical time in Raft.
// It is a monotonically increasing number used to identify
// the "era" of leadership. Each new election increments the term.
type Term uint64

// AsUint64 returns the Term as a uint64.
func (t Term) AsUint64() uint64 {
	return uint64(t)
}

// State represents the state of a node in the Raft consensus algorithm.
type State uint8

const (
	// FOLLOWER is the default state of a node.
	// It follows the leader and responds to requests from candidates/leaders.
	FOLLOWER State = iota

	// CANDIDATE is when a node times out waiting for heartbeats
	// and starts a new election to try to become the leader.
	CANDIDATE

	// LEADER is the node that won the election.
	// It handles client requests and replicates log entries.
	LEADER
)

// AsUint8 returns the State as a uint8.
func (s State) AsUint8() uint8 {
	return uint8(s)
}

// String returns a human-readable representation of the State.
func (s State) String() string {
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

// IsValid checks whether the State is within the valid Raft states.
func (s State) IsValid() bool {
	return s >= FOLLOWER && s <= LEADER
}

// Encode converts a domain State into its protobuf representation.
func (s State) Encode() graftpb.NodeState {
	return graftpb.NodeState(s)
}

// Decode populates a domain State from its protobuf representation.
func (s *State) Decode(pb graftpb.NodeState) {
	*s = State(pb)
}
