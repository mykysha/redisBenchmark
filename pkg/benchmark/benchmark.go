package benchmark

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mykysha/redisBenchmark/domain"
	"github.com/mykysha/redisBenchmark/pkg/clients"
	"github.com/mykysha/redisBenchmark/pkg/clients/postgres"
	"github.com/mykysha/redisBenchmark/pkg/clients/redis"
	"github.com/mykysha/redisBenchmark/pkg/workerpool"
)

type Benchmark struct {
	numberOfRuns   int
	numberOfTests  int
	postgresClient *postgres.Client
	redisClient    *redis.Client
	workerPool     workerpool.Pool
}

func NewBenchmark(numberOfRuns, numberOfTests int, postgresClient *postgres.Client, redisClient *redis.Client, workerPool workerpool.Pool) *Benchmark {
	workerPool.Run()

	return &Benchmark{
		numberOfRuns:   numberOfRuns,
		numberOfTests:  numberOfTests,
		postgresClient: postgresClient,
		redisClient:    redisClient,
		workerPool:     workerPool,
	}
}

func (b *Benchmark) Run() error {
	postgresTime := b.measureClient(b.postgresClient)
	redisTime := b.measureClient(b.redisClient)

	msg := "Postgres: " + postgresTime.String() + "\nRedis: " + redisTime.String()

	_, err := os.Stdout.WriteString(msg)
	if err != nil {
		return fmt.Errorf("failed to write to stdout: %w", err)
	}

	return nil
}

func (b *Benchmark) measureClient(client clients.Client) time.Duration {
	runTimes := make([]time.Duration, 0, b.numberOfRuns)

	for i := 0; i < b.numberOfRuns; i++ {
		start := time.Now()

		for i := 0; i < b.numberOfTests; i++ {
			run := i + 1

			b.workerPool.Add(func() {
				databaseClientBenchmark(run, client)
			})
		}

		b.workerPool.Wait()

		runTimes = append(runTimes, time.Since(start))
	}

	totalTime := time.Duration(0)

	for _, t := range runTimes {
		totalTime += t
	}

	avgTime := totalTime / time.Duration(b.numberOfRuns)

	return avgTime
}

func databaseClientBenchmark(code int, client clients.Client) {
	author := domain.Author{
		Name:         "Author#" + strconv.Itoa(code),
		Surname:      "Surname#" + strconv.Itoa(code),
		BirthCountry: "BirthCountry#" + strconv.Itoa(code),
	}

	id, err := client.CreateAuthor(author)
	if err != nil {
		panic(err)
	}

	author.ID = id

	book := domain.Book{
		AuthorID: author.ID,
		Year:     code,
		Pages:    code,
		Title:    "Title#" + strconv.Itoa(code),
		Genre:    "Genre#" + strconv.Itoa(code),
	}

	id, err = client.CreateBook(book)
	if err != nil {
		panic(err)
	}

	book.ID = id
	author.Name = "Author#" + strconv.Itoa(code) + "_updated"
	book.Pages = code + 1

	err = client.UpdateAuthor(author)
	if err != nil {
		panic(err)
	}

	err = client.UpdateBook(book)
	if err != nil {
		panic(err)
	}

	_, err = client.GetAuthor(author.ID)
	if err != nil {
		panic(err)
	}

	_, err = client.GetBook(book.ID)
	if err != nil {
		panic(err)
	}

	err = client.DeleteBook(book.ID)
	if err != nil {
		panic(err)
	}

	err = client.DeleteAuthor(author.ID)
	if err != nil {
		panic(err)
	}
}
