package main

// "time"

// Task struct - jo task worker ko dena hai
type Task struct {
	id   int
	name string
}

// Worker function - jo har task ko process karega
// func worker(id int, jobs <-chan Task, wg *sync.WaitGroup) {
// 	defer wg.Done() // Task complete hone par WaitGroup se Done call karenge
// 	for job := range jobs {
// 		fmt.Printf("Worker %d starting task %d: %s\n", id, job.id, job.name)
// 		// Simulating task processing (like image processing)
// 		// time.Sleep(2 * time.Second)
// 		fmt.Printf("Worker %d completed task %d: %s\n", id, job.id, job.name)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	jobs := make(chan Task, 10) // Channel to hold the tasks

// 	// Start workers (goroutines)
// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1) // Add workers to the WaitGroup
// 		go worker(i, jobs, &wg)
// 	}

// 	// Add tasks to the job channel
// 	tasks := []Task{
// 		{1, "Image1.jpg"},
// 		{2, "Image2.jpg"},
// 		{3, "Image3.jpg"},
// 		{4, "Image4.jpg"},
// 		{5, "Image5.jpg"},
// 		{6, "Image6.jpg"},
// 		{7, "Image7.jpg"},
// 		{8, "Image8.jpg"},
// 		{9, "Image9.jpg"},
// 		{10, "Image10.jpg"},
// 		{11, "Image11.jpg"},
// 		{12, "Image12.jpg"},
// 	}

// 	for _, task := range tasks {
// 		jobs <- task // Send task to the job channel
// 	}

// 	close(jobs) // Close channel once all tasks are sent
// 	wg.Wait()   // Wait until all workers are done

// 	fmt.Println("All tasks completed.")
// }
