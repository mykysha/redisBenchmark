package workerpool

type Pool interface {
	// Add adds a new task to the WorkerPool.
	Add(task func()) error
	// Run runs all tasks in the WorkerPool.
	Run()
	// Wait waits for all tasks to finish.
	Wait()
	// Stop stops the WorkerPool.
	Stop()
}

// NewPool creates a new WorkerPool.
func NewPool(numberOfWorkers int) *WorkerPool {
	return &WorkerPool{
		numberOfWorkers: numberOfWorkers,
		tasks:           make(chan func()),
		stop:            make(chan struct{}),
	}
}

type WorkerPool struct {
	numberOfWorkers int
	tasks           chan func()
	stop            chan struct{}
}

func (p *WorkerPool) Add(task func()) error {
	select {
	case <-p.stop:
		return ErrPoolStopped
	case p.tasks <- task:
		return nil
	}
}

func (p *WorkerPool) Run() {
	for i := 0; i < p.numberOfWorkers; i++ {
		go func() {
			for {
				select {
				case <-p.stop:
					return
				case task := <-p.tasks:
					task()
				}
			}
		}()
	}
}

func (p *WorkerPool) Wait() {
	for {
		if len(p.tasks) == 0 {
			return
		}
	}
}

func (p *WorkerPool) Stop() {
	close(p.stop)
}
