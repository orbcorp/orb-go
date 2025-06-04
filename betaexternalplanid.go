// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// BetaExternalPlanIDService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaExternalPlanIDService] method instead.
type BetaExternalPlanIDService struct {
	Options []option.RequestOption
}

// NewBetaExternalPlanIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBetaExternalPlanIDService(opts ...option.RequestOption) (r *BetaExternalPlanIDService) {
	r = &BetaExternalPlanIDService{}
	r.Options = opts
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows the creation of a new plan version for an existing plan.
func (r *BetaExternalPlanIDService) NewPlanVersion(ctx context.Context, externalPlanID string, body BetaExternalPlanIDNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *BetaExternalPlanIDService) FetchPlanVersion(ctx context.Context, externalPlanID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions/%s", externalPlanID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows setting the default version of a plan.
func (r *BetaExternalPlanIDService) SetDefaultPlanVersion(ctx context.Context, externalPlanID string, body BetaExternalPlanIDSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/set_default_version", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BetaExternalPlanIDNewPlanVersionParams struct {
	// New version number.
	Version param.Field[int64] `json:"version,required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]BetaExternalPlanIDNewPlanVersionParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]BetaExternalPlanIDNewPlanVersionParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r BetaExternalPlanIDNewPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                       `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                                `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                                `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string]      `json:"currency"`
	Filters  param.Field[interface{}] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType     param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                                 `json:"usage_discount"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment].
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                                           `json:"percentage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                                      `json:"usage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                                        `json:"amount_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimum) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMinimumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                                 `json:"maximum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximum) implementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewMaximumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPrice] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// The custom expiration for the allocation.
	CustomExpiration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration] `json:"custom_expiration"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period. Set to null if using custom_expiration.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The custom expiration for the allocation.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration struct {
	Duration     param.Field[int64]                                                                                      `json:"duration,required"`
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesAllocationPriceCustomExpirationDurationUnitMonth:
		return true
	}
	return false
}

// The price to add to the plan
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance           param.Field[bool]        `json:"billed_in_advance"`
	BillingCycleConfiguration param.Field[interface{}] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[interface{}] `json:"bps_config"`
	BulkBpsConfig             param.Field[interface{}] `json:"bulk_bps_config"`
	BulkConfig                param.Field[interface{}] `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}] `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion interface {
	implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                         `json:"name,required"`
	UnitConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                               `json:"name,required"`
	PackageConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelTypePackage BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType = "package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                             `json:"item_id,required"`
	MatrixConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                             `json:"name,required"`
	TieredConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                                   `json:"name,required"`
	TieredBpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelTypeBps BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                 `json:"name,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                 `json:"name,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                            `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                              `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                           `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                             `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                         `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                                         `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                                      `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                        `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                         `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                                         `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                   `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                                       `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                           param.Field[string]                 `json:"name,required"`
	TieredPackageWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_package_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                                         `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                               `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPrice) implementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit                            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered                          BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredBps                       BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_bps"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBps                             BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bps"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkBps                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_bps"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk                            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredBps, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBps, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkBps, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id,required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                           `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                                    `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                                    `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string]      `json:"currency"`
	Filters  param.Field[interface{}] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID             param.Field[string]  `json:"item_id"`
	MaximumAmount      param.Field[string]  `json:"maximum_amount"`
	MinimumAmount      param.Field[string]  `json:"minimum_amount"`
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType     param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                                     `json:"usage_discount"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType     param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                                               `json:"percentage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType] `json:"adjustment_type,required"`
	UsageDiscount  param.Field[float64]                                                                                          `json:"usage_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType = "usage_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewUsageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                                            `json:"amount_discount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewAmountDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimum) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMinimumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	MaximumAmount  param.Field[string]                                                                                     `json:"maximum_amount,required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType] `json:"price_type"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximum) implementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumAppliesToAllTrue:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter struct {
	// The property of the price to filter on.
	Field param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField] `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator] `json:"operator,required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID       BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "price_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldItemID        BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "item_id"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "price_type"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "currency"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField = "pricing_unit_id"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersField) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldItemID, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPriceType, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldCurrency, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator = "includes"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator = "excludes"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperator) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorIncludes, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewMaximumPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPrice] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The allocation price to add to the plan.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// The custom expiration for the allocation.
	CustomExpiration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration] `json:"custom_expiration"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period. Set to null if using custom_expiration.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The custom expiration for the allocation.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration struct {
	Duration     param.Field[int64]                                                                                          `json:"duration,required"`
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesAllocationPriceCustomExpirationDurationUnitMonth:
		return true
	}
	return false
}

// The price to add to the plan
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance           param.Field[bool]        `json:"billed_in_advance"`
	BillingCycleConfiguration param.Field[interface{}] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[interface{}] `json:"bps_config"`
	BulkBpsConfig             param.Field[interface{}] `json:"bulk_bps_config"`
	BulkConfig                param.Field[interface{}] `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}] `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                      param.Field[string]      `json:"currency"`
	DimensionalPriceConfiguration param.Field[interface{}] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig       param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey                    param.Field[string]      `json:"invoice_grouping_key"`
	InvoicingCycleConfiguration           param.Field[interface{}] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[interface{}] `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[interface{}] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}] `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}] `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}] `json:"metadata"`
	PackageConfig                         param.Field[interface{}] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}] `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[interface{}] `json:"tiered_bps_config"`
	TieredConfig                          param.Field[interface{}] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}] `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}] `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}] `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}] `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[interface{}] `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}] `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}] `json:"unit_with_proration_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion interface {
	implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                             `json:"name,required"`
	UnitConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType = "unit"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                                   `json:"name,required"`
	PackageConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType = "package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                                                                 `json:"item_id,required"`
	MatrixConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType = "matrix"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                                 `json:"name,required"`
	TieredConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier struct {
	// Exclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Inclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                                       `json:"name,required"`
	TieredBpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Exclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Inclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice struct {
	BpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType = "bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice struct {
	BulkConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType = "bulk"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                 `json:"name,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                             `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                 `json:"name,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                                       `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                                `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                  `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                               `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                             `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                                                             `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}]                                                                          `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                            `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                                                             `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                                                             `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMaxGroupTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithUnitPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                       `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                 `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[map[string]interface{}] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanScalableMatrixWithTieredPricingPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                                                           `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                             `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                           param.Field[string]                 `json:"name,required"`
	TieredPackageWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_package_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                                                                             `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence             param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence] `json:"cadence,required"`
	GroupedTieredConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPrice) implementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// For dimensional price: specifies a price group and dimension values
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration struct {
	// The list of dimension values matching (in order) the dimensions of the price
	// group
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// The id of the dimensional price group to include this price in
	DimensionalPriceGroupID param.Field[string] `json:"dimensional_price_group_id"`
	// The external id of the dimensional price group to include this price in
	ExternalDimensionalPriceGroupID param.Field[string] `json:"external_dimensional_price_group_id"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceDimensionalPriceConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredBps                       BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_bps"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBps                             BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bps"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkBps                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_bps"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredBps, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBps, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkBps, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version,required"`
}

func (r BetaExternalPlanIDSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
