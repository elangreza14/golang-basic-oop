package internet

import (
	"errors"

	"github.com/elangreza14/golang-basic-oop/bank"
)

type (
	internet struct {
		Amount float64
		Name   string
	}
	IInternet interface {
		bank.IMethod
	}
)

func NewInternet(name string) IInternet {
	return &internet{
		Amount: 0,
		Name:   name,
	}
}

func (i *internet) AddMoney(reqAmount float64) {
	i.Amount = i.Amount + reqAmount
}

func (i *internet) DecMoney(reqAmount float64) error {
	if reqAmount > i.Amount {
		return errors.New("reqAmount cannot exceed current amount")
	}

	i.Amount = i.Amount - reqAmount

	return nil
}

func (i *internet) GetAmount() float64 {
	return i.Amount
}
