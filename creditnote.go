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
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// CreditNoteService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditNoteService] method instead.
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
func (r *CreditNoteService) List(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) (res *pagination.Page[CreditNote], err error) {
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
func (r *CreditNoteService) ListAutoPaging(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CreditNote] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch a single
// [`Credit Note`](../guides/invoicing/credit-notes) given an identifier.
func (r *CreditNoteService) Fetch(ctx context.Context, creditNoteID string, opts ...option.RequestOption) (res *CreditNote, err error) {
	opts = append(r.Options[:], opts...)
	if creditNoteID == "" {
		err = errors.New("missing required credit_note_id parameter")
		return
	}
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
	// The id of the invoice resource that this credit note is applied to.
	InvoiceID string `json:"invoice_id,required"`
	// All of the line items associated with this credit note.
	LineItems []CreditNoteLineItem `json:"line_items,required"`
	// The maximum amount applied on the original invoice
	MaximumAmountAdjustment CreditNoteMaximumAmountAdjustment `json:"maximum_amount_adjustment,required,nullable"`
	// An optional memo supplied on the credit note.
	Memo string `json:"memo,required,nullable"`
	// Any credited amount from the applied minimum on the invoice.
	MinimumAmountRefunded string           `json:"minimum_amount_refunded,required,nullable"`
	Reason                CreditNoteReason `json:"reason,required,nullable"`
	// The total prior to any creditable invoice-level discounts or minimums.
	Subtotal string `json:"subtotal,required"`
	// The total including creditable invoice-level discounts or minimums, and tax.
	Total string         `json:"total,required"`
	Type  CreditNoteType `json:"type,required"`
	// The time at which the credit note was voided in Orb, if applicable.
	VoidedAt time.Time `json:"voided_at,required,nullable" format:"date-time"`
	// Any discounts applied on the original invoice.
	Discounts []CreditNoteDiscount `json:"discounts"`
	JSON      creditNoteJSON       `json:"-"`
}

// creditNoteJSON contains the JSON metadata for the struct [CreditNote]
type creditNoteJSON struct {
	ID                      apijson.Field
	CreatedAt               apijson.Field
	CreditNoteNumber        apijson.Field
	CreditNotePdf           apijson.Field
	Customer                apijson.Field
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
	Discounts               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CreditNote) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteJSON) RawJSON() string {
	return r.raw
}

type CreditNoteCustomer struct {
	ID                 string                 `json:"id,required"`
	ExternalCustomerID string                 `json:"external_customer_id,required,nullable"`
	JSON               creditNoteCustomerJSON `json:"-"`
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

func (r creditNoteCustomerJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItem struct {
	// The Orb id of this resource.
	ID string `json:"id,required"`
	// The amount of the line item, including any line item minimums and discounts.
	Amount string `json:"amount,required"`
	// The name of the corresponding invoice line item.
	Name string `json:"name,required"`
	// An optional quantity credited.
	Quantity float64 `json:"quantity,required,nullable"`
	// The amount of the line item, excluding any line item minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Any tax amounts applied onto the line item.
	TaxAmounts []CreditNoteLineItemsTaxAmount `json:"tax_amounts,required"`
	// Any line item discounts from the invoice's line item.
	Discounts []CreditNoteLineItemsDiscount `json:"discounts"`
	JSON      creditNoteLineItemJSON        `json:"-"`
}

// creditNoteLineItemJSON contains the JSON metadata for the struct
// [CreditNoteLineItem]
type creditNoteLineItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	Name        apijson.Field
	Quantity    apijson.Field
	Subtotal    apijson.Field
	TaxAmounts  apijson.Field
	Discounts   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteLineItemJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItemsTaxAmount struct {
	// The amount of additional tax incurred by this tax rate.
	Amount string `json:"amount,required"`
	// The human-readable description of the applied tax rate.
	TaxRateDescription string `json:"tax_rate_description,required"`
	// The tax rate percentage, out of 100.
	TaxRatePercentage string                           `json:"tax_rate_percentage,required,nullable"`
	JSON              creditNoteLineItemsTaxAmountJSON `json:"-"`
}

// creditNoteLineItemsTaxAmountJSON contains the JSON metadata for the struct
// [CreditNoteLineItemsTaxAmount]
type creditNoteLineItemsTaxAmountJSON struct {
	Amount             apijson.Field
	TaxRateDescription apijson.Field
	TaxRatePercentage  apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteLineItemsTaxAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteLineItemsTaxAmountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItemsDiscount struct {
	ID                 string                                   `json:"id,required"`
	AmountApplied      string                                   `json:"amount_applied,required"`
	AppliesToPriceIDs  []string                                 `json:"applies_to_price_ids,required"`
	DiscountType       CreditNoteLineItemsDiscountsDiscountType `json:"discount_type,required"`
	PercentageDiscount float64                                  `json:"percentage_discount,required"`
	AmountDiscount     string                                   `json:"amount_discount,nullable"`
	Reason             string                                   `json:"reason,nullable"`
	JSON               creditNoteLineItemsDiscountJSON          `json:"-"`
}

// creditNoteLineItemsDiscountJSON contains the JSON metadata for the struct
// [CreditNoteLineItemsDiscount]
type creditNoteLineItemsDiscountJSON struct {
	ID                 apijson.Field
	AmountApplied      apijson.Field
	AppliesToPriceIDs  apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AmountDiscount     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteLineItemsDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteLineItemsDiscountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteLineItemsDiscountsDiscountType string

const (
	CreditNoteLineItemsDiscountsDiscountTypePercentage CreditNoteLineItemsDiscountsDiscountType = "percentage"
	CreditNoteLineItemsDiscountsDiscountTypeAmount     CreditNoteLineItemsDiscountsDiscountType = "amount"
)

func (r CreditNoteLineItemsDiscountsDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteLineItemsDiscountsDiscountTypePercentage, CreditNoteLineItemsDiscountsDiscountTypeAmount:
		return true
	}
	return false
}

// The maximum amount applied on the original invoice
type CreditNoteMaximumAmountAdjustment struct {
	AmountApplied      string                                            `json:"amount_applied,required"`
	DiscountType       CreditNoteMaximumAmountAdjustmentDiscountType     `json:"discount_type,required"`
	PercentageDiscount float64                                           `json:"percentage_discount,required"`
	AppliesToPrices    []CreditNoteMaximumAmountAdjustmentAppliesToPrice `json:"applies_to_prices,nullable"`
	Reason             string                                            `json:"reason,nullable"`
	JSON               creditNoteMaximumAmountAdjustmentJSON             `json:"-"`
}

// creditNoteMaximumAmountAdjustmentJSON contains the JSON metadata for the struct
// [CreditNoteMaximumAmountAdjustment]
type creditNoteMaximumAmountAdjustmentJSON struct {
	AmountApplied      apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPrices    apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteMaximumAmountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteMaximumAmountAdjustmentJSON) RawJSON() string {
	return r.raw
}

type CreditNoteMaximumAmountAdjustmentDiscountType string

const (
	CreditNoteMaximumAmountAdjustmentDiscountTypePercentage CreditNoteMaximumAmountAdjustmentDiscountType = "percentage"
)

func (r CreditNoteMaximumAmountAdjustmentDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteMaximumAmountAdjustmentDiscountTypePercentage:
		return true
	}
	return false
}

type CreditNoteMaximumAmountAdjustmentAppliesToPrice struct {
	ID   string                                              `json:"id,required"`
	Name string                                              `json:"name,required"`
	JSON creditNoteMaximumAmountAdjustmentAppliesToPriceJSON `json:"-"`
}

// creditNoteMaximumAmountAdjustmentAppliesToPriceJSON contains the JSON metadata
// for the struct [CreditNoteMaximumAmountAdjustmentAppliesToPrice]
type creditNoteMaximumAmountAdjustmentAppliesToPriceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteMaximumAmountAdjustmentAppliesToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteMaximumAmountAdjustmentAppliesToPriceJSON) RawJSON() string {
	return r.raw
}

type CreditNoteReason string

const (
	CreditNoteReasonDuplicate             CreditNoteReason = "Duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "Fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "Order change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "Product unsatisfactory"
)

func (r CreditNoteReason) IsKnown() bool {
	switch r {
	case CreditNoteReasonDuplicate, CreditNoteReasonFraudulent, CreditNoteReasonOrderChange, CreditNoteReasonProductUnsatisfactory:
		return true
	}
	return false
}

type CreditNoteType string

const (
	CreditNoteTypeRefund     CreditNoteType = "refund"
	CreditNoteTypeAdjustment CreditNoteType = "adjustment"
)

func (r CreditNoteType) IsKnown() bool {
	switch r {
	case CreditNoteTypeRefund, CreditNoteTypeAdjustment:
		return true
	}
	return false
}

type CreditNoteDiscount struct {
	AmountApplied      string                              `json:"amount_applied,required"`
	DiscountType       CreditNoteDiscountsDiscountType     `json:"discount_type,required"`
	PercentageDiscount float64                             `json:"percentage_discount,required"`
	AppliesToPrices    []CreditNoteDiscountsAppliesToPrice `json:"applies_to_prices,nullable"`
	Reason             string                              `json:"reason,nullable"`
	JSON               creditNoteDiscountJSON              `json:"-"`
}

// creditNoteDiscountJSON contains the JSON metadata for the struct
// [CreditNoteDiscount]
type creditNoteDiscountJSON struct {
	AmountApplied      apijson.Field
	DiscountType       apijson.Field
	PercentageDiscount apijson.Field
	AppliesToPrices    apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CreditNoteDiscount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountJSON) RawJSON() string {
	return r.raw
}

type CreditNoteDiscountsDiscountType string

const (
	CreditNoteDiscountsDiscountTypePercentage CreditNoteDiscountsDiscountType = "percentage"
)

func (r CreditNoteDiscountsDiscountType) IsKnown() bool {
	switch r {
	case CreditNoteDiscountsDiscountTypePercentage:
		return true
	}
	return false
}

type CreditNoteDiscountsAppliesToPrice struct {
	ID   string                                `json:"id,required"`
	Name string                                `json:"name,required"`
	JSON creditNoteDiscountsAppliesToPriceJSON `json:"-"`
}

// creditNoteDiscountsAppliesToPriceJSON contains the JSON metadata for the struct
// [CreditNoteDiscountsAppliesToPrice]
type creditNoteDiscountsAppliesToPriceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditNoteDiscountsAppliesToPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditNoteDiscountsAppliesToPriceJSON) RawJSON() string {
	return r.raw
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
