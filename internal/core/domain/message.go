package domain

// RequestVotePayload represents the arguments sent in a RequestVote RPC.
type RequestVotePayload struct {
	// Term is the candidate’s current term.
	Term Term

	// LastLogTerm is the term of the candidate’s last log entry.
	LastLogTerm Term

	// LastLogIndex is the index of the candidate’s last log entry.
	LastLogIndex int64

	// CandidateID is the identifier of the requesting candidate.
	CandidateID NodeID
}

// AppendEntriesPayload represents the arguments sent in an AppendEntries RPC.
type AppendEntriesPayload struct {
	// Term is the leader’s current term.
	Term Term

	// PrevLogTerm is the term of the log entry immediately before the new entries.
	PrevLogTerm Term

	// PrevLogIndex is the index of the log entry immediately before the new entries.
	PrevLogIndex int64

	// LeaderCommitIndex is the leader’s commit index.
	LeaderCommitIndex int64

	// LeaderID identifies the leader sending this RPC.
	LeaderID NodeID

	// LogEntries holds the new log entries to be appended.
	// Empty for heartbeats (no-op AppendEntries).
	LogEntries []*LogEntry
}
