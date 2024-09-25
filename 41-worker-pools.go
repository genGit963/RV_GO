// package main

// import (
// 	"fmt"
// 	"time"
// )

// /*
// Here’s the worker, of which we’ll run several concurrent
// instances. These workers will receive work on the jobs channel
//  and send the corresponding results on results.
//  We’ll sleep a second per job to simulate an expensive task.
// */

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	smthCalculation := 0
// 	for j := range jobs {
// 		fmt.Println("worker ", id, "started job: ", j)
// 		time.Sleep(time.Second)
// 		smthCalculation = j * 10
// 		fmt.Println("worker ", id, "finished job: ", j)
// 		results <- smthCalculation
// 	}
// }

// func learn_Worker_pools() {
// 	fmt.Println("\n----------- learn_Worker_pools -------------")

// 	const numWorkers = 3
// 	const numJobs = 5
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// setting workers
// 	for w := 1; w <= numWorkers; w++ {
// 		go worker(w, jobs, results)
// 	}

// 	// setting jobs
// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}

// 	close(jobs) //close jobs channel

// 	// fetch the results of each worker's jobs
// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}
// }

// func main() {
// 	learn_Worker_pools()
// 	println("\n-------------------------------")
// }

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
