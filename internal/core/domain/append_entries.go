package domain

import graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"

// AppendEntriesRequest represents the arguments sent in an AppendEntries RPC.
type AppendEntriesRequest struct {
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
	LogEntries []*LogEntry
}

func (aep *AppendEntriesRequest) Encode() *graftpb.AppendEntriesRequest {
	pb := graftpb.AppendEntriesRequest{
		LeaderId:     aep.LeaderID,
		PrevLogIndex: aep.PrevLogIndex,
		Term:         aep.Term.AsUint64(),
		LeaderCommit: aep.LeaderCommitIndex,
		PrevLogTerm:  aep.PrevLogTerm.AsUint64(),
		LogEntries:   make([]*graftpb.LogEntry, len(aep.LogEntries)),
	}

	for i, entry := range aep.LogEntries {
		pb.LogEntries[i] = entry.Encode()
	}

	return &pb
}

func (aep *AppendEntriesRequest) Decode(pb *graftpb.AppendEntriesRequest) {
	aep.Term = Term(pb.Term)
	aep.LeaderID = pb.LeaderId
	aep.PrevLogIndex = pb.PrevLogIndex
	aep.PrevLogTerm = Term(pb.PrevLogTerm)
	aep.LeaderCommitIndex = pb.LeaderCommit
	aep.LogEntries = make([]*LogEntry, len(pb.LogEntries))

	for i, entry := range pb.LogEntries {
		logEntry := &LogEntry{}
		logEntry.Decode(entry)
		aep.LogEntries[i] = logEntry
	}
}

// AppendEntriesResponse represents the reply to an AppendEntries RPC.
type AppendEntriesResponse struct {
	// Term is the responder’s current term.
	Term Term

	// Success indicates whether the entries were appended successfully.
	Success bool

	// ConflictTerm is the term of the conflicting entry if Success is false.
	ConflictTerm Term

	// ConflictIndex is the index of the first entry with ConflictTerm.
	ConflictIndex int64

	// LastLogIndex is the index of the follower’s last log entry.
	LastLogIndex int64
}

func (aer *AppendEntriesResponse) Encode() *graftpb.AppendEntriesResponse {
	return &graftpb.AppendEntriesResponse{
		Success:       aer.Success,
		LastLogIndex:  aer.LastLogIndex,
		ConflictIndex: aer.ConflictIndex,
		Term:          aer.Term.AsUint64(),
		ConflictTerm:  aer.ConflictTerm.AsUint64(),
	}
}

func (aer *AppendEntriesResponse) Decode(pb *graftpb.AppendEntriesResponse) {
	aer.Term = Term(pb.Term)
	aer.Success = pb.Success
	aer.LastLogIndex = pb.LastLogIndex
	aer.ConflictIndex = pb.ConflictIndex
	aer.ConflictTerm = Term(pb.ConflictTerm)
}
