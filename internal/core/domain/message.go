package domain

import graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"

// RequestVotePayload represents the arguments sent in a RequestVote RPC.
type RequestVotePayload struct {
	// Term is the candidate’s current term.
	Term Term

	// LastLogTerm is the term of the candidate’s last log entry.
	LastLogTerm Term

	// LastLogIndex is the index of the candidate’s last log entry.
	LastLogIndex int64

	// CandidateID is the identifier of the requesting candidate.
	CandidateID string
}

// ToPB converts a domain level RequestVotePayload into its protobuf representation.
func (rvp *RequestVotePayload) ToPB() *graftpb.RequestVoteRequest {
	return &graftpb.RequestVoteRequest{
		Term:         uint64(rvp.Term),
		LastLogTerm:  uint64(rvp.LastLogTerm),
		LastLogIndex: rvp.LastLogIndex,
		CandidateId:  rvp.CandidateID,
	}
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
	LeaderID string

	// LogEntries holds the new log entries to be appended.
	// Empty for heartbeats (no-op AppendEntries).
	LogEntries []*LogEntry
}

// ToPB converts a domain level AppendEntriesPayload into its protobuf representation.
func (aep *AppendEntriesPayload) ToPB() *graftpb.AppendEntriesRequest {
	pb := graftpb.AppendEntriesRequest{
		Term:         uint64(aep.Term),
		LeaderId:     aep.LeaderID,
		PrevLogIndex: aep.PrevLogIndex,
		PrevLogTerm:  uint64(aep.PrevLogTerm),
		LeaderCommit: aep.LeaderCommitIndex,
		LogEntries:   make([]*graftpb.LogEntry, len(aep.LogEntries)),
	}

	for i, entry := range aep.LogEntries {
		pb.LogEntries[i] = entry.ToPB()
	}

	return &pb
}
