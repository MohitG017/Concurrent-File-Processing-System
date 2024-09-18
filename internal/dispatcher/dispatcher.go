package dispatcher

import (
	"advancedProject/internal/task"
	"log"
	"os"
)

type Dispatcher struct {
	TaskQueue   chan *task.Task
	ResultQueue chan string
	WorkerPool  []*Worker
	MaxWorkers  int
}

// NewDispatcher creates a new Dispatcher
func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{
		TaskQueue:   make(chan *task.Task),
		ResultQueue: make(chan string),
		WorkerPool:  make([]*Worker, maxWorkers),
		MaxWorkers:  maxWorkers,
	}
}

// Run starts the worker pool
func (d *Dispatcher) Run(searchWord string) {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := &Worker{
			ID:          i + 1,
			TaskQueue:   d.TaskQueue,
			ResultQueue: d.ResultQueue,
		}
		worker.Start(searchWord)
		d.WorkerPool[i] = worker
	}
}

// AddTask adds a task to the task queue
func (d *Dispatcher) AddTask(t *task.Task) {
	d.TaskQueue <- t
}

// CollectResultsToFile collects results from workers and saves them to an output file
func (d *Dispatcher) CollectResultsToFile(done chan bool, outputFile string) {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	for result := range d.ResultQueue {
		_, err := file.WriteString(result + "\n")
		if err != nil {
			log.Printf("Failed to write to file: %v", err)
		}
	}
	done <- true
}

// Stop stops all workers
func (d *Dispatcher) Stop() {
	for _, worker := range d.WorkerPool {
		worker.Stop()
	}
}
