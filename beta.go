// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// BetaService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	Options        []option.RequestOption
	ExternalPlanID *BetaExternalPlanIDService
}

// NewBetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaService(opts ...option.RequestOption) (r *BetaService) {
	r = &BetaService{}
	r.Options = opts
	r.ExternalPlanID = NewBetaExternalPlanIDService(opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows the creation of a new plan version for an existing plan.
func (r *BetaService) NewPlanVersion(ctx context.Context, planID string, body BetaNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *BetaService) FetchPlanVersion(ctx context.Context, planID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/versions/%s", planID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API endpoint is in beta and its interface may change. It is recommended for
// use only in test mode.
//
// This endpoint allows setting the default version of a plan.
func (r *BetaService) SetDefaultPlanVersion(ctx context.Context, planID string, body BetaSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s/set_default_version", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersion struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionAdjustment `json:"adjustments,required"`
	CreatedAt   time.Time               `json:"created_at,required" format:"date-time"`
	PlanPhases  []PlanVersionPhase      `json:"plan_phases,required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []shared.Price  `json:"prices,required"`
	Version int64           `json:"version,required"`
	JSON    planVersionJSON `json:"-"`
}

// planVersionJSON contains the JSON metadata for the struct [PlanVersion]
type planVersionJSON struct {
	Adjustments apijson.Field
	CreatedAt   apijson.Field
	PlanPhases  apijson.Field
	Prices      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlanVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionJSON) RawJSON() string {
	return r.raw
}

type PlanVersionAdjustment struct {
	ID             string                               `json:"id,required"`
	AdjustmentType PlanVersionAdjustmentsAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of [[]shared.TransformPriceFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invoice, false for adjustments that
	// apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The adjustment id this adjustment replaces. This adjustment will take the place
	// of the replaced adjustment in plan version migrations.
	ReplacesAdjustmentID string `json:"replaces_adjustment_id,required,nullable"`
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
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                   `json:"usage_discount"`
	JSON          planVersionAdjustmentJSON `json:"-"`
	union         PlanVersionAdjustmentsUnion
}

// planVersionAdjustmentJSON contains the JSON metadata for the struct
// [PlanVersionAdjustment]
type planVersionAdjustmentJSON struct {
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
	UsageDiscount        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r planVersionAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *PlanVersionAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = PlanVersionAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PlanVersionAdjustmentsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment], [shared.PlanPhaseMaximumAdjustment].
func (r PlanVersionAdjustment) AsUnion() PlanVersionAdjustmentsUnion {
	return r.union
}

// Union satisfied by [shared.PlanPhaseUsageDiscountAdjustment],
// [shared.PlanPhaseAmountDiscountAdjustment],
// [shared.PlanPhasePercentageDiscountAdjustment],
// [shared.PlanPhaseMinimumAdjustment] or [shared.PlanPhaseMaximumAdjustment].
type PlanVersionAdjustmentsUnion interface {
	ImplementsPlanVersionAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanVersionAdjustmentsUnion)(nil)).Elem(),
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

type PlanVersionAdjustmentsAdjustmentType string

const (
	PlanVersionAdjustmentsAdjustmentTypeUsageDiscount      PlanVersionAdjustmentsAdjustmentType = "usage_discount"
	PlanVersionAdjustmentsAdjustmentTypeAmountDiscount     PlanVersionAdjustmentsAdjustmentType = "amount_discount"
	PlanVersionAdjustmentsAdjustmentTypePercentageDiscount PlanVersionAdjustmentsAdjustmentType = "percentage_discount"
	PlanVersionAdjustmentsAdjustmentTypeMinimum            PlanVersionAdjustmentsAdjustmentType = "minimum"
	PlanVersionAdjustmentsAdjustmentTypeMaximum            PlanVersionAdjustmentsAdjustmentType = "maximum"
)

func (r PlanVersionAdjustmentsAdjustmentType) IsKnown() bool {
	switch r {
	case PlanVersionAdjustmentsAdjustmentTypeUsageDiscount, PlanVersionAdjustmentsAdjustmentTypeAmountDiscount, PlanVersionAdjustmentsAdjustmentTypePercentageDiscount, PlanVersionAdjustmentsAdjustmentTypeMinimum, PlanVersionAdjustmentsAdjustmentTypeMaximum:
		return true
	}
	return false
}

type PlanVersionPhase struct {
	ID          string `json:"id,required"`
	Description string `json:"description,required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                        `json:"duration,required,nullable"`
	DurationUnit PlanVersionPhaseDurationUnit `json:"duration_unit,required,nullable"`
	Name         string                       `json:"name,required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                `json:"order,required"`
	JSON  planVersionPhaseJSON `json:"-"`
}

// planVersionPhaseJSON contains the JSON metadata for the struct
// [PlanVersionPhase]
type planVersionPhaseJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Duration     apijson.Field
	DurationUnit apijson.Field
	Name         apijson.Field
	Order        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PlanVersionPhase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planVersionPhaseJSON) RawJSON() string {
	return r.raw
}

type PlanVersionPhaseDurationUnit string

const (
	PlanVersionPhaseDurationUnitDaily      PlanVersionPhaseDurationUnit = "daily"
	PlanVersionPhaseDurationUnitMonthly    PlanVersionPhaseDurationUnit = "monthly"
	PlanVersionPhaseDurationUnitQuarterly  PlanVersionPhaseDurationUnit = "quarterly"
	PlanVersionPhaseDurationUnitSemiAnnual PlanVersionPhaseDurationUnit = "semi_annual"
	PlanVersionPhaseDurationUnitAnnual     PlanVersionPhaseDurationUnit = "annual"
)

func (r PlanVersionPhaseDurationUnit) IsKnown() bool {
	switch r {
	case PlanVersionPhaseDurationUnitDaily, PlanVersionPhaseDurationUnitMonthly, PlanVersionPhaseDurationUnitQuarterly, PlanVersionPhaseDurationUnitSemiAnnual, PlanVersionPhaseDurationUnitAnnual:
		return true
	}
	return false
}

type BetaNewPlanVersionParams struct {
	// New version number.
	Version param.Field[int64] `json:"version,required"`
	// Additional adjustments to be added to the plan.
	AddAdjustments param.Field[[]BetaNewPlanVersionParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the plan.
	AddPrices param.Field[[]BetaNewPlanVersionParamsAddPrice] `json:"add_prices"`
	// Adjustments to be removed from the plan.
	RemoveAdjustments param.Field[[]BetaNewPlanVersionParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Prices to be removed from the plan.
	RemovePrices param.Field[[]BetaNewPlanVersionParamsRemovePrice] `json:"remove_prices"`
	// Adjustments to be replaced with additional adjustments on the plan.
	ReplaceAdjustments param.Field[[]BetaNewPlanVersionParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Prices to be replaced with additional prices on the plan.
	ReplacePrices param.Field[[]BetaNewPlanVersionParamsReplacePrice] `json:"replace_prices"`
	// Set this new plan version as the default
	SetAsDefault param.Field[bool] `json:"set_as_default"`
}

func (r BetaNewPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                         `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                  `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                  `json:"applies_to_price_ids"`
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
	PriceType     param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                   `json:"usage_discount"`
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustment) ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [BetaNewPlanVersionParamsAddAdjustmentsAdjustment].
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsBetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
}

type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType string

const (
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "usage"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed          BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears      BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeUsage, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeFixed, BetaNewPlanVersionParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPrice struct {
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The price to add to the plan
type BetaNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	BulkConfig                param.Field[shared.BulkConfigParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}]                              `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig        param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity                param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig           param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig               param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[shared.MatrixConfigParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                              `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                              `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                              `json:"metadata"`
	MinimumConfig               param.Field[interface{}]                              `json:"minimum_config"`
	PackageConfig               param.Field[shared.PackageConfigParam]                `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}]                              `json:"package_with_allocation_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]                   `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]              `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]              `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]              `json:"threshold_total_amount_config"`
	TieredConfig                          param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}]              `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigParam]   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}]              `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}]              `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanPackagePriceParam],
// [shared.NewPlanMatrixPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam], [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanTierWithProrationPriceParam],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaNewPlanVersionParamsAddPricesPrice].
type BetaNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence                           param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}]                                                               `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit                            BetaNewPlanVersionParamsAddPricesPriceModelType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackage                         BetaNewPlanVersionParamsAddPricesPriceModelType = "package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered                          BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk                            BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMinimum                         BetaNewPlanVersionParamsAddPricesPriceModelType = "minimum"
)

func (r BetaNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsAddPricesPriceModelTypeMinimum:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id,required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                             `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                      `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                      `json:"applies_to_price_ids"`
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
	PriceType     param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                       `json:"usage_discount"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment) ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsBetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
}

type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePrice struct {
	// The id of the price on the plan to replace in the plan.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The price to add to the plan
	Price param.Field[BetaNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The price to add to the plan
type BetaNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                              `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	BulkConfig                param.Field[shared.BulkConfigParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[interface{}]                              `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig        param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// For dimensional price: specifies a price group and dimension values
	DimensionalPriceConfiguration param.Field[shared.NewDimensionalPriceConfigurationParam] `json:"dimensional_price_configuration"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity                param.Field[float64]     `json:"fixed_price_quantity"`
	GroupedAllocationConfig           param.Field[interface{}] `json:"grouped_allocation_config"`
	GroupedTieredConfig               param.Field[interface{}] `json:"grouped_tiered_config"`
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                param.Field[shared.MatrixConfigParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                              `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                              `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                              `json:"metadata"`
	MinimumConfig               param.Field[interface{}]                              `json:"minimum_config"`
	PackageConfig               param.Field[shared.PackageConfigParam]                `json:"package_config"`
	PackageWithAllocationConfig param.Field[interface{}]                              `json:"package_with_allocation_config"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID                           param.Field[string]                   `json:"reference_id"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]              `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]              `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]              `json:"threshold_total_amount_config"`
	TieredConfig                          param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}]              `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigParam]   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}]              `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}]              `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The price to add to the plan
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanPackagePriceParam],
// [shared.NewPlanMatrixPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam], [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanTierWithProrationPriceParam],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaNewPlanVersionParamsReplacePricesPrice].
type BetaNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence                           param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMinimum                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMinimum:
		return true
	}
	return false
}

type BetaSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version,required"`
}

func (r BetaSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
