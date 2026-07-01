// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
// credits within Orb.
//
// CustomerCreditService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerCreditService] method instead.
type CustomerCreditService struct {
	Options []option.RequestOption
	// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
	// credits within Orb.
	Ledger *CustomerCreditLedgerService
	// The [Credit Ledger Entry resource](/product-catalog/prepurchase) models prepaid
	// credits within Orb.
	TopUps *CustomerCreditTopUpService
}

// NewCustomerCreditService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCreditService(opts ...option.RequestOption) (r *CustomerCreditService) {
	r = &CustomerCreditService{}
	r.Options = opts
	r.Ledger = NewCustomerCreditLedgerService(opts...)
	r.TopUps = NewCustomerCreditTopUpService(opts...)
	return
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
//
// Results can be filtered by the block's `effective_date` using the
// `effective_date[gte]`, `effective_date[gt]`, `effective_date[lt]`, and
// `effective_date[lte]` query parameters. This filters on when the credit block
// becomes effective, which may differ from creation time for backdated credits.
func (r *CustomerCreditService) List(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/%s/credits", customerID)
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

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
//
// Results can be filtered by the block's `effective_date` using the
// `effective_date[gte]`, `effective_date[gt]`, `effective_date[lt]`, and
// `effective_date[lte]` query parameters. This filters on when the credit block
// becomes effective, which may differ from creation time for backdated credits.
func (r *CustomerCreditService) ListAutoPaging(ctx context.Context, customerID string, query CustomerCreditListParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditListResponse] {
	return pagination.NewPageAutoPager(r.List(ctx, customerID, query, opts...))
}

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
//
// Results can be filtered by the block's `effective_date` using the
// `effective_date[gte]`, `effective_date[gt]`, `effective_date[lt]`, and
// `effective_date[lte]` query parameters. This filters on when the credit block
// becomes effective, which may differ from creation time for backdated credits.
func (r *CustomerCreditService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) (res *pagination.Page[CustomerCreditListByExternalIDResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("customers/external_customer_id/%s/credits", externalCustomerID)
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

// Returns a paginated list of unexpired, non-zero credit blocks for a customer.
//
// If `include_all_blocks` is set to `true`, all credit blocks (including expired
// and depleted blocks) will be included in the response.
//
// Note that `currency` defaults to credits if not specified. To use a real world
// currency, set `currency` to an ISO 4217 string.
//
// Results can be filtered by the block's `effective_date` using the
// `effective_date[gte]`, `effective_date[gt]`, `effective_date[lt]`, and
// `effective_date[lte]` query parameters. This filters on when the credit block
// becomes effective, which may differ from creation time for backdated credits.
func (r *CustomerCreditService) ListByExternalIDAutoPaging(ctx context.Context, externalCustomerID string, query CustomerCreditListByExternalIDParams, opts ...option.RequestOption) *pagination.PageAutoPager[CustomerCreditListByExternalIDResponse] {
	return pagination.NewPageAutoPager(r.ListByExternalID(ctx, externalCustomerID, query, opts...))
}

type CustomerCreditListResponse struct {
	ID      string  `json:"id" api:"required"`
	Balance float64 `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CustomerCreditListResponseCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                   `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                   `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CustomerCreditListResponseFilter          `json:"filters" api:"required"`
	MaximumInitialBalance float64                                     `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                `json:"metadata" api:"required"`
	PerUnitCostBasis string                           `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CustomerCreditListResponseStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CustomerCreditListResponseCreditAllocation `json:"credit_allocation" api:"nullable"`
	JSON             customerCreditListResponseJSON             `json:"-"`
}

// customerCreditListResponseJSON contains the JSON metadata for the struct
// [CustomerCreditListResponse]
type customerCreditListResponseJSON struct {
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

func (r *CustomerCreditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CustomerCreditListResponseCreditBlockSource string

const (
	CustomerCreditListResponseCreditBlockSourceAllocation CustomerCreditListResponseCreditBlockSource = "allocation"
	CustomerCreditListResponseCreditBlockSourceTopUp      CustomerCreditListResponseCreditBlockSource = "top_up"
	CustomerCreditListResponseCreditBlockSourceManual     CustomerCreditListResponseCreditBlockSource = "manual"
)

func (r CustomerCreditListResponseCreditBlockSource) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseCreditBlockSourceAllocation, CustomerCreditListResponseCreditBlockSourceTopUp, CustomerCreditListResponseCreditBlockSourceManual:
		return true
	}
	return false
}

// A PriceFilter that only allows item_id field for block filters.
type CustomerCreditListResponseFilter struct {
	// The property of the price the block applies to. Only item_id is supported.
	Field CustomerCreditListResponseFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditListResponseFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                             `json:"values" api:"required"`
	JSON   customerCreditListResponseFilterJSON `json:"-"`
}

// customerCreditListResponseFilterJSON contains the JSON metadata for the struct
// [CustomerCreditListResponseFilter]
type customerCreditListResponseFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price the block applies to. Only item_id is supported.
type CustomerCreditListResponseFiltersField string

const (
	CustomerCreditListResponseFiltersFieldItemID CustomerCreditListResponseFiltersField = "item_id"
)

func (r CustomerCreditListResponseFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseFiltersFieldItemID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditListResponseFiltersOperator string

const (
	CustomerCreditListResponseFiltersOperatorIncludes CustomerCreditListResponseFiltersOperator = "includes"
	CustomerCreditListResponseFiltersOperatorExcludes CustomerCreditListResponseFiltersOperator = "excludes"
)

func (r CustomerCreditListResponseFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseFiltersOperatorIncludes, CustomerCreditListResponseFiltersOperatorExcludes:
		return true
	}
	return false
}

type CustomerCreditListResponseStatus string

const (
	CustomerCreditListResponseStatusActive         CustomerCreditListResponseStatus = "active"
	CustomerCreditListResponseStatusPendingPayment CustomerCreditListResponseStatus = "pending_payment"
)

func (r CustomerCreditListResponseStatus) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseStatusActive, CustomerCreditListResponseStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CustomerCreditListResponseCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                             `json:"item_id" api:"required"`
	Filters       []CustomerCreditListResponseCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                             `json:"license_type_id" api:"nullable"`
	JSON          customerCreditListResponseCreditAllocationJSON     `json:"-"`
}

// customerCreditListResponseCreditAllocationJSON contains the JSON metadata for
// the struct [CustomerCreditListResponseCreditAllocation]
type customerCreditListResponseCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditListResponseCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListResponseCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditListResponseCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditListResponseCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                             `json:"values" api:"required"`
	JSON   customerCreditListResponseCreditAllocationFilterJSON `json:"-"`
}

// customerCreditListResponseCreditAllocationFilterJSON contains the JSON metadata
// for the struct [CustomerCreditListResponseCreditAllocationFilter]
type customerCreditListResponseCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListResponseCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListResponseCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditListResponseCreditAllocationFiltersField string

const (
	CustomerCreditListResponseCreditAllocationFiltersFieldPriceID       CustomerCreditListResponseCreditAllocationFiltersField = "price_id"
	CustomerCreditListResponseCreditAllocationFiltersFieldItemID        CustomerCreditListResponseCreditAllocationFiltersField = "item_id"
	CustomerCreditListResponseCreditAllocationFiltersFieldPriceType     CustomerCreditListResponseCreditAllocationFiltersField = "price_type"
	CustomerCreditListResponseCreditAllocationFiltersFieldCurrency      CustomerCreditListResponseCreditAllocationFiltersField = "currency"
	CustomerCreditListResponseCreditAllocationFiltersFieldPricingUnitID CustomerCreditListResponseCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CustomerCreditListResponseCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseCreditAllocationFiltersFieldPriceID, CustomerCreditListResponseCreditAllocationFiltersFieldItemID, CustomerCreditListResponseCreditAllocationFiltersFieldPriceType, CustomerCreditListResponseCreditAllocationFiltersFieldCurrency, CustomerCreditListResponseCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditListResponseCreditAllocationFiltersOperator string

const (
	CustomerCreditListResponseCreditAllocationFiltersOperatorIncludes CustomerCreditListResponseCreditAllocationFiltersOperator = "includes"
	CustomerCreditListResponseCreditAllocationFiltersOperatorExcludes CustomerCreditListResponseCreditAllocationFiltersOperator = "excludes"
)

func (r CustomerCreditListResponseCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditListResponseCreditAllocationFiltersOperatorIncludes, CustomerCreditListResponseCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

type CustomerCreditListByExternalIDResponse struct {
	ID      string  `json:"id" api:"required"`
	Balance float64 `json:"balance" api:"required"`
	// How this credit block was created: `allocation` (a subscription's recurring
	// credit allocation), `top_up` (an automatic balance-threshold top-up), or
	// `manual` (a manual credit ledger increment, including credits voided or expired
	// off another block).
	CreditBlockSource     CustomerCreditListByExternalIDResponseCreditBlockSource `json:"credit_block_source" api:"required"`
	EffectiveDate         time.Time                                               `json:"effective_date" api:"required,nullable" format:"date-time"`
	ExpiryDate            time.Time                                               `json:"expiry_date" api:"required,nullable" format:"date-time"`
	Filters               []CustomerCreditListByExternalIDResponseFilter          `json:"filters" api:"required"`
	MaximumInitialBalance float64                                                 `json:"maximum_initial_balance" api:"required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata         map[string]string                            `json:"metadata" api:"required"`
	PerUnitCostBasis string                                       `json:"per_unit_cost_basis" api:"required,nullable"`
	Status           CustomerCreditListByExternalIDResponseStatus `json:"status" api:"required"`
	// The credit allocation that funded a block. Extends the allocation resource
	// serialized on prices with the catalog-item attribution of the funding price.
	CreditAllocation CustomerCreditListByExternalIDResponseCreditAllocation `json:"credit_allocation" api:"nullable"`
	JSON             customerCreditListByExternalIDResponseJSON             `json:"-"`
}

// customerCreditListByExternalIDResponseJSON contains the JSON metadata for the
// struct [CustomerCreditListByExternalIDResponse]
type customerCreditListByExternalIDResponseJSON struct {
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

func (r *CustomerCreditListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

// How this credit block was created: `allocation` (a subscription's recurring
// credit allocation), `top_up` (an automatic balance-threshold top-up), or
// `manual` (a manual credit ledger increment, including credits voided or expired
// off another block).
type CustomerCreditListByExternalIDResponseCreditBlockSource string

const (
	CustomerCreditListByExternalIDResponseCreditBlockSourceAllocation CustomerCreditListByExternalIDResponseCreditBlockSource = "allocation"
	CustomerCreditListByExternalIDResponseCreditBlockSourceTopUp      CustomerCreditListByExternalIDResponseCreditBlockSource = "top_up"
	CustomerCreditListByExternalIDResponseCreditBlockSourceManual     CustomerCreditListByExternalIDResponseCreditBlockSource = "manual"
)

func (r CustomerCreditListByExternalIDResponseCreditBlockSource) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseCreditBlockSourceAllocation, CustomerCreditListByExternalIDResponseCreditBlockSourceTopUp, CustomerCreditListByExternalIDResponseCreditBlockSourceManual:
		return true
	}
	return false
}

// A PriceFilter that only allows item_id field for block filters.
type CustomerCreditListByExternalIDResponseFilter struct {
	// The property of the price the block applies to. Only item_id is supported.
	Field CustomerCreditListByExternalIDResponseFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditListByExternalIDResponseFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                         `json:"values" api:"required"`
	JSON   customerCreditListByExternalIDResponseFilterJSON `json:"-"`
}

// customerCreditListByExternalIDResponseFilterJSON contains the JSON metadata for
// the struct [CustomerCreditListByExternalIDResponseFilter]
type customerCreditListByExternalIDResponseFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListByExternalIDResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListByExternalIDResponseFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price the block applies to. Only item_id is supported.
type CustomerCreditListByExternalIDResponseFiltersField string

const (
	CustomerCreditListByExternalIDResponseFiltersFieldItemID CustomerCreditListByExternalIDResponseFiltersField = "item_id"
)

func (r CustomerCreditListByExternalIDResponseFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseFiltersFieldItemID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditListByExternalIDResponseFiltersOperator string

const (
	CustomerCreditListByExternalIDResponseFiltersOperatorIncludes CustomerCreditListByExternalIDResponseFiltersOperator = "includes"
	CustomerCreditListByExternalIDResponseFiltersOperatorExcludes CustomerCreditListByExternalIDResponseFiltersOperator = "excludes"
)

func (r CustomerCreditListByExternalIDResponseFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseFiltersOperatorIncludes, CustomerCreditListByExternalIDResponseFiltersOperatorExcludes:
		return true
	}
	return false
}

type CustomerCreditListByExternalIDResponseStatus string

const (
	CustomerCreditListByExternalIDResponseStatusActive         CustomerCreditListByExternalIDResponseStatus = "active"
	CustomerCreditListByExternalIDResponseStatusPendingPayment CustomerCreditListByExternalIDResponseStatus = "pending_payment"
)

func (r CustomerCreditListByExternalIDResponseStatus) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseStatusActive, CustomerCreditListByExternalIDResponseStatusPendingPayment:
		return true
	}
	return false
}

// The credit allocation that funded a block. Extends the allocation resource
// serialized on prices with the catalog-item attribution of the funding price.
type CustomerCreditListByExternalIDResponseCreditAllocation struct {
	AllowsRollover   bool                    `json:"allows_rollover" api:"required"`
	Currency         string                  `json:"currency" api:"required"`
	CustomExpiration shared.CustomExpiration `json:"custom_expiration" api:"required,nullable"`
	// The ID of the catalog item this block was allocated from, derived from the
	// allocation's price.
	ItemID        string                                                         `json:"item_id" api:"required"`
	Filters       []CustomerCreditListByExternalIDResponseCreditAllocationFilter `json:"filters"`
	LicenseTypeID string                                                         `json:"license_type_id" api:"nullable"`
	JSON          customerCreditListByExternalIDResponseCreditAllocationJSON     `json:"-"`
}

// customerCreditListByExternalIDResponseCreditAllocationJSON contains the JSON
// metadata for the struct [CustomerCreditListByExternalIDResponseCreditAllocation]
type customerCreditListByExternalIDResponseCreditAllocationJSON struct {
	AllowsRollover   apijson.Field
	Currency         apijson.Field
	CustomExpiration apijson.Field
	ItemID           apijson.Field
	Filters          apijson.Field
	LicenseTypeID    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomerCreditListByExternalIDResponseCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListByExternalIDResponseCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type CustomerCreditListByExternalIDResponseCreditAllocationFilter struct {
	// The property of the price to filter on.
	Field CustomerCreditListByExternalIDResponseCreditAllocationFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                         `json:"values" api:"required"`
	JSON   customerCreditListByExternalIDResponseCreditAllocationFilterJSON `json:"-"`
}

// customerCreditListByExternalIDResponseCreditAllocationFilterJSON contains the
// JSON metadata for the struct
// [CustomerCreditListByExternalIDResponseCreditAllocationFilter]
type customerCreditListByExternalIDResponseCreditAllocationFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCreditListByExternalIDResponseCreditAllocationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCreditListByExternalIDResponseCreditAllocationFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type CustomerCreditListByExternalIDResponseCreditAllocationFiltersField string

const (
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPriceID       CustomerCreditListByExternalIDResponseCreditAllocationFiltersField = "price_id"
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldItemID        CustomerCreditListByExternalIDResponseCreditAllocationFiltersField = "item_id"
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPriceType     CustomerCreditListByExternalIDResponseCreditAllocationFiltersField = "price_type"
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldCurrency      CustomerCreditListByExternalIDResponseCreditAllocationFiltersField = "currency"
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPricingUnitID CustomerCreditListByExternalIDResponseCreditAllocationFiltersField = "pricing_unit_id"
)

func (r CustomerCreditListByExternalIDResponseCreditAllocationFiltersField) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPriceID, CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldItemID, CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPriceType, CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldCurrency, CustomerCreditListByExternalIDResponseCreditAllocationFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperator string

const (
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperatorIncludes CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperator = "includes"
	CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperatorExcludes CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperator = "excludes"
)

func (r CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperator) IsKnown() bool {
	switch r {
	case CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperatorIncludes, CustomerCreditListByExternalIDResponseCreditAllocationFiltersOperatorExcludes:
		return true
	}
	return false
}

type CustomerCreditListParams struct {
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor           param.Field[string]    `query:"cursor"`
	EffectiveDateGt  param.Field[time.Time] `query:"effective_date[gt]" format:"date-time"`
	EffectiveDateGte param.Field[time.Time] `query:"effective_date[gte]" format:"date-time"`
	EffectiveDateLt  param.Field[time.Time] `query:"effective_date[lt]" format:"date-time"`
	EffectiveDateLte param.Field[time.Time] `query:"effective_date[lte]" format:"date-time"`
	// If set to True, all expired and depleted blocks, as well as active block will be
	// returned.
	IncludeAllBlocks param.Field[bool] `query:"include_all_blocks"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerCreditListByExternalIDParams struct {
	// The ledger currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor           param.Field[string]    `query:"cursor"`
	EffectiveDateGt  param.Field[time.Time] `query:"effective_date[gt]" format:"date-time"`
	EffectiveDateGte param.Field[time.Time] `query:"effective_date[gte]" format:"date-time"`
	EffectiveDateLt  param.Field[time.Time] `query:"effective_date[lt]" format:"date-time"`
	EffectiveDateLte param.Field[time.Time] `query:"effective_date[lte]" format:"date-time"`
	// If set to True, all expired and depleted blocks, as well as active block will be
	// returned.
	IncludeAllBlocks param.Field[bool] `query:"include_all_blocks"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [CustomerCreditListByExternalIDParams]'s query parameters as
// `url.Values`.
func (r CustomerCreditListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
