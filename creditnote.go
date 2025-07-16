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

// This endpoint is used to create a single
// [`Credit Note`](/invoicing/credit-notes).
//
// The credit note service period configuration supports two explicit modes:
//
//  1. Global service periods: Specify start_date and end_date at the credit note
//     level. These dates will be applied to all line items uniformly.
//
//  2. Individual service periods: Specify start_date and end_date for each line
//     item. When using this mode, ALL line items must have individual periods
//     specified.
//
//  3. Default behavior: If no service periods are specified (neither global nor
//     individual), the original invoice line item service periods will be used.
//
// Note: Mixing global and individual service periods in the same request is not
// allowed to prevent confusion.
//
// Service period dates are normalized to the start of the day in the customer's
// timezone to ensure consistent handling across different timezones.
//
// Date Format: Use start_date and end_date with format "YYYY-MM-DD" (e.g.,
// "2023-09-22") to match other Orb APIs like /v1/invoice_line_items.
//
// Note: Both start_date and end_date are inclusive - the service period will cover
// both the start date and end date completely (from start of start_date to end of
// end_date).
func (r *CreditNoteService) New(ctx context.Context, body CreditNoteNewParams, opts ...option.RequestOption) (res *shared.CreditNote, err error) {
	opts = append(r.Options[:], opts...)
	path := "credit_notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a paginated list of CreditNotes. Users can also filter by customer_id,
// subscription_id, or external_customer_id. The credit notes will be returned in
// reverse chronological order by `creation_time`.
func (r *CreditNoteService) List(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CreditNote], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *CreditNoteService) ListAutoPaging(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CreditNote] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch a single [`Credit Note`](/invoicing/credit-notes)
// given an identifier.
func (r *CreditNoteService) Fetch(ctx context.Context, creditNoteID string, opts ...option.RequestOption) (res *shared.CreditNote, err error) {
	opts = append(r.Options[:], opts...)
	if creditNoteID == "" {
		err = errors.New("missing required credit_note_id parameter")
		return
	}
	path := fmt.Sprintf("credit_notes/%s", creditNoteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CreditNoteNewParams struct {
	LineItems param.Field[[]CreditNoteNewParamsLineItem] `json:"line_items,required"`
	// An optional reason for the credit note.
	Reason param.Field[CreditNoteNewParamsReason] `json:"reason,required"`
	// A date string to specify the global credit note service period end date in the
	// customer's timezone. This will be applied to all line items that don't have
	// their own individual service periods specified. If not provided, line items will
	// use their original invoice line item service periods. This date is inclusive.
	EndDate param.Field[time.Time] `json:"end_date" format:"date"`
	// An optional memo to attach to the credit note.
	Memo param.Field[string] `json:"memo"`
	// A date string to specify the global credit note service period start date in the
	// customer's timezone. This will be applied to all line items that don't have
	// their own individual service periods specified. If not provided, line items will
	// use their original invoice line item service periods. This date is inclusive.
	StartDate param.Field[time.Time] `json:"start_date" format:"date"`
}

func (r CreditNoteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditNoteNewParamsLineItem struct {
	// The total amount in the invoice's currency to credit this line item.
	Amount param.Field[string] `json:"amount,required"`
	// The ID of the line item to credit.
	InvoiceLineItemID param.Field[string] `json:"invoice_line_item_id,required"`
	// A date string to specify this line item's credit note service period end date in
	// the customer's timezone. If provided, this will be used for this specific line
	// item. If not provided, will use the global end_date if available, otherwise
	// defaults to the original invoice line item's end date. This date is inclusive.
	EndDate param.Field[time.Time] `json:"end_date" format:"date"`
	// A date string to specify this line item's credit note service period start date
	// in the customer's timezone. If provided, this will be used for this specific
	// line item. If not provided, will use the global start_date if available,
	// otherwise defaults to the original invoice line item's start date. This date is
	// inclusive.
	StartDate param.Field[time.Time] `json:"start_date" format:"date"`
}

func (r CreditNoteNewParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An optional reason for the credit note.
type CreditNoteNewParamsReason string

const (
	CreditNoteNewParamsReasonDuplicate             CreditNoteNewParamsReason = "duplicate"
	CreditNoteNewParamsReasonFraudulent            CreditNoteNewParamsReason = "fraudulent"
	CreditNoteNewParamsReasonOrderChange           CreditNoteNewParamsReason = "order_change"
	CreditNoteNewParamsReasonProductUnsatisfactory CreditNoteNewParamsReason = "product_unsatisfactory"
)

func (r CreditNoteNewParamsReason) IsKnown() bool {
	switch r {
	case CreditNoteNewParamsReasonDuplicate, CreditNoteNewParamsReasonFraudulent, CreditNoteNewParamsReasonOrderChange, CreditNoteNewParamsReasonProductUnsatisfactory:
		return true
	}
	return false
}

type CreditNoteListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
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
