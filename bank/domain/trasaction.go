package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ITransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(CreditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	newTransaction := &Transaction{
		ID:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
	}

	return newTransaction
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {

	if creditCard.Balance+t.Amount >= creditCard.Limit {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		creditCard.Balance += t.Amount
	}
}
