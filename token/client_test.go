package token

import (
	"testing"

	stripe "github.com/channelmeter/stripe-go"
	"github.com/channelmeter/stripe-go/bankaccount"
	. "github.com/channelmeter/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestTokenNew(t *testing.T) {
	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
	}

	target, err := New(tokenParams)

	if err != nil {
		t.Error(err)
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Type != Card {
		t.Errorf("Type %v does not match expected value\n", target.Type)
	}

	if target.Card == nil {
		t.Errorf("Card is not set\n")
	}

	if target.Card.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.Card.LastFour, tokenParams.Card.Number)
	}
}

func TestTokenGet(t *testing.T) {
	tokenParams := &stripe.TokenParams{
		Bank: &stripe.BankAccountParams{
			Country: "US",
			Routing: "110000000",
			Account: "000123456789",
		},
	}

	tok, _ := New(tokenParams)

	target, err := Get(tok.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if target.Type != Bank {
		t.Errorf("Type %v does not match expected value\n", target.Type)
	}

	if target.Bank == nil {
		t.Errorf("Bank account is not set\n")
	}

	if target.Bank.Status != bankaccount.NewAccount {
		t.Errorf("Bank account status %q does not match expected value\n", target.Bank.Status)
	}
}
