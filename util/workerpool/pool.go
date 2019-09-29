package workerpool

type Job struct {
	F   func()
	Idx int
}

type Worker struct {
	JobChannel chan Job
	Stop       chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		JobChannel: make(chan Job),
		Stop:       make(chan struct{}),
	}
}

func (w *Worker) Start() {
	go func() {
		var job Job
		for {
			select {
			case job = <-w.JobChannel:
				job.F()
			case <-w.Stop:
				w.Stop <- struct{}{}
				return
			}
		}
	}()
}

type Pool struct {
	JobQueue    chan Job
	WorkerQueue map[int]*Worker
	NumWorkers  int
	Stop        chan struct{}
}

func NewPool(numWorkers int, jobQueueLen int) *Pool {
	if numWorkers == 0 {
		numWorkers = 1
	}

	jobQueue := make(chan Job, jobQueueLen)
	workerQueue := make(map[int]*Worker, numWorkers)

	pool := &Pool{
		JobQueue:    jobQueue,
		WorkerQueue: workerQueue,
		NumWorkers:  numWorkers,
		Stop:        make(chan struct{}),
	}
	pool.Start()
	return pool
}

func (p *Pool) Start() {
	for i := 0; i < p.NumWorkers; i++ {
		worker := NewWorker()
		p.WorkerQueue[i] = worker
		worker.Start()
	}

	go p.Dispatch()
}

func (p *Pool) Dispatch() {
	for {
		select {
		case job := <-p.JobQueue:
			worker := p.FindWorker(job.Idx)
			worker.JobChannel <- job
		case <-p.Stop:
			for i := 0; i < p.NumWorkers; i++ {
				worker := p.FindWorker(i)
				worker.Stop <- struct{}{}
				<-worker.Stop
			}

			p.Stop <- struct{}{}
			return
		}
	}
}

func (p *Pool) FindWorker(idx int) *Worker {
	key := idx % p.NumWorkers
	if worker, ok := p.WorkerQueue[key]; ok {
		return worker
	} else {
		panic("Worker escaped from WorkerQueue")
	}
}

func (p *Pool) StopAll() {
	p.Stop <- struct{}{}
	<-p.Stop
}
