package domain

import (
	graftpb "github.com/iamBelugax/graft/internal/adapters/primary/grpc/proto/__gen__"
)

// LogEntryKind represents the category of a log entry.
type LogEntryKind uint8

const (
	// KindNormal represents a standard log entry containing
	// client commands that should be replicated across the cluster.
	KindNormal LogEntryKind = iota
)

// AsUint8 returns the LogEntryKind as a uint8.
func (l LogEntryKind) AsUint8() uint8 {
	return uint8(l)
}

// Encode converts a domain LogEntryKind into its protobuf equivalent.
func (l LogEntryKind) Encode() graftpb.EntryKind {
	return graftpb.EntryKind(l)
}

// Decode populates a LogEntryKind from its protobuf equivalent.
func (l *LogEntryKind) Decode(pb graftpb.EntryKind) {
	*l = LogEntryKind(pb)
}

// LogEntry represents a single entry in the Raft log.
type LogEntry struct {
	// Term is the term number when the entry was created by the leader.
	Term Term

	// Index is the position of the entry in the log.
	Index int64

	// Kind specifies the type of log entry.
	Kind LogEntryKind

	// Payload holds the actual command to be applied to the state machine.
	Payload []byte
}

// NewLogEntry creates and returns a new log entry.
func NewLogEntry(term Term, kind LogEntryKind, index int64, payload []byte) *LogEntry {
	return &LogEntry{
		Term:    term,
		Kind:    kind,
		Index:   index,
		Payload: payload,
	}
}

// Encode converts a domain LogEntry into its protobuf representation.
func (le *LogEntry) Encode() *graftpb.LogEntry {
	return &graftpb.LogEntry{
		Index:   le.Index,
		Payload: le.Payload,
		Term:    le.Term.AsUint64(),
		Kind:    le.Kind.Encode(),
	}
}

// Decode populates a domain LogEntry from its protobuf representation.
func (le *LogEntry) Decode(pb *graftpb.LogEntry) {
	le.Index = pb.Index
	le.Payload = pb.Payload
	le.Term = Term(pb.Term)
	le.Kind = LogEntryKind(pb.Kind)
}
