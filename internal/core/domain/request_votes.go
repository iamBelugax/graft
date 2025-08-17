package domain

import graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"

// RequestVoteRequest represents the arguments sent in a RequestVote RPC.
type RequestVoteRequest struct {
	Term         Term   // Term is the candidate’s current term.
	LastLogTerm  Term   // LastLogTerm is the term of the candidate’s last log entry.
	LastLogIndex int64  // LastLogIndex is the index of the candidate’s last log entry.
	CandidateID  string // CandidateID is the identifier of the requesting candidate.
}

// Encode converts a domain RequestVotePayload into its protobuf representation.
func (rvp *RequestVoteRequest) Encode() *graftpb.RequestVoteRequest {
	return &graftpb.RequestVoteRequest{
		CandidateId:  rvp.CandidateID,
		LastLogIndex: rvp.LastLogIndex,
		Term:         rvp.Term.AsUint64(),
		LastLogTerm:  rvp.LastLogTerm.AsUint64(),
	}
}

// Decode populates a domain RequestVotePayload from its protobuf representation.
func (rvp *RequestVoteRequest) Decode(pb *graftpb.RequestVoteRequest) {
	rvp.Term = Term(pb.Term)
	rvp.CandidateID = pb.CandidateId
	rvp.LastLogIndex = pb.LastLogIndex
	rvp.LastLogTerm = Term(pb.LastLogTerm)
}

// RequestVoteResponse represents the reply to a RequestVote RPC.
type RequestVoteResponse struct {
	Term        Term // Term is the current term of the responder.
	VoteGranted bool // VoteGranted indicates whether the vote was granted.
}

// Encode converts a domain RequestVoteResponse into its protobuf representation.
func (rvr *RequestVoteResponse) Encode() *graftpb.RequestVoteResponse {
	return &graftpb.RequestVoteResponse{
		Term:        rvr.Term.AsUint64(),
		VoteGranted: rvr.VoteGranted,
	}
}

// Decode populates a domain RequestVoteResponse from its protobuf representation.
func (rvr *RequestVoteResponse) Decode(pb *graftpb.RequestVoteResponse) {
	rvr.Term = Term(pb.Term)
	rvr.VoteGranted = pb.VoteGranted
}
