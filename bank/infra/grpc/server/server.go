package server

import (
	"log"
	"net"

	usecases "github.com/j-keven/codeBank/UseCases"
	"github.com/j-keven/codeBank/infra/grpc/pb"
	service "github.com/j-keven/codeBank/infra/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecases.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (g *GRPCServer) Serve() {
	list, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("could not list tcp port")
	}

	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = g.ProcessTransactionUseCase
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, transactionService)
	reflection.Register(grpcServer)
	grpcServer.Serve(list)
}
