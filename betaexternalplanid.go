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

// This endpoint allows the creation of a new plan version for an existing plan.
func (r *BetaExternalPlanIDService) NewPlanVersion(ctx context.Context, externalPlanID string, body BetaExternalPlanIDNewPlanVersionParams, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalPlanID == "" {
		err = errors.New("missing required external_plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/external_plan_id/%s/versions", externalPlanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch a plan version. It returns the phases, prices,
// and adjustments present on this version of the plan.
func (r *BetaExternalPlanIDService) FetchPlanVersion(ctx context.Context, externalPlanID string, version string, opts ...option.RequestOption) (res *PlanVersion, err error) {
	opts = slices.Concat(r.Options, opts)
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

// This endpoint allows setting the default version of a plan.
func (r *BetaExternalPlanIDService) SetDefaultPlanVersion(ctx context.Context, externalPlanID string, body BetaExternalPlanIDSetDefaultPlanVersionParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = slices.Concat(r.Options, opts)
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustment].
type BetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddAdjustmentsAdjustmentUnion()
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
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New plan price request body params.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType] `json:"model_type,required"`
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
	// Configuration for bulk pricing
	BulkConfig              param.Field[shared.BulkConfigParam] `json:"bulk_config"`
	BulkWithFiltersConfig   param.Field[interface{}]            `json:"bulk_with_filters_config"`
	BulkWithProrationConfig param.Field[interface{}]            `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig        param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
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
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumConfig               param.Field[interface{}]                            `json:"minimum_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                              `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                   `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
	// The key in the event data to extract the unit rate from.
	UnitRatingKey param.Field[string] `json:"unit_rating_key,required"`
	// If provided, this amount will be used as the unit rate when an event does not
	// have a value for the `unit_rating_key`. If not provided, events missing a unit
	// rate will be ignored.
	DefaultUnitRate param.Field[string] `json:"default_unit_rate"`
	// An optional key in the event data to group by (e.g., event ID). All events will
	// also be grouped by their unit rate.
	GroupingKey param.Field[string] `json:"grouping_key"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit                            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered                          BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk                            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters                 BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_filters"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix                          BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "threshold_total_amount"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage                   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered                   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum        BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "package_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent                 BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_percent"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration             BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "tiered_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "unit_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_allocation"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration               BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "bulk_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "matrix_with_display_name"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage            BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "max_group_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk           BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimum                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "minimum"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePercent                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "percent"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeEventOutput                     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "event_output"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithFilters, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePercent, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeEventOutput:
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the plan.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustment].
type BetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplaceAdjustmentsAdjustmentUnion()
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
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// The phase to replace this price from.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New plan price request body params.
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New plan price request body params.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType] `json:"model_type,required"`
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
	// Configuration for bulk pricing
	BulkConfig              param.Field[shared.BulkConfigParam] `json:"bulk_config"`
	BulkWithFiltersConfig   param.Field[interface{}]            `json:"bulk_with_filters_config"`
	BulkWithProrationConfig param.Field[interface{}]            `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]     `json:"conversion_rate"`
	ConversionRateConfig        param.Field[interface{}] `json:"conversion_rate_config"`
	CumulativeGroupedBulkConfig param.Field[interface{}] `json:"cumulative_grouped_bulk_config"`
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
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig  param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig param.Field[interface{}]                            `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig param.Field[interface{}]                            `json:"max_group_tiered_package_config"`
	Metadata                    param.Field[interface{}]                            `json:"metadata"`
	MinimumConfig               param.Field[interface{}]                            `json:"minimum_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// New plan price request body params.
//
// Satisfied by [shared.NewPlanUnitPriceParam], [shared.NewPlanTieredPriceParam],
// [shared.NewPlanBulkPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice],
// [shared.NewPlanPackagePriceParam], [shared.NewPlanMatrixPriceParam],
// [shared.NewPlanThresholdTotalAmountPriceParam],
// [shared.NewPlanTieredPackagePriceParam],
// [shared.NewPlanTieredWithMinimumPriceParam],
// [shared.NewPlanGroupedTieredPriceParam],
// [shared.NewPlanTieredPackageWithMinimumPriceParam],
// [shared.NewPlanPackageWithAllocationPriceParam],
// [shared.NewPlanUnitWithPercentPriceParam],
// [shared.NewPlanMatrixWithAllocationPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice],
// [shared.NewPlanUnitWithProrationPriceParam],
// [shared.NewPlanGroupedAllocationPriceParam],
// [shared.NewPlanBulkWithProrationPriceParam],
// [shared.NewPlanGroupedWithProratedMinimumPriceParam],
// [shared.NewPlanGroupedWithMeteredMinimumPriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
// [shared.NewPlanMatrixWithDisplayNamePriceParam],
// [shared.NewPlanGroupedTieredPackagePriceParam],
// [shared.NewPlanMaxGroupTieredPackagePriceParam],
// [shared.NewPlanScalableMatrixWithUnitPricingPriceParam],
// [shared.NewPlanScalableMatrixWithTieredPricingPriceParam],
// [shared.NewPlanCumulativeGroupedBulkPriceParam],
// [shared.NewPlanMinimumCompositePriceParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                          `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                  `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType = "percent"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                       `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig struct {
	// The key in the event data to extract the unit rate from.
	UnitRatingKey param.Field[string] `json:"unit_rating_key,required"`
	// If provided, this amount will be used as the unit rate when an event does not
	// have a value for the `unit_rating_key`. If not provided, events missing a unit
	// rate will be ignored.
	DefaultUnitRate param.Field[string] `json:"default_unit_rate"`
	// An optional key in the event data to group by (e.g., event ID). All events will
	// also be grouped by their unit rate.
	GroupingKey param.Field[string] `json:"grouping_key"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType = "event_output"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                  `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfig].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit                            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered                          BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk                            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters                 BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_filters"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix                          BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "threshold_total_amount"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage                   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered                   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "package_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent                 BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_percent"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration             BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "tiered_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "unit_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_allocation"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration               BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "bulk_with_proration"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "matrix_with_display_name"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage            BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "max_group_tiered_package"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimum                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "minimum"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePercent                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "percent"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput                     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "event_output"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithFilters, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePercent, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeEventOutput:
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
