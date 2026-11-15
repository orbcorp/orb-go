// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
)

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
//
// PlanExternalPlanIDDraftService contains methods and other services that help
// with interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanExternalPlanIDDraftService] method instead.
type PlanExternalPlanIDDraftService struct {
	Options []option.RequestOption
}

// NewPlanExternalPlanIDDraftService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPlanExternalPlanIDDraftService(opts ...option.RequestOption) (r *PlanExternalPlanIDDraftService) {
	r = &PlanExternalPlanIDDraftService{}
	r.Options = opts
	return
}

// This endpoint discards the plan's draft version. Prices and adjustments the
// draft created stay in the Pricebook.
func (r *PlanExternalPlanIDDraftService) Discard(ctx context.Context, externalPlanID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/draft", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint returns the plan's unpublished draft version, with its changes
// applied on top of the latest published version. It returns a 404 if the plan has
// no draft.
func (r *PlanExternalPlanIDDraftService) Fetch(ctx context.Context, externalPlanID string, opts ...option.RequestOption) (res *PlanVersionDraft, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/draft", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint publishes the plan's draft as the next numbered version. A draft
// that fails plan version validation is a 400 and is left as it was.
func (r *PlanExternalPlanIDDraftService) Publish(ctx context.Context, externalPlanID string, body PlanExternalPlanIDDraftPublishParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/draft/publish", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint creates the plan's draft version or fully replaces its changes.
// The changes are always relative to the latest published version. A save replaces
// the draft; last write wins.
func (r *PlanExternalPlanIDDraftService) Save(ctx context.Context, externalPlanID string, body PlanExternalPlanIDDraftSaveParams, opts ...option.RequestOption) (res *PlanVersionDraft, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/draft", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type PlanExternalPlanIDDraftPublishParams struct {
	// Also make the new version the plan's default.
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanExternalPlanIDDraftPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParams struct {
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanExternalPlanIDDraftSaveParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPrice] `json:"add_prices"`
	// Optional description of what changes in this draft
	Description param.Field[string] `json:"description"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanExternalPlanIDDraftSaveParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanExternalPlanIDDraftSaveParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanExternalPlanIDDraftSaveParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePrice] `json:"replace_prices"`
}

func (r PlanExternalPlanIDDraftSaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                                  `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                           `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                           `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                                        `json:"tiers"`
	UsageDiscount param.Field[float64]                                                            `json:"usage_discount"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustment) ImplementsPlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustment].
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAllTrue PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType string

const (
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeUsage          PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "usage"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixed          PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeInArrears      PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeUsage, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixed, PlanExternalPlanIDDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit pricing
	UnitConfig param.Field[shared.UnitConfigParam] `json:"unit_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelTypeUnit PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanExternalPlanIDDraftSaveParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// Configuration for bulk pricing
	BulkConfig              param.Field[shared.BulkConfigParam] `json:"bulk_config"`
	BulkWithFiltersConfig   param.Field[interface{}]            `json:"bulk_with_filters_config"`
	BulkWithProrationConfig param.Field[interface{}]            `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate                    param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig              param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedAllocationConfig param.Field[interface{}] `json:"cumulative_grouped_allocation_config"`
	CumulativeGroupedBulkConfig       param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                   param.Field[string]      `json:"currency"`
	DailyCreditAllowanceConfig param.Field[interface{}] `json:"daily_credit_allowance_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	EventOutputConfig             param.Field[interface{}]                                  `json:"event_output_config"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity                param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig           param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig               param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredMatrixConfig         param.Field[interface{}] `json:"grouped_tiered_matrix_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig         param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig        param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MatrixWithThresholdDiscountsConfig param.Field[interface{}]                            `json:"matrix_with_threshold_discounts_config"`
	MaxGroupTieredPackageConfig        param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                           param.Field[interface{}]                            `json:"metadata"`
	MeteredAllowanceConfig             param.Field[interface{}]                            `json:"metered_allowance_config"`
	MinimumCompositeConfig             param.Field[interface{}]                            `json:"minimum_composite_config"`
	// Configuration for package pricing
	PackageConfig               param.Field[shared.PackageConfigParam] `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}]               `json:"package_with_allocation_config"`
	PercentConfig               param.Field[interface{}]               `json:"percent_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]      `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	// Configuration for tiered pricing
	TieredConfig                     param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredMatrixWithAllocationConfig param.Field[interface{}]              `json:"tiered_matrix_with_allocation_config"`
	TieredPackageConfig              param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig   param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig          param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig        param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig               param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithMultiplierConfig param.Field[interface{}]            `json:"unit_with_multiplier_config"`
	UnitWithPercentConfig    param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig  param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPrice].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                             `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                 `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                        `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                            `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
	// Discount rate applied to spend above the threshold.
	AboveThresholdDiscountPercentage param.Field[string] `json:"above_threshold_discount_percentage" api:"required"`
	// Discount rate applied to spend at or below the threshold. Set to 0 for no
	// baseline discount.
	BelowThresholdDiscountPercentage param.Field[string] `json:"below_threshold_discount_percentage" api:"required"`
	// Semicolon-separated list of matrix cell coordinates targeted by this group. Each
	// coordinate is `first,second` when the matrix has two dimensions, or just `first`
	// for a single-dimension matrix. Example: `blue,circle;green,triangle`.
	CellCoordinates param.Field[string] `json:"cell_coordinates" api:"required"`
	ThresholdAmount param.Field[string] `json:"threshold_amount" api:"required"`
	Description     param.Field[string] `json:"description"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                 `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
	// Credits granted per day. Lose-it-or-use-it; does not roll over.
	DailyAllowance param.Field[string] `json:"daily_allowance" api:"required"`
	// Default per-unit credit rate for any usage not bucketed into a specified
	// matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Event property whose value identifies the day bucket the event belongs to (e.g.
	// 'event_day' set to an ISO date string in the customer's timezone). The allowance
	// resets per distinct value of this property.
	EventDayProperty param.Field[string] `json:"event_day_property" api:"required"`
	// Per-dimension credit rates
	MatrixValues param.Field[[]PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
	// The grouping_key value whose summed quantity represents the allowance for this
	// period (e.g. 'storage_snapshot' emitting 3 × avg storage). Capped at consumption
	// — credit can never exceed actual usage.
	AllowanceGroupingValue param.Field[string] `json:"allowance_grouping_value" api:"required"`
	// The grouping_key value whose summed quantity represents consumption (e.g.
	// 'download'). Charged at unit_amount.
	ConsumptionGroupingValue param.Field[string] `json:"consumption_grouping_value" api:"required"`
	// Event property used to partition the metric into consumption and allowance
	// quantities (e.g. 'event_name'). The metric is queried with this key and the two
	// values below select which partition is which.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Per-unit price applied to gross consumption and to the allowance credit.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// Sub-line label for the credit row (e.g. 'Up to 3x free egress').
	AllowanceDisplayName param.Field[string] `json:"allowance_display_name"`
	// Sub-line label for the gross consumption row (e.g. 'bytes gotten').
	ConsumptionDisplayName param.Field[string] `json:"consumption_display_name"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// Fraction of the component subtotals to charge (0 < percent <= 1).
	Percent param.Field[float64] `json:"percent" api:"required"`
	// Maximum amount to charge. If unset, the fee has no upper bound.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Minimum amount to charge. If unset, the fee is bounded below by 0.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// If true, the minimum_amount is prorated based on the service period. The
	// maximum_amount is an absolute cap (never prorated), and the percent applied to
	// upstream subtotals is never prorated either.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
	// The key in the event data to extract the unit rate from.
	UnitRatingKey param.Field[string] `json:"unit_rating_key" api:"required"`
	// If provided, this amount will be used as the unit rate when an event does not
	// have a value for the `unit_rating_key`. If not provided, events missing a unit
	// rate will be ignored.
	DefaultUnitRate param.Field[string] `json:"default_unit_rate"`
	// An optional key in the event data to group by (e.g., event ID). All events will
	// also be grouped by their unit rate.
	GroupingKey param.Field[string] `json:"grouping_key"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
	// Sub-line label for list-price events that don't set display_name_key. Unlabeled
	// events with another multiplier are labeled by their multiplier (e.g. '0.8×
	// rate').
	DefaultDisplayName param.Field[string] `json:"default_display_name" api:"required"`
	// Event property holding the sub-line label for the event's rate (e.g.
	// 'Off-peak').
	DisplayNameKey param.Field[string] `json:"display_name_key" api:"required"`
	// Event property holding the multiplier applied to unit_amount (e.g. 0.8 for a 20%
	// discount). Events without a number between 0 and 1000 with at most 12 decimal
	// places are charged at unit_amount.
	MultiplierKey param.Field[string] `json:"multiplier_key" api:"required"`
	// List price per unit, before the event's multiplier is applied.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnit                            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "unit"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTiered                          PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulk                            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "bulk"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulkWithFilters                 PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "bulk_with_filters"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePackage                         PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "package"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrix                          PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "matrix"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredPackage                   PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered_package"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_tiered"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTieredMatrix             PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_tiered_matrix"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredMatrixWithAllocation      PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered_matrix_with_allocation"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts    PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "matrix_with_threshold_discounts"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredWithProration             PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithProration               PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulkWithProration               PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedAllocation     PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "cumulative_grouped_allocation"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeDailyCreditAllowance            PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "daily_credit_allowance"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMeteredAllowance                PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "metered_allowance"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMinimumComposite                PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "minimum_composite"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePercent                         PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "percent"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeEventOutput                     PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "event_output"
	PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithMultiplier              PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnit, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTiered, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulk, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulkWithFilters, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePackage, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrix, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredPackage, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTiered, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTieredMatrix, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredMatrixWithAllocation, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeTieredWithProration, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithProration, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeBulkWithProration, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedAllocation, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeDailyCreditAllowance, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMeteredAllowance, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeMinimumComposite, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypePercent, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeEventOutput, PlanExternalPlanIDDraftSaveParamsAddPricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id" api:"required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDDraftSaveParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id" api:"required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDDraftSaveParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id" api:"required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                                      `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                               `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                               `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                                            `json:"tiers"`
	UsageDiscount param.Field[float64]                                                                `json:"usage_discount"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustment) ImplementsPlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustment].
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, PlanExternalPlanIDDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id" api:"required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit pricing
	UnitConfig param.Field[shared.UnitConfigParam] `json:"unit_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelTypeUnit PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                       `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// Configuration for bulk pricing
	BulkConfig              param.Field[shared.BulkConfigParam] `json:"bulk_config"`
	BulkWithFiltersConfig   param.Field[interface{}]            `json:"bulk_with_filters_config"`
	BulkWithProrationConfig param.Field[interface{}]            `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate                    param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig              param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedAllocationConfig param.Field[interface{}] `json:"cumulative_grouped_allocation_config"`
	CumulativeGroupedBulkConfig       param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                   param.Field[string]      `json:"currency"`
	DailyCreditAllowanceConfig param.Field[interface{}] `json:"daily_credit_allowance_config"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	EventOutputConfig             param.Field[interface{}]                                  `json:"event_output_config"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity                param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig           param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig               param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredMatrixConfig         param.Field[interface{}] `json:"grouped_tiered_matrix_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig         param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig        param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MatrixWithThresholdDiscountsConfig param.Field[interface{}]                            `json:"matrix_with_threshold_discounts_config"`
	MaxGroupTieredPackageConfig        param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                           param.Field[interface{}]                            `json:"metadata"`
	MeteredAllowanceConfig             param.Field[interface{}]                            `json:"metered_allowance_config"`
	MinimumCompositeConfig             param.Field[interface{}]                            `json:"minimum_composite_config"`
	// Configuration for package pricing
	PackageConfig               param.Field[shared.PackageConfigParam] `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}]               `json:"package_with_allocation_config"`
	PercentConfig               param.Field[interface{}]               `json:"percent_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]      `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}] `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}] `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}] `json:"threshold_total_amount_config"`
	// Configuration for tiered pricing
	TieredConfig                     param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredMatrixWithAllocationConfig param.Field[interface{}]              `json:"tiered_matrix_with_allocation_config"`
	TieredPackageConfig              param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig   param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig          param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig        param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig               param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithMultiplierConfig param.Field[interface{}]            `json:"unit_with_multiplier_config"`
	UnitWithPercentConfig    param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig  param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPrice].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                 `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                            `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                                `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
	// Discount rate applied to spend above the threshold.
	AboveThresholdDiscountPercentage param.Field[string] `json:"above_threshold_discount_percentage" api:"required"`
	// Discount rate applied to spend at or below the threshold. Set to 0 for no
	// baseline discount.
	BelowThresholdDiscountPercentage param.Field[string] `json:"below_threshold_discount_percentage" api:"required"`
	// Semicolon-separated list of matrix cell coordinates targeted by this group. Each
	// coordinate is `first,second` when the matrix has two dimensions, or just `first`
	// for a single-dimension matrix. Example: `blue,circle;green,triangle`.
	CellCoordinates param.Field[string] `json:"cell_coordinates" api:"required"`
	ThresholdAmount param.Field[string] `json:"threshold_amount" api:"required"`
	Description     param.Field[string] `json:"description"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
	// Credits granted per day. Lose-it-or-use-it; does not roll over.
	DailyAllowance param.Field[string] `json:"daily_allowance" api:"required"`
	// Default per-unit credit rate for any usage not bucketed into a specified
	// matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Event property whose value identifies the day bucket the event belongs to (e.g.
	// 'event_day' set to an ISO date string in the customer's timezone). The allowance
	// resets per distinct value of this property.
	EventDayProperty param.Field[string] `json:"event_day_property" api:"required"`
	// Per-dimension credit rates
	MatrixValues param.Field[[]PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
	// The grouping_key value whose summed quantity represents the allowance for this
	// period (e.g. 'storage_snapshot' emitting 3 × avg storage). Capped at consumption
	// — credit can never exceed actual usage.
	AllowanceGroupingValue param.Field[string] `json:"allowance_grouping_value" api:"required"`
	// The grouping_key value whose summed quantity represents consumption (e.g.
	// 'download'). Charged at unit_amount.
	ConsumptionGroupingValue param.Field[string] `json:"consumption_grouping_value" api:"required"`
	// Event property used to partition the metric into consumption and allowance
	// quantities (e.g. 'event_name'). The metric is queried with this key and the two
	// values below select which partition is which.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Per-unit price applied to gross consumption and to the allowance credit.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// Sub-line label for the credit row (e.g. 'Up to 3x free egress').
	AllowanceDisplayName param.Field[string] `json:"allowance_display_name"`
	// Sub-line label for the gross consumption row (e.g. 'bytes gotten').
	ConsumptionDisplayName param.Field[string] `json:"consumption_display_name"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// Fraction of the component subtotals to charge (0 < percent <= 1).
	Percent param.Field[float64] `json:"percent" api:"required"`
	// Maximum amount to charge. If unset, the fee has no upper bound.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Minimum amount to charge. If unset, the fee is bounded below by 0.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// If true, the minimum_amount is prorated based on the service period. The
	// maximum_amount is an absolute cap (never prorated), and the percent applied to
	// upstream subtotals is never prorated either.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
	// The key in the event data to extract the unit rate from.
	UnitRatingKey param.Field[string] `json:"unit_rating_key" api:"required"`
	// If provided, this amount will be used as the unit rate when an event does not
	// have a value for the `unit_rating_key`. If not provided, events missing a unit
	// rate will be ignored.
	DefaultUnitRate param.Field[string] `json:"default_unit_rate"`
	// An optional key in the event data to group by (e.g., event ID). All events will
	// also be grouped by their unit rate.
	GroupingKey param.Field[string] `json:"grouping_key"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                             `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price. On a usage price this
	// also marks the price as eligible to draw down from that license type's
	// allocation; a usage price created without it is billed normally. Usage prices
	// with a license type must use the `unit` model, and only draw down when their
	// currency matches the allocation's.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
	// Sub-line label for list-price events that don't set display_name_key. Unlabeled
	// events with another multiplier are labeled by their multiplier (e.g. '0.8×
	// rate').
	DefaultDisplayName param.Field[string] `json:"default_display_name" api:"required"`
	// Event property holding the sub-line label for the event's rate (e.g.
	// 'Off-peak').
	DisplayNameKey param.Field[string] `json:"display_name_key" api:"required"`
	// Event property holding the multiplier applied to unit_amount (e.g. 0.8 for a 20%
	// discount). Events without a number between 0 and 1000 with at most 12 decimal
	// places are charged at unit_amount.
	MultiplierKey param.Field[string] `json:"multiplier_key" api:"required"`
	// List price per unit, before the event's multiplier is applied.
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceOneTime    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "one_time"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceMonthly    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "monthly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceQuarterly  PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "quarterly"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceSemiAnnual PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "semi_annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceAnnual     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "annual"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceCustom     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceOneTime, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceMonthly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceQuarterly, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceSemiAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceAnnual, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType string

const (
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnit                            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "unit"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTiered                          PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulk                            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "bulk"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulkWithFilters                 PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "bulk_with_filters"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePackage                         PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "package"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrix                          PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "matrix"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredPackage                   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered_package"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredMatrix             PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered_matrix"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation      PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered_matrix_with_allocation"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts    PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "matrix_with_threshold_discounts"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredWithProration             PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithProration               PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulkWithProration               PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "cumulative_grouped_allocation"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeDailyCreditAllowance            PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "daily_credit_allowance"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMeteredAllowance                PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "metered_allowance"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMinimumComposite                PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "minimum_composite"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePercent                         PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "percent"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeEventOutput                     PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "event_output"
	PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithMultiplier              PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnit, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTiered, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulk, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulkWithFilters, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePackage, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrix, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredPackage, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTiered, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredMatrix, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeTieredWithProration, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithProration, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeBulkWithProration, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeDailyCreditAllowance, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMeteredAllowance, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeMinimumComposite, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypePercent, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeEventOutput, PlanExternalPlanIDDraftSaveParamsReplacePricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}
