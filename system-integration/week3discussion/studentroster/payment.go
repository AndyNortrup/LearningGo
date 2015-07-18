package roster

//Payment represents the status of a student's payment to the University
type Payment struct {
	PaymentID     int
	UnpaidAmmount float32
	Status        PaymentStatus
}

//PaymentStatus is a status type for when a student has paid their bills
type PaymentStatus int

const (
	//PaidInFull indicates that a student has paid their bill
	PaidInFull PaymentStatus = iota

	//PaymentDue indicates that a student has a payment due
	PaymentDue PaymentStatus = iota

	//Late indicates that a student is overduue for a payment
	Late PaymentStatus = iota
)

//MakePayment reduces the UnpaidAmmount of the Payment and returns
// the new Unpaid Amount.  If the UnpaidAmmount is reduced below zero
// then the PaymentStatus is changed to PaidInFull
func (p Payment) MakePayment(ammount float32) float32 {
	p.UnpaidAmmount = p.UnpaidAmmount - ammount
	if p.UnpaidAmmount < 0 {
		p.Status = PaidInFull
	}
	return p.UnpaidAmmount
}
