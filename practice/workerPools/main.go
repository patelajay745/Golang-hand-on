package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	ImageId    int
	ImageBytes []byte
}

func processImage(workerId int, job Job, results chan<- int) {

	fmt.Printf("Worker %d is processing image %d...\n", workerId, job.ImageId)
	time.Sleep(time.Second * 15)
	fmt.Println("Image %d processing completed \n", workerId, job.ImageId)

	results <- job.ImageId
}

func imageHandler(w http.ResponseWriter, r *http.Request, jobs chan<- Job) {

	imageId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid image Id", http.StatusBadRequest)
		return
	}

	imageByte := make([]byte, 1024)

	job := Job{
		ImageId:    imageId,
		ImageBytes: imageByte,
	}

	jobs <- job

	fmt.Fprintf(w, "Image %d processing started ... \n", imageId)
}

func workerPool(numWorkers int, jobs <-chan Job, results chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= numWorkers; i++ {
		go func(workerId int) {
			for job := range jobs {
				processImage(workerId, job, results)
			}
		}(i)
	}

}

func main() {

	const numWorkers = 3
	const jobQueueSize = 10

	jobs := make(chan Job, jobQueueSize)
	results := make(chan int, jobQueueSize)

	var wg sync.WaitGroup

	wg.Add(numWorkers)

	go workerPool(numWorkers, jobs, results, &wg)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		imageHandler(w, r, jobs)
	})

	// Listen for incoming HTTP requests
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)

	// Wait for all workers to finish
	wg.Wait()

	// Close channels after all workers finish
	close(jobs)
	close(results)

}
