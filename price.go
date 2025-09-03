// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/packages/pagination"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// PriceService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPriceService] method instead.
type PriceService struct {
	Options         []option.RequestOption
	ExternalPriceID *PriceExternalPriceIDService
}

// NewPriceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPriceService(opts ...option.RequestOption) (r *PriceService) {
	r = &PriceService{}
	r.Options = opts
	r.ExternalPriceID = NewPriceExternalPriceIDService(opts...)
	return
}

// This endpoint is used to create a [price](/product-catalog/price-configuration).
// A price created using this endpoint is always an add-on, meaning that it's not
// associated with a specific plan and can instead be individually added to
// subscriptions, including subscriptions on different plans.
//
// An `external_price_id` can be optionally specified as an alias to allow
// ergonomic interaction with prices in the Orb API.
//
// See the [Price resource](/product-catalog/price-configuration) for the
// specification of different price model configurations possible in this endpoint.
func (r *PriceService) New(ctx context.Context, body PriceNewParams, opts ...option.RequestOption) (res *shared.Price, err error) {
	opts = append(r.Options[:], opts...)
	path := "prices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to update the `metadata` property on a price. If you
// pass null for the metadata value, it will clear any existing metadata for that
// price.
func (r *PriceService) Update(ctx context.Context, priceID string, body PriceUpdateParams, opts ...option.RequestOption) (res *shared.Price, err error) {
	opts = append(r.Options[:], opts...)
	if priceID == "" {
		err = errors.New("missing required price_id parameter")
		return
	}
	path := fmt.Sprintf("prices/%s", priceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint is used to list all add-on prices created using the
// [price creation endpoint](/api-reference/price/create-price).
func (r *PriceService) List(ctx context.Context, query PriceListParams, opts ...option.RequestOption) (res *pagination.Page[shared.Price], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "prices"
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

// This endpoint is used to list all add-on prices created using the
// [price creation endpoint](/api-reference/price/create-price).
func (r *PriceService) ListAutoPaging(ctx context.Context, query PriceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.Price] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// [NOTE] It is recommended to use the `/v1/prices/evaluate` which offers further
// functionality, such as multiple prices, inline price definitions, and querying
// over preview events.
//
// This endpoint is used to evaluate the output of a price for a given customer and
// time range. It enables filtering and grouping the output using
// [computed properties](/extensibility/advanced-metrics#computed-properties),
// supporting the following workflows:
//
// 1. Showing detailed usage and costs to the end customer.
// 2. Auditing subtotals on invoice line items.
//
// For these workflows, the expressiveness of computed properties in both the
// filters and grouping is critical. For example, if you'd like to show your
// customer their usage grouped by hour and another property, you can do so with
// the following `grouping_keys`:
// `["hour_floor_timestamp_millis(timestamp_millis)", "my_property"]`. If you'd
// like to examine a customer's usage for a specific property value, you can do so
// with the following `filter`:
// `my_property = 'foo' AND my_other_property = 'bar'`.
//
// By default, the start of the time range must be no more than 100 days ago and
// the length of the results must be no greater than 1000. Note that this is a POST
// endpoint rather than a GET endpoint because it employs a JSON body rather than
// query parameters.
func (r *PriceService) Evaluate(ctx context.Context, priceID string, body PriceEvaluateParams, opts ...option.RequestOption) (res *PriceEvaluateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if priceID == "" {
		err = errors.New("missing required price_id parameter")
		return
	}
	path := fmt.Sprintf("prices/%s/evaluate", priceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to evaluate the output of price(s) for a given customer
// and time range over ingested events. It enables filtering and grouping the
// output using
// [computed properties](/extensibility/advanced-metrics#computed-properties),
// supporting the following workflows:
//
// 1. Showing detailed usage and costs to the end customer.
// 2. Auditing subtotals on invoice line items.
//
// For these workflows, the expressiveness of computed properties in both the
// filters and grouping is critical. For example, if you'd like to show your
// customer their usage grouped by hour and another property, you can do so with
// the following `grouping_keys`:
// `["hour_floor_timestamp_millis(timestamp_millis)", "my_property"]`. If you'd
// like to examine a customer's usage for a specific property value, you can do so
// with the following `filter`:
// `my_property = 'foo' AND my_other_property = 'bar'`.
//
// Prices may either reference existing prices in your Orb account or be defined
// inline in the request body. Up to 100 prices can be evaluated in a single
// request.
//
// Prices are evaluated on ingested events and the start of the time range must be
// no more than 100 days ago. To evaluate based off a set of provided events, the
// [evaluate preview events](/api-reference/price/evaluate-preview-events) endpoint
// can be used instead.
//
// Note that this is a POST endpoint rather than a GET endpoint because it employs
// a JSON body rather than query parameters.
func (r *PriceService) EvaluateMultiple(ctx context.Context, body PriceEvaluateMultipleParams, opts ...option.RequestOption) (res *PriceEvaluateMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "prices/evaluate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint evaluates prices on preview events instead of actual usage, making
// it ideal for building price calculators and cost estimation tools. You can
// filter and group results using
// [computed properties](/extensibility/advanced-metrics#computed-properties) to
// analyze pricing across different dimensions.
//
// Prices may either reference existing prices in your Orb account or be defined
// inline in the request body. The endpoint has the following limitations:
//
// 1. Up to 100 prices can be evaluated in a single request.
// 2. Up to 500 preview events can be provided in a single request.
//
// A top-level customer_id is required to evaluate the preview events.
// Additionally, all events without a customer_id will have the top-level
// customer_id added.
//
// Note that this is a POST endpoint rather than a GET endpoint because it employs
// a JSON body rather than query parameters.
func (r *PriceService) EvaluatePreviewEvents(ctx context.Context, body PriceEvaluatePreviewEventsParams, opts ...option.RequestOption) (res *PriceEvaluatePreviewEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "prices/evaluate_preview_events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint returns a price given an identifier.
func (r *PriceService) Fetch(ctx context.Context, priceID string, opts ...option.RequestOption) (res *shared.Price, err error) {
	opts = append(r.Options[:], opts...)
	if priceID == "" {
		err = errors.New("missing required price_id parameter")
		return
	}
	path := fmt.Sprintf("prices/%s", priceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EvaluatePriceGroup struct {
	// The price's output for the group
	Amount string `json:"amount,required"`
	// The values for the group in the order specified by `grouping_keys`
	GroupingValues []EvaluatePriceGroupGroupingValuesUnion `json:"grouping_values,required"`
	// The price's usage quantity for the group
	Quantity float64                `json:"quantity,required"`
	JSON     evaluatePriceGroupJSON `json:"-"`
}

// evaluatePriceGroupJSON contains the JSON metadata for the struct
// [EvaluatePriceGroup]
type evaluatePriceGroupJSON struct {
	Amount         apijson.Field
	GroupingValues apijson.Field
	Quantity       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EvaluatePriceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluatePriceGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluatePriceGroupGroupingValuesUnion interface {
	ImplementsEvaluatePriceGroupGroupingValuesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluatePriceGroupGroupingValuesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type PriceEvaluateResponse struct {
	Data []EvaluatePriceGroup      `json:"data,required"`
	JSON priceEvaluateResponseJSON `json:"-"`
}

// priceEvaluateResponseJSON contains the JSON metadata for the struct
// [PriceEvaluateResponse]
type priceEvaluateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceEvaluateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEvaluateResponseJSON) RawJSON() string {
	return r.raw
}

type PriceEvaluateMultipleResponse struct {
	Data []PriceEvaluateMultipleResponseData `json:"data,required"`
	JSON priceEvaluateMultipleResponseJSON   `json:"-"`
}

// priceEvaluateMultipleResponseJSON contains the JSON metadata for the struct
// [PriceEvaluateMultipleResponse]
type priceEvaluateMultipleResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceEvaluateMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEvaluateMultipleResponseJSON) RawJSON() string {
	return r.raw
}

type PriceEvaluateMultipleResponseData struct {
	// The currency of the price
	Currency string `json:"currency,required"`
	// The computed price groups associated with input price.
	PriceGroups []EvaluatePriceGroup `json:"price_groups,required"`
	// The external ID of the price
	ExternalPriceID string `json:"external_price_id,nullable"`
	// The index of the inline price
	InlinePriceIndex int64 `json:"inline_price_index,nullable"`
	// The ID of the price
	PriceID string                                `json:"price_id,nullable"`
	JSON    priceEvaluateMultipleResponseDataJSON `json:"-"`
}

// priceEvaluateMultipleResponseDataJSON contains the JSON metadata for the struct
// [PriceEvaluateMultipleResponseData]
type priceEvaluateMultipleResponseDataJSON struct {
	Currency         apijson.Field
	PriceGroups      apijson.Field
	ExternalPriceID  apijson.Field
	InlinePriceIndex apijson.Field
	PriceID          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PriceEvaluateMultipleResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEvaluateMultipleResponseDataJSON) RawJSON() string {
	return r.raw
}

type PriceEvaluatePreviewEventsResponse struct {
	Data []PriceEvaluatePreviewEventsResponseData `json:"data,required"`
	JSON priceEvaluatePreviewEventsResponseJSON   `json:"-"`
}

// priceEvaluatePreviewEventsResponseJSON contains the JSON metadata for the struct
// [PriceEvaluatePreviewEventsResponse]
type priceEvaluatePreviewEventsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceEvaluatePreviewEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEvaluatePreviewEventsResponseJSON) RawJSON() string {
	return r.raw
}

type PriceEvaluatePreviewEventsResponseData struct {
	// The currency of the price
	Currency string `json:"currency,required"`
	// The computed price groups associated with input price.
	PriceGroups []EvaluatePriceGroup `json:"price_groups,required"`
	// The external ID of the price
	ExternalPriceID string `json:"external_price_id,nullable"`
	// The index of the inline price
	InlinePriceIndex int64 `json:"inline_price_index,nullable"`
	// The ID of the price
	PriceID string                                     `json:"price_id,nullable"`
	JSON    priceEvaluatePreviewEventsResponseDataJSON `json:"-"`
}

// priceEvaluatePreviewEventsResponseDataJSON contains the JSON metadata for the
// struct [PriceEvaluatePreviewEventsResponseData]
type priceEvaluatePreviewEventsResponseDataJSON struct {
	Currency         apijson.Field
	PriceGroups      apijson.Field
	ExternalPriceID  apijson.Field
	InlinePriceIndex apijson.Field
	PriceID          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PriceEvaluatePreviewEventsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceEvaluatePreviewEventsResponseDataJSON) RawJSON() string {
	return r.raw
}

// This interface is a union satisfied by one of the following:
// [PriceNewParamsNewFloatingUnitPrice], [PriceNewParamsNewFloatingPackagePrice],
// [PriceNewParamsNewFloatingMatrixPrice],
// [PriceNewParamsNewFloatingMatrixWithAllocationPrice],
// [PriceNewParamsNewFloatingTieredPrice], [PriceNewParamsNewFloatingBulkPrice],
// [PriceNewParamsNewFloatingThresholdTotalAmountPrice],
// [PriceNewParamsNewFloatingTieredPackagePrice],
// [PriceNewParamsNewFloatingGroupedTieredPrice],
// [PriceNewParamsNewFloatingMaxGroupTieredPackagePrice],
// [PriceNewParamsNewFloatingTieredWithMinimumPrice],
// [PriceNewParamsNewFloatingPackageWithAllocationPrice],
// [PriceNewParamsNewFloatingTieredPackageWithMinimumPrice],
// [PriceNewParamsNewFloatingUnitWithPercentPrice],
// [PriceNewParamsNewFloatingTieredWithProrationPrice],
// [PriceNewParamsNewFloatingUnitWithProrationPrice],
// [PriceNewParamsNewFloatingGroupedAllocationPrice],
// [PriceNewParamsNewFloatingGroupedWithProratedMinimumPrice],
// [PriceNewParamsNewFloatingGroupedWithMeteredMinimumPrice],
// [PriceNewParamsNewFloatingMatrixWithDisplayNamePrice],
// [PriceNewParamsNewFloatingBulkWithProrationPrice],
// [PriceNewParamsNewFloatingGroupedTieredPackagePrice],
// [PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPrice],
// [PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPrice],
// [PriceNewParamsNewFloatingCumulativeGroupedBulkPrice],
// [PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPrice],
// [PriceNewParamsNewFloatingMinimumCompositePrice].
type PriceNewParams interface {
	ImplementsPriceNewParams()
}

type PriceNewParamsNewFloatingUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                 `json:"name,required"`
	UnitConfig param.Field[shared.UnitConfigParam] `json:"unit_config,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingUnitPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingUnitPriceCadence string

const (
	PriceNewParamsNewFloatingUnitPriceCadenceAnnual     PriceNewParamsNewFloatingUnitPriceCadence = "annual"
	PriceNewParamsNewFloatingUnitPriceCadenceSemiAnnual PriceNewParamsNewFloatingUnitPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingUnitPriceCadenceMonthly    PriceNewParamsNewFloatingUnitPriceCadence = "monthly"
	PriceNewParamsNewFloatingUnitPriceCadenceQuarterly  PriceNewParamsNewFloatingUnitPriceCadence = "quarterly"
	PriceNewParamsNewFloatingUnitPriceCadenceOneTime    PriceNewParamsNewFloatingUnitPriceCadence = "one_time"
	PriceNewParamsNewFloatingUnitPriceCadenceCustom     PriceNewParamsNewFloatingUnitPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingUnitPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitPriceCadenceAnnual, PriceNewParamsNewFloatingUnitPriceCadenceSemiAnnual, PriceNewParamsNewFloatingUnitPriceCadenceMonthly, PriceNewParamsNewFloatingUnitPriceCadenceQuarterly, PriceNewParamsNewFloatingUnitPriceCadenceOneTime, PriceNewParamsNewFloatingUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitPriceModelType string

const (
	PriceNewParamsNewFloatingUnitPriceModelTypeUnit PriceNewParamsNewFloatingUnitPriceModelType = "unit"
)

func (r PriceNewParamsNewFloatingUnitPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitPriceConversionRateConfig struct {
	ConversionRateType param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                     `json:"unit_config"`
}

func (r PriceNewParamsNewFloatingUnitPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceNewParamsNewFloatingUnitPriceConversionRateConfig) ImplementsPriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PriceNewParamsNewFloatingUnitPriceConversionRateConfig].
type PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion interface {
	ImplementsPriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion()
}

type PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateType string

const (
	PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateTypeUnit   PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateType = "unit"
	PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateTypeTiered PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateType = "tiered"
)

func (r PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateTypeUnit, PriceNewParamsNewFloatingUnitPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                    `json:"name,required"`
	PackageConfig param.Field[shared.PackageConfigParam] `json:"package_config,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingPackagePriceCadence string

const (
	PriceNewParamsNewFloatingPackagePriceCadenceAnnual     PriceNewParamsNewFloatingPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingPackagePriceCadenceSemiAnnual PriceNewParamsNewFloatingPackagePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingPackagePriceCadenceMonthly    PriceNewParamsNewFloatingPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingPackagePriceCadenceQuarterly  PriceNewParamsNewFloatingPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingPackagePriceCadenceOneTime    PriceNewParamsNewFloatingPackagePriceCadence = "one_time"
	PriceNewParamsNewFloatingPackagePriceCadenceCustom     PriceNewParamsNewFloatingPackagePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackagePriceCadenceAnnual, PriceNewParamsNewFloatingPackagePriceCadenceSemiAnnual, PriceNewParamsNewFloatingPackagePriceCadenceMonthly, PriceNewParamsNewFloatingPackagePriceCadenceQuarterly, PriceNewParamsNewFloatingPackagePriceCadenceOneTime, PriceNewParamsNewFloatingPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackagePriceModelType string

const (
	PriceNewParamsNewFloatingPackagePriceModelTypePackage PriceNewParamsNewFloatingPackagePriceModelType = "package"
)

func (r PriceNewParamsNewFloatingPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackagePriceModelTypePackage:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID       param.Field[string]                                        `json:"item_id,required"`
	MatrixConfig param.Field[shared.MatrixConfigParam]                      `json:"matrix_config,required"`
	ModelType    param.Field[PriceNewParamsNewFloatingMatrixPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMatrixPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMatrixPriceCadence string

const (
	PriceNewParamsNewFloatingMatrixPriceCadenceAnnual     PriceNewParamsNewFloatingMatrixPriceCadence = "annual"
	PriceNewParamsNewFloatingMatrixPriceCadenceSemiAnnual PriceNewParamsNewFloatingMatrixPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingMatrixPriceCadenceMonthly    PriceNewParamsNewFloatingMatrixPriceCadence = "monthly"
	PriceNewParamsNewFloatingMatrixPriceCadenceQuarterly  PriceNewParamsNewFloatingMatrixPriceCadence = "quarterly"
	PriceNewParamsNewFloatingMatrixPriceCadenceOneTime    PriceNewParamsNewFloatingMatrixPriceCadence = "one_time"
	PriceNewParamsNewFloatingMatrixPriceCadenceCustom     PriceNewParamsNewFloatingMatrixPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixPriceCadenceAnnual, PriceNewParamsNewFloatingMatrixPriceCadenceSemiAnnual, PriceNewParamsNewFloatingMatrixPriceCadenceMonthly, PriceNewParamsNewFloatingMatrixPriceCadenceQuarterly, PriceNewParamsNewFloatingMatrixPriceCadenceOneTime, PriceNewParamsNewFloatingMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixPriceModelType string

const (
	PriceNewParamsNewFloatingMatrixPriceModelTypeMatrix PriceNewParamsNewFloatingMatrixPriceModelType = "matrix"
)

func (r PriceNewParamsNewFloatingMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                     param.Field[string]                                                      `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[shared.MatrixWithAllocationConfigParam]                      `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMatrixWithAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceAnnual     PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "annual"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceMonthly    PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "monthly"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceQuarterly  PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceOneTime    PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "one_time"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceCustom     PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceAnnual, PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceMonthly, PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceQuarterly, PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceOneTime, PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                        `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                   `json:"name,required"`
	TieredConfig param.Field[shared.TieredConfigParam] `json:"tiered_config,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPriceCadence string

const (
	PriceNewParamsNewFloatingTieredPriceCadenceAnnual     PriceNewParamsNewFloatingTieredPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredPriceCadenceMonthly    PriceNewParamsNewFloatingTieredPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPriceCadenceQuarterly  PriceNewParamsNewFloatingTieredPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPriceCadenceOneTime    PriceNewParamsNewFloatingTieredPriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredPriceCadenceCustom     PriceNewParamsNewFloatingTieredPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPriceCadenceAnnual, PriceNewParamsNewFloatingTieredPriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredPriceCadenceMonthly, PriceNewParamsNewFloatingTieredPriceCadenceQuarterly, PriceNewParamsNewFloatingTieredPriceCadenceOneTime, PriceNewParamsNewFloatingTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPriceModelType string

const (
	PriceNewParamsNewFloatingTieredPriceModelTypeTiered PriceNewParamsNewFloatingTieredPriceModelType = "tiered"
)

func (r PriceNewParamsNewFloatingTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkPrice struct {
	BulkConfig param.Field[shared.BulkConfigParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingBulkPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingBulkPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingBulkPriceCadence string

const (
	PriceNewParamsNewFloatingBulkPriceCadenceAnnual     PriceNewParamsNewFloatingBulkPriceCadence = "annual"
	PriceNewParamsNewFloatingBulkPriceCadenceSemiAnnual PriceNewParamsNewFloatingBulkPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingBulkPriceCadenceMonthly    PriceNewParamsNewFloatingBulkPriceCadence = "monthly"
	PriceNewParamsNewFloatingBulkPriceCadenceQuarterly  PriceNewParamsNewFloatingBulkPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBulkPriceCadenceOneTime    PriceNewParamsNewFloatingBulkPriceCadence = "one_time"
	PriceNewParamsNewFloatingBulkPriceCadenceCustom     PriceNewParamsNewFloatingBulkPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkPriceCadenceAnnual, PriceNewParamsNewFloatingBulkPriceCadenceSemiAnnual, PriceNewParamsNewFloatingBulkPriceCadenceMonthly, PriceNewParamsNewFloatingBulkPriceCadenceQuarterly, PriceNewParamsNewFloatingBulkPriceCadenceOneTime, PriceNewParamsNewFloatingBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkPriceModelType string

const (
	PriceNewParamsNewFloatingBulkPriceModelTypeBulk PriceNewParamsNewFloatingBulkPriceModelType = "bulk"
)

func (r PriceNewParamsNewFloatingBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingThresholdTotalAmountPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceAnnual     PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "annual"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceMonthly    PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "monthly"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceQuarterly  PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "quarterly"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceOneTime    PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "one_time"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceCustom     PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceAnnual, PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual, PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceMonthly, PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceQuarterly, PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceOneTime, PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                               `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPackagePriceCadence string

const (
	PriceNewParamsNewFloatingTieredPackagePriceCadenceAnnual     PriceNewParamsNewFloatingTieredPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredPackagePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceMonthly    PriceNewParamsNewFloatingTieredPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceQuarterly  PriceNewParamsNewFloatingTieredPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceOneTime    PriceNewParamsNewFloatingTieredPackagePriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceCustom     PriceNewParamsNewFloatingTieredPackagePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackagePriceCadenceAnnual, PriceNewParamsNewFloatingTieredPackagePriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredPackagePriceCadenceMonthly, PriceNewParamsNewFloatingTieredPackagePriceCadenceQuarterly, PriceNewParamsNewFloatingTieredPackagePriceCadenceOneTime, PriceNewParamsNewFloatingTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackagePriceModelType string

const (
	PriceNewParamsNewFloatingTieredPackagePriceModelTypeTieredPackage PriceNewParamsNewFloatingTieredPackagePriceModelType = "tiered_package"
)

func (r PriceNewParamsNewFloatingTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency            param.Field[string]                 `json:"currency,required"`
	GroupedTieredConfig param.Field[map[string]interface{}] `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                               `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedTieredPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedTieredPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedTieredPriceCadence string

const (
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceAnnual     PriceNewParamsNewFloatingGroupedTieredPriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedTieredPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceMonthly    PriceNewParamsNewFloatingGroupedTieredPriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedTieredPriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceOneTime    PriceNewParamsNewFloatingGroupedTieredPriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedTieredPriceCadenceCustom     PriceNewParamsNewFloatingGroupedTieredPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPriceCadenceAnnual, PriceNewParamsNewFloatingGroupedTieredPriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedTieredPriceCadenceMonthly, PriceNewParamsNewFloatingGroupedTieredPriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedTieredPriceCadenceOneTime, PriceNewParamsNewFloatingGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedTieredPriceModelType string

const (
	PriceNewParamsNewFloatingGroupedTieredPriceModelTypeGroupedTiered PriceNewParamsNewFloatingGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PriceNewParamsNewFloatingGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMaxGroupTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                       `json:"item_id,required"`
	MaxGroupTieredPackageConfig param.Field[map[string]interface{}]                                       `json:"max_group_tiered_package_config,required"`
	ModelType                   param.Field[PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingMaxGroupTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMaxGroupTieredPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence string

const (
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceAnnual     PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceMonthly    PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly  PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceOneTime    PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "one_time"
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceCustom     PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceAnnual, PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceSemiAnnual, PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceMonthly, PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceQuarterly, PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceOneTime, PriceNewParamsNewFloatingMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelType string

const (
	PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                   `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredWithMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredWithMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceAnnual     PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceMonthly    PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceQuarterly  PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceOneTime    PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceCustom     PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceAnnual, PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceMonthly, PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceQuarterly, PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceOneTime, PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum PriceNewParamsNewFloatingTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingPackageWithAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingPackageWithAllocationPriceCadence string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceAnnual     PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "annual"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceSemiAnnual PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceMonthly    PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "monthly"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceQuarterly  PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceOneTime    PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "one_time"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceCustom     PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceAnnual, PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceMonthly, PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceQuarterly, PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceOneTime, PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackageWithAllocationPriceModelType string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation PriceNewParamsNewFloatingPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                          `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPackageWithMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceAnnual     PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceMonthly    PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly  PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceOneTime    PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceCustom     PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceAnnual, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceMonthly, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceOneTime, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                 `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingUnitWithPercentPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingUnitWithPercentPriceCadence string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceAnnual     PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "annual"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceSemiAnnual PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceMonthly    PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "monthly"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceQuarterly  PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "quarterly"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceOneTime    PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "one_time"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceCustom     PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithPercentPriceCadenceAnnual, PriceNewParamsNewFloatingUnitWithPercentPriceCadenceSemiAnnual, PriceNewParamsNewFloatingUnitWithPercentPriceCadenceMonthly, PriceNewParamsNewFloatingUnitWithPercentPriceCadenceQuarterly, PriceNewParamsNewFloatingUnitWithPercentPriceCadenceOneTime, PriceNewParamsNewFloatingUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithPercentPriceModelType string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent PriceNewParamsNewFloatingUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PriceNewParamsNewFloatingUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                     `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredWithProrationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredWithProrationPriceCadence string

const (
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceAnnual     PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceMonthly    PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceQuarterly  PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceOneTime    PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredWithProrationPriceCadenceCustom     PriceNewParamsNewFloatingTieredWithProrationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithProrationPriceCadenceAnnual, PriceNewParamsNewFloatingTieredWithProrationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredWithProrationPriceCadenceMonthly, PriceNewParamsNewFloatingTieredWithProrationPriceCadenceQuarterly, PriceNewParamsNewFloatingTieredWithProrationPriceCadenceOneTime, PriceNewParamsNewFloatingTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithProrationPriceModelType string

const (
	PriceNewParamsNewFloatingTieredWithProrationPriceModelTypeTieredWithProration PriceNewParamsNewFloatingTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PriceNewParamsNewFloatingTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                   `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingUnitWithProrationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingUnitWithProrationPriceCadence string

const (
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceAnnual     PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "annual"
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceSemiAnnual PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceMonthly    PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "monthly"
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceQuarterly  PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceOneTime    PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "one_time"
	PriceNewParamsNewFloatingUnitWithProrationPriceCadenceCustom     PriceNewParamsNewFloatingUnitWithProrationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithProrationPriceCadenceAnnual, PriceNewParamsNewFloatingUnitWithProrationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingUnitWithProrationPriceCadenceMonthly, PriceNewParamsNewFloatingUnitWithProrationPriceCadenceQuarterly, PriceNewParamsNewFloatingUnitWithProrationPriceCadenceOneTime, PriceNewParamsNewFloatingUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithProrationPriceModelType string

const (
	PriceNewParamsNewFloatingUnitWithProrationPriceModelTypeUnitWithProration PriceNewParamsNewFloatingUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PriceNewParamsNewFloatingUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                param.Field[string]                 `json:"currency,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}] `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                   `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedAllocationPriceCadence string

const (
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceAnnual     PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceMonthly    PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceOneTime    PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedAllocationPriceCadenceCustom     PriceNewParamsNewFloatingGroupedAllocationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedAllocationPriceCadenceAnnual, PriceNewParamsNewFloatingGroupedAllocationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedAllocationPriceCadenceMonthly, PriceNewParamsNewFloatingGroupedAllocationPriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedAllocationPriceCadenceOneTime, PriceNewParamsNewFloatingGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedAllocationPriceModelType string

const (
	PriceNewParamsNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation PriceNewParamsNewFloatingGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PriceNewParamsNewFloatingGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                         param.Field[string]                 `json:"currency,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                            `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedWithProratedMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual     PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly    PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime    PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceCustom     PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                        param.Field[string]                 `json:"currency,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                           `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedWithMeteredMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual     PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly    PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime    PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom     PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID                      param.Field[string]                                                       `json:"item_id,required"`
	MatrixWithDisplayNameConfig param.Field[map[string]interface{}]                                       `json:"matrix_with_display_name_config,required"`
	ModelType                   param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMatrixWithDisplayNamePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence string

const (
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceAnnual     PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "annual"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceMonthly    PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "monthly"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly  PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "quarterly"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceOneTime    PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "one_time"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceCustom     PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceAnnual, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceSemiAnnual, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceMonthly, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceQuarterly, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceOneTime, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelType string

const (
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                   `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingBulkWithProrationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingBulkWithProrationPriceCadence string

const (
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceAnnual     PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "annual"
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceSemiAnnual PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceMonthly    PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "monthly"
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceQuarterly  PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceOneTime    PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "one_time"
	PriceNewParamsNewFloatingBulkWithProrationPriceCadenceCustom     PriceNewParamsNewFloatingBulkWithProrationPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkWithProrationPriceCadenceAnnual, PriceNewParamsNewFloatingBulkWithProrationPriceCadenceSemiAnnual, PriceNewParamsNewFloatingBulkWithProrationPriceCadenceMonthly, PriceNewParamsNewFloatingBulkWithProrationPriceCadenceQuarterly, PriceNewParamsNewFloatingBulkWithProrationPriceCadenceOneTime, PriceNewParamsNewFloatingBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkWithProrationPriceModelType string

const (
	PriceNewParamsNewFloatingBulkWithProrationPriceModelTypeBulkWithProration PriceNewParamsNewFloatingBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PriceNewParamsNewFloatingBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                   param.Field[string]                 `json:"currency,required"`
	GroupedTieredPackageConfig param.Field[map[string]interface{}] `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                      `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedTieredPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence string

const (
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceAnnual     PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceMonthly    PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceOneTime    PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceCustom     PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceAnnual, PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceMonthly, PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceOneTime, PriceNewParamsNewFloatingGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedTieredPackagePriceModelType string

const (
	PriceNewParamsNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage PriceNewParamsNewFloatingGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence string

const (
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual     PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "annual"
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly    PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "monthly"
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly  PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime    PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "one_time"
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom     PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceAnnual, PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceMonthly, PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceQuarterly, PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceOneTime, PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelType string

const (
	PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[shared.NewBillingCycleConfigurationParam] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The configuration for the rate of the price currency to the invoicing currency.
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence string

const (
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual     PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "annual"
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly    PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "monthly"
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly  PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime    PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "one_time"
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom     PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceAnnual, PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceMonthly, PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceQuarterly, PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceOneTime, PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelType string

const (
	PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingCumulativeGroupedBulkPrice struct {
	// The cadence to bill for this price on.
	Cadence                     param.Field[PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	CumulativeGroupedBulkConfig param.Field[map[string]interface{}]                                     `json:"cumulative_grouped_bulk_config,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                       `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingCumulativeGroupedBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingCumulativeGroupedBulkPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence string

const (
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceAnnual     PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "annual"
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceMonthly    PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "monthly"
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly  PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "quarterly"
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceOneTime    PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "one_time"
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceCustom     PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceAnnual, PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceSemiAnnual, PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceMonthly, PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceQuarterly, PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceOneTime, PriceNewParamsNewFloatingCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelType string

const (
	PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                          param.Field[string]                 `json:"currency,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                             `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence string

const (
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual     PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly    PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime    PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom     PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual, PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly, PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime, PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelType string

const (
	PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMinimumCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMinimumCompositePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID        param.Field[string]                                                      `json:"item_id,required"`
	MinimumConfig param.Field[PriceNewParamsNewFloatingMinimumCompositePriceMinimumConfig] `json:"minimum_config,required"`
	ModelType     param.Field[PriceNewParamsNewFloatingMinimumCompositePriceModelType]     `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceNewParamsNewFloatingUnitPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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
}

func (r PriceNewParamsNewFloatingMinimumCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMinimumCompositePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMinimumCompositePriceCadence string

const (
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceAnnual     PriceNewParamsNewFloatingMinimumCompositePriceCadence = "annual"
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceSemiAnnual PriceNewParamsNewFloatingMinimumCompositePriceCadence = "semi_annual"
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceMonthly    PriceNewParamsNewFloatingMinimumCompositePriceCadence = "monthly"
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceQuarterly  PriceNewParamsNewFloatingMinimumCompositePriceCadence = "quarterly"
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceOneTime    PriceNewParamsNewFloatingMinimumCompositePriceCadence = "one_time"
	PriceNewParamsNewFloatingMinimumCompositePriceCadenceCustom     PriceNewParamsNewFloatingMinimumCompositePriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMinimumCompositePriceCadenceAnnual, PriceNewParamsNewFloatingMinimumCompositePriceCadenceSemiAnnual, PriceNewParamsNewFloatingMinimumCompositePriceCadenceMonthly, PriceNewParamsNewFloatingMinimumCompositePriceCadenceQuarterly, PriceNewParamsNewFloatingMinimumCompositePriceCadenceOneTime, PriceNewParamsNewFloatingMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMinimumCompositePriceMinimumConfig struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// By default, subtotals from minimum composite prices are prorated based on the
	// service period. Set to false to disable proration.
	Prorated param.Field[bool] `json:"prorated"`
}

func (r PriceNewParamsNewFloatingMinimumCompositePriceMinimumConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingMinimumCompositePriceModelType string

const (
	PriceNewParamsNewFloatingMinimumCompositePriceModelTypeMinimum PriceNewParamsNewFloatingMinimumCompositePriceModelType = "minimum"
)

func (r PriceNewParamsNewFloatingMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMinimumCompositePriceModelTypeMinimum:
		return true
	}
	return false
}

type PriceUpdateParams struct {
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PriceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceListParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [PriceListParams]'s query parameters as `url.Values`.
func (r PriceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PriceEvaluateParams struct {
	// The exclusive upper bound for event timestamps
	TimeframeEnd param.Field[time.Time] `json:"timeframe_end,required" format:"date-time"`
	// The inclusive lower bound for event timestamps
	TimeframeStart param.Field[time.Time] `json:"timeframe_start,required" format:"date-time"`
	// The ID of the customer to which this evaluation is scoped.
	CustomerID param.Field[string] `json:"customer_id"`
	// The external customer ID of the customer to which this evaluation is scoped.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the underlying billable metric
	Filter param.Field[string] `json:"filter"`
	// Properties (or
	// [computed properties](/extensibility/advanced-metrics#computed-properties)) used
	// to group the underlying billable metric
	GroupingKeys param.Field[[]string] `json:"grouping_keys"`
}

func (r PriceEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceEvaluateMultipleParams struct {
	// The exclusive upper bound for event timestamps
	TimeframeEnd param.Field[time.Time] `json:"timeframe_end,required" format:"date-time"`
	// The inclusive lower bound for event timestamps
	TimeframeStart param.Field[time.Time] `json:"timeframe_start,required" format:"date-time"`
	// The ID of the customer to which this evaluation is scoped.
	CustomerID param.Field[string] `json:"customer_id"`
	// The external customer ID of the customer to which this evaluation is scoped.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// List of prices to evaluate (max 100)
	PriceEvaluations param.Field[[]PriceEvaluateMultipleParamsPriceEvaluation] `json:"price_evaluations"`
}

func (r PriceEvaluateMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceEvaluateMultipleParamsPriceEvaluation struct {
	// The external ID of a price to evaluate that exists in your Orb account.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the underlying billable metric
	Filter param.Field[string] `json:"filter"`
	// Properties (or
	// [computed properties](/extensibility/advanced-metrics#computed-properties)) used
	// to group the underlying billable metric
	GroupingKeys param.Field[[]string] `json:"grouping_keys"`
	// An inline price definition to evaluate, allowing you to test price
	// configurations before adding them to Orb.
	Price param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceUnion] `json:"price"`
	// The ID of a price to evaluate that exists in your Orb account.
	PriceID param.Field[string] `json:"price_id"`
}

func (r PriceEvaluateMultipleParamsPriceEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An inline price definition to evaluate, allowing you to test price
// configurations before adding them to Orb.
type PriceEvaluateMultipleParamsPriceEvaluationsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                    `json:"item_id,required"`
	ModelType param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType] `json:"model_type,required"`
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
	InvoicingCycleConfiguration           param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[shared.MatrixConfigParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[shared.MatrixWithAllocationConfigParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}]                              `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}]                              `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}]                              `json:"metadata"`
	MinimumConfig                         param.Field[interface{}]                              `json:"minimum_config"`
	PackageConfig                         param.Field[shared.PackageConfigParam]                `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}]                              `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]                              `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]                              `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]                              `json:"threshold_total_amount_config"`
	TieredConfig                          param.Field[shared.TieredConfigParam]                 `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}]                              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}]                              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}]                              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}]                              `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigParam]                   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}]                              `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}]                              `json:"unit_with_proration_config"`
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPrice) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

// An inline price definition to evaluate, allowing you to test price
// configurations before adding them to Orb.
//
// Satisfied by [shared.NewFloatingUnitPriceParam],
// [shared.NewFloatingPackagePriceParam], [shared.NewFloatingMatrixPriceParam],
// [shared.NewFloatingMatrixWithAllocationPriceParam],
// [shared.NewFloatingTieredPriceParam], [shared.NewFloatingBulkPriceParam],
// [shared.NewFloatingThresholdTotalAmountPriceParam],
// [shared.NewFloatingTieredPackagePriceParam],
// [shared.NewFloatingGroupedTieredPriceParam],
// [shared.NewFloatingMaxGroupTieredPackagePriceParam],
// [shared.NewFloatingTieredWithMinimumPriceParam],
// [shared.NewFloatingPackageWithAllocationPriceParam],
// [shared.NewFloatingTieredPackageWithMinimumPriceParam],
// [shared.NewFloatingUnitWithPercentPriceParam],
// [shared.NewFloatingTieredWithProrationPriceParam],
// [shared.NewFloatingUnitWithProrationPriceParam],
// [shared.NewFloatingGroupedAllocationPriceParam],
// [shared.NewFloatingGroupedWithProratedMinimumPriceParam],
// [shared.NewFloatingGroupedWithMeteredMinimumPriceParam],
// [shared.NewFloatingMatrixWithDisplayNamePriceParam],
// [shared.NewFloatingBulkWithProrationPriceParam],
// [shared.NewFloatingGroupedTieredPackagePriceParam],
// [shared.NewFloatingScalableMatrixWithUnitPricingPriceParam],
// [shared.NewFloatingScalableMatrixWithTieredPricingPriceParam],
// [shared.NewFloatingCumulativeGroupedBulkPriceParam],
// [PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice],
// [shared.NewFloatingMinimumCompositePriceParam],
// [PriceEvaluateMultipleParamsPriceEvaluationsPrice].
type PriceEvaluateMultipleParamsPriceEvaluationsPriceUnion interface {
	ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion()
}

type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                          param.Field[string]                 `json:"currency,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                               `json:"item_id,required"`
	ModelType param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion] `json:"conversion_rate_config"`
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
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceUnion() {
}

// The cadence to bill for this price on.
type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence string

const (
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual     PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly    PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime    PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom     PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType string

const (
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	ConversionRateType param.Field[PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                          `json:"unit_config"`
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice].
type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion interface {
	ImplementsPriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion()
}

type PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType string

const (
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit   PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType = "unit"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType = "tiered"
)

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType) IsKnown() bool {
	switch r {
	case PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit, PriceEvaluateMultipleParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence string

const (
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceAnnual     PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "annual"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceSemiAnnual PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "semi_annual"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceMonthly    PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "monthly"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceQuarterly  PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "quarterly"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceOneTime    PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "one_time"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceCustom     PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence = "custom"
)

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceCadence) IsKnown() bool {
	switch r {
	case PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceAnnual, PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceSemiAnnual, PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceMonthly, PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceQuarterly, PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceOneTime, PriceEvaluateMultipleParamsPriceEvaluationsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType string

const (
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnit                            PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "unit"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypePackage                         PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "package"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrix                          PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "matrix"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrixWithAllocation            PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "matrix_with_allocation"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTiered                          PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "tiered"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeBulk                            PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "bulk"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeThresholdTotalAmount            PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "threshold_total_amount"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredPackage                   PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "tiered_package"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedTiered                   PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_tiered"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMaxGroupTieredPackage           PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "max_group_tiered_package"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredWithMinimum               PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "tiered_with_minimum"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypePackageWithAllocation           PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "package_with_allocation"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredPackageWithMinimum        PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "tiered_package_with_minimum"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnitWithPercent                 PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "unit_with_percent"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredWithProration             PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "tiered_with_proration"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnitWithProration               PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "unit_with_proration"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedAllocation               PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_allocation"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithProratedMinimum      PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_with_prorated_minimum"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithMeteredMinimum       PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_with_metered_minimum"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrixWithDisplayName           PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "matrix_with_display_name"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeBulkWithProration               PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "bulk_with_proration"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedTieredPackage            PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_tiered_package"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeScalableMatrixWithUnitPricing   PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "scalable_matrix_with_unit_pricing"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeScalableMatrixWithTieredPricing PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "scalable_matrix_with_tiered_pricing"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeCumulativeGroupedBulk           PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "cumulative_grouped_bulk"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithMinMaxThresholds     PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "grouped_with_min_max_thresholds"
	PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMinimum                         PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType = "minimum"
)

func (r PriceEvaluateMultipleParamsPriceEvaluationsPriceModelType) IsKnown() bool {
	switch r {
	case PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnit, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypePackage, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrix, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrixWithAllocation, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTiered, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeBulk, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeThresholdTotalAmount, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredPackage, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedTiered, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMaxGroupTieredPackage, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredWithMinimum, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypePackageWithAllocation, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredPackageWithMinimum, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnitWithPercent, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeTieredWithProration, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeUnitWithProration, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedAllocation, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithProratedMinimum, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithMeteredMinimum, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMatrixWithDisplayName, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeBulkWithProration, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedTieredPackage, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeScalableMatrixWithUnitPricing, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeScalableMatrixWithTieredPricing, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeCumulativeGroupedBulk, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeGroupedWithMinMaxThresholds, PriceEvaluateMultipleParamsPriceEvaluationsPriceModelTypeMinimum:
		return true
	}
	return false
}

type PriceEvaluatePreviewEventsParams struct {
	// The exclusive upper bound for event timestamps
	TimeframeEnd param.Field[time.Time] `json:"timeframe_end,required" format:"date-time"`
	// The inclusive lower bound for event timestamps
	TimeframeStart param.Field[time.Time] `json:"timeframe_start,required" format:"date-time"`
	// The ID of the customer to which this evaluation is scoped.
	CustomerID param.Field[string] `json:"customer_id"`
	// List of preview events to use instead of actual usage data
	Events param.Field[[]PriceEvaluatePreviewEventsParamsEvent] `json:"events"`
	// The external customer ID of the customer to which this evaluation is scoped.
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
	// List of prices to evaluate (max 100)
	PriceEvaluations param.Field[[]PriceEvaluatePreviewEventsParamsPriceEvaluation] `json:"price_evaluations"`
}

func (r PriceEvaluatePreviewEventsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceEvaluatePreviewEventsParamsEvent struct {
	// A name to meaningfully identify the action or event type.
	EventName param.Field[string] `json:"event_name,required"`
	// A dictionary of custom properties. Values in this dictionary must be numeric,
	// boolean, or strings. Nested dictionaries are disallowed.
	Properties param.Field[map[string]interface{}] `json:"properties,required"`
	// An ISO 8601 format date with no timezone offset (i.e. UTC). This should
	// represent the time that usage was recorded, and is particularly important to
	// attribute usage to a given billing period.
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// The Orb Customer identifier
	CustomerID param.Field[string] `json:"customer_id"`
	// An alias for the Orb customer, whose mapping is specified when creating the
	// customer
	ExternalCustomerID param.Field[string] `json:"external_customer_id"`
}

func (r PriceEvaluatePreviewEventsParamsEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceEvaluatePreviewEventsParamsPriceEvaluation struct {
	// The external ID of a price to evaluate that exists in your Orb account.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// A boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties) used to
	// filter the underlying billable metric
	Filter param.Field[string] `json:"filter"`
	// Properties (or
	// [computed properties](/extensibility/advanced-metrics#computed-properties)) used
	// to group the underlying billable metric
	GroupingKeys param.Field[[]string] `json:"grouping_keys"`
	// An inline price definition to evaluate, allowing you to test price
	// configurations before adding them to Orb.
	Price param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion] `json:"price"`
	// The ID of a price to evaluate that exists in your Orb account.
	PriceID param.Field[string] `json:"price_id"`
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An inline price definition to evaluate, allowing you to test price
// configurations before adding them to Orb.
type PriceEvaluatePreviewEventsParamsPriceEvaluationsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                         `json:"item_id,required"`
	ModelType param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType] `json:"model_type,required"`
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
	InvoicingCycleConfiguration           param.Field[shared.NewBillingCycleConfigurationParam] `json:"invoicing_cycle_configuration"`
	MatrixConfig                          param.Field[shared.MatrixConfigParam]                 `json:"matrix_config"`
	MatrixWithAllocationConfig            param.Field[shared.MatrixWithAllocationConfigParam]   `json:"matrix_with_allocation_config"`
	MatrixWithDisplayNameConfig           param.Field[interface{}]                              `json:"matrix_with_display_name_config"`
	MaxGroupTieredPackageConfig           param.Field[interface{}]                              `json:"max_group_tiered_package_config"`
	Metadata                              param.Field[interface{}]                              `json:"metadata"`
	MinimumConfig                         param.Field[interface{}]                              `json:"minimum_config"`
	PackageConfig                         param.Field[shared.PackageConfigParam]                `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}]                              `json:"package_with_allocation_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]                              `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]                              `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]                              `json:"threshold_total_amount_config"`
	TieredConfig                          param.Field[shared.TieredConfigParam]                 `json:"tiered_config"`
	TieredPackageConfig                   param.Field[interface{}]                              `json:"tiered_package_config"`
	TieredPackageWithMinimumConfig        param.Field[interface{}]                              `json:"tiered_package_with_minimum_config"`
	TieredWithMinimumConfig               param.Field[interface{}]                              `json:"tiered_with_minimum_config"`
	TieredWithProrationConfig             param.Field[interface{}]                              `json:"tiered_with_proration_config"`
	UnitConfig                            param.Field[shared.UnitConfigParam]                   `json:"unit_config"`
	UnitWithPercentConfig                 param.Field[interface{}]                              `json:"unit_with_percent_config"`
	UnitWithProrationConfig               param.Field[interface{}]                              `json:"unit_with_proration_config"`
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPrice) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

// An inline price definition to evaluate, allowing you to test price
// configurations before adding them to Orb.
//
// Satisfied by [shared.NewFloatingUnitPriceParam],
// [shared.NewFloatingPackagePriceParam], [shared.NewFloatingMatrixPriceParam],
// [shared.NewFloatingMatrixWithAllocationPriceParam],
// [shared.NewFloatingTieredPriceParam], [shared.NewFloatingBulkPriceParam],
// [shared.NewFloatingThresholdTotalAmountPriceParam],
// [shared.NewFloatingTieredPackagePriceParam],
// [shared.NewFloatingGroupedTieredPriceParam],
// [shared.NewFloatingMaxGroupTieredPackagePriceParam],
// [shared.NewFloatingTieredWithMinimumPriceParam],
// [shared.NewFloatingPackageWithAllocationPriceParam],
// [shared.NewFloatingTieredPackageWithMinimumPriceParam],
// [shared.NewFloatingUnitWithPercentPriceParam],
// [shared.NewFloatingTieredWithProrationPriceParam],
// [shared.NewFloatingUnitWithProrationPriceParam],
// [shared.NewFloatingGroupedAllocationPriceParam],
// [shared.NewFloatingGroupedWithProratedMinimumPriceParam],
// [shared.NewFloatingGroupedWithMeteredMinimumPriceParam],
// [shared.NewFloatingMatrixWithDisplayNamePriceParam],
// [shared.NewFloatingBulkWithProrationPriceParam],
// [shared.NewFloatingGroupedTieredPackagePriceParam],
// [shared.NewFloatingScalableMatrixWithUnitPricingPriceParam],
// [shared.NewFloatingScalableMatrixWithTieredPricingPriceParam],
// [shared.NewFloatingCumulativeGroupedBulkPriceParam],
// [PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice],
// [shared.NewFloatingMinimumCompositePriceParam],
// [PriceEvaluatePreviewEventsParamsPriceEvaluationsPrice].
type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion interface {
	ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion()
}

type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                          param.Field[string]                 `json:"currency,required"`
	GroupedWithMinMaxThresholdsConfig param.Field[map[string]interface{}] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID    param.Field[string]                                                                                                    `json:"item_id,required"`
	ModelType param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion] `json:"conversion_rate_config"`
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
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceUnion() {
}

// The cadence to bill for this price on.
type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence string

const (
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual     PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "annual"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly    PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly  PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime    PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom     PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType string

const (
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	ConversionRateType param.Field[PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                               `json:"unit_config"`
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPrice].
type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion interface {
	ImplementsPriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceUnion()
}

type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType string

const (
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit   PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType = "unit"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType = "tiered"
)

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateType) IsKnown() bool {
	switch r {
	case PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeUnit, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence string

const (
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceAnnual     PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "annual"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceSemiAnnual PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "semi_annual"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceMonthly    PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "monthly"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceQuarterly  PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "quarterly"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceOneTime    PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "one_time"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceCustom     PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence = "custom"
)

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadence) IsKnown() bool {
	switch r {
	case PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceAnnual, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceSemiAnnual, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceMonthly, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceQuarterly, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceOneTime, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType string

const (
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnit                            PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "unit"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypePackage                         PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "package"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrix                          PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "matrix"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrixWithAllocation            PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "matrix_with_allocation"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTiered                          PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "tiered"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeBulk                            PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "bulk"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeThresholdTotalAmount            PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "threshold_total_amount"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredPackage                   PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "tiered_package"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedTiered                   PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_tiered"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMaxGroupTieredPackage           PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "max_group_tiered_package"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredWithMinimum               PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "tiered_with_minimum"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypePackageWithAllocation           PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "package_with_allocation"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredPackageWithMinimum        PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "tiered_package_with_minimum"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnitWithPercent                 PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "unit_with_percent"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredWithProration             PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "tiered_with_proration"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnitWithProration               PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "unit_with_proration"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedAllocation               PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_allocation"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithProratedMinimum      PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_with_prorated_minimum"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithMeteredMinimum       PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_with_metered_minimum"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrixWithDisplayName           PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "matrix_with_display_name"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeBulkWithProration               PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "bulk_with_proration"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedTieredPackage            PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_tiered_package"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeScalableMatrixWithUnitPricing   PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "scalable_matrix_with_unit_pricing"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeScalableMatrixWithTieredPricing PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "scalable_matrix_with_tiered_pricing"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeCumulativeGroupedBulk           PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "cumulative_grouped_bulk"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithMinMaxThresholds     PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "grouped_with_min_max_thresholds"
	PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMinimum                         PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType = "minimum"
)

func (r PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelType) IsKnown() bool {
	switch r {
	case PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnit, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypePackage, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrix, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrixWithAllocation, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTiered, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeBulk, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeThresholdTotalAmount, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredPackage, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedTiered, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMaxGroupTieredPackage, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredWithMinimum, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypePackageWithAllocation, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredPackageWithMinimum, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnitWithPercent, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeTieredWithProration, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeUnitWithProration, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedAllocation, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithProratedMinimum, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithMeteredMinimum, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMatrixWithDisplayName, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeBulkWithProration, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedTieredPackage, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeScalableMatrixWithUnitPricing, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeScalableMatrixWithTieredPricing, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeCumulativeGroupedBulk, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeGroupedWithMinMaxThresholds, PriceEvaluatePreviewEventsParamsPriceEvaluationsPriceModelTypeMinimum:
		return true
	}
	return false
}
