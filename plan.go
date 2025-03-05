// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
)

// PlanService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options        []option.RequestOption
	ExternalPlanID *PlanExternalPlanIDService
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	r.ExternalPlanID = NewPlanExternalPlanIDService(opts...)
	return
}

// This endpoint allows creation of plans including their prices.
func (r *PlanService) New(ctx context.Context, body PlanNewParams, opts ...option.RequestOption) (res *shared.PlanModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "plans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `external_plan_id`, and `metadata` of an
// existing plan.
//
// Other fields on a customer are currently immutable.
func (r *PlanService) Update(ctx context.Context, planID string, body PlanUpdateParams, opts ...option.RequestOption) (res *shared.PlanModel, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all [plans](/core-concepts#plan-and-price) for
// an account in a list format. The list of plans is ordered starting from the most
// recently created plan. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
func (r *PlanService) List(ctx context.Context, query PlanListParams, opts ...option.RequestOption) (res *pagination.Page[shared.PlanModel], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "plans"
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

// This endpoint returns a list of all [plans](/core-concepts#plan-and-price) for
// an account in a list format. The list of plans is ordered starting from the most
// recently created plan. The response also includes
// [`pagination_metadata`](/api-reference/pagination), which lets the caller
// retrieve the next page of results if they exist.
func (r *PlanService) ListAutoPaging(ctx context.Context, query PlanListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.PlanModel] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to fetch [plan](/core-concepts#plan-and-price) details
// given a plan identifier. It returns information about the prices included in the
// plan and their configuration, as well as the product that the plan is attached
// to.
//
// ## Serialized prices
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given [Price](/core-concepts#plan-and-price)
// object. The `model_type` field determines the key for the configuration object
// that is present. A detailed explanation of price types can be found in the
// [Price schema](/core-concepts#plan-and-price).
//
// ## Phases
//
// Orb supports plan phases, also known as contract ramps. For plans with phases,
// the serialized prices refer to all prices across all phases.
func (r *PlanService) Fetch(ctx context.Context, planID string, opts ...option.RequestOption) (res *shared.PlanModel, err error) {
	opts = append(r.Options[:], opts...)
	if planID == "" {
		err = errors.New("missing required plan_id parameter")
		return
	}
	path := fmt.Sprintf("plans/%s", planID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PlanNewParams struct {
	// An ISO 4217 currency string for invoices generated by subscriptions on this
	// plan.
	Currency param.Field[string] `json:"currency,required"`
	Name     param.Field[string] `json:"name,required"`
	// Prices for this plan. If the plan has phases, this includes prices across all
	// phases of the plan.
	Prices param.Field[[]PlanNewParamsPriceUnion] `json:"prices,required"`
	// Free-form text which is available on the invoice PDF and the Orb invoice portal.
	DefaultInvoiceMemo param.Field[string] `json:"default_invoice_memo"`
	ExternalPlanID     param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0.
	NetTerms param.Field[int64] `json:"net_terms"`
	// The status of the plan to create (either active or draft). If not specified,
	// this defaults to active.
	Status param.Field[PlanNewParamsStatus] `json:"status"`
}

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanNewParamsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	BpsConfig                 param.Field[shared.BpsConfigModelParam]                    `json:"bps_config"`
	BulkBpsConfig             param.Field[shared.BulkBpsConfigModelParam]                `json:"bulk_bps_config"`
	BulkConfig                param.Field[shared.BulkConfigModelParam]                   `json:"bulk_config"`
	BulkWithProrationConfig   param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"bulk_with_proration_config"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate              param.Field[float64]                                     `json:"conversion_rate"`
	CumulativeGroupedBulkConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"cumulative_grouped_bulk_config"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity               param.Field[float64]                                     `json:"fixed_price_quantity"`
	GroupedAllocationConfig          param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"grouped_allocation_config"`
	GroupedTieredPackageConfig       param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"grouped_tiered_package_config"`
	GroupedWithMeteredMinimumConfig  param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"grouped_with_metered_minimum_config"`
	GroupedWithProratedMinimumConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"grouped_with_prorated_minimum_config"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration           param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[shared.MatrixConfigModelParam]                 `json:"matrix_config"`
	MatrixWithDisplayNameConfig           param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}]                                   `json:"metadata"`
	PackageConfig                         param.Field[shared.PackageConfigModelParam]                `json:"package_config"`
	PackageWithAllocationConfig           param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"threshold_total_amount_config"`
	TieredBpsConfig                       param.Field[shared.TieredBpsConfigModelParam]              `json:"tiered_bps_config"`
	TieredConfig                          param.Field[shared.TieredConfigModelParam]                 `json:"tiered_config"`
	TieredPackageConfig                   param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"tiered_package_config"`
	TieredWithMinimumConfig               param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigModelParam]                   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[shared.CustomRatingFunctionConfigModelParam]   `json:"unit_with_proration_config"`
}

func (r PlanNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPrice) implementsPlanNewParamsPriceUnion() {}

// Satisfied by [PlanNewParamsPricesNewPlanUnitPrice],
// [PlanNewParamsPricesNewPlanPackagePrice],
// [PlanNewParamsPricesNewPlanMatrixPrice],
// [PlanNewParamsPricesNewPlanTieredPrice],
// [PlanNewParamsPricesNewPlanTieredBpsPrice],
// [PlanNewParamsPricesNewPlanBpsPrice], [PlanNewParamsPricesNewPlanBulkBpsPrice],
// [PlanNewParamsPricesNewPlanBulkPrice],
// [PlanNewParamsPricesNewPlanThresholdTotalAmountPrice],
// [PlanNewParamsPricesNewPlanTieredPackagePrice],
// [PlanNewParamsPricesNewPlanTieredWithMinimumPrice],
// [PlanNewParamsPricesNewPlanUnitWithPercentPrice],
// [PlanNewParamsPricesNewPlanPackageWithAllocationPrice],
// [PlanNewParamsPricesNewPlanTierWithProrationPrice],
// [PlanNewParamsPricesNewPlanUnitWithProrationPrice],
// [PlanNewParamsPricesNewPlanGroupedAllocationPrice],
// [PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice],
// [PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice],
// [PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice],
// [PlanNewParamsPricesNewPlanBulkWithProrationPrice],
// [PlanNewParamsPricesNewPlanGroupedTieredPackagePrice],
// [PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice],
// [PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPrice],
// [PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPrice],
// [PlanNewParamsPricesNewPlanCumulativeGroupedBulkPrice], [PlanNewParamsPrice].
type PlanNewParamsPriceUnion interface {
	implementsPlanNewParamsPriceUnion()
}

type PlanNewParamsPricesNewPlanUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                      `json:"name,required"`
	UnitConfig param.Field[shared.UnitConfigModelParam] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit PlanNewParamsPricesNewPlanUnitPriceModelType = "unit"
)

func (r PlanNewParamsPricesNewPlanUnitPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                         `json:"name,required"`
	PackageConfig param.Field[shared.PackageConfigModelParam] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanPackagePriceModelTypePackage PlanNewParamsPricesNewPlanPackagePriceModelType = "package"
)

func (r PlanNewParamsPricesNewPlanPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                         `json:"item_id,required"`
	MatrixConfig param.Field[shared.MatrixConfigModelParam]                  `json:"matrix_config,required"`
	ModelType    param.Field[PlanNewParamsPricesNewPlanMatrixPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMatrixPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMatrixPriceCadence string

const (
	PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual     PlanNewParamsPricesNewPlanMatrixPriceCadence = "annual"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMatrixPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly    PlanNewParamsPricesNewPlanMatrixPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly  PlanNewParamsPricesNewPlanMatrixPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime    PlanNewParamsPricesNewPlanMatrixPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMatrixPriceCadenceCustom     PlanNewParamsPricesNewPlanMatrixPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceCadenceAnnual, PlanNewParamsPricesNewPlanMatrixPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMatrixPriceCadenceMonthly, PlanNewParamsPricesNewPlanMatrixPriceCadenceQuarterly, PlanNewParamsPricesNewPlanMatrixPriceCadenceOneTime, PlanNewParamsPricesNewPlanMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixPriceModelType string

const (
	PlanNewParamsPricesNewPlanMatrixPriceModelTypeMatrix PlanNewParamsPricesNewPlanMatrixPriceModelType = "matrix"
)

func (r PlanNewParamsPricesNewPlanMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                        `json:"name,required"`
	TieredConfig param.Field[shared.TieredConfigModelParam] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredPriceModelTypeTiered PlanNewParamsPricesNewPlanTieredPriceModelType = "tiered"
)

func (r PlanNewParamsPricesNewPlanTieredPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                            `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                           `json:"name,required"`
	TieredBpsConfig param.Field[shared.TieredBpsConfigModelParam] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredBpsPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredBpsPriceModelTypeTieredBps PlanNewParamsPricesNewPlanTieredBpsPriceModelType = "tiered_bps"
)

func (r PlanNewParamsPricesNewPlanTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBpsPrice struct {
	BpsConfig param.Field[shared.BpsConfigModelParam] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBpsPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanBpsPriceModelTypeBps PlanNewParamsPricesNewPlanBpsPriceModelType = "bps"
)

func (r PlanNewParamsPricesNewPlanBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBpsPriceModelTypeBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkBpsPrice struct {
	BulkBpsConfig param.Field[shared.BulkBpsConfigModelParam] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                          `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkBpsPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkBpsPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkBpsPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkBpsPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkBpsPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkBpsPriceModelTypeBulkBps PlanNewParamsPricesNewPlanBulkBpsPriceModelType = "bulk_bps"
)

func (r PlanNewParamsPricesNewPlanBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkPrice struct {
	BulkConfig param.Field[shared.BulkConfigModelParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkPriceModelTypeBulk PlanNewParamsPricesNewPlanBulkPriceModelType = "bulk"
)

func (r PlanNewParamsPricesNewPlanBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                                      `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual     PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "annual"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly    PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly  PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime    PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceCustom     PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceAnnual, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceMonthly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceQuarterly, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceOneTime, PlanNewParamsPricesNewPlanThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType string

const (
	PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                                      `json:"name,required"`
	TieredPackageConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredPackagePriceModelTypeTieredPackage PlanNewParamsPricesNewPlanTieredPackagePriceModelType = "tiered_package"
)

func (r PlanNewParamsPricesNewPlanTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                                      `json:"name,required"`
	TieredWithMinimumConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                  `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                  param.Field[string]                                      `json:"name,required"`
	UnitWithPercentConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"unit_with_percent_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitWithPercentPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitWithPercentPriceModelTypeUnitWithPercent PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PlanNewParamsPricesNewPlanUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                                      `json:"name,required"`
	PackageWithAllocationConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual     PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly    PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime    PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceCustom     PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceAnnual, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceMonthly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceOneTime, PlanNewParamsPricesNewPlanPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType string

const (
	PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanTierWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                      param.Field[string]                                      `json:"name,required"`
	TieredWithProrationConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"tiered_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanTierWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanTierWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanTierWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanTierWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanTierWithProrationPriceModelTypeTieredWithProration PlanNewParamsPricesNewPlanTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r PlanNewParamsPricesNewPlanTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                                      `json:"name,required"`
	UnitWithProrationConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"unit_with_proration_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanUnitWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanUnitWithProrationPriceModelTypeUnitWithProration PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PlanNewParamsPricesNewPlanUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[shared.CustomRatingFunctionConfigModelParam]             `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedAllocationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedAllocationPriceModelTypeGroupedAllocation PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PlanNewParamsPricesNewPlanGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                      `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                         param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithMeteredMinimumConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                     `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                        `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                   `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceAnnual     PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "annual"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceMonthly    PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceQuarterly  PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceOneTime    PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceCustom     PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceAnnual, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceMonthly, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceQuarterly, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceOneTime, PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType string

const (
	PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanBulkWithProrationPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceAnnual     PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "annual"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceMonthly    PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceQuarterly  PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceOneTime    PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceCustom     PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceAnnual, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceMonthly, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceQuarterly, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceOneTime, PlanNewParamsPricesNewPlanBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType string

const (
	PlanNewParamsPricesNewPlanBulkWithProrationPriceModelTypeBulkWithProration PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PlanNewParamsPricesNewPlanBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence                    param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	GroupedTieredPackageConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                        `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                   `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceAnnual     PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "annual"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceMonthly    PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "monthly"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceQuarterly  PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceOneTime    PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "one_time"
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceCustom     PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceAnnual, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceMonthly, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceQuarterly, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceOneTime, PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType string

const (
	PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                param.Field[string]                                      `json:"name,required"`
	ScalableMatrixWithUnitPricingConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_unit_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence string

const (
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual     PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "annual"
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly    PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly  PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime    PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom     PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceAnnual, PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceMonthly, PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceQuarterly, PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceOneTime, PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelType string

const (
	PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                                  param.Field[string]                                      `json:"name,required"`
	ScalableMatrixWithTieredPricingConfig param.Field[shared.CustomRatingFunctionConfigModelParam] `json:"scalable_matrix_with_tiered_pricing_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPrice) implementsPlanNewParamsPriceUnion() {
}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence string

const (
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual     PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "annual"
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly    PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly  PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime    PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom     PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceAnnual, PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceMonthly, PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceQuarterly, PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceOneTime, PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelType string

const (
	PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[shared.CustomRatingFunctionConfigModelParam]                 `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                        `json:"item_id,required"`
	ModelType param.Field[PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency param.Field[string] `json:"currency"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationModelParam] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanNewParamsPricesNewPlanCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PlanNewParamsPricesNewPlanCumulativeGroupedBulkPrice) implementsPlanNewParamsPriceUnion() {}

// The cadence to bill for this price on.
type PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence string

const (
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceAnnual     PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "annual"
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "semi_annual"
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceMonthly    PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "monthly"
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceQuarterly  PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "quarterly"
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceOneTime    PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "one_time"
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceCustom     PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceAnnual, PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceSemiAnnual, PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceMonthly, PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceQuarterly, PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceOneTime, PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelType string

const (
	PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesNewPlanCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PlanNewParamsPricesCadence string

const (
	PlanNewParamsPricesCadenceAnnual     PlanNewParamsPricesCadence = "annual"
	PlanNewParamsPricesCadenceSemiAnnual PlanNewParamsPricesCadence = "semi_annual"
	PlanNewParamsPricesCadenceMonthly    PlanNewParamsPricesCadence = "monthly"
	PlanNewParamsPricesCadenceQuarterly  PlanNewParamsPricesCadence = "quarterly"
	PlanNewParamsPricesCadenceOneTime    PlanNewParamsPricesCadence = "one_time"
	PlanNewParamsPricesCadenceCustom     PlanNewParamsPricesCadence = "custom"
)

func (r PlanNewParamsPricesCadence) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesCadenceAnnual, PlanNewParamsPricesCadenceSemiAnnual, PlanNewParamsPricesCadenceMonthly, PlanNewParamsPricesCadenceQuarterly, PlanNewParamsPricesCadenceOneTime, PlanNewParamsPricesCadenceCustom:
		return true
	}
	return false
}

type PlanNewParamsPricesModelType string

const (
	PlanNewParamsPricesModelTypeUnit                            PlanNewParamsPricesModelType = "unit"
	PlanNewParamsPricesModelTypePackage                         PlanNewParamsPricesModelType = "package"
	PlanNewParamsPricesModelTypeMatrix                          PlanNewParamsPricesModelType = "matrix"
	PlanNewParamsPricesModelTypeTiered                          PlanNewParamsPricesModelType = "tiered"
	PlanNewParamsPricesModelTypeTieredBps                       PlanNewParamsPricesModelType = "tiered_bps"
	PlanNewParamsPricesModelTypeBps                             PlanNewParamsPricesModelType = "bps"
	PlanNewParamsPricesModelTypeBulkBps                         PlanNewParamsPricesModelType = "bulk_bps"
	PlanNewParamsPricesModelTypeBulk                            PlanNewParamsPricesModelType = "bulk"
	PlanNewParamsPricesModelTypeThresholdTotalAmount            PlanNewParamsPricesModelType = "threshold_total_amount"
	PlanNewParamsPricesModelTypeTieredPackage                   PlanNewParamsPricesModelType = "tiered_package"
	PlanNewParamsPricesModelTypeTieredWithMinimum               PlanNewParamsPricesModelType = "tiered_with_minimum"
	PlanNewParamsPricesModelTypeUnitWithPercent                 PlanNewParamsPricesModelType = "unit_with_percent"
	PlanNewParamsPricesModelTypePackageWithAllocation           PlanNewParamsPricesModelType = "package_with_allocation"
	PlanNewParamsPricesModelTypeTieredWithProration             PlanNewParamsPricesModelType = "tiered_with_proration"
	PlanNewParamsPricesModelTypeUnitWithProration               PlanNewParamsPricesModelType = "unit_with_proration"
	PlanNewParamsPricesModelTypeGroupedAllocation               PlanNewParamsPricesModelType = "grouped_allocation"
	PlanNewParamsPricesModelTypeGroupedWithProratedMinimum      PlanNewParamsPricesModelType = "grouped_with_prorated_minimum"
	PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum       PlanNewParamsPricesModelType = "grouped_with_metered_minimum"
	PlanNewParamsPricesModelTypeMatrixWithDisplayName           PlanNewParamsPricesModelType = "matrix_with_display_name"
	PlanNewParamsPricesModelTypeBulkWithProration               PlanNewParamsPricesModelType = "bulk_with_proration"
	PlanNewParamsPricesModelTypeGroupedTieredPackage            PlanNewParamsPricesModelType = "grouped_tiered_package"
	PlanNewParamsPricesModelTypeMaxGroupTieredPackage           PlanNewParamsPricesModelType = "max_group_tiered_package"
	PlanNewParamsPricesModelTypeScalableMatrixWithUnitPricing   PlanNewParamsPricesModelType = "scalable_matrix_with_unit_pricing"
	PlanNewParamsPricesModelTypeScalableMatrixWithTieredPricing PlanNewParamsPricesModelType = "scalable_matrix_with_tiered_pricing"
	PlanNewParamsPricesModelTypeCumulativeGroupedBulk           PlanNewParamsPricesModelType = "cumulative_grouped_bulk"
)

func (r PlanNewParamsPricesModelType) IsKnown() bool {
	switch r {
	case PlanNewParamsPricesModelTypeUnit, PlanNewParamsPricesModelTypePackage, PlanNewParamsPricesModelTypeMatrix, PlanNewParamsPricesModelTypeTiered, PlanNewParamsPricesModelTypeTieredBps, PlanNewParamsPricesModelTypeBps, PlanNewParamsPricesModelTypeBulkBps, PlanNewParamsPricesModelTypeBulk, PlanNewParamsPricesModelTypeThresholdTotalAmount, PlanNewParamsPricesModelTypeTieredPackage, PlanNewParamsPricesModelTypeTieredWithMinimum, PlanNewParamsPricesModelTypeUnitWithPercent, PlanNewParamsPricesModelTypePackageWithAllocation, PlanNewParamsPricesModelTypeTieredWithProration, PlanNewParamsPricesModelTypeUnitWithProration, PlanNewParamsPricesModelTypeGroupedAllocation, PlanNewParamsPricesModelTypeGroupedWithProratedMinimum, PlanNewParamsPricesModelTypeGroupedWithMeteredMinimum, PlanNewParamsPricesModelTypeMatrixWithDisplayName, PlanNewParamsPricesModelTypeBulkWithProration, PlanNewParamsPricesModelTypeGroupedTieredPackage, PlanNewParamsPricesModelTypeMaxGroupTieredPackage, PlanNewParamsPricesModelTypeScalableMatrixWithUnitPricing, PlanNewParamsPricesModelTypeScalableMatrixWithTieredPricing, PlanNewParamsPricesModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

// The status of the plan to create (either active or draft). If not specified,
// this defaults to active.
type PlanNewParamsStatus string

const (
	PlanNewParamsStatusActive PlanNewParamsStatus = "active"
	PlanNewParamsStatusDraft  PlanNewParamsStatus = "draft"
)

func (r PlanNewParamsStatus) IsKnown() bool {
	switch r {
	case PlanNewParamsStatusActive, PlanNewParamsStatusDraft:
		return true
	}
	return false
}

type PlanUpdateParams struct {
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PlanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// The plan status to filter to ('active', 'archived', or 'draft').
	Status param.Field[PlanListParamsStatus] `query:"status"`
}

// URLQuery serializes [PlanListParams]'s query parameters as `url.Values`.
func (r PlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The plan status to filter to ('active', 'archived', or 'draft').
type PlanListParamsStatus string

const (
	PlanListParamsStatusActive   PlanListParamsStatus = "active"
	PlanListParamsStatusArchived PlanListParamsStatus = "archived"
	PlanListParamsStatusDraft    PlanListParamsStatus = "draft"
)

func (r PlanListParamsStatus) IsKnown() bool {
	switch r {
	case PlanListParamsStatusActive, PlanListParamsStatusArchived, PlanListParamsStatusDraft:
		return true
	}
	return false
}
