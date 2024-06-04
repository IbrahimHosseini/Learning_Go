package closures

// The BankService type in Go contains a transfer function that takes two float64 parameters and
// returns a float64.
// @property transfer - The `BankService` struct has a single property `transfer` which is a function
// that takes two `float64` arguments and returns a `float64`.
type BankService struct {
	transfer func(float64, float64) float64
}

// The NewBankService function creates a new instance of BankService with a transfer method that
// subtracts the transfer amount from the sender's balance.
func NewBanckService() *BankService {
	service := &BankService{}

	service.transfer = func(from, to float64) float64 {
		return from - to
	}

	return service
}

// The `Transfer` method in the `BankService` struct is calling the `transfer` function that was
// defined when creating a new instance of `BankService` using the `NewBankService` function. This
// `transfer` function subtracts the transfer amount (`to`) from the sender's balance (`from`) and
// returns the result. So, the `Transfer` method essentially facilitates transferring funds from one
// account to another by subtracting the transfer amount from the sender's balance.
func (service *BankService) Transfer(from, to float64) float64 {
	return service.transfer(from, to)
}
