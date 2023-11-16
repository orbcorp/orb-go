// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
)

// CustomerBalanceTransactionService contains methods and other services that help
// with interacting with the orb API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the
// [NewCustomerBalanceTransactionService] method instead.
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
func (r *CustomerBalanceTransactionService) New(ctx context.Context, customerID string, body CustomerBalanceTransactionNewParams, opts ...option.RequestOption) (res *CustomerBalanceTransactionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *CustomerBalanceTransactionService) List(ctx context.Context, customerID string, query CustomerBalanceTransactionListParams, opts ...option.RequestOption) (res *shared.Page[CustomerBalanceTransactionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *CustomerBalanceTransactionService) ListAutoPaging(ctx context.Context, customerID string, query CustomerBalanceTransactionListParams, opts ...option.RequestOption) *shared.PageAutoPager[CustomerBalanceTransactionListResponse] {
	return shared.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

type CustomerBalanceTransactionNewResponse struct {
	// A unique id for this transaction.
	ID     string                                      `json:"id,required"`
	Action CustomerBalanceTransactionNewResponseAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                       `json:"created_at,required" format:"date-time"`
	CreditNote CustomerBalanceTransactionNewResponseCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                       `json:"ending_balance,required"`
	Invoice       CustomerBalanceTransactionNewResponseInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                    `json:"starting_balance,required"`
	Type            CustomerBalanceTransactionNewResponseType `json:"type,required"`
	JSON            customerBalanceTransactionNewResponseJSON `json:"-"`
}

// customerBalanceTransactionNewResponseJSON contains the JSON metadata for the
// struct [CustomerBalanceTransactionNewResponse]
type customerBalanceTransactionNewResponseJSON struct {
	ID              apijson.Field
	Action          apijson.Field
	Amount          apijson.Field
	CreatedAt       apijson.Field
	CreditNote      apijson.Field
	Description     apijson.Field
	EndingBalance   apijson.Field
	Invoice         apijson.Field
	StartingBalance apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerBalanceTransactionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionNewResponseAction string

const (
	CustomerBalanceTransactionNewResponseActionAppliedToInvoice     CustomerBalanceTransactionNewResponseAction = "applied_to_invoice"
	CustomerBalanceTransactionNewResponseActionManualAdjustment     CustomerBalanceTransactionNewResponseAction = "manual_adjustment"
	CustomerBalanceTransactionNewResponseActionProratedRefund       CustomerBalanceTransactionNewResponseAction = "prorated_refund"
	CustomerBalanceTransactionNewResponseActionRevertProratedRefund CustomerBalanceTransactionNewResponseAction = "revert_prorated_refund"
	CustomerBalanceTransactionNewResponseActionReturnFromVoiding    CustomerBalanceTransactionNewResponseAction = "return_from_voiding"
	CustomerBalanceTransactionNewResponseActionCreditNoteApplied    CustomerBalanceTransactionNewResponseAction = "credit_note_applied"
	CustomerBalanceTransactionNewResponseActionCreditNoteVoided     CustomerBalanceTransactionNewResponseAction = "credit_note_voided"
	CustomerBalanceTransactionNewResponseActionOverpaymentRefund    CustomerBalanceTransactionNewResponseAction = "overpayment_refund"
)

type CustomerBalanceTransactionNewResponseCreditNote struct {
	// The id of the Credit note
	ID   string                                              `json:"id,required"`
	JSON customerBalanceTransactionNewResponseCreditNoteJSON `json:"-"`
}

// customerBalanceTransactionNewResponseCreditNoteJSON contains the JSON metadata
// for the struct [CustomerBalanceTransactionNewResponseCreditNote]
type customerBalanceTransactionNewResponseCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionNewResponseCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionNewResponseInvoice struct {
	// The Invoice id
	ID   string                                           `json:"id,required"`
	JSON customerBalanceTransactionNewResponseInvoiceJSON `json:"-"`
}

// customerBalanceTransactionNewResponseInvoiceJSON contains the JSON metadata for
// the struct [CustomerBalanceTransactionNewResponseInvoice]
type customerBalanceTransactionNewResponseInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionNewResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionNewResponseType string

const (
	CustomerBalanceTransactionNewResponseTypeIncrement CustomerBalanceTransactionNewResponseType = "increment"
	CustomerBalanceTransactionNewResponseTypeDecrement CustomerBalanceTransactionNewResponseType = "decrement"
)

type CustomerBalanceTransactionListResponse struct {
	// A unique id for this transaction.
	ID     string                                       `json:"id,required"`
	Action CustomerBalanceTransactionListResponseAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                        `json:"created_at,required" format:"date-time"`
	CreditNote CustomerBalanceTransactionListResponseCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                        `json:"ending_balance,required"`
	Invoice       CustomerBalanceTransactionListResponseInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                     `json:"starting_balance,required"`
	Type            CustomerBalanceTransactionListResponseType `json:"type,required"`
	JSON            customerBalanceTransactionListResponseJSON `json:"-"`
}

// customerBalanceTransactionListResponseJSON contains the JSON metadata for the
// struct [CustomerBalanceTransactionListResponse]
type customerBalanceTransactionListResponseJSON struct {
	ID              apijson.Field
	Action          apijson.Field
	Amount          apijson.Field
	CreatedAt       apijson.Field
	CreditNote      apijson.Field
	Description     apijson.Field
	EndingBalance   apijson.Field
	Invoice         apijson.Field
	StartingBalance apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CustomerBalanceTransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionListResponseAction string

const (
	CustomerBalanceTransactionListResponseActionAppliedToInvoice     CustomerBalanceTransactionListResponseAction = "applied_to_invoice"
	CustomerBalanceTransactionListResponseActionManualAdjustment     CustomerBalanceTransactionListResponseAction = "manual_adjustment"
	CustomerBalanceTransactionListResponseActionProratedRefund       CustomerBalanceTransactionListResponseAction = "prorated_refund"
	CustomerBalanceTransactionListResponseActionRevertProratedRefund CustomerBalanceTransactionListResponseAction = "revert_prorated_refund"
	CustomerBalanceTransactionListResponseActionReturnFromVoiding    CustomerBalanceTransactionListResponseAction = "return_from_voiding"
	CustomerBalanceTransactionListResponseActionCreditNoteApplied    CustomerBalanceTransactionListResponseAction = "credit_note_applied"
	CustomerBalanceTransactionListResponseActionCreditNoteVoided     CustomerBalanceTransactionListResponseAction = "credit_note_voided"
	CustomerBalanceTransactionListResponseActionOverpaymentRefund    CustomerBalanceTransactionListResponseAction = "overpayment_refund"
)

type CustomerBalanceTransactionListResponseCreditNote struct {
	// The id of the Credit note
	ID   string                                               `json:"id,required"`
	JSON customerBalanceTransactionListResponseCreditNoteJSON `json:"-"`
}

// customerBalanceTransactionListResponseCreditNoteJSON contains the JSON metadata
// for the struct [CustomerBalanceTransactionListResponseCreditNote]
type customerBalanceTransactionListResponseCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionListResponseCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionListResponseInvoice struct {
	// The Invoice id
	ID   string                                            `json:"id,required"`
	JSON customerBalanceTransactionListResponseInvoiceJSON `json:"-"`
}

// customerBalanceTransactionListResponseInvoiceJSON contains the JSON metadata for
// the struct [CustomerBalanceTransactionListResponseInvoice]
type customerBalanceTransactionListResponseInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerBalanceTransactionListResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomerBalanceTransactionListResponseType string

const (
	CustomerBalanceTransactionListResponseTypeIncrement CustomerBalanceTransactionListResponseType = "increment"
	CustomerBalanceTransactionListResponseTypeDecrement CustomerBalanceTransactionListResponseType = "decrement"
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
