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

// CreditNoteService contains methods and other services that help with interacting
// with the orb API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCreditNoteService] method instead.
type CreditNoteService struct {
	Options []option.RequestOption
}

// NewCreditNoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditNoteService(opts ...option.RequestOption) (r *CreditNoteService) {
	r = &CreditNoteService{}
	r.Options = opts
	return
}

// Get a paginated list of CreditNotes. Users can also filter by customer_id,
// subscription_id, or external_customer_id. The credit notes will be returned in
// reverse chronological order by `creation_time`.
func (r *CreditNoteService) List(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) (res *shared.Page[CreditNote], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "credit_notes"
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

// Get a paginated list of CreditNotes. Users can also filter by customer_id,
// subscription_id, or external_customer_id. The credit notes will be returned in
// reverse chronological order by `creation_time`.
func (r *CreditNoteService) ListAutoPaging(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) *shared.PageAutoPager[CreditNote] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch a single
// [`Credit Note`](../guides/invoicing/credit-notes) given an identifier.
func (r *CreditNoteService) Fetch(ctx context.Context, creditNoteID string, opts ...option.RequestOption) (res *CreditNote, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("credit_notes/%s", creditNoteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// The [Credit Note](/guides/invoicing/credit-notes) resource represents a credit
// that has been applied to a particular invoice.
type CreditNote struct {
	// The Orb id of this credit note.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The unique identifier for credit notes.
	CreditNoteNumber string `json:"credit_note_number,required"`
	// A URL to a PDF of the credit note.
	CreditNotePdf string             `json:"credit_note_pdf,required,nullable"`
	Customer      CreditNoteCustomer `json:"customer,required"`
	// Any discounts applied on the original invoice.
	Discounts []interface{} `json:"discounts,required"`
	// The id of the invoice resource that this credit note is applied to.
	InvoiceID string `json:"invoice_id,required"`
	// All of the line items associated with this credit note.
	LineItems []CreditNoteLineItem `json:"line_items,required"`
	// The maximum amount applied on the original invoice
	MaximumAmountAdjustment interface{} `json:"maximum_amount_adjustment,required,nullable"`
	// An optional memo supplied on the credit note.
	Memo string `json:"memo,required,nullable"`
	// Any credited amount from the applied minimum on the invoice.
	MinimumAmountRefunded string           `json:"minimum_amount_refunded,required,nullable"`
	Reason                CreditNoteReason `json:"reason,required"`
	// The total prior to any creditable invoice-level discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// The total including creditable invoice-level discounts or minimums, and tax.
	Total string         `json:"total,required"`
	Type  CreditNoteType `json:"type,required"`
	// The time at which the credit note was voided in Orb, if applicable.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	JSON     creditNoteJSON
}

// creditNoteJSON contains the JSON metadata for the struct [CreditNote]
type creditNoteJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	CreditNoteNumber        apijson.Field
	CreditNotePdf           apijson.Field
	Customer                apijson.Field
	Discounts               apijson.Field
	InvoiceID               apijson.Field
	LineItems               apijson.Field
	MaximumAmountAdjustment apijson.Field
	Memo                    apijson.Field
	MinimumAmountRefunded   apijson.Field
	Reason                  apijson.Field
	Subtotal                apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	VoidedAt                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreditNoteCustomer struct {
	ID                 string `json:"id,required"`
	ExternalCustomerID string `json:"external_customer_id,required,nullable"`
	JSON               creditNoteCustomerJSON
}

// creditNoteCustomerJSON contains the JSON metadata for the struct
// [CreditNoteCustomer]
type creditNoteCustomerJSON struct {
	ID                 apijson.Field
	ExternalCustomerID apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteCustomer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreditNoteLineItem struct {
	// The Orb id of this resource.
	ID string `json:"id,required"`
	// The amount of the line item, including any line item minimums and discounts.
	Amount string `json:"amount,required"`
	// Any line items discounts from the invoice's line item.
	Discounts []interface{} `json:"discounts,required"`
	// The name of the corresponding invoice line item.
	Name string `json:"name,required"`
	// An optional quantity credited.
	Quantity float64 `json:"quantity,required,nullable"`
	// Any sub line items that may be credited.
	SubLineItems []CreditNoteLineItemsSubLineItem `json:"sub_line_items,required"`
	// The amount of the line item, excluding any line item minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Any tax amounts applied onto the line item.
	TaxAmounts []interface{} `json:"tax_amounts,required"`
	JSON       creditNoteLineItemJSON
}

// creditNoteLineItemJSON contains the JSON metadata for the struct
// [CreditNoteLineItem]
type creditNoteLineItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	Discounts    apijson.Field
	Name         apijson.Field
	Quantity     apijson.Field
	SubLineItems apijson.Field
	Subtotal     apijson.Field
	TaxAmounts   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CreditNoteLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreditNoteLineItemsSubLineItem struct {
	Amount   string  `json:"amount,required"`
	Name     string  `json:"name,required"`
	Quantity float64 `json:"quantity,required,nullable"`
	JSON     creditNoteLineItemsSubLineItemJSON
}

// creditNoteLineItemsSubLineItemJSON contains the JSON metadata for the struct
// [CreditNoteLineItemsSubLineItem]
type creditNoteLineItemsSubLineItemJSON struct {
	Amount      apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteLineItemsSubLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreditNoteReason string

const (
	CreditNoteReasonDuplicate             CreditNoteReason = "Duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "Fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "Order change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "Product unsatisfactory"
)

type CreditNoteType string

const (
	CreditNoteTypeRefund     CreditNoteType = "refund"
	CreditNoteTypeAdjustment CreditNoteType = "adjustment"
)

type CreditNoteListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CreditNoteListParams]'s query parameters as `url.Values`.
func (r CreditNoteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
