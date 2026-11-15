// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
// subscribed to by a customer. Plans define the billing behavior of the
// subscription. You can see more about how to configure prices in the
// [Price resource](/reference/price).
//
// PlanDraftService contains methods and other services that help with interacting
// with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanDraftService] method instead.
type PlanDraftService struct {
	Options []option.RequestOption
}

// NewPlanDraftService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanDraftService(opts ...option.RequestOption) (r *PlanDraftService) {
	r = &PlanDraftService{}
	r.Options = opts
	return
}

// This endpoint discards the plan's draft version. Prices and adjustments the
// draft created stay in the Pricebook.
func (r *PlanDraftService) Discard(ctx context.Context, planID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return err
	}
	path := fmt.Sprintf("plans/%s/draft", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint returns the plan's unpublished draft version, with its changes
// applied on top of the latest published version. It returns a 404 if the plan has
// no draft.
func (r *PlanDraftService) Fetch(ctx context.Context, planID string, opts ...option.RequestOption) (res *PlanVersionDraft, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/draft", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint publishes the plan's draft as the next numbered version. A draft
// that fails plan version validation is a 400 and is left as it was.
func (r *PlanDraftService) Publish(ctx context.Context, planID string, body PlanDraftPublishParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/draft/publish", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint creates the plan's draft version or fully replaces its changes.
// The changes are always relative to the latest published version. A save replaces
// the draft; last write wins.
func (r *PlanDraftService) Save(ctx context.Context, planID string, body PlanDraftSaveParams, opts ...option.RequestOption) (res *PlanVersionDraft, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/draft", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// The PlanVersionDraft resource represents a plan's unpublished next version: the
// prices and adjustments it will have once published, built on top of
// `base_version`.
type PlanVersionDraft struct {
	// Adjustments on the draft. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionDraftAdjustment `json:"adjustments" api:"required"`
	// The published version the draft is built on, the latest at read time.
	BaseVersion int64     `json:"base_version" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	// Optional description of what changes in this draft
	Description string             `json:"description" api:"required,nullable"`
	PlanID      string             `json:"plan_id" api:"required"`
	PlanPhases  []PlanVersionPhase `json:"plan_phases" api:"required,nullable"`
	// Prices on the draft. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices []shared.Price       `json:"prices" api:"required"`
	JSON   planVersionDraftJSON `json:"-"`
}

// planVersionDraftJSON contains the JSON metadata for the struct
// [PlanVersionDraft]
type planVersionDraftJSON struct {
	Adjustments apijson.Field
	BaseVersion apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	PlanID      apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionDraft) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionDraftJSON) RawJSON() string {
	return r.raw
}

type PlanVersionDraftAdjustment struct {
	ID             string                                    `json:"id" api:"required"`
	AdjustmentType PlanVersionDraftAdjustmentsAdjustmentType `json:"adjustment_type" api:"required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids" api:"required"`
	// This field can have the runtime type of
	// [[]shared.PlanPhaseUsageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseAmountDiscountAdjustmentFilter],
	// [[]shared.PlanPhasePercentageDiscountAdjustmentFilter],
	// [[]PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseMinimumAdjustmentFilter],
	// [[]shared.PlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters" api:"required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level" api:"required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order" api:"required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason" api:"required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id" api:"required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// This field can have the runtime type of
	// [[]PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier].
	Tiers interface{} `json:"tiers"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                        `json:"usage_discount"`
	JSON          planVersionDraftAdjustmentJSON `json:"-"`
	union         PlanVersionDraftAdjustmentsUnion
}

// planVersionDraftAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionDraftAdjustment]
type planVersionDraftAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	AmountDiscount       apijson.Field
	ItemID               apijson.Field
	MaximumAmount        apijson.Field
	MinimumAmount        apijson.Field
	PercentageDiscount   apijson.Field
	Tiers                apijson.Field
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r planVersionDraftAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionDraftAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionDraftAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionDraftAdjustmentsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment], [shared.PlanPhaseMaximumAdjustment].
func (r PlanVersionDraftAdjustment) AsUnion() PlanVersionDraftAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment] or [shared.PlanPhaseMaximumAdjustment].
type PlanVersionDraftAdjustmentsUnion interface {
	ImplementsPlanVersionDraftAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionDraftAdjustmentsUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment{}),
			DiscriminatorValue: "tiered_percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment struct {
	ID             string                                                                               `json:"id" api:"required"`
	AdjustmentType PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type" api:"required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids" api:"required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter `json:"filters" api:"required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level" api:"required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order" api:"required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason" api:"required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id" api:"required,nullable"`
	// The ordered, contiguous bands of cumulative eligible spend, each discounted at
	// its own percentage (progressive fill-a-tier), applied to the prices this
	// adjustment covers in a given billing period.
	Tiers []PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier `json:"tiers" api:"required"`
	JSON  planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON   `json:"-"`
}

// planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment]
type planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON struct {
	ID                   apijson.Field
	AdjustmentType       apijson.Field
	AppliesToPriceIDs    apijson.Field
	Filters              apijson.Field
	IsInvoiceLevel       apijson.Field
	PlanPhaseOrder       apijson.Field
	Reason               apijson.Field
	ReplacesAdjustmentID apijson.Field
	Tiers                apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustment) ImplementsPlanVersionDraftAdjustment() {
}

type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType string

const (
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values []string                                                                         `json:"values" api:"required"`
	JSON   planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter]
type planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField string

const (
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID       PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_id"
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID        PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "item_id"
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType     PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "price_type"
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency      PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "currency"
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceID, PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldItemID, PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPriceType, PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldCurrency, PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator string

const (
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "includes"
	PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorIncludes, PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

// One band of a tiered percentage discount. Bounds are denominated in the
// discount's currency. `lower_bound` is the exclusive start of the band and
// `upper_bound` is the inclusive end; `upper_bound` is null only for the
// open-ended final tier.
type PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound float64 `json:"lower_bound" api:"required"`
	// The percentage (between 0 and 1) discounted from spend that falls within this
	// tier.
	Percentage float64 `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend for this tier; null for the final
	// open-ended tier.
	UpperBound float64                                                                        `json:"upper_bound" api:"nullable"`
	JSON       planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON `json:"-"`
}

// planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON
// contains the JSON metadata for the struct
// [PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier]
type planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON struct {
	LowerBound  apijson.Field
	Percentage  apijson.Field
	UpperBound  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionDraftAdjustmentsPlanPhaseTieredPercentageDiscountAdjustmentTierJSON) RawJSON() string {
	return r.raw
}

type PlanVersionDraftAdjustmentsAdjustmentType string

const (
	PlanVersionDraftAdjustmentsAdjustmentTypeUsageDiscount            PlanVersionDraftAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionDraftAdjustmentsAdjustmentTypeAmountDiscount           PlanVersionDraftAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionDraftAdjustmentsAdjustmentTypePercentageDiscount       PlanVersionDraftAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionDraftAdjustmentsAdjustmentTypeTieredPercentageDiscount PlanVersionDraftAdjustmentsAdjustmentType = "tiered_percentage_discount"
	PlanVersionDraftAdjustmentsAdjustmentTypeMinimum                  PlanVersionDraftAdjustmentsAdjustmentType = "minimum"
	PlanVersionDraftAdjustmentsAdjustmentTypeMaximum                  PlanVersionDraftAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionDraftAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionDraftAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionDraftAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionDraftAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionDraftAdjustmentsAdjustmentTypeTieredPercentageDiscount, PlanVersionDraftAdjustmentsAdjustmentTypeMinimum, PlanVersionDraftAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanDraftPublishParams struct {
	// Also make the new version the plan's default.
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r PlanDraftPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParams struct {
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]PlanDraftSaveParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]PlanDraftSaveParamsAddPrice] `json:"add_prices"`
	// Optional description of what changes in this draft
	Description param.Field[string] `json:"description"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]PlanDraftSaveParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]PlanDraftSaveParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]PlanDraftSaveParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]PlanDraftSaveParamsReplacePrice] `json:"replace_prices"`
}

func (r PlanDraftSaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanDraftSaveParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanDraftSaveParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                    `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                             `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                             `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                          `json:"tiers"`
	UsageDiscount param.Field[float64]                                              `json:"usage_discount"`
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustment) ImplementsPlanDraftSaveParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanDraftSaveParamsAddAdjustmentsAdjustment].
type PlanDraftSaveParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsPlanDraftSaveParamsAddAdjustmentsAdjustmentUnion()
}

type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanDraftSaveParamsAddAdjustmentsAdjustmentUnion() {
}

type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanDraftSaveParamsAddAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanDraftSaveParamsAddAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAllTrue PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType string

const (
	PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeUsage          PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "usage"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixed          PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeInArrears      PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeUsage, PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeFixed, PlanDraftSaveParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanDraftSaveParamsAddPricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanDraftSaveParamsAddPricesPriceUnion] `json:"price"`
}

func (r PlanDraftSaveParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanDraftSaveParamsAddPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence string

const (
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceOneTime    PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceMonthly    PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceAnnual     PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceCustom     PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceOneTime, PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceMonthly, PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceAnnual, PlanDraftSaveParamsAddPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelType string

const (
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelTypeUnit PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                     `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanDraftSaveParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceModelType] `json:"model_type" api:"required"`
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

func (r PlanDraftSaveParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice],
// [PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanDraftSaveParamsAddPricesPrice].
type PlanDraftSaveParamsAddPricesPriceUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                               `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                          `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                              `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
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
	MatrixValues param.Field[[]PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                           `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanDraftSaveParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
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

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                  `json:"unit_config"`
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsAddPricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsAddPricesPriceCadence string

const (
	PlanDraftSaveParamsAddPricesPriceCadenceOneTime    PlanDraftSaveParamsAddPricesPriceCadence = "one_time"
	PlanDraftSaveParamsAddPricesPriceCadenceMonthly    PlanDraftSaveParamsAddPricesPriceCadence = "monthly"
	PlanDraftSaveParamsAddPricesPriceCadenceQuarterly  PlanDraftSaveParamsAddPricesPriceCadence = "quarterly"
	PlanDraftSaveParamsAddPricesPriceCadenceSemiAnnual PlanDraftSaveParamsAddPricesPriceCadence = "semi_annual"
	PlanDraftSaveParamsAddPricesPriceCadenceAnnual     PlanDraftSaveParamsAddPricesPriceCadence = "annual"
	PlanDraftSaveParamsAddPricesPriceCadenceCustom     PlanDraftSaveParamsAddPricesPriceCadence = "custom"
)

func (r PlanDraftSaveParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceCadenceOneTime, PlanDraftSaveParamsAddPricesPriceCadenceMonthly, PlanDraftSaveParamsAddPricesPriceCadenceQuarterly, PlanDraftSaveParamsAddPricesPriceCadenceSemiAnnual, PlanDraftSaveParamsAddPricesPriceCadenceAnnual, PlanDraftSaveParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsAddPricesPriceModelType string

const (
	PlanDraftSaveParamsAddPricesPriceModelTypeUnit                            PlanDraftSaveParamsAddPricesPriceModelType = "unit"
	PlanDraftSaveParamsAddPricesPriceModelTypeTiered                          PlanDraftSaveParamsAddPricesPriceModelType = "tiered"
	PlanDraftSaveParamsAddPricesPriceModelTypeBulk                            PlanDraftSaveParamsAddPricesPriceModelType = "bulk"
	PlanDraftSaveParamsAddPricesPriceModelTypeBulkWithFilters                 PlanDraftSaveParamsAddPricesPriceModelType = "bulk_with_filters"
	PlanDraftSaveParamsAddPricesPriceModelTypePackage                         PlanDraftSaveParamsAddPricesPriceModelType = "package"
	PlanDraftSaveParamsAddPricesPriceModelTypeMatrix                          PlanDraftSaveParamsAddPricesPriceModelType = "matrix"
	PlanDraftSaveParamsAddPricesPriceModelTypeThresholdTotalAmount            PlanDraftSaveParamsAddPricesPriceModelType = "threshold_total_amount"
	PlanDraftSaveParamsAddPricesPriceModelTypeTieredPackage                   PlanDraftSaveParamsAddPricesPriceModelType = "tiered_package"
	PlanDraftSaveParamsAddPricesPriceModelTypeTieredWithMinimum               PlanDraftSaveParamsAddPricesPriceModelType = "tiered_with_minimum"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTiered                   PlanDraftSaveParamsAddPricesPriceModelType = "grouped_tiered"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTieredMatrix             PlanDraftSaveParamsAddPricesPriceModelType = "grouped_tiered_matrix"
	PlanDraftSaveParamsAddPricesPriceModelTypeTieredPackageWithMinimum        PlanDraftSaveParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	PlanDraftSaveParamsAddPricesPriceModelTypePackageWithAllocation           PlanDraftSaveParamsAddPricesPriceModelType = "package_with_allocation"
	PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithPercent                 PlanDraftSaveParamsAddPricesPriceModelType = "unit_with_percent"
	PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithAllocation            PlanDraftSaveParamsAddPricesPriceModelType = "matrix_with_allocation"
	PlanDraftSaveParamsAddPricesPriceModelTypeTieredMatrixWithAllocation      PlanDraftSaveParamsAddPricesPriceModelType = "tiered_matrix_with_allocation"
	PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts    PlanDraftSaveParamsAddPricesPriceModelType = "matrix_with_threshold_discounts"
	PlanDraftSaveParamsAddPricesPriceModelTypeTieredWithProration             PlanDraftSaveParamsAddPricesPriceModelType = "tiered_with_proration"
	PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithProration               PlanDraftSaveParamsAddPricesPriceModelType = "unit_with_proration"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedAllocation               PlanDraftSaveParamsAddPricesPriceModelType = "grouped_allocation"
	PlanDraftSaveParamsAddPricesPriceModelTypeBulkWithProration               PlanDraftSaveParamsAddPricesPriceModelType = "bulk_with_proration"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      PlanDraftSaveParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       PlanDraftSaveParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     PlanDraftSaveParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithDisplayName           PlanDraftSaveParamsAddPricesPriceModelType = "matrix_with_display_name"
	PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTieredPackage            PlanDraftSaveParamsAddPricesPriceModelType = "grouped_tiered_package"
	PlanDraftSaveParamsAddPricesPriceModelTypeMaxGroupTieredPackage           PlanDraftSaveParamsAddPricesPriceModelType = "max_group_tiered_package"
	PlanDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   PlanDraftSaveParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing PlanDraftSaveParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedBulk           PlanDraftSaveParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	PlanDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedAllocation     PlanDraftSaveParamsAddPricesPriceModelType = "cumulative_grouped_allocation"
	PlanDraftSaveParamsAddPricesPriceModelTypeDailyCreditAllowance            PlanDraftSaveParamsAddPricesPriceModelType = "daily_credit_allowance"
	PlanDraftSaveParamsAddPricesPriceModelTypeMeteredAllowance                PlanDraftSaveParamsAddPricesPriceModelType = "metered_allowance"
	PlanDraftSaveParamsAddPricesPriceModelTypeMinimumComposite                PlanDraftSaveParamsAddPricesPriceModelType = "minimum_composite"
	PlanDraftSaveParamsAddPricesPriceModelTypePercent                         PlanDraftSaveParamsAddPricesPriceModelType = "percent"
	PlanDraftSaveParamsAddPricesPriceModelTypeEventOutput                     PlanDraftSaveParamsAddPricesPriceModelType = "event_output"
	PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithMultiplier              PlanDraftSaveParamsAddPricesPriceModelType = "unit_with_multiplier"
)

func (r PlanDraftSaveParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsAddPricesPriceModelTypeUnit, PlanDraftSaveParamsAddPricesPriceModelTypeTiered, PlanDraftSaveParamsAddPricesPriceModelTypeBulk, PlanDraftSaveParamsAddPricesPriceModelTypeBulkWithFilters, PlanDraftSaveParamsAddPricesPriceModelTypePackage, PlanDraftSaveParamsAddPricesPriceModelTypeMatrix, PlanDraftSaveParamsAddPricesPriceModelTypeThresholdTotalAmount, PlanDraftSaveParamsAddPricesPriceModelTypeTieredPackage, PlanDraftSaveParamsAddPricesPriceModelTypeTieredWithMinimum, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTiered, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTieredMatrix, PlanDraftSaveParamsAddPricesPriceModelTypeTieredPackageWithMinimum, PlanDraftSaveParamsAddPricesPriceModelTypePackageWithAllocation, PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithPercent, PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithAllocation, PlanDraftSaveParamsAddPricesPriceModelTypeTieredMatrixWithAllocation, PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithThresholdDiscounts, PlanDraftSaveParamsAddPricesPriceModelTypeTieredWithProration, PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithProration, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedAllocation, PlanDraftSaveParamsAddPricesPriceModelTypeBulkWithProration, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, PlanDraftSaveParamsAddPricesPriceModelTypeMatrixWithDisplayName, PlanDraftSaveParamsAddPricesPriceModelTypeGroupedTieredPackage, PlanDraftSaveParamsAddPricesPriceModelTypeMaxGroupTieredPackage, PlanDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, PlanDraftSaveParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, PlanDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedBulk, PlanDraftSaveParamsAddPricesPriceModelTypeCumulativeGroupedAllocation, PlanDraftSaveParamsAddPricesPriceModelTypeDailyCreditAllowance, PlanDraftSaveParamsAddPricesPriceModelTypeMeteredAllowance, PlanDraftSaveParamsAddPricesPriceModelTypeMinimumComposite, PlanDraftSaveParamsAddPricesPriceModelTypePercent, PlanDraftSaveParamsAddPricesPriceModelTypeEventOutput, PlanDraftSaveParamsAddPricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

type PlanDraftSaveParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id" api:"required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanDraftSaveParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id" api:"required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanDraftSaveParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id" api:"required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r PlanDraftSaveParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
	AmountDiscount param.Field[string]                                                        `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                 `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                 `json:"applies_to_price_ids"`
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
	PriceType     param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	Tiers         param.Field[interface{}]                                              `json:"tiers"`
	UsageDiscount param.Field[float64]                                                  `json:"usage_discount"`
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustment) ImplementsPlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount],
// [PlanDraftSaveParamsReplaceAdjustmentsAdjustment].
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsPlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion()
}

type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount struct {
	AdjustmentType param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType] `json:"adjustment_type" api:"required"`
	Tiers          param.Field[[]PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier]         `json:"tiers" api:"required"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll] `json:"applies_to_all"`
	// The set of item IDs to which this adjustment applies.
	AppliesToItemIDs param.Field[[]string] `json:"applies_to_item_ids"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// If set, only prices in the specified currency will have the adjustment applied.
	Currency param.Field[string] `json:"currency"`
	// A list of filters that determine which prices this adjustment will apply to.
	Filters param.Field[[]PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter] `json:"filters"`
	// When false, this adjustment will be applied to a single price. Otherwise, it
	// will be applied at the invoice level, possibly to multiple prices.
	IsInvoiceLevel param.Field[bool] `json:"is_invoice_level"`
	// If set, only prices of the specified type will have the adjustment applied.
	PriceType param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType] `json:"price_type"`
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscount) ImplementsPlanDraftSaveParamsReplaceAdjustmentsAdjustmentUnion() {
}

type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType = "tiered_percentage_discount"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier struct {
	// Exclusive lower bound of cumulative spend for this tier.
	LowerBound param.Field[float64] `json:"lower_bound" api:"required"`
	// The percentage (0-1) discounted from spend in this tier.
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Inclusive upper bound of cumulative spend; null for the final open-ended tier.
	UpperBound param.Field[float64] `json:"upper_bound"`
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If set, the adjustment will apply to every price on the subscription.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll bool

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll = true
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAll) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountAppliesToAllTrue:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter struct {
	// The property of the price to filter on.
	Field param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField] `json:"field" api:"required"`
	// Should prices that match the filter be included or excluded.
	Operator param.Field[PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator] `json:"operator" api:"required"`
	// The IDs or values that match this filter.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The property of the price to filter on.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID       PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_id"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID        PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "item_id"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType     PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "price_type"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency      PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "currency"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField = "pricing_unit_id"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersField) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceID, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldItemID, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPriceType, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldCurrency, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "includes"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator = "excludes"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperator) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorIncludes, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountFiltersOperatorExcludes:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage          PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "usage"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_advance"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed_in_arrears"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed          PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "fixed"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears      PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType = "in_arrears"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeUsage, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInAdvance, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixedInArrears, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeFixed, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentNewTieredPercentageDiscountPriceTypeInArrears:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount       PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount            PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount           PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum                  PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum                  PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType = "tiered_percentage_discount"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAdjustmentTypeTieredPercentageDiscount:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, PlanDraftSaveParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id" api:"required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[PlanDraftSaveParamsReplacePricesLicenseAllocationPrice] `json:"license_allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[PlanDraftSaveParamsReplacePricesPriceUnion] `json:"price"`
}

func (r PlanDraftSaveParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type PlanDraftSaveParamsReplacePricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set. Usage prices draw down from these
	// allocations only if they set the same license type.
	LicenseAllocations param.Field[[]PlanDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceCustom     PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelTypeUnit PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelType = "unit"
)

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesLicenseAllocationPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                         `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesLicenseAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// New plan price request body params.
type PlanDraftSaveParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceModelType] `json:"model_type" api:"required"`
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

func (r PlanDraftSaveParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice],
// [PlanDraftSaveParamsReplacePricesPrice].
type PlanDraftSaveParamsReplacePricesPriceUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_matrix pricing
	GroupedTieredMatrixConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig] `json:"grouped_tiered_matrix_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_matrix pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig struct {
	// Per unit rate for usage whose dimension value has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// The billable metric property used to group usage before tiering
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Graduated tiers keyed by dimension value; usage for a value is tiered only
	// against its own rows
	Tiers param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to a dimension value
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier struct {
	// The dimension value this tier applies to
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceGroupedTieredMatrixConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType = "grouped_tiered_matrix"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceModelTypeGroupedTieredMatrix:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                       `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedTieredMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_matrix_with_allocation pricing
	TieredMatrixWithAllocationConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig] `json:"tiered_matrix_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType = "tiered_matrix_with_allocation"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceModelTypeTieredMatrixWithAllocation:
		return true
	}
	return false
}

// Configuration for tiered_matrix_with_allocation pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig struct {
	// Usage allocation, pooled across all matrix cells
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// Per unit rate for usage whose matrix cell has no configured tiers
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// One or two event property values to evaluate matrix cells by
	Dimensions param.Field[[]string] `json:"dimensions" api:"required"`
	// Graduated tiers keyed by matrix cell; usage in a cell is tiered only against its
	// own rows
	Tiers param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier scoped to one matrix cell
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier struct {
	// The matrix cell this tier applies to, as one or two dimension values
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Exclusive tier starting value. The tier runs up to and including the next bound
	// configured for the same matrix cell.
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceTieredMatrixWithAllocationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for matrix_with_threshold_discounts pricing
	MatrixWithThresholdDiscountsConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig] `json:"matrix_with_threshold_discounts_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_threshold_discounts pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig struct {
	// Unit price used for usage that does not match any defined matrix cell.
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount" api:"required"`
	// First matrix dimension key.
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Per-cell unit prices.
	MatrixValues param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue] `json:"matrix_values" api:"required"`
	// Optional second matrix dimension key.
	SecondDimension         param.Field[string]                                                                                                                                  `json:"second_dimension"`
	ThresholdDiscountGroups param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup] `json:"threshold_discount_groups"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	UnitAmount           param.Field[string] `json:"unit_amount" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup struct {
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceMatrixWithThresholdDiscountsConfigThresholdDiscountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType = "matrix_with_threshold_discounts"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceModelTypeMatrixWithThresholdDiscounts:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanMatrixWithThresholdDiscountsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                       `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                               `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                               `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence] `json:"cadence" api:"required"`
	// Configuration for daily_credit_allowance pricing
	DailyCreditAllowanceConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig] `json:"daily_credit_allowance_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for daily_credit_allowance pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig struct {
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
	MatrixValues param.Field[[]PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue] `json:"matrix_values" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Per-dimension credit price for the daily credit allowance model.
type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this value by. For example, ["model"]
	// could be used to apply a different credit rate to each AI model.
	DimensionValues param.Field[[]string] `json:"dimension_values" api:"required"`
	// Credits charged per unit of usage matching the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceDailyCreditAllowanceConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType = "daily_credit_allowance"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceModelTypeDailyCreditAllowance:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanDailyCreditAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// Configuration for metered_allowance pricing
	MeteredAllowanceConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig] `json:"metered_allowance_config" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for metered_allowance pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig struct {
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceMeteredAllowanceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType = "metered_allowance"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceModelTypeMeteredAllowance:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanMeteredAllowancePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                               `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_multiplier pricing
	UnitWithMultiplierConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig] `json:"unit_with_multiplier_config" api:"required"`
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
	ConversionRateConfig param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPrice) ImplementsPlanDraftSaveParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType = "unit_with_multiplier"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}

// Configuration for unit_with_multiplier pricing
type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig struct {
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

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceUnitWithMultiplierConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig struct {
	ConversionRateType param.Field[PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                      `json:"unit_config"`
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig) ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfig].
type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion interface {
	ImplementsPlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigUnion()
}

type PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType string

const (
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit   PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "unit"
	PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeUnit, PlanDraftSaveParamsReplacePricesPriceNewPlanUnitWithMultiplierPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanDraftSaveParamsReplacePricesPriceCadence string

const (
	PlanDraftSaveParamsReplacePricesPriceCadenceOneTime    PlanDraftSaveParamsReplacePricesPriceCadence = "one_time"
	PlanDraftSaveParamsReplacePricesPriceCadenceMonthly    PlanDraftSaveParamsReplacePricesPriceCadence = "monthly"
	PlanDraftSaveParamsReplacePricesPriceCadenceQuarterly  PlanDraftSaveParamsReplacePricesPriceCadence = "quarterly"
	PlanDraftSaveParamsReplacePricesPriceCadenceSemiAnnual PlanDraftSaveParamsReplacePricesPriceCadence = "semi_annual"
	PlanDraftSaveParamsReplacePricesPriceCadenceAnnual     PlanDraftSaveParamsReplacePricesPriceCadence = "annual"
	PlanDraftSaveParamsReplacePricesPriceCadenceCustom     PlanDraftSaveParamsReplacePricesPriceCadence = "custom"
)

func (r PlanDraftSaveParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceCadenceOneTime, PlanDraftSaveParamsReplacePricesPriceCadenceMonthly, PlanDraftSaveParamsReplacePricesPriceCadenceQuarterly, PlanDraftSaveParamsReplacePricesPriceCadenceSemiAnnual, PlanDraftSaveParamsReplacePricesPriceCadenceAnnual, PlanDraftSaveParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type PlanDraftSaveParamsReplacePricesPriceModelType string

const (
	PlanDraftSaveParamsReplacePricesPriceModelTypeUnit                            PlanDraftSaveParamsReplacePricesPriceModelType = "unit"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTiered                          PlanDraftSaveParamsReplacePricesPriceModelType = "tiered"
	PlanDraftSaveParamsReplacePricesPriceModelTypeBulk                            PlanDraftSaveParamsReplacePricesPriceModelType = "bulk"
	PlanDraftSaveParamsReplacePricesPriceModelTypeBulkWithFilters                 PlanDraftSaveParamsReplacePricesPriceModelType = "bulk_with_filters"
	PlanDraftSaveParamsReplacePricesPriceModelTypePackage                         PlanDraftSaveParamsReplacePricesPriceModelType = "package"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMatrix                          PlanDraftSaveParamsReplacePricesPriceModelType = "matrix"
	PlanDraftSaveParamsReplacePricesPriceModelTypeThresholdTotalAmount            PlanDraftSaveParamsReplacePricesPriceModelType = "threshold_total_amount"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTieredPackage                   PlanDraftSaveParamsReplacePricesPriceModelType = "tiered_package"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTieredWithMinimum               PlanDraftSaveParamsReplacePricesPriceModelType = "tiered_with_minimum"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTiered                   PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredMatrix             PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered_matrix"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        PlanDraftSaveParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	PlanDraftSaveParamsReplacePricesPriceModelTypePackageWithAllocation           PlanDraftSaveParamsReplacePricesPriceModelType = "package_with_allocation"
	PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithPercent                 PlanDraftSaveParamsReplacePricesPriceModelType = "unit_with_percent"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithAllocation            PlanDraftSaveParamsReplacePricesPriceModelType = "matrix_with_allocation"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation      PlanDraftSaveParamsReplacePricesPriceModelType = "tiered_matrix_with_allocation"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts    PlanDraftSaveParamsReplacePricesPriceModelType = "matrix_with_threshold_discounts"
	PlanDraftSaveParamsReplacePricesPriceModelTypeTieredWithProration             PlanDraftSaveParamsReplacePricesPriceModelType = "tiered_with_proration"
	PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithProration               PlanDraftSaveParamsReplacePricesPriceModelType = "unit_with_proration"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedAllocation               PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_allocation"
	PlanDraftSaveParamsReplacePricesPriceModelTypeBulkWithProration               PlanDraftSaveParamsReplacePricesPriceModelType = "bulk_with_proration"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithDisplayName           PlanDraftSaveParamsReplacePricesPriceModelType = "matrix_with_display_name"
	PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredPackage            PlanDraftSaveParamsReplacePricesPriceModelType = "grouped_tiered_package"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           PlanDraftSaveParamsReplacePricesPriceModelType = "max_group_tiered_package"
	PlanDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   PlanDraftSaveParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	PlanDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing PlanDraftSaveParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	PlanDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           PlanDraftSaveParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	PlanDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation     PlanDraftSaveParamsReplacePricesPriceModelType = "cumulative_grouped_allocation"
	PlanDraftSaveParamsReplacePricesPriceModelTypeDailyCreditAllowance            PlanDraftSaveParamsReplacePricesPriceModelType = "daily_credit_allowance"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMeteredAllowance                PlanDraftSaveParamsReplacePricesPriceModelType = "metered_allowance"
	PlanDraftSaveParamsReplacePricesPriceModelTypeMinimumComposite                PlanDraftSaveParamsReplacePricesPriceModelType = "minimum_composite"
	PlanDraftSaveParamsReplacePricesPriceModelTypePercent                         PlanDraftSaveParamsReplacePricesPriceModelType = "percent"
	PlanDraftSaveParamsReplacePricesPriceModelTypeEventOutput                     PlanDraftSaveParamsReplacePricesPriceModelType = "event_output"
	PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithMultiplier              PlanDraftSaveParamsReplacePricesPriceModelType = "unit_with_multiplier"
)

func (r PlanDraftSaveParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case PlanDraftSaveParamsReplacePricesPriceModelTypeUnit, PlanDraftSaveParamsReplacePricesPriceModelTypeTiered, PlanDraftSaveParamsReplacePricesPriceModelTypeBulk, PlanDraftSaveParamsReplacePricesPriceModelTypeBulkWithFilters, PlanDraftSaveParamsReplacePricesPriceModelTypePackage, PlanDraftSaveParamsReplacePricesPriceModelTypeMatrix, PlanDraftSaveParamsReplacePricesPriceModelTypeThresholdTotalAmount, PlanDraftSaveParamsReplacePricesPriceModelTypeTieredPackage, PlanDraftSaveParamsReplacePricesPriceModelTypeTieredWithMinimum, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTiered, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredMatrix, PlanDraftSaveParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, PlanDraftSaveParamsReplacePricesPriceModelTypePackageWithAllocation, PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithPercent, PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithAllocation, PlanDraftSaveParamsReplacePricesPriceModelTypeTieredMatrixWithAllocation, PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithThresholdDiscounts, PlanDraftSaveParamsReplacePricesPriceModelTypeTieredWithProration, PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithProration, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedAllocation, PlanDraftSaveParamsReplacePricesPriceModelTypeBulkWithProration, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, PlanDraftSaveParamsReplacePricesPriceModelTypeMatrixWithDisplayName, PlanDraftSaveParamsReplacePricesPriceModelTypeGroupedTieredPackage, PlanDraftSaveParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, PlanDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, PlanDraftSaveParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, PlanDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, PlanDraftSaveParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation, PlanDraftSaveParamsReplacePricesPriceModelTypeDailyCreditAllowance, PlanDraftSaveParamsReplacePricesPriceModelTypeMeteredAllowance, PlanDraftSaveParamsReplacePricesPriceModelTypeMinimumComposite, PlanDraftSaveParamsReplacePricesPriceModelTypePercent, PlanDraftSaveParamsReplacePricesPriceModelTypeEventOutput, PlanDraftSaveParamsReplacePricesPriceModelTypeUnitWithMultiplier:
		return true
	}
	return false
}
