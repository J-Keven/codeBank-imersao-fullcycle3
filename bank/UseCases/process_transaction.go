package usecases

import (
	"github.com/j-keven/codeBank/domain"
	"github.com/j-keven/codeBank/dto"
)

type UseCaseTransaction struct {
	TransactionRepository domain.ITransactionRepository
}

func NewUseCaseTransaction(transactionRepository domain.ITransactionRepository) UseCaseTransaction {
	return UseCaseTransaction{TransactionRepository: transactionRepository}
}

func (u UseCaseTransaction) ProcessTransaction(transactionDTO dto.Transaction) (domain.Transaction, error) {
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
		return domain.Transaction{}, nil
	}

	return *t, nil
}

func (u UseCaseTransaction) hydrateCreditCard(transactionDTO dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()
	creditCard.Name = transactionDTO.Name
	creditCard.Number = transactionDTO.Number
	creditCard.CVV = transactionDTO.CVV
	creditCard.ExpirationMonth = transactionDTO.ExpirationMonth
	creditCard.ExpirationYear = transactionDTO.ExpirationYear

	return creditCard
}

func (u UseCaseTransaction) newTransaction(transactionDTO dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.Amount = transactionDTO.Amount
	t.CreditCardId = cc.ID
	t.Status = transactionDTO.Status
	t.Store = transactionDTO.Store
	t.Description = transactionDTO.Description

	return t
}
