package future

import (
	"context"
	"math"
)

// Scheduler represents an executor for a pool of go-routines and is akin to
// thread-pools in other languages. Scheduler must be started before it can
// run jobs.
//  func main() {
//  	poolSize = 5
//  	workBufferSize = 100
//
//  	// Pick an implemenation of Scheduler
//  	var sc Scheduler = NewFixedRoutineScheduler(poolSize, workBufferSize)
//  	if err := sc.Start(); err != nil {
//  		os.Exit(1)
//  	}
//  	// make sure to shut down the pool when you're done using it
//  	defer sc.Stop()
//  }
type Scheduler interface {
	Run(f func()) error
	Start() error
	Stop() error
}

// FixedRoutineScheduler is a Scheduler with a fixed number of go-routines.
type FixedRoutineScheduler struct {
	workerPool     []worker
	workerPoolSize int
	workBuffer     chan func()
	workBufferSize int
}

// NewFixedRoutineScheduler creates a scheduler with a fixed number of
// go-routines in it's pool. Buffer size represents the size of pending
// jobs the scheduler can hold.
// Implements the Scheduler interface
func NewFixedRoutineScheduler(numRoutines, bufferSize int) *FixedRoutineScheduler {
	scheduler := &FixedRoutineScheduler{
		workerPool:     make([]worker, numRoutines),
		workerPoolSize: numRoutines,
		workBuffer:     make(chan func(), bufferSize),
		workBufferSize: bufferSize,
	}

	return scheduler
}

// Start creates the scheduler's monitoring routing as well as all of the
// worker routines.
func (s *FixedRoutineScheduler) Start() {
	// Start monitoring routine
	go func() {
		// 	for {
		// 		select {
		// 		case <-ctx.ctx.Done():
		// 			break
		// 		case _ = <-ctx.fs:
		// 			// TODO: schedule work onto workers
		// 			break
		// 		}
		// 	}
	}()

	// start worker routines
	for i := 0; i < s.workerPoolSize; i++ {
		workQueueSize := int(math.Ceil(float64(s.workBufferSize) / float64(s.workerPoolSize)))
		s.workerPool[i] = newWorker(workQueueSize)
	}
}

type worker struct {
	// Work is communicated with a ring buffer
	workQueue      []func()
	workQueueSize  int
	workWriteIndex int
	workReadIndex  int

	// Work is started and stopped by the work in the buffer. But need to
	// send signals to/from the schduler to signal if work has stopped or
	// should start.
	startChan chan workStart
	stopChan  chan workStop

	// signal to shutdown worker
	ctx context.Context
}

type workStop int
type workStart int

func newWorker(workQueueSize int) worker {
	worker := worker{
		workQueue:      make([]func(), workQueueSize),
		workQueueSize:  workQueueSize,
		workWriteIndex: 0,
		workReadIndex:  0,
		startChan:      make(chan workStart, 1),
		stopChan:       make(chan workStop, 1),
		ctx:            context.Background(),
	}

	go worker.startRoutine()

	return worker
}

// TOOD: implement
func (w *worker) startRoutine() {
	for {
		for w.hasWork() {

		}
	}
}

// TODO: implement
// TODO: write with ringbuffer type
func (w *worker) hasWork() bool {
	writeIndex := w.workWriteIndex
	if writeIndex < w.workReadIndex {
		writeIndex = writeIndex + w.workQueueSize
	}

	return true
}
