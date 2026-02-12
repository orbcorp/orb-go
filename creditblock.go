// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// CreditBlockService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditBlockService] method instead.
type CreditBlockService struct {
	Options []option.RequestOption
}

// NewCreditBlockService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditBlockService(opts ...option.RequestOption) (r *CreditBlockService) {
	r = &CreditBlockService{}
	r.Options = opts
	return
}

// This endpoint returns a credit block identified by its block_id.
func (r *CreditBlockService) Get(ctx context.Context, blockID string, opts ...option.RequestOption) (res *CreditBlockGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if blockID == "" {
		err = errors.New("missing required block_id parameter")
		return
	}
	path := fmt.Sprintf("credit_blocks/%s", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint deletes a credit block by its ID.
//
// When a credit block is deleted:
//
//   - The block is removed from the customer's credit ledger.
//   - Any usage of the credit block is reversed, and the ledger is replayed as if
//     the block never existed.
//   - If invoices were generated from the purchase of the credit block, they will be
//     deleted if in draft status, voided if issued, or a credit note will be issued
//     if the invoice is paid.
//
// <Note>
// Issued invoices that had credits applied from this block will not be regenerated, but the ledger will
// reflect the state as if credits from the deleted block were never applied.
// </Note>
func (r *CreditBlockService) Delete(ctx context.Context, blockID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if blockID == "" {
		err = errors.New("missing required block_id parameter")
		return
	}
	path := fmt.Sprintf("credit_blocks/%s", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// The Credit Block resource models prepaid credits within Orb.
type CreditBlockGetResponse struct {
	ID                    string                         `json:"id,required"`
	Balance               float64                        `json:"balance,required"`
	EffectiveDate         time.Time                      `json:"effective_date,required,nullable" format:"date-time"`
	ExpiryDate            time.Time                      `json:"expiry_date,required,nullable" format:"date-time"`
	Filters               []CreditBlockGetResponseFilter `json:"filters,required"`
	MaximumInitialBalance float64                        `json:"maximum_initial_balance,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string            `json:"metadata,required"`
	PerUnitCostBasis string                       `json:"per_unit_cost_basis,required,nullable"`
	Status           CreditBlockGetResponseStatus `json:"status,required"`
	JSON             creditBlockGetResponseJSON   `json:"-"`
}

// creditBlockGetResponseJSON contains the JSON metadata for the struct
// [CreditBlockGetResponse]
type creditBlockGetResponseJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CreditBlockGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockGetResponseJSON) RawJSON() string {
	return r.raw
}

type CreditBlockGetResponseFilter struct {
	// The property of the price to filter on.
	Field CreditBlockGetResponseFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockGetResponseFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                         `json:"values,required"`
	JSON   creditBlockGetResponseFilterJSON `json:"-"`
}

// creditBlockGetResponseFilterJSON contains the JSON metadata for the struct
// [CreditBlockGetResponseFilter]
type creditBlockGetResponseFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockGetResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockGetResponseFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockGetResponseFiltersField string

const (
	CreditBlockGetResponseFiltersFieldPriceID       CreditBlockGetResponseFiltersField = "price_id"
	CreditBlockGetResponseFiltersFieldItemID        CreditBlockGetResponseFiltersField = "item_id"
	CreditBlockGetResponseFiltersFieldPriceType     CreditBlockGetResponseFiltersField = "price_type"
	CreditBlockGetResponseFiltersFieldCurrency      CreditBlockGetResponseFiltersField = "currency"
	CreditBlockGetResponseFiltersFieldPricingUnitID CreditBlockGetResponseFiltersField = "pricing_unit_id"
)

func (r CreditBlockGetResponseFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseFiltersFieldPriceID, CreditBlockGetResponseFiltersFieldItemID, CreditBlockGetResponseFiltersFieldPriceType, CreditBlockGetResponseFiltersFieldCurrency, CreditBlockGetResponseFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockGetResponseFiltersOperator string

const (
	CreditBlockGetResponseFiltersOperatorIncludes CreditBlockGetResponseFiltersOperator = "includes"
	CreditBlockGetResponseFiltersOperatorExcludes CreditBlockGetResponseFiltersOperator = "excludes"
)

func (r CreditBlockGetResponseFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseFiltersOperatorIncludes, CreditBlockGetResponseFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockGetResponseStatus string

const (
	CreditBlockGetResponseStatusActive         CreditBlockGetResponseStatus = "active"
	CreditBlockGetResponseStatusPendingPayment CreditBlockGetResponseStatus = "pending_payment"
)

func (r CreditBlockGetResponseStatus) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseStatusActive, CreditBlockGetResponseStatusPendingPayment:
		return true
	}
	return false
}
