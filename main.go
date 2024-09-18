package main

import (
	"advancedProject/internal/config"
	"advancedProject/internal/dispatcher"
	"advancedProject/internal/task"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig() // Load configuration (worker count, paths, etc.)
	disp := dispatcher.NewDispatcher(cfg.MaxWorkers)
	disp.Run(cfg.SearchWord)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)

	go func() {
		// Graceful shutdown on signal
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		cancel()
	}()

	// Submit tasks for each file
	go func() {
		for _, file := range cfg.InputFiles {
			t := &task.Task{FilePath: file}
			disp.AddTask(t)
		}
	}()

	// Collect and save results
	go disp.CollectResultsToFile(done, cfg.OutputFile)

	// Wait for graceful shutdown
	select {
	case <-ctx.Done():
		log.Println("Shutting down...")
		disp.Stop()
	}

	// Wait for all tasks to finish
	<-done
	log.Println("All tasks completed.")
}
