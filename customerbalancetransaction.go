// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// CustomerBalanceTransactionService contains methods and other services that help
// with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerBalanceTransactionService] method instead.
type CustomerBalanceTransactionService struct {
	Options []option.RequestOption
}

// NewCustomerBalanceTransactionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomerBalanceTransactionService(opts ...option.RequestOption) (r *CustomerBalanceTransactionService) {
	r = &CustomerBalanceTransactionService{}
	r.Options = opts
	return
}

// Creates an immutable balance transaction that updates the customer's balance and
// returns back the newly created transaction.
func (r *CustomerBalanceTransactionService) New(ctx context.Context, customerID string, body CustomerBalanceTransactionNewParams, opts ...option.RequestOption) (res *shared.CustomerBalanceTransactionModel, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/balance_transactions", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// ## The customer balance
//
// The customer balance is an amount in the customer's currency, which Orb
// automatically applies to subsequent invoices. This balance can be adjusted
// manually via Orb's webapp on the customer details page. You can use this balance
// to provide a fixed mid-period credit to the customer. Commonly, this is done due
// to system downtime/SLA violation, or an adhoc adjustment discussed with the
// customer.
//
// If the balance is a positive value at the time of invoicing, it represents that
// the customer has credit that should be used to offset the amount due on the next
// issued invoice. In this case, Orb will automatically reduce the next invoice by
// the balance amount, and roll over any remaining balance if the invoice is fully
// discounted.
//
// If the balance is a negative value at the time of invoicing, Orb will increase
// the invoice's amount due with a positive adjustment, and reset the balance to 0.
//
// This endpoint retrieves all customer balance transactions in reverse
// chronological order for a single customer, providing a complete audit trail of
// all adjustments and invoice applications.
//
// ## Eligibility
//
// The customer balance can only be applied to invoices or adjusted manually if
// invoices are not synced to a separate invoicing provider. If a payment gateway
// such as Stripe is used, the balance will be applied to the invoice before
// forwarding payment to the gateway.
func (r *CustomerBalanceTransactionService) List(ctx context.Context, customerID string, query CustomerBalanceTransactionListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CustomerBalanceTransactionModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("customers/%s/balance_transactions", customerID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// ## The customer balance
//
// The customer balance is an amount in the customer's currency, which Orb
// automatically applies to subsequent invoices. This balance can be adjusted
// manually via Orb's webapp on the customer details page. You can use this balance
// to provide a fixed mid-period credit to the customer. Commonly, this is done due
// to system downtime/SLA violation, or an adhoc adjustment discussed with the
// customer.
//
// If the balance is a positive value at the time of invoicing, it represents that
// the customer has credit that should be used to offset the amount due on the next
// issued invoice. In this case, Orb will automatically reduce the next invoice by
// the balance amount, and roll over any remaining balance if the invoice is fully
// discounted.
//
// If the balance is a negative value at the time of invoicing, Orb will increase
// the invoice's amount due with a positive adjustment, and reset the balance to 0.
//
// This endpoint retrieves all customer balance transactions in reverse
// chronological order for a single customer, providing a complete audit trail of
// all adjustments and invoice applications.
//
// ## Eligibility
//
// The customer balance can only be applied to invoices or adjusted manually if
// invoices are not synced to a separate invoicing provider. If a payment gateway
// such as Stripe is used, the balance will be applied to the invoice before
// forwarding payment to the gateway.
func (r *CustomerBalanceTransactionService) ListAutoPaging(ctx context.Context, customerID string, query CustomerBalanceTransactionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CustomerBalanceTransactionModel] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

type CustomerBalanceTransactionNewParams struct {
	Amount param.Field[string]                                  `json:"amount,required"`
	Type   param.Field[CustomerBalanceTransactionNewParamsType] `json:"type,required"`
	// An optional description that can be specified around this entry.
	Description param.Field[string] `json:"description"`
}

func (r CustomerBalanceTransactionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerBalanceTransactionNewParamsType string

const (
	CustomerBalanceTransactionNewParamsTypeIncrement CustomerBalanceTransactionNewParamsType = "increment"
	CustomerBalanceTransactionNewParamsTypeDecrement CustomerBalanceTransactionNewParamsType = "decrement"
)

func (r CustomerBalanceTransactionNewParamsType) IsKnown() bool {
	switch r {
	case CustomerBalanceTransactionNewParamsTypeIncrement, CustomerBalanceTransactionNewParamsTypeDecrement:
		return true
	}
	return false
}

type CustomerBalanceTransactionListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit            param.Field[int64]     `query:"limit"`
	OperationTimeGt  param.Field[time.Time] `query:"operation_time[gt]" format:"date-time"`
	OperationTimeGte param.Field[time.Time] `query:"operation_time[gte]" format:"date-time"`
	OperationTimeLt  param.Field[time.Time] `query:"operation_time[lt]" format:"date-time"`
	OperationTimeLte param.Field[time.Time] `query:"operation_time[lte]" format:"date-time"`
}

// URLQuery serializes [CustomerBalanceTransactionListParams]'s query parameters as
// `url.Values`.
func (r CustomerBalanceTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
