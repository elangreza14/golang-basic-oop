package conventional

import (
	"errors"

	"github.com/elangreza14/golang-basic-oop/bank"
)

type (
	conventional struct {
		Amount      float64
		Name        string
		withdrawFee float64
	}

	IConventional interface {
		bank.IMethod
	}
)

func NewConventional(name string,
	withdrawFee float64) IConventional {
	return &conventional{
		Amount:      0,
		Name:        name,
		withdrawFee: withdrawFee,
	}
}

func (c *conventional) AddMoney(reqAmount float64) {
	c.Amount = c.Amount + reqAmount
}

func (c *conventional) DecMoney(reqAmount float64) error {
	baseReq := reqAmount + reqAmount*c.getWithdrawFee()

	if reqAmount > baseReq {
		return errors.New("reqAmount cannot exceed current amount")
	}

	c.Amount = baseReq - reqAmount

	return nil
}

func (c *conventional) GetAmount() float64 {
	return c.Amount
}

func (c *conventional) getWithdrawFee() float64 {
	return c.withdrawFee
}
