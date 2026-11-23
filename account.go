// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"github.com/orbcorp/orb-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
	// A billing entity is a legal entity an account invoices from, with its own legal
	// name, address and tax registrations.
	BillingEntities *AccountBillingEntityService
	// A payment gateway connection is one Stripe account or one Adyen merchant account
	// that a billing entity collects payments through.
	PaymentGatewayConnections *AccountPaymentGatewayConnectionService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.BillingEntities = NewAccountBillingEntityService(opts...)
	r.PaymentGatewayConnections = NewAccountPaymentGatewayConnectionService(opts...)
	return
}
