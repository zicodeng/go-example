package main

import (
	"sync"
)

// Stack represents a stack data structure
// that is safe for concurrent use.
type Stack struct {
	// Lower-case names are used here
	// because we want to make them unexported
	// so that code in other packages cannot access them directly.
	entries []string   // Slice for the entries.
	mx      sync.Mutex // Mutex to protect the slice.
}

// Push pushes a new entry on to the stack.
func (s *Stack) Push(entry string) {
	// Obtain the exclusive lock.
	s.mx.Lock()
	// Append the new entry to the slice.
	s.entries = append(s.entries, entry)
	// Release the exclusive lock.
	s.mx.Unlock()
}

// Pop pops the last entry off of the stack.
func (s *Stack) Pop() string {
	// Obtain the exclusive lock.
	s.mx.Lock()
	// Use defer to ensure that we release the lock,
	// Regardless of how we exit this function.
	defer s.mx.Unlock()
	// If there are no entries, just return "".
	if len(s.entries) == 0 {
		return ""
	}
	// Get the last entry in the slice.
	e := s.entries[len(s.entries)-1]
	// Remove that entry from the slice.
	s.entries = s.entries[:len(s.entries)-1]
	// Return the entry.
	return e
}
