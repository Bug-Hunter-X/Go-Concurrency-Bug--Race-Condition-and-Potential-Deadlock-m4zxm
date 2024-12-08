# Go Concurrency Bug: Race Condition and Potential Deadlock

This repository demonstrates a common concurrency issue in Go involving race conditions and potential deadlocks when using goroutines and channels. The `bug.go` file contains the buggy code, while `bugSolution.go` provides a corrected version.

## Bug Description

The code creates ten goroutines that incrementally send integers to a channel. A separate goroutine waits for all the sending goroutines to complete before closing the channel. While seemingly straightforward, it harbors a race condition and a potential deadlock if not properly handled.

## Solution

The solution in `bugSolution.go` addresses the concurrency issues by ensuring proper synchronization using mutexes and wait groups.