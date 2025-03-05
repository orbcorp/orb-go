// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
func (r *CreditNoteService) New(ctx context.Context, body CreditNoteNewParams, opts ...option.RequestOption) (res *shared.CreditNoteModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "credit_notes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a paginated list of CreditNotes. Users can also filter by customer_id,
// subscription_id, or external_customer_id. The credit notes will be returned in
// reverse chronological order by `creation_time`.
func (r *CreditNoteService) List(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) (res *pagination.Page[shared.CreditNoteModel], err error) {
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
func (r *CreditNoteService) ListAutoPaging(ctx context.Context, query CreditNoteListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.CreditNoteModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch a single [`Credit Note`](/invoicing/credit-notes)
// given an identifier.
func (r *CreditNoteService) Fetch(ctx context.Context, creditNoteID string, opts ...option.RequestOption) (res *shared.CreditNoteModel, err error) {
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
	// An optional memo to attach to the credit note.
	Memo param.Field[string] `json:"memo"`
	// An optional reason for the credit note.
	Reason param.Field[CreditNoteNewParamsReason] `json:"reason"`
}

func (r CreditNoteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditNoteNewParamsLineItem struct {
	// The total amount in the invoice's currency to credit this line item.
	Amount param.Field[string] `json:"amount,required"`
	// The ID of the line item to credit.
	InvoiceLineItemID param.Field[string] `json:"invoice_line_item_id,required"`
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
