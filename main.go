package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d got job: %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <num_workers>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	jobs := make(chan int)

	// Запускаем N воркеров
	for i := 1; i <= n; i++ {
		go worker(i, jobs)
	}

	// Главная горутина пишет данные в канал бесконечно
	counter := 1
	for {
		jobs <- counter
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}
