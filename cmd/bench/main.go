package main

import (
	"log"

	"github.com/mykysha/redisBenchmark/pkg/benchmark"
	"github.com/mykysha/redisBenchmark/pkg/clients/postgres"
	"github.com/mykysha/redisBenchmark/pkg/clients/redis"
	"github.com/mykysha/redisBenchmark/pkg/config"
	"github.com/mykysha/redisBenchmark/pkg/workerpool"
)

const configFile = "config.yaml"

func main() {
	configReader, err := config.NewReaderWithPath(configFile)
	if err != nil {
		log.Fatalf("failed to create config reader: %v", err)
	}

	workerNumber := configReader.GetInt("pool.workerCount")
	workerPool := workerpool.NewPool(workerNumber)

	postgresHost := configReader.GetString("postgres.host")
	postgresPort := configReader.GetString("postgres.port")
	postgresUser := configReader.GetString("postgres.user")
	postgresPassword := configReader.GetString("postgres.password")
	postgresDatabase := configReader.GetString("postgres.database")
	postgresSSLMode := configReader.GetString("postgres.sslMode")

	postgresClient, err := postgres.NewClient(postgresHost, postgresPort, postgresUser, postgresPassword, postgresDatabase, postgresSSLMode)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}

	redisAddress := configReader.GetString("redis.address")

	redisClient := redis.NewClient(redisAddress)

	numberOfRuns := configReader.GetInt("benchmark.runs")
	numberOfTests := configReader.GetInt("benchmark.tests")

	bench := benchmark.NewBenchmark(numberOfRuns, numberOfTests, postgresClient, redisClient, workerPool)

	err = bench.Run()
	if err != nil {
		log.Fatalf("failed to run benchmark: %v", err)
	}
}
