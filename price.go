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

// This endpoint is used to create a [price](../reference/price). A price created
// using this endpoint is always an add-on, meaning that it’s not associated with a
// specific plan and can instead be individually added to subscriptions, including
// subscriptions on different plans.
//
// An `external_price_id` can be optionally specified as an alias to allow
// ergonomic interaction with prices in the Orb API.
//
// See the [Price resource](../reference/price) for the specification of different
// price model configurations possible in this endpoint.
func (r *PriceService) New(ctx context.Context, body PriceNewParams, opts ...option.RequestOption) (res *Price, err error) {
	opts = append(r.Options[:], opts...)
	path := "prices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to update the `metadata` property on a price. If you
// pass null for the metadata value, it will clear any existing metadata for that
// price.
func (r *PriceService) Update(ctx context.Context, priceID string, body PriceUpdateParams, opts ...option.RequestOption) (res *Price, err error) {
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
// [price creation endpoint](../reference/create-price).
func (r *PriceService) List(ctx context.Context, query PriceListParams, opts ...option.RequestOption) (res *pagination.Page[Price], err error) {
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
// [price creation endpoint](../reference/create-price).
func (r *PriceService) ListAutoPaging(ctx context.Context, query PriceListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Price] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint is used to evaluate the output of a price for a given customer and
// time range. It enables filtering and grouping the output using
// [computed properties](../guides/extensibility/advanced-metrics#computed-properties),
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

// This endpoint returns a price given an identifier.
func (r *PriceService) Fetch(ctx context.Context, priceID string, opts ...option.RequestOption) (res *Price, err error) {
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

// The Price resource represents a price that can be billed on a subscription,
// resulting in a charge on an invoice in the form of an invoice line item. Prices
// take a quantity and determine an amount to bill.
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given Price object. The model_type field
// determines the key for the configuration object that is present.
//
// ## Unit pricing
//
// With unit pricing, each unit costs a fixed amount.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "unit",
//	    "unit_config": {
//	        "unit_amount": "0.50"
//	    }
//	    ...
//	}
//
// ```
//
// ## Tiered pricing
//
// In tiered pricing, the cost of a given unit depends on the tier range that it
// falls into, where each tier range is defined by an upper and lower bound. For
// example, the first ten units may cost $0.50 each and all units thereafter may
// cost $0.10 each.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "tiered",
//	    "tiered_config": {
//	        "tiers": [
//	            {
//	                "first_unit": 1,
//	                "last_unit": 10,
//	                "unit_amount": "0.50"
//	            },
//	            {
//	                "first_unit": 11,
//	                "last_unit": null,
//	                "unit_amount": "0.10"
//	            }
//	        ]
//	    }
//	    ...
//
// ```
//
// ## Bulk pricing
//
// Bulk pricing applies when the number of units determine the cost of all units.
// For example, if you've bought less than 10 units, they may each be $0.50 for a
// total of $5.00. Once you've bought more than 10 units, all units may now be
// priced at $0.40 (i.e. 101 units total would be $40.40).
//
// ```json
//
//	{
//	    ...
//	    "model_type": "bulk",
//	    "bulk_config": {
//	        "tiers": [
//	            {
//	                "maximum_units": 10,
//	                "unit_amount": "0.50"
//	            },
//	            {
//	                "maximum_units": 1000,
//	                "unit_amount": "0.40"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Package pricing
//
// Package pricing defines the size or granularity of a unit for billing purposes.
// For example, if the package size is set to 5, then 4 units will be billed as 5
// and 6 units will be billed at 10.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "package",
//	    "package_config": {
//	        "package_amount": "0.80",
//	        "package_size": 10
//	    }
//	    ...
//	}
//
// ```
//
// ## BPS pricing
//
// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
// percent (the number of basis points to charge), as well as a cap per event to
// assess. For example, this would allow you to assess a fee of 0.25% on every
// payment you process, with a maximum charge of $25 per payment.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "bps",
//	    "bps_config": {
//	       "bps": 125,
//	       "per_unit_maximum": "11.00"
//	    }
//	    ...
//	 }
//
// ```
//
// ## Bulk BPS pricing
//
// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
// total quantity across all events. Similar to bulk pricing, the BPS parameters of
// a given event depends on the tier range that the billing period falls into. Each
// tier range is defined by an upper bound. For example, after $1.5M of payment
// volume is reached, each individual payment may have a lower cap or a smaller
// take-rate.
//
// ```json
//
//	    ...
//	    "model_type": "bulk_bps",
//	    "bulk_bps_config": {
//	        "tiers": [
//	           {
//	                "maximum_amount": "1000000.00",
//	                "bps": 125,
//	                "per_unit_maximum": "19.00"
//	           },
//	          {
//	                "maximum_amount": null,
//	                "bps": 115,
//	                "per_unit_maximum": "4.00"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Tiered BPS pricing
//
// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
// event's applicable parameter is a function of its marginal addition to the
// period total. Similar to tiered pricing, the BPS parameters of a given event
// depends on the tier range that it falls into, where each tier range is defined
// by an upper and lower bound. For example, the first few payments may have a 0.8
// BPS take-rate and all payments after a specific volume may incur a take-rate of
// 0.5 BPS each.
//
// ```json
//
//	    ...
//	    "model_type": "tiered_bps",
//	    "tiered_bps_config": {
//	        "tiers": [
//	           {
//	                "minimum_amount": "0",
//	                "maximum_amount": "1000000.00",
//	                "bps": 125,
//	                "per_unit_maximum": "19.00"
//	           },
//	          {
//	                "minimum_amount": "1000000.00",
//	                "maximum_amount": null,
//	                "bps": 115,
//	                "per_unit_maximum": "4.00"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Matrix pricing
//
// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
// `dimensions` defines the two event property values evaluated in this pricing
// model. In a one-dimensional matrix, the second value is `null`. Every
// configuration has a list of `matrix_values` which give the unit prices for
// specified property values. In a one-dimensional matrix, the matrix values will
// have `dimension_values` where the second value of the pair is null. If an event
// does not match any of the dimension values in the matrix, it will resort to the
// `default_unit_amount`.
//
// ```json
//
//	{
//	    "model_type": "matrix"
//	    "matrix_config": {
//	        "default_unit_amount": "3.00",
//	        "dimensions": [
//	            "cluster_name",
//	            "region"
//	        ],
//	        "matrix_values": [
//	            {
//	                "dimension_values": [
//	                    "alpha",
//	                    "west"
//	                ],
//	                "unit_amount": "2.00"
//	            },
//	            ...
//	        ]
//	    }
//	}
//
// ```
//
// ## Fixed fees
//
// Fixed fees are prices that are applied independent of usage quantities, and
// follow unit pricing. They also have an additional parameter
// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
// the quantity of units applied.
//
// ```json
//
//	{
//	    ...
//	    "id": "price_id",
//	    "model_type": "unit",
//	    "unit_config": {
//	       "unit_amount": "2.00"
//	    },
//	    "fixed_price_quantity": 3.0
//	    ...
//	}
//
// ```
type Price struct {
	ID string `json:"id,required"`
	// This field can have the runtime type of [PriceUnitPriceBillableMetric],
	// [PricePackagePriceBillableMetric], [PriceMatrixPriceBillableMetric],
	// [PriceTieredPriceBillableMetric], [PriceTieredBpsPriceBillableMetric],
	// [PriceBpsPriceBillableMetric], [PriceBulkBpsPriceBillableMetric],
	// [PriceBulkPriceBillableMetric], [PriceThresholdTotalAmountPriceBillableMetric],
	// [PriceTieredPackagePriceBillableMetric],
	// [PriceGroupedTieredPriceBillableMetric],
	// [PriceTieredWithMinimumPriceBillableMetric],
	// [PriceTieredPackageWithMinimumPriceBillableMetric],
	// [PricePackageWithAllocationPriceBillableMetric],
	// [PriceUnitWithPercentPriceBillableMetric],
	// [PriceMatrixWithAllocationPriceBillableMetric],
	// [PriceTieredWithProrationPriceBillableMetric],
	// [PriceUnitWithProrationPriceBillableMetric],
	// [PriceGroupedAllocationPriceBillableMetric],
	// [PriceGroupedWithProratedMinimumPriceBillableMetric],
	// [PriceGroupedWithMeteredMinimumPriceBillableMetric],
	// [PriceMatrixWithDisplayNamePriceBillableMetric],
	// [PriceBulkWithProrationPriceBillableMetric],
	// [PriceGroupedTieredPackagePriceBillableMetric].
	BillableMetric interface{} `json:"billable_metric,required"`
	// This field can have the runtime type of
	// [PriceUnitPriceBillingCycleConfiguration],
	// [PricePackagePriceBillingCycleConfiguration],
	// [PriceMatrixPriceBillingCycleConfiguration],
	// [PriceTieredPriceBillingCycleConfiguration],
	// [PriceTieredBpsPriceBillingCycleConfiguration],
	// [PriceBpsPriceBillingCycleConfiguration],
	// [PriceBulkBpsPriceBillingCycleConfiguration],
	// [PriceBulkPriceBillingCycleConfiguration],
	// [PriceThresholdTotalAmountPriceBillingCycleConfiguration],
	// [PriceTieredPackagePriceBillingCycleConfiguration],
	// [PriceGroupedTieredPriceBillingCycleConfiguration],
	// [PriceTieredWithMinimumPriceBillingCycleConfiguration],
	// [PriceTieredPackageWithMinimumPriceBillingCycleConfiguration],
	// [PricePackageWithAllocationPriceBillingCycleConfiguration],
	// [PriceUnitWithPercentPriceBillingCycleConfiguration],
	// [PriceMatrixWithAllocationPriceBillingCycleConfiguration],
	// [PriceTieredWithProrationPriceBillingCycleConfiguration],
	// [PriceUnitWithProrationPriceBillingCycleConfiguration],
	// [PriceGroupedAllocationPriceBillingCycleConfiguration],
	// [PriceGroupedWithProratedMinimumPriceBillingCycleConfiguration],
	// [PriceGroupedWithMeteredMinimumPriceBillingCycleConfiguration],
	// [PriceMatrixWithDisplayNamePriceBillingCycleConfiguration],
	// [PriceBulkWithProrationPriceBillingCycleConfiguration],
	// [PriceGroupedTieredPackagePriceBillingCycleConfiguration].
	BillingCycleConfiguration interface{}  `json:"billing_cycle_configuration,required"`
	Cadence                   PriceCadence `json:"cadence,required"`
	ConversionRate            float64      `json:"conversion_rate,required,nullable"`
	CreatedAt                 time.Time    `json:"created_at,required" format:"date-time"`
	// This field can have the runtime type of [PriceUnitPriceCreditAllocation],
	// [PricePackagePriceCreditAllocation], [PriceMatrixPriceCreditAllocation],
	// [PriceTieredPriceCreditAllocation], [PriceTieredBpsPriceCreditAllocation],
	// [PriceBpsPriceCreditAllocation], [PriceBulkBpsPriceCreditAllocation],
	// [PriceBulkPriceCreditAllocation],
	// [PriceThresholdTotalAmountPriceCreditAllocation],
	// [PriceTieredPackagePriceCreditAllocation],
	// [PriceGroupedTieredPriceCreditAllocation],
	// [PriceTieredWithMinimumPriceCreditAllocation],
	// [PriceTieredPackageWithMinimumPriceCreditAllocation],
	// [PricePackageWithAllocationPriceCreditAllocation],
	// [PriceUnitWithPercentPriceCreditAllocation],
	// [PriceMatrixWithAllocationPriceCreditAllocation],
	// [PriceTieredWithProrationPriceCreditAllocation],
	// [PriceUnitWithProrationPriceCreditAllocation],
	// [PriceGroupedAllocationPriceCreditAllocation],
	// [PriceGroupedWithProratedMinimumPriceCreditAllocation],
	// [PriceGroupedWithMeteredMinimumPriceCreditAllocation],
	// [PriceMatrixWithDisplayNamePriceCreditAllocation],
	// [PriceBulkWithProrationPriceCreditAllocation],
	// [PriceGroupedTieredPackagePriceCreditAllocation].
	CreditAllocation   interface{}     `json:"credit_allocation,required"`
	Currency           string          `json:"currency,required"`
	Discount           shared.Discount `json:"discount,required,nullable"`
	ExternalPriceID    string          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64         `json:"fixed_price_quantity,required,nullable"`
	// This field can have the runtime type of
	// [PriceUnitPriceInvoicingCycleConfiguration],
	// [PricePackagePriceInvoicingCycleConfiguration],
	// [PriceMatrixPriceInvoicingCycleConfiguration],
	// [PriceTieredPriceInvoicingCycleConfiguration],
	// [PriceTieredBpsPriceInvoicingCycleConfiguration],
	// [PriceBpsPriceInvoicingCycleConfiguration],
	// [PriceBulkBpsPriceInvoicingCycleConfiguration],
	// [PriceBulkPriceInvoicingCycleConfiguration],
	// [PriceThresholdTotalAmountPriceInvoicingCycleConfiguration],
	// [PriceTieredPackagePriceInvoicingCycleConfiguration],
	// [PriceGroupedTieredPriceInvoicingCycleConfiguration],
	// [PriceTieredWithMinimumPriceInvoicingCycleConfiguration],
	// [PriceTieredPackageWithMinimumPriceInvoicingCycleConfiguration],
	// [PricePackageWithAllocationPriceInvoicingCycleConfiguration],
	// [PriceUnitWithPercentPriceInvoicingCycleConfiguration],
	// [PriceMatrixWithAllocationPriceInvoicingCycleConfiguration],
	// [PriceTieredWithProrationPriceInvoicingCycleConfiguration],
	// [PriceUnitWithProrationPriceInvoicingCycleConfiguration],
	// [PriceGroupedAllocationPriceInvoicingCycleConfiguration],
	// [PriceGroupedWithProratedMinimumPriceInvoicingCycleConfiguration],
	// [PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration],
	// [PriceMatrixWithDisplayNamePriceInvoicingCycleConfiguration],
	// [PriceBulkWithProrationPriceInvoicingCycleConfiguration],
	// [PriceGroupedTieredPackagePriceInvoicingCycleConfiguration].
	InvoicingCycleConfiguration interface{} `json:"invoicing_cycle_configuration,required"`
	// This field can have the runtime type of [PriceUnitPriceItem],
	// [PricePackagePriceItem], [PriceMatrixPriceItem], [PriceTieredPriceItem],
	// [PriceTieredBpsPriceItem], [PriceBpsPriceItem], [PriceBulkBpsPriceItem],
	// [PriceBulkPriceItem], [PriceThresholdTotalAmountPriceItem],
	// [PriceTieredPackagePriceItem], [PriceGroupedTieredPriceItem],
	// [PriceTieredWithMinimumPriceItem], [PriceTieredPackageWithMinimumPriceItem],
	// [PricePackageWithAllocationPriceItem], [PriceUnitWithPercentPriceItem],
	// [PriceMatrixWithAllocationPriceItem], [PriceTieredWithProrationPriceItem],
	// [PriceUnitWithProrationPriceItem], [PriceGroupedAllocationPriceItem],
	// [PriceGroupedWithProratedMinimumPriceItem],
	// [PriceGroupedWithMeteredMinimumPriceItem],
	// [PriceMatrixWithDisplayNamePriceItem], [PriceBulkWithProrationPriceItem],
	// [PriceGroupedTieredPackagePriceItem].
	Item interface{} `json:"item,required"`
	// This field can have the runtime type of [PriceUnitPriceMaximum],
	// [PricePackagePriceMaximum], [PriceMatrixPriceMaximum],
	// [PriceTieredPriceMaximum], [PriceTieredBpsPriceMaximum], [PriceBpsPriceMaximum],
	// [PriceBulkBpsPriceMaximum], [PriceBulkPriceMaximum],
	// [PriceThresholdTotalAmountPriceMaximum], [PriceTieredPackagePriceMaximum],
	// [PriceGroupedTieredPriceMaximum], [PriceTieredWithMinimumPriceMaximum],
	// [PriceTieredPackageWithMinimumPriceMaximum],
	// [PricePackageWithAllocationPriceMaximum], [PriceUnitWithPercentPriceMaximum],
	// [PriceMatrixWithAllocationPriceMaximum], [PriceTieredWithProrationPriceMaximum],
	// [PriceUnitWithProrationPriceMaximum], [PriceGroupedAllocationPriceMaximum],
	// [PriceGroupedWithProratedMinimumPriceMaximum],
	// [PriceGroupedWithMeteredMinimumPriceMaximum],
	// [PriceMatrixWithDisplayNamePriceMaximum], [PriceBulkWithProrationPriceMaximum],
	// [PriceGroupedTieredPackagePriceMaximum].
	Maximum       interface{} `json:"maximum,required"`
	MaximumAmount string      `json:"maximum_amount,required,nullable"`
	// This field can have the runtime type of [map[string]string].
	Metadata interface{} `json:"metadata,required"`
	// This field can have the runtime type of [PriceUnitPriceMinimum],
	// [PricePackagePriceMinimum], [PriceMatrixPriceMinimum],
	// [PriceTieredPriceMinimum], [PriceTieredBpsPriceMinimum], [PriceBpsPriceMinimum],
	// [PriceBulkBpsPriceMinimum], [PriceBulkPriceMinimum],
	// [PriceThresholdTotalAmountPriceMinimum], [PriceTieredPackagePriceMinimum],
	// [PriceGroupedTieredPriceMinimum], [PriceTieredWithMinimumPriceMinimum],
	// [PriceTieredPackageWithMinimumPriceMinimum],
	// [PricePackageWithAllocationPriceMinimum], [PriceUnitWithPercentPriceMinimum],
	// [PriceMatrixWithAllocationPriceMinimum], [PriceTieredWithProrationPriceMinimum],
	// [PriceUnitWithProrationPriceMinimum], [PriceGroupedAllocationPriceMinimum],
	// [PriceGroupedWithProratedMinimumPriceMinimum],
	// [PriceGroupedWithMeteredMinimumPriceMinimum],
	// [PriceMatrixWithDisplayNamePriceMinimum], [PriceBulkWithProrationPriceMinimum],
	// [PriceGroupedTieredPackagePriceMinimum].
	Minimum        interface{}    `json:"minimum,required"`
	MinimumAmount  string         `json:"minimum_amount,required,nullable"`
	ModelType      PriceModelType `json:"model_type,required"`
	Name           string         `json:"name,required"`
	PlanPhaseOrder int64          `json:"plan_phase_order,required,nullable"`
	PriceType      PricePriceType `json:"price_type,required"`
	// This field can have the runtime type of [PriceBpsPriceBpsConfig].
	BpsConfig interface{} `json:"bps_config"`
	// This field can have the runtime type of [PriceBulkBpsPriceBulkBpsConfig].
	BulkBpsConfig interface{} `json:"bulk_bps_config"`
	// This field can have the runtime type of [PriceBulkPriceBulkConfig].
	BulkConfig interface{} `json:"bulk_config"`
	// This field can have the runtime type of [map[string]interface{}].
	BulkWithProrationConfig interface{} `json:"bulk_with_proration_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedAllocationConfig interface{} `json:"grouped_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedTieredConfig interface{} `json:"grouped_tiered_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedTieredPackageConfig interface{} `json:"grouped_tiered_package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedWithMeteredMinimumConfig interface{} `json:"grouped_with_metered_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	GroupedWithProratedMinimumConfig interface{} `json:"grouped_with_prorated_minimum_config"`
	// This field can have the runtime type of [PriceMatrixPriceMatrixConfig].
	MatrixConfig interface{} `json:"matrix_config"`
	// This field can have the runtime type of
	// [PriceMatrixWithAllocationPriceMatrixWithAllocationConfig].
	MatrixWithAllocationConfig interface{} `json:"matrix_with_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	MatrixWithDisplayNameConfig interface{} `json:"matrix_with_display_name_config"`
	// This field can have the runtime type of [PricePackagePricePackageConfig].
	PackageConfig interface{} `json:"package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	PackageWithAllocationConfig interface{} `json:"package_with_allocation_config"`
	// This field can have the runtime type of [map[string]interface{}].
	ThresholdTotalAmountConfig interface{} `json:"threshold_total_amount_config"`
	// This field can have the runtime type of [PriceTieredBpsPriceTieredBpsConfig].
	TieredBpsConfig interface{} `json:"tiered_bps_config"`
	// This field can have the runtime type of [PriceTieredPriceTieredConfig].
	TieredConfig interface{} `json:"tiered_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredPackageConfig interface{} `json:"tiered_package_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredPackageWithMinimumConfig interface{} `json:"tiered_package_with_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredWithMinimumConfig interface{} `json:"tiered_with_minimum_config"`
	// This field can have the runtime type of [map[string]interface{}].
	TieredWithProrationConfig interface{} `json:"tiered_with_proration_config"`
	// This field can have the runtime type of [PriceUnitPriceUnitConfig].
	UnitConfig interface{} `json:"unit_config"`
	// This field can have the runtime type of [map[string]interface{}].
	UnitWithPercentConfig interface{} `json:"unit_with_percent_config"`
	// This field can have the runtime type of [map[string]interface{}].
	UnitWithProrationConfig interface{} `json:"unit_with_proration_config"`
	JSON                    priceJSON   `json:"-"`
	union                   PriceUnion
}

// priceJSON contains the JSON metadata for the struct [Price]
type priceJSON struct {
	ID                               apijson.Field
	BillableMetric                   apijson.Field
	BillingCycleConfiguration        apijson.Field
	Cadence                          apijson.Field
	ConversionRate                   apijson.Field
	CreatedAt                        apijson.Field
	CreditAllocation                 apijson.Field
	Currency                         apijson.Field
	Discount                         apijson.Field
	ExternalPriceID                  apijson.Field
	FixedPriceQuantity               apijson.Field
	InvoicingCycleConfiguration      apijson.Field
	Item                             apijson.Field
	Maximum                          apijson.Field
	MaximumAmount                    apijson.Field
	Metadata                         apijson.Field
	Minimum                          apijson.Field
	MinimumAmount                    apijson.Field
	ModelType                        apijson.Field
	Name                             apijson.Field
	PlanPhaseOrder                   apijson.Field
	PriceType                        apijson.Field
	BpsConfig                        apijson.Field
	BulkBpsConfig                    apijson.Field
	BulkConfig                       apijson.Field
	BulkWithProrationConfig          apijson.Field
	GroupedAllocationConfig          apijson.Field
	GroupedTieredConfig              apijson.Field
	GroupedTieredPackageConfig       apijson.Field
	GroupedWithMeteredMinimumConfig  apijson.Field
	GroupedWithProratedMinimumConfig apijson.Field
	MatrixConfig                     apijson.Field
	MatrixWithAllocationConfig       apijson.Field
	MatrixWithDisplayNameConfig      apijson.Field
	PackageConfig                    apijson.Field
	PackageWithAllocationConfig      apijson.Field
	ThresholdTotalAmountConfig       apijson.Field
	TieredBpsConfig                  apijson.Field
	TieredConfig                     apijson.Field
	TieredPackageConfig              apijson.Field
	TieredPackageWithMinimumConfig   apijson.Field
	TieredWithMinimumConfig          apijson.Field
	TieredWithProrationConfig        apijson.Field
	UnitConfig                       apijson.Field
	UnitWithPercentConfig            apijson.Field
	UnitWithProrationConfig          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r priceJSON) RawJSON() string {
	return r.raw
}

func (r *Price) UnmarshalJSON(data []byte) (err error) {
	*r = Price{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PriceUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PriceUnitPrice], [PricePackagePrice],
// [PriceMatrixPrice], [PriceTieredPrice], [PriceTieredBpsPrice], [PriceBpsPrice],
// [PriceBulkBpsPrice], [PriceBulkPrice], [PriceThresholdTotalAmountPrice],
// [PriceTieredPackagePrice], [PriceGroupedTieredPrice],
// [PriceTieredWithMinimumPrice], [PriceTieredPackageWithMinimumPrice],
// [PricePackageWithAllocationPrice], [PriceUnitWithPercentPrice],
// [PriceMatrixWithAllocationPrice], [PriceTieredWithProrationPrice],
// [PriceUnitWithProrationPrice], [PriceGroupedAllocationPrice],
// [PriceGroupedWithProratedMinimumPrice], [PriceGroupedWithMeteredMinimumPrice],
// [PriceMatrixWithDisplayNamePrice], [PriceBulkWithProrationPrice],
// [PriceGroupedTieredPackagePrice].
func (r Price) AsUnion() PriceUnion {
	return r.union
}

// The Price resource represents a price that can be billed on a subscription,
// resulting in a charge on an invoice in the form of an invoice line item. Prices
// take a quantity and determine an amount to bill.
//
// Orb supports a few different pricing models out of the box. Each of these models
// is serialized differently in a given Price object. The model_type field
// determines the key for the configuration object that is present.
//
// ## Unit pricing
//
// With unit pricing, each unit costs a fixed amount.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "unit",
//	    "unit_config": {
//	        "unit_amount": "0.50"
//	    }
//	    ...
//	}
//
// ```
//
// ## Tiered pricing
//
// In tiered pricing, the cost of a given unit depends on the tier range that it
// falls into, where each tier range is defined by an upper and lower bound. For
// example, the first ten units may cost $0.50 each and all units thereafter may
// cost $0.10 each.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "tiered",
//	    "tiered_config": {
//	        "tiers": [
//	            {
//	                "first_unit": 1,
//	                "last_unit": 10,
//	                "unit_amount": "0.50"
//	            },
//	            {
//	                "first_unit": 11,
//	                "last_unit": null,
//	                "unit_amount": "0.10"
//	            }
//	        ]
//	    }
//	    ...
//
// ```
//
// ## Bulk pricing
//
// Bulk pricing applies when the number of units determine the cost of all units.
// For example, if you've bought less than 10 units, they may each be $0.50 for a
// total of $5.00. Once you've bought more than 10 units, all units may now be
// priced at $0.40 (i.e. 101 units total would be $40.40).
//
// ```json
//
//	{
//	    ...
//	    "model_type": "bulk",
//	    "bulk_config": {
//	        "tiers": [
//	            {
//	                "maximum_units": 10,
//	                "unit_amount": "0.50"
//	            },
//	            {
//	                "maximum_units": 1000,
//	                "unit_amount": "0.40"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Package pricing
//
// Package pricing defines the size or granularity of a unit for billing purposes.
// For example, if the package size is set to 5, then 4 units will be billed as 5
// and 6 units will be billed at 10.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "package",
//	    "package_config": {
//	        "package_amount": "0.80",
//	        "package_size": 10
//	    }
//	    ...
//	}
//
// ```
//
// ## BPS pricing
//
// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
// percent (the number of basis points to charge), as well as a cap per event to
// assess. For example, this would allow you to assess a fee of 0.25% on every
// payment you process, with a maximum charge of $25 per payment.
//
// ```json
//
//	{
//	    ...
//	    "model_type": "bps",
//	    "bps_config": {
//	       "bps": 125,
//	       "per_unit_maximum": "11.00"
//	    }
//	    ...
//	 }
//
// ```
//
// ## Bulk BPS pricing
//
// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
// total quantity across all events. Similar to bulk pricing, the BPS parameters of
// a given event depends on the tier range that the billing period falls into. Each
// tier range is defined by an upper bound. For example, after $1.5M of payment
// volume is reached, each individual payment may have a lower cap or a smaller
// take-rate.
//
// ```json
//
//	    ...
//	    "model_type": "bulk_bps",
//	    "bulk_bps_config": {
//	        "tiers": [
//	           {
//	                "maximum_amount": "1000000.00",
//	                "bps": 125,
//	                "per_unit_maximum": "19.00"
//	           },
//	          {
//	                "maximum_amount": null,
//	                "bps": 115,
//	                "per_unit_maximum": "4.00"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Tiered BPS pricing
//
// Tiered BPS pricing specifies BPS parameters in a graduated manner, where an
// event's applicable parameter is a function of its marginal addition to the
// period total. Similar to tiered pricing, the BPS parameters of a given event
// depends on the tier range that it falls into, where each tier range is defined
// by an upper and lower bound. For example, the first few payments may have a 0.8
// BPS take-rate and all payments after a specific volume may incur a take-rate of
// 0.5 BPS each.
//
// ```json
//
//	    ...
//	    "model_type": "tiered_bps",
//	    "tiered_bps_config": {
//	        "tiers": [
//	           {
//	                "minimum_amount": "0",
//	                "maximum_amount": "1000000.00",
//	                "bps": 125,
//	                "per_unit_maximum": "19.00"
//	           },
//	          {
//	                "minimum_amount": "1000000.00",
//	                "maximum_amount": null,
//	                "bps": 115,
//	                "per_unit_maximum": "4.00"
//	            }
//	        ]
//	    }
//	    ...
//	}
//
// ```
//
// ## Matrix pricing
//
// Matrix pricing defines a set of unit prices in a one or two-dimensional matrix.
// `dimensions` defines the two event property values evaluated in this pricing
// model. In a one-dimensional matrix, the second value is `null`. Every
// configuration has a list of `matrix_values` which give the unit prices for
// specified property values. In a one-dimensional matrix, the matrix values will
// have `dimension_values` where the second value of the pair is null. If an event
// does not match any of the dimension values in the matrix, it will resort to the
// `default_unit_amount`.
//
// ```json
//
//	{
//	    "model_type": "matrix"
//	    "matrix_config": {
//	        "default_unit_amount": "3.00",
//	        "dimensions": [
//	            "cluster_name",
//	            "region"
//	        ],
//	        "matrix_values": [
//	            {
//	                "dimension_values": [
//	                    "alpha",
//	                    "west"
//	                ],
//	                "unit_amount": "2.00"
//	            },
//	            ...
//	        ]
//	    }
//	}
//
// ```
//
// ## Fixed fees
//
// Fixed fees are prices that are applied independent of usage quantities, and
// follow unit pricing. They also have an additional parameter
// `fixed_price_quantity`. If the Price represents a fixed cost, this represents
// the quantity of units applied.
//
// ```json
//
//	{
//	    ...
//	    "id": "price_id",
//	    "model_type": "unit",
//	    "unit_config": {
//	       "unit_amount": "2.00"
//	    },
//	    "fixed_price_quantity": 3.0
//	    ...
//	}
//
// ```
//
// Union satisfied by [PriceUnitPrice], [PricePackagePrice], [PriceMatrixPrice],
// [PriceTieredPrice], [PriceTieredBpsPrice], [PriceBpsPrice], [PriceBulkBpsPrice],
// [PriceBulkPrice], [PriceThresholdTotalAmountPrice], [PriceTieredPackagePrice],
// [PriceGroupedTieredPrice], [PriceTieredWithMinimumPrice],
// [PriceTieredPackageWithMinimumPrice], [PricePackageWithAllocationPrice],
// [PriceUnitWithPercentPrice], [PriceMatrixWithAllocationPrice],
// [PriceTieredWithProrationPrice], [PriceUnitWithProrationPrice],
// [PriceGroupedAllocationPrice], [PriceGroupedWithProratedMinimumPrice],
// [PriceGroupedWithMeteredMinimumPrice], [PriceMatrixWithDisplayNamePrice],
// [PriceBulkWithProrationPrice] or [PriceGroupedTieredPackagePrice].
type PriceUnion interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PriceUnion)(nil)).Elem(),
		"model_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitPrice{}),
			DiscriminatorValue: "unit",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PricePackagePrice{}),
			DiscriminatorValue: "package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixPrice{}),
			DiscriminatorValue: "matrix",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPrice{}),
			DiscriminatorValue: "tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredBpsPrice{}),
			DiscriminatorValue: "tiered_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBpsPrice{}),
			DiscriminatorValue: "bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkBpsPrice{}),
			DiscriminatorValue: "bulk_bps",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkPrice{}),
			DiscriminatorValue: "bulk",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceThresholdTotalAmountPrice{}),
			DiscriminatorValue: "threshold_total_amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPackagePrice{}),
			DiscriminatorValue: "tiered_package",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedTieredPrice{}),
			DiscriminatorValue: "grouped_tiered",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredWithMinimumPrice{}),
			DiscriminatorValue: "tiered_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredPackageWithMinimumPrice{}),
			DiscriminatorValue: "tiered_package_with_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PricePackageWithAllocationPrice{}),
			DiscriminatorValue: "package_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitWithPercentPrice{}),
			DiscriminatorValue: "unit_with_percent",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixWithAllocationPrice{}),
			DiscriminatorValue: "matrix_with_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceTieredWithProrationPrice{}),
			DiscriminatorValue: "tiered_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceUnitWithProrationPrice{}),
			DiscriminatorValue: "unit_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedAllocationPrice{}),
			DiscriminatorValue: "grouped_allocation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedWithProratedMinimumPrice{}),
			DiscriminatorValue: "grouped_with_prorated_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedWithMeteredMinimumPrice{}),
			DiscriminatorValue: "grouped_with_metered_minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceMatrixWithDisplayNamePrice{}),
			DiscriminatorValue: "matrix_with_display_name",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceBulkWithProrationPrice{}),
			DiscriminatorValue: "bulk_with_proration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PriceGroupedTieredPackagePrice{}),
			DiscriminatorValue: "grouped_tiered_package",
		},
	)
}

type PriceUnitPrice struct {
	ID                          string                                    `json:"id,required"`
	BillableMetric              PriceUnitPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceUnitPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceUnitPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceUnitPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                    `json:"currency,required"`
	Discount                    shared.Discount                           `json:"discount,required,nullable"`
	ExternalPriceID             string                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceUnitPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceUnitPriceItem                        `json:"item,required"`
	Maximum                     PriceUnitPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string        `json:"metadata,required"`
	Minimum        PriceUnitPriceMinimum    `json:"minimum,required,nullable"`
	MinimumAmount  string                   `json:"minimum_amount,required,nullable"`
	ModelType      PriceUnitPriceModelType  `json:"model_type,required"`
	Name           string                   `json:"name,required"`
	PlanPhaseOrder int64                    `json:"plan_phase_order,required,nullable"`
	PriceType      PriceUnitPricePriceType  `json:"price_type,required"`
	UnitConfig     PriceUnitPriceUnitConfig `json:"unit_config,required"`
	JSON           priceUnitPriceJSON       `json:"-"`
}

// priceUnitPriceJSON contains the JSON metadata for the struct [PriceUnitPrice]
type priceUnitPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	UnitConfig                  apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceUnitPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitPrice) implementsPrice() {}

type PriceUnitPriceBillableMetric struct {
	ID   string                           `json:"id,required"`
	JSON priceUnitPriceBillableMetricJSON `json:"-"`
}

// priceUnitPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceUnitPriceBillableMetric]
type priceUnitPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceBillingCycleConfiguration struct {
	Duration     int64                                               `json:"duration,required"`
	DurationUnit PriceUnitPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceUnitPriceBillingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceUnitPriceBillingCycleConfiguration]
type priceUnitPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PriceUnitPriceBillingCycleConfigurationDurationUnitDay   PriceUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PriceUnitPriceBillingCycleConfigurationDurationUnitMonth PriceUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitPriceBillingCycleConfigurationDurationUnitDay, PriceUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitPriceCadence string

const (
	PriceUnitPriceCadenceOneTime    PriceUnitPriceCadence = "one_time"
	PriceUnitPriceCadenceMonthly    PriceUnitPriceCadence = "monthly"
	PriceUnitPriceCadenceQuarterly  PriceUnitPriceCadence = "quarterly"
	PriceUnitPriceCadenceSemiAnnual PriceUnitPriceCadence = "semi_annual"
	PriceUnitPriceCadenceAnnual     PriceUnitPriceCadence = "annual"
	PriceUnitPriceCadenceCustom     PriceUnitPriceCadence = "custom"
)

func (r PriceUnitPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitPriceCadenceOneTime, PriceUnitPriceCadenceMonthly, PriceUnitPriceCadenceQuarterly, PriceUnitPriceCadenceSemiAnnual, PriceUnitPriceCadenceAnnual, PriceUnitPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitPriceCreditAllocation struct {
	AllowsRollover bool                               `json:"allows_rollover,required"`
	Currency       string                             `json:"currency,required"`
	JSON           priceUnitPriceCreditAllocationJSON `json:"-"`
}

// priceUnitPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceUnitPriceCreditAllocation]
type priceUnitPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceUnitPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                 `json:"duration,required"`
	DurationUnit PriceUnitPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceUnitPriceInvoicingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceUnitPriceInvoicingCycleConfiguration]
type priceUnitPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceUnitPriceInvoicingCycleConfigurationDurationUnitDay   PriceUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceUnitPriceInvoicingCycleConfigurationDurationUnitMonth PriceUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitPriceInvoicingCycleConfigurationDurationUnitDay, PriceUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitPriceItem struct {
	ID   string                 `json:"id,required"`
	Name string                 `json:"name,required"`
	JSON priceUnitPriceItemJSON `json:"-"`
}

// priceUnitPriceItemJSON contains the JSON metadata for the struct
// [PriceUnitPriceItem]
type priceUnitPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                    `json:"maximum_amount,required"`
	JSON          priceUnitPriceMaximumJSON `json:"-"`
}

// priceUnitPriceMaximumJSON contains the JSON metadata for the struct
// [PriceUnitPriceMaximum]
type priceUnitPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                    `json:"minimum_amount,required"`
	JSON          priceUnitPriceMinimumJSON `json:"-"`
}

// priceUnitPriceMinimumJSON contains the JSON metadata for the struct
// [PriceUnitPriceMinimum]
type priceUnitPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitPriceModelType string

const (
	PriceUnitPriceModelTypeUnit PriceUnitPriceModelType = "unit"
)

func (r PriceUnitPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type PriceUnitPricePriceType string

const (
	PriceUnitPricePriceTypeUsagePrice PriceUnitPricePriceType = "usage_price"
	PriceUnitPricePriceTypeFixedPrice PriceUnitPricePriceType = "fixed_price"
)

func (r PriceUnitPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitPricePriceTypeUsagePrice, PriceUnitPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount string                       `json:"unit_amount,required"`
	JSON       priceUnitPriceUnitConfigJSON `json:"-"`
}

// priceUnitPriceUnitConfigJSON contains the JSON metadata for the struct
// [PriceUnitPriceUnitConfig]
type priceUnitPriceUnitConfigJSON struct {
	UnitAmount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitPriceUnitConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitPriceUnitConfigJSON) RawJSON() string {
	return r.raw
}

type PricePackagePrice struct {
	ID                          string                                       `json:"id,required"`
	BillableMetric              PricePackagePriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PricePackagePriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PricePackagePriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                      `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                    `json:"created_at,required" format:"date-time"`
	CreditAllocation            PricePackagePriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                       `json:"currency,required"`
	Discount                    shared.Discount                              `json:"discount,required,nullable"`
	ExternalPriceID             string                                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                      `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PricePackagePriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PricePackagePriceItem                        `json:"item,required"`
	Maximum                     PricePackagePriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                       `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string              `json:"metadata,required"`
	Minimum        PricePackagePriceMinimum       `json:"minimum,required,nullable"`
	MinimumAmount  string                         `json:"minimum_amount,required,nullable"`
	ModelType      PricePackagePriceModelType     `json:"model_type,required"`
	Name           string                         `json:"name,required"`
	PackageConfig  PricePackagePricePackageConfig `json:"package_config,required"`
	PlanPhaseOrder int64                          `json:"plan_phase_order,required,nullable"`
	PriceType      PricePackagePricePriceType     `json:"price_type,required"`
	JSON           pricePackagePriceJSON          `json:"-"`
}

// pricePackagePriceJSON contains the JSON metadata for the struct
// [PricePackagePrice]
type pricePackagePriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PackageConfig               apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PricePackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PricePackagePrice) implementsPrice() {}

type PricePackagePriceBillableMetric struct {
	ID   string                              `json:"id,required"`
	JSON pricePackagePriceBillableMetricJSON `json:"-"`
}

// pricePackagePriceBillableMetricJSON contains the JSON metadata for the struct
// [PricePackagePriceBillableMetric]
type pricePackagePriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricePackagePriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceBillingCycleConfiguration struct {
	Duration     int64                                                  `json:"duration,required"`
	DurationUnit PricePackagePriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         pricePackagePriceBillingCycleConfigurationJSON         `json:"-"`
}

// pricePackagePriceBillingCycleConfigurationJSON contains the JSON metadata for
// the struct [PricePackagePriceBillingCycleConfiguration]
type pricePackagePriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricePackagePriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceBillingCycleConfigurationDurationUnit string

const (
	PricePackagePriceBillingCycleConfigurationDurationUnitDay   PricePackagePriceBillingCycleConfigurationDurationUnit = "day"
	PricePackagePriceBillingCycleConfigurationDurationUnitMonth PricePackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PricePackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PricePackagePriceBillingCycleConfigurationDurationUnitDay, PricePackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PricePackagePriceCadence string

const (
	PricePackagePriceCadenceOneTime    PricePackagePriceCadence = "one_time"
	PricePackagePriceCadenceMonthly    PricePackagePriceCadence = "monthly"
	PricePackagePriceCadenceQuarterly  PricePackagePriceCadence = "quarterly"
	PricePackagePriceCadenceSemiAnnual PricePackagePriceCadence = "semi_annual"
	PricePackagePriceCadenceAnnual     PricePackagePriceCadence = "annual"
	PricePackagePriceCadenceCustom     PricePackagePriceCadence = "custom"
)

func (r PricePackagePriceCadence) IsKnown() bool {
	switch r {
	case PricePackagePriceCadenceOneTime, PricePackagePriceCadenceMonthly, PricePackagePriceCadenceQuarterly, PricePackagePriceCadenceSemiAnnual, PricePackagePriceCadenceAnnual, PricePackagePriceCadenceCustom:
		return true
	}
	return false
}

type PricePackagePriceCreditAllocation struct {
	AllowsRollover bool                                  `json:"allows_rollover,required"`
	Currency       string                                `json:"currency,required"`
	JSON           pricePackagePriceCreditAllocationJSON `json:"-"`
}

// pricePackagePriceCreditAllocationJSON contains the JSON metadata for the struct
// [PricePackagePriceCreditAllocation]
type pricePackagePriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PricePackagePriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceInvoicingCycleConfiguration struct {
	Duration     int64                                                    `json:"duration,required"`
	DurationUnit PricePackagePriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         pricePackagePriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// pricePackagePriceInvoicingCycleConfigurationJSON contains the JSON metadata for
// the struct [PricePackagePriceInvoicingCycleConfiguration]
type pricePackagePriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricePackagePriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PricePackagePriceInvoicingCycleConfigurationDurationUnitDay   PricePackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PricePackagePriceInvoicingCycleConfigurationDurationUnitMonth PricePackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PricePackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PricePackagePriceInvoicingCycleConfigurationDurationUnitDay, PricePackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PricePackagePriceItem struct {
	ID   string                    `json:"id,required"`
	Name string                    `json:"name,required"`
	JSON pricePackagePriceItemJSON `json:"-"`
}

// pricePackagePriceItemJSON contains the JSON metadata for the struct
// [PricePackagePriceItem]
type pricePackagePriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricePackagePriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceItemJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                       `json:"maximum_amount,required"`
	JSON          pricePackagePriceMaximumJSON `json:"-"`
}

// pricePackagePriceMaximumJSON contains the JSON metadata for the struct
// [PricePackagePriceMaximum]
type pricePackagePriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PricePackagePriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                       `json:"minimum_amount,required"`
	JSON          pricePackagePriceMinimumJSON `json:"-"`
}

// pricePackagePriceMinimumJSON contains the JSON metadata for the struct
// [PricePackagePriceMinimum]
type pricePackagePriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PricePackagePriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PricePackagePriceModelType string

const (
	PricePackagePriceModelTypePackage PricePackagePriceModelType = "package"
)

func (r PricePackagePriceModelType) IsKnown() bool {
	switch r {
	case PricePackagePriceModelTypePackage:
		return true
	}
	return false
}

type PricePackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount string `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize int64                              `json:"package_size,required"`
	JSON        pricePackagePricePackageConfigJSON `json:"-"`
}

// pricePackagePricePackageConfigJSON contains the JSON metadata for the struct
// [PricePackagePricePackageConfig]
type pricePackagePricePackageConfigJSON struct {
	PackageAmount apijson.Field
	PackageSize   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PricePackagePricePackageConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackagePricePackageConfigJSON) RawJSON() string {
	return r.raw
}

type PricePackagePricePriceType string

const (
	PricePackagePricePriceTypeUsagePrice PricePackagePricePriceType = "usage_price"
	PricePackagePricePriceTypeFixedPrice PricePackagePricePriceType = "fixed_price"
)

func (r PricePackagePricePriceType) IsKnown() bool {
	switch r {
	case PricePackagePricePriceTypeUsagePrice, PricePackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixPrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              PriceMatrixPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceMatrixPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceMatrixPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceMatrixPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    shared.Discount                             `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceMatrixPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceMatrixPriceItem                        `json:"item,required"`
	MatrixConfig                PriceMatrixPriceMatrixConfig                `json:"matrix_config,required"`
	Maximum                     PriceMatrixPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string         `json:"metadata,required"`
	Minimum        PriceMatrixPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                    `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixPriceModelType `json:"model_type,required"`
	Name           string                    `json:"name,required"`
	PlanPhaseOrder int64                     `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixPricePriceType `json:"price_type,required"`
	JSON           priceMatrixPriceJSON      `json:"-"`
}

// priceMatrixPriceJSON contains the JSON metadata for the struct
// [PriceMatrixPrice]
type priceMatrixPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	MatrixConfig                apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceMatrixPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixPrice) implementsPrice() {}

type PriceMatrixPriceBillableMetric struct {
	ID   string                             `json:"id,required"`
	JSON priceMatrixPriceBillableMetricJSON `json:"-"`
}

// priceMatrixPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceMatrixPriceBillableMetric]
type priceMatrixPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceBillingCycleConfiguration struct {
	Duration     int64                                                 `json:"duration,required"`
	DurationUnit PriceMatrixPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixPriceBillingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceMatrixPriceBillingCycleConfiguration]
type priceMatrixPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PriceMatrixPriceBillingCycleConfigurationDurationUnitDay   PriceMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PriceMatrixPriceBillingCycleConfigurationDurationUnitMonth PriceMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixPriceBillingCycleConfigurationDurationUnitDay, PriceMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixPriceCadence string

const (
	PriceMatrixPriceCadenceOneTime    PriceMatrixPriceCadence = "one_time"
	PriceMatrixPriceCadenceMonthly    PriceMatrixPriceCadence = "monthly"
	PriceMatrixPriceCadenceQuarterly  PriceMatrixPriceCadence = "quarterly"
	PriceMatrixPriceCadenceSemiAnnual PriceMatrixPriceCadence = "semi_annual"
	PriceMatrixPriceCadenceAnnual     PriceMatrixPriceCadence = "annual"
	PriceMatrixPriceCadenceCustom     PriceMatrixPriceCadence = "custom"
)

func (r PriceMatrixPriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixPriceCadenceOneTime, PriceMatrixPriceCadenceMonthly, PriceMatrixPriceCadenceQuarterly, PriceMatrixPriceCadenceSemiAnnual, PriceMatrixPriceCadenceAnnual, PriceMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixPriceCreditAllocation struct {
	AllowsRollover bool                                 `json:"allows_rollover,required"`
	Currency       string                               `json:"currency,required"`
	JSON           priceMatrixPriceCreditAllocationJSON `json:"-"`
}

// priceMatrixPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceMatrixPriceCreditAllocation]
type priceMatrixPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceMatrixPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                   `json:"duration,required"`
	DurationUnit PriceMatrixPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixPriceInvoicingCycleConfigurationJSON contains the JSON metadata for
// the struct [PriceMatrixPriceInvoicingCycleConfiguration]
type priceMatrixPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PriceMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PriceMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PriceMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixPriceItem struct {
	ID   string                   `json:"id,required"`
	Name string                   `json:"name,required"`
	JSON priceMatrixPriceItemJSON `json:"-"`
}

// priceMatrixPriceItemJSON contains the JSON metadata for the struct
// [PriceMatrixPriceItem]
type priceMatrixPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []PriceMatrixPriceMatrixConfigMatrixValue `json:"matrix_values,required"`
	JSON         priceMatrixPriceMatrixConfigJSON          `json:"-"`
}

// priceMatrixPriceMatrixConfigJSON contains the JSON metadata for the struct
// [PriceMatrixPriceMatrixConfig]
type priceMatrixPriceMatrixConfigJSON struct {
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixPriceMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string                                      `json:"unit_amount,required"`
	JSON       priceMatrixPriceMatrixConfigMatrixValueJSON `json:"-"`
}

// priceMatrixPriceMatrixConfigMatrixValueJSON contains the JSON metadata for the
// struct [PriceMatrixPriceMatrixConfigMatrixValue]
type priceMatrixPriceMatrixConfigMatrixValueJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PriceMatrixPriceMatrixConfigMatrixValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceMatrixConfigMatrixValueJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                      `json:"maximum_amount,required"`
	JSON          priceMatrixPriceMaximumJSON `json:"-"`
}

// priceMatrixPriceMaximumJSON contains the JSON metadata for the struct
// [PriceMatrixPriceMaximum]
type priceMatrixPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                      `json:"minimum_amount,required"`
	JSON          priceMatrixPriceMinimumJSON `json:"-"`
}

// priceMatrixPriceMinimumJSON contains the JSON metadata for the struct
// [PriceMatrixPriceMinimum]
type priceMatrixPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceModelType string

const (
	PriceMatrixPriceModelTypeMatrix PriceMatrixPriceModelType = "matrix"
)

func (r PriceMatrixPriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type PriceMatrixPricePriceType string

const (
	PriceMatrixPricePriceTypeUsagePrice PriceMatrixPricePriceType = "usage_price"
	PriceMatrixPricePriceTypeFixedPrice PriceMatrixPricePriceType = "fixed_price"
)

func (r PriceMatrixPricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixPricePriceTypeUsagePrice, PriceMatrixPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPrice struct {
	ID                          string                                      `json:"id,required"`
	BillableMetric              PriceTieredPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                     `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                   `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                      `json:"currency,required"`
	Discount                    shared.Discount                             `json:"discount,required,nullable"`
	ExternalPriceID             string                                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                     `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredPriceItem                        `json:"item,required"`
	Maximum                     PriceTieredPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                      `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string            `json:"metadata,required"`
	Minimum        PriceTieredPriceMinimum      `json:"minimum,required,nullable"`
	MinimumAmount  string                       `json:"minimum_amount,required,nullable"`
	ModelType      PriceTieredPriceModelType    `json:"model_type,required"`
	Name           string                       `json:"name,required"`
	PlanPhaseOrder int64                        `json:"plan_phase_order,required,nullable"`
	PriceType      PriceTieredPricePriceType    `json:"price_type,required"`
	TieredConfig   PriceTieredPriceTieredConfig `json:"tiered_config,required"`
	JSON           priceTieredPriceJSON         `json:"-"`
}

// priceTieredPriceJSON contains the JSON metadata for the struct
// [PriceTieredPrice]
type priceTieredPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	TieredConfig                apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPrice) implementsPrice() {}

type PriceTieredPriceBillableMetric struct {
	ID   string                             `json:"id,required"`
	JSON priceTieredPriceBillableMetricJSON `json:"-"`
}

// priceTieredPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceTieredPriceBillableMetric]
type priceTieredPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceBillingCycleConfiguration struct {
	Duration     int64                                                 `json:"duration,required"`
	DurationUnit PriceTieredPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPriceBillingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceTieredPriceBillingCycleConfiguration]
type priceTieredPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredPriceBillingCycleConfigurationDurationUnitDay   PriceTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredPriceBillingCycleConfigurationDurationUnitMonth PriceTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPriceBillingCycleConfigurationDurationUnitDay, PriceTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPriceCadence string

const (
	PriceTieredPriceCadenceOneTime    PriceTieredPriceCadence = "one_time"
	PriceTieredPriceCadenceMonthly    PriceTieredPriceCadence = "monthly"
	PriceTieredPriceCadenceQuarterly  PriceTieredPriceCadence = "quarterly"
	PriceTieredPriceCadenceSemiAnnual PriceTieredPriceCadence = "semi_annual"
	PriceTieredPriceCadenceAnnual     PriceTieredPriceCadence = "annual"
	PriceTieredPriceCadenceCustom     PriceTieredPriceCadence = "custom"
)

func (r PriceTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPriceCadenceOneTime, PriceTieredPriceCadenceMonthly, PriceTieredPriceCadenceQuarterly, PriceTieredPriceCadenceSemiAnnual, PriceTieredPriceCadenceAnnual, PriceTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPriceCreditAllocation struct {
	AllowsRollover bool                                 `json:"allows_rollover,required"`
	Currency       string                               `json:"currency,required"`
	JSON           priceTieredPriceCreditAllocationJSON `json:"-"`
}

// priceTieredPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceTieredPriceCreditAllocation]
type priceTieredPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                   `json:"duration,required"`
	DurationUnit PriceTieredPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPriceInvoicingCycleConfigurationJSON contains the JSON metadata for
// the struct [PriceTieredPriceInvoicingCycleConfiguration]
type priceTieredPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredPriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredPriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPriceItem struct {
	ID   string                   `json:"id,required"`
	Name string                   `json:"name,required"`
	JSON priceTieredPriceItemJSON `json:"-"`
}

// priceTieredPriceItemJSON contains the JSON metadata for the struct
// [PriceTieredPriceItem]
type priceTieredPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                      `json:"maximum_amount,required"`
	JSON          priceTieredPriceMaximumJSON `json:"-"`
}

// priceTieredPriceMaximumJSON contains the JSON metadata for the struct
// [PriceTieredPriceMaximum]
type priceTieredPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                      `json:"minimum_amount,required"`
	JSON          priceTieredPriceMinimumJSON `json:"-"`
}

// priceTieredPriceMinimumJSON contains the JSON metadata for the struct
// [PriceTieredPriceMinimum]
type priceTieredPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceModelType string

const (
	PriceTieredPriceModelTypeTiered PriceTieredPriceModelType = "tiered"
)

func (r PriceTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type PriceTieredPricePriceType string

const (
	PriceTieredPricePriceTypeUsagePrice PriceTieredPricePriceType = "usage_price"
	PriceTieredPricePriceTypeFixedPrice PriceTieredPricePriceType = "fixed_price"
)

func (r PriceTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPricePriceTypeUsagePrice, PriceTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers []PriceTieredPriceTieredConfigTier `json:"tiers,required"`
	JSON  priceTieredPriceTieredConfigJSON   `json:"-"`
}

// priceTieredPriceTieredConfigJSON contains the JSON metadata for the struct
// [PriceTieredPriceTieredConfig]
type priceTieredPriceTieredConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPriceTieredConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceTieredConfigJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit float64 `json:"first_unit,required"`
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit float64                              `json:"last_unit,nullable"`
	JSON     priceTieredPriceTieredConfigTierJSON `json:"-"`
}

// priceTieredPriceTieredConfigTierJSON contains the JSON metadata for the struct
// [PriceTieredPriceTieredConfigTier]
type priceTieredPriceTieredConfigTierJSON struct {
	FirstUnit   apijson.Field
	UnitAmount  apijson.Field
	LastUnit    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPriceTieredConfigTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPriceTieredConfigTierJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPrice struct {
	ID                          string                                         `json:"id,required"`
	BillableMetric              PriceTieredBpsPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredBpsPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredBpsPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                        `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                      `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredBpsPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                         `json:"currency,required"`
	Discount                    shared.Discount                                `json:"discount,required,nullable"`
	ExternalPriceID             string                                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                        `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredBpsPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredBpsPriceItem                        `json:"item,required"`
	Maximum                     PriceTieredBpsPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata        map[string]string                  `json:"metadata,required"`
	Minimum         PriceTieredBpsPriceMinimum         `json:"minimum,required,nullable"`
	MinimumAmount   string                             `json:"minimum_amount,required,nullable"`
	ModelType       PriceTieredBpsPriceModelType       `json:"model_type,required"`
	Name            string                             `json:"name,required"`
	PlanPhaseOrder  int64                              `json:"plan_phase_order,required,nullable"`
	PriceType       PriceTieredBpsPricePriceType       `json:"price_type,required"`
	TieredBpsConfig PriceTieredBpsPriceTieredBpsConfig `json:"tiered_bps_config,required"`
	JSON            priceTieredBpsPriceJSON            `json:"-"`
}

// priceTieredBpsPriceJSON contains the JSON metadata for the struct
// [PriceTieredBpsPrice]
type priceTieredBpsPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	TieredBpsConfig             apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceTieredBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredBpsPrice) implementsPrice() {}

type PriceTieredBpsPriceBillableMetric struct {
	ID   string                                `json:"id,required"`
	JSON priceTieredBpsPriceBillableMetricJSON `json:"-"`
}

// priceTieredBpsPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceTieredBpsPriceBillableMetric]
type priceTieredBpsPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredBpsPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceBillingCycleConfiguration struct {
	Duration     int64                                                    `json:"duration,required"`
	DurationUnit PriceTieredBpsPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredBpsPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredBpsPriceBillingCycleConfigurationJSON contains the JSON metadata for
// the struct [PriceTieredBpsPriceBillingCycleConfiguration]
type priceTieredBpsPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredBpsPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PriceTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PriceTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PriceTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredBpsPriceCadence string

const (
	PriceTieredBpsPriceCadenceOneTime    PriceTieredBpsPriceCadence = "one_time"
	PriceTieredBpsPriceCadenceMonthly    PriceTieredBpsPriceCadence = "monthly"
	PriceTieredBpsPriceCadenceQuarterly  PriceTieredBpsPriceCadence = "quarterly"
	PriceTieredBpsPriceCadenceSemiAnnual PriceTieredBpsPriceCadence = "semi_annual"
	PriceTieredBpsPriceCadenceAnnual     PriceTieredBpsPriceCadence = "annual"
	PriceTieredBpsPriceCadenceCustom     PriceTieredBpsPriceCadence = "custom"
)

func (r PriceTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredBpsPriceCadenceOneTime, PriceTieredBpsPriceCadenceMonthly, PriceTieredBpsPriceCadenceQuarterly, PriceTieredBpsPriceCadenceSemiAnnual, PriceTieredBpsPriceCadenceAnnual, PriceTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredBpsPriceCreditAllocation struct {
	AllowsRollover bool                                    `json:"allows_rollover,required"`
	Currency       string                                  `json:"currency,required"`
	JSON           priceTieredBpsPriceCreditAllocationJSON `json:"-"`
}

// priceTieredBpsPriceCreditAllocationJSON contains the JSON metadata for the
// struct [PriceTieredBpsPriceCreditAllocation]
type priceTieredBpsPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredBpsPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                      `json:"duration,required"`
	DurationUnit PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredBpsPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredBpsPriceInvoicingCycleConfigurationJSON contains the JSON metadata
// for the struct [PriceTieredBpsPriceInvoicingCycleConfiguration]
type priceTieredBpsPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredBpsPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredBpsPriceItem struct {
	ID   string                      `json:"id,required"`
	Name string                      `json:"name,required"`
	JSON priceTieredBpsPriceItemJSON `json:"-"`
}

// priceTieredBpsPriceItemJSON contains the JSON metadata for the struct
// [PriceTieredBpsPriceItem]
type priceTieredBpsPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredBpsPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                         `json:"maximum_amount,required"`
	JSON          priceTieredBpsPriceMaximumJSON `json:"-"`
}

// priceTieredBpsPriceMaximumJSON contains the JSON metadata for the struct
// [PriceTieredBpsPriceMaximum]
type priceTieredBpsPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredBpsPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                         `json:"minimum_amount,required"`
	JSON          priceTieredBpsPriceMinimumJSON `json:"-"`
}

// priceTieredBpsPriceMinimumJSON contains the JSON metadata for the struct
// [PriceTieredBpsPriceMinimum]
type priceTieredBpsPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredBpsPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceModelType string

const (
	PriceTieredBpsPriceModelTypeTieredBps PriceTieredBpsPriceModelType = "tiered_bps"
)

func (r PriceTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PriceTieredBpsPricePriceType string

const (
	PriceTieredBpsPricePriceTypeUsagePrice PriceTieredBpsPricePriceType = "usage_price"
	PriceTieredBpsPricePriceTypeFixedPrice PriceTieredBpsPricePriceType = "fixed_price"
)

func (r PriceTieredBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredBpsPricePriceTypeUsagePrice, PriceTieredBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers []PriceTieredBpsPriceTieredBpsConfigTier `json:"tiers,required"`
	JSON  priceTieredBpsPriceTieredBpsConfigJSON   `json:"-"`
}

// priceTieredBpsPriceTieredBpsConfigJSON contains the JSON metadata for the struct
// [PriceTieredBpsPriceTieredBpsConfig]
type priceTieredBpsPriceTieredBpsConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredBpsPriceTieredBpsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceTieredBpsConfigJSON) RawJSON() string {
	return r.raw
}

type PriceTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps float64 `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount string `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount string `json:"maximum_amount,nullable"`
	// Per unit maximum to charge
	PerUnitMaximum string                                     `json:"per_unit_maximum,nullable"`
	JSON           priceTieredBpsPriceTieredBpsConfigTierJSON `json:"-"`
}

// priceTieredBpsPriceTieredBpsConfigTierJSON contains the JSON metadata for the
// struct [PriceTieredBpsPriceTieredBpsConfigTier]
type priceTieredBpsPriceTieredBpsConfigTierJSON struct {
	Bps            apijson.Field
	MinimumAmount  apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredBpsPriceTieredBpsConfigTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredBpsPriceTieredBpsConfigTierJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPrice struct {
	ID                          string                                   `json:"id,required"`
	BillableMetric              PriceBpsPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceBpsPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	BpsConfig                   PriceBpsPriceBpsConfig                   `json:"bps_config,required"`
	Cadence                     PriceBpsPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                  `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceBpsPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                   `json:"currency,required"`
	Discount                    shared.Discount                          `json:"discount,required,nullable"`
	ExternalPriceID             string                                   `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                  `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceBpsPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceBpsPriceItem                        `json:"item,required"`
	Maximum                     PriceBpsPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                   `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string      `json:"metadata,required"`
	Minimum        PriceBpsPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                 `json:"minimum_amount,required,nullable"`
	ModelType      PriceBpsPriceModelType `json:"model_type,required"`
	Name           string                 `json:"name,required"`
	PlanPhaseOrder int64                  `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBpsPricePriceType `json:"price_type,required"`
	JSON           priceBpsPriceJSON      `json:"-"`
}

// priceBpsPriceJSON contains the JSON metadata for the struct [PriceBpsPrice]
type priceBpsPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	BpsConfig                   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBpsPrice) implementsPrice() {}

type PriceBpsPriceBillableMetric struct {
	ID   string                          `json:"id,required"`
	JSON priceBpsPriceBillableMetricJSON `json:"-"`
}

// priceBpsPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceBpsPriceBillableMetric]
type priceBpsPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBpsPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceBillingCycleConfiguration struct {
	Duration     int64                                              `json:"duration,required"`
	DurationUnit PriceBpsPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBpsPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceBpsPriceBillingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceBpsPriceBillingCycleConfiguration]
type priceBpsPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBpsPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceBpsPriceBillingCycleConfigurationDurationUnitDay   PriceBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceBpsPriceBillingCycleConfigurationDurationUnitMonth PriceBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBpsPriceBillingCycleConfigurationDurationUnitDay, PriceBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps float64 `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum string                     `json:"per_unit_maximum,nullable"`
	JSON           priceBpsPriceBpsConfigJSON `json:"-"`
}

// priceBpsPriceBpsConfigJSON contains the JSON metadata for the struct
// [PriceBpsPriceBpsConfig]
type priceBpsPriceBpsConfigJSON struct {
	Bps            apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBpsPriceBpsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceBpsConfigJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceCadence string

const (
	PriceBpsPriceCadenceOneTime    PriceBpsPriceCadence = "one_time"
	PriceBpsPriceCadenceMonthly    PriceBpsPriceCadence = "monthly"
	PriceBpsPriceCadenceQuarterly  PriceBpsPriceCadence = "quarterly"
	PriceBpsPriceCadenceSemiAnnual PriceBpsPriceCadence = "semi_annual"
	PriceBpsPriceCadenceAnnual     PriceBpsPriceCadence = "annual"
	PriceBpsPriceCadenceCustom     PriceBpsPriceCadence = "custom"
)

func (r PriceBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceBpsPriceCadenceOneTime, PriceBpsPriceCadenceMonthly, PriceBpsPriceCadenceQuarterly, PriceBpsPriceCadenceSemiAnnual, PriceBpsPriceCadenceAnnual, PriceBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBpsPriceCreditAllocation struct {
	AllowsRollover bool                              `json:"allows_rollover,required"`
	Currency       string                            `json:"currency,required"`
	JSON           priceBpsPriceCreditAllocationJSON `json:"-"`
}

// priceBpsPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceBpsPriceCreditAllocation]
type priceBpsPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBpsPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                `json:"duration,required"`
	DurationUnit PriceBpsPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBpsPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceBpsPriceInvoicingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceBpsPriceInvoicingCycleConfiguration]
type priceBpsPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBpsPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBpsPriceItem struct {
	ID   string                `json:"id,required"`
	Name string                `json:"name,required"`
	JSON priceBpsPriceItemJSON `json:"-"`
}

// priceBpsPriceItemJSON contains the JSON metadata for the struct
// [PriceBpsPriceItem]
type priceBpsPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBpsPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                   `json:"maximum_amount,required"`
	JSON          priceBpsPriceMaximumJSON `json:"-"`
}

// priceBpsPriceMaximumJSON contains the JSON metadata for the struct
// [PriceBpsPriceMaximum]
type priceBpsPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBpsPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                   `json:"minimum_amount,required"`
	JSON          priceBpsPriceMinimumJSON `json:"-"`
}

// priceBpsPriceMinimumJSON contains the JSON metadata for the struct
// [PriceBpsPriceMinimum]
type priceBpsPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBpsPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBpsPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceBpsPriceModelType string

const (
	PriceBpsPriceModelTypeBps PriceBpsPriceModelType = "bps"
)

func (r PriceBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceBpsPriceModelTypeBps:
		return true
	}
	return false
}

type PriceBpsPricePriceType string

const (
	PriceBpsPricePriceTypeUsagePrice PriceBpsPricePriceType = "usage_price"
	PriceBpsPricePriceTypeFixedPrice PriceBpsPricePriceType = "fixed_price"
)

func (r PriceBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceBpsPricePriceTypeUsagePrice, PriceBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkBpsPrice struct {
	ID                          string                                       `json:"id,required"`
	BillableMetric              PriceBulkBpsPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceBulkBpsPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	BulkBpsConfig               PriceBulkBpsPriceBulkBpsConfig               `json:"bulk_bps_config,required"`
	Cadence                     PriceBulkBpsPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                      `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                    `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceBulkBpsPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                       `json:"currency,required"`
	Discount                    shared.Discount                              `json:"discount,required,nullable"`
	ExternalPriceID             string                                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                      `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceBulkBpsPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceBulkBpsPriceItem                        `json:"item,required"`
	Maximum                     PriceBulkBpsPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                       `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string          `json:"metadata,required"`
	Minimum        PriceBulkBpsPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                     `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkBpsPriceModelType `json:"model_type,required"`
	Name           string                     `json:"name,required"`
	PlanPhaseOrder int64                      `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkBpsPricePriceType `json:"price_type,required"`
	JSON           priceBulkBpsPriceJSON      `json:"-"`
}

// priceBulkBpsPriceJSON contains the JSON metadata for the struct
// [PriceBulkBpsPrice]
type priceBulkBpsPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	BulkBpsConfig               apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceBulkBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkBpsPrice) implementsPrice() {}

type PriceBulkBpsPriceBillableMetric struct {
	ID   string                              `json:"id,required"`
	JSON priceBulkBpsPriceBillableMetricJSON `json:"-"`
}

// priceBulkBpsPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceBillableMetric]
type priceBulkBpsPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkBpsPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceBillingCycleConfiguration struct {
	Duration     int64                                                  `json:"duration,required"`
	DurationUnit PriceBulkBpsPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkBpsPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceBulkBpsPriceBillingCycleConfigurationJSON contains the JSON metadata for
// the struct [PriceBulkBpsPriceBillingCycleConfiguration]
type priceBulkBpsPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkBpsPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PriceBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PriceBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PriceBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers []PriceBulkBpsPriceBulkBpsConfigTier `json:"tiers,required"`
	JSON  priceBulkBpsPriceBulkBpsConfigJSON   `json:"-"`
}

// priceBulkBpsPriceBulkBpsConfigJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceBulkBpsConfig]
type priceBulkBpsPriceBulkBpsConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkBpsPriceBulkBpsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceBulkBpsConfigJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps float64 `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount string `json:"maximum_amount,nullable"`
	// The maximum amount to charge for any one event
	PerUnitMaximum string                                 `json:"per_unit_maximum,nullable"`
	JSON           priceBulkBpsPriceBulkBpsConfigTierJSON `json:"-"`
}

// priceBulkBpsPriceBulkBpsConfigTierJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceBulkBpsConfigTier]
type priceBulkBpsPriceBulkBpsConfigTierJSON struct {
	Bps            apijson.Field
	MaximumAmount  apijson.Field
	PerUnitMaximum apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBulkBpsPriceBulkBpsConfigTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceBulkBpsConfigTierJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceCadence string

const (
	PriceBulkBpsPriceCadenceOneTime    PriceBulkBpsPriceCadence = "one_time"
	PriceBulkBpsPriceCadenceMonthly    PriceBulkBpsPriceCadence = "monthly"
	PriceBulkBpsPriceCadenceQuarterly  PriceBulkBpsPriceCadence = "quarterly"
	PriceBulkBpsPriceCadenceSemiAnnual PriceBulkBpsPriceCadence = "semi_annual"
	PriceBulkBpsPriceCadenceAnnual     PriceBulkBpsPriceCadence = "annual"
	PriceBulkBpsPriceCadenceCustom     PriceBulkBpsPriceCadence = "custom"
)

func (r PriceBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkBpsPriceCadenceOneTime, PriceBulkBpsPriceCadenceMonthly, PriceBulkBpsPriceCadenceQuarterly, PriceBulkBpsPriceCadenceSemiAnnual, PriceBulkBpsPriceCadenceAnnual, PriceBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkBpsPriceCreditAllocation struct {
	AllowsRollover bool                                  `json:"allows_rollover,required"`
	Currency       string                                `json:"currency,required"`
	JSON           priceBulkBpsPriceCreditAllocationJSON `json:"-"`
}

// priceBulkBpsPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceCreditAllocation]
type priceBulkBpsPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBulkBpsPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                    `json:"duration,required"`
	DurationUnit PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkBpsPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceBulkBpsPriceInvoicingCycleConfigurationJSON contains the JSON metadata for
// the struct [PriceBulkBpsPriceInvoicingCycleConfiguration]
type priceBulkBpsPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkBpsPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkBpsPriceItem struct {
	ID   string                    `json:"id,required"`
	Name string                    `json:"name,required"`
	JSON priceBulkBpsPriceItemJSON `json:"-"`
}

// priceBulkBpsPriceItemJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceItem]
type priceBulkBpsPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkBpsPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                       `json:"maximum_amount,required"`
	JSON          priceBulkBpsPriceMaximumJSON `json:"-"`
}

// priceBulkBpsPriceMaximumJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceMaximum]
type priceBulkBpsPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkBpsPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                       `json:"minimum_amount,required"`
	JSON          priceBulkBpsPriceMinimumJSON `json:"-"`
}

// priceBulkBpsPriceMinimumJSON contains the JSON metadata for the struct
// [PriceBulkBpsPriceMinimum]
type priceBulkBpsPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkBpsPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkBpsPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkBpsPriceModelType string

const (
	PriceBulkBpsPriceModelTypeBulkBps PriceBulkBpsPriceModelType = "bulk_bps"
)

func (r PriceBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

type PriceBulkBpsPricePriceType string

const (
	PriceBulkBpsPricePriceTypeUsagePrice PriceBulkBpsPricePriceType = "usage_price"
	PriceBulkBpsPricePriceTypeFixedPrice PriceBulkBpsPricePriceType = "fixed_price"
)

func (r PriceBulkBpsPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkBpsPricePriceTypeUsagePrice, PriceBulkBpsPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkPrice struct {
	ID                          string                                    `json:"id,required"`
	BillableMetric              PriceBulkPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceBulkPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	BulkConfig                  PriceBulkPriceBulkConfig                  `json:"bulk_config,required"`
	Cadence                     PriceBulkPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceBulkPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                    `json:"currency,required"`
	Discount                    shared.Discount                           `json:"discount,required,nullable"`
	ExternalPriceID             string                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceBulkPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceBulkPriceItem                        `json:"item,required"`
	Maximum                     PriceBulkPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string       `json:"metadata,required"`
	Minimum        PriceBulkPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkPriceModelType `json:"model_type,required"`
	Name           string                  `json:"name,required"`
	PlanPhaseOrder int64                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkPricePriceType `json:"price_type,required"`
	JSON           priceBulkPriceJSON      `json:"-"`
}

// priceBulkPriceJSON contains the JSON metadata for the struct [PriceBulkPrice]
type priceBulkPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	BulkConfig                  apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkPrice) implementsPrice() {}

type PriceBulkPriceBillableMetric struct {
	ID   string                           `json:"id,required"`
	JSON priceBulkPriceBillableMetricJSON `json:"-"`
}

// priceBulkPriceBillableMetricJSON contains the JSON metadata for the struct
// [PriceBulkPriceBillableMetric]
type priceBulkPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceBillingCycleConfiguration struct {
	Duration     int64                                               `json:"duration,required"`
	DurationUnit PriceBulkPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceBulkPriceBillingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceBulkPriceBillingCycleConfiguration]
type priceBulkPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PriceBulkPriceBillingCycleConfigurationDurationUnitDay   PriceBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PriceBulkPriceBillingCycleConfigurationDurationUnitMonth PriceBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkPriceBillingCycleConfigurationDurationUnitDay, PriceBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers []PriceBulkPriceBulkConfigTier `json:"tiers,required"`
	JSON  priceBulkPriceBulkConfigJSON   `json:"-"`
}

// priceBulkPriceBulkConfigJSON contains the JSON metadata for the struct
// [PriceBulkPriceBulkConfig]
type priceBulkPriceBulkConfigJSON struct {
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkPriceBulkConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceBulkConfigJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits float64                          `json:"maximum_units,nullable"`
	JSON         priceBulkPriceBulkConfigTierJSON `json:"-"`
}

// priceBulkPriceBulkConfigTierJSON contains the JSON metadata for the struct
// [PriceBulkPriceBulkConfigTier]
type priceBulkPriceBulkConfigTierJSON struct {
	UnitAmount   apijson.Field
	MaximumUnits apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkPriceBulkConfigTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceBulkConfigTierJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceCadence string

const (
	PriceBulkPriceCadenceOneTime    PriceBulkPriceCadence = "one_time"
	PriceBulkPriceCadenceMonthly    PriceBulkPriceCadence = "monthly"
	PriceBulkPriceCadenceQuarterly  PriceBulkPriceCadence = "quarterly"
	PriceBulkPriceCadenceSemiAnnual PriceBulkPriceCadence = "semi_annual"
	PriceBulkPriceCadenceAnnual     PriceBulkPriceCadence = "annual"
	PriceBulkPriceCadenceCustom     PriceBulkPriceCadence = "custom"
)

func (r PriceBulkPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkPriceCadenceOneTime, PriceBulkPriceCadenceMonthly, PriceBulkPriceCadenceQuarterly, PriceBulkPriceCadenceSemiAnnual, PriceBulkPriceCadenceAnnual, PriceBulkPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkPriceCreditAllocation struct {
	AllowsRollover bool                               `json:"allows_rollover,required"`
	Currency       string                             `json:"currency,required"`
	JSON           priceBulkPriceCreditAllocationJSON `json:"-"`
}

// priceBulkPriceCreditAllocationJSON contains the JSON metadata for the struct
// [PriceBulkPriceCreditAllocation]
type priceBulkPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBulkPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                 `json:"duration,required"`
	DurationUnit PriceBulkPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceBulkPriceInvoicingCycleConfigurationJSON contains the JSON metadata for the
// struct [PriceBulkPriceInvoicingCycleConfiguration]
type priceBulkPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceBulkPriceInvoicingCycleConfigurationDurationUnitDay   PriceBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceBulkPriceInvoicingCycleConfigurationDurationUnitMonth PriceBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkPriceInvoicingCycleConfigurationDurationUnitDay, PriceBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkPriceItem struct {
	ID   string                 `json:"id,required"`
	Name string                 `json:"name,required"`
	JSON priceBulkPriceItemJSON `json:"-"`
}

// priceBulkPriceItemJSON contains the JSON metadata for the struct
// [PriceBulkPriceItem]
type priceBulkPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                    `json:"maximum_amount,required"`
	JSON          priceBulkPriceMaximumJSON `json:"-"`
}

// priceBulkPriceMaximumJSON contains the JSON metadata for the struct
// [PriceBulkPriceMaximum]
type priceBulkPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                    `json:"minimum_amount,required"`
	JSON          priceBulkPriceMinimumJSON `json:"-"`
}

// priceBulkPriceMinimumJSON contains the JSON metadata for the struct
// [PriceBulkPriceMinimum]
type priceBulkPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkPriceModelType string

const (
	PriceBulkPriceModelTypeBulk PriceBulkPriceModelType = "bulk"
)

func (r PriceBulkPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type PriceBulkPricePriceType string

const (
	PriceBulkPricePriceTypeUsagePrice PriceBulkPricePriceType = "usage_price"
	PriceBulkPricePriceTypeFixedPrice PriceBulkPricePriceType = "fixed_price"
)

func (r PriceBulkPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkPricePriceTypeUsagePrice, PriceBulkPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPrice struct {
	ID                          string                                                    `json:"id,required"`
	BillableMetric              PriceThresholdTotalAmountPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceThresholdTotalAmountPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceThresholdTotalAmountPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceThresholdTotalAmountPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                    `json:"currency,required"`
	Discount                    shared.Discount                                           `json:"discount,required,nullable"`
	ExternalPriceID             string                                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceThresholdTotalAmountPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceThresholdTotalAmountPriceItem                        `json:"item,required"`
	Maximum                     PriceThresholdTotalAmountPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                   map[string]string                       `json:"metadata,required"`
	Minimum                    PriceThresholdTotalAmountPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount              string                                  `json:"minimum_amount,required,nullable"`
	ModelType                  PriceThresholdTotalAmountPriceModelType `json:"model_type,required"`
	Name                       string                                  `json:"name,required"`
	PlanPhaseOrder             int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType                  PriceThresholdTotalAmountPricePriceType `json:"price_type,required"`
	ThresholdTotalAmountConfig map[string]interface{}                  `json:"threshold_total_amount_config,required"`
	JSON                       priceThresholdTotalAmountPriceJSON      `json:"-"`
}

// priceThresholdTotalAmountPriceJSON contains the JSON metadata for the struct
// [PriceThresholdTotalAmountPrice]
type priceThresholdTotalAmountPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	ThresholdTotalAmountConfig  apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceThresholdTotalAmountPrice) implementsPrice() {}

type PriceThresholdTotalAmountPriceBillableMetric struct {
	ID   string                                           `json:"id,required"`
	JSON priceThresholdTotalAmountPriceBillableMetricJSON `json:"-"`
}

// priceThresholdTotalAmountPriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceThresholdTotalAmountPriceBillableMetric]
type priceThresholdTotalAmountPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceBillingCycleConfiguration struct {
	Duration     int64                                                               `json:"duration,required"`
	DurationUnit PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceThresholdTotalAmountPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceThresholdTotalAmountPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceThresholdTotalAmountPriceBillingCycleConfiguration]
type priceThresholdTotalAmountPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PriceThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPriceCadence string

const (
	PriceThresholdTotalAmountPriceCadenceOneTime    PriceThresholdTotalAmountPriceCadence = "one_time"
	PriceThresholdTotalAmountPriceCadenceMonthly    PriceThresholdTotalAmountPriceCadence = "monthly"
	PriceThresholdTotalAmountPriceCadenceQuarterly  PriceThresholdTotalAmountPriceCadence = "quarterly"
	PriceThresholdTotalAmountPriceCadenceSemiAnnual PriceThresholdTotalAmountPriceCadence = "semi_annual"
	PriceThresholdTotalAmountPriceCadenceAnnual     PriceThresholdTotalAmountPriceCadence = "annual"
	PriceThresholdTotalAmountPriceCadenceCustom     PriceThresholdTotalAmountPriceCadence = "custom"
)

func (r PriceThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceCadenceOneTime, PriceThresholdTotalAmountPriceCadenceMonthly, PriceThresholdTotalAmountPriceCadenceQuarterly, PriceThresholdTotalAmountPriceCadenceSemiAnnual, PriceThresholdTotalAmountPriceCadenceAnnual, PriceThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPriceCreditAllocation struct {
	AllowsRollover bool                                               `json:"allows_rollover,required"`
	Currency       string                                             `json:"currency,required"`
	JSON           priceThresholdTotalAmountPriceCreditAllocationJSON `json:"-"`
}

// priceThresholdTotalAmountPriceCreditAllocationJSON contains the JSON metadata
// for the struct [PriceThresholdTotalAmountPriceCreditAllocation]
type priceThresholdTotalAmountPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                 `json:"duration,required"`
	DurationUnit PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceThresholdTotalAmountPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceThresholdTotalAmountPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceThresholdTotalAmountPriceInvoicingCycleConfiguration]
type priceThresholdTotalAmountPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PriceThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPriceItem struct {
	ID   string                                 `json:"id,required"`
	Name string                                 `json:"name,required"`
	JSON priceThresholdTotalAmountPriceItemJSON `json:"-"`
}

// priceThresholdTotalAmountPriceItemJSON contains the JSON metadata for the struct
// [PriceThresholdTotalAmountPriceItem]
type priceThresholdTotalAmountPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                    `json:"maximum_amount,required"`
	JSON          priceThresholdTotalAmountPriceMaximumJSON `json:"-"`
}

// priceThresholdTotalAmountPriceMaximumJSON contains the JSON metadata for the
// struct [PriceThresholdTotalAmountPriceMaximum]
type priceThresholdTotalAmountPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                    `json:"minimum_amount,required"`
	JSON          priceThresholdTotalAmountPriceMinimumJSON `json:"-"`
}

// priceThresholdTotalAmountPriceMinimumJSON contains the JSON metadata for the
// struct [PriceThresholdTotalAmountPriceMinimum]
type priceThresholdTotalAmountPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceThresholdTotalAmountPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceThresholdTotalAmountPriceModelType string

const (
	PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r PriceThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

type PriceThresholdTotalAmountPricePriceType string

const (
	PriceThresholdTotalAmountPricePriceTypeUsagePrice PriceThresholdTotalAmountPricePriceType = "usage_price"
	PriceThresholdTotalAmountPricePriceTypeFixedPrice PriceThresholdTotalAmountPricePriceType = "fixed_price"
)

func (r PriceThresholdTotalAmountPricePriceType) IsKnown() bool {
	switch r {
	case PriceThresholdTotalAmountPricePriceTypeUsagePrice, PriceThresholdTotalAmountPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPackagePrice struct {
	ID                          string                                             `json:"id,required"`
	BillableMetric              PriceTieredPackagePriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredPackagePriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredPackagePriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                            `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                          `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredPackagePriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                             `json:"currency,required"`
	Discount                    shared.Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                             `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                            `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredPackagePriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredPackagePriceItem                        `json:"item,required"`
	Maximum                     PriceTieredPackagePriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                             `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata            map[string]string                `json:"metadata,required"`
	Minimum             PriceTieredPackagePriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount       string                           `json:"minimum_amount,required,nullable"`
	ModelType           PriceTieredPackagePriceModelType `json:"model_type,required"`
	Name                string                           `json:"name,required"`
	PlanPhaseOrder      int64                            `json:"plan_phase_order,required,nullable"`
	PriceType           PriceTieredPackagePricePriceType `json:"price_type,required"`
	TieredPackageConfig map[string]interface{}           `json:"tiered_package_config,required"`
	JSON                priceTieredPackagePriceJSON      `json:"-"`
}

// priceTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceTieredPackagePrice]
type priceTieredPackagePriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	TieredPackageConfig         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPackagePrice) implementsPrice() {}

type PriceTieredPackagePriceBillableMetric struct {
	ID   string                                    `json:"id,required"`
	JSON priceTieredPackagePriceBillableMetricJSON `json:"-"`
}

// priceTieredPackagePriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceTieredPackagePriceBillableMetric]
type priceTieredPackagePriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPackagePriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceBillingCycleConfiguration struct {
	Duration     int64                                                        `json:"duration,required"`
	DurationUnit PriceTieredPackagePriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPackagePriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPackagePriceBillingCycleConfigurationJSON contains the JSON metadata
// for the struct [PriceTieredPackagePriceBillingCycleConfiguration]
type priceTieredPackagePriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPackagePriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PriceTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PriceTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PriceTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPackagePriceCadence string

const (
	PriceTieredPackagePriceCadenceOneTime    PriceTieredPackagePriceCadence = "one_time"
	PriceTieredPackagePriceCadenceMonthly    PriceTieredPackagePriceCadence = "monthly"
	PriceTieredPackagePriceCadenceQuarterly  PriceTieredPackagePriceCadence = "quarterly"
	PriceTieredPackagePriceCadenceSemiAnnual PriceTieredPackagePriceCadence = "semi_annual"
	PriceTieredPackagePriceCadenceAnnual     PriceTieredPackagePriceCadence = "annual"
	PriceTieredPackagePriceCadenceCustom     PriceTieredPackagePriceCadence = "custom"
)

func (r PriceTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceCadenceOneTime, PriceTieredPackagePriceCadenceMonthly, PriceTieredPackagePriceCadenceQuarterly, PriceTieredPackagePriceCadenceSemiAnnual, PriceTieredPackagePriceCadenceAnnual, PriceTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPackagePriceCreditAllocation struct {
	AllowsRollover bool                                        `json:"allows_rollover,required"`
	Currency       string                                      `json:"currency,required"`
	JSON           priceTieredPackagePriceCreditAllocationJSON `json:"-"`
}

// priceTieredPackagePriceCreditAllocationJSON contains the JSON metadata for the
// struct [PriceTieredPackagePriceCreditAllocation]
type priceTieredPackagePriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredPackagePriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceInvoicingCycleConfiguration struct {
	Duration     int64                                                          `json:"duration,required"`
	DurationUnit PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPackagePriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPackagePriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceTieredPackagePriceInvoicingCycleConfiguration]
type priceTieredPackagePriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPackagePriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPackagePriceItem struct {
	ID   string                          `json:"id,required"`
	Name string                          `json:"name,required"`
	JSON priceTieredPackagePriceItemJSON `json:"-"`
}

// priceTieredPackagePriceItemJSON contains the JSON metadata for the struct
// [PriceTieredPackagePriceItem]
type priceTieredPackagePriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPackagePriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                             `json:"maximum_amount,required"`
	JSON          priceTieredPackagePriceMaximumJSON `json:"-"`
}

// priceTieredPackagePriceMaximumJSON contains the JSON metadata for the struct
// [PriceTieredPackagePriceMaximum]
type priceTieredPackagePriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPackagePriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                             `json:"minimum_amount,required"`
	JSON          priceTieredPackagePriceMinimumJSON `json:"-"`
}

// priceTieredPackagePriceMinimumJSON contains the JSON metadata for the struct
// [PriceTieredPackagePriceMinimum]
type priceTieredPackagePriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPackagePriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackagePriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackagePriceModelType string

const (
	PriceTieredPackagePriceModelTypeTieredPackage PriceTieredPackagePriceModelType = "tiered_package"
)

func (r PriceTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

type PriceTieredPackagePricePriceType string

const (
	PriceTieredPackagePricePriceTypeUsagePrice PriceTieredPackagePricePriceType = "usage_price"
	PriceTieredPackagePricePriceTypeFixedPrice PriceTieredPackagePricePriceType = "fixed_price"
)

func (r PriceTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPackagePricePriceTypeUsagePrice, PriceTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedTieredPrice struct {
	ID                          string                                             `json:"id,required"`
	BillableMetric              PriceGroupedTieredPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceGroupedTieredPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceGroupedTieredPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                            `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                          `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceGroupedTieredPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                             `json:"currency,required"`
	Discount                    shared.Discount                                    `json:"discount,required,nullable"`
	ExternalPriceID             string                                             `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                            `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredConfig         map[string]interface{}                             `json:"grouped_tiered_config,required"`
	InvoicingCycleConfiguration PriceGroupedTieredPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceGroupedTieredPriceItem                        `json:"item,required"`
	Maximum                     PriceGroupedTieredPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                             `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                `json:"metadata,required"`
	Minimum        PriceGroupedTieredPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                           `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedTieredPriceModelType `json:"model_type,required"`
	Name           string                           `json:"name,required"`
	PlanPhaseOrder int64                            `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedTieredPricePriceType `json:"price_type,required"`
	JSON           priceGroupedTieredPriceJSON      `json:"-"`
}

// priceGroupedTieredPriceJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPrice]
type priceGroupedTieredPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	GroupedTieredConfig         apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceGroupedTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedTieredPrice) implementsPrice() {}

type PriceGroupedTieredPriceBillableMetric struct {
	ID   string                                    `json:"id,required"`
	JSON priceGroupedTieredPriceBillableMetricJSON `json:"-"`
}

// priceGroupedTieredPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceGroupedTieredPriceBillableMetric]
type priceGroupedTieredPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceBillingCycleConfiguration struct {
	Duration     int64                                                        `json:"duration,required"`
	DurationUnit PriceGroupedTieredPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedTieredPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedTieredPriceBillingCycleConfigurationJSON contains the JSON metadata
// for the struct [PriceGroupedTieredPriceBillingCycleConfiguration]
type priceGroupedTieredPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PriceGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PriceGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PriceGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PriceGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PriceGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedTieredPriceCadence string

const (
	PriceGroupedTieredPriceCadenceOneTime    PriceGroupedTieredPriceCadence = "one_time"
	PriceGroupedTieredPriceCadenceMonthly    PriceGroupedTieredPriceCadence = "monthly"
	PriceGroupedTieredPriceCadenceQuarterly  PriceGroupedTieredPriceCadence = "quarterly"
	PriceGroupedTieredPriceCadenceSemiAnnual PriceGroupedTieredPriceCadence = "semi_annual"
	PriceGroupedTieredPriceCadenceAnnual     PriceGroupedTieredPriceCadence = "annual"
	PriceGroupedTieredPriceCadenceCustom     PriceGroupedTieredPriceCadence = "custom"
)

func (r PriceGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceCadenceOneTime, PriceGroupedTieredPriceCadenceMonthly, PriceGroupedTieredPriceCadenceQuarterly, PriceGroupedTieredPriceCadenceSemiAnnual, PriceGroupedTieredPriceCadenceAnnual, PriceGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedTieredPriceCreditAllocation struct {
	AllowsRollover bool                                        `json:"allows_rollover,required"`
	Currency       string                                      `json:"currency,required"`
	JSON           priceGroupedTieredPriceCreditAllocationJSON `json:"-"`
}

// priceGroupedTieredPriceCreditAllocationJSON contains the JSON metadata for the
// struct [PriceGroupedTieredPriceCreditAllocation]
type priceGroupedTieredPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                          `json:"duration,required"`
	DurationUnit PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedTieredPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedTieredPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceGroupedTieredPriceInvoicingCycleConfiguration]
type priceGroupedTieredPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PriceGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedTieredPriceItem struct {
	ID   string                          `json:"id,required"`
	Name string                          `json:"name,required"`
	JSON priceGroupedTieredPriceItemJSON `json:"-"`
}

// priceGroupedTieredPriceItemJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPriceItem]
type priceGroupedTieredPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                             `json:"maximum_amount,required"`
	JSON          priceGroupedTieredPriceMaximumJSON `json:"-"`
}

// priceGroupedTieredPriceMaximumJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPriceMaximum]
type priceGroupedTieredPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                             `json:"minimum_amount,required"`
	JSON          priceGroupedTieredPriceMinimumJSON `json:"-"`
}

// priceGroupedTieredPriceMinimumJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPriceMinimum]
type priceGroupedTieredPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedTieredPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPriceModelType string

const (
	PriceGroupedTieredPriceModelTypeGroupedTiered PriceGroupedTieredPriceModelType = "grouped_tiered"
)

func (r PriceGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type PriceGroupedTieredPricePriceType string

const (
	PriceGroupedTieredPricePriceTypeUsagePrice PriceGroupedTieredPricePriceType = "usage_price"
	PriceGroupedTieredPricePriceTypeFixedPrice PriceGroupedTieredPricePriceType = "fixed_price"
)

func (r PriceGroupedTieredPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPricePriceTypeUsagePrice, PriceGroupedTieredPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredWithMinimumPrice struct {
	ID                          string                                                 `json:"id,required"`
	BillableMetric              PriceTieredWithMinimumPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredWithMinimumPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredWithMinimumPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                              `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredWithMinimumPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                 `json:"currency,required"`
	Discount                    shared.Discount                                        `json:"discount,required,nullable"`
	ExternalPriceID             string                                                 `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredWithMinimumPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredWithMinimumPriceItem                        `json:"item,required"`
	Maximum                     PriceTieredWithMinimumPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                 `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                map[string]string                    `json:"metadata,required"`
	Minimum                 PriceTieredWithMinimumPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount           string                               `json:"minimum_amount,required,nullable"`
	ModelType               PriceTieredWithMinimumPriceModelType `json:"model_type,required"`
	Name                    string                               `json:"name,required"`
	PlanPhaseOrder          int64                                `json:"plan_phase_order,required,nullable"`
	PriceType               PriceTieredWithMinimumPricePriceType `json:"price_type,required"`
	TieredWithMinimumConfig map[string]interface{}               `json:"tiered_with_minimum_config,required"`
	JSON                    priceTieredWithMinimumPriceJSON      `json:"-"`
}

// priceTieredWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPrice]
type priceTieredWithMinimumPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	TieredWithMinimumConfig     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredWithMinimumPrice) implementsPrice() {}

type PriceTieredWithMinimumPriceBillableMetric struct {
	ID   string                                        `json:"id,required"`
	JSON priceTieredWithMinimumPriceBillableMetricJSON `json:"-"`
}

// priceTieredWithMinimumPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceTieredWithMinimumPriceBillableMetric]
type priceTieredWithMinimumPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceBillingCycleConfiguration struct {
	Duration     int64                                                            `json:"duration,required"`
	DurationUnit PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredWithMinimumPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredWithMinimumPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceTieredWithMinimumPriceBillingCycleConfiguration]
type priceTieredWithMinimumPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredWithMinimumPriceCadence string

const (
	PriceTieredWithMinimumPriceCadenceOneTime    PriceTieredWithMinimumPriceCadence = "one_time"
	PriceTieredWithMinimumPriceCadenceMonthly    PriceTieredWithMinimumPriceCadence = "monthly"
	PriceTieredWithMinimumPriceCadenceQuarterly  PriceTieredWithMinimumPriceCadence = "quarterly"
	PriceTieredWithMinimumPriceCadenceSemiAnnual PriceTieredWithMinimumPriceCadence = "semi_annual"
	PriceTieredWithMinimumPriceCadenceAnnual     PriceTieredWithMinimumPriceCadence = "annual"
	PriceTieredWithMinimumPriceCadenceCustom     PriceTieredWithMinimumPriceCadence = "custom"
)

func (r PriceTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceCadenceOneTime, PriceTieredWithMinimumPriceCadenceMonthly, PriceTieredWithMinimumPriceCadenceQuarterly, PriceTieredWithMinimumPriceCadenceSemiAnnual, PriceTieredWithMinimumPriceCadenceAnnual, PriceTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredWithMinimumPriceCreditAllocation struct {
	AllowsRollover bool                                            `json:"allows_rollover,required"`
	Currency       string                                          `json:"currency,required"`
	JSON           priceTieredWithMinimumPriceCreditAllocationJSON `json:"-"`
}

// priceTieredWithMinimumPriceCreditAllocationJSON contains the JSON metadata for
// the struct [PriceTieredWithMinimumPriceCreditAllocation]
type priceTieredWithMinimumPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                              `json:"duration,required"`
	DurationUnit PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredWithMinimumPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredWithMinimumPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceTieredWithMinimumPriceInvoicingCycleConfiguration]
type priceTieredWithMinimumPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredWithMinimumPriceItem struct {
	ID   string                              `json:"id,required"`
	Name string                              `json:"name,required"`
	JSON priceTieredWithMinimumPriceItemJSON `json:"-"`
}

// priceTieredWithMinimumPriceItemJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPriceItem]
type priceTieredWithMinimumPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                 `json:"maximum_amount,required"`
	JSON          priceTieredWithMinimumPriceMaximumJSON `json:"-"`
}

// priceTieredWithMinimumPriceMaximumJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPriceMaximum]
type priceTieredWithMinimumPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                 `json:"minimum_amount,required"`
	JSON          priceTieredWithMinimumPriceMinimumJSON `json:"-"`
}

// priceTieredWithMinimumPriceMinimumJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPriceMinimum]
type priceTieredWithMinimumPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithMinimumPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithMinimumPriceModelType string

const (
	PriceTieredWithMinimumPriceModelTypeTieredWithMinimum PriceTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r PriceTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

type PriceTieredWithMinimumPricePriceType string

const (
	PriceTieredWithMinimumPricePriceTypeUsagePrice PriceTieredWithMinimumPricePriceType = "usage_price"
	PriceTieredWithMinimumPricePriceTypeFixedPrice PriceTieredWithMinimumPricePriceType = "fixed_price"
)

func (r PriceTieredWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredWithMinimumPricePriceTypeUsagePrice, PriceTieredWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPrice struct {
	ID                          string                                                        `json:"id,required"`
	BillableMetric              PriceTieredPackageWithMinimumPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredPackageWithMinimumPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredPackageWithMinimumPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                       `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                     `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredPackageWithMinimumPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                        `json:"currency,required"`
	Discount                    shared.Discount                                               `json:"discount,required,nullable"`
	ExternalPriceID             string                                                        `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                       `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredPackageWithMinimumPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredPackageWithMinimumPriceItem                        `json:"item,required"`
	Maximum                     PriceTieredPackageWithMinimumPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                        `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                       map[string]string                           `json:"metadata,required"`
	Minimum                        PriceTieredPackageWithMinimumPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount                  string                                      `json:"minimum_amount,required,nullable"`
	ModelType                      PriceTieredPackageWithMinimumPriceModelType `json:"model_type,required"`
	Name                           string                                      `json:"name,required"`
	PlanPhaseOrder                 int64                                       `json:"plan_phase_order,required,nullable"`
	PriceType                      PriceTieredPackageWithMinimumPricePriceType `json:"price_type,required"`
	TieredPackageWithMinimumConfig map[string]interface{}                      `json:"tiered_package_with_minimum_config,required"`
	JSON                           priceTieredPackageWithMinimumPriceJSON      `json:"-"`
}

// priceTieredPackageWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredPackageWithMinimumPrice]
type priceTieredPackageWithMinimumPriceJSON struct {
	ID                             apijson.Field
	BillableMetric                 apijson.Field
	BillingCycleConfiguration      apijson.Field
	Cadence                        apijson.Field
	ConversionRate                 apijson.Field
	CreatedAt                      apijson.Field
	CreditAllocation               apijson.Field
	Currency                       apijson.Field
	Discount                       apijson.Field
	ExternalPriceID                apijson.Field
	FixedPriceQuantity             apijson.Field
	InvoicingCycleConfiguration    apijson.Field
	Item                           apijson.Field
	Maximum                        apijson.Field
	MaximumAmount                  apijson.Field
	Metadata                       apijson.Field
	Minimum                        apijson.Field
	MinimumAmount                  apijson.Field
	ModelType                      apijson.Field
	Name                           apijson.Field
	PlanPhaseOrder                 apijson.Field
	PriceType                      apijson.Field
	TieredPackageWithMinimumConfig apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredPackageWithMinimumPrice) implementsPrice() {}

type PriceTieredPackageWithMinimumPriceBillableMetric struct {
	ID   string                                               `json:"id,required"`
	JSON priceTieredPackageWithMinimumPriceBillableMetricJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceBillableMetricJSON contains the JSON metadata
// for the struct [PriceTieredPackageWithMinimumPriceBillableMetric]
type priceTieredPackageWithMinimumPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	Duration     int64                                                                   `json:"duration,required"`
	DurationUnit PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPackageWithMinimumPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPackageWithMinimumPriceBillingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceTieredPackageWithMinimumPriceBillingCycleConfiguration]
type priceTieredPackageWithMinimumPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPriceCadence string

const (
	PriceTieredPackageWithMinimumPriceCadenceOneTime    PriceTieredPackageWithMinimumPriceCadence = "one_time"
	PriceTieredPackageWithMinimumPriceCadenceMonthly    PriceTieredPackageWithMinimumPriceCadence = "monthly"
	PriceTieredPackageWithMinimumPriceCadenceQuarterly  PriceTieredPackageWithMinimumPriceCadence = "quarterly"
	PriceTieredPackageWithMinimumPriceCadenceSemiAnnual PriceTieredPackageWithMinimumPriceCadence = "semi_annual"
	PriceTieredPackageWithMinimumPriceCadenceAnnual     PriceTieredPackageWithMinimumPriceCadence = "annual"
	PriceTieredPackageWithMinimumPriceCadenceCustom     PriceTieredPackageWithMinimumPriceCadence = "custom"
)

func (r PriceTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceCadenceOneTime, PriceTieredPackageWithMinimumPriceCadenceMonthly, PriceTieredPackageWithMinimumPriceCadenceQuarterly, PriceTieredPackageWithMinimumPriceCadenceSemiAnnual, PriceTieredPackageWithMinimumPriceCadenceAnnual, PriceTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPriceCreditAllocation struct {
	AllowsRollover bool                                                   `json:"allows_rollover,required"`
	Currency       string                                                 `json:"currency,required"`
	JSON           priceTieredPackageWithMinimumPriceCreditAllocationJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceCreditAllocationJSON contains the JSON
// metadata for the struct [PriceTieredPackageWithMinimumPriceCreditAllocation]
type priceTieredPackageWithMinimumPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                     `json:"duration,required"`
	DurationUnit PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredPackageWithMinimumPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredPackageWithMinimumPriceInvoicingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceTieredPackageWithMinimumPriceInvoicingCycleConfiguration]
type priceTieredPackageWithMinimumPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPriceItem struct {
	ID   string                                     `json:"id,required"`
	Name string                                     `json:"name,required"`
	JSON priceTieredPackageWithMinimumPriceItemJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceItemJSON contains the JSON metadata for the
// struct [PriceTieredPackageWithMinimumPriceItem]
type priceTieredPackageWithMinimumPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                        `json:"maximum_amount,required"`
	JSON          priceTieredPackageWithMinimumPriceMaximumJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceMaximumJSON contains the JSON metadata for the
// struct [PriceTieredPackageWithMinimumPriceMaximum]
type priceTieredPackageWithMinimumPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                        `json:"minimum_amount,required"`
	JSON          priceTieredPackageWithMinimumPriceMinimumJSON `json:"-"`
}

// priceTieredPackageWithMinimumPriceMinimumJSON contains the JSON metadata for the
// struct [PriceTieredPackageWithMinimumPriceMinimum]
type priceTieredPackageWithMinimumPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredPackageWithMinimumPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredPackageWithMinimumPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredPackageWithMinimumPriceModelType string

const (
	PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PriceTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r PriceTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

type PriceTieredPackageWithMinimumPricePriceType string

const (
	PriceTieredPackageWithMinimumPricePriceTypeUsagePrice PriceTieredPackageWithMinimumPricePriceType = "usage_price"
	PriceTieredPackageWithMinimumPricePriceTypeFixedPrice PriceTieredPackageWithMinimumPricePriceType = "fixed_price"
)

func (r PriceTieredPackageWithMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredPackageWithMinimumPricePriceTypeUsagePrice, PriceTieredPackageWithMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PricePackageWithAllocationPrice struct {
	ID                          string                                                     `json:"id,required"`
	BillableMetric              PricePackageWithAllocationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PricePackageWithAllocationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PricePackageWithAllocationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                    `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation            PricePackageWithAllocationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                     `json:"currency,required"`
	Discount                    shared.Discount                                            `json:"discount,required,nullable"`
	ExternalPriceID             string                                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                    `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PricePackageWithAllocationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PricePackageWithAllocationPriceItem                        `json:"item,required"`
	Maximum                     PricePackageWithAllocationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                     `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                    map[string]string                        `json:"metadata,required"`
	Minimum                     PricePackageWithAllocationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount               string                                   `json:"minimum_amount,required,nullable"`
	ModelType                   PricePackageWithAllocationPriceModelType `json:"model_type,required"`
	Name                        string                                   `json:"name,required"`
	PackageWithAllocationConfig map[string]interface{}                   `json:"package_with_allocation_config,required"`
	PlanPhaseOrder              int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType                   PricePackageWithAllocationPricePriceType `json:"price_type,required"`
	JSON                        pricePackageWithAllocationPriceJSON      `json:"-"`
}

// pricePackageWithAllocationPriceJSON contains the JSON metadata for the struct
// [PricePackageWithAllocationPrice]
type pricePackageWithAllocationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PackageWithAllocationConfig apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PricePackageWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PricePackageWithAllocationPrice) implementsPrice() {}

type PricePackageWithAllocationPriceBillableMetric struct {
	ID   string                                            `json:"id,required"`
	JSON pricePackageWithAllocationPriceBillableMetricJSON `json:"-"`
}

// pricePackageWithAllocationPriceBillableMetricJSON contains the JSON metadata for
// the struct [PricePackageWithAllocationPriceBillableMetric]
type pricePackageWithAllocationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceBillingCycleConfiguration struct {
	Duration     int64                                                                `json:"duration,required"`
	DurationUnit PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         pricePackageWithAllocationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// pricePackageWithAllocationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PricePackageWithAllocationPriceBillingCycleConfiguration]
type pricePackageWithAllocationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PricePackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PricePackageWithAllocationPriceCadence string

const (
	PricePackageWithAllocationPriceCadenceOneTime    PricePackageWithAllocationPriceCadence = "one_time"
	PricePackageWithAllocationPriceCadenceMonthly    PricePackageWithAllocationPriceCadence = "monthly"
	PricePackageWithAllocationPriceCadenceQuarterly  PricePackageWithAllocationPriceCadence = "quarterly"
	PricePackageWithAllocationPriceCadenceSemiAnnual PricePackageWithAllocationPriceCadence = "semi_annual"
	PricePackageWithAllocationPriceCadenceAnnual     PricePackageWithAllocationPriceCadence = "annual"
	PricePackageWithAllocationPriceCadenceCustom     PricePackageWithAllocationPriceCadence = "custom"
)

func (r PricePackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceCadenceOneTime, PricePackageWithAllocationPriceCadenceMonthly, PricePackageWithAllocationPriceCadenceQuarterly, PricePackageWithAllocationPriceCadenceSemiAnnual, PricePackageWithAllocationPriceCadenceAnnual, PricePackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PricePackageWithAllocationPriceCreditAllocation struct {
	AllowsRollover bool                                                `json:"allows_rollover,required"`
	Currency       string                                              `json:"currency,required"`
	JSON           pricePackageWithAllocationPriceCreditAllocationJSON `json:"-"`
}

// pricePackageWithAllocationPriceCreditAllocationJSON contains the JSON metadata
// for the struct [PricePackageWithAllocationPriceCreditAllocation]
type pricePackageWithAllocationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                  `json:"duration,required"`
	DurationUnit PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         pricePackageWithAllocationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// pricePackageWithAllocationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PricePackageWithAllocationPriceInvoicingCycleConfiguration]
type pricePackageWithAllocationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PricePackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PricePackageWithAllocationPriceItem struct {
	ID   string                                  `json:"id,required"`
	Name string                                  `json:"name,required"`
	JSON pricePackageWithAllocationPriceItemJSON `json:"-"`
}

// pricePackageWithAllocationPriceItemJSON contains the JSON metadata for the
// struct [PricePackageWithAllocationPriceItem]
type pricePackageWithAllocationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                     `json:"maximum_amount,required"`
	JSON          pricePackageWithAllocationPriceMaximumJSON `json:"-"`
}

// pricePackageWithAllocationPriceMaximumJSON contains the JSON metadata for the
// struct [PricePackageWithAllocationPriceMaximum]
type pricePackageWithAllocationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                     `json:"minimum_amount,required"`
	JSON          pricePackageWithAllocationPriceMinimumJSON `json:"-"`
}

// pricePackageWithAllocationPriceMinimumJSON contains the JSON metadata for the
// struct [PricePackageWithAllocationPriceMinimum]
type pricePackageWithAllocationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PricePackageWithAllocationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricePackageWithAllocationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PricePackageWithAllocationPriceModelType string

const (
	PricePackageWithAllocationPriceModelTypePackageWithAllocation PricePackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r PricePackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

type PricePackageWithAllocationPricePriceType string

const (
	PricePackageWithAllocationPricePriceTypeUsagePrice PricePackageWithAllocationPricePriceType = "usage_price"
	PricePackageWithAllocationPricePriceTypeFixedPrice PricePackageWithAllocationPricePriceType = "fixed_price"
)

func (r PricePackageWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PricePackageWithAllocationPricePriceTypeUsagePrice, PricePackageWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceUnitWithPercentPrice struct {
	ID                          string                                               `json:"id,required"`
	BillableMetric              PriceUnitWithPercentPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceUnitWithPercentPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceUnitWithPercentPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                              `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                            `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceUnitWithPercentPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                               `json:"currency,required"`
	Discount                    shared.Discount                                      `json:"discount,required,nullable"`
	ExternalPriceID             string                                               `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                              `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceUnitWithPercentPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceUnitWithPercentPriceItem                        `json:"item,required"`
	Maximum                     PriceUnitWithPercentPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                               `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata              map[string]string                  `json:"metadata,required"`
	Minimum               PriceUnitWithPercentPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount         string                             `json:"minimum_amount,required,nullable"`
	ModelType             PriceUnitWithPercentPriceModelType `json:"model_type,required"`
	Name                  string                             `json:"name,required"`
	PlanPhaseOrder        int64                              `json:"plan_phase_order,required,nullable"`
	PriceType             PriceUnitWithPercentPricePriceType `json:"price_type,required"`
	UnitWithPercentConfig map[string]interface{}             `json:"unit_with_percent_config,required"`
	JSON                  priceUnitWithPercentPriceJSON      `json:"-"`
}

// priceUnitWithPercentPriceJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPrice]
type priceUnitWithPercentPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	UnitWithPercentConfig       apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceUnitWithPercentPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitWithPercentPrice) implementsPrice() {}

type PriceUnitWithPercentPriceBillableMetric struct {
	ID   string                                      `json:"id,required"`
	JSON priceUnitWithPercentPriceBillableMetricJSON `json:"-"`
}

// priceUnitWithPercentPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceUnitWithPercentPriceBillableMetric]
type priceUnitWithPercentPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceBillingCycleConfiguration struct {
	Duration     int64                                                          `json:"duration,required"`
	DurationUnit PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitWithPercentPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceUnitWithPercentPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceUnitWithPercentPriceBillingCycleConfiguration]
type priceUnitWithPercentPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PriceUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitWithPercentPriceCadence string

const (
	PriceUnitWithPercentPriceCadenceOneTime    PriceUnitWithPercentPriceCadence = "one_time"
	PriceUnitWithPercentPriceCadenceMonthly    PriceUnitWithPercentPriceCadence = "monthly"
	PriceUnitWithPercentPriceCadenceQuarterly  PriceUnitWithPercentPriceCadence = "quarterly"
	PriceUnitWithPercentPriceCadenceSemiAnnual PriceUnitWithPercentPriceCadence = "semi_annual"
	PriceUnitWithPercentPriceCadenceAnnual     PriceUnitWithPercentPriceCadence = "annual"
	PriceUnitWithPercentPriceCadenceCustom     PriceUnitWithPercentPriceCadence = "custom"
)

func (r PriceUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceCadenceOneTime, PriceUnitWithPercentPriceCadenceMonthly, PriceUnitWithPercentPriceCadenceQuarterly, PriceUnitWithPercentPriceCadenceSemiAnnual, PriceUnitWithPercentPriceCadenceAnnual, PriceUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitWithPercentPriceCreditAllocation struct {
	AllowsRollover bool                                          `json:"allows_rollover,required"`
	Currency       string                                        `json:"currency,required"`
	JSON           priceUnitWithPercentPriceCreditAllocationJSON `json:"-"`
}

// priceUnitWithPercentPriceCreditAllocationJSON contains the JSON metadata for the
// struct [PriceUnitWithPercentPriceCreditAllocation]
type priceUnitWithPercentPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                            `json:"duration,required"`
	DurationUnit PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitWithPercentPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceUnitWithPercentPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceUnitWithPercentPriceInvoicingCycleConfiguration]
type priceUnitWithPercentPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PriceUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitWithPercentPriceItem struct {
	ID   string                            `json:"id,required"`
	Name string                            `json:"name,required"`
	JSON priceUnitWithPercentPriceItemJSON `json:"-"`
}

// priceUnitWithPercentPriceItemJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPriceItem]
type priceUnitWithPercentPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                               `json:"maximum_amount,required"`
	JSON          priceUnitWithPercentPriceMaximumJSON `json:"-"`
}

// priceUnitWithPercentPriceMaximumJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPriceMaximum]
type priceUnitWithPercentPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                               `json:"minimum_amount,required"`
	JSON          priceUnitWithPercentPriceMinimumJSON `json:"-"`
}

// priceUnitWithPercentPriceMinimumJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPriceMinimum]
type priceUnitWithPercentPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitWithPercentPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithPercentPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithPercentPriceModelType string

const (
	PriceUnitWithPercentPriceModelTypeUnitWithPercent PriceUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r PriceUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

type PriceUnitWithPercentPricePriceType string

const (
	PriceUnitWithPercentPricePriceTypeUsagePrice PriceUnitWithPercentPricePriceType = "usage_price"
	PriceUnitWithPercentPricePriceTypeFixedPrice PriceUnitWithPercentPricePriceType = "fixed_price"
)

func (r PriceUnitWithPercentPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitWithPercentPricePriceTypeUsagePrice, PriceUnitWithPercentPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPrice struct {
	ID                          string                                                    `json:"id,required"`
	BillableMetric              PriceMatrixWithAllocationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceMatrixWithAllocationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceMatrixWithAllocationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceMatrixWithAllocationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                    `json:"currency,required"`
	Discount                    shared.Discount                                           `json:"discount,required,nullable"`
	ExternalPriceID             string                                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                   `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceMatrixWithAllocationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceMatrixWithAllocationPriceItem                        `json:"item,required"`
	MatrixWithAllocationConfig  PriceMatrixWithAllocationPriceMatrixWithAllocationConfig  `json:"matrix_with_allocation_config,required"`
	Maximum                     PriceMatrixWithAllocationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                       `json:"metadata,required"`
	Minimum        PriceMatrixWithAllocationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixWithAllocationPriceModelType `json:"model_type,required"`
	Name           string                                  `json:"name,required"`
	PlanPhaseOrder int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixWithAllocationPricePriceType `json:"price_type,required"`
	JSON           priceMatrixWithAllocationPriceJSON      `json:"-"`
}

// priceMatrixWithAllocationPriceJSON contains the JSON metadata for the struct
// [PriceMatrixWithAllocationPrice]
type priceMatrixWithAllocationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	MatrixWithAllocationConfig  apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixWithAllocationPrice) implementsPrice() {}

type PriceMatrixWithAllocationPriceBillableMetric struct {
	ID   string                                           `json:"id,required"`
	JSON priceMatrixWithAllocationPriceBillableMetricJSON `json:"-"`
}

// priceMatrixWithAllocationPriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceMatrixWithAllocationPriceBillableMetric]
type priceMatrixWithAllocationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceBillingCycleConfiguration struct {
	Duration     int64                                                               `json:"duration,required"`
	DurationUnit PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixWithAllocationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixWithAllocationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithAllocationPriceBillingCycleConfiguration]
type priceMatrixWithAllocationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PriceMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPriceCadence string

const (
	PriceMatrixWithAllocationPriceCadenceOneTime    PriceMatrixWithAllocationPriceCadence = "one_time"
	PriceMatrixWithAllocationPriceCadenceMonthly    PriceMatrixWithAllocationPriceCadence = "monthly"
	PriceMatrixWithAllocationPriceCadenceQuarterly  PriceMatrixWithAllocationPriceCadence = "quarterly"
	PriceMatrixWithAllocationPriceCadenceSemiAnnual PriceMatrixWithAllocationPriceCadence = "semi_annual"
	PriceMatrixWithAllocationPriceCadenceAnnual     PriceMatrixWithAllocationPriceCadence = "annual"
	PriceMatrixWithAllocationPriceCadenceCustom     PriceMatrixWithAllocationPriceCadence = "custom"
)

func (r PriceMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceCadenceOneTime, PriceMatrixWithAllocationPriceCadenceMonthly, PriceMatrixWithAllocationPriceCadenceQuarterly, PriceMatrixWithAllocationPriceCadenceSemiAnnual, PriceMatrixWithAllocationPriceCadenceAnnual, PriceMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPriceCreditAllocation struct {
	AllowsRollover bool                                               `json:"allows_rollover,required"`
	Currency       string                                             `json:"currency,required"`
	JSON           priceMatrixWithAllocationPriceCreditAllocationJSON `json:"-"`
}

// priceMatrixWithAllocationPriceCreditAllocationJSON contains the JSON metadata
// for the struct [PriceMatrixWithAllocationPriceCreditAllocation]
type priceMatrixWithAllocationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                 `json:"duration,required"`
	DurationUnit PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixWithAllocationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixWithAllocationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithAllocationPriceInvoicingCycleConfiguration]
type priceMatrixWithAllocationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PriceMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPriceItem struct {
	ID   string                                 `json:"id,required"`
	Name string                                 `json:"name,required"`
	JSON priceMatrixWithAllocationPriceItemJSON `json:"-"`
}

// priceMatrixWithAllocationPriceItemJSON contains the JSON metadata for the struct
// [PriceMatrixWithAllocationPriceItem]
type priceMatrixWithAllocationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation float64 `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []PriceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue `json:"matrix_values,required"`
	JSON         priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON          `json:"-"`
}

// priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithAllocationPriceMatrixWithAllocationConfig]
type priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON struct {
	Allocation        apijson.Field
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceMatrixWithAllocationConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string                                                                  `json:"unit_amount,required"`
	JSON       priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON `json:"-"`
}

// priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON contains
// the JSON metadata for the struct
// [PriceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue]
type priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                    `json:"maximum_amount,required"`
	JSON          priceMatrixWithAllocationPriceMaximumJSON `json:"-"`
}

// priceMatrixWithAllocationPriceMaximumJSON contains the JSON metadata for the
// struct [PriceMatrixWithAllocationPriceMaximum]
type priceMatrixWithAllocationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                    `json:"minimum_amount,required"`
	JSON          priceMatrixWithAllocationPriceMinimumJSON `json:"-"`
}

// priceMatrixWithAllocationPriceMinimumJSON contains the JSON metadata for the
// struct [PriceMatrixWithAllocationPriceMinimum]
type priceMatrixWithAllocationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixWithAllocationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithAllocationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithAllocationPriceModelType string

const (
	PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation PriceMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r PriceMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type PriceMatrixWithAllocationPricePriceType string

const (
	PriceMatrixWithAllocationPricePriceTypeUsagePrice PriceMatrixWithAllocationPricePriceType = "usage_price"
	PriceMatrixWithAllocationPricePriceTypeFixedPrice PriceMatrixWithAllocationPricePriceType = "fixed_price"
)

func (r PriceMatrixWithAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixWithAllocationPricePriceTypeUsagePrice, PriceMatrixWithAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceTieredWithProrationPrice struct {
	ID                          string                                                   `json:"id,required"`
	BillableMetric              PriceTieredWithProrationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceTieredWithProrationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceTieredWithProrationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                  `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceTieredWithProrationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                   `json:"currency,required"`
	Discount                    shared.Discount                                          `json:"discount,required,nullable"`
	ExternalPriceID             string                                                   `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                  `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceTieredWithProrationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceTieredWithProrationPriceItem                        `json:"item,required"`
	Maximum                     PriceTieredWithProrationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                   `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                  map[string]string                      `json:"metadata,required"`
	Minimum                   PriceTieredWithProrationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount             string                                 `json:"minimum_amount,required,nullable"`
	ModelType                 PriceTieredWithProrationPriceModelType `json:"model_type,required"`
	Name                      string                                 `json:"name,required"`
	PlanPhaseOrder            int64                                  `json:"plan_phase_order,required,nullable"`
	PriceType                 PriceTieredWithProrationPricePriceType `json:"price_type,required"`
	TieredWithProrationConfig map[string]interface{}                 `json:"tiered_with_proration_config,required"`
	JSON                      priceTieredWithProrationPriceJSON      `json:"-"`
}

// priceTieredWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithProrationPrice]
type priceTieredWithProrationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	TieredWithProrationConfig   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceTieredWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceTieredWithProrationPrice) implementsPrice() {}

type PriceTieredWithProrationPriceBillableMetric struct {
	ID   string                                          `json:"id,required"`
	JSON priceTieredWithProrationPriceBillableMetricJSON `json:"-"`
}

// priceTieredWithProrationPriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceTieredWithProrationPriceBillableMetric]
type priceTieredWithProrationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceBillingCycleConfiguration struct {
	Duration     int64                                                              `json:"duration,required"`
	DurationUnit PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredWithProrationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceTieredWithProrationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceTieredWithProrationPriceBillingCycleConfiguration]
type priceTieredWithProrationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredWithProrationPriceCadence string

const (
	PriceTieredWithProrationPriceCadenceOneTime    PriceTieredWithProrationPriceCadence = "one_time"
	PriceTieredWithProrationPriceCadenceMonthly    PriceTieredWithProrationPriceCadence = "monthly"
	PriceTieredWithProrationPriceCadenceQuarterly  PriceTieredWithProrationPriceCadence = "quarterly"
	PriceTieredWithProrationPriceCadenceSemiAnnual PriceTieredWithProrationPriceCadence = "semi_annual"
	PriceTieredWithProrationPriceCadenceAnnual     PriceTieredWithProrationPriceCadence = "annual"
	PriceTieredWithProrationPriceCadenceCustom     PriceTieredWithProrationPriceCadence = "custom"
)

func (r PriceTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceCadenceOneTime, PriceTieredWithProrationPriceCadenceMonthly, PriceTieredWithProrationPriceCadenceQuarterly, PriceTieredWithProrationPriceCadenceSemiAnnual, PriceTieredWithProrationPriceCadenceAnnual, PriceTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceTieredWithProrationPriceCreditAllocation struct {
	AllowsRollover bool                                              `json:"allows_rollover,required"`
	Currency       string                                            `json:"currency,required"`
	JSON           priceTieredWithProrationPriceCreditAllocationJSON `json:"-"`
}

// priceTieredWithProrationPriceCreditAllocationJSON contains the JSON metadata for
// the struct [PriceTieredWithProrationPriceCreditAllocation]
type priceTieredWithProrationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                `json:"duration,required"`
	DurationUnit PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceTieredWithProrationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceTieredWithProrationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceTieredWithProrationPriceInvoicingCycleConfiguration]
type priceTieredWithProrationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceTieredWithProrationPriceItem struct {
	ID   string                                `json:"id,required"`
	Name string                                `json:"name,required"`
	JSON priceTieredWithProrationPriceItemJSON `json:"-"`
}

// priceTieredWithProrationPriceItemJSON contains the JSON metadata for the struct
// [PriceTieredWithProrationPriceItem]
type priceTieredWithProrationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                   `json:"maximum_amount,required"`
	JSON          priceTieredWithProrationPriceMaximumJSON `json:"-"`
}

// priceTieredWithProrationPriceMaximumJSON contains the JSON metadata for the
// struct [PriceTieredWithProrationPriceMaximum]
type priceTieredWithProrationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                   `json:"minimum_amount,required"`
	JSON          priceTieredWithProrationPriceMinimumJSON `json:"-"`
}

// priceTieredWithProrationPriceMinimumJSON contains the JSON metadata for the
// struct [PriceTieredWithProrationPriceMinimum]
type priceTieredWithProrationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTieredWithProrationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceTieredWithProrationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceTieredWithProrationPriceModelType string

const (
	PriceTieredWithProrationPriceModelTypeTieredWithProration PriceTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r PriceTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

type PriceTieredWithProrationPricePriceType string

const (
	PriceTieredWithProrationPricePriceTypeUsagePrice PriceTieredWithProrationPricePriceType = "usage_price"
	PriceTieredWithProrationPricePriceTypeFixedPrice PriceTieredWithProrationPricePriceType = "fixed_price"
)

func (r PriceTieredWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceTieredWithProrationPricePriceTypeUsagePrice, PriceTieredWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceUnitWithProrationPrice struct {
	ID                          string                                                 `json:"id,required"`
	BillableMetric              PriceUnitWithProrationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceUnitWithProrationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceUnitWithProrationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                              `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceUnitWithProrationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                 `json:"currency,required"`
	Discount                    shared.Discount                                        `json:"discount,required,nullable"`
	ExternalPriceID             string                                                 `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceUnitWithProrationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceUnitWithProrationPriceItem                        `json:"item,required"`
	Maximum                     PriceUnitWithProrationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                 `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata                map[string]string                    `json:"metadata,required"`
	Minimum                 PriceUnitWithProrationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount           string                               `json:"minimum_amount,required,nullable"`
	ModelType               PriceUnitWithProrationPriceModelType `json:"model_type,required"`
	Name                    string                               `json:"name,required"`
	PlanPhaseOrder          int64                                `json:"plan_phase_order,required,nullable"`
	PriceType               PriceUnitWithProrationPricePriceType `json:"price_type,required"`
	UnitWithProrationConfig map[string]interface{}               `json:"unit_with_proration_config,required"`
	JSON                    priceUnitWithProrationPriceJSON      `json:"-"`
}

// priceUnitWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceUnitWithProrationPrice]
type priceUnitWithProrationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	UnitWithProrationConfig     apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceUnitWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceUnitWithProrationPrice) implementsPrice() {}

type PriceUnitWithProrationPriceBillableMetric struct {
	ID   string                                        `json:"id,required"`
	JSON priceUnitWithProrationPriceBillableMetricJSON `json:"-"`
}

// priceUnitWithProrationPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceUnitWithProrationPriceBillableMetric]
type priceUnitWithProrationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceBillingCycleConfiguration struct {
	Duration     int64                                                            `json:"duration,required"`
	DurationUnit PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitWithProrationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceUnitWithProrationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceUnitWithProrationPriceBillingCycleConfiguration]
type priceUnitWithProrationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitWithProrationPriceCadence string

const (
	PriceUnitWithProrationPriceCadenceOneTime    PriceUnitWithProrationPriceCadence = "one_time"
	PriceUnitWithProrationPriceCadenceMonthly    PriceUnitWithProrationPriceCadence = "monthly"
	PriceUnitWithProrationPriceCadenceQuarterly  PriceUnitWithProrationPriceCadence = "quarterly"
	PriceUnitWithProrationPriceCadenceSemiAnnual PriceUnitWithProrationPriceCadence = "semi_annual"
	PriceUnitWithProrationPriceCadenceAnnual     PriceUnitWithProrationPriceCadence = "annual"
	PriceUnitWithProrationPriceCadenceCustom     PriceUnitWithProrationPriceCadence = "custom"
)

func (r PriceUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceCadenceOneTime, PriceUnitWithProrationPriceCadenceMonthly, PriceUnitWithProrationPriceCadenceQuarterly, PriceUnitWithProrationPriceCadenceSemiAnnual, PriceUnitWithProrationPriceCadenceAnnual, PriceUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceUnitWithProrationPriceCreditAllocation struct {
	AllowsRollover bool                                            `json:"allows_rollover,required"`
	Currency       string                                          `json:"currency,required"`
	JSON           priceUnitWithProrationPriceCreditAllocationJSON `json:"-"`
}

// priceUnitWithProrationPriceCreditAllocationJSON contains the JSON metadata for
// the struct [PriceUnitWithProrationPriceCreditAllocation]
type priceUnitWithProrationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                              `json:"duration,required"`
	DurationUnit PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceUnitWithProrationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceUnitWithProrationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceUnitWithProrationPriceInvoicingCycleConfiguration]
type priceUnitWithProrationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceUnitWithProrationPriceItem struct {
	ID   string                              `json:"id,required"`
	Name string                              `json:"name,required"`
	JSON priceUnitWithProrationPriceItemJSON `json:"-"`
}

// priceUnitWithProrationPriceItemJSON contains the JSON metadata for the struct
// [PriceUnitWithProrationPriceItem]
type priceUnitWithProrationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                 `json:"maximum_amount,required"`
	JSON          priceUnitWithProrationPriceMaximumJSON `json:"-"`
}

// priceUnitWithProrationPriceMaximumJSON contains the JSON metadata for the struct
// [PriceUnitWithProrationPriceMaximum]
type priceUnitWithProrationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                 `json:"minimum_amount,required"`
	JSON          priceUnitWithProrationPriceMinimumJSON `json:"-"`
}

// priceUnitWithProrationPriceMinimumJSON contains the JSON metadata for the struct
// [PriceUnitWithProrationPriceMinimum]
type priceUnitWithProrationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceUnitWithProrationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceUnitWithProrationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceUnitWithProrationPriceModelType string

const (
	PriceUnitWithProrationPriceModelTypeUnitWithProration PriceUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r PriceUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

type PriceUnitWithProrationPricePriceType string

const (
	PriceUnitWithProrationPricePriceTypeUsagePrice PriceUnitWithProrationPricePriceType = "usage_price"
	PriceUnitWithProrationPricePriceTypeFixedPrice PriceUnitWithProrationPricePriceType = "fixed_price"
)

func (r PriceUnitWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceUnitWithProrationPricePriceTypeUsagePrice, PriceUnitWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedAllocationPrice struct {
	ID                          string                                                 `json:"id,required"`
	BillableMetric              PriceGroupedAllocationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceGroupedAllocationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceGroupedAllocationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                              `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceGroupedAllocationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                 `json:"currency,required"`
	Discount                    shared.Discount                                        `json:"discount,required,nullable"`
	ExternalPriceID             string                                                 `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                `json:"fixed_price_quantity,required,nullable"`
	GroupedAllocationConfig     map[string]interface{}                                 `json:"grouped_allocation_config,required"`
	InvoicingCycleConfiguration PriceGroupedAllocationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceGroupedAllocationPriceItem                        `json:"item,required"`
	Maximum                     PriceGroupedAllocationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                 `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                    `json:"metadata,required"`
	Minimum        PriceGroupedAllocationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedAllocationPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedAllocationPricePriceType `json:"price_type,required"`
	JSON           priceGroupedAllocationPriceJSON      `json:"-"`
}

// priceGroupedAllocationPriceJSON contains the JSON metadata for the struct
// [PriceGroupedAllocationPrice]
type priceGroupedAllocationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	GroupedAllocationConfig     apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceGroupedAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedAllocationPrice) implementsPrice() {}

type PriceGroupedAllocationPriceBillableMetric struct {
	ID   string                                        `json:"id,required"`
	JSON priceGroupedAllocationPriceBillableMetricJSON `json:"-"`
}

// priceGroupedAllocationPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceGroupedAllocationPriceBillableMetric]
type priceGroupedAllocationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceBillingCycleConfiguration struct {
	Duration     int64                                                            `json:"duration,required"`
	DurationUnit PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedAllocationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedAllocationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceGroupedAllocationPriceBillingCycleConfiguration]
type priceGroupedAllocationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PriceGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedAllocationPriceCadence string

const (
	PriceGroupedAllocationPriceCadenceOneTime    PriceGroupedAllocationPriceCadence = "one_time"
	PriceGroupedAllocationPriceCadenceMonthly    PriceGroupedAllocationPriceCadence = "monthly"
	PriceGroupedAllocationPriceCadenceQuarterly  PriceGroupedAllocationPriceCadence = "quarterly"
	PriceGroupedAllocationPriceCadenceSemiAnnual PriceGroupedAllocationPriceCadence = "semi_annual"
	PriceGroupedAllocationPriceCadenceAnnual     PriceGroupedAllocationPriceCadence = "annual"
	PriceGroupedAllocationPriceCadenceCustom     PriceGroupedAllocationPriceCadence = "custom"
)

func (r PriceGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceCadenceOneTime, PriceGroupedAllocationPriceCadenceMonthly, PriceGroupedAllocationPriceCadenceQuarterly, PriceGroupedAllocationPriceCadenceSemiAnnual, PriceGroupedAllocationPriceCadenceAnnual, PriceGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedAllocationPriceCreditAllocation struct {
	AllowsRollover bool                                            `json:"allows_rollover,required"`
	Currency       string                                          `json:"currency,required"`
	JSON           priceGroupedAllocationPriceCreditAllocationJSON `json:"-"`
}

// priceGroupedAllocationPriceCreditAllocationJSON contains the JSON metadata for
// the struct [PriceGroupedAllocationPriceCreditAllocation]
type priceGroupedAllocationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                              `json:"duration,required"`
	DurationUnit PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedAllocationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedAllocationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceGroupedAllocationPriceInvoicingCycleConfiguration]
type priceGroupedAllocationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PriceGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedAllocationPriceItem struct {
	ID   string                              `json:"id,required"`
	Name string                              `json:"name,required"`
	JSON priceGroupedAllocationPriceItemJSON `json:"-"`
}

// priceGroupedAllocationPriceItemJSON contains the JSON metadata for the struct
// [PriceGroupedAllocationPriceItem]
type priceGroupedAllocationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                 `json:"maximum_amount,required"`
	JSON          priceGroupedAllocationPriceMaximumJSON `json:"-"`
}

// priceGroupedAllocationPriceMaximumJSON contains the JSON metadata for the struct
// [PriceGroupedAllocationPriceMaximum]
type priceGroupedAllocationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                 `json:"minimum_amount,required"`
	JSON          priceGroupedAllocationPriceMinimumJSON `json:"-"`
}

// priceGroupedAllocationPriceMinimumJSON contains the JSON metadata for the struct
// [PriceGroupedAllocationPriceMinimum]
type priceGroupedAllocationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedAllocationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedAllocationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedAllocationPriceModelType string

const (
	PriceGroupedAllocationPriceModelTypeGroupedAllocation PriceGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r PriceGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type PriceGroupedAllocationPricePriceType string

const (
	PriceGroupedAllocationPricePriceTypeUsagePrice PriceGroupedAllocationPricePriceType = "usage_price"
	PriceGroupedAllocationPricePriceTypeFixedPrice PriceGroupedAllocationPricePriceType = "fixed_price"
)

func (r PriceGroupedAllocationPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedAllocationPricePriceTypeUsagePrice, PriceGroupedAllocationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPrice struct {
	ID                               string                                                          `json:"id,required"`
	BillableMetric                   PriceGroupedWithProratedMinimumPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration        PriceGroupedWithProratedMinimumPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                          PriceGroupedWithProratedMinimumPriceCadence                     `json:"cadence,required"`
	ConversionRate                   float64                                                         `json:"conversion_rate,required,nullable"`
	CreatedAt                        time.Time                                                       `json:"created_at,required" format:"date-time"`
	CreditAllocation                 PriceGroupedWithProratedMinimumPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                         string                                                          `json:"currency,required"`
	Discount                         shared.Discount                                                 `json:"discount,required,nullable"`
	ExternalPriceID                  string                                                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity               float64                                                         `json:"fixed_price_quantity,required,nullable"`
	GroupedWithProratedMinimumConfig map[string]interface{}                                          `json:"grouped_with_prorated_minimum_config,required"`
	InvoicingCycleConfiguration      PriceGroupedWithProratedMinimumPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                             PriceGroupedWithProratedMinimumPriceItem                        `json:"item,required"`
	Maximum                          PriceGroupedWithProratedMinimumPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount                    string                                                          `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                             `json:"metadata,required"`
	Minimum        PriceGroupedWithProratedMinimumPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                                        `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedWithProratedMinimumPriceModelType `json:"model_type,required"`
	Name           string                                        `json:"name,required"`
	PlanPhaseOrder int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedWithProratedMinimumPricePriceType `json:"price_type,required"`
	JSON           priceGroupedWithProratedMinimumPriceJSON      `json:"-"`
}

// priceGroupedWithProratedMinimumPriceJSON contains the JSON metadata for the
// struct [PriceGroupedWithProratedMinimumPrice]
type priceGroupedWithProratedMinimumPriceJSON struct {
	ID                               apijson.Field
	BillableMetric                   apijson.Field
	BillingCycleConfiguration        apijson.Field
	Cadence                          apijson.Field
	ConversionRate                   apijson.Field
	CreatedAt                        apijson.Field
	CreditAllocation                 apijson.Field
	Currency                         apijson.Field
	Discount                         apijson.Field
	ExternalPriceID                  apijson.Field
	FixedPriceQuantity               apijson.Field
	GroupedWithProratedMinimumConfig apijson.Field
	InvoicingCycleConfiguration      apijson.Field
	Item                             apijson.Field
	Maximum                          apijson.Field
	MaximumAmount                    apijson.Field
	Metadata                         apijson.Field
	Minimum                          apijson.Field
	MinimumAmount                    apijson.Field
	ModelType                        apijson.Field
	Name                             apijson.Field
	PlanPhaseOrder                   apijson.Field
	PriceType                        apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedWithProratedMinimumPrice) implementsPrice() {}

type PriceGroupedWithProratedMinimumPriceBillableMetric struct {
	ID   string                                                 `json:"id,required"`
	JSON priceGroupedWithProratedMinimumPriceBillableMetricJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceBillableMetricJSON contains the JSON
// metadata for the struct [PriceGroupedWithProratedMinimumPriceBillableMetric]
type priceGroupedWithProratedMinimumPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	Duration     int64                                                                     `json:"duration,required"`
	DurationUnit PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedWithProratedMinimumPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedWithProratedMinimumPriceBillingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceGroupedWithProratedMinimumPriceBillingCycleConfiguration]
type priceGroupedWithProratedMinimumPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPriceCadence string

const (
	PriceGroupedWithProratedMinimumPriceCadenceOneTime    PriceGroupedWithProratedMinimumPriceCadence = "one_time"
	PriceGroupedWithProratedMinimumPriceCadenceMonthly    PriceGroupedWithProratedMinimumPriceCadence = "monthly"
	PriceGroupedWithProratedMinimumPriceCadenceQuarterly  PriceGroupedWithProratedMinimumPriceCadence = "quarterly"
	PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual PriceGroupedWithProratedMinimumPriceCadence = "semi_annual"
	PriceGroupedWithProratedMinimumPriceCadenceAnnual     PriceGroupedWithProratedMinimumPriceCadence = "annual"
	PriceGroupedWithProratedMinimumPriceCadenceCustom     PriceGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r PriceGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceCadenceOneTime, PriceGroupedWithProratedMinimumPriceCadenceMonthly, PriceGroupedWithProratedMinimumPriceCadenceQuarterly, PriceGroupedWithProratedMinimumPriceCadenceSemiAnnual, PriceGroupedWithProratedMinimumPriceCadenceAnnual, PriceGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPriceCreditAllocation struct {
	AllowsRollover bool                                                     `json:"allows_rollover,required"`
	Currency       string                                                   `json:"currency,required"`
	JSON           priceGroupedWithProratedMinimumPriceCreditAllocationJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceCreditAllocationJSON contains the JSON
// metadata for the struct [PriceGroupedWithProratedMinimumPriceCreditAllocation]
type priceGroupedWithProratedMinimumPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                       `json:"duration,required"`
	DurationUnit PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceGroupedWithProratedMinimumPriceInvoicingCycleConfiguration]
type priceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPriceItem struct {
	ID   string                                       `json:"id,required"`
	Name string                                       `json:"name,required"`
	JSON priceGroupedWithProratedMinimumPriceItemJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceItemJSON contains the JSON metadata for the
// struct [PriceGroupedWithProratedMinimumPriceItem]
type priceGroupedWithProratedMinimumPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                          `json:"maximum_amount,required"`
	JSON          priceGroupedWithProratedMinimumPriceMaximumJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceMaximumJSON contains the JSON metadata for
// the struct [PriceGroupedWithProratedMinimumPriceMaximum]
type priceGroupedWithProratedMinimumPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                          `json:"minimum_amount,required"`
	JSON          priceGroupedWithProratedMinimumPriceMinimumJSON `json:"-"`
}

// priceGroupedWithProratedMinimumPriceMinimumJSON contains the JSON metadata for
// the struct [PriceGroupedWithProratedMinimumPriceMinimum]
type priceGroupedWithProratedMinimumPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedWithProratedMinimumPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithProratedMinimumPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithProratedMinimumPriceModelType string

const (
	PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum PriceGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r PriceGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type PriceGroupedWithProratedMinimumPricePriceType string

const (
	PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice PriceGroupedWithProratedMinimumPricePriceType = "usage_price"
	PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice PriceGroupedWithProratedMinimumPricePriceType = "fixed_price"
)

func (r PriceGroupedWithProratedMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedWithProratedMinimumPricePriceTypeUsagePrice, PriceGroupedWithProratedMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPrice struct {
	ID                              string                                                         `json:"id,required"`
	BillableMetric                  PriceGroupedWithMeteredMinimumPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration       PriceGroupedWithMeteredMinimumPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                         PriceGroupedWithMeteredMinimumPriceCadence                     `json:"cadence,required"`
	ConversionRate                  float64                                                        `json:"conversion_rate,required,nullable"`
	CreatedAt                       time.Time                                                      `json:"created_at,required" format:"date-time"`
	CreditAllocation                PriceGroupedWithMeteredMinimumPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                        string                                                         `json:"currency,required"`
	Discount                        shared.Discount                                                `json:"discount,required,nullable"`
	ExternalPriceID                 string                                                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity              float64                                                        `json:"fixed_price_quantity,required,nullable"`
	GroupedWithMeteredMinimumConfig map[string]interface{}                                         `json:"grouped_with_metered_minimum_config,required"`
	InvoicingCycleConfiguration     PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                            PriceGroupedWithMeteredMinimumPriceItem                        `json:"item,required"`
	Maximum                         PriceGroupedWithMeteredMinimumPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount                   string                                                         `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                            `json:"metadata,required"`
	Minimum        PriceGroupedWithMeteredMinimumPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                                       `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedWithMeteredMinimumPriceModelType `json:"model_type,required"`
	Name           string                                       `json:"name,required"`
	PlanPhaseOrder int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedWithMeteredMinimumPricePriceType `json:"price_type,required"`
	JSON           priceGroupedWithMeteredMinimumPriceJSON      `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceJSON contains the JSON metadata for the
// struct [PriceGroupedWithMeteredMinimumPrice]
type priceGroupedWithMeteredMinimumPriceJSON struct {
	ID                              apijson.Field
	BillableMetric                  apijson.Field
	BillingCycleConfiguration       apijson.Field
	Cadence                         apijson.Field
	ConversionRate                  apijson.Field
	CreatedAt                       apijson.Field
	CreditAllocation                apijson.Field
	Currency                        apijson.Field
	Discount                        apijson.Field
	ExternalPriceID                 apijson.Field
	FixedPriceQuantity              apijson.Field
	GroupedWithMeteredMinimumConfig apijson.Field
	InvoicingCycleConfiguration     apijson.Field
	Item                            apijson.Field
	Maximum                         apijson.Field
	MaximumAmount                   apijson.Field
	Metadata                        apijson.Field
	Minimum                         apijson.Field
	MinimumAmount                   apijson.Field
	ModelType                       apijson.Field
	Name                            apijson.Field
	PlanPhaseOrder                  apijson.Field
	PriceType                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedWithMeteredMinimumPrice) implementsPrice() {}

type PriceGroupedWithMeteredMinimumPriceBillableMetric struct {
	ID   string                                                `json:"id,required"`
	JSON priceGroupedWithMeteredMinimumPriceBillableMetricJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceBillableMetricJSON contains the JSON metadata
// for the struct [PriceGroupedWithMeteredMinimumPriceBillableMetric]
type priceGroupedWithMeteredMinimumPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	Duration     int64                                                                    `json:"duration,required"`
	DurationUnit PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedWithMeteredMinimumPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceBillingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceGroupedWithMeteredMinimumPriceBillingCycleConfiguration]
type priceGroupedWithMeteredMinimumPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPriceCadence string

const (
	PriceGroupedWithMeteredMinimumPriceCadenceOneTime    PriceGroupedWithMeteredMinimumPriceCadence = "one_time"
	PriceGroupedWithMeteredMinimumPriceCadenceMonthly    PriceGroupedWithMeteredMinimumPriceCadence = "monthly"
	PriceGroupedWithMeteredMinimumPriceCadenceQuarterly  PriceGroupedWithMeteredMinimumPriceCadence = "quarterly"
	PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual PriceGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	PriceGroupedWithMeteredMinimumPriceCadenceAnnual     PriceGroupedWithMeteredMinimumPriceCadence = "annual"
	PriceGroupedWithMeteredMinimumPriceCadenceCustom     PriceGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r PriceGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceCadenceOneTime, PriceGroupedWithMeteredMinimumPriceCadenceMonthly, PriceGroupedWithMeteredMinimumPriceCadenceQuarterly, PriceGroupedWithMeteredMinimumPriceCadenceSemiAnnual, PriceGroupedWithMeteredMinimumPriceCadenceAnnual, PriceGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPriceCreditAllocation struct {
	AllowsRollover bool                                                    `json:"allows_rollover,required"`
	Currency       string                                                  `json:"currency,required"`
	JSON           priceGroupedWithMeteredMinimumPriceCreditAllocationJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceCreditAllocationJSON contains the JSON
// metadata for the struct [PriceGroupedWithMeteredMinimumPriceCreditAllocation]
type priceGroupedWithMeteredMinimumPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                      `json:"duration,required"`
	DurationUnit PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationJSON contains the
// JSON metadata for the struct
// [PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration]
type priceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPriceItem struct {
	ID   string                                      `json:"id,required"`
	Name string                                      `json:"name,required"`
	JSON priceGroupedWithMeteredMinimumPriceItemJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceItemJSON contains the JSON metadata for the
// struct [PriceGroupedWithMeteredMinimumPriceItem]
type priceGroupedWithMeteredMinimumPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                         `json:"maximum_amount,required"`
	JSON          priceGroupedWithMeteredMinimumPriceMaximumJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceMaximumJSON contains the JSON metadata for
// the struct [PriceGroupedWithMeteredMinimumPriceMaximum]
type priceGroupedWithMeteredMinimumPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                         `json:"minimum_amount,required"`
	JSON          priceGroupedWithMeteredMinimumPriceMinimumJSON `json:"-"`
}

// priceGroupedWithMeteredMinimumPriceMinimumJSON contains the JSON metadata for
// the struct [PriceGroupedWithMeteredMinimumPriceMinimum]
type priceGroupedWithMeteredMinimumPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedWithMeteredMinimumPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedWithMeteredMinimumPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedWithMeteredMinimumPriceModelType string

const (
	PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum PriceGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r PriceGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type PriceGroupedWithMeteredMinimumPricePriceType string

const (
	PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice PriceGroupedWithMeteredMinimumPricePriceType = "usage_price"
	PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice PriceGroupedWithMeteredMinimumPricePriceType = "fixed_price"
)

func (r PriceGroupedWithMeteredMinimumPricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedWithMeteredMinimumPricePriceTypeUsagePrice, PriceGroupedWithMeteredMinimumPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePrice struct {
	ID                          string                                                     `json:"id,required"`
	BillableMetric              PriceMatrixWithDisplayNamePriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceMatrixWithDisplayNamePriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceMatrixWithDisplayNamePriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                    `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                  `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceMatrixWithDisplayNamePriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                     `json:"currency,required"`
	Discount                    shared.Discount                                            `json:"discount,required,nullable"`
	ExternalPriceID             string                                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                    `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceMatrixWithDisplayNamePriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceMatrixWithDisplayNamePriceItem                        `json:"item,required"`
	MatrixWithDisplayNameConfig map[string]interface{}                                     `json:"matrix_with_display_name_config,required"`
	Maximum                     PriceMatrixWithDisplayNamePriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                     `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                        `json:"metadata,required"`
	Minimum        PriceMatrixWithDisplayNamePriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                                   `json:"minimum_amount,required,nullable"`
	ModelType      PriceMatrixWithDisplayNamePriceModelType `json:"model_type,required"`
	Name           string                                   `json:"name,required"`
	PlanPhaseOrder int64                                    `json:"plan_phase_order,required,nullable"`
	PriceType      PriceMatrixWithDisplayNamePricePriceType `json:"price_type,required"`
	JSON           priceMatrixWithDisplayNamePriceJSON      `json:"-"`
}

// priceMatrixWithDisplayNamePriceJSON contains the JSON metadata for the struct
// [PriceMatrixWithDisplayNamePrice]
type priceMatrixWithDisplayNamePriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	MatrixWithDisplayNameConfig apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceMatrixWithDisplayNamePrice) implementsPrice() {}

type PriceMatrixWithDisplayNamePriceBillableMetric struct {
	ID   string                                            `json:"id,required"`
	JSON priceMatrixWithDisplayNamePriceBillableMetricJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceMatrixWithDisplayNamePriceBillableMetric]
type priceMatrixWithDisplayNamePriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	Duration     int64                                                                `json:"duration,required"`
	DurationUnit PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixWithDisplayNamePriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixWithDisplayNamePriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithDisplayNamePriceBillingCycleConfiguration]
type priceMatrixWithDisplayNamePriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PriceMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePriceCadence string

const (
	PriceMatrixWithDisplayNamePriceCadenceOneTime    PriceMatrixWithDisplayNamePriceCadence = "one_time"
	PriceMatrixWithDisplayNamePriceCadenceMonthly    PriceMatrixWithDisplayNamePriceCadence = "monthly"
	PriceMatrixWithDisplayNamePriceCadenceQuarterly  PriceMatrixWithDisplayNamePriceCadence = "quarterly"
	PriceMatrixWithDisplayNamePriceCadenceSemiAnnual PriceMatrixWithDisplayNamePriceCadence = "semi_annual"
	PriceMatrixWithDisplayNamePriceCadenceAnnual     PriceMatrixWithDisplayNamePriceCadence = "annual"
	PriceMatrixWithDisplayNamePriceCadenceCustom     PriceMatrixWithDisplayNamePriceCadence = "custom"
)

func (r PriceMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceCadenceOneTime, PriceMatrixWithDisplayNamePriceCadenceMonthly, PriceMatrixWithDisplayNamePriceCadenceQuarterly, PriceMatrixWithDisplayNamePriceCadenceSemiAnnual, PriceMatrixWithDisplayNamePriceCadenceAnnual, PriceMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePriceCreditAllocation struct {
	AllowsRollover bool                                                `json:"allows_rollover,required"`
	Currency       string                                              `json:"currency,required"`
	JSON           priceMatrixWithDisplayNamePriceCreditAllocationJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceCreditAllocationJSON contains the JSON metadata
// for the struct [PriceMatrixWithDisplayNamePriceCreditAllocation]
type priceMatrixWithDisplayNamePriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                  `json:"duration,required"`
	DurationUnit PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceMatrixWithDisplayNamePriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceMatrixWithDisplayNamePriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithDisplayNamePriceInvoicingCycleConfiguration]
type priceMatrixWithDisplayNamePriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PriceMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePriceItem struct {
	ID   string                                  `json:"id,required"`
	Name string                                  `json:"name,required"`
	JSON priceMatrixWithDisplayNamePriceItemJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceItemJSON contains the JSON metadata for the
// struct [PriceMatrixWithDisplayNamePriceItem]
type priceMatrixWithDisplayNamePriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                     `json:"maximum_amount,required"`
	JSON          priceMatrixWithDisplayNamePriceMaximumJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceMaximumJSON contains the JSON metadata for the
// struct [PriceMatrixWithDisplayNamePriceMaximum]
type priceMatrixWithDisplayNamePriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                     `json:"minimum_amount,required"`
	JSON          priceMatrixWithDisplayNamePriceMinimumJSON `json:"-"`
}

// priceMatrixWithDisplayNamePriceMinimumJSON contains the JSON metadata for the
// struct [PriceMatrixWithDisplayNamePriceMinimum]
type priceMatrixWithDisplayNamePriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixWithDisplayNamePriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceMatrixWithDisplayNamePriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixWithDisplayNamePriceModelType string

const (
	PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName PriceMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r PriceMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type PriceMatrixWithDisplayNamePricePriceType string

const (
	PriceMatrixWithDisplayNamePricePriceTypeUsagePrice PriceMatrixWithDisplayNamePricePriceType = "usage_price"
	PriceMatrixWithDisplayNamePricePriceTypeFixedPrice PriceMatrixWithDisplayNamePricePriceType = "fixed_price"
)

func (r PriceMatrixWithDisplayNamePricePriceType) IsKnown() bool {
	switch r {
	case PriceMatrixWithDisplayNamePricePriceTypeUsagePrice, PriceMatrixWithDisplayNamePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceBulkWithProrationPrice struct {
	ID                          string                                                 `json:"id,required"`
	BillableMetric              PriceBulkWithProrationPriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceBulkWithProrationPriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	BulkWithProrationConfig     map[string]interface{}                                 `json:"bulk_with_proration_config,required"`
	Cadence                     PriceBulkWithProrationPriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                              `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceBulkWithProrationPriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                 `json:"currency,required"`
	Discount                    shared.Discount                                        `json:"discount,required,nullable"`
	ExternalPriceID             string                                                 `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                `json:"fixed_price_quantity,required,nullable"`
	InvoicingCycleConfiguration PriceBulkWithProrationPriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceBulkWithProrationPriceItem                        `json:"item,required"`
	Maximum                     PriceBulkWithProrationPriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                 `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                    `json:"metadata,required"`
	Minimum        PriceBulkWithProrationPriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                               `json:"minimum_amount,required,nullable"`
	ModelType      PriceBulkWithProrationPriceModelType `json:"model_type,required"`
	Name           string                               `json:"name,required"`
	PlanPhaseOrder int64                                `json:"plan_phase_order,required,nullable"`
	PriceType      PriceBulkWithProrationPricePriceType `json:"price_type,required"`
	JSON           priceBulkWithProrationPriceJSON      `json:"-"`
}

// priceBulkWithProrationPriceJSON contains the JSON metadata for the struct
// [PriceBulkWithProrationPrice]
type priceBulkWithProrationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	BulkWithProrationConfig     apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceBulkWithProrationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceBulkWithProrationPrice) implementsPrice() {}

type PriceBulkWithProrationPriceBillableMetric struct {
	ID   string                                        `json:"id,required"`
	JSON priceBulkWithProrationPriceBillableMetricJSON `json:"-"`
}

// priceBulkWithProrationPriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceBulkWithProrationPriceBillableMetric]
type priceBulkWithProrationPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceBillingCycleConfiguration struct {
	Duration     int64                                                            `json:"duration,required"`
	DurationUnit PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkWithProrationPriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceBulkWithProrationPriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceBulkWithProrationPriceBillingCycleConfiguration]
type priceBulkWithProrationPriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkWithProrationPriceCadence string

const (
	PriceBulkWithProrationPriceCadenceOneTime    PriceBulkWithProrationPriceCadence = "one_time"
	PriceBulkWithProrationPriceCadenceMonthly    PriceBulkWithProrationPriceCadence = "monthly"
	PriceBulkWithProrationPriceCadenceQuarterly  PriceBulkWithProrationPriceCadence = "quarterly"
	PriceBulkWithProrationPriceCadenceSemiAnnual PriceBulkWithProrationPriceCadence = "semi_annual"
	PriceBulkWithProrationPriceCadenceAnnual     PriceBulkWithProrationPriceCadence = "annual"
	PriceBulkWithProrationPriceCadenceCustom     PriceBulkWithProrationPriceCadence = "custom"
)

func (r PriceBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceCadenceOneTime, PriceBulkWithProrationPriceCadenceMonthly, PriceBulkWithProrationPriceCadenceQuarterly, PriceBulkWithProrationPriceCadenceSemiAnnual, PriceBulkWithProrationPriceCadenceAnnual, PriceBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type PriceBulkWithProrationPriceCreditAllocation struct {
	AllowsRollover bool                                            `json:"allows_rollover,required"`
	Currency       string                                          `json:"currency,required"`
	JSON           priceBulkWithProrationPriceCreditAllocationJSON `json:"-"`
}

// priceBulkWithProrationPriceCreditAllocationJSON contains the JSON metadata for
// the struct [PriceBulkWithProrationPriceCreditAllocation]
type priceBulkWithProrationPriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceInvoicingCycleConfiguration struct {
	Duration     int64                                                              `json:"duration,required"`
	DurationUnit PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceBulkWithProrationPriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceBulkWithProrationPriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct [PriceBulkWithProrationPriceInvoicingCycleConfiguration]
type priceBulkWithProrationPriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceBulkWithProrationPriceItem struct {
	ID   string                              `json:"id,required"`
	Name string                              `json:"name,required"`
	JSON priceBulkWithProrationPriceItemJSON `json:"-"`
}

// priceBulkWithProrationPriceItemJSON contains the JSON metadata for the struct
// [PriceBulkWithProrationPriceItem]
type priceBulkWithProrationPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                 `json:"maximum_amount,required"`
	JSON          priceBulkWithProrationPriceMaximumJSON `json:"-"`
}

// priceBulkWithProrationPriceMaximumJSON contains the JSON metadata for the struct
// [PriceBulkWithProrationPriceMaximum]
type priceBulkWithProrationPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                 `json:"minimum_amount,required"`
	JSON          priceBulkWithProrationPriceMinimumJSON `json:"-"`
}

// priceBulkWithProrationPriceMinimumJSON contains the JSON metadata for the struct
// [PriceBulkWithProrationPriceMinimum]
type priceBulkWithProrationPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceBulkWithProrationPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceBulkWithProrationPriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceBulkWithProrationPriceModelType string

const (
	PriceBulkWithProrationPriceModelTypeBulkWithProration PriceBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r PriceBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type PriceBulkWithProrationPricePriceType string

const (
	PriceBulkWithProrationPricePriceTypeUsagePrice PriceBulkWithProrationPricePriceType = "usage_price"
	PriceBulkWithProrationPricePriceTypeFixedPrice PriceBulkWithProrationPricePriceType = "fixed_price"
)

func (r PriceBulkWithProrationPricePriceType) IsKnown() bool {
	switch r {
	case PriceBulkWithProrationPricePriceTypeUsagePrice, PriceBulkWithProrationPricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePrice struct {
	ID                          string                                                    `json:"id,required"`
	BillableMetric              PriceGroupedTieredPackagePriceBillableMetric              `json:"billable_metric,required,nullable"`
	BillingCycleConfiguration   PriceGroupedTieredPackagePriceBillingCycleConfiguration   `json:"billing_cycle_configuration,required"`
	Cadence                     PriceGroupedTieredPackagePriceCadence                     `json:"cadence,required"`
	ConversionRate              float64                                                   `json:"conversion_rate,required,nullable"`
	CreatedAt                   time.Time                                                 `json:"created_at,required" format:"date-time"`
	CreditAllocation            PriceGroupedTieredPackagePriceCreditAllocation            `json:"credit_allocation,required,nullable"`
	Currency                    string                                                    `json:"currency,required"`
	Discount                    shared.Discount                                           `json:"discount,required,nullable"`
	ExternalPriceID             string                                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                                   `json:"fixed_price_quantity,required,nullable"`
	GroupedTieredPackageConfig  map[string]interface{}                                    `json:"grouped_tiered_package_config,required"`
	InvoicingCycleConfiguration PriceGroupedTieredPackagePriceInvoicingCycleConfiguration `json:"invoicing_cycle_configuration,required,nullable"`
	Item                        PriceGroupedTieredPackagePriceItem                        `json:"item,required"`
	Maximum                     PriceGroupedTieredPackagePriceMaximum                     `json:"maximum,required,nullable"`
	MaximumAmount               string                                                    `json:"maximum_amount,required,nullable"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata       map[string]string                       `json:"metadata,required"`
	Minimum        PriceGroupedTieredPackagePriceMinimum   `json:"minimum,required,nullable"`
	MinimumAmount  string                                  `json:"minimum_amount,required,nullable"`
	ModelType      PriceGroupedTieredPackagePriceModelType `json:"model_type,required"`
	Name           string                                  `json:"name,required"`
	PlanPhaseOrder int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType      PriceGroupedTieredPackagePricePriceType `json:"price_type,required"`
	JSON           priceGroupedTieredPackagePriceJSON      `json:"-"`
}

// priceGroupedTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPackagePrice]
type priceGroupedTieredPackagePriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	BillingCycleConfiguration   apijson.Field
	Cadence                     apijson.Field
	ConversionRate              apijson.Field
	CreatedAt                   apijson.Field
	CreditAllocation            apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	GroupedTieredPackageConfig  apijson.Field
	InvoicingCycleConfiguration apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Metadata                    apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceJSON) RawJSON() string {
	return r.raw
}

func (r PriceGroupedTieredPackagePrice) implementsPrice() {}

type PriceGroupedTieredPackagePriceBillableMetric struct {
	ID   string                                           `json:"id,required"`
	JSON priceGroupedTieredPackagePriceBillableMetricJSON `json:"-"`
}

// priceGroupedTieredPackagePriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceGroupedTieredPackagePriceBillableMetric]
type priceGroupedTieredPackagePriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceBillableMetricJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceBillingCycleConfiguration struct {
	Duration     int64                                                               `json:"duration,required"`
	DurationUnit PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedTieredPackagePriceBillingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedTieredPackagePriceBillingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceGroupedTieredPackagePriceBillingCycleConfiguration]
type priceGroupedTieredPackagePriceBillingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceBillingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceBillingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PriceGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePriceCadence string

const (
	PriceGroupedTieredPackagePriceCadenceOneTime    PriceGroupedTieredPackagePriceCadence = "one_time"
	PriceGroupedTieredPackagePriceCadenceMonthly    PriceGroupedTieredPackagePriceCadence = "monthly"
	PriceGroupedTieredPackagePriceCadenceQuarterly  PriceGroupedTieredPackagePriceCadence = "quarterly"
	PriceGroupedTieredPackagePriceCadenceSemiAnnual PriceGroupedTieredPackagePriceCadence = "semi_annual"
	PriceGroupedTieredPackagePriceCadenceAnnual     PriceGroupedTieredPackagePriceCadence = "annual"
	PriceGroupedTieredPackagePriceCadenceCustom     PriceGroupedTieredPackagePriceCadence = "custom"
)

func (r PriceGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceCadenceOneTime, PriceGroupedTieredPackagePriceCadenceMonthly, PriceGroupedTieredPackagePriceCadenceQuarterly, PriceGroupedTieredPackagePriceCadenceSemiAnnual, PriceGroupedTieredPackagePriceCadenceAnnual, PriceGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePriceCreditAllocation struct {
	AllowsRollover bool                                               `json:"allows_rollover,required"`
	Currency       string                                             `json:"currency,required"`
	JSON           priceGroupedTieredPackagePriceCreditAllocationJSON `json:"-"`
}

// priceGroupedTieredPackagePriceCreditAllocationJSON contains the JSON metadata
// for the struct [PriceGroupedTieredPackagePriceCreditAllocation]
type priceGroupedTieredPackagePriceCreditAllocationJSON struct {
	AllowsRollover apijson.Field
	Currency       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceCreditAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceCreditAllocationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	Duration     int64                                                                 `json:"duration,required"`
	DurationUnit PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit `json:"duration_unit,required"`
	JSON         priceGroupedTieredPackagePriceInvoicingCycleConfigurationJSON         `json:"-"`
}

// priceGroupedTieredPackagePriceInvoicingCycleConfigurationJSON contains the JSON
// metadata for the struct
// [PriceGroupedTieredPackagePriceInvoicingCycleConfiguration]
type priceGroupedTieredPackagePriceInvoicingCycleConfigurationJSON struct {
	Duration     apijson.Field
	DurationUnit apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceInvoicingCycleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceInvoicingCycleConfigurationJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PriceGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePriceItem struct {
	ID   string                                 `json:"id,required"`
	Name string                                 `json:"name,required"`
	JSON priceGroupedTieredPackagePriceItemJSON `json:"-"`
}

// priceGroupedTieredPackagePriceItemJSON contains the JSON metadata for the struct
// [PriceGroupedTieredPackagePriceItem]
type priceGroupedTieredPackagePriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceItemJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string                                    `json:"maximum_amount,required"`
	JSON          priceGroupedTieredPackagePriceMaximumJSON `json:"-"`
}

// priceGroupedTieredPackagePriceMaximumJSON contains the JSON metadata for the
// struct [PriceGroupedTieredPackagePriceMaximum]
type priceGroupedTieredPackagePriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceMaximumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string                                    `json:"minimum_amount,required"`
	JSON          priceGroupedTieredPackagePriceMinimumJSON `json:"-"`
}

// priceGroupedTieredPackagePriceMinimumJSON contains the JSON metadata for the
// struct [PriceGroupedTieredPackagePriceMinimum]
type priceGroupedTieredPackagePriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceGroupedTieredPackagePriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r priceGroupedTieredPackagePriceMinimumJSON) RawJSON() string {
	return r.raw
}

type PriceGroupedTieredPackagePriceModelType string

const (
	PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage PriceGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r PriceGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PriceGroupedTieredPackagePricePriceType string

const (
	PriceGroupedTieredPackagePricePriceTypeUsagePrice PriceGroupedTieredPackagePricePriceType = "usage_price"
	PriceGroupedTieredPackagePricePriceTypeFixedPrice PriceGroupedTieredPackagePricePriceType = "fixed_price"
)

func (r PriceGroupedTieredPackagePricePriceType) IsKnown() bool {
	switch r {
	case PriceGroupedTieredPackagePricePriceTypeUsagePrice, PriceGroupedTieredPackagePricePriceTypeFixedPrice:
		return true
	}
	return false
}

type PriceCadence string

const (
	PriceCadenceOneTime    PriceCadence = "one_time"
	PriceCadenceMonthly    PriceCadence = "monthly"
	PriceCadenceQuarterly  PriceCadence = "quarterly"
	PriceCadenceSemiAnnual PriceCadence = "semi_annual"
	PriceCadenceAnnual     PriceCadence = "annual"
	PriceCadenceCustom     PriceCadence = "custom"
)

func (r PriceCadence) IsKnown() bool {
	switch r {
	case PriceCadenceOneTime, PriceCadenceMonthly, PriceCadenceQuarterly, PriceCadenceSemiAnnual, PriceCadenceAnnual, PriceCadenceCustom:
		return true
	}
	return false
}

type PriceModelType string

const (
	PriceModelTypeUnit                       PriceModelType = "unit"
	PriceModelTypePackage                    PriceModelType = "package"
	PriceModelTypeMatrix                     PriceModelType = "matrix"
	PriceModelTypeTiered                     PriceModelType = "tiered"
	PriceModelTypeTieredBps                  PriceModelType = "tiered_bps"
	PriceModelTypeBps                        PriceModelType = "bps"
	PriceModelTypeBulkBps                    PriceModelType = "bulk_bps"
	PriceModelTypeBulk                       PriceModelType = "bulk"
	PriceModelTypeThresholdTotalAmount       PriceModelType = "threshold_total_amount"
	PriceModelTypeTieredPackage              PriceModelType = "tiered_package"
	PriceModelTypeGroupedTiered              PriceModelType = "grouped_tiered"
	PriceModelTypeTieredWithMinimum          PriceModelType = "tiered_with_minimum"
	PriceModelTypeTieredPackageWithMinimum   PriceModelType = "tiered_package_with_minimum"
	PriceModelTypePackageWithAllocation      PriceModelType = "package_with_allocation"
	PriceModelTypeUnitWithPercent            PriceModelType = "unit_with_percent"
	PriceModelTypeMatrixWithAllocation       PriceModelType = "matrix_with_allocation"
	PriceModelTypeTieredWithProration        PriceModelType = "tiered_with_proration"
	PriceModelTypeUnitWithProration          PriceModelType = "unit_with_proration"
	PriceModelTypeGroupedAllocation          PriceModelType = "grouped_allocation"
	PriceModelTypeGroupedWithProratedMinimum PriceModelType = "grouped_with_prorated_minimum"
	PriceModelTypeGroupedWithMeteredMinimum  PriceModelType = "grouped_with_metered_minimum"
	PriceModelTypeMatrixWithDisplayName      PriceModelType = "matrix_with_display_name"
	PriceModelTypeBulkWithProration          PriceModelType = "bulk_with_proration"
	PriceModelTypeGroupedTieredPackage       PriceModelType = "grouped_tiered_package"
)

func (r PriceModelType) IsKnown() bool {
	switch r {
	case PriceModelTypeUnit, PriceModelTypePackage, PriceModelTypeMatrix, PriceModelTypeTiered, PriceModelTypeTieredBps, PriceModelTypeBps, PriceModelTypeBulkBps, PriceModelTypeBulk, PriceModelTypeThresholdTotalAmount, PriceModelTypeTieredPackage, PriceModelTypeGroupedTiered, PriceModelTypeTieredWithMinimum, PriceModelTypeTieredPackageWithMinimum, PriceModelTypePackageWithAllocation, PriceModelTypeUnitWithPercent, PriceModelTypeMatrixWithAllocation, PriceModelTypeTieredWithProration, PriceModelTypeUnitWithProration, PriceModelTypeGroupedAllocation, PriceModelTypeGroupedWithProratedMinimum, PriceModelTypeGroupedWithMeteredMinimum, PriceModelTypeMatrixWithDisplayName, PriceModelTypeBulkWithProration, PriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type PricePriceType string

const (
	PricePriceTypeUsagePrice PricePriceType = "usage_price"
	PricePriceTypeFixedPrice PricePriceType = "fixed_price"
)

func (r PricePriceType) IsKnown() bool {
	switch r {
	case PricePriceTypeUsagePrice, PricePriceTypeFixedPrice:
		return true
	}
	return false
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

// This interface is a union satisfied by one of the following:
// [PriceNewParamsNewFloatingUnitPrice], [PriceNewParamsNewFloatingPackagePrice],
// [PriceNewParamsNewFloatingMatrixPrice],
// [PriceNewParamsNewFloatingMatrixWithAllocationPrice],
// [PriceNewParamsNewFloatingTieredPrice],
// [PriceNewParamsNewFloatingTieredBpsPrice], [PriceNewParamsNewFloatingBpsPrice],
// [PriceNewParamsNewFloatingBulkBpsPrice], [PriceNewParamsNewFloatingBulkPrice],
// [PriceNewParamsNewFloatingThresholdTotalAmountPrice],
// [PriceNewParamsNewFloatingTieredPackagePrice],
// [PriceNewParamsNewFloatingGroupedTieredPrice],
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
// [PriceNewParamsNewFloatingGroupedTieredPackagePrice].
type PriceNewParams interface {
	ImplementsPriceNewParams()
}

type PriceNewParamsNewFloatingUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                      `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                       `json:"name,required"`
	UnitConfig param.Field[PriceNewParamsNewFloatingUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PriceNewParamsNewFloatingUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                             `json:"name,required"`
	PackageConfig param.Field[PriceNewParamsNewFloatingPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r PriceNewParamsNewFloatingPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                           `json:"item_id,required"`
	MatrixConfig param.Field[PriceNewParamsNewFloatingMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PriceNewParamsNewFloatingMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PriceNewParamsNewFloatingMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PriceNewParamsNewFloatingMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PriceNewParamsNewFloatingMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID                     param.Field[string]                                                                       `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                        `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                           `json:"name,required"`
	TieredConfig param.Field[PriceNewParamsNewFloatingTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PriceNewParamsNewFloatingTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewFloatingTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PriceNewParamsNewFloatingTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                           `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                 `json:"name,required"`
	TieredBpsConfig param.Field[PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PriceNewParamsNewFloatingTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredBpsPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredBpsPriceCadence string

const (
	PriceNewParamsNewFloatingTieredBpsPriceCadenceAnnual     PriceNewParamsNewFloatingTieredBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceSemiAnnual PriceNewParamsNewFloatingTieredBpsPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceMonthly    PriceNewParamsNewFloatingTieredBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceQuarterly  PriceNewParamsNewFloatingTieredBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceOneTime    PriceNewParamsNewFloatingTieredBpsPriceCadence = "one_time"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceCustom     PriceNewParamsNewFloatingTieredBpsPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredBpsPriceCadenceAnnual, PriceNewParamsNewFloatingTieredBpsPriceCadenceSemiAnnual, PriceNewParamsNewFloatingTieredBpsPriceCadenceMonthly, PriceNewParamsNewFloatingTieredBpsPriceCadenceQuarterly, PriceNewParamsNewFloatingTieredBpsPriceCadenceOneTime, PriceNewParamsNewFloatingTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredBpsPriceModelType string

const (
	PriceNewParamsNewFloatingTieredBpsPriceModelTypeTieredBps PriceNewParamsNewFloatingTieredBpsPriceModelType = "tiered_bps"
)

func (r PriceNewParamsNewFloatingTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewFloatingTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBpsPrice struct {
	BpsConfig param.Field[PriceNewParamsNewFloatingBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                     `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PriceNewParamsNewFloatingBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingBpsPrice) ImplementsPriceNewParams() {

}

type PriceNewParamsNewFloatingBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewFloatingBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingBpsPriceCadence string

const (
	PriceNewParamsNewFloatingBpsPriceCadenceAnnual     PriceNewParamsNewFloatingBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingBpsPriceCadenceSemiAnnual PriceNewParamsNewFloatingBpsPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingBpsPriceCadenceMonthly    PriceNewParamsNewFloatingBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingBpsPriceCadenceQuarterly  PriceNewParamsNewFloatingBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBpsPriceCadenceOneTime    PriceNewParamsNewFloatingBpsPriceCadence = "one_time"
	PriceNewParamsNewFloatingBpsPriceCadenceCustom     PriceNewParamsNewFloatingBpsPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBpsPriceCadenceAnnual, PriceNewParamsNewFloatingBpsPriceCadenceSemiAnnual, PriceNewParamsNewFloatingBpsPriceCadenceMonthly, PriceNewParamsNewFloatingBpsPriceCadenceQuarterly, PriceNewParamsNewFloatingBpsPriceCadenceOneTime, PriceNewParamsNewFloatingBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBpsPriceModelType string

const (
	PriceNewParamsNewFloatingBpsPriceModelTypeBps PriceNewParamsNewFloatingBpsPriceModelType = "bps"
)

func (r PriceNewParamsNewFloatingBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkBpsPrice struct {
	BulkBpsConfig param.Field[PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingBulkBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                         `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewFloatingBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r PriceNewParamsNewFloatingBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingBulkBpsPrice) ImplementsPriceNewParams() {

}

type PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewFloatingBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingBulkBpsPriceCadence string

const (
	PriceNewParamsNewFloatingBulkBpsPriceCadenceAnnual     PriceNewParamsNewFloatingBulkBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceSemiAnnual PriceNewParamsNewFloatingBulkBpsPriceCadence = "semi_annual"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceMonthly    PriceNewParamsNewFloatingBulkBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceQuarterly  PriceNewParamsNewFloatingBulkBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceOneTime    PriceNewParamsNewFloatingBulkBpsPriceCadence = "one_time"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceCustom     PriceNewParamsNewFloatingBulkBpsPriceCadence = "custom"
)

func (r PriceNewParamsNewFloatingBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkBpsPriceCadenceAnnual, PriceNewParamsNewFloatingBulkBpsPriceCadenceSemiAnnual, PriceNewParamsNewFloatingBulkBpsPriceCadenceMonthly, PriceNewParamsNewFloatingBulkBpsPriceCadenceQuarterly, PriceNewParamsNewFloatingBulkBpsPriceCadenceOneTime, PriceNewParamsNewFloatingBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkBpsPriceModelType string

const (
	PriceNewParamsNewFloatingBulkBpsPriceModelTypeBulkBps PriceNewParamsNewFloatingBulkBpsPriceModelType = "bulk_bps"
)

func (r PriceNewParamsNewFloatingBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingBulkPrice struct {
	BulkConfig param.Field[PriceNewParamsNewFloatingBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

type PriceNewParamsNewFloatingBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PriceNewParamsNewFloatingBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewFloatingBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PriceNewParamsNewFloatingBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type PriceNewParamsNewFloatingMatrixWithDisplayNamePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingMatrixWithDisplayNamePriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// The id of the item the plan will be associated with.
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
	BillingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// Within each billing cycle, specifies the cadence at which invoices are produced.
	// If unspecified, a single invoice is produced per billing cycle.
	InvoicingCycleConfiguration param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
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

// For custom cadence: specifies the duration of the billing period in days or
// months.
type PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, PriceNewParamsNewFloatingGroupedTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
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
	// [computed property](../guides/extensibility/advanced-metrics#computed-properties)
	// used to filter the underlying billable metric
	Filter param.Field[string] `json:"filter"`
	// Properties (or
	// [computed properties](../guides/extensibility/advanced-metrics#computed-properties))
	// used to group the underlying billable metric
	GroupingKeys param.Field[[]string] `json:"grouping_keys"`
}

func (r PriceEvaluateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
