package closures

type BackAccount struct {
	Balance float64
}

func NewAccount(initialBalance float64) func(float64) float64 {
	account := BackAccount{Balance: initialBalance}

	return func(amount float64) float64 {
		account.Balance += amount
		return account.Balance
	}
}
