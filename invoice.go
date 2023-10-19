// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// InvoiceService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewInvoiceService] method instead.
type InvoiceService struct {
	Options []option.RequestOption
}

// NewInvoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvoiceService(opts ...option.RequestOption) (r *InvoiceService) {
	r = &InvoiceService{}
	r.Options = opts
	return
}

// This endpoint is used to create a one-off invoice for a customer.
func (r *InvoiceService) New(ctx context.Context, body InvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all [`Invoice`](../guides/concepts#invoice)s for
// an account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
func (r *InvoiceService) List(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) (res *shared.Page[Invoice], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "invoices"
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

// This endpoint returns a list of all [`Invoice`](../guides/concepts#invoice)s for
// an account in a list format.
//
// The list of invoices is ordered starting from the most recently issued invoice
// date. The response also includes
// [`pagination_metadata`](../reference/pagination), which lets the caller retrieve
// the next page of results if they exist.
//
// By default, this only returns invoices that are `issued`, `paid`, or `synced`.
func (r *InvoiceService) ListAutoPaging(ctx context.Context, query InvoiceListParams, opts ...option.RequestOption) *shared.PageAutoPager[Invoice] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch an [`Invoice`](../guides/concepts#invoice) given
// an identifier.
func (r *InvoiceService) Fetch(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint can be used to fetch the upcoming
// [invoice](../guides/concepts#invoice) for the current billing period given a
// subscription.
func (r *InvoiceService) FetchUpcoming(ctx context.Context, query InvoiceFetchUpcomingParams, opts ...option.RequestOption) (res *InvoiceFetchUpcomingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "invoices/upcoming"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint allows an eligible invoice to be issued manually. This is only
// possible with invoices where status is `draft`, `will_auto_issue` is true, and
// an `eligible_to_issue_at` is a time in the past. Issuing an invoice could
// possibly trigger side effects, some of which could be customer-visible (e.g.
// sending emails, auto-collecting payment, syncing the invoice to external
// providers, etc).
func (r *InvoiceService) Issue(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/issue", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `paid` status. This can
// only be done to invoices that are in the `issued` status.
func (r *InvoiceService) MarkPaid(ctx context.Context, invoiceID string, body InvoiceMarkPaidParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/mark_paid", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows an invoice's status to be set the `void` status. This can
// only be done to invoices that are in the `issued` status.
//
// If the associated invoice has used the customer balance to change the amount
// due, the customer balance operation will be reverted. For example, if the
// invoice used $10 of customer balance, that amount will be added back to the
// customer balance upon voiding.
func (r *InvoiceService) Void(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("invoices/%s/void", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An [`Invoice`](../guides/concepts#invoice) is a fundamental billing entity,
// representing the request for payment for a single subscription. This includes a
// set of line items, which correspond to prices in the subscription's plan and can
// represent fixed recurring fees or usage-based fees. They are generated at the
// end of a billing period, or as the result of an action, such as a cancellation.
type Invoice struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                `json:"amount_due,required"`
	AutoCollection InvoiceAutoCollection `json:"auto_collection,required"`
	BillingAddress InvoiceBillingAddress `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                              `json:"currency,required"`
	Customer                    InvoiceCustomer                     `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
	CustomerTaxID               InvoiceCustomerTaxID                `json:"customer_tax_id,required,nullable"`
	Discount                    InvoiceDiscount                     `json:"discount,required,nullable"`
	Discounts                   []InvoiceDiscount                   `json:"discounts,required"`
	// When the invoice payment is due.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the invoice portal.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// The scheduled date of the invoice
	InvoiceDate time.Time `json:"invoice_date,required" format:"date-time"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf string `json:"invoice_pdf,required,nullable"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceLineItem `json:"line_items,required"`
	Maximum       InvoiceMaximum    `json:"maximum,required,nullable"`
	MaximumAmount string            `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo          string         `json:"memo,required,nullable"`
	Metadata      interface{}    `json:"metadata,required"`
	Minimum       InvoiceMinimum `json:"minimum,required,nullable"`
	MinimumAmount string         `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time              `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  InvoiceShippingAddress `json:"shipping_address,required,nullable"`
	Status           InvoiceStatus          `json:"status,required"`
	Subscription     InvoiceSubscription    `json:"subscription,required,nullable"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal,required"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at,required,nullable" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total,required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool `json:"will_auto_issue,required"`
	JSON          invoiceJSON
}

// invoiceJSON contains the JSON metadata for the struct [Invoice]
type invoiceJSON struct {
	ID                          apijson.Field
	AmountDue                   apijson.Field
	AutoCollection              apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	CreditNotes                 apijson.Field
	Currency                    apijson.Field
	Customer                    apijson.Field
	CustomerBalanceTransactions apijson.Field
	CustomerTaxID               apijson.Field
	Discount                    apijson.Field
	Discounts                   apijson.Field
	DueDate                     apijson.Field
	EligibleToIssueAt           apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceDate                 apijson.Field
	InvoiceNumber               apijson.Field
	InvoicePdf                  apijson.Field
	IssueFailedAt               apijson.Field
	IssuedAt                    apijson.Field
	LineItems                   apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Memo                        apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	PaidAt                      apijson.Field
	PaymentFailedAt             apijson.Field
	PaymentStartedAt            apijson.Field
	ScheduledIssueAt            apijson.Field
	ShippingAddress             apijson.Field
	Status                      apijson.Field
	Subscription                apijson.Field
	Subtotal                    apijson.Field
	SyncFailedAt                apijson.Field
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Invoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceAutoCollection struct {
	// True only if auto-collection is enabled for this invoice.
	Enabled bool `json:"enabled,required,nullable"`
	// If the invoice is scheduled for auto-collection, this field will reflect when
	// the next attempt will occur. If dunning has been exhausted, or auto-collection
	// is not enabled for this invoice, this field will be `null`.
	NextAttemptAt time.Time `json:"next_attempt_at,required,nullable" format:"date-time"`
	// If Orb has ever attempted payment auto-collection for this invoice, this field
	// will reflect when that attempt occurred. In conjunction with `next_attempt_at`,
	// this can be used to tell whether the invoice is currently in dunning (that is,
	// `previously_attempted_at` is non-null, and `next_attempt_time` is non-null), or
	// if dunning has been exhausted (`previously_attempted_at` is non-null, but
	// `next_attempt_time` is null).
	PreviouslyAttemptedAt time.Time `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  invoiceAutoCollectionJSON
}

// invoiceAutoCollectionJSON contains the JSON metadata for the struct
// [InvoiceAutoCollection]
type invoiceAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceBillingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       invoiceBillingAddressJSON
}

// invoiceBillingAddressJSON contains the JSON metadata for the struct
// [InvoiceBillingAddress]
type invoiceBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCreditNote struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	Reason           string `json:"reason,required"`
	Total            string `json:"total,required"`
	Type             string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	JSON     invoiceCreditNoteJSON
}

// invoiceCreditNoteJSON contains the JSON metadata for the struct
// [InvoiceCreditNote]
type invoiceCreditNoteJSON struct {
	ID               apijson.Field
	CreditNoteNumber apijson.Field
	Reason           apijson.Field
	Total            apijson.Field
	Type             apijson.Field
	VoidedAt         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               invoiceCustomerJSON
}

// invoiceCustomerJSON contains the JSON metadata for the struct [InvoiceCustomer]
type invoiceCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                   `json:"id,required"`
	Action InvoiceCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                    `json:"created_at,required" format:"date-time"`
	CreditNote InvoiceCustomerBalanceTransactionsCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                    `json:"ending_balance,required"`
	Invoice       InvoiceCustomerBalanceTransactionsInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                 `json:"starting_balance,required"`
	Type            InvoiceCustomerBalanceTransactionsType `json:"type,required"`
	JSON            invoiceCustomerBalanceTransactionJSON
}

// invoiceCustomerBalanceTransactionJSON contains the JSON metadata for the struct
// [InvoiceCustomerBalanceTransaction]
type invoiceCustomerBalanceTransactionJSON struct {
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

func (r *InvoiceCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomerBalanceTransactionsAction string

const (
	InvoiceCustomerBalanceTransactionsActionAppliedToInvoice InvoiceCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceCustomerBalanceTransactionsActionProratedRefund   InvoiceCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceCustomerBalanceTransactionsActionManualAdjustment InvoiceCustomerBalanceTransactionsAction = "manual_adjustment"
)

type InvoiceCustomerBalanceTransactionsCreditNote struct {
	// The id of the Credit note
	ID   string `json:"id,required"`
	JSON invoiceCustomerBalanceTransactionsCreditNoteJSON
}

// invoiceCustomerBalanceTransactionsCreditNoteJSON contains the JSON metadata for
// the struct [InvoiceCustomerBalanceTransactionsCreditNote]
type invoiceCustomerBalanceTransactionsCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerBalanceTransactionsCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomerBalanceTransactionsInvoice struct {
	// The Invoice id
	ID   string `json:"id,required"`
	JSON invoiceCustomerBalanceTransactionsInvoiceJSON
}

// invoiceCustomerBalanceTransactionsInvoiceJSON contains the JSON metadata for the
// struct [InvoiceCustomerBalanceTransactionsInvoice]
type invoiceCustomerBalanceTransactionsInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerBalanceTransactionsInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceCustomerBalanceTransactionsType string

const (
	InvoiceCustomerBalanceTransactionsTypeIncrement InvoiceCustomerBalanceTransactionsType = "increment"
	InvoiceCustomerBalanceTransactionsTypeDecrement InvoiceCustomerBalanceTransactionsType = "decrement"
)

type InvoiceCustomerTaxID struct {
	Country string `json:"country,required"`
	Type    string `json:"type,required"`
	Value   string `json:"value,required"`
	JSON    invoiceCustomerTaxIDJSON
}

// invoiceCustomerTaxIDJSON contains the JSON metadata for the struct
// [InvoiceCustomerTaxID]
type invoiceCustomerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceCustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The final amount after any discounts or minimums.
	Amount   string          `json:"amount,required"`
	Discount InvoiceDiscount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping      string                  `json:"grouping,required,nullable"`
	Maximum       InvoiceLineItemsMaximum `json:"maximum,required,nullable"`
	MaximumAmount string                  `json:"maximum_amount,required,nullable"`
	Minimum       InvoiceLineItemsMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                  `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// ## Unit pricing
	//
	// With unit pricing, each unit costs a fixed amount.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	        "unit_amount": "0.50"
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered pricing
	//
	// In tiered pricing, the cost of a given unit depends on the tier range that it
	// falls into, where each tier range is defined by an upper and lower bound. For
	// example, the first ten units may cost $0.50 each and all units thereafter may
	// cost $0.10 each.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "tiered",
	//	    "tiered_config": {
	//	        "tiers": [
	//	            {
	//	                "first_unit": 1,
	//	                "last_unit": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "first_unit": 11,
	//	                "last_unit": null,
	//	                "unit_amount": "0.10"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//
	// ```
	//
	// ## Bulk pricing
	//
	// Bulk pricing applies when the number of units determine the cost of all units.
	// For example, if you've bought less than 10 units, they may each be $0.50 for a
	// total of $5.00. Once you've bought more than 10 units, all units may now be
	// priced at $0.40 (i.e. 101 units total would be $40.40).
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bulk",
	//	    "bulk_config": {
	//	        "tiers": [
	//	            {
	//	                "maximum_units": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "maximum_units": 1000,
	//	                "unit_amount": "0.40"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Package pricing
	//
	// Package pricing defines the size or granularity of a unit for billing purposes.
	// For example, if the package size is set to 5, then 4 units will be billed as 5
	// and 6 units will be billed at 10.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "package",
	//	    "package_config": {
	//	        "package_amount": "0.80",
	//	        "package_size": 10
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## BPS pricing
	//
	// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
	// percent (the number of basis points to charge), as well as a cap per event to
	// assess. For example, this would allow you to assess a fee of 0.25% on every
	// payment you process, with a maximum charge of $25 per payment.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bps",
	//	    "bps_config": {
	//	       "bps": 125,
	//	       "per_unit_maximum": "11.00"
	//	    }
	//	    ...
	//	 }
	//
	// ```
	//
	// ## Bulk BPS pricing
	//
	// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
	// total quantity across all events. Similar to bulk pricing, the BPS parameters of
	// a given event depends on the tier range that the billing period falls into. Each
	// tier range is defined by an upper bound. For example, after $1.5M of payment
	// volume is reached, each individual payment may have a lower cap or a smaller
	// take-rate.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "bulk_bps",
	//	    "bulk_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered BPS pricing
	//
	// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
	// event's applicable parameter is a function of its marginal addition to the
	// period total. Similar to tiered pricing, the BPS parameters of a given event
	// depends on the tier range that it falls into, where each tier range is defined
	// by an upper and lower bound. For example, the first few payments may have a 0.8
	// BPS take-rate and all payments after a specific volume may incur a take-rate of
	// 0.5 BPS each.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "tiered_bps",
	//	    "tiered_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "minimum_amount": "0",
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "minimum_amount": "1000000.00",
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Matrix pricing
	//
	// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
	// `dimensions` defines the two event property values evaluated in this pricing
	// model. In a one-dimensional matrix, the second value is `null`. Every
	// configuration has a list of `matrix_values` which give the unit prices for
	// specified property values. In a one-dimensional matrix, the matrix values will
	// have `dimension_values` where the second value of the pair is null. If an event
	// does not match any of the dimension values in the matrix, it will resort to the
	// `default_unit_amount`.
	//
	// ```json
	//
	//	{
	//	    "model_type": "matrix"
	//	    "matrix_config": {
	//	        "default_unit_amount": "3.00",
	//	        "dimensions": [
	//	            "cluster_name",
	//	            "region"
	//	        ],
	//	        "matrix_values": [
	//	            {
	//	                "dimension_values": [
	//	                    "alpha",
	//	                    "west"
	//	                ],
	//	                "unit_amount": "2.00"
	//	            },
	//	            ...
	//	        ]
	//	    }
	//	}
	//
	// ```
	//
	// ### Fixed fees
	//
	// Fixed fees are prices that are applied independent of usage quantities, and
	// follow unit pricing. They also have an additional parameter
	// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
	// the quantity of units applied.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "id": "price_id",
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	       "unit_amount": "2.00"
	//	    },
	//	    "fixed_price_quantity": 3.0
	//	    ...
	//	}
	//
	// ```
	Price    Price   `json:"price,required,nullable"`
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceLineItemsTaxAmount `json:"tax_amounts,required"`
	JSON       invoiceLineItemJSON
}

// invoiceLineItemJSON contains the JSON metadata for the struct [InvoiceLineItem]
type invoiceLineItemJSON struct {
	ID            apijson.Field
	Amount        apijson.Field
	Discount      apijson.Field
	EndDate       apijson.Field
	Grouping      apijson.Field
	Maximum       apijson.Field
	MaximumAmount apijson.Field
	Minimum       apijson.Field
	MinimumAmount apijson.Field
	Name          apijson.Field
	Price         apijson.Field
	Quantity      apijson.Field
	StartDate     apijson.Field
	SubLineItems  apijson.Field
	Subtotal      apijson.Field
	TaxAmounts    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          invoiceLineItemsMaximumJSON
}

// invoiceLineItemsMaximumJSON contains the JSON metadata for the struct
// [InvoiceLineItemsMaximum]
type invoiceLineItemsMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          invoiceLineItemsMinimumJSON
}

// invoiceLineItemsMinimumJSON contains the JSON metadata for the struct
// [InvoiceLineItemsMinimum]
type invoiceLineItemsMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceLineItemsMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [InvoiceLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceLineItemsSubLineItemsOtherSubLineItem].
type InvoiceLineItemsSubLineItem interface {
	implementsInvoiceLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceLineItemsSubLineItem)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"matrix\"",
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsMatrixSubLineItem{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tier\"",
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsTierSubLineItem{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"'null'\"",
			Type:               reflect.TypeOf(InvoiceLineItemsSubLineItemsOtherSubLineItem{}),
		},
	)
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                    `json:"amount,required"`
	Grouping     InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                    `json:"name,required"`
	Quantity     float64                                                   `json:"quantity,required"`
	Type         InvoiceLineItemsSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceLineItemsSubLineItemsMatrixSubLineItemJSON
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsMatrixSubLineItem]
type invoiceLineItemsSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceLineItemsSubLineItemsMatrixSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping]
type invoiceLineItemsSubLineItemsMatrixSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string `json:"dimension_values,required"`
	JSON            invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON
}

// invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON contains the JSON
// metadata for the struct
// [InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsMatrixSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsMatrixSubLineItemTypeMatrix InvoiceLineItemsSubLineItemsMatrixSubLineItemType = "matrix"
)

type InvoiceLineItemsSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                `json:"amount,required"`
	Grouping   InvoiceLineItemsSubLineItemsTierSubLineItemGrouping   `json:"grouping,required,nullable"`
	Name       string                                                `json:"name,required"`
	Quantity   float64                                               `json:"quantity,required"`
	TierConfig InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceLineItemsSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceLineItemsSubLineItemsTierSubLineItemJSON
}

// invoiceLineItemsSubLineItemsTierSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsTierSubLineItem]
type invoiceLineItemsSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceLineItemsSubLineItemsTierSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsTierSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON
}

// invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsTierSubLineItemGrouping]
type invoiceLineItemsSubLineItemsTierSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64 `json:"first_unit,required"`
	LastUnit   float64 `json:"last_unit,required,nullable"`
	UnitAmount string  `json:"unit_amount,required"`
	JSON       invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON
}

// invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig]
type invoiceLineItemsSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsTierSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsTierSubLineItemTypeTier InvoiceLineItemsSubLineItemsTierSubLineItemType = "tier"
)

type InvoiceLineItemsSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                               `json:"amount,required"`
	Grouping InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping `json:"grouping,required,nullable"`
	Name     string                                               `json:"name,required"`
	Quantity float64                                              `json:"quantity,required"`
	Type     InvoiceLineItemsSubLineItemsOtherSubLineItemType     `json:"type,required"`
	JSON     invoiceLineItemsSubLineItemsOtherSubLineItemJSON
}

// invoiceLineItemsSubLineItemsOtherSubLineItemJSON contains the JSON metadata for
// the struct [InvoiceLineItemsSubLineItemsOtherSubLineItem]
type invoiceLineItemsSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceLineItemsSubLineItemsOtherSubLineItem) implementsInvoiceLineItemsSubLineItem() {}

type InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON
}

// invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON contains the JSON
// metadata for the struct [InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping]
type invoiceLineItemsSubLineItemsOtherSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceLineItemsSubLineItemsOtherSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceLineItemsSubLineItemsOtherSubLineItemType string

const (
	InvoiceLineItemsSubLineItemsOtherSubLineItemTypeNull InvoiceLineItemsSubLineItemsOtherSubLineItemType = "'null'"
)

type InvoiceLineItemsTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string `json:"tax_rate_percentage,required,nullable"`
	JSON              invoiceLineItemsTaxAmountJSON
}

// invoiceLineItemsTaxAmountJSON contains the JSON metadata for the struct
// [InvoiceLineItemsTaxAmount]
type invoiceLineItemsTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceLineItemsTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          invoiceMaximumJSON
}

// invoiceMaximumJSON contains the JSON metadata for the struct [InvoiceMaximum]
type invoiceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          invoiceMinimumJSON
}

// invoiceMinimumJSON contains the JSON metadata for the struct [InvoiceMinimum]
type invoiceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceShippingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       invoiceShippingAddressJSON
}

// invoiceShippingAddressJSON contains the JSON metadata for the struct
// [InvoiceShippingAddress]
type invoiceShippingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceStatus string

const (
	InvoiceStatusIssued InvoiceStatus = "issued"
	InvoiceStatusPaid   InvoiceStatus = "paid"
	InvoiceStatusSynced InvoiceStatus = "synced"
	InvoiceStatusVoid   InvoiceStatus = "void"
	InvoiceStatusDraft  InvoiceStatus = "draft"
)

type InvoiceSubscription struct {
	ID   string `json:"id,required"`
	JSON invoiceSubscriptionJSON
}

// invoiceSubscriptionJSON contains the JSON metadata for the struct
// [InvoiceSubscription]
type invoiceSubscriptionJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [InvoiceDiscountPercentageDiscount],
// [InvoiceDiscountTrialDiscount], [InvoiceDiscountUsageDiscount] or
// [InvoiceDiscountAmountDiscount].
type InvoiceDiscount interface {
	implementsInvoiceDiscount()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*InvoiceDiscount)(nil)).Elem(), "")
}

type InvoiceDiscountPercentageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                      `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceDiscountPercentageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	JSON               invoiceDiscountPercentageDiscountJSON
}

// invoiceDiscountPercentageDiscountJSON contains the JSON metadata for the struct
// [InvoiceDiscountPercentageDiscount]
type invoiceDiscountPercentageDiscountJSON struct {
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceDiscountPercentageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceDiscountPercentageDiscount) implementsInvoiceDiscount() {}

type InvoiceDiscountPercentageDiscountDiscountType string

const (
	InvoiceDiscountPercentageDiscountDiscountTypePercentage InvoiceDiscountPercentageDiscountDiscountType = "percentage"
)

type InvoiceDiscountTrialDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                 `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceDiscountTrialDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount string `json:"trial_amount_discount,nullable"`
	// Only available if discount_type is `trial`
	TrialPercentageDiscount float64 `json:"trial_percentage_discount,nullable"`
	JSON                    invoiceDiscountTrialDiscountJSON
}

// invoiceDiscountTrialDiscountJSON contains the JSON metadata for the struct
// [InvoiceDiscountTrialDiscount]
type invoiceDiscountTrialDiscountJSON struct {
	AppliesToPriceIDs       apijson.Field
	DiscountType            apijson.Field
	TrialAmountDiscount     apijson.Field
	TrialPercentageDiscount apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InvoiceDiscountTrialDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceDiscountTrialDiscount) implementsInvoiceDiscount() {}

type InvoiceDiscountTrialDiscountDiscountType string

const (
	InvoiceDiscountTrialDiscountDiscountTypeTrial InvoiceDiscountTrialDiscountDiscountType = "trial"
)

type InvoiceDiscountUsageDiscount struct {
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                 `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceDiscountUsageDiscountDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64 `json:"usage_discount,required"`
	JSON          invoiceDiscountUsageDiscountJSON
}

// invoiceDiscountUsageDiscountJSON contains the JSON metadata for the struct
// [InvoiceDiscountUsageDiscount]
type invoiceDiscountUsageDiscountJSON struct {
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceDiscountUsageDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceDiscountUsageDiscount) implementsInvoiceDiscount() {}

type InvoiceDiscountUsageDiscountDiscountType string

const (
	InvoiceDiscountUsageDiscountDiscountTypeUsage InvoiceDiscountUsageDiscountDiscountType = "usage"
)

type InvoiceDiscountAmountDiscount struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs []string                                  `json:"applies_to_price_ids,required"`
	DiscountType      InvoiceDiscountAmountDiscountDiscountType `json:"discount_type,required"`
	JSON              invoiceDiscountAmountDiscountJSON
}

// invoiceDiscountAmountDiscountJSON contains the JSON metadata for the struct
// [InvoiceDiscountAmountDiscount]
type invoiceDiscountAmountDiscountJSON struct {
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	DiscountType      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceDiscountAmountDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceDiscountAmountDiscount) implementsInvoiceDiscount() {}

type InvoiceDiscountAmountDiscountDiscountType string

const (
	InvoiceDiscountAmountDiscountDiscountTypeAmount InvoiceDiscountAmountDiscountDiscountType = "amount"
)

type InvoiceFetchUpcomingResponse struct {
	ID string `json:"id,required"`
	// This is the final amount required to be charged to the customer and reflects the
	// application of the customer balance to the `total` of the invoice.
	AmountDue      string                                     `json:"amount_due,required"`
	AutoCollection InvoiceFetchUpcomingResponseAutoCollection `json:"auto_collection,required"`
	BillingAddress InvoiceFetchUpcomingResponseBillingAddress `json:"billing_address,required,nullable"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A list of credit notes associated with the invoice
	CreditNotes []InvoiceFetchUpcomingResponseCreditNote `json:"credit_notes,required"`
	// An ISO 4217 currency string or `credits`
	Currency                    string                                                   `json:"currency,required"`
	Customer                    InvoiceFetchUpcomingResponseCustomer                     `json:"customer,required"`
	CustomerBalanceTransactions []InvoiceFetchUpcomingResponseCustomerBalanceTransaction `json:"customer_balance_transactions,required"`
	CustomerTaxID               InvoiceFetchUpcomingResponseCustomerTaxID                `json:"customer_tax_id,required,nullable"`
	Discount                    InvoiceDiscount                                          `json:"discount,required,nullable"`
	Discounts                   []InvoiceDiscount                                        `json:"discounts,required"`
	// When the invoice payment is due.
	DueDate time.Time `json:"due_date,required" format:"date-time"`
	// If the invoice has a status of `draft`, this will be the time that the invoice
	// will be eligible to be issued, otherwise it will be `null`. If `auto-issue` is
	// true, the invoice will automatically begin issuing at this time.
	EligibleToIssueAt time.Time `json:"eligible_to_issue_at,required,nullable" format:"date-time"`
	// A URL for the invoice portal.
	HostedInvoiceURL string `json:"hosted_invoice_url,required,nullable"`
	// Automatically generated invoice number to help track and reconcile invoices.
	// Invoice numbers have a prefix such as `RFOBWG`. These can be sequential per
	// account or customer.
	InvoiceNumber string `json:"invoice_number,required"`
	// The link to download the PDF representation of the `Invoice`.
	InvoicePdf string `json:"invoice_pdf,required,nullable"`
	// If the invoice failed to issue, this will be the last time it failed to issue
	// (even if it is now in a different state.)
	IssueFailedAt time.Time `json:"issue_failed_at,required,nullable" format:"date-time"`
	// If the invoice has been issued, this will be the time it transitioned to
	// `issued` (even if it is now in a different state.)
	IssuedAt time.Time `json:"issued_at,required,nullable" format:"date-time"`
	// The breakdown of prices in this invoice.
	LineItems     []InvoiceFetchUpcomingResponseLineItem `json:"line_items,required"`
	Maximum       InvoiceFetchUpcomingResponseMaximum    `json:"maximum,required,nullable"`
	MaximumAmount string                                 `json:"maximum_amount,required,nullable"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	Memo          string                              `json:"memo,required,nullable"`
	Metadata      interface{}                         `json:"metadata,required"`
	Minimum       InvoiceFetchUpcomingResponseMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                              `json:"minimum_amount,required,nullable"`
	// If the invoice has a status of `paid`, this gives a timestamp when the invoice
	// was paid.
	PaidAt time.Time `json:"paid_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice but failed, this will be the time of
	// the most recent attempt.
	PaymentFailedAt time.Time `json:"payment_failed_at,required,nullable" format:"date-time"`
	// If payment was attempted on this invoice, this will be the start time of the
	// most recent attempt. This field is especially useful for delayed-notification
	// payment mechanisms (like bank transfers), where payment can take 3 days or more.
	PaymentStartedAt time.Time `json:"payment_started_at,required,nullable" format:"date-time"`
	// If the invoice is in draft, this timestamp will reflect when the invoice is
	// scheduled to be issued.
	ScheduledIssueAt time.Time                                   `json:"scheduled_issue_at,required,nullable" format:"date-time"`
	ShippingAddress  InvoiceFetchUpcomingResponseShippingAddress `json:"shipping_address,required,nullable"`
	Status           InvoiceFetchUpcomingResponseStatus          `json:"status,required"`
	Subscription     InvoiceFetchUpcomingResponseSubscription    `json:"subscription,required,nullable"`
	// The total before any discounts and minimums are applied.
	Subtotal string `json:"subtotal,required"`
	// If the invoice failed to sync, this will be the last time an external invoicing
	// provider sync was attempted. This field will always be `null` for invoices using
	// Orb Invoicing.
	SyncFailedAt time.Time `json:"sync_failed_at,required,nullable" format:"date-time"`
	// The scheduled date of the invoice
	TargetDate time.Time `json:"target_date,required" format:"date-time"`
	// The total after any minimums and discounts have been applied.
	Total string `json:"total,required"`
	// If the invoice has a status of `void`, this gives a timestamp when the invoice
	// was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// This is true if the invoice will be automatically issued in the future, and
	// false otherwise.
	WillAutoIssue bool `json:"will_auto_issue,required"`
	JSON          invoiceFetchUpcomingResponseJSON
}

// invoiceFetchUpcomingResponseJSON contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponse]
type invoiceFetchUpcomingResponseJSON struct {
	ID                          apijson.Field
	AmountDue                   apijson.Field
	AutoCollection              apijson.Field
	BillingAddress              apijson.Field
	CreatedAt                   apijson.Field
	CreditNotes                 apijson.Field
	Currency                    apijson.Field
	Customer                    apijson.Field
	CustomerBalanceTransactions apijson.Field
	CustomerTaxID               apijson.Field
	Discount                    apijson.Field
	Discounts                   apijson.Field
	DueDate                     apijson.Field
	EligibleToIssueAt           apijson.Field
	HostedInvoiceURL            apijson.Field
	InvoiceNumber               apijson.Field
	InvoicePdf                  apijson.Field
	IssueFailedAt               apijson.Field
	IssuedAt                    apijson.Field
	LineItems                   apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Memo                        apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	PaidAt                      apijson.Field
	PaymentFailedAt             apijson.Field
	PaymentStartedAt            apijson.Field
	ScheduledIssueAt            apijson.Field
	ShippingAddress             apijson.Field
	Status                      apijson.Field
	Subscription                apijson.Field
	Subtotal                    apijson.Field
	SyncFailedAt                apijson.Field
	TargetDate                  apijson.Field
	Total                       apijson.Field
	VoidedAt                    apijson.Field
	WillAutoIssue               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseAutoCollection struct {
	// True only if auto-collection is enabled for this invoice.
	Enabled bool `json:"enabled,required,nullable"`
	// If the invoice is scheduled for auto-collection, this field will reflect when
	// the next attempt will occur. If dunning has been exhausted, or auto-collection
	// is not enabled for this invoice, this field will be `null`.
	NextAttemptAt time.Time `json:"next_attempt_at,required,nullable" format:"date-time"`
	// If Orb has ever attempted payment auto-collection for this invoice, this field
	// will reflect when that attempt occurred. In conjunction with `next_attempt_at`,
	// this can be used to tell whether the invoice is currently in dunning (that is,
	// `previously_attempted_at` is non-null, and `next_attempt_time` is non-null), or
	// if dunning has been exhausted (`previously_attempted_at` is non-null, but
	// `next_attempt_time` is null).
	PreviouslyAttemptedAt time.Time `json:"previously_attempted_at,required,nullable" format:"date-time"`
	JSON                  invoiceFetchUpcomingResponseAutoCollectionJSON
}

// invoiceFetchUpcomingResponseAutoCollectionJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseAutoCollection]
type invoiceFetchUpcomingResponseAutoCollectionJSON struct {
	Enabled               apijson.Field
	NextAttemptAt         apijson.Field
	PreviouslyAttemptedAt apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseAutoCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseBillingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       invoiceFetchUpcomingResponseBillingAddressJSON
}

// invoiceFetchUpcomingResponseBillingAddressJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseBillingAddress]
type invoiceFetchUpcomingResponseBillingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseBillingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCreditNote struct {
	ID               string `json:"id,required"`
	CreditNoteNumber string `json:"credit_note_number,required"`
	Reason           string `json:"reason,required"`
	Total            string `json:"total,required"`
	Type             string `json:"type,required"`
	// If the credit note has a status of `void`, this gives a timestamp when the
	// credit note was voided.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	JSON     invoiceFetchUpcomingResponseCreditNoteJSON
}

// invoiceFetchUpcomingResponseCreditNoteJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCreditNote]
type invoiceFetchUpcomingResponseCreditNoteJSON struct {
	ID               apijson.Field
	CreditNoteNumber apijson.Field
	Reason           apijson.Field
	Total            apijson.Field
	Type             apijson.Field
	VoidedAt         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               invoiceFetchUpcomingResponseCustomerJSON
}

// invoiceFetchUpcomingResponseCustomerJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCustomer]
type invoiceFetchUpcomingResponseCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransaction struct {
	// A unique id for this transaction.
	ID     string                                                        `json:"id,required"`
	Action InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction `json:"action,required"`
	// The value of the amount changed in the transaction.
	Amount string `json:"amount,required"`
	// The creation time of this transaction.
	CreatedAt  time.Time                                                         `json:"created_at,required" format:"date-time"`
	CreditNote InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote `json:"credit_note,required,nullable"`
	// An optional description provided for manual customer balance adjustments.
	Description string `json:"description,required,nullable"`
	// The new value of the customer's balance prior to the transaction, in the
	// customer's currency.
	EndingBalance string                                                         `json:"ending_balance,required"`
	Invoice       InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice `json:"invoice,required,nullable"`
	// The original value of the customer's balance prior to the transaction, in the
	// customer's currency.
	StartingBalance string                                                      `json:"starting_balance,required"`
	Type            InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType `json:"type,required"`
	JSON            invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON contains the JSON
// metadata for the struct [InvoiceFetchUpcomingResponseCustomerBalanceTransaction]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionJSON struct {
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

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction string

const (
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionAppliedToInvoice InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "applied_to_invoice"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionProratedRefund   InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "prorated_refund"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsActionManualAdjustment InvoiceFetchUpcomingResponseCustomerBalanceTransactionsAction = "manual_adjustment"
)

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote struct {
	// The id of the Credit note
	ID   string `json:"id,required"`
	JSON invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNoteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransactionsCreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice struct {
	// The Invoice id
	ID   string `json:"id,required"`
	JSON invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON
}

// invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON contains the
// JSON metadata for the struct
// [InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice]
type invoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoiceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerBalanceTransactionsInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType string

const (
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeIncrement InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType = "increment"
	InvoiceFetchUpcomingResponseCustomerBalanceTransactionsTypeDecrement InvoiceFetchUpcomingResponseCustomerBalanceTransactionsType = "decrement"
)

type InvoiceFetchUpcomingResponseCustomerTaxID struct {
	Country string `json:"country,required"`
	Type    string `json:"type,required"`
	Value   string `json:"value,required"`
	JSON    invoiceFetchUpcomingResponseCustomerTaxIDJSON
}

// invoiceFetchUpcomingResponseCustomerTaxIDJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseCustomerTaxID]
type invoiceFetchUpcomingResponseCustomerTaxIDJSON struct {
	Country     apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseCustomerTaxID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItem struct {
	// A unique ID for this line item.
	ID string `json:"id,required"`
	// The final amount after any discounts or minimums.
	Amount   string          `json:"amount,required"`
	Discount InvoiceDiscount `json:"discount,required,nullable"`
	// The end date of the range of time applied for this line item's price.
	EndDate time.Time `json:"end_date,required" format:"date-time"`
	// [DEPRECATED] For configured prices that are split by a grouping key, this will
	// be populated with the key and a value. The `amount` and `subtotal` will be the
	// values for this particular grouping.
	Grouping      string                                       `json:"grouping,required,nullable"`
	Maximum       InvoiceFetchUpcomingResponseLineItemsMaximum `json:"maximum,required,nullable"`
	MaximumAmount string                                       `json:"maximum_amount,required,nullable"`
	Minimum       InvoiceFetchUpcomingResponseLineItemsMinimum `json:"minimum,required,nullable"`
	MinimumAmount string                                       `json:"minimum_amount,required,nullable"`
	// The name of the price associated with this line item.
	Name string `json:"name,required"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// ## Unit pricing
	//
	// With unit pricing, each unit costs a fixed amount.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	        "unit_amount": "0.50"
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered pricing
	//
	// In tiered pricing, the cost of a given unit depends on the tier range that it
	// falls into, where each tier range is defined by an upper and lower bound. For
	// example, the first ten units may cost $0.50 each and all units thereafter may
	// cost $0.10 each.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "tiered",
	//	    "tiered_config": {
	//	        "tiers": [
	//	            {
	//	                "first_unit": 1,
	//	                "last_unit": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "first_unit": 11,
	//	                "last_unit": null,
	//	                "unit_amount": "0.10"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//
	// ```
	//
	// ## Bulk pricing
	//
	// Bulk pricing applies when the number of units determine the cost of all units.
	// For example, if you've bought less than 10 units, they may each be $0.50 for a
	// total of $5.00. Once you've bought more than 10 units, all units may now be
	// priced at $0.40 (i.e. 101 units total would be $40.40).
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bulk",
	//	    "bulk_config": {
	//	        "tiers": [
	//	            {
	//	                "maximum_units": 10,
	//	                "unit_amount": "0.50"
	//	            },
	//	            {
	//	                "maximum_units": 1000,
	//	                "unit_amount": "0.40"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Package pricing
	//
	// Package pricing defines the size or granularity of a unit for billing purposes.
	// For example, if the package size is set to 5, then 4 units will be billed as 5
	// and 6 units will be billed at 10.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "package",
	//	    "package_config": {
	//	        "package_amount": "0.80",
	//	        "package_size": 10
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## BPS pricing
	//
	// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
	// percent (the number of basis points to charge), as well as a cap per event to
	// assess. For example, this would allow you to assess a fee of 0.25% on every
	// payment you process, with a maximum charge of $25 per payment.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "model_type": "bps",
	//	    "bps_config": {
	//	       "bps": 125,
	//	       "per_unit_maximum": "11.00"
	//	    }
	//	    ...
	//	 }
	//
	// ```
	//
	// ## Bulk BPS pricing
	//
	// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
	// total quantity across all events. Similar to bulk pricing, the BPS parameters of
	// a given event depends on the tier range that the billing period falls into. Each
	// tier range is defined by an upper bound. For example, after $1.5M of payment
	// volume is reached, each individual payment may have a lower cap or a smaller
	// take-rate.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "bulk_bps",
	//	    "bulk_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Tiered BPS pricing
	//
	// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
	// event's applicable parameter is a function of its marginal addition to the
	// period total. Similar to tiered pricing, the BPS parameters of a given event
	// depends on the tier range that it falls into, where each tier range is defined
	// by an upper and lower bound. For example, the first few payments may have a 0.8
	// BPS take-rate and all payments after a specific volume may incur a take-rate of
	// 0.5 BPS each.
	//
	// ```json
	//
	//	    ...
	//	    "model_type": "tiered_bps",
	//	    "tiered_bps_config": {
	//	        "tiers": [
	//	           {
	//	                "minimum_amount": "0",
	//	                "maximum_amount": "1000000.00",
	//	                "bps": 125,
	//	                "per_unit_maximum": "19.00"
	//	           },
	//	          {
	//	                "minimum_amount": "1000000.00",
	//	                "maximum_amount": null,
	//	                "bps": 115,
	//	                "per_unit_maximum": "4.00"
	//	            }
	//	        ]
	//	    }
	//	    ...
	//	}
	//
	// ```
	//
	// ## Matrix pricing
	//
	// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
	// `dimensions` defines the two event property values evaluated in this pricing
	// model. In a one-dimensional matrix, the second value is `null`. Every
	// configuration has a list of `matrix_values` which give the unit prices for
	// specified property values. In a one-dimensional matrix, the matrix values will
	// have `dimension_values` where the second value of the pair is null. If an event
	// does not match any of the dimension values in the matrix, it will resort to the
	// `default_unit_amount`.
	//
	// ```json
	//
	//	{
	//	    "model_type": "matrix"
	//	    "matrix_config": {
	//	        "default_unit_amount": "3.00",
	//	        "dimensions": [
	//	            "cluster_name",
	//	            "region"
	//	        ],
	//	        "matrix_values": [
	//	            {
	//	                "dimension_values": [
	//	                    "alpha",
	//	                    "west"
	//	                ],
	//	                "unit_amount": "2.00"
	//	            },
	//	            ...
	//	        ]
	//	    }
	//	}
	//
	// ```
	//
	// ### Fixed fees
	//
	// Fixed fees are prices that are applied independent of usage quantities, and
	// follow unit pricing. They also have an additional parameter
	// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
	// the quantity of units applied.
	//
	// ```json
	//
	//	{
	//	    ...
	//	    "id": "price_id",
	//	    "model_type": "unit",
	//	    "unit_config": {
	//	       "unit_amount": "2.00"
	//	    },
	//	    "fixed_price_quantity": 3.0
	//	    ...
	//	}
	//
	// ```
	Price    Price   `json:"price,required,nullable"`
	Quantity float64 `json:"quantity,required"`
	// The start date of the range of time applied for this line item's price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// For complex pricing structures, the line item can be broken down further in
	// `sub_line_items`.
	SubLineItems []InvoiceFetchUpcomingResponseLineItemsSubLineItem `json:"sub_line_items,required"`
	// The line amount before any line item-specific discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// An array of tax rates and their incurred tax amounts. Empty if no tax
	// integration is configured.
	TaxAmounts []InvoiceFetchUpcomingResponseLineItemsTaxAmount `json:"tax_amounts,required"`
	JSON       invoiceFetchUpcomingResponseLineItemJSON
}

// invoiceFetchUpcomingResponseLineItemJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseLineItem]
type invoiceFetchUpcomingResponseLineItemJSON struct {
	ID            apijson.Field
	Amount        apijson.Field
	Discount      apijson.Field
	EndDate       apijson.Field
	Grouping      apijson.Field
	Maximum       apijson.Field
	MaximumAmount apijson.Field
	Minimum       apijson.Field
	MinimumAmount apijson.Field
	Name          apijson.Field
	Price         apijson.Field
	Quantity      apijson.Field
	StartDate     apijson.Field
	SubLineItems  apijson.Field
	Subtotal      apijson.Field
	TaxAmounts    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          invoiceFetchUpcomingResponseLineItemsMaximumJSON
}

// invoiceFetchUpcomingResponseLineItemsMaximumJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseLineItemsMaximum]
type invoiceFetchUpcomingResponseLineItemsMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          invoiceFetchUpcomingResponseLineItemsMinimumJSON
}

// invoiceFetchUpcomingResponseLineItemsMinimumJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseLineItemsMinimum]
type invoiceFetchUpcomingResponseLineItemsMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem],
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem] or
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem].
type InvoiceFetchUpcomingResponseLineItemsSubLineItem interface {
	implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvoiceFetchUpcomingResponseLineItemsSubLineItem)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"matrix\"",
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tier\"",
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"'null'\"",
			Type:               reflect.TypeOf(InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem{}),
		},
	)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem struct {
	// The total amount for this sub line item.
	Amount       string                                                                         `json:"amount,required"`
	Grouping     InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping     `json:"grouping,required,nullable"`
	MatrixConfig InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig `json:"matrix_config,required"`
	Name         string                                                                         `json:"name,required"`
	Quantity     float64                                                                        `json:"quantity,required"`
	Type         InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType         `json:"type,required"`
	JSON         invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemJSON struct {
	Amount       apijson.Field
	Grouping     apijson.Field
	MatrixConfig apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig struct {
	// The ordered dimension values for this line item.
	DimensionValues []string `json:"dimension_values,required"`
	JSON            invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfigJSON struct {
	DimensionValues apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemTypeMatrix InvoiceFetchUpcomingResponseLineItemsSubLineItemsMatrixSubLineItemType = "matrix"
)

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem struct {
	// The total amount for this sub line item.
	Amount     string                                                                     `json:"amount,required"`
	Grouping   InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping   `json:"grouping,required,nullable"`
	Name       string                                                                     `json:"name,required"`
	Quantity   float64                                                                    `json:"quantity,required"`
	TierConfig InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig `json:"tier_config,required"`
	Type       InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType       `json:"type,required"`
	JSON       invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	TierConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig struct {
	FirstUnit  float64 `json:"first_unit,required"`
	LastUnit   float64 `json:"last_unit,required,nullable"`
	UnitAmount string  `json:"unit_amount,required"`
	JSON       invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfigJSON struct {
	FirstUnit   apijson.Field
	LastUnit    apijson.Field
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTierConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemTypeTier InvoiceFetchUpcomingResponseLineItemsSubLineItemsTierSubLineItemType = "tier"
)

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem struct {
	// The total amount for this sub line item.
	Amount   string                                                                    `json:"amount,required"`
	Grouping InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping `json:"grouping,required,nullable"`
	Name     string                                                                    `json:"name,required"`
	Quantity float64                                                                   `json:"quantity,required"`
	Type     InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType     `json:"type,required"`
	JSON     invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON contains
// the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemJSON struct {
	Amount      apijson.Field
	Grouping    apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItem) implementsInvoiceFetchUpcomingResponseLineItemsSubLineItem() {
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping struct {
	Key string `json:"key,required"`
	// No value indicates the default group
	Value string `json:"value,required,nullable"`
	JSON  invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON
}

// invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON
// contains the JSON metadata for the struct
// [InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping]
type invoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGroupingJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemGrouping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType string

const (
	InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemTypeNull InvoiceFetchUpcomingResponseLineItemsSubLineItemsOtherSubLineItemType = "'null'"
)

type InvoiceFetchUpcomingResponseLineItemsTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string `json:"tax_rate_percentage,required,nullable"`
	JSON              invoiceFetchUpcomingResponseLineItemsTaxAmountJSON
}

// invoiceFetchUpcomingResponseLineItemsTaxAmountJSON contains the JSON metadata
// for the struct [InvoiceFetchUpcomingResponseLineItemsTaxAmount]
type invoiceFetchUpcomingResponseLineItemsTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseLineItemsTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          invoiceFetchUpcomingResponseMaximumJSON
}

// invoiceFetchUpcomingResponseMaximumJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseMaximum]
type invoiceFetchUpcomingResponseMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          invoiceFetchUpcomingResponseMinimumJSON
}

// invoiceFetchUpcomingResponseMinimumJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseMinimum]
type invoiceFetchUpcomingResponseMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseShippingAddress struct {
	City       string `json:"city,required,nullable"`
	Country    string `json:"country,required,nullable"`
	Line1      string `json:"line1,required,nullable"`
	Line2      string `json:"line2,required,nullable"`
	PostalCode string `json:"postal_code,required,nullable"`
	State      string `json:"state,required,nullable"`
	JSON       invoiceFetchUpcomingResponseShippingAddressJSON
}

// invoiceFetchUpcomingResponseShippingAddressJSON contains the JSON metadata for
// the struct [InvoiceFetchUpcomingResponseShippingAddress]
type invoiceFetchUpcomingResponseShippingAddressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	PostalCode  apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseShippingAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceFetchUpcomingResponseStatus string

const (
	InvoiceFetchUpcomingResponseStatusIssued InvoiceFetchUpcomingResponseStatus = "issued"
	InvoiceFetchUpcomingResponseStatusPaid   InvoiceFetchUpcomingResponseStatus = "paid"
	InvoiceFetchUpcomingResponseStatusSynced InvoiceFetchUpcomingResponseStatus = "synced"
	InvoiceFetchUpcomingResponseStatusVoid   InvoiceFetchUpcomingResponseStatus = "void"
	InvoiceFetchUpcomingResponseStatusDraft  InvoiceFetchUpcomingResponseStatus = "draft"
)

type InvoiceFetchUpcomingResponseSubscription struct {
	ID   string `json:"id,required"`
	JSON invoiceFetchUpcomingResponseSubscriptionJSON
}

// invoiceFetchUpcomingResponseSubscriptionJSON contains the JSON metadata for the
// struct [InvoiceFetchUpcomingResponseSubscription]
type invoiceFetchUpcomingResponseSubscriptionJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvoiceFetchUpcomingResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceNewParams struct {
	// An ISO 4217 currency string. Must be the same as the customer's currency if it
	// is set.
	Currency param.Field[string] `json:"currency,required"`
	// Optional invoice date to set. Must be in the past, if not set, `invoice_date` is
	// set to the current time in the customer's timezone.
	InvoiceDate param.Field[time.Time]                  `json:"invoice_date,required" format:"date-time"`
	LineItems   param.Field[[]InvoiceNewParamsLineItem] `json:"line_items,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of '0' here represents that the
	// invoice is due on issue, whereas a value of 30 represents that the customer has
	// 30 days to pay the invoice.
	NetTerms param.Field[int64] `json:"net_terms,required"`
	// The id of the `Customer` to create this invoice for. One of `customer_id` and
	// `external_customer_id` are required.
	CustomerID param.Field[string] `json:"customer_id"`
	// The `external_customer_id` of the `Customer` to create this invoice for. One of
	// `customer_id` and `external_customer_id` are required.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// An optional memo to attach to the invoice.
	Memo param.Field[string] `json:"memo"`
	// When true, this invoice will automatically be issued upon creation. When false,
	// the resulting invoice will require manual review to issue. Defaulted to false.
	WillAutoIssue param.Field[bool] `json:"will_auto_issue"`
}

func (r InvoiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsLineItem struct {
	// A date string to specify the line item's end date in the customer's timezone.
	EndDate   param.Field[time.Time]                          `json:"end_date,required" format:"date"`
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[InvoiceNewParamsLineItemsModelType] `json:"model_type,required"`
	// The name of the line item.
	Name param.Field[string] `json:"name,required"`
	// The number of units on the line item
	Quantity param.Field[float64] `json:"quantity,required"`
	// A date string to specify the line item's start date in the customer's timezone.
	StartDate  param.Field[time.Time]                           `json:"start_date,required" format:"date"`
	UnitConfig param.Field[InvoiceNewParamsLineItemsUnitConfig] `json:"unit_config,required"`
}

func (r InvoiceNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceNewParamsLineItemsModelType string

const (
	InvoiceNewParamsLineItemsModelTypeUnit InvoiceNewParamsLineItemsModelType = "unit"
)

type InvoiceNewParamsLineItemsUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r InvoiceNewParamsLineItemsUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvoiceListParams struct {
	Amount   param.Field[string] `query:"amount"`
	AmountGt param.Field[string] `query:"amount[gt]"`
	AmountLt param.Field[string] `query:"amount[lt]"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor             param.Field[string]                    `query:"cursor"`
	CustomerID         param.Field[string]                    `query:"customer_id"`
	DateType           param.Field[InvoiceListParamsDateType] `query:"date_type"`
	DueDate            param.Field[time.Time]                 `query:"due_date" format:"date"`
	DueDateWindow      param.Field[string]                    `query:"due_date_window"`
	DueDateGt          param.Field[time.Time]                 `query:"due_date[gt]" format:"date"`
	DueDateLt          param.Field[time.Time]                 `query:"due_date[lt]" format:"date"`
	ExternalCustomerID param.Field[string]                    `query:"external_customer_id"`
	InvoiceDateGt      param.Field[time.Time]                 `query:"invoice_date[gt]" format:"date-time"`
	InvoiceDateGte     param.Field[time.Time]                 `query:"invoice_date[gte]" format:"date-time"`
	InvoiceDateLt      param.Field[time.Time]                 `query:"invoice_date[lt]" format:"date-time"`
	InvoiceDateLte     param.Field[time.Time]                 `query:"invoice_date[lte]" format:"date-time"`
	IsRecurring        param.Field[bool]                      `query:"is_recurring"`
	// The number of items to fetch. Defaults to 20.
	Limit          param.Field[int64]                   `query:"limit"`
	Status         param.Field[InvoiceListParamsStatus] `query:"status"`
	SubscriptionID param.Field[string]                  `query:"subscription_id"`
}

// URLQuery serializes [InvoiceListParams]'s query parameters as `url.Values`.
func (r InvoiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceListParamsDateType string

const (
	InvoiceListParamsDateTypeDueDate     InvoiceListParamsDateType = "due_date"
	InvoiceListParamsDateTypeInvoiceDate InvoiceListParamsDateType = "invoice_date"
)

type InvoiceListParamsStatus string

const (
	InvoiceListParamsStatusDraft  InvoiceListParamsStatus = "draft"
	InvoiceListParamsStatusIssued InvoiceListParamsStatus = "issued"
	InvoiceListParamsStatusPaid   InvoiceListParamsStatus = "paid"
	InvoiceListParamsStatusSynced InvoiceListParamsStatus = "synced"
	InvoiceListParamsStatusVoid   InvoiceListParamsStatus = "void"
)

type InvoiceFetchUpcomingParams struct {
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [InvoiceFetchUpcomingParams]'s query parameters as
// `url.Values`.
func (r InvoiceFetchUpcomingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvoiceMarkPaidParams struct {
	// An optional external ID to associate with the payment.
	ExternalID param.Field[string] `json:"external_id,required"`
	// An optional note to associate with the payment.
	Notes param.Field[string] `json:"notes,required"`
	// A date string to specify the date of the payment.
	PaymentReceivedDate param.Field[time.Time] `json:"payment_received_date,required" format:"date"`
}

func (r InvoiceMarkPaidParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
