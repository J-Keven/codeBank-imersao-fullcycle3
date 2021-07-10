package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	usecases "github.com/j-keven/codeBank/UseCases"
	"github.com/j-keven/codeBank/infra/grpc/server"
	"github.com/j-keven/codeBank/infra/kafka"
	"github.com/j-keven/codeBank/infra/repositories"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()

	kafkaProducer := setupKafkaProducer()

	processTransactionUseCase := setupTransactionUseCase(db, kafkaProducer)

	fmt.Printf("runing grpc server")
	serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB, kafkaProducer kafka.KafkaProducer) usecases.UseCaseTransaction {
	tranactionRepository := repositories.NewTransactionRepositoryDB(db)
	useCaseTransaction := usecases.NewUseCaseTransaction(tranactionRepository)
	useCaseTransaction.SetupProducer(kafkaProducer)

	return useCaseTransaction
}

func setupKafkaProducer() kafka.KafkaProducer {
	kafkaProducer := kafka.NewProducerKafka()

	kafkaProducer.SetupProducer("host.docker.internal:9094")

	return kafkaProducer
}

func setupDb() *sql.DB {
	pgCredentials := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DBNAME"),
	)

	fmt.Println(pgCredentials)
	db, err := sql.Open("postgres", pgCredentials)

	if err != nil {
		log.Fatal("error connection to database", err)
	}
	return db
}

func serveGrpc(processTransactionUseCase usecases.UseCaseTransaction) {
	server := server.NewGRPCServer()
	server.ProcessTransactionUseCase = processTransactionUseCase
	server.Serve()
}
