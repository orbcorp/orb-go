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
	"github.com/orbcorp/orb-go/shared"
)

// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
// credits within Orb.
//
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
		return nil, err
	}
	path := fmt.Sprintf("credit_blocks/%s", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
		return err
	}
	path := fmt.Sprintf("credit_blocks/%s", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint returns the credit block and its associated purchasing invoices.
//
// If a credit block was purchased (as opposed to being manually added), this
// endpoint returns the invoices that were created to charge the customer for the
// credit block. For credit blocks with payment schedules spanning multiple periods
// (e.g., monthly payments over 12 months), multiple invoices will be returned.
//
// For credit blocks created by subscription allocation prices, this endpoint
// returns the subscription invoice containing the allocation line item that
// created the block.
//
// If the credit block was not purchased (e.g., manual increment), an empty
// invoices list is returned.
//
// **Note: This endpoint is currently experimental and its interface may change in
// future releases. Please contact support before building production integrations
// against this endpoint.**
func (r *CreditBlockService) ListInvoices(ctx context.Context, blockID string, opts ...option.RequestOption) (res *CreditBlockListInvoicesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if blockID == "" {
		err = errors.New("missing required block_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("credit_blocks/%s/invoices", blockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The Credit Block resource models prepaid credits within Orb.
type CreditBlockGetResponse struct {
	ID      string  `json:"id" api:"required"`
	Balance float64 `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CreditBlockGetResponseCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                               `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                               `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CreditBlockGetResponseFilter          `json:"filters" api:"required"`
	MaximumInitialBalance float64                                 `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string            `json:"metadata" api:"required"`
	PerUnitCostBasis string                       `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CreditBlockGetResponseStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CreditBlockGetResponseCreditAllocation `json:"credit_allocation" api:"nullable"`
	JSON             creditBlockGetResponseJSON             `json:"-"`
}

// creditBlockGetResponseJSON contains the JSON metadata for the struct
// [CreditBlockGetResponse]
type creditBlockGetResponseJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	CreditBlockSource     apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	CreditAllocation      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CreditBlockGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockGetResponseJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CreditBlockGetResponseCreditBlockSource string

const (
	CreditBlockGetResponseCreditBlockSourceAllocation CreditBlockGetResponseCreditBlockSource = "allocation"
	CreditBlockGetResponseCreditBlockSourceTopUp      CreditBlockGetResponseCreditBlockSource = "top_up"
	CreditBlockGetResponseCreditBlockSourceManual     CreditBlockGetResponseCreditBlockSource = "manual"
)

func (r CreditBlockGetResponseCreditBlockSource) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseCreditBlockSourceAllocation, CreditBlockGetResponseCreditBlockSourceTopUp, CreditBlockGetResponseCreditBlockSourceManual:
		return true
	}
	return false
}

type CreditBlockGetResponseFilter struct {
	// The property of the price to filter on.
	Field CreditBlockGetResponseFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockGetResponseFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                         `json:"values" api:"required"`
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

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CreditBlockGetResponseCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                         `json:"item_id" api:"required"`
	Filters       []CreditBlockGetResponseCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                         `json:"license_type_id" api:"nullable"`
	JSON          creditBlockGetResponseCreditAllocationJSON     `json:"-"`
}

// creditBlockGetResponseCreditAllocationJSON contains the JSON metadata for the
// struct [CreditBlockGetResponseCreditAllocation]
type creditBlockGetResponseCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditBlockGetResponseCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockGetResponseCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CreditBlockGetResponseCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CreditBlockGetResponseCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockGetResponseCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                         `json:"values" api:"required"`
	JSON   creditBlockGetResponseCreditAllocationFilterJSON `json:"-"`
}

// creditBlockGetResponseCreditAllocationFilterJSON contains the JSON metadata for
// the struct [CreditBlockGetResponseCreditAllocationFilter]
type creditBlockGetResponseCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockGetResponseCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockGetResponseCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockGetResponseCreditAllocationFiltersField string

const (
	CreditBlockGetResponseCreditAllocationFiltersFieldPriceID       CreditBlockGetResponseCreditAllocationFiltersField = "price_id"
	CreditBlockGetResponseCreditAllocationFiltersFieldItemID        CreditBlockGetResponseCreditAllocationFiltersField = "item_id"
	CreditBlockGetResponseCreditAllocationFiltersFieldPriceType     CreditBlockGetResponseCreditAllocationFiltersField = "price_type"
	CreditBlockGetResponseCreditAllocationFiltersFieldCurrency      CreditBlockGetResponseCreditAllocationFiltersField = "currency"
	CreditBlockGetResponseCreditAllocationFiltersFieldPricingUnitID CreditBlockGetResponseCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CreditBlockGetResponseCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseCreditAllocationFiltersFieldPriceID, CreditBlockGetResponseCreditAllocationFiltersFieldItemID, CreditBlockGetResponseCreditAllocationFiltersFieldPriceType, CreditBlockGetResponseCreditAllocationFiltersFieldCurrency, CreditBlockGetResponseCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockGetResponseCreditAllocationFiltersOperator string

const (
	CreditBlockGetResponseCreditAllocationFiltersOperatorIncludes CreditBlockGetResponseCreditAllocationFiltersOperator = "includes"
	CreditBlockGetResponseCreditAllocationFiltersOperatorExcludes CreditBlockGetResponseCreditAllocationFiltersOperator = "excludes"
)

func (r CreditBlockGetResponseCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockGetResponseCreditAllocationFiltersOperatorIncludes, CreditBlockGetResponseCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockListInvoicesResponse struct {
	// The Credit Block resource models prepaid credits within Orb.
	Block    CreditBlockListInvoicesResponseBlock     `json:"block" api:"required"`
	Invoices []CreditBlockListInvoicesResponseInvoice `json:"invoices" api:"required"`
	JSON     creditBlockListInvoicesResponseJSON      `json:"-"`
}

// creditBlockListInvoicesResponseJSON contains the JSON metadata for the struct
// [CreditBlockListInvoicesResponse]
type creditBlockListInvoicesResponseJSON struct {
	Block       apijson.Field
	Invoices    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseJSON) RawJSON() string {
	return r.raw
}

// The Credit Block resource models prepaid credits within Orb.
type CreditBlockListInvoicesResponseBlock struct {
	ID      string  `json:"id" api:"required"`
	Balance float64 `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CreditBlockListInvoicesResponseBlockCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                             `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                             `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CreditBlockListInvoicesResponseBlockFilter          `json:"filters" api:"required"`
	MaximumInitialBalance float64                                               `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                          `json:"metadata" api:"required"`
	PerUnitCostBasis string                                     `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CreditBlockListInvoicesResponseBlockStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CreditBlockListInvoicesResponseBlockCreditAllocation `json:"credit_allocation" api:"nullable"`
	JSON             creditBlockListInvoicesResponseBlockJSON             `json:"-"`
}

// creditBlockListInvoicesResponseBlockJSON contains the JSON metadata for the
// struct [CreditBlockListInvoicesResponseBlock]
type creditBlockListInvoicesResponseBlockJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	CreditBlockSource     apijson.Field
	EffectiveDate         apijson.Field
	ExpiryDate            apijson.Field
	Filters               apijson.Field
	MaximumInitialBalance apijson.Field
	Metadata              apijson.Field
	PerUnitCostBasis      apijson.Field
	Status                apijson.Field
	CreditAllocation      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponseBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseBlockJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CreditBlockListInvoicesResponseBlockCreditBlockSource string

const (
	CreditBlockListInvoicesResponseBlockCreditBlockSourceAllocation CreditBlockListInvoicesResponseBlockCreditBlockSource = "allocation"
	CreditBlockListInvoicesResponseBlockCreditBlockSourceTopUp      CreditBlockListInvoicesResponseBlockCreditBlockSource = "top_up"
	CreditBlockListInvoicesResponseBlockCreditBlockSourceManual     CreditBlockListInvoicesResponseBlockCreditBlockSource = "manual"
)

func (r CreditBlockListInvoicesResponseBlockCreditBlockSource) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockCreditBlockSourceAllocation, CreditBlockListInvoicesResponseBlockCreditBlockSourceTopUp, CreditBlockListInvoicesResponseBlockCreditBlockSourceManual:
		return true
	}
	return false
}

type CreditBlockListInvoicesResponseBlockFilter struct {
	// The property of the price to filter on.
	Field CreditBlockListInvoicesResponseBlockFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockListInvoicesResponseBlockFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                       `json:"values" api:"required"`
	JSON   creditBlockListInvoicesResponseBlockFilterJSON `json:"-"`
}

// creditBlockListInvoicesResponseBlockFilterJSON contains the JSON metadata for
// the struct [CreditBlockListInvoicesResponseBlockFilter]
type creditBlockListInvoicesResponseBlockFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponseBlockFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseBlockFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockListInvoicesResponseBlockFiltersField string

const (
	CreditBlockListInvoicesResponseBlockFiltersFieldPriceID       CreditBlockListInvoicesResponseBlockFiltersField = "price_id"
	CreditBlockListInvoicesResponseBlockFiltersFieldItemID        CreditBlockListInvoicesResponseBlockFiltersField = "item_id"
	CreditBlockListInvoicesResponseBlockFiltersFieldPriceType     CreditBlockListInvoicesResponseBlockFiltersField = "price_type"
	CreditBlockListInvoicesResponseBlockFiltersFieldCurrency      CreditBlockListInvoicesResponseBlockFiltersField = "currency"
	CreditBlockListInvoicesResponseBlockFiltersFieldPricingUnitID CreditBlockListInvoicesResponseBlockFiltersField = "pricing_unit_id"
)

func (r CreditBlockListInvoicesResponseBlockFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockFiltersFieldPriceID, CreditBlockListInvoicesResponseBlockFiltersFieldItemID, CreditBlockListInvoicesResponseBlockFiltersFieldPriceType, CreditBlockListInvoicesResponseBlockFiltersFieldCurrency, CreditBlockListInvoicesResponseBlockFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockListInvoicesResponseBlockFiltersOperator string

const (
	CreditBlockListInvoicesResponseBlockFiltersOperatorIncludes CreditBlockListInvoicesResponseBlockFiltersOperator = "includes"
	CreditBlockListInvoicesResponseBlockFiltersOperatorExcludes CreditBlockListInvoicesResponseBlockFiltersOperator = "excludes"
)

func (r CreditBlockListInvoicesResponseBlockFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockFiltersOperatorIncludes, CreditBlockListInvoicesResponseBlockFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockListInvoicesResponseBlockStatus string

const (
	CreditBlockListInvoicesResponseBlockStatusActive         CreditBlockListInvoicesResponseBlockStatus = "active"
	CreditBlockListInvoicesResponseBlockStatusPendingPayment CreditBlockListInvoicesResponseBlockStatus = "pending_payment"
)

func (r CreditBlockListInvoicesResponseBlockStatus) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockStatusActive, CreditBlockListInvoicesResponseBlockStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CreditBlockListInvoicesResponseBlockCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                                       `json:"item_id" api:"required"`
	Filters       []CreditBlockListInvoicesResponseBlockCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                                       `json:"license_type_id" api:"nullable"`
	JSON          creditBlockListInvoicesResponseBlockCreditAllocationJSON     `json:"-"`
}

// creditBlockListInvoicesResponseBlockCreditAllocationJSON contains the JSON
// metadata for the struct [CreditBlockListInvoicesResponseBlockCreditAllocation]
type creditBlockListInvoicesResponseBlockCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponseBlockCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseBlockCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CreditBlockListInvoicesResponseBlockCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                       `json:"values" api:"required"`
	JSON   creditBlockListInvoicesResponseBlockCreditAllocationFilterJSON `json:"-"`
}

// creditBlockListInvoicesResponseBlockCreditAllocationFilterJSON contains the JSON
// metadata for the struct
// [CreditBlockListInvoicesResponseBlockCreditAllocationFilter]
type creditBlockListInvoicesResponseBlockCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponseBlockCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseBlockCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField string

const (
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPriceID       CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField = "price_id"
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldItemID        CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField = "item_id"
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPriceType     CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField = "price_type"
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldCurrency      CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField = "currency"
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPricingUnitID CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CreditBlockListInvoicesResponseBlockCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPriceID, CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldItemID, CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPriceType, CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldCurrency, CreditBlockListInvoicesResponseBlockCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperator string

const (
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperatorIncludes CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperator = "includes"
	CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperatorExcludes CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperator = "excludes"
)

func (r CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperatorIncludes, CreditBlockListInvoicesResponseBlockCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

type CreditBlockListInvoicesResponseInvoice struct {
	ID            string                                        `json:"id" api:"required"`
	Customer      shared.CustomerMinified                       `json:"customer" api:"required"`
	InvoiceNumber string                                        `json:"invoice_number" api:"required"`
	Status        CreditBlockListInvoicesResponseInvoicesStatus `json:"status" api:"required"`
	Subscription  shared.SubscriptionMinified                   `json:"subscription" api:"required,nullable"`
	JSON          creditBlockListInvoicesResponseInvoiceJSON    `json:"-"`
}

// creditBlockListInvoicesResponseInvoiceJSON contains the JSON metadata for the
// struct [CreditBlockListInvoicesResponseInvoice]
type creditBlockListInvoicesResponseInvoiceJSON struct {
	ID            apijson.Field
	Customer      apijson.Field
	InvoiceNumber apijson.Field
	Status        apijson.Field
	Subscription  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CreditBlockListInvoicesResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditBlockListInvoicesResponseInvoiceJSON) RawJSON() string {
	return r.raw
}

type CreditBlockListInvoicesResponseInvoicesStatus string

const (
	CreditBlockListInvoicesResponseInvoicesStatusIssued CreditBlockListInvoicesResponseInvoicesStatus = "issued"
	CreditBlockListInvoicesResponseInvoicesStatusPaid   CreditBlockListInvoicesResponseInvoicesStatus = "paid"
	CreditBlockListInvoicesResponseInvoicesStatusSynced CreditBlockListInvoicesResponseInvoicesStatus = "synced"
	CreditBlockListInvoicesResponseInvoicesStatusVoid   CreditBlockListInvoicesResponseInvoicesStatus = "void"
	CreditBlockListInvoicesResponseInvoicesStatusDraft  CreditBlockListInvoicesResponseInvoicesStatus = "draft"
)

func (r CreditBlockListInvoicesResponseInvoicesStatus) IsKnown() bool {
	switch r {
	case CreditBlockListInvoicesResponseInvoicesStatusIssued, CreditBlockListInvoicesResponseInvoicesStatusPaid, CreditBlockListInvoicesResponseInvoicesStatusSynced, CreditBlockListInvoicesResponseInvoicesStatusVoid, CreditBlockListInvoicesResponseInvoicesStatusDraft:
		return true
	}
	return false
}
