package domain

// NodeID uniquely identifies a node in the Raft cluster.
type NodeID string

// Node represents a single Raft participant.
type Node struct {
	// ID is the unique identifier for the node.
	ID NodeID

	// Address holds the network address (host:port) used to communicate with this node.
	Address string

	// CurrentTerm stores the latest term number known to this node.
	CurrentTerm Term

	// VotedFor records the candidate ID that this node voted for in its CurrentTerm.
	VotedFor NodeID

	// State indicates whether the node is a Follower, Candidate, or Leader.
	State NodeState

	// LogEntries stores the sequence of log entries replicated via the Raft protocol.
	LogEntries []*LogEntry
}
