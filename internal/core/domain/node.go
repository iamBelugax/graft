package domain

import (
	"time"
)

// Peer represents metadata about another node in the cluster.
type Peer struct {
	// ID is the unique identifier for the peer node.
	ID string

	// Address is the network location (host:port) of the peer node.
	Address string
}

// Node represents a single Raft participant.
type Node struct {
	// ID is the unique identifier for the node.
	ID string

	// Address holds the network address (host:port) used to communicate with this node.
	Address string

	// CurrentTerm stores the latest term number known to this node.
	CurrentTerm Term

	// VotedFor records the candidate ID that this node voted for in its CurrentTerm.
	VotedFor string

	// State indicates whether the node is a Follower, Candidate, or Leader.
	State NodeState

	// LogEntries stores the sequence of log entries replicated via the Raft protocol.
	LogEntries []*LogEntry

	// Peers holds references to other nodes in the cluster.
	Peers []*Peer

	// Leader points to the current known leader of the cluster.
	Leader *Peer

	// ElectionTimeout defines how long a follower waits without hearing from a leader
	// before starting a new election.
	ElectionTimeout time.Duration

	// HeartbeatInterval defines how often the leader sends heartbeats
	// to maintain authority and prevent new elections.
	HeartbeatInterval time.Duration

	// LastHeartbeatAt records the last time a heartbeat was received.
	// Used by followers to detect leader failures.
	LastHeartbeatAt time.Time
}

// Is checks if the node is in the given state.
func (n *Node) Is(state NodeState) bool {
	return n.State.IsValid() && state.IsValid() && n.State == state
}

// IsFollower checks if the node is in the Follower state.
func (n *Node) IsFollower() bool {
	return n.Is(FOLLOWER)
}

// IsCandidate checks if the node is in the Candidate state.
func (n *Node) IsCandidate() bool {
	return n.Is(CANDIDATE)
}

// IsLeader checks if the node is in the Leader state.
func (n *Node) IsLeader() bool {
	return n.Is(LEADER)
}
