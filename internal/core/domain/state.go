package domain

import graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"

// Term represents the logical time in Raft.
// It is a monotonically increasing number.
// Each new election increments the term.
type Term uint64

func (t Term) AsUint64() uint64 {
	return uint64(t)
}

// NodeState represents the state of a node in the Raft consensus algorithm.
type NodeState uint8

const (
	// FOLLOWER is the default state of a node.
	// It follows the leader and responds to requests from candidates/leaders.
	FOLLOWER NodeState = iota

	// CANDIDATE is when a node times out waiting for heartbeats
	// and starts a new election to try to become the leader.
	CANDIDATE

	// LEADER is the node that won the election.
	// It handles client requests and replicates log entries.
	LEADER
)

func (s NodeState) AsUint8() uint8 {
	return uint8(s)
}

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

func (s NodeState) IsValid() bool {
	return s >= FOLLOWER && s <= LEADER
}

func (s NodeState) Encode() graftpb.NodeState {
	return graftpb.NodeState(s)
}

func (s *NodeState) Decode(pb graftpb.NodeState) {
	*s = NodeState(pb)
}
