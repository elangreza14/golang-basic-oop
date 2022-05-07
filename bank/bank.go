package bank

type (
	IBank interface {
		FlowMoney(reqAmount float64) error
		GetBallance() float64
	}

	Bank struct {
		method IMethod
	}

	IMethod interface {
		AddMoney(reqAmount float64)
		DecMoney(reqAmount float64) error
		GetAmount() float64
	}
)

func NewBank(method IMethod) IBank {
	return &Bank{
		method: method,
	}
}

func (b *Bank) FlowMoney(reqAmount float64) error {
	if reqAmount < 0 {
		if err := b.method.DecMoney(-reqAmount); err != nil {
			return err
		}
		return nil
	}

	b.method.AddMoney(reqAmount)

	return nil
}

func (b *Bank) GetBallance() float64 {
	return b.method.GetAmount()
}
