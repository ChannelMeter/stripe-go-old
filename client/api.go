// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	. "github.com/channelmeter/stripe-go"
	"github.com/channelmeter/stripe-go/account"
	"github.com/channelmeter/stripe-go/balance"
	"github.com/channelmeter/stripe-go/bankaccount"
	"github.com/channelmeter/stripe-go/bitcoinreceiver"
	"github.com/channelmeter/stripe-go/bitcointransaction"
	"github.com/channelmeter/stripe-go/card"
	"github.com/channelmeter/stripe-go/charge"
	"github.com/channelmeter/stripe-go/coupon"
	"github.com/channelmeter/stripe-go/customer"
	"github.com/channelmeter/stripe-go/discount"
	"github.com/channelmeter/stripe-go/dispute"
	"github.com/channelmeter/stripe-go/event"
	"github.com/channelmeter/stripe-go/fee"
	"github.com/channelmeter/stripe-go/feerefund"
	"github.com/channelmeter/stripe-go/fileupload"
	"github.com/channelmeter/stripe-go/invoice"
	"github.com/channelmeter/stripe-go/invoiceitem"
	"github.com/channelmeter/stripe-go/plan"
	"github.com/channelmeter/stripe-go/recipient"
	"github.com/channelmeter/stripe-go/refund"
	"github.com/channelmeter/stripe-go/reversal"
	"github.com/channelmeter/stripe-go/sub"
	"github.com/channelmeter/stripe-go/token"
	"github.com/channelmeter/stripe-go/transfer"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// Charges is the client used to invoke /charges APIs.
	// For more details see https://stripe.com/docs/api#charges.
	Charges *charge.Client
	// Customers is the client used to invoke /customers APIs.
	// For more details see https://stripe.com/docs/api#customers.
	Customers *customer.Client
	// Cards is the client used to invoke /cards APIs.
	// For more details see https://stripe.com/docs/api#cards.
	Cards *card.Client
	// Subs is the client used to invoke /subscriptions APIs.
	// For more details see https://stripe.com/docs/api#subscriptions.
	Subs *sub.Client
	// Plans is the client used to invoke /plans APIs.
	// For more details see https://stripe.com/docs/api#plans.
	Plans *plan.Client
	// Coupons is the client used to invoke /coupons APIs.
	// For more details see https://stripe.com/docs/api#coupons.
	Coupons *coupon.Client
	// Discounts is the client used to invoke discount-related APIs.
	// For mode details see https://stripe.com/docs/api#discounts.
	Discounts *discount.Client
	// Invoices is the client used to invoke /invoices APIs.
	// For more details see https://stripe.com/docs/api#invoices.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	// For more details see https://stripe.com/docs/api#invoiceitems.
	InvoiceItems *invoiceitem.Client
	// Disputes is the client used to invoke dispute-related APIs.
	// For more details see https://stripe.com/docs/api#disputes.
	Disputes *dispute.Client
	// Transfers is the client used to invoke /transfers APIs.
	// For more details see https://stripe.com/docs/api#transfers.
	Transfers *transfer.Client
	// Recipients is the client used to invoke /recipients APIs.
	// For more details see https://stripe.com/docs/api#recipients.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	// For more details see https://stripe.com/docs/api#refunds.
	Refunds *refund.Client
	// Fees is the client used to invoke /application_fees APIs.
	// For more details see https://stripe.com/docs/api#application_fees.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	// For more details see https://stripe.com/docs/api#fee_refundss.
	FeeRefunds *feerefund.Client
	// Account is the client used to invoke /account APIs.
	// For more details see https://stripe.com/docs/api#account.
	Account *account.Client
	// Balance is the client used to invoke /balance and transaction-related APIs.
	// For more details see https://stripe.com/docs/api#balance.
	Balance *balance.Client
	// Events is the client used to invoke /events APIs.
	// For more details see https://stripe.com/docs/api#events.
	Events *event.Client
	// Tokens is the client used to invoke /tokens APIs.
	// For more details see https://stripe.com/docs/api#tokens.
	Tokens *token.Client
	// FileUploads is the client used to invoke the uploads /files APIs.
	// For more details see https://stripe.com/docs/api#file_uploads.
	FileUploads *fileupload.Client
	// BitcoinReceivers is the client used to invoke /bitcoin/receivers APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinReceivers *bitcoinreceiver.Client
	// BitcoinTransactions is the client used to invoke /bitcoin/transactions APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinTransactions *bitcointransaction.Client
	// Reversals is the client used to invoke /transfers/reversals APIs.
	Reversals *reversal.Client
	// BankAccounts is the client used to invoke /accounts/bank_accounts APIs.
	BankAccounts *bankaccount.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *Backends) {
	if backends == nil {
		backends = &Backends{GetBackend(APIBackend), GetBackend(UploadsBackend)}
	}

	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.Subs = &sub.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.Discounts = &discount.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
	a.Recipients = &recipient.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.Fees = &fee.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.Account = &account.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.FileUploads = &fileupload.Client{B: backends.Uploads, Key: key}
	a.BitcoinReceivers = &bitcoinreceiver.Client{B: backends.API, Key: key}
	a.BitcoinTransactions = &bitcointransaction.Client{B: backends.API, Key: key}
	a.Reversals = &reversal.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
}
