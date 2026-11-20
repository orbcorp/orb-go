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
// PlanExternalPlanIDService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanExternalPlanIDService] method instead.
type PlanExternalPlanIDService struct {
	Options []option.RequestOption
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Draft *PlanExternalPlanIDDraftService
}

// NewPlanExternalPlanIDService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPlanExternalPlanIDService(opts ...option.RequestOption) (r *PlanExternalPlanIDService) {
	r = &PlanExternalPlanIDService{}
	r.Options = opts
	r.Draft = NewPlanExternalPlanIDDraftService(opts...)
	return
}

// This endpoint can be used to update the `name`, `external_plan_id`,
// `description`, and `metadata` of an existing plan. A plan's `name` cannot be
// empty.
//
// Other fields on a plan are currently immutable.
func (r *PlanExternalPlanIDService) Update(ctx context.Context, otherExternalPlanID string, body PlanExternalPlanIDUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if otherExternalPlanID == "" {
		err = errors.New("missing required other_external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s", otherExternalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// This endpoint allows the creation of a new plan version for an existing plan.
func (r *PlanExternalPlanIDService) NewPlanVersion(ctx context.Context, externalPlanID string, body PlanExternalPlanIDNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint is used to fetch [plan](/core-concepts##plan-and-price) details
// given an external_plan_id identifier. It returns information about the prices
// included in the plan and their configuration, as well as the product that the
// plan is attached to.
//
// If multiple plans are found to contain the specified external_plan_id, the
// active plans will take priority over archived ones, and among those, the
// endpoint will return the most recently created plan.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](/core-concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](/core-concepts#plan-and-price). "
func (r *PlanExternalPlanIDService) Fetch(ctx context.Context, externalPlanID string, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *PlanExternalPlanIDService) FetchPlanVersion(ctx context.Context, externalPlanID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions/%s", externalPlanID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows setting the default version of a plan.
func (r *PlanExternalPlanIDService) SetDefaultPlanVersion(ctx context.Context, externalPlanID string, body PlanExternalPlanIDSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/set_default_version", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PlanExternalPlanIDUpdateParams struct {
	// The Orb ID of the active billing entity to set as this plan's default entity.
	// Pass `null` to clear it.
	BillingEntityID param.Field[string] `json:"billing_entity_id"`
	// An optional user-defined description of the plan.
	Description param.Field[string] `json:"description"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The name of the plan.
	Name param.Field[string] `json:"name"`
}

func (r PlanExternalPlanIDUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParams struct {
	// New version number.
	Version param.Field[int64] `json:"version" api:"required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanExternalPlanIDNewPlanVersionParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanExternalPlanIDNewPlanVersionParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanExternalPlanIDNewPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                                       `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
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
	PriceType     param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                                             `json:"tiers"`
	UsageDiscount param.Field[float64]                                                                 `json:"usage_discount"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment].
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage          PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "usage"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed          PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears      PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed, PlanExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type" api:"required"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPrice].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                                 `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                               `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
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
	MatrixValues param.Field[[]PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                   `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                   `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit                            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered                          PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk                            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters                 PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_filters"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage                         PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredMatrix             PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_matrix"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredMatrixWithAllocation      PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_matrix_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts    PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_threshold_discounts"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_allocation"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeDailyCreditAllowance            PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "daily_credit_allowance"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMeteredAllowance                PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "metered_allowance"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite                PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "minimum_composite"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePercent                         PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "percent"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeEventOutput                     PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "event_output"
	PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithMultiplier              PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredMatrix, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredMatrixWithAllocation, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeDailyCreditAllowance, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMeteredAllowance, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePercent, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeEventOutput, PlanExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id" api:"required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id" api:"required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id" api:"required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                                           `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
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
	PriceType     param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                                                 `json:"tiers"`
	UsageDiscount param.Field[float64]                                                                     `json:"usage_discount"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, PlanExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id" api:"required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type" api:"required"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPrice].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                 `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                                     `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                   `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
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
	MatrixValues param.Field[[]PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                           `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
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

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceOneTime    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "one_time"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceMonthly    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "monthly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceQuarterly  PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "quarterly"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "semi_annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceAnnual     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "annual"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceCustom     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence = "custom"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceOneTime, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceMonthly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceQuarterly, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceAnnual, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType string

const (
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters                 PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_filters"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage                         PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredMatrix             PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_matrix"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation      PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_matrix_with_allocation"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts    PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_threshold_discounts"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_allocation"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeDailyCreditAllowance            PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "daily_credit_allowance"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMeteredAllowance                PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "metered_allowance"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite                PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "minimum_composite"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePercent                         PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "percent"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput                     PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "event_output"
	PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithMultiplier              PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_multiplier"
)

func (r PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredMatrix, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeDailyCreditAllowance, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMeteredAllowance, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePercent, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput, PlanExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanExternalPlanIDSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version" api:"required"`
}

func (r PlanExternalPlanIDSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
