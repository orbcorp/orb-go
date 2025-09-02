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
	// The price to add to the plan
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
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
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
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
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence                           param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}]                                                                             `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                               `json:"item_id,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion] `json:"conversion_rate_config"`
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

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                          `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPrice].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID        param.Field[string]                                                                                        `json:"item_id,required"`
	MinimumConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceMinimumConfig] `json:"minimum_config,required"`
	ModelType     param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelType]     `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceMinimumConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// By default, subtotals from minimum composite prices are prorated based on the
	// service period. Set to false to disable proration.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelTypeMinimum BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceModelTypeMinimum:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                               `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePrice].
type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceNewPlanMinimumCompositePriceConversionRateTypeTiered:
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
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
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
	BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimum                         BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeGroupedTiered, BetaExternalPlanIDNewPlanVersionParamsAddPricesPriceModelTypeMinimum:
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
	// The price to add to the plan
	Price param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion] `json:"price"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
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
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice],
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
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPrice].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence                           param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}]                                                                                 `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                   `json:"item_id,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion] `json:"conversion_rate_config"`
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

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPrice].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID        param.Field[string]                                                                                            `json:"item_id,required"`
	MinimumConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceMinimumConfig] `json:"minimum_config,required"`
	ModelType     param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelType]     `json:"model_type,required"`
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
	ConversionRateConfig param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceUnion] `json:"conversion_rate_config"`
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

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceAnnual     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceSemiAnnual BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "semi_annual"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceMonthly    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "monthly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceQuarterly  BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "quarterly"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceOneTime    BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "one_time"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceCustom     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence = "custom"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceSemiAnnual, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceMonthly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceQuarterly, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceOneTime, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceMinimumConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// By default, subtotals from minimum composite prices are prorated based on the
	// service period. Set to false to disable proration.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelTypeMinimum BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceModelTypeMinimum:
		return true
	}
	return false
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice struct {
	ConversionRateType param.Field[BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                 `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                   `json:"unit_config"`
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice) ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePrice].
type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceUnion interface {
	ImplementsBetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceUnion()
}

type BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateType string

const (
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateTypeUnit   BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateType = "unit"
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateTypeTiered BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateType = "tiered"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceNewPlanMinimumCompositePriceConversionRateTypeTiered:
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
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
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
	BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimum                         BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType = "minimum"
)

func (r BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnit, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrix, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTiered, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeThresholdTotalAmount, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithPercent, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypePackageWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeUnitWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithDisplayName, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeBulkWithProration, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMatrixWithAllocation, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeGroupedTiered, BetaExternalPlanIDNewPlanVersionParamsReplacePricesPriceModelTypeMinimum:
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
