// File generated from our OpenAPI spec by Stainless.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/internal/shared"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// PriceService contains methods and other services that help with interacting with
// the orb API. Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly, and
// instead use the [NewPriceService] method instead.
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

// This endpoint is used to list all add-on prices created using the
// [price creation endpoint](../reference/create-price).
func (r *PriceService) List(ctx context.Context, query PriceListParams, opts ...option.RequestOption) (res *shared.Page[Price], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *PriceService) ListAutoPaging(ctx context.Context, query PriceListParams, opts ...option.RequestOption) *shared.PageAutoPager[Price] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint returns a price given an identifier.
func (r *PriceService) Fetch(ctx context.Context, priceID string, opts ...option.RequestOption) (res *Price, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("prices/%s", priceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DiscountParam struct {
	DiscountType param.Field[DiscountDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
	// List of price_ids that this discount applies to. For plan/plan phase discounts,
	// this can be a subset of prices.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `trial`
	TrialAmountDiscount param.Field[string] `json:"trial_amount_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r DiscountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiscountDiscountType string

const (
	DiscountDiscountTypePercentage DiscountDiscountType = "percentage"
	DiscountDiscountTypeTrial      DiscountDiscountType = "trial"
	DiscountDiscountTypeUsage      DiscountDiscountType = "usage"
	DiscountDiscountTypeAmount     DiscountDiscountType = "amount"
)

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
// ### Fixed fees
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
// [PriceBulkPrice], [PriceTestRatingFunctionPrice], [PriceFivetranExamplePrice],
// [PriceThresholdTotalAmountPrice], [PriceTieredPackagePrice],
// [PriceTieredWithMinimumPrice] or [PricePackageWithAllocationPrice].
type Price interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Price)(nil)).Elem(),
		"model_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"unit\"",
			Type:               reflect.TypeOf(PriceUnitPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"package\"",
			Type:               reflect.TypeOf(PricePackagePrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"matrix\"",
			Type:               reflect.TypeOf(PriceMatrixPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tiered\"",
			Type:               reflect.TypeOf(PriceTieredPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tiered_bps\"",
			Type:               reflect.TypeOf(PriceTieredBpsPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"bps\"",
			Type:               reflect.TypeOf(PriceBpsPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"bulk_bps\"",
			Type:               reflect.TypeOf(PriceBulkBpsPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"bulk\"",
			Type:               reflect.TypeOf(PriceBulkPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"test_rating_function\"",
			Type:               reflect.TypeOf(PriceTestRatingFunctionPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"fivetran_example\"",
			Type:               reflect.TypeOf(PriceFivetranExamplePrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"threshold_total_amount\"",
			Type:               reflect.TypeOf(PriceThresholdTotalAmountPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tiered_package\"",
			Type:               reflect.TypeOf(PriceTieredPackagePrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"tiered_with_minimum\"",
			Type:               reflect.TypeOf(PriceTieredWithMinimumPrice{}),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			DiscriminatorValue: "\"package_with_allocation\"",
			Type:               reflect.TypeOf(PricePackageWithAllocationPrice{}),
		},
	)
}

type PriceUnitPrice struct {
	ID                 string                       `json:"id,required"`
	BillableMetric     PriceUnitPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceUnitPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                    `json:"created_at,required" format:"date-time"`
	Currency           string                       `json:"currency,required"`
	ExternalPriceID    string                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                      `json:"fixed_price_quantity,required,nullable"`
	Item               PriceUnitPriceItem           `json:"item,required"`
	ModelType          PriceUnitPriceModelType      `json:"model_type,required"`
	Name               string                       `json:"name,required"`
	PlanPhaseOrder     int64                        `json:"plan_phase_order,required,nullable"`
	PriceType          PriceUnitPricePriceType      `json:"price_type,required"`
	UnitConfig         PriceUnitPriceUnitConfig     `json:"unit_config,required"`
	Discount           InvoiceDiscount              `json:"discount,nullable"`
	Maximum            PriceUnitPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                       `json:"maximum_amount,nullable"`
	Minimum            PriceUnitPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                       `json:"minimum_amount,nullable"`
	JSON               priceUnitPriceJSON
}

// priceUnitPriceJSON contains the JSON metadata for the struct [PriceUnitPrice]
type priceUnitPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	UnitConfig         apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceUnitPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceUnitPrice) implementsPrice() {}

type PriceUnitPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceUnitPriceBillableMetricJSON
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

type PriceUnitPriceCadence string

const (
	PriceUnitPriceCadenceOneTime   PriceUnitPriceCadence = "one_time"
	PriceUnitPriceCadenceMonthly   PriceUnitPriceCadence = "monthly"
	PriceUnitPriceCadenceQuarterly PriceUnitPriceCadence = "quarterly"
	PriceUnitPriceCadenceAnnual    PriceUnitPriceCadence = "annual"
)

type PriceUnitPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceUnitPriceItemJSON
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

type PriceUnitPriceModelType string

const (
	PriceUnitPriceModelTypeUnit PriceUnitPriceModelType = "unit"
)

type PriceUnitPricePriceType string

const (
	PriceUnitPricePriceTypeUsagePrice PriceUnitPricePriceType = "usage_price"
	PriceUnitPricePriceTypeFixedPrice PriceUnitPricePriceType = "fixed_price"
)

type PriceUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount string `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor float64 `json:"scaling_factor,nullable"`
	JSON          priceUnitPriceUnitConfigJSON
}

// priceUnitPriceUnitConfigJSON contains the JSON metadata for the struct
// [PriceUnitPriceUnitConfig]
type priceUnitPriceUnitConfigJSON struct {
	UnitAmount    apijson.Field
	ScalingFactor apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PriceUnitPriceUnitConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceUnitPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceUnitPriceMaximumJSON
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

type PriceUnitPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceUnitPriceMinimumJSON
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

type PricePackagePrice struct {
	ID                 string                          `json:"id,required"`
	BillableMetric     PricePackagePriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PricePackagePriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                       `json:"created_at,required" format:"date-time"`
	Currency           string                          `json:"currency,required"`
	ExternalPriceID    string                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                         `json:"fixed_price_quantity,required,nullable"`
	Item               PricePackagePriceItem           `json:"item,required"`
	ModelType          PricePackagePriceModelType      `json:"model_type,required"`
	Name               string                          `json:"name,required"`
	PackageConfig      PricePackagePricePackageConfig  `json:"package_config,required"`
	PlanPhaseOrder     int64                           `json:"plan_phase_order,required,nullable"`
	PriceType          PricePackagePricePriceType      `json:"price_type,required"`
	Discount           InvoiceDiscount                 `json:"discount,nullable"`
	Maximum            PricePackagePriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                          `json:"maximum_amount,nullable"`
	Minimum            PricePackagePriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                          `json:"minimum_amount,nullable"`
	JSON               pricePackagePriceJSON
}

// pricePackagePriceJSON contains the JSON metadata for the struct
// [PricePackagePrice]
type pricePackagePriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PackageConfig      apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PricePackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PricePackagePrice) implementsPrice() {}

type PricePackagePriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON pricePackagePriceBillableMetricJSON
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

type PricePackagePriceCadence string

const (
	PricePackagePriceCadenceOneTime   PricePackagePriceCadence = "one_time"
	PricePackagePriceCadenceMonthly   PricePackagePriceCadence = "monthly"
	PricePackagePriceCadenceQuarterly PricePackagePriceCadence = "quarterly"
	PricePackagePriceCadenceAnnual    PricePackagePriceCadence = "annual"
)

type PricePackagePriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON pricePackagePriceItemJSON
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

type PricePackagePriceModelType string

const (
	PricePackagePriceModelTypePackage PricePackagePriceModelType = "package"
)

type PricePackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount string `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize int64 `json:"package_size,nullable"`
	JSON        pricePackagePricePackageConfigJSON
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

type PricePackagePricePriceType string

const (
	PricePackagePricePriceTypeUsagePrice PricePackagePricePriceType = "usage_price"
	PricePackagePricePriceTypeFixedPrice PricePackagePricePriceType = "fixed_price"
)

type PricePackagePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          pricePackagePriceMaximumJSON
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

type PricePackagePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          pricePackagePriceMinimumJSON
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

type PriceMatrixPrice struct {
	ID                 string                         `json:"id,required"`
	BillableMetric     PriceMatrixPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceMatrixPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                      `json:"created_at,required" format:"date-time"`
	Currency           string                         `json:"currency,required"`
	ExternalPriceID    string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                        `json:"fixed_price_quantity,required,nullable"`
	Item               PriceMatrixPriceItem           `json:"item,required"`
	MatrixConfig       PriceMatrixPriceMatrixConfig   `json:"matrix_config,required"`
	ModelType          PriceMatrixPriceModelType      `json:"model_type,required"`
	Name               string                         `json:"name,required"`
	PlanPhaseOrder     int64                          `json:"plan_phase_order,required,nullable"`
	PriceType          PriceMatrixPricePriceType      `json:"price_type,required"`
	Discount           InvoiceDiscount                `json:"discount,nullable"`
	Maximum            PriceMatrixPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                         `json:"maximum_amount,nullable"`
	Minimum            PriceMatrixPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                         `json:"minimum_amount,nullable"`
	JSON               priceMatrixPriceJSON
}

// priceMatrixPriceJSON contains the JSON metadata for the struct
// [PriceMatrixPrice]
type priceMatrixPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	MatrixConfig       apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceMatrixPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceMatrixPrice) implementsPrice() {}

type PriceMatrixPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceMatrixPriceBillableMetricJSON
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

type PriceMatrixPriceCadence string

const (
	PriceMatrixPriceCadenceOneTime   PriceMatrixPriceCadence = "one_time"
	PriceMatrixPriceCadenceMonthly   PriceMatrixPriceCadence = "monthly"
	PriceMatrixPriceCadenceQuarterly PriceMatrixPriceCadence = "quarterly"
	PriceMatrixPriceCadenceAnnual    PriceMatrixPriceCadence = "annual"
)

type PriceMatrixPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceMatrixPriceItemJSON
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

type PriceMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount string `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions []string `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues []PriceMatrixPriceMatrixConfigMatrixValue `json:"matrix_values,required"`
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor float64 `json:"scaling_factor,nullable"`
	JSON          priceMatrixPriceMatrixConfigJSON
}

// priceMatrixPriceMatrixConfigJSON contains the JSON metadata for the struct
// [PriceMatrixPriceMatrixConfig]
type priceMatrixPriceMatrixConfigJSON struct {
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	ScalingFactor     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceMatrixPriceMatrixConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string `json:"unit_amount,required"`
	// Optional multiplier to scale rated quantities by
	ScalingFactor float64 `json:"scaling_factor,nullable"`
	JSON          priceMatrixPriceMatrixConfigMatrixValueJSON
}

// priceMatrixPriceMatrixConfigMatrixValueJSON contains the JSON metadata for the
// struct [PriceMatrixPriceMatrixConfigMatrixValue]
type priceMatrixPriceMatrixConfigMatrixValueJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	ScalingFactor   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PriceMatrixPriceMatrixConfigMatrixValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceMatrixPriceModelType string

const (
	PriceMatrixPriceModelTypeMatrix PriceMatrixPriceModelType = "matrix"
)

type PriceMatrixPricePriceType string

const (
	PriceMatrixPricePriceTypeUsagePrice PriceMatrixPricePriceType = "usage_price"
	PriceMatrixPricePriceTypeFixedPrice PriceMatrixPricePriceType = "fixed_price"
)

type PriceMatrixPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceMatrixPriceMaximumJSON
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

type PriceMatrixPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceMatrixPriceMinimumJSON
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

type PriceTieredPrice struct {
	ID                 string                         `json:"id,required"`
	BillableMetric     PriceTieredPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceTieredPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                      `json:"created_at,required" format:"date-time"`
	Currency           string                         `json:"currency,required"`
	ExternalPriceID    string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                        `json:"fixed_price_quantity,required,nullable"`
	Item               PriceTieredPriceItem           `json:"item,required"`
	ModelType          PriceTieredPriceModelType      `json:"model_type,required"`
	Name               string                         `json:"name,required"`
	PlanPhaseOrder     int64                          `json:"plan_phase_order,required,nullable"`
	PriceType          PriceTieredPricePriceType      `json:"price_type,required"`
	TieredConfig       PriceTieredPriceTieredConfig   `json:"tiered_config,required"`
	Discount           InvoiceDiscount                `json:"discount,nullable"`
	Maximum            PriceTieredPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                         `json:"maximum_amount,nullable"`
	Minimum            PriceTieredPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                         `json:"minimum_amount,nullable"`
	JSON               priceTieredPriceJSON
}

// priceTieredPriceJSON contains the JSON metadata for the struct
// [PriceTieredPrice]
type priceTieredPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	TieredConfig       apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceTieredPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceTieredPrice) implementsPrice() {}

type PriceTieredPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceTieredPriceBillableMetricJSON
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

type PriceTieredPriceCadence string

const (
	PriceTieredPriceCadenceOneTime   PriceTieredPriceCadence = "one_time"
	PriceTieredPriceCadenceMonthly   PriceTieredPriceCadence = "monthly"
	PriceTieredPriceCadenceQuarterly PriceTieredPriceCadence = "quarterly"
	PriceTieredPriceCadenceAnnual    PriceTieredPriceCadence = "annual"
)

type PriceTieredPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceTieredPriceItemJSON
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

type PriceTieredPriceModelType string

const (
	PriceTieredPriceModelTypeTiered PriceTieredPriceModelType = "tiered"
)

type PriceTieredPricePriceType string

const (
	PriceTieredPricePriceTypeUsagePrice PriceTieredPricePriceType = "usage_price"
	PriceTieredPricePriceTypeFixedPrice PriceTieredPricePriceType = "fixed_price"
)

type PriceTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers []PriceTieredPriceTieredConfigTier `json:"tiers,required"`
	JSON  priceTieredPriceTieredConfigJSON
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

type PriceTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit float64 `json:"first_unit,required"`
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit float64 `json:"last_unit,nullable"`
	JSON     priceTieredPriceTieredConfigTierJSON
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

type PriceTieredPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceTieredPriceMaximumJSON
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

type PriceTieredPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceTieredPriceMinimumJSON
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

type PriceTieredBpsPrice struct {
	ID                 string                             `json:"id,required"`
	BillableMetric     PriceTieredBpsPriceBillableMetric  `json:"billable_metric,required,nullable"`
	Cadence            PriceTieredBpsPriceCadence         `json:"cadence,required"`
	CreatedAt          time.Time                          `json:"created_at,required" format:"date-time"`
	Currency           string                             `json:"currency,required"`
	ExternalPriceID    string                             `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                            `json:"fixed_price_quantity,required,nullable"`
	Item               PriceTieredBpsPriceItem            `json:"item,required"`
	ModelType          PriceTieredBpsPriceModelType       `json:"model_type,required"`
	Name               string                             `json:"name,required"`
	PlanPhaseOrder     int64                              `json:"plan_phase_order,required,nullable"`
	PriceType          PriceTieredBpsPricePriceType       `json:"price_type,required"`
	TieredBpsConfig    PriceTieredBpsPriceTieredBpsConfig `json:"tiered_bps_config,required"`
	Discount           InvoiceDiscount                    `json:"discount,nullable"`
	Maximum            PriceTieredBpsPriceMaximum         `json:"maximum,nullable"`
	MaximumAmount      string                             `json:"maximum_amount,nullable"`
	Minimum            PriceTieredBpsPriceMinimum         `json:"minimum,nullable"`
	MinimumAmount      string                             `json:"minimum_amount,nullable"`
	JSON               priceTieredBpsPriceJSON
}

// priceTieredBpsPriceJSON contains the JSON metadata for the struct
// [PriceTieredBpsPrice]
type priceTieredBpsPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	TieredBpsConfig    apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceTieredBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceTieredBpsPrice) implementsPrice() {}

type PriceTieredBpsPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceTieredBpsPriceBillableMetricJSON
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

type PriceTieredBpsPriceCadence string

const (
	PriceTieredBpsPriceCadenceOneTime   PriceTieredBpsPriceCadence = "one_time"
	PriceTieredBpsPriceCadenceMonthly   PriceTieredBpsPriceCadence = "monthly"
	PriceTieredBpsPriceCadenceQuarterly PriceTieredBpsPriceCadence = "quarterly"
	PriceTieredBpsPriceCadenceAnnual    PriceTieredBpsPriceCadence = "annual"
)

type PriceTieredBpsPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceTieredBpsPriceItemJSON
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

type PriceTieredBpsPriceModelType string

const (
	PriceTieredBpsPriceModelTypeTieredBps PriceTieredBpsPriceModelType = "tiered_bps"
)

type PriceTieredBpsPricePriceType string

const (
	PriceTieredBpsPricePriceTypeUsagePrice PriceTieredBpsPricePriceType = "usage_price"
	PriceTieredBpsPricePriceTypeFixedPrice PriceTieredBpsPricePriceType = "fixed_price"
)

type PriceTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers []PriceTieredBpsPriceTieredBpsConfigTier `json:"tiers,required"`
	JSON  priceTieredBpsPriceTieredBpsConfigJSON
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

type PriceTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps float64 `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount string `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount string `json:"maximum_amount,nullable"`
	// Per unit maximum to charge
	PerUnitMaximum string `json:"per_unit_maximum,nullable"`
	JSON           priceTieredBpsPriceTieredBpsConfigTierJSON
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

type PriceTieredBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceTieredBpsPriceMaximumJSON
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

type PriceTieredBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceTieredBpsPriceMinimumJSON
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

type PriceBpsPrice struct {
	ID                 string                      `json:"id,required"`
	BillableMetric     PriceBpsPriceBillableMetric `json:"billable_metric,required,nullable"`
	BpsConfig          PriceBpsPriceBpsConfig      `json:"bps_config,required"`
	Cadence            PriceBpsPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                   `json:"created_at,required" format:"date-time"`
	Currency           string                      `json:"currency,required"`
	ExternalPriceID    string                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                     `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBpsPriceItem           `json:"item,required"`
	ModelType          PriceBpsPriceModelType      `json:"model_type,required"`
	Name               string                      `json:"name,required"`
	PlanPhaseOrder     int64                       `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBpsPricePriceType      `json:"price_type,required"`
	Discount           InvoiceDiscount             `json:"discount,nullable"`
	Maximum            PriceBpsPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                      `json:"maximum_amount,nullable"`
	Minimum            PriceBpsPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                      `json:"minimum_amount,nullable"`
	JSON               priceBpsPriceJSON
}

// priceBpsPriceJSON contains the JSON metadata for the struct [PriceBpsPrice]
type priceBpsPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	BpsConfig          apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceBpsPrice) implementsPrice() {}

type PriceBpsPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceBpsPriceBillableMetricJSON
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

type PriceBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps float64 `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum string `json:"per_unit_maximum,nullable"`
	JSON           priceBpsPriceBpsConfigJSON
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

type PriceBpsPriceCadence string

const (
	PriceBpsPriceCadenceOneTime   PriceBpsPriceCadence = "one_time"
	PriceBpsPriceCadenceMonthly   PriceBpsPriceCadence = "monthly"
	PriceBpsPriceCadenceQuarterly PriceBpsPriceCadence = "quarterly"
	PriceBpsPriceCadenceAnnual    PriceBpsPriceCadence = "annual"
)

type PriceBpsPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceBpsPriceItemJSON
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

type PriceBpsPriceModelType string

const (
	PriceBpsPriceModelTypeBps PriceBpsPriceModelType = "bps"
)

type PriceBpsPricePriceType string

const (
	PriceBpsPricePriceTypeUsagePrice PriceBpsPricePriceType = "usage_price"
	PriceBpsPricePriceTypeFixedPrice PriceBpsPricePriceType = "fixed_price"
)

type PriceBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceBpsPriceMaximumJSON
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

type PriceBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceBpsPriceMinimumJSON
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

type PriceBulkBpsPrice struct {
	ID                 string                          `json:"id,required"`
	BillableMetric     PriceBulkBpsPriceBillableMetric `json:"billable_metric,required,nullable"`
	BulkBpsConfig      PriceBulkBpsPriceBulkBpsConfig  `json:"bulk_bps_config,required"`
	Cadence            PriceBulkBpsPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                       `json:"created_at,required" format:"date-time"`
	Currency           string                          `json:"currency,required"`
	ExternalPriceID    string                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                         `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBulkBpsPriceItem           `json:"item,required"`
	ModelType          PriceBulkBpsPriceModelType      `json:"model_type,required"`
	Name               string                          `json:"name,required"`
	PlanPhaseOrder     int64                           `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBulkBpsPricePriceType      `json:"price_type,required"`
	Discount           InvoiceDiscount                 `json:"discount,nullable"`
	Maximum            PriceBulkBpsPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                          `json:"maximum_amount,nullable"`
	Minimum            PriceBulkBpsPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                          `json:"minimum_amount,nullable"`
	JSON               priceBulkBpsPriceJSON
}

// priceBulkBpsPriceJSON contains the JSON metadata for the struct
// [PriceBulkBpsPrice]
type priceBulkBpsPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	BulkBpsConfig      apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceBulkBpsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceBulkBpsPrice) implementsPrice() {}

type PriceBulkBpsPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceBulkBpsPriceBillableMetricJSON
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

type PriceBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers []PriceBulkBpsPriceBulkBpsConfigTier `json:"tiers,required"`
	JSON  priceBulkBpsPriceBulkBpsConfigJSON
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

type PriceBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps float64 `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount string `json:"maximum_amount,nullable"`
	// The maximum amount to charge for any one event
	PerUnitMaximum string `json:"per_unit_maximum,nullable"`
	JSON           priceBulkBpsPriceBulkBpsConfigTierJSON
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

type PriceBulkBpsPriceCadence string

const (
	PriceBulkBpsPriceCadenceOneTime   PriceBulkBpsPriceCadence = "one_time"
	PriceBulkBpsPriceCadenceMonthly   PriceBulkBpsPriceCadence = "monthly"
	PriceBulkBpsPriceCadenceQuarterly PriceBulkBpsPriceCadence = "quarterly"
	PriceBulkBpsPriceCadenceAnnual    PriceBulkBpsPriceCadence = "annual"
)

type PriceBulkBpsPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceBulkBpsPriceItemJSON
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

type PriceBulkBpsPriceModelType string

const (
	PriceBulkBpsPriceModelTypeBulkBps PriceBulkBpsPriceModelType = "bulk_bps"
)

type PriceBulkBpsPricePriceType string

const (
	PriceBulkBpsPricePriceTypeUsagePrice PriceBulkBpsPricePriceType = "usage_price"
	PriceBulkBpsPricePriceTypeFixedPrice PriceBulkBpsPricePriceType = "fixed_price"
)

type PriceBulkBpsPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceBulkBpsPriceMaximumJSON
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

type PriceBulkBpsPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceBulkBpsPriceMinimumJSON
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

type PriceBulkPrice struct {
	ID                 string                       `json:"id,required"`
	BillableMetric     PriceBulkPriceBillableMetric `json:"billable_metric,required,nullable"`
	BulkConfig         PriceBulkPriceBulkConfig     `json:"bulk_config,required"`
	Cadence            PriceBulkPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                    `json:"created_at,required" format:"date-time"`
	Currency           string                       `json:"currency,required"`
	ExternalPriceID    string                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                      `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBulkPriceItem           `json:"item,required"`
	ModelType          PriceBulkPriceModelType      `json:"model_type,required"`
	Name               string                       `json:"name,required"`
	PlanPhaseOrder     int64                        `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBulkPricePriceType      `json:"price_type,required"`
	Discount           InvoiceDiscount              `json:"discount,nullable"`
	Maximum            PriceBulkPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount      string                       `json:"maximum_amount,nullable"`
	Minimum            PriceBulkPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount      string                       `json:"minimum_amount,nullable"`
	JSON               priceBulkPriceJSON
}

// priceBulkPriceJSON contains the JSON metadata for the struct [PriceBulkPrice]
type priceBulkPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	BulkConfig         apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	Discount           apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PriceBulkPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceBulkPrice) implementsPrice() {}

type PriceBulkPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceBulkPriceBillableMetricJSON
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

type PriceBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers []PriceBulkPriceBulkConfigTier `json:"tiers,required"`
	JSON  priceBulkPriceBulkConfigJSON
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

type PriceBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount string `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits float64 `json:"maximum_units,nullable"`
	JSON         priceBulkPriceBulkConfigTierJSON
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

type PriceBulkPriceCadence string

const (
	PriceBulkPriceCadenceOneTime   PriceBulkPriceCadence = "one_time"
	PriceBulkPriceCadenceMonthly   PriceBulkPriceCadence = "monthly"
	PriceBulkPriceCadenceQuarterly PriceBulkPriceCadence = "quarterly"
	PriceBulkPriceCadenceAnnual    PriceBulkPriceCadence = "annual"
)

type PriceBulkPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceBulkPriceItemJSON
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

type PriceBulkPriceModelType string

const (
	PriceBulkPriceModelTypeBulk PriceBulkPriceModelType = "bulk"
)

type PriceBulkPricePriceType string

const (
	PriceBulkPricePriceTypeUsagePrice PriceBulkPricePriceType = "usage_price"
	PriceBulkPricePriceTypeFixedPrice PriceBulkPricePriceType = "fixed_price"
)

type PriceBulkPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceBulkPriceMaximumJSON
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

type PriceBulkPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceBulkPriceMinimumJSON
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

type PriceTestRatingFunctionPrice struct {
	ID                       string                                     `json:"id,required"`
	BillableMetric           PriceTestRatingFunctionPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                  PriceTestRatingFunctionPriceCadence        `json:"cadence,required"`
	CreatedAt                time.Time                                  `json:"created_at,required" format:"date-time"`
	Currency                 string                                     `json:"currency,required"`
	ExternalPriceID          string                                     `json:"external_price_id,required,nullable"`
	FixedPriceQuantity       float64                                    `json:"fixed_price_quantity,required,nullable"`
	Item                     PriceTestRatingFunctionPriceItem           `json:"item,required"`
	ModelType                PriceTestRatingFunctionPriceModelType      `json:"model_type,required"`
	Name                     string                                     `json:"name,required"`
	PlanPhaseOrder           int64                                      `json:"plan_phase_order,required,nullable"`
	PriceType                PriceTestRatingFunctionPricePriceType      `json:"price_type,required"`
	TestRatingFunctionConfig map[string]interface{}                     `json:"test_rating_function_config,required"`
	Discount                 InvoiceDiscount                            `json:"discount,nullable"`
	Maximum                  PriceTestRatingFunctionPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount            string                                     `json:"maximum_amount,nullable"`
	Minimum                  PriceTestRatingFunctionPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount            string                                     `json:"minimum_amount,nullable"`
	JSON                     priceTestRatingFunctionPriceJSON
}

// priceTestRatingFunctionPriceJSON contains the JSON metadata for the struct
// [PriceTestRatingFunctionPrice]
type priceTestRatingFunctionPriceJSON struct {
	ID                       apijson.Field
	BillableMetric           apijson.Field
	Cadence                  apijson.Field
	CreatedAt                apijson.Field
	Currency                 apijson.Field
	ExternalPriceID          apijson.Field
	FixedPriceQuantity       apijson.Field
	Item                     apijson.Field
	ModelType                apijson.Field
	Name                     apijson.Field
	PlanPhaseOrder           apijson.Field
	PriceType                apijson.Field
	TestRatingFunctionConfig apijson.Field
	Discount                 apijson.Field
	Maximum                  apijson.Field
	MaximumAmount            apijson.Field
	Minimum                  apijson.Field
	MinimumAmount            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *PriceTestRatingFunctionPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceTestRatingFunctionPrice) implementsPrice() {}

type PriceTestRatingFunctionPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceTestRatingFunctionPriceBillableMetricJSON
}

// priceTestRatingFunctionPriceBillableMetricJSON contains the JSON metadata for
// the struct [PriceTestRatingFunctionPriceBillableMetric]
type priceTestRatingFunctionPriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTestRatingFunctionPriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceTestRatingFunctionPriceCadence string

const (
	PriceTestRatingFunctionPriceCadenceOneTime   PriceTestRatingFunctionPriceCadence = "one_time"
	PriceTestRatingFunctionPriceCadenceMonthly   PriceTestRatingFunctionPriceCadence = "monthly"
	PriceTestRatingFunctionPriceCadenceQuarterly PriceTestRatingFunctionPriceCadence = "quarterly"
	PriceTestRatingFunctionPriceCadenceAnnual    PriceTestRatingFunctionPriceCadence = "annual"
)

type PriceTestRatingFunctionPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceTestRatingFunctionPriceItemJSON
}

// priceTestRatingFunctionPriceItemJSON contains the JSON metadata for the struct
// [PriceTestRatingFunctionPriceItem]
type priceTestRatingFunctionPriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceTestRatingFunctionPriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceTestRatingFunctionPriceModelType string

const (
	PriceTestRatingFunctionPriceModelTypeTestRatingFunction PriceTestRatingFunctionPriceModelType = "test_rating_function"
)

type PriceTestRatingFunctionPricePriceType string

const (
	PriceTestRatingFunctionPricePriceTypeUsagePrice PriceTestRatingFunctionPricePriceType = "usage_price"
	PriceTestRatingFunctionPricePriceTypeFixedPrice PriceTestRatingFunctionPricePriceType = "fixed_price"
)

type PriceTestRatingFunctionPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceTestRatingFunctionPriceMaximumJSON
}

// priceTestRatingFunctionPriceMaximumJSON contains the JSON metadata for the
// struct [PriceTestRatingFunctionPriceMaximum]
type priceTestRatingFunctionPriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTestRatingFunctionPriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceTestRatingFunctionPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceTestRatingFunctionPriceMinimumJSON
}

// priceTestRatingFunctionPriceMinimumJSON contains the JSON metadata for the
// struct [PriceTestRatingFunctionPriceMinimum]
type priceTestRatingFunctionPriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceTestRatingFunctionPriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceFivetranExamplePrice struct {
	ID                    string                                  `json:"id,required"`
	BillableMetric        PriceFivetranExamplePriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence               PriceFivetranExamplePriceCadence        `json:"cadence,required"`
	CreatedAt             time.Time                               `json:"created_at,required" format:"date-time"`
	Currency              string                                  `json:"currency,required"`
	ExternalPriceID       string                                  `json:"external_price_id,required,nullable"`
	FivetranExampleConfig map[string]interface{}                  `json:"fivetran_example_config,required"`
	FixedPriceQuantity    float64                                 `json:"fixed_price_quantity,required,nullable"`
	Item                  PriceFivetranExamplePriceItem           `json:"item,required"`
	ModelType             PriceFivetranExamplePriceModelType      `json:"model_type,required"`
	Name                  string                                  `json:"name,required"`
	PlanPhaseOrder        int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType             PriceFivetranExamplePricePriceType      `json:"price_type,required"`
	Discount              InvoiceDiscount                         `json:"discount,nullable"`
	Maximum               PriceFivetranExamplePriceMaximum        `json:"maximum,nullable"`
	MaximumAmount         string                                  `json:"maximum_amount,nullable"`
	Minimum               PriceFivetranExamplePriceMinimum        `json:"minimum,nullable"`
	MinimumAmount         string                                  `json:"minimum_amount,nullable"`
	JSON                  priceFivetranExamplePriceJSON
}

// priceFivetranExamplePriceJSON contains the JSON metadata for the struct
// [PriceFivetranExamplePrice]
type priceFivetranExamplePriceJSON struct {
	ID                    apijson.Field
	BillableMetric        apijson.Field
	Cadence               apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	ExternalPriceID       apijson.Field
	FivetranExampleConfig apijson.Field
	FixedPriceQuantity    apijson.Field
	Item                  apijson.Field
	ModelType             apijson.Field
	Name                  apijson.Field
	PlanPhaseOrder        apijson.Field
	PriceType             apijson.Field
	Discount              apijson.Field
	Maximum               apijson.Field
	MaximumAmount         apijson.Field
	Minimum               apijson.Field
	MinimumAmount         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PriceFivetranExamplePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceFivetranExamplePrice) implementsPrice() {}

type PriceFivetranExamplePriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceFivetranExamplePriceBillableMetricJSON
}

// priceFivetranExamplePriceBillableMetricJSON contains the JSON metadata for the
// struct [PriceFivetranExamplePriceBillableMetric]
type priceFivetranExamplePriceBillableMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceFivetranExamplePriceBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceFivetranExamplePriceCadence string

const (
	PriceFivetranExamplePriceCadenceOneTime   PriceFivetranExamplePriceCadence = "one_time"
	PriceFivetranExamplePriceCadenceMonthly   PriceFivetranExamplePriceCadence = "monthly"
	PriceFivetranExamplePriceCadenceQuarterly PriceFivetranExamplePriceCadence = "quarterly"
	PriceFivetranExamplePriceCadenceAnnual    PriceFivetranExamplePriceCadence = "annual"
)

type PriceFivetranExamplePriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceFivetranExamplePriceItemJSON
}

// priceFivetranExamplePriceItemJSON contains the JSON metadata for the struct
// [PriceFivetranExamplePriceItem]
type priceFivetranExamplePriceItemJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PriceFivetranExamplePriceItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceFivetranExamplePriceModelType string

const (
	PriceFivetranExamplePriceModelTypeFivetranExample PriceFivetranExamplePriceModelType = "fivetran_example"
)

type PriceFivetranExamplePricePriceType string

const (
	PriceFivetranExamplePricePriceTypeUsagePrice PriceFivetranExamplePricePriceType = "usage_price"
	PriceFivetranExamplePricePriceTypeFixedPrice PriceFivetranExamplePricePriceType = "fixed_price"
)

type PriceFivetranExamplePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceFivetranExamplePriceMaximumJSON
}

// priceFivetranExamplePriceMaximumJSON contains the JSON metadata for the struct
// [PriceFivetranExamplePriceMaximum]
type priceFivetranExamplePriceMaximumJSON struct {
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceFivetranExamplePriceMaximum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceFivetranExamplePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceFivetranExamplePriceMinimumJSON
}

// priceFivetranExamplePriceMinimumJSON contains the JSON metadata for the struct
// [PriceFivetranExamplePriceMinimum]
type priceFivetranExamplePriceMinimumJSON struct {
	AppliesToPriceIDs apijson.Field
	MinimumAmount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PriceFivetranExamplePriceMinimum) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PriceThresholdTotalAmountPrice struct {
	ID                         string                                       `json:"id,required"`
	BillableMetric             PriceThresholdTotalAmountPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                    PriceThresholdTotalAmountPriceCadence        `json:"cadence,required"`
	CreatedAt                  time.Time                                    `json:"created_at,required" format:"date-time"`
	Currency                   string                                       `json:"currency,required"`
	ExternalPriceID            string                                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity         float64                                      `json:"fixed_price_quantity,required,nullable"`
	Item                       PriceThresholdTotalAmountPriceItem           `json:"item,required"`
	ModelType                  PriceThresholdTotalAmountPriceModelType      `json:"model_type,required"`
	Name                       string                                       `json:"name,required"`
	PlanPhaseOrder             int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType                  PriceThresholdTotalAmountPricePriceType      `json:"price_type,required"`
	ThresholdTotalAmountConfig map[string]interface{}                       `json:"threshold_total_amount_config,required"`
	Discount                   InvoiceDiscount                              `json:"discount,nullable"`
	Maximum                    PriceThresholdTotalAmountPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount              string                                       `json:"maximum_amount,nullable"`
	Minimum                    PriceThresholdTotalAmountPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount              string                                       `json:"minimum_amount,nullable"`
	JSON                       priceThresholdTotalAmountPriceJSON
}

// priceThresholdTotalAmountPriceJSON contains the JSON metadata for the struct
// [PriceThresholdTotalAmountPrice]
type priceThresholdTotalAmountPriceJSON struct {
	ID                         apijson.Field
	BillableMetric             apijson.Field
	Cadence                    apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	ExternalPriceID            apijson.Field
	FixedPriceQuantity         apijson.Field
	Item                       apijson.Field
	ModelType                  apijson.Field
	Name                       apijson.Field
	PlanPhaseOrder             apijson.Field
	PriceType                  apijson.Field
	ThresholdTotalAmountConfig apijson.Field
	Discount                   apijson.Field
	Maximum                    apijson.Field
	MaximumAmount              apijson.Field
	Minimum                    apijson.Field
	MinimumAmount              apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *PriceThresholdTotalAmountPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceThresholdTotalAmountPrice) implementsPrice() {}

type PriceThresholdTotalAmountPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceThresholdTotalAmountPriceBillableMetricJSON
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

type PriceThresholdTotalAmountPriceCadence string

const (
	PriceThresholdTotalAmountPriceCadenceOneTime   PriceThresholdTotalAmountPriceCadence = "one_time"
	PriceThresholdTotalAmountPriceCadenceMonthly   PriceThresholdTotalAmountPriceCadence = "monthly"
	PriceThresholdTotalAmountPriceCadenceQuarterly PriceThresholdTotalAmountPriceCadence = "quarterly"
	PriceThresholdTotalAmountPriceCadenceAnnual    PriceThresholdTotalAmountPriceCadence = "annual"
)

type PriceThresholdTotalAmountPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceThresholdTotalAmountPriceItemJSON
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

type PriceThresholdTotalAmountPriceModelType string

const (
	PriceThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

type PriceThresholdTotalAmountPricePriceType string

const (
	PriceThresholdTotalAmountPricePriceTypeUsagePrice PriceThresholdTotalAmountPricePriceType = "usage_price"
	PriceThresholdTotalAmountPricePriceTypeFixedPrice PriceThresholdTotalAmountPricePriceType = "fixed_price"
)

type PriceThresholdTotalAmountPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceThresholdTotalAmountPriceMaximumJSON
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

type PriceThresholdTotalAmountPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceThresholdTotalAmountPriceMinimumJSON
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

type PriceTieredPackagePrice struct {
	ID                  string                                `json:"id,required"`
	BillableMetric      PriceTieredPackagePriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence             PriceTieredPackagePriceCadence        `json:"cadence,required"`
	CreatedAt           time.Time                             `json:"created_at,required" format:"date-time"`
	Currency            string                                `json:"currency,required"`
	ExternalPriceID     string                                `json:"external_price_id,required,nullable"`
	FixedPriceQuantity  float64                               `json:"fixed_price_quantity,required,nullable"`
	Item                PriceTieredPackagePriceItem           `json:"item,required"`
	ModelType           PriceTieredPackagePriceModelType      `json:"model_type,required"`
	Name                string                                `json:"name,required"`
	PlanPhaseOrder      int64                                 `json:"plan_phase_order,required,nullable"`
	PriceType           PriceTieredPackagePricePriceType      `json:"price_type,required"`
	TieredPackageConfig map[string]interface{}                `json:"tiered_package_config,required"`
	Discount            InvoiceDiscount                       `json:"discount,nullable"`
	Maximum             PriceTieredPackagePriceMaximum        `json:"maximum,nullable"`
	MaximumAmount       string                                `json:"maximum_amount,nullable"`
	Minimum             PriceTieredPackagePriceMinimum        `json:"minimum,nullable"`
	MinimumAmount       string                                `json:"minimum_amount,nullable"`
	JSON                priceTieredPackagePriceJSON
}

// priceTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceTieredPackagePrice]
type priceTieredPackagePriceJSON struct {
	ID                  apijson.Field
	BillableMetric      apijson.Field
	Cadence             apijson.Field
	CreatedAt           apijson.Field
	Currency            apijson.Field
	ExternalPriceID     apijson.Field
	FixedPriceQuantity  apijson.Field
	Item                apijson.Field
	ModelType           apijson.Field
	Name                apijson.Field
	PlanPhaseOrder      apijson.Field
	PriceType           apijson.Field
	TieredPackageConfig apijson.Field
	Discount            apijson.Field
	Maximum             apijson.Field
	MaximumAmount       apijson.Field
	Minimum             apijson.Field
	MinimumAmount       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PriceTieredPackagePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceTieredPackagePrice) implementsPrice() {}

type PriceTieredPackagePriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceTieredPackagePriceBillableMetricJSON
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

type PriceTieredPackagePriceCadence string

const (
	PriceTieredPackagePriceCadenceOneTime   PriceTieredPackagePriceCadence = "one_time"
	PriceTieredPackagePriceCadenceMonthly   PriceTieredPackagePriceCadence = "monthly"
	PriceTieredPackagePriceCadenceQuarterly PriceTieredPackagePriceCadence = "quarterly"
	PriceTieredPackagePriceCadenceAnnual    PriceTieredPackagePriceCadence = "annual"
)

type PriceTieredPackagePriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceTieredPackagePriceItemJSON
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

type PriceTieredPackagePriceModelType string

const (
	PriceTieredPackagePriceModelTypeTieredPackage PriceTieredPackagePriceModelType = "tiered_package"
)

type PriceTieredPackagePricePriceType string

const (
	PriceTieredPackagePricePriceTypeUsagePrice PriceTieredPackagePricePriceType = "usage_price"
	PriceTieredPackagePricePriceTypeFixedPrice PriceTieredPackagePricePriceType = "fixed_price"
)

type PriceTieredPackagePriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceTieredPackagePriceMaximumJSON
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

type PriceTieredPackagePriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceTieredPackagePriceMinimumJSON
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

type PriceTieredWithMinimumPrice struct {
	ID                      string                                    `json:"id,required"`
	BillableMetric          PriceTieredWithMinimumPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                 PriceTieredWithMinimumPriceCadence        `json:"cadence,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	Currency                string                                    `json:"currency,required"`
	ExternalPriceID         string                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity      float64                                   `json:"fixed_price_quantity,required,nullable"`
	Item                    PriceTieredWithMinimumPriceItem           `json:"item,required"`
	ModelType               PriceTieredWithMinimumPriceModelType      `json:"model_type,required"`
	Name                    string                                    `json:"name,required"`
	PlanPhaseOrder          int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType               PriceTieredWithMinimumPricePriceType      `json:"price_type,required"`
	TieredWithMinimumConfig map[string]interface{}                    `json:"tiered_with_minimum_config,required"`
	Discount                InvoiceDiscount                           `json:"discount,nullable"`
	Maximum                 PriceTieredWithMinimumPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount           string                                    `json:"maximum_amount,nullable"`
	Minimum                 PriceTieredWithMinimumPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount           string                                    `json:"minimum_amount,nullable"`
	JSON                    priceTieredWithMinimumPriceJSON
}

// priceTieredWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPrice]
type priceTieredWithMinimumPriceJSON struct {
	ID                      apijson.Field
	BillableMetric          apijson.Field
	Cadence                 apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	ExternalPriceID         apijson.Field
	FixedPriceQuantity      apijson.Field
	Item                    apijson.Field
	ModelType               apijson.Field
	Name                    apijson.Field
	PlanPhaseOrder          apijson.Field
	PriceType               apijson.Field
	TieredWithMinimumConfig apijson.Field
	Discount                apijson.Field
	Maximum                 apijson.Field
	MaximumAmount           apijson.Field
	Minimum                 apijson.Field
	MinimumAmount           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PriceTieredWithMinimumPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PriceTieredWithMinimumPrice) implementsPrice() {}

type PriceTieredWithMinimumPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON priceTieredWithMinimumPriceBillableMetricJSON
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

type PriceTieredWithMinimumPriceCadence string

const (
	PriceTieredWithMinimumPriceCadenceOneTime   PriceTieredWithMinimumPriceCadence = "one_time"
	PriceTieredWithMinimumPriceCadenceMonthly   PriceTieredWithMinimumPriceCadence = "monthly"
	PriceTieredWithMinimumPriceCadenceQuarterly PriceTieredWithMinimumPriceCadence = "quarterly"
	PriceTieredWithMinimumPriceCadenceAnnual    PriceTieredWithMinimumPriceCadence = "annual"
)

type PriceTieredWithMinimumPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON priceTieredWithMinimumPriceItemJSON
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

type PriceTieredWithMinimumPriceModelType string

const (
	PriceTieredWithMinimumPriceModelTypeTieredWithMinimum PriceTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

type PriceTieredWithMinimumPricePriceType string

const (
	PriceTieredWithMinimumPricePriceTypeUsagePrice PriceTieredWithMinimumPricePriceType = "usage_price"
	PriceTieredWithMinimumPricePriceTypeFixedPrice PriceTieredWithMinimumPricePriceType = "fixed_price"
)

type PriceTieredWithMinimumPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          priceTieredWithMinimumPriceMaximumJSON
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

type PriceTieredWithMinimumPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          priceTieredWithMinimumPriceMinimumJSON
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

type PricePackageWithAllocationPrice struct {
	ID                          string                                        `json:"id,required"`
	BillableMetric              PricePackageWithAllocationPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                     PricePackageWithAllocationPriceCadence        `json:"cadence,required"`
	CreatedAt                   time.Time                                     `json:"created_at,required" format:"date-time"`
	Currency                    string                                        `json:"currency,required"`
	ExternalPriceID             string                                        `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                       `json:"fixed_price_quantity,required,nullable"`
	Item                        PricePackageWithAllocationPriceItem           `json:"item,required"`
	ModelType                   PricePackageWithAllocationPriceModelType      `json:"model_type,required"`
	Name                        string                                        `json:"name,required"`
	PackageWithAllocationConfig map[string]interface{}                        `json:"package_with_allocation_config,required"`
	PlanPhaseOrder              int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                   PricePackageWithAllocationPricePriceType      `json:"price_type,required"`
	Discount                    InvoiceDiscount                               `json:"discount,nullable"`
	Maximum                     PricePackageWithAllocationPriceMaximum        `json:"maximum,nullable"`
	MaximumAmount               string                                        `json:"maximum_amount,nullable"`
	Minimum                     PricePackageWithAllocationPriceMinimum        `json:"minimum,nullable"`
	MinimumAmount               string                                        `json:"minimum_amount,nullable"`
	JSON                        pricePackageWithAllocationPriceJSON
}

// pricePackageWithAllocationPriceJSON contains the JSON metadata for the struct
// [PricePackageWithAllocationPrice]
type pricePackageWithAllocationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	Cadence                     apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	Item                        apijson.Field
	ModelType                   apijson.Field
	Name                        apijson.Field
	PackageWithAllocationConfig apijson.Field
	PlanPhaseOrder              apijson.Field
	PriceType                   apijson.Field
	Discount                    apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
	Minimum                     apijson.Field
	MinimumAmount               apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PricePackageWithAllocationPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PricePackageWithAllocationPrice) implementsPrice() {}

type PricePackageWithAllocationPriceBillableMetric struct {
	ID   string `json:"id,required"`
	JSON pricePackageWithAllocationPriceBillableMetricJSON
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

type PricePackageWithAllocationPriceCadence string

const (
	PricePackageWithAllocationPriceCadenceOneTime   PricePackageWithAllocationPriceCadence = "one_time"
	PricePackageWithAllocationPriceCadenceMonthly   PricePackageWithAllocationPriceCadence = "monthly"
	PricePackageWithAllocationPriceCadenceQuarterly PricePackageWithAllocationPriceCadence = "quarterly"
	PricePackageWithAllocationPriceCadenceAnnual    PricePackageWithAllocationPriceCadence = "annual"
)

type PricePackageWithAllocationPriceItem struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	JSON pricePackageWithAllocationPriceItemJSON
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

type PricePackageWithAllocationPriceModelType string

const (
	PricePackageWithAllocationPriceModelTypePackageWithAllocation PricePackageWithAllocationPriceModelType = "package_with_allocation"
)

type PricePackageWithAllocationPricePriceType string

const (
	PricePackageWithAllocationPricePriceTypeUsagePrice PricePackageWithAllocationPricePriceType = "usage_price"
	PricePackageWithAllocationPricePriceTypeFixedPrice PricePackageWithAllocationPricePriceType = "fixed_price"
)

type PricePackageWithAllocationPriceMaximum struct {
	// List of price_ids that this maximum amount applies to. For plan/plan phase
	// maximums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Maximum amount applied
	MaximumAmount string `json:"maximum_amount,required"`
	JSON          pricePackageWithAllocationPriceMaximumJSON
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

type PricePackageWithAllocationPriceMinimum struct {
	// List of price_ids that this minimum amount applies to. For plan/plan phase
	// minimums, this can be a subset of prices.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// Minimum amount applied
	MinimumAmount string `json:"minimum_amount,required"`
	JSON          pricePackageWithAllocationPriceMinimumJSON
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

// This interface is a union satisfied by one of the following:
// [PriceNewParamsNewUnitPrice], [PriceNewParamsNewPackagePrice],
// [PriceNewParamsNewMatrixPrice], [PriceNewParamsNewTieredPrice],
// [PriceNewParamsNewTieredBpsPrice], [PriceNewParamsNewBpsPrice],
// [PriceNewParamsNewBulkBpsPrice], [PriceNewParamsNewBulkPrice],
// [PriceNewParamsNewThresholdTotalAmountPrice],
// [PriceNewParamsNewTieredPackagePrice],
// [PriceNewParamsNewTieredWithMinimumPrice],
// [PriceNewParamsNewPackageWithAllocationPrice].
type PriceNewParams interface {
	ImplementsPriceNewParams()
}

type PriceNewParamsNewUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                              `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                               `json:"name,required"`
	UnitConfig param.Field[PriceNewParamsNewUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewUnitPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewUnitPriceCadence string

const (
	PriceNewParamsNewUnitPriceCadenceAnnual    PriceNewParamsNewUnitPriceCadence = "annual"
	PriceNewParamsNewUnitPriceCadenceMonthly   PriceNewParamsNewUnitPriceCadence = "monthly"
	PriceNewParamsNewUnitPriceCadenceQuarterly PriceNewParamsNewUnitPriceCadence = "quarterly"
)

type PriceNewParamsNewUnitPriceModelType string

const (
	PriceNewParamsNewUnitPriceModelTypeUnit PriceNewParamsNewUnitPriceModelType = "unit"
)

type PriceNewParamsNewUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                     `json:"name,required"`
	PackageConfig param.Field[PriceNewParamsNewPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewPackagePriceCadence string

const (
	PriceNewParamsNewPackagePriceCadenceAnnual    PriceNewParamsNewPackagePriceCadence = "annual"
	PriceNewParamsNewPackagePriceCadenceMonthly   PriceNewParamsNewPackagePriceCadence = "monthly"
	PriceNewParamsNewPackagePriceCadenceQuarterly PriceNewParamsNewPackagePriceCadence = "quarterly"
)

type PriceNewParamsNewPackagePriceModelType string

const (
	PriceNewParamsNewPackagePriceModelTypePackage PriceNewParamsNewPackagePriceModelType = "package"
)

type PriceNewParamsNewPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size"`
}

func (r PriceNewParamsNewPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                   `json:"item_id,required"`
	MatrixConfig param.Field[PriceNewParamsNewMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[PriceNewParamsNewMatrixPriceModelType]    `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewMatrixPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewMatrixPriceCadence string

const (
	PriceNewParamsNewMatrixPriceCadenceAnnual    PriceNewParamsNewMatrixPriceCadence = "annual"
	PriceNewParamsNewMatrixPriceCadenceMonthly   PriceNewParamsNewMatrixPriceCadence = "monthly"
	PriceNewParamsNewMatrixPriceCadenceQuarterly PriceNewParamsNewMatrixPriceCadence = "quarterly"
)

type PriceNewParamsNewMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PriceNewParamsNewMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Optional multiplier to scale rated quantities by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewMatrixPriceModelType string

const (
	PriceNewParamsNewMatrixPriceModelTypeMatrix PriceNewParamsNewMatrixPriceModelType = "matrix"
)

type PriceNewParamsNewTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                   `json:"name,required"`
	TieredConfig param.Field[PriceNewParamsNewTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewTieredPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewTieredPriceCadence string

const (
	PriceNewParamsNewTieredPriceCadenceAnnual    PriceNewParamsNewTieredPriceCadence = "annual"
	PriceNewParamsNewTieredPriceCadenceMonthly   PriceNewParamsNewTieredPriceCadence = "monthly"
	PriceNewParamsNewTieredPriceCadenceQuarterly PriceNewParamsNewTieredPriceCadence = "quarterly"
)

type PriceNewParamsNewTieredPriceModelType string

const (
	PriceNewParamsNewTieredPriceModelTypeTiered PriceNewParamsNewTieredPriceModelType = "tiered"
)

type PriceNewParamsNewTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]PriceNewParamsNewTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r PriceNewParamsNewTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewTieredBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                   `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                         `json:"name,required"`
	TieredBpsConfig param.Field[PriceNewParamsNewTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewTieredBpsPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewTieredBpsPriceCadence string

const (
	PriceNewParamsNewTieredBpsPriceCadenceAnnual    PriceNewParamsNewTieredBpsPriceCadence = "annual"
	PriceNewParamsNewTieredBpsPriceCadenceMonthly   PriceNewParamsNewTieredBpsPriceCadence = "monthly"
	PriceNewParamsNewTieredBpsPriceCadenceQuarterly PriceNewParamsNewTieredBpsPriceCadence = "quarterly"
)

type PriceNewParamsNewTieredBpsPriceModelType string

const (
	PriceNewParamsNewTieredBpsPriceModelTypeTieredBps PriceNewParamsNewTieredBpsPriceModelType = "tiered_bps"
)

type PriceNewParamsNewTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]PriceNewParamsNewTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewBpsPrice struct {
	BpsConfig param.Field[PriceNewParamsNewBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                             `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewBpsPrice) ImplementsPriceNewParams() {

}

type PriceNewParamsNewBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PriceNewParamsNewBpsPriceCadence string

const (
	PriceNewParamsNewBpsPriceCadenceAnnual    PriceNewParamsNewBpsPriceCadence = "annual"
	PriceNewParamsNewBpsPriceCadenceMonthly   PriceNewParamsNewBpsPriceCadence = "monthly"
	PriceNewParamsNewBpsPriceCadenceQuarterly PriceNewParamsNewBpsPriceCadence = "quarterly"
)

type PriceNewParamsNewBpsPriceModelType string

const (
	PriceNewParamsNewBpsPriceModelTypeBps PriceNewParamsNewBpsPriceModelType = "bps"
)

type PriceNewParamsNewBulkBpsPrice struct {
	BulkBpsConfig param.Field[PriceNewParamsNewBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewBulkBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                 `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewBulkBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewBulkBpsPrice) ImplementsPriceNewParams() {

}

type PriceNewParamsNewBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]PriceNewParamsNewBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r PriceNewParamsNewBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PriceNewParamsNewBulkBpsPriceCadence string

const (
	PriceNewParamsNewBulkBpsPriceCadenceAnnual    PriceNewParamsNewBulkBpsPriceCadence = "annual"
	PriceNewParamsNewBulkBpsPriceCadenceMonthly   PriceNewParamsNewBulkBpsPriceCadence = "monthly"
	PriceNewParamsNewBulkBpsPriceCadenceQuarterly PriceNewParamsNewBulkBpsPriceCadence = "quarterly"
)

type PriceNewParamsNewBulkBpsPriceModelType string

const (
	PriceNewParamsNewBulkBpsPriceModelTypeBulkBps PriceNewParamsNewBulkBpsPriceModelType = "bulk_bps"
)

type PriceNewParamsNewBulkPrice struct {
	BulkConfig param.Field[PriceNewParamsNewBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                              `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewBulkPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewBulkPrice) ImplementsPriceNewParams() {

}

type PriceNewParamsNewBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]PriceNewParamsNewBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r PriceNewParamsNewBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r PriceNewParamsNewBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type PriceNewParamsNewBulkPriceCadence string

const (
	PriceNewParamsNewBulkPriceCadenceAnnual    PriceNewParamsNewBulkPriceCadence = "annual"
	PriceNewParamsNewBulkPriceCadenceMonthly   PriceNewParamsNewBulkPriceCadence = "monthly"
	PriceNewParamsNewBulkPriceCadenceQuarterly PriceNewParamsNewBulkPriceCadence = "quarterly"
)

type PriceNewParamsNewBulkPriceModelType string

const (
	PriceNewParamsNewBulkPriceModelTypeBulk PriceNewParamsNewBulkPriceModelType = "bulk"
)

type PriceNewParamsNewThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                              `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                       param.Field[string]                 `json:"name,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}] `json:"threshold_total_amount_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewThresholdTotalAmountPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewThresholdTotalAmountPriceCadence string

const (
	PriceNewParamsNewThresholdTotalAmountPriceCadenceAnnual    PriceNewParamsNewThresholdTotalAmountPriceCadence = "annual"
	PriceNewParamsNewThresholdTotalAmountPriceCadenceMonthly   PriceNewParamsNewThresholdTotalAmountPriceCadence = "monthly"
	PriceNewParamsNewThresholdTotalAmountPriceCadenceQuarterly PriceNewParamsNewThresholdTotalAmountPriceCadence = "quarterly"
)

type PriceNewParamsNewThresholdTotalAmountPriceModelType string

const (
	PriceNewParamsNewThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceNewParamsNewThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

type PriceNewParamsNewTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                       `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                param.Field[string]                 `json:"name,required"`
	TieredPackageConfig param.Field[map[string]interface{}] `json:"tiered_package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewTieredPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewTieredPackagePriceCadence string

const (
	PriceNewParamsNewTieredPackagePriceCadenceAnnual    PriceNewParamsNewTieredPackagePriceCadence = "annual"
	PriceNewParamsNewTieredPackagePriceCadenceMonthly   PriceNewParamsNewTieredPackagePriceCadence = "monthly"
	PriceNewParamsNewTieredPackagePriceCadenceQuarterly PriceNewParamsNewTieredPackagePriceCadence = "quarterly"
)

type PriceNewParamsNewTieredPackagePriceModelType string

const (
	PriceNewParamsNewTieredPackagePriceModelTypeTieredPackage PriceNewParamsNewTieredPackagePriceModelType = "tiered_package"
)

type PriceNewParamsNewTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                           `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                    param.Field[string]                 `json:"name,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}] `json:"tiered_with_minimum_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewTieredWithMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewTieredWithMinimumPriceCadence string

const (
	PriceNewParamsNewTieredWithMinimumPriceCadenceAnnual    PriceNewParamsNewTieredWithMinimumPriceCadence = "annual"
	PriceNewParamsNewTieredWithMinimumPriceCadenceMonthly   PriceNewParamsNewTieredWithMinimumPriceCadence = "monthly"
	PriceNewParamsNewTieredWithMinimumPriceCadenceQuarterly PriceNewParamsNewTieredWithMinimumPriceCadence = "quarterly"
)

type PriceNewParamsNewTieredWithMinimumPriceModelType string

const (
	PriceNewParamsNewTieredWithMinimumPriceModelTypeTieredWithMinimum PriceNewParamsNewTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

type PriceNewParamsNewPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[PriceNewParamsNewPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                               `json:"item_id,required"`
	ModelType param.Field[PriceNewParamsNewPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name                        param.Field[string]                 `json:"name,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}] `json:"package_with_allocation_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewPackageWithAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewPackageWithAllocationPriceCadence string

const (
	PriceNewParamsNewPackageWithAllocationPriceCadenceAnnual    PriceNewParamsNewPackageWithAllocationPriceCadence = "annual"
	PriceNewParamsNewPackageWithAllocationPriceCadenceMonthly   PriceNewParamsNewPackageWithAllocationPriceCadence = "monthly"
	PriceNewParamsNewPackageWithAllocationPriceCadenceQuarterly PriceNewParamsNewPackageWithAllocationPriceCadence = "quarterly"
)

type PriceNewParamsNewPackageWithAllocationPriceModelType string

const (
	PriceNewParamsNewPackageWithAllocationPriceModelTypePackageWithAllocation PriceNewParamsNewPackageWithAllocationPriceModelType = "package_with_allocation"
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
