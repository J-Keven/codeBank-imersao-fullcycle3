package usecases

import (
	"encoding/json"

	"github.com/j-keven/codeBank/domain"
	"github.com/j-keven/codeBank/dto"
	"github.com/j-keven/codeBank/infra/kafka"
)

type UseCaseTransaction struct {
	TransactionRepository domain.ITransactionRepository
	KafkaProducer         kafka.KafkaProducer
}

func NewUseCaseTransaction(transactionRepository domain.ITransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u *UseCaseTransaction) ProcessTransaction(transactionDTO dto.Transaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(transactionDTO)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Balance = ccBalanceAndLimit.Balance
	creditCard.Limit = ccBalanceAndLimit.Limit

	t := u.newTransaction(transactionDTO, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.SaveTransaction(*t, *creditCard)

	if err != nil {
		return domain.Transaction{}, err
	}

	transactionDTO.ID = t.ID
	transactionDTO.CreatedAt = t.CreatedAt

	transactionJson, err := json.Marshal(transactionDTO)

	if err != nil {
		return domain.Transaction{}, err
	}

	err = u.KafkaProducer.Publish(string(transactionJson), "Payments")

	if err != nil {
		return domain.Transaction{}, err
	}
	return *t, nil
}

func (u *UseCaseTransaction) hydrateCreditCard(transactionDTO dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDTO.Name
	creditCard.Number = transactionDTO.Number
	creditCard.CVV = transactionDTO.CVV
	creditCard.ExpirationMonth = transactionDTO.ExpirationMonth
	creditCard.ExpirationYear = transactionDTO.ExpirationYear

	return creditCard
}

func (u *UseCaseTransaction) newTransaction(transactionDTO dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.Amount = transactionDTO.Amount
	t.CreditCardId = cc.ID
	t.Store = transactionDTO.Store
	t.Description = transactionDTO.Description

	return t
}

func (u *UseCaseTransaction) SetupProducer(kafkaProducer kafka.KafkaProducer) {
	u.KafkaProducer = kafkaProducer
}
