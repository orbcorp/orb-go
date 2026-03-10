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
// BetaService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	Options []option.RequestOption
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
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

// This endpoint allows the creation of a new plan version for an existing plan.
func (r *BetaService) NewPlanVersion(ctx context.Context, planID string, body BetaNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/versions", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *BetaService) FetchPlanVersion(ctx context.Context, planID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	if version == "" {
		err = errors.New("missing required version parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/versions/%s", planID, version)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows setting the default version of a plan.
func (r *BetaService) SetDefaultPlanVersion(ctx context.Context, planID string, body BetaSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("plans/%s/set_default_version", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The PlanVersion resource represents the prices and adjustments present on a
// specific version of a plan.
type PlanVersion struct {
	// Adjustments for this plan. If the plan has phases, this includes adjustments
	// across all phases of the plan.
	Adjustments []PlanVersionAdjustment `json:"adjustments" api:"required"`
	CreatedAt   time.Time               `json:"created_at" api:"required" format:"date-time"`
	PlanPhases  []PlanVersionPhase      `json:"plan_phases" api:"required,nullable"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices  []shared.Price  `json:"prices" api:"required"`
	Version int64           `json:"version" api:"required"`
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
	ID             string                               `json:"id" api:"required"`
	AdjustmentType PlanVersionAdjustmentsAdjustmentType `json:"adjustment_type" api:"required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids" api:"required"`
	// This field can have the runtime type of
	// [[]shared.PlanPhaseUsageDiscountAdjustmentFilter],
	// [[]shared.PlanPhaseAmountDiscountAdjustmentFilter],
	// [[]shared.PlanPhasePercentageDiscountAdjustmentFilter],
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
	ID          string `json:"id" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// How many terms of length `duration_unit` this phase is active for. If null, this
	// phase is evergreen and active indefinitely
	Duration     int64                        `json:"duration" api:"required,nullable"`
	DurationUnit PlanVersionPhaseDurationUnit `json:"duration_unit" api:"required,nullable"`
	Name         string                       `json:"name" api:"required"`
	// Determines the ordering of the phase in a plan's lifecycle. 1 = first phase.
	Order int64                `json:"order" api:"required"`
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
	Version param.Field[int64] `json:"version" api:"required"`
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
	Adjustment param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
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
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion] `json:"license_allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[BetaNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID             param.Field[string]      `json:"item_id" api:"required"`
	LicenseAllocations param.Field[interface{}] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	Currency param.Field[string] `json:"currency"`
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
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                   param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig            param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig        param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig      param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig              param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithPercentConfig   param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The license allocation price to add to the plan.
//
// Satisfied by
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPrice].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion interface {
	implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType = "unit"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered pricing
	TieredConfig param.Field[shared.TieredConfigParam] `json:"tiered_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice struct {
	// Configuration for bulk pricing
	BulkConfig param.Field[shared.BulkConfigParam] `json:"bulk_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType = "bulk"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package pricing
	PackageConfig param.Field[shared.PackageConfigParam] `json:"package_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType = "package"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType = "matrix"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for threshold_total_amount pricing
	ThresholdTotalAmountConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig] `json:"threshold_total_amount_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// Configuration for threshold_total_amount pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig struct {
	// When the quantity consumed passes a provided threshold, the configured total
	// will be charged
	ConsumptionTable param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable] `json:"consumption_table" api:"required"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single threshold
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable struct {
	Threshold param.Field[string] `json:"threshold" api:"required"`
	// Total amount for this threshold
	TotalAmount param.Field[string] `json:"total_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package pricing
	TieredPackageConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig] `json:"tiered_package_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType = "tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// Configuration for tiered_package pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig struct {
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds. The tier bounds are defined based on
	// the total quantity rather than the number of packages, so they must be multiples
	// of the package size.
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier with business logic
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier struct {
	// Price per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_minimum pricing
	TieredWithMinimumConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig] `json:"tiered_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_with_minimum pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig struct {
	// Tiered pricing with a minimum amount dependent on the volume tier. Tiers are
	// defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier] `json:"tiers" api:"required"`
	// If true, tiers with an accrued amount of 0 will not be included in the rating.
	HideZeroAmountTiers param.Field[bool] `json:"hide_zero_amount_tiers"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered pricing
	GroupedTieredConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig] `json:"grouped_tiered_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig struct {
	// The billable metric property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Apply tiered pricing to each segment generated after grouping with the provided
	// key
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package_with_minimum pricing
	TieredPackageWithMinimumConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig] `json:"tiered_package_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_package_with_minimum pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig struct {
	PackageSize param.Field[float64] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package_with_allocation pricing
	PackageWithAllocationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig] `json:"package_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// Configuration for package_with_allocation pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig struct {
	Allocation    param.Field[string] `json:"allocation" api:"required"`
	PackageAmount param.Field[string] `json:"package_amount" api:"required"`
	PackageSize   param.Field[string] `json:"package_size" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_percent pricing
	UnitWithPercentConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig] `json:"unit_with_percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// Configuration for unit_with_percent pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig struct {
	// What percent, out of 100, of the calculated total to charge
	Percent param.Field[string] `json:"percent" api:"required"`
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                      `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_proration pricing
	UnitWithProrationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig] `json:"unit_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// Configuration for unit_with_proration pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_allocation pricing
	GroupedAllocationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig] `json:"grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_allocation pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig struct {
	// Usage allocation per group
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// How to determine the groups that should each be allocated some quantity
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Unit rate for post-allocation
	OverageUnitRate param.Field[string] `json:"overage_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice struct {
	// Configuration for bulk_with_proration pricing
	BulkWithProrationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig] `json:"bulk_with_proration_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_proration pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier with proration
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier struct {
	// Cost per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_prorated_minimum pricing
	GroupedWithProratedMinimumConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig] `json:"grouped_with_prorated_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_prorated_minimum pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig struct {
	// How to determine the groups that should each have a minimum
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group
	Minimum param.Field[string] `json:"minimum" api:"required"`
	// The amount to charge per unit
	UnitRate param.Field[string] `json:"unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                             `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_metered_minimum pricing
	GroupedWithMeteredMinimumConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig] `json:"grouped_with_metered_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_metered_minimum pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig struct {
	// Used to partition the usage into groups. The minimum amount is applied to each
	// group.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group per unit
	MinimumUnitAmount param.Field[string] `json:"minimum_unit_amount" api:"required"`
	// Used to determine the unit rate
	PricingKey param.Field[string] `json:"pricing_key" api:"required"`
	// Scale the unit rates by the scaling factor.
	ScalingFactors param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor] `json:"scaling_factors" api:"required"`
	// Used to determine the unit rate scaling factor
	ScalingKey param.Field[string] `json:"scaling_key" api:"required"`
	// Apply per unit pricing to each pricing value. The minimum amount is applied any
	// unmatched usage.
	UnitAmounts param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a scaling factor
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor struct {
	ScalingFactor param.Field[string] `json:"scaling_factor" api:"required"`
	ScalingValue  param.Field[string] `json:"scaling_value" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount struct {
	PricingValue param.Field[string] `json:"pricing_value" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                              `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_display_name pricing
	MatrixWithDisplayNameConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig] `json:"matrix_with_display_name_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for matrix_with_display_name pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig struct {
	// Used to determine the unit rate
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Apply per unit pricing to each dimension value
	UnitAmounts param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount item
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount struct {
	// The dimension value
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	// Display name for this dimension value
	DisplayName param.Field[string] `json:"display_name" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_package pricing
	GroupedTieredPackageConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig] `json:"grouped_tiered_package_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_package pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig struct {
	// The event property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier struct {
	// Per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for max_group_tiered_package pricing
	MaxGroupTieredPackageConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig] `json:"max_group_tiered_package_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for max_group_tiered_package pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig struct {
	// The event property used to group before tiering the group with the highest value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing to the largest group after grouping with the provided key.
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_unit_pricing pricing
	ScalableMatrixWithUnitPricingConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig] `json:"scalable_matrix_with_unit_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_unit_pricing pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig struct {
	// Used to determine the unit rate
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	// The final unit price to rate against the output of the matrix
	UnitPrice param.Field[string] `json:"unit_price" api:"required"`
	// The property used to group this price
	GroupingKey param.Field[string] `json:"grouping_key"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
	// Used to determine the unit rate (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_tiered_pricing pricing
	ScalableMatrixWithTieredPricingConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig] `json:"scalable_matrix_with_tiered_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_tiered_pricing pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig struct {
	// Used for the scalable matrix first dimension
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	Tiers                param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier]                `json:"tiers" api:"required"`
	// Used for the scalable matrix second dimension (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier entry with business logic
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	UnitAmount     param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_bulk pricing
	CumulativeGroupedBulkConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig] `json:"cumulative_grouped_bulk_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_bulk pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig struct {
	// Each tier lower bound must have the same group of values.
	DimensionValues param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue] `json:"dimension_values" api:"required"`
	Group           param.Field[string]                                                                                                                                           `json:"group" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a dimension value entry
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue struct {
	// Grouping key value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Unit amount for this combination
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                              `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for minimum_composite pricing
	MinimumCompositeConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig] `json:"minimum_composite_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for minimum_composite pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount" api:"required"`
	// If true, subtotals from this price are prorated based on the service period
	Prorated param.Field[bool] `json:"prorated"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType = "minimum_composite"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                   `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType = "percent"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                   `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) implementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig struct {
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

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit                            BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "unit"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTiered                          BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "tiered"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulk                            BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "bulk"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulkWithFilters                 BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "bulk_with_filters"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePackage                         BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "package"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrix                          BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "matrix"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "cumulative_grouped_allocation"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMinimumComposite                BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "minimum_composite"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePercent                         BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "percent"
	BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeEventOutput                     BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnit, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTiered, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulk, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulkWithFilters, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePackage, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrix, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredPackage, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeMinimumComposite, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypePercent, BetaNewPlanVersionParamsAddPricesLicenseAllocationPriceModelTypeEventOutput:
		return true
	}
	return false
}

// New plan price request body params.
type BetaNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type" api:"required"`
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
	Currency param.Field[string] `json:"currency"`
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
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                   param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig            param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig        param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig      param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig              param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithPercentConfig   param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice],
// [BetaNewPlanVersionParamsAddPricesPrice].
type BetaNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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

// Configuration for grouped_with_min_max_thresholds pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
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
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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

type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                     `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsBetaNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type BetaNewPlanVersionParamsAddPricesPriceModelType string

const (
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit                            BetaNewPlanVersionParamsAddPricesPriceModelType = "unit"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered                          BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk                            BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters                 BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_filters"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackage                         BetaNewPlanVersionParamsAddPricesPriceModelType = "package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation     BetaNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_allocation"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite                BetaNewPlanVersionParamsAddPricesPriceModelType = "minimum_composite"
	BetaNewPlanVersionParamsAddPricesPriceModelTypePercent                         BetaNewPlanVersionParamsAddPricesPriceModelType = "percent"
	BetaNewPlanVersionParamsAddPricesPriceModelTypeEventOutput                     BetaNewPlanVersionParamsAddPricesPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters, BetaNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedAllocation, BetaNewPlanVersionParamsAddPricesPriceModelTypeMinimumComposite, BetaNewPlanVersionParamsAddPricesPriceModelTypePercent, BetaNewPlanVersionParamsAddPricesPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsRemoveAdjustment struct {
	// The id of the adjustment to remove from on the plan.
	AdjustmentID param.Field[string] `json:"adjustment_id" api:"required"`
	// The phase to remove this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsRemovePrice struct {
	// The id of the price to remove from the plan.
	PriceID param.Field[string] `json:"price_id" api:"required"`
	// The phase to remove this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the plan.
	Adjustment param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment" api:"required"`
	// The id of the adjustment on the plan to replace in the plan.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id" api:"required"`
	// The phase to replace this adjustment from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
}

func (r BetaNewPlanVersionParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the plan.
type BetaNewPlanVersionParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[BetaNewPlanVersionParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type" api:"required"`
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
	ReplacesPriceID param.Field[string] `json:"replaces_price_id" api:"required"`
	// The allocation price to add to the plan.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The license allocation price to add to the plan.
	LicenseAllocationPrice param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion] `json:"license_allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[BetaNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The license allocation price to add to the plan.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID             param.Field[string]      `json:"item_id" api:"required"`
	LicenseAllocations param.Field[interface{}] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType] `json:"model_type" api:"required"`
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
	Currency param.Field[string] `json:"currency"`
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
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                   param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig            param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig        param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig      param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig              param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithPercentConfig   param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The license allocation price to add to the plan.
//
// Satisfied by
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPrice].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion interface {
	implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType = "unit"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered pricing
	TieredConfig param.Field[shared.TieredConfigParam] `json:"tiered_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPrice struct {
	// Configuration for bulk pricing
	BulkConfig param.Field[shared.BulkConfigParam] `json:"bulk_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType = "bulk"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                      `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package pricing
	PackageConfig param.Field[shared.PackageConfigParam] `json:"package_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType = "package"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceModelTypePackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType = "matrix"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for threshold_total_amount pricing
	ThresholdTotalAmountConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig] `json:"threshold_total_amount_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// Configuration for threshold_total_amount pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig struct {
	// When the quantity consumed passes a provided threshold, the configured total
	// will be charged
	ConsumptionTable param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable] `json:"consumption_table" api:"required"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single threshold
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable struct {
	Threshold param.Field[string] `json:"threshold" api:"required"`
	// Total amount for this threshold
	TotalAmount param.Field[string] `json:"total_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package pricing
	TieredPackageConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig] `json:"tiered_package_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType = "tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// Configuration for tiered_package pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig struct {
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds. The tier bounds are defined based on
	// the total quantity rather than the number of packages, so they must be multiples
	// of the package size.
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier with business logic
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier struct {
	// Price per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_minimum pricing
	TieredWithMinimumConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig] `json:"tiered_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_with_minimum pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig struct {
	// Tiered pricing with a minimum amount dependent on the volume tier. Tiers are
	// defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier] `json:"tiers" api:"required"`
	// If true, tiers with an accrued amount of 0 will not be included in the rating.
	HideZeroAmountTiers param.Field[bool] `json:"hide_zero_amount_tiers"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceTieredWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered pricing
	GroupedTieredConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig] `json:"grouped_tiered_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig struct {
	// The billable metric property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Apply tiered pricing to each segment generated after grouping with the provided
	// key
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceGroupedTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType = "grouped_tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_package_with_minimum pricing
	TieredPackageWithMinimumConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig] `json:"tiered_package_with_minimum_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_package_with_minimum pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig struct {
	PackageSize param.Field[float64] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier struct {
	MinimumAmount  param.Field[string] `json:"minimum_amount" api:"required"`
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                               `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for package_with_allocation pricing
	PackageWithAllocationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig] `json:"package_with_allocation_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// Configuration for package_with_allocation pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig struct {
	Allocation    param.Field[string] `json:"allocation" api:"required"`
	PackageAmount param.Field[string] `json:"package_amount" api:"required"`
	PackageSize   param.Field[string] `json:"package_size" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPricePackageWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_percent pricing
	UnitWithPercentConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig] `json:"unit_with_percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// Configuration for unit_with_percent pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig struct {
	// What percent, out of 100, of the calculated total to charge
	Percent param.Field[string] `json:"percent" api:"required"`
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceUnitWithPercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                      `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                          `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for unit_with_proration pricing
	UnitWithProrationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig] `json:"unit_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// Configuration for unit_with_proration pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceUnitWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_allocation pricing
	GroupedAllocationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig] `json:"grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_allocation pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig struct {
	// Usage allocation per group
	Allocation param.Field[string] `json:"allocation" api:"required"`
	// How to determine the groups that should each be allocated some quantity
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Unit rate for post-allocation
	OverageUnitRate param.Field[string] `json:"overage_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice struct {
	// Configuration for bulk_with_proration pricing
	BulkWithProrationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig] `json:"bulk_with_proration_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// Configuration for bulk_with_proration pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier with proration
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier struct {
	// Cost per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceBulkWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_prorated_minimum pricing
	GroupedWithProratedMinimumConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig] `json:"grouped_with_prorated_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_prorated_minimum pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig struct {
	// How to determine the groups that should each have a minimum
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group
	Minimum param.Field[string] `json:"minimum" api:"required"`
	// The amount to charge per unit
	UnitRate param.Field[string] `json:"unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                 `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_metered_minimum pricing
	GroupedWithMeteredMinimumConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig] `json:"grouped_with_metered_minimum_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_metered_minimum pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig struct {
	// Used to partition the usage into groups. The minimum amount is applied to each
	// group.
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The minimum amount to charge per group per unit
	MinimumUnitAmount param.Field[string] `json:"minimum_unit_amount" api:"required"`
	// Used to determine the unit rate
	PricingKey param.Field[string] `json:"pricing_key" api:"required"`
	// Scale the unit rates by the scaling factor.
	ScalingFactors param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor] `json:"scaling_factors" api:"required"`
	// Used to determine the unit rate scaling factor
	ScalingKey param.Field[string] `json:"scaling_key" api:"required"`
	// Apply per unit pricing to each pricing value. The minimum amount is applied any
	// unmatched usage.
	UnitAmounts param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a scaling factor
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor struct {
	ScalingFactor param.Field[string] `json:"scaling_factor" api:"required"`
	ScalingValue  param.Field[string] `json:"scaling_value" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount struct {
	PricingValue param.Field[string] `json:"pricing_value" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for matrix_with_display_name pricing
	MatrixWithDisplayNameConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig] `json:"matrix_with_display_name_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for matrix_with_display_name pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig struct {
	// Used to determine the unit rate
	Dimension param.Field[string] `json:"dimension" api:"required"`
	// Apply per unit pricing to each dimension value
	UnitAmounts param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount] `json:"unit_amounts" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount item
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount struct {
	// The dimension value
	DimensionValue param.Field[string] `json:"dimension_value" api:"required"`
	// Display name for this dimension value
	DisplayName param.Field[string] `json:"display_name" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_tiered_package pricing
	GroupedTieredPackageConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig] `json:"grouped_tiered_package_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_package pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig struct {
	// The event property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier struct {
	// Per package
	PerUnit        param.Field[string] `json:"per_unit" api:"required"`
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceGroupedTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                           `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for max_group_tiered_package pricing
	MaxGroupTieredPackageConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig] `json:"max_group_tiered_package_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for max_group_tiered_package pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig struct {
	// The event property used to group before tiering the group with the highest value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	PackageSize param.Field[string] `json:"package_size" api:"required"`
	// Apply tiered pricing to the largest group after grouping with the provided key.
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_unit_pricing pricing
	ScalableMatrixWithUnitPricingConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig] `json:"scalable_matrix_with_unit_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_unit_pricing pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig struct {
	// Used to determine the unit rate
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	// The final unit price to rate against the output of the matrix
	UnitPrice param.Field[string] `json:"unit_price" api:"required"`
	// The property used to group this price
	GroupingKey param.Field[string] `json:"grouping_key"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
	// Used to determine the unit rate (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for scalable_matrix_with_tiered_pricing pricing
	ScalableMatrixWithTieredPricingConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig] `json:"scalable_matrix_with_tiered_pricing_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_tiered_pricing pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig struct {
	// Used for the scalable matrix first dimension
	FirstDimension param.Field[string] `json:"first_dimension" api:"required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor] `json:"matrix_scaling_factors" api:"required"`
	Tiers                param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier]                `json:"tiers" api:"required"`
	// Used for the scalable matrix second dimension (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor struct {
	FirstDimensionValue  param.Field[string] `json:"first_dimension_value" api:"required"`
	ScalingFactor        param.Field[string] `json:"scaling_factor" api:"required"`
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactor) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier entry with business logic
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier struct {
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	UnitAmount     param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                      `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_bulk pricing
	CumulativeGroupedBulkConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig] `json:"cumulative_grouped_bulk_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_bulk pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig struct {
	// Each tier lower bound must have the same group of values.
	DimensionValues param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue] `json:"dimension_values" api:"required"`
	Group           param.Field[string]                                                                                                                                               `json:"group" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a dimension value entry
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue struct {
	// Grouping key value
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Unit amount for this combination
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// Configuration for minimum_composite pricing
	MinimumCompositeConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig] `json:"minimum_composite_config" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for minimum_composite pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount" api:"required"`
	// If true, subtotals from this price are prorated based on the service period
	Prorated param.Field[bool] `json:"prorated"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceMinimumCompositeConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType = "minimum_composite"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceModelTypeMinimumComposite:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationMinimumCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType = "percent"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                       `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// License allocations to associate with this price. Each entry defines a
	// per-license credit pool granted each cadence. Requires license_type_id or
	// license_type_configuration to be set.
	LicenseAllocations param.Field[[]BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation] `json:"license_allocations" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPrice) implementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig struct {
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

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation struct {
	// The amount of credits granted per active license per cadence.
	Amount param.Field[string] `json:"amount" api:"required"`
	// The currency of the license allocation.
	Currency param.Field[string] `json:"currency" api:"required"`
	// When True, overage beyond the allocation is written off.
	WriteOffOverage param.Field[bool] `json:"write_off_overage"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceLicenseAllocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceNewLicenseAllocationEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit                            BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "unit"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTiered                          BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "tiered"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulk                            BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "bulk"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulkWithFilters                 BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "bulk_with_filters"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePackage                         BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "package"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrix                          BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "matrix"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "cumulative_grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMinimumComposite                BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "minimum_composite"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePercent                         BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "percent"
	BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeEventOutput                     BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnit, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTiered, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulk, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulkWithFilters, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePackage, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrix, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredPackage, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeCumulativeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeMinimumComposite, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypePercent, BetaNewPlanVersionParamsReplacePricesLicenseAllocationPriceModelTypeEventOutput:
		return true
	}
	return false
}

// New plan price request body params.
type BetaNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type" api:"required"`
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
	Currency param.Field[string] `json:"currency"`
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
	GroupedTieredPackageConfig        param.Field[interface{}] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig   param.Field[interface{}] `json:"grouped_with_metered_minimum_config"`
	GroupedWithMinMaxThresholdsConfig param.Field[interface{}] `json:"grouped_with_min_max_thresholds_config"`
	GroupedWithProratedMinimumConfig  param.Field[interface{}] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumCompositeConfig      param.Field[interface{}]                            `json:"minimum_composite_config"`
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
	TieredConfig                   param.Field[shared.TieredConfigParam] `json:"tiered_config"`
	TieredPackageConfig            param.Field[interface{}]              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig param.Field[interface{}]              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig        param.Field[interface{}]              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig      param.Field[interface{}]              `json:"tiered_with_proration_config"`
	// Configuration for unit pricing
	UnitConfig              param.Field[shared.UnitConfigParam] `json:"unit_config"`
	UnitWithPercentConfig   param.Field[interface{}]            `json:"unit_with_percent_config"`
	UnitWithProrationConfig param.Field[interface{}]            `json:"unit_with_proration_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice],
// [BetaNewPlanVersionParamsReplacePricesPrice].
type BetaNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config" api:"required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters" api:"required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key" api:"required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                        `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound" api:"required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                            `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence" api:"required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type" api:"required"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
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

// Configuration for grouped_with_min_max_thresholds pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge" api:"required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge" api:"required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
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
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
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

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence] `json:"cadence" api:"required"`
	// Configuration for cumulative_grouped_allocation pricing
	CumulativeGroupedAllocationConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig] `json:"cumulative_grouped_allocation_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_allocation pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig struct {
	// The overall allocation across all groups
	CumulativeAllocation param.Field[string] `json:"cumulative_allocation" api:"required"`
	// The allocation per individual group
	GroupAllocation param.Field[string] `json:"group_allocation" api:"required"`
	// The event property used to group usage before applying allocations
	GroupingKey param.Field[string] `json:"grouping_key" api:"required"`
	// The amount to charge for each unit outside of the allocation
	UnitAmount param.Field[string] `json:"unit_amount" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceCumulativeGroupedAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType = "cumulative_grouped_allocation"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceModelTypeCumulativeGroupedAllocation:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanCumulativeGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type" api:"required"`
	// The name of the price.
	Name param.Field[string] `json:"name" api:"required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent" api:"required"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence" api:"required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config" api:"required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id" api:"required"`
	// The pricing model type
	ModelType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type" api:"required"`
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
	ConversionRateConfig param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
	// The ID of the license type to associate with this price.
	LicenseTypeID param.Field[string] `json:"license_type_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// A transient ID that can be used to reference this price when adding adjustments
	// in the same API call.
	ReferenceID param.Field[string] `json:"reference_id"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsBetaNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
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

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type" api:"required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                    `json:"unit_config"`
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type BetaNewPlanVersionParamsReplacePricesPriceModelType string

const (
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters                 BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_filters"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            BetaNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           BetaNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             BetaNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               BetaNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           BetaNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            BetaNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           BetaNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing BetaNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           BetaNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation     BetaNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_allocation"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite                BetaNewPlanVersionParamsReplacePricesPriceModelType = "minimum_composite"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypePercent                         BetaNewPlanVersionParamsReplacePricesPriceModelType = "percent"
	BetaNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput                     BetaNewPlanVersionParamsReplacePricesPriceModelType = "event_output"
)

func (r BetaNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedAllocation, BetaNewPlanVersionParamsReplacePricesPriceModelTypeMinimumComposite, BetaNewPlanVersionParamsReplacePricesPriceModelTypePercent, BetaNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaSetDefaultPlanVersionParams struct {
	// Plan version to set as the default.
	Version param.Field[int64] `json:"version" api:"required"`
}

func (r BetaSetDefaultPlanVersionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
