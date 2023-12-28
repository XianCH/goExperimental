package main

import (
	"fmt"
	"log"
	"sync"
)

type Job interface {
	Run()
}

type JobChan chan Job

type Worker struct {
	ID         int
	JobQueue   JobChan
	WorkerPool chan *Worker
}

func NewWorker(id int, workerpool chan *Worker) *Worker {
	return &Worker{
		ID:         id,
		JobQueue:   make(JobChan),
		WorkerPool: workerpool,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w
			select {
			case job := <-w.JobQueue:
				log.Printf("Worker %d processing job\n", w.ID)
				job.Run()
			}
		}
	}()
}

type WorkerPool struct {
	Size       int
	JobQueue   JobChan
	WorkerPool chan *Worker
	WaitGroup  *sync.WaitGroup
	Mutex      sync.Mutex
}

func NewWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		Size:       size,
		JobQueue:   make(JobChan),
		WorkerPool: make(chan *Worker, size),
		WaitGroup:  &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) StartPool() {
	for i := 0; i < wp.Size; i++ {
		worker := NewWorker(i, wp.WorkerPool)
		go worker.Start()
		wp.WaitGroup.Add(1)
	}

	go func() {
		for job := range wp.JobQueue {
			worker := <-wp.WorkerPool
			go func(j Job, w *Worker) {
				defer func() {
					wp.WorkerPool <- w
					wp.WaitGroup.Done() // Decrement the wait group when the job is completed
				}()
				w.JobQueue <- j
			}(job, worker)
		}
	}()
}

func (wp *WorkerPool) SubmitJob(job Job) {
	wp.WaitGroup.Add(1) // Increment the wait group when submitting a job
	wp.JobQueue <- job
}

func (wp *WorkerPool) Close() {
	close(wp.JobQueue)
	wp.WaitGroup.Wait()
	close(wp.WorkerPool)

	for worker := range wp.WorkerPool {
		close(worker.JobQueue)
	}
}

type TickerSys struct {
	tickerCount int
	Mutex       *sync.Mutex
}

func NewTickerSys(initialCount int) *TickerSys {
	return &TickerSys{
		tickerCount: initialCount,
		Mutex:       &sync.Mutex{},
	}
}

func (t *TickerSys) Run(request int) {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	t.tickerCount -= request
}

type PurchaseJob struct {
	TickerSys   *TickerSys
	TicketCount int
}

func (p *PurchaseJob) Run() {
	p.TickerSys.Run(p.TicketCount)
}

func main() {
	wp := NewWorkerPool(10)
	wp.StartPool()

	tickerSys := NewTickerSys(100)

	for i := 0; i < 1000; i++ {
		job := &PurchaseJob{
			TickerSys:   tickerSys,
			TicketCount: 1,
		}
		wp.SubmitJob(job)
		if tickerSys.tickerCount == 0 {
			fmt.Println("售票结束")
			break
		}
	}
	wp.Close()
}
