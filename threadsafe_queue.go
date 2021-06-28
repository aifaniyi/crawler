package main

import (
	"sync"

	"github.com/aifaniyi/algorithms/datastructures/queue"
)

// SafeQueue : using this implementation (i.e dynamically resized queue)
// instead of a regular go channel (which would have been an ideal substitute)
// because it frees us from having to preallocate n siezd buffer chan
// given we do not know the maximum number of links that can be present in
// a fetched page
type SafeQueue struct {
	queue *queue.Queue
	sync.Mutex
}

func NewSafeQueue() *SafeQueue {
	q := queue.NewQueue()
	return &SafeQueue{
		queue: q,
	}
}

func (s *SafeQueue) enqueue(elem interface{}) {
	s.Lock()
	s.queue.Enqueue(elem)
	s.Unlock()
}

func (s *SafeQueue) dequeue() interface{} {
	s.Lock()
	elem := s.queue.Dequeue()
	s.Unlock()
	return elem
}

func (s *SafeQueue) isEmpty() bool {
	s.Lock()
	status := s.queue.IsEmpty()
	s.Unlock()
	return status
}
