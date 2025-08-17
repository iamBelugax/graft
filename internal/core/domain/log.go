package domain

// LogEntryKind represents the category of a log entry.
type LogEntryKind uint8

const (
	// KindNormal represents a standard log entry containing
	// client commands that should be replicated across the cluster.
	KindNormal LogEntryKind = iota
)

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
