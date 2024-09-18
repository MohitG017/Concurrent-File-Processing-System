package dispatcher

import (
	"advancedProject/internal/task"
	"fmt"
	"log"
)

type Worker struct {
	ID          int
	TaskQueue   chan *task.Task
	ResultQueue chan string
	Quit        chan bool
}

// Start runs the worker in a Goroutine
func (w *Worker) Start(searchWord string) {
	go func() {
		for {
			select {
			case t := <-w.TaskQueue:
				count, err := t.ProcessFile(searchWord)
				if err != nil {
					log.Printf("Worker %d: Error processing file %s: %v", w.ID, t.FilePath, err)
					continue
				}
				// Properly format the count using fmt.Sprintf
				w.ResultQueue <- fmt.Sprintf("%s: %s occurs %d times", t.FilePath, searchWord, count)
			case <-w.Quit:
				log.Printf("Worker %d is stopping", w.ID)
				return
			}
		}
	}()
}

// Stop signals the worker to stop
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}
