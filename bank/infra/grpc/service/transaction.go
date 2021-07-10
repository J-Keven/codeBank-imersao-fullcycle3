package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	usecases "github.com/j-keven/codeBank/UseCases"
	"github.com/j-keven/codeBank/dto"
	"github.com/j-keven/codeBank/infra/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecases.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            in.GetCreditCardd().GetName(),
		Number:          in.GetCreditCardd().GetNumber(),
		ExpirationMonth: in.GetCreditCardd().GetExpirationMonth(),
		ExpirationYear:  in.GetCreditCardd().GetExpirationYear(),
		CVV:             in.GetCreditCardd().GetCvv(),
		Store:           in.GetStore(),
		Amount:          in.GetAmount(),
		Description:     in.GetDescritpion(),
	}
	// example
	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDto)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	return &empty.Empty{}, nil
}
