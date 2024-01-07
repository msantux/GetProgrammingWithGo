package main

import "sync"

// Visited tracks wheter web pages have been visited.
// Its methods may be used concurrently from multiple goroutines.
type Visited struct {
	// mu guards the visited map.
	mu      sync.Mutex
	visited map[string]int
}

// VisitLink tracks that the page with the given URl has
// been visited, and returns the updated link count.
func (v *Visited) VisitedLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()

	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}
