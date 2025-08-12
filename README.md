# graft

## About This Project

This project is a Go implementation of Raft, built for learning and
experimentation. The focus is on clarity rather than performance or
production-readiness.

## What is Raft?

Raft is a consensus algorithm used to manage a replicated log across multiple
computers. It ensures that all nodes in a cluster agree on the same sequence of
operations, even if some nodes fail or messages are delayed.

Raft works by electing a leader, who handles client requests and replicates
changes to the other nodes (followers). If the leader fails, the followers hold
an election to choose a new leader. This process makes Raft easier to understand
compared to algorithms like Paxos while providing the same guarantees of safety
and fault tolerance.
