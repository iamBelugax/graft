package domain

// EntryKind represents the category of a log entry.
type EntryKind uint8

const (
	// KindNormal represents a standard log entry containing
	// client commands that should be replicated across the cluster.
	KindNormal EntryKind = iota
)
