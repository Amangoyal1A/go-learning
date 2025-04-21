package worker

import (
	"fmt"
	"time"
	"worker-app/config"
	"worker-app/models"
)

func StartWorkerPool(workerCount int) {
	jobs := make(chan models.Job, 10)

	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs)
	}

	go func() {
		for {
			var pendingJobs []models.Job
			config.DB.Where("status = ?", "pending").Find(&pendingJobs)
			for _, job := range pendingJobs {
				jobs <- job
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func worker(id int, jobs <-chan models.Job) {
	for job := range jobs {
		fmt.Printf("👷 Worker %d picked job %d: %s\n", id, job.ID, job.Task)
		// time.Sleep(1 * time.Second) // simulate job processing

		config.DB.Model(&job).Update("status", "done")
		fmt.Printf("✅ Worker %d completed job %d\n", id, job.ID)
	}
}
