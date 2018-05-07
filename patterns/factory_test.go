package study_go_pattern

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Errorf("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Errorf("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Errorf("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:", err)
}
