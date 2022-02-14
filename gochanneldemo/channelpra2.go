package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	res := make(chan int, 100)
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 3; j++ {
		wg.Add(1)
		go worker(j, jobs, res)
	}
	wg.Wait()
	close(res)
	for val := range res {
		fmt.Println(val)
	}
}

func worker(job int, jobs, res chan int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("%d start job %d\n", job, j)
		res <- j * 2
	}
}
