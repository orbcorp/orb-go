// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
// [PriceBulkPrice], [PriceThresholdTotalAmountPrice], [PriceTieredPackagePrice],
// [PriceTieredWithMinimumPrice], [PricePackageWithAllocationPrice],
// [PriceUnitWithPercentPrice] or [PriceMatrixWithAllocationPrice].
type Price interface {
	implementsPrice()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*Price)(nil)).Elem(),
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
			Type:               reflect.TypeOf(PriceTieredWithMinimumPrice{}),
			DiscriminatorValue: "tiered_with_minimum",
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
	)
}

type PriceUnitPrice struct {
	ID                 string                       `json:"id,required"`
	BillableMetric     PriceUnitPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceUnitPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                    `json:"created_at,required" format:"date-time"`
	Currency           string                       `json:"currency,required"`
	Discount           shared.Discount              `json:"discount,required,nullable"`
	ExternalPriceID    string                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                      `json:"fixed_price_quantity,required,nullable"`
	Item               PriceUnitPriceItem           `json:"item,required"`
	Maximum            PriceUnitPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                       `json:"maximum_amount,required,nullable"`
	Minimum            PriceUnitPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                       `json:"minimum_amount,required,nullable"`
	ModelType          PriceUnitPriceModelType      `json:"model_type,required"`
	Name               string                       `json:"name,required"`
	PlanPhaseOrder     int64                        `json:"plan_phase_order,required,nullable"`
	PriceType          PriceUnitPricePriceType      `json:"price_type,required"`
	UnitConfig         PriceUnitPriceUnitConfig     `json:"unit_config,required"`
	JSON               priceUnitPriceJSON           `json:"-"`
}

// priceUnitPriceJSON contains the JSON metadata for the struct [PriceUnitPrice]
type priceUnitPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	UnitConfig         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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

type PriceUnitPriceCadence string

const (
	PriceUnitPriceCadenceOneTime   PriceUnitPriceCadence = "one_time"
	PriceUnitPriceCadenceMonthly   PriceUnitPriceCadence = "monthly"
	PriceUnitPriceCadenceQuarterly PriceUnitPriceCadence = "quarterly"
	PriceUnitPriceCadenceAnnual    PriceUnitPriceCadence = "annual"
)

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

type PriceUnitPricePriceType string

const (
	PriceUnitPricePriceTypeUsagePrice PriceUnitPricePriceType = "usage_price"
	PriceUnitPricePriceTypeFixedPrice PriceUnitPricePriceType = "fixed_price"
)

type PriceUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount string `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor float64                      `json:"scaling_factor,nullable"`
	JSON          priceUnitPriceUnitConfigJSON `json:"-"`
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

func (r priceUnitPriceUnitConfigJSON) RawJSON() string {
	return r.raw
}

type PricePackagePrice struct {
	ID                 string                          `json:"id,required"`
	BillableMetric     PricePackagePriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PricePackagePriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                       `json:"created_at,required" format:"date-time"`
	Currency           string                          `json:"currency,required"`
	Discount           shared.Discount                 `json:"discount,required,nullable"`
	ExternalPriceID    string                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                         `json:"fixed_price_quantity,required,nullable"`
	Item               PricePackagePriceItem           `json:"item,required"`
	Maximum            PricePackagePriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                          `json:"maximum_amount,required,nullable"`
	Minimum            PricePackagePriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                          `json:"minimum_amount,required,nullable"`
	ModelType          PricePackagePriceModelType      `json:"model_type,required"`
	Name               string                          `json:"name,required"`
	PackageConfig      PricePackagePricePackageConfig  `json:"package_config,required"`
	PlanPhaseOrder     int64                           `json:"plan_phase_order,required,nullable"`
	PriceType          PricePackagePricePriceType      `json:"price_type,required"`
	JSON               pricePackagePriceJSON           `json:"-"`
}

// pricePackagePriceJSON contains the JSON metadata for the struct
// [PricePackagePrice]
type pricePackagePriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PackageConfig      apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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

type PricePackagePriceCadence string

const (
	PricePackagePriceCadenceOneTime   PricePackagePriceCadence = "one_time"
	PricePackagePriceCadenceMonthly   PricePackagePriceCadence = "monthly"
	PricePackagePriceCadenceQuarterly PricePackagePriceCadence = "quarterly"
	PricePackagePriceCadenceAnnual    PricePackagePriceCadence = "annual"
)

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

type PricePackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount string `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize int64                              `json:"package_size,nullable"`
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

type PriceMatrixPrice struct {
	ID                 string                         `json:"id,required"`
	BillableMetric     PriceMatrixPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceMatrixPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                      `json:"created_at,required" format:"date-time"`
	Currency           string                         `json:"currency,required"`
	Discount           shared.Discount                `json:"discount,required,nullable"`
	ExternalPriceID    string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                        `json:"fixed_price_quantity,required,nullable"`
	Item               PriceMatrixPriceItem           `json:"item,required"`
	MatrixConfig       PriceMatrixPriceMatrixConfig   `json:"matrix_config,required"`
	Maximum            PriceMatrixPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                         `json:"maximum_amount,required,nullable"`
	Minimum            PriceMatrixPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                         `json:"minimum_amount,required,nullable"`
	ModelType          PriceMatrixPriceModelType      `json:"model_type,required"`
	Name               string                         `json:"name,required"`
	PlanPhaseOrder     int64                          `json:"plan_phase_order,required,nullable"`
	PriceType          PriceMatrixPricePriceType      `json:"price_type,required"`
	JSON               priceMatrixPriceJSON           `json:"-"`
}

// priceMatrixPriceJSON contains the JSON metadata for the struct
// [PriceMatrixPrice]
type priceMatrixPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	MatrixConfig       apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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

type PriceMatrixPriceCadence string

const (
	PriceMatrixPriceCadenceOneTime   PriceMatrixPriceCadence = "one_time"
	PriceMatrixPriceCadenceMonthly   PriceMatrixPriceCadence = "monthly"
	PriceMatrixPriceCadenceQuarterly PriceMatrixPriceCadence = "quarterly"
	PriceMatrixPriceCadenceAnnual    PriceMatrixPriceCadence = "annual"
)

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
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor float64                          `json:"scaling_factor,nullable"`
	JSON          priceMatrixPriceMatrixConfigJSON `json:"-"`
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

func (r priceMatrixPriceMatrixConfigJSON) RawJSON() string {
	return r.raw
}

type PriceMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues []string `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount string `json:"unit_amount,required"`
	// Optional multiplier to scale rated quantities by
	ScalingFactor float64                                     `json:"scaling_factor,nullable"`
	JSON          priceMatrixPriceMatrixConfigMatrixValueJSON `json:"-"`
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

type PriceMatrixPricePriceType string

const (
	PriceMatrixPricePriceTypeUsagePrice PriceMatrixPricePriceType = "usage_price"
	PriceMatrixPricePriceTypeFixedPrice PriceMatrixPricePriceType = "fixed_price"
)

type PriceTieredPrice struct {
	ID                 string                         `json:"id,required"`
	BillableMetric     PriceTieredPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence            PriceTieredPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                      `json:"created_at,required" format:"date-time"`
	Currency           string                         `json:"currency,required"`
	Discount           shared.Discount                `json:"discount,required,nullable"`
	ExternalPriceID    string                         `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                        `json:"fixed_price_quantity,required,nullable"`
	Item               PriceTieredPriceItem           `json:"item,required"`
	Maximum            PriceTieredPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                         `json:"maximum_amount,required,nullable"`
	Minimum            PriceTieredPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                         `json:"minimum_amount,required,nullable"`
	ModelType          PriceTieredPriceModelType      `json:"model_type,required"`
	Name               string                         `json:"name,required"`
	PlanPhaseOrder     int64                          `json:"plan_phase_order,required,nullable"`
	PriceType          PriceTieredPricePriceType      `json:"price_type,required"`
	TieredConfig       PriceTieredPriceTieredConfig   `json:"tiered_config,required"`
	JSON               priceTieredPriceJSON           `json:"-"`
}

// priceTieredPriceJSON contains the JSON metadata for the struct
// [PriceTieredPrice]
type priceTieredPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	TieredConfig       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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

type PriceTieredPriceCadence string

const (
	PriceTieredPriceCadenceOneTime   PriceTieredPriceCadence = "one_time"
	PriceTieredPriceCadenceMonthly   PriceTieredPriceCadence = "monthly"
	PriceTieredPriceCadenceQuarterly PriceTieredPriceCadence = "quarterly"
	PriceTieredPriceCadenceAnnual    PriceTieredPriceCadence = "annual"
)

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

type PriceTieredPricePriceType string

const (
	PriceTieredPricePriceTypeUsagePrice PriceTieredPricePriceType = "usage_price"
	PriceTieredPricePriceTypeFixedPrice PriceTieredPricePriceType = "fixed_price"
)

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
	ID                 string                             `json:"id,required"`
	BillableMetric     PriceTieredBpsPriceBillableMetric  `json:"billable_metric,required,nullable"`
	Cadence            PriceTieredBpsPriceCadence         `json:"cadence,required"`
	CreatedAt          time.Time                          `json:"created_at,required" format:"date-time"`
	Currency           string                             `json:"currency,required"`
	Discount           shared.Discount                    `json:"discount,required,nullable"`
	ExternalPriceID    string                             `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                            `json:"fixed_price_quantity,required,nullable"`
	Item               PriceTieredBpsPriceItem            `json:"item,required"`
	Maximum            PriceTieredBpsPriceMaximum         `json:"maximum,required,nullable"`
	MaximumAmount      string                             `json:"maximum_amount,required,nullable"`
	Minimum            PriceTieredBpsPriceMinimum         `json:"minimum,required,nullable"`
	MinimumAmount      string                             `json:"minimum_amount,required,nullable"`
	ModelType          PriceTieredBpsPriceModelType       `json:"model_type,required"`
	Name               string                             `json:"name,required"`
	PlanPhaseOrder     int64                              `json:"plan_phase_order,required,nullable"`
	PriceType          PriceTieredBpsPricePriceType       `json:"price_type,required"`
	TieredBpsConfig    PriceTieredBpsPriceTieredBpsConfig `json:"tiered_bps_config,required"`
	JSON               priceTieredBpsPriceJSON            `json:"-"`
}

// priceTieredBpsPriceJSON contains the JSON metadata for the struct
// [PriceTieredBpsPrice]
type priceTieredBpsPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	TieredBpsConfig    apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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

type PriceTieredBpsPriceCadence string

const (
	PriceTieredBpsPriceCadenceOneTime   PriceTieredBpsPriceCadence = "one_time"
	PriceTieredBpsPriceCadenceMonthly   PriceTieredBpsPriceCadence = "monthly"
	PriceTieredBpsPriceCadenceQuarterly PriceTieredBpsPriceCadence = "quarterly"
	PriceTieredBpsPriceCadenceAnnual    PriceTieredBpsPriceCadence = "annual"
)

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

type PriceTieredBpsPricePriceType string

const (
	PriceTieredBpsPricePriceTypeUsagePrice PriceTieredBpsPricePriceType = "usage_price"
	PriceTieredBpsPricePriceTypeFixedPrice PriceTieredBpsPricePriceType = "fixed_price"
)

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
	ID                 string                      `json:"id,required"`
	BillableMetric     PriceBpsPriceBillableMetric `json:"billable_metric,required,nullable"`
	BpsConfig          PriceBpsPriceBpsConfig      `json:"bps_config,required"`
	Cadence            PriceBpsPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                   `json:"created_at,required" format:"date-time"`
	Currency           string                      `json:"currency,required"`
	Discount           shared.Discount             `json:"discount,required,nullable"`
	ExternalPriceID    string                      `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                     `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBpsPriceItem           `json:"item,required"`
	Maximum            PriceBpsPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                      `json:"maximum_amount,required,nullable"`
	Minimum            PriceBpsPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                      `json:"minimum_amount,required,nullable"`
	ModelType          PriceBpsPriceModelType      `json:"model_type,required"`
	Name               string                      `json:"name,required"`
	PlanPhaseOrder     int64                       `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBpsPricePriceType      `json:"price_type,required"`
	JSON               priceBpsPriceJSON           `json:"-"`
}

// priceBpsPriceJSON contains the JSON metadata for the struct [PriceBpsPrice]
type priceBpsPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	BpsConfig          apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	PriceBpsPriceCadenceOneTime   PriceBpsPriceCadence = "one_time"
	PriceBpsPriceCadenceMonthly   PriceBpsPriceCadence = "monthly"
	PriceBpsPriceCadenceQuarterly PriceBpsPriceCadence = "quarterly"
	PriceBpsPriceCadenceAnnual    PriceBpsPriceCadence = "annual"
)

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

type PriceBpsPricePriceType string

const (
	PriceBpsPricePriceTypeUsagePrice PriceBpsPricePriceType = "usage_price"
	PriceBpsPricePriceTypeFixedPrice PriceBpsPricePriceType = "fixed_price"
)

type PriceBulkBpsPrice struct {
	ID                 string                          `json:"id,required"`
	BillableMetric     PriceBulkBpsPriceBillableMetric `json:"billable_metric,required,nullable"`
	BulkBpsConfig      PriceBulkBpsPriceBulkBpsConfig  `json:"bulk_bps_config,required"`
	Cadence            PriceBulkBpsPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                       `json:"created_at,required" format:"date-time"`
	Currency           string                          `json:"currency,required"`
	Discount           shared.Discount                 `json:"discount,required,nullable"`
	ExternalPriceID    string                          `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                         `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBulkBpsPriceItem           `json:"item,required"`
	Maximum            PriceBulkBpsPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                          `json:"maximum_amount,required,nullable"`
	Minimum            PriceBulkBpsPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                          `json:"minimum_amount,required,nullable"`
	ModelType          PriceBulkBpsPriceModelType      `json:"model_type,required"`
	Name               string                          `json:"name,required"`
	PlanPhaseOrder     int64                           `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBulkBpsPricePriceType      `json:"price_type,required"`
	JSON               priceBulkBpsPriceJSON           `json:"-"`
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
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	PriceBulkBpsPriceCadenceOneTime   PriceBulkBpsPriceCadence = "one_time"
	PriceBulkBpsPriceCadenceMonthly   PriceBulkBpsPriceCadence = "monthly"
	PriceBulkBpsPriceCadenceQuarterly PriceBulkBpsPriceCadence = "quarterly"
	PriceBulkBpsPriceCadenceAnnual    PriceBulkBpsPriceCadence = "annual"
)

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

type PriceBulkBpsPricePriceType string

const (
	PriceBulkBpsPricePriceTypeUsagePrice PriceBulkBpsPricePriceType = "usage_price"
	PriceBulkBpsPricePriceTypeFixedPrice PriceBulkBpsPricePriceType = "fixed_price"
)

type PriceBulkPrice struct {
	ID                 string                       `json:"id,required"`
	BillableMetric     PriceBulkPriceBillableMetric `json:"billable_metric,required,nullable"`
	BulkConfig         PriceBulkPriceBulkConfig     `json:"bulk_config,required"`
	Cadence            PriceBulkPriceCadence        `json:"cadence,required"`
	CreatedAt          time.Time                    `json:"created_at,required" format:"date-time"`
	Currency           string                       `json:"currency,required"`
	Discount           shared.Discount              `json:"discount,required,nullable"`
	ExternalPriceID    string                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity float64                      `json:"fixed_price_quantity,required,nullable"`
	Item               PriceBulkPriceItem           `json:"item,required"`
	Maximum            PriceBulkPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount      string                       `json:"maximum_amount,required,nullable"`
	Minimum            PriceBulkPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount      string                       `json:"minimum_amount,required,nullable"`
	ModelType          PriceBulkPriceModelType      `json:"model_type,required"`
	Name               string                       `json:"name,required"`
	PlanPhaseOrder     int64                        `json:"plan_phase_order,required,nullable"`
	PriceType          PriceBulkPricePriceType      `json:"price_type,required"`
	JSON               priceBulkPriceJSON           `json:"-"`
}

// priceBulkPriceJSON contains the JSON metadata for the struct [PriceBulkPrice]
type priceBulkPriceJSON struct {
	ID                 apijson.Field
	BillableMetric     apijson.Field
	BulkConfig         apijson.Field
	Cadence            apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Discount           apijson.Field
	ExternalPriceID    apijson.Field
	FixedPriceQuantity apijson.Field
	Item               apijson.Field
	Maximum            apijson.Field
	MaximumAmount      apijson.Field
	Minimum            apijson.Field
	MinimumAmount      apijson.Field
	ModelType          apijson.Field
	Name               apijson.Field
	PlanPhaseOrder     apijson.Field
	PriceType          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
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
	PriceBulkPriceCadenceOneTime   PriceBulkPriceCadence = "one_time"
	PriceBulkPriceCadenceMonthly   PriceBulkPriceCadence = "monthly"
	PriceBulkPriceCadenceQuarterly PriceBulkPriceCadence = "quarterly"
	PriceBulkPriceCadenceAnnual    PriceBulkPriceCadence = "annual"
)

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

type PriceBulkPricePriceType string

const (
	PriceBulkPricePriceTypeUsagePrice PriceBulkPricePriceType = "usage_price"
	PriceBulkPricePriceTypeFixedPrice PriceBulkPricePriceType = "fixed_price"
)

type PriceThresholdTotalAmountPrice struct {
	ID                         string                                       `json:"id,required"`
	BillableMetric             PriceThresholdTotalAmountPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                    PriceThresholdTotalAmountPriceCadence        `json:"cadence,required"`
	CreatedAt                  time.Time                                    `json:"created_at,required" format:"date-time"`
	Currency                   string                                       `json:"currency,required"`
	Discount                   shared.Discount                              `json:"discount,required,nullable"`
	ExternalPriceID            string                                       `json:"external_price_id,required,nullable"`
	FixedPriceQuantity         float64                                      `json:"fixed_price_quantity,required,nullable"`
	Item                       PriceThresholdTotalAmountPriceItem           `json:"item,required"`
	Maximum                    PriceThresholdTotalAmountPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount              string                                       `json:"maximum_amount,required,nullable"`
	Minimum                    PriceThresholdTotalAmountPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount              string                                       `json:"minimum_amount,required,nullable"`
	ModelType                  PriceThresholdTotalAmountPriceModelType      `json:"model_type,required"`
	Name                       string                                       `json:"name,required"`
	PlanPhaseOrder             int64                                        `json:"plan_phase_order,required,nullable"`
	PriceType                  PriceThresholdTotalAmountPricePriceType      `json:"price_type,required"`
	ThresholdTotalAmountConfig map[string]interface{}                       `json:"threshold_total_amount_config,required"`
	JSON                       priceThresholdTotalAmountPriceJSON           `json:"-"`
}

// priceThresholdTotalAmountPriceJSON contains the JSON metadata for the struct
// [PriceThresholdTotalAmountPrice]
type priceThresholdTotalAmountPriceJSON struct {
	ID                         apijson.Field
	BillableMetric             apijson.Field
	Cadence                    apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	Discount                   apijson.Field
	ExternalPriceID            apijson.Field
	FixedPriceQuantity         apijson.Field
	Item                       apijson.Field
	Maximum                    apijson.Field
	MaximumAmount              apijson.Field
	Minimum                    apijson.Field
	MinimumAmount              apijson.Field
	ModelType                  apijson.Field
	Name                       apijson.Field
	PlanPhaseOrder             apijson.Field
	PriceType                  apijson.Field
	ThresholdTotalAmountConfig apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
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

type PriceThresholdTotalAmountPriceCadence string

const (
	PriceThresholdTotalAmountPriceCadenceOneTime   PriceThresholdTotalAmountPriceCadence = "one_time"
	PriceThresholdTotalAmountPriceCadenceMonthly   PriceThresholdTotalAmountPriceCadence = "monthly"
	PriceThresholdTotalAmountPriceCadenceQuarterly PriceThresholdTotalAmountPriceCadence = "quarterly"
	PriceThresholdTotalAmountPriceCadenceAnnual    PriceThresholdTotalAmountPriceCadence = "annual"
)

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

type PriceThresholdTotalAmountPricePriceType string

const (
	PriceThresholdTotalAmountPricePriceTypeUsagePrice PriceThresholdTotalAmountPricePriceType = "usage_price"
	PriceThresholdTotalAmountPricePriceTypeFixedPrice PriceThresholdTotalAmountPricePriceType = "fixed_price"
)

type PriceTieredPackagePrice struct {
	ID                  string                                `json:"id,required"`
	BillableMetric      PriceTieredPackagePriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence             PriceTieredPackagePriceCadence        `json:"cadence,required"`
	CreatedAt           time.Time                             `json:"created_at,required" format:"date-time"`
	Currency            string                                `json:"currency,required"`
	Discount            shared.Discount                       `json:"discount,required,nullable"`
	ExternalPriceID     string                                `json:"external_price_id,required,nullable"`
	FixedPriceQuantity  float64                               `json:"fixed_price_quantity,required,nullable"`
	Item                PriceTieredPackagePriceItem           `json:"item,required"`
	Maximum             PriceTieredPackagePriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount       string                                `json:"maximum_amount,required,nullable"`
	Minimum             PriceTieredPackagePriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount       string                                `json:"minimum_amount,required,nullable"`
	ModelType           PriceTieredPackagePriceModelType      `json:"model_type,required"`
	Name                string                                `json:"name,required"`
	PlanPhaseOrder      int64                                 `json:"plan_phase_order,required,nullable"`
	PriceType           PriceTieredPackagePricePriceType      `json:"price_type,required"`
	TieredPackageConfig map[string]interface{}                `json:"tiered_package_config,required"`
	JSON                priceTieredPackagePriceJSON           `json:"-"`
}

// priceTieredPackagePriceJSON contains the JSON metadata for the struct
// [PriceTieredPackagePrice]
type priceTieredPackagePriceJSON struct {
	ID                  apijson.Field
	BillableMetric      apijson.Field
	Cadence             apijson.Field
	CreatedAt           apijson.Field
	Currency            apijson.Field
	Discount            apijson.Field
	ExternalPriceID     apijson.Field
	FixedPriceQuantity  apijson.Field
	Item                apijson.Field
	Maximum             apijson.Field
	MaximumAmount       apijson.Field
	Minimum             apijson.Field
	MinimumAmount       apijson.Field
	ModelType           apijson.Field
	Name                apijson.Field
	PlanPhaseOrder      apijson.Field
	PriceType           apijson.Field
	TieredPackageConfig apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
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

type PriceTieredPackagePriceCadence string

const (
	PriceTieredPackagePriceCadenceOneTime   PriceTieredPackagePriceCadence = "one_time"
	PriceTieredPackagePriceCadenceMonthly   PriceTieredPackagePriceCadence = "monthly"
	PriceTieredPackagePriceCadenceQuarterly PriceTieredPackagePriceCadence = "quarterly"
	PriceTieredPackagePriceCadenceAnnual    PriceTieredPackagePriceCadence = "annual"
)

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

type PriceTieredPackagePricePriceType string

const (
	PriceTieredPackagePricePriceTypeUsagePrice PriceTieredPackagePricePriceType = "usage_price"
	PriceTieredPackagePricePriceTypeFixedPrice PriceTieredPackagePricePriceType = "fixed_price"
)

type PriceTieredWithMinimumPrice struct {
	ID                      string                                    `json:"id,required"`
	BillableMetric          PriceTieredWithMinimumPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                 PriceTieredWithMinimumPriceCadence        `json:"cadence,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	Currency                string                                    `json:"currency,required"`
	Discount                shared.Discount                           `json:"discount,required,nullable"`
	ExternalPriceID         string                                    `json:"external_price_id,required,nullable"`
	FixedPriceQuantity      float64                                   `json:"fixed_price_quantity,required,nullable"`
	Item                    PriceTieredWithMinimumPriceItem           `json:"item,required"`
	Maximum                 PriceTieredWithMinimumPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount           string                                    `json:"maximum_amount,required,nullable"`
	Minimum                 PriceTieredWithMinimumPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount           string                                    `json:"minimum_amount,required,nullable"`
	ModelType               PriceTieredWithMinimumPriceModelType      `json:"model_type,required"`
	Name                    string                                    `json:"name,required"`
	PlanPhaseOrder          int64                                     `json:"plan_phase_order,required,nullable"`
	PriceType               PriceTieredWithMinimumPricePriceType      `json:"price_type,required"`
	TieredWithMinimumConfig map[string]interface{}                    `json:"tiered_with_minimum_config,required"`
	JSON                    priceTieredWithMinimumPriceJSON           `json:"-"`
}

// priceTieredWithMinimumPriceJSON contains the JSON metadata for the struct
// [PriceTieredWithMinimumPrice]
type priceTieredWithMinimumPriceJSON struct {
	ID                      apijson.Field
	BillableMetric          apijson.Field
	Cadence                 apijson.Field
	CreatedAt               apijson.Field
	Currency                apijson.Field
	Discount                apijson.Field
	ExternalPriceID         apijson.Field
	FixedPriceQuantity      apijson.Field
	Item                    apijson.Field
	Maximum                 apijson.Field
	MaximumAmount           apijson.Field
	Minimum                 apijson.Field
	MinimumAmount           apijson.Field
	ModelType               apijson.Field
	Name                    apijson.Field
	PlanPhaseOrder          apijson.Field
	PriceType               apijson.Field
	TieredWithMinimumConfig apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
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

type PriceTieredWithMinimumPriceCadence string

const (
	PriceTieredWithMinimumPriceCadenceOneTime   PriceTieredWithMinimumPriceCadence = "one_time"
	PriceTieredWithMinimumPriceCadenceMonthly   PriceTieredWithMinimumPriceCadence = "monthly"
	PriceTieredWithMinimumPriceCadenceQuarterly PriceTieredWithMinimumPriceCadence = "quarterly"
	PriceTieredWithMinimumPriceCadenceAnnual    PriceTieredWithMinimumPriceCadence = "annual"
)

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

type PriceTieredWithMinimumPricePriceType string

const (
	PriceTieredWithMinimumPricePriceTypeUsagePrice PriceTieredWithMinimumPricePriceType = "usage_price"
	PriceTieredWithMinimumPricePriceTypeFixedPrice PriceTieredWithMinimumPricePriceType = "fixed_price"
)

type PricePackageWithAllocationPrice struct {
	ID                          string                                        `json:"id,required"`
	BillableMetric              PricePackageWithAllocationPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence                     PricePackageWithAllocationPriceCadence        `json:"cadence,required"`
	CreatedAt                   time.Time                                     `json:"created_at,required" format:"date-time"`
	Currency                    string                                        `json:"currency,required"`
	Discount                    shared.Discount                               `json:"discount,required,nullable"`
	ExternalPriceID             string                                        `json:"external_price_id,required,nullable"`
	FixedPriceQuantity          float64                                       `json:"fixed_price_quantity,required,nullable"`
	Item                        PricePackageWithAllocationPriceItem           `json:"item,required"`
	Maximum                     PricePackageWithAllocationPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount               string                                        `json:"maximum_amount,required,nullable"`
	Minimum                     PricePackageWithAllocationPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount               string                                        `json:"minimum_amount,required,nullable"`
	ModelType                   PricePackageWithAllocationPriceModelType      `json:"model_type,required"`
	Name                        string                                        `json:"name,required"`
	PackageWithAllocationConfig map[string]interface{}                        `json:"package_with_allocation_config,required"`
	PlanPhaseOrder              int64                                         `json:"plan_phase_order,required,nullable"`
	PriceType                   PricePackageWithAllocationPricePriceType      `json:"price_type,required"`
	JSON                        pricePackageWithAllocationPriceJSON           `json:"-"`
}

// pricePackageWithAllocationPriceJSON contains the JSON metadata for the struct
// [PricePackageWithAllocationPrice]
type pricePackageWithAllocationPriceJSON struct {
	ID                          apijson.Field
	BillableMetric              apijson.Field
	Cadence                     apijson.Field
	CreatedAt                   apijson.Field
	Currency                    apijson.Field
	Discount                    apijson.Field
	ExternalPriceID             apijson.Field
	FixedPriceQuantity          apijson.Field
	Item                        apijson.Field
	Maximum                     apijson.Field
	MaximumAmount               apijson.Field
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

type PricePackageWithAllocationPriceCadence string

const (
	PricePackageWithAllocationPriceCadenceOneTime   PricePackageWithAllocationPriceCadence = "one_time"
	PricePackageWithAllocationPriceCadenceMonthly   PricePackageWithAllocationPriceCadence = "monthly"
	PricePackageWithAllocationPriceCadenceQuarterly PricePackageWithAllocationPriceCadence = "quarterly"
	PricePackageWithAllocationPriceCadenceAnnual    PricePackageWithAllocationPriceCadence = "annual"
)

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

type PricePackageWithAllocationPricePriceType string

const (
	PricePackageWithAllocationPricePriceTypeUsagePrice PricePackageWithAllocationPricePriceType = "usage_price"
	PricePackageWithAllocationPricePriceTypeFixedPrice PricePackageWithAllocationPricePriceType = "fixed_price"
)

type PriceUnitWithPercentPrice struct {
	ID                    string                                  `json:"id,required"`
	BillableMetric        PriceUnitWithPercentPriceBillableMetric `json:"billable_metric,required,nullable"`
	Cadence               PriceUnitWithPercentPriceCadence        `json:"cadence,required"`
	CreatedAt             time.Time                               `json:"created_at,required" format:"date-time"`
	Currency              string                                  `json:"currency,required"`
	Discount              shared.Discount                         `json:"discount,required,nullable"`
	ExternalPriceID       string                                  `json:"external_price_id,required,nullable"`
	FixedPriceQuantity    float64                                 `json:"fixed_price_quantity,required,nullable"`
	Item                  PriceUnitWithPercentPriceItem           `json:"item,required"`
	Maximum               PriceUnitWithPercentPriceMaximum        `json:"maximum,required,nullable"`
	MaximumAmount         string                                  `json:"maximum_amount,required,nullable"`
	Minimum               PriceUnitWithPercentPriceMinimum        `json:"minimum,required,nullable"`
	MinimumAmount         string                                  `json:"minimum_amount,required,nullable"`
	ModelType             PriceUnitWithPercentPriceModelType      `json:"model_type,required"`
	Name                  string                                  `json:"name,required"`
	PlanPhaseOrder        int64                                   `json:"plan_phase_order,required,nullable"`
	PriceType             PriceUnitWithPercentPricePriceType      `json:"price_type,required"`
	UnitWithPercentConfig map[string]interface{}                  `json:"unit_with_percent_config,required"`
	JSON                  priceUnitWithPercentPriceJSON           `json:"-"`
}

// priceUnitWithPercentPriceJSON contains the JSON metadata for the struct
// [PriceUnitWithPercentPrice]
type priceUnitWithPercentPriceJSON struct {
	ID                    apijson.Field
	BillableMetric        apijson.Field
	Cadence               apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	Discount              apijson.Field
	ExternalPriceID       apijson.Field
	FixedPriceQuantity    apijson.Field
	Item                  apijson.Field
	Maximum               apijson.Field
	MaximumAmount         apijson.Field
	Minimum               apijson.Field
	MinimumAmount         apijson.Field
	ModelType             apijson.Field
	Name                  apijson.Field
	PlanPhaseOrder        apijson.Field
	PriceType             apijson.Field
	UnitWithPercentConfig apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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

type PriceUnitWithPercentPriceCadence string

const (
	PriceUnitWithPercentPriceCadenceOneTime   PriceUnitWithPercentPriceCadence = "one_time"
	PriceUnitWithPercentPriceCadenceMonthly   PriceUnitWithPercentPriceCadence = "monthly"
	PriceUnitWithPercentPriceCadenceQuarterly PriceUnitWithPercentPriceCadence = "quarterly"
	PriceUnitWithPercentPriceCadenceAnnual    PriceUnitWithPercentPriceCadence = "annual"
)

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

type PriceUnitWithPercentPricePriceType string

const (
	PriceUnitWithPercentPricePriceTypeUsagePrice PriceUnitWithPercentPricePriceType = "usage_price"
	PriceUnitWithPercentPricePriceTypeFixedPrice PriceUnitWithPercentPricePriceType = "fixed_price"
)

type PriceMatrixWithAllocationPrice struct {
	ID                         string                                                   `json:"id,required"`
	BillableMetric             PriceMatrixWithAllocationPriceBillableMetric             `json:"billable_metric,required,nullable"`
	Cadence                    PriceMatrixWithAllocationPriceCadence                    `json:"cadence,required"`
	CreatedAt                  time.Time                                                `json:"created_at,required" format:"date-time"`
	Currency                   string                                                   `json:"currency,required"`
	Discount                   shared.Discount                                          `json:"discount,required,nullable"`
	ExternalPriceID            string                                                   `json:"external_price_id,required,nullable"`
	FixedPriceQuantity         float64                                                  `json:"fixed_price_quantity,required,nullable"`
	Item                       PriceMatrixWithAllocationPriceItem                       `json:"item,required"`
	MatrixWithAllocationConfig PriceMatrixWithAllocationPriceMatrixWithAllocationConfig `json:"matrix_with_allocation_config,required"`
	Maximum                    PriceMatrixWithAllocationPriceMaximum                    `json:"maximum,required,nullable"`
	MaximumAmount              string                                                   `json:"maximum_amount,required,nullable"`
	Minimum                    PriceMatrixWithAllocationPriceMinimum                    `json:"minimum,required,nullable"`
	MinimumAmount              string                                                   `json:"minimum_amount,required,nullable"`
	ModelType                  PriceMatrixWithAllocationPriceModelType                  `json:"model_type,required"`
	Name                       string                                                   `json:"name,required"`
	PlanPhaseOrder             int64                                                    `json:"plan_phase_order,required,nullable"`
	PriceType                  PriceMatrixWithAllocationPricePriceType                  `json:"price_type,required"`
	JSON                       priceMatrixWithAllocationPriceJSON                       `json:"-"`
}

// priceMatrixWithAllocationPriceJSON contains the JSON metadata for the struct
// [PriceMatrixWithAllocationPrice]
type priceMatrixWithAllocationPriceJSON struct {
	ID                         apijson.Field
	BillableMetric             apijson.Field
	Cadence                    apijson.Field
	CreatedAt                  apijson.Field
	Currency                   apijson.Field
	Discount                   apijson.Field
	ExternalPriceID            apijson.Field
	FixedPriceQuantity         apijson.Field
	Item                       apijson.Field
	MatrixWithAllocationConfig apijson.Field
	Maximum                    apijson.Field
	MaximumAmount              apijson.Field
	Minimum                    apijson.Field
	MinimumAmount              apijson.Field
	ModelType                  apijson.Field
	Name                       apijson.Field
	PlanPhaseOrder             apijson.Field
	PriceType                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
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

type PriceMatrixWithAllocationPriceCadence string

const (
	PriceMatrixWithAllocationPriceCadenceOneTime   PriceMatrixWithAllocationPriceCadence = "one_time"
	PriceMatrixWithAllocationPriceCadenceMonthly   PriceMatrixWithAllocationPriceCadence = "monthly"
	PriceMatrixWithAllocationPriceCadenceQuarterly PriceMatrixWithAllocationPriceCadence = "quarterly"
	PriceMatrixWithAllocationPriceCadenceAnnual    PriceMatrixWithAllocationPriceCadence = "annual"
)

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
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor float64                                                      `json:"scaling_factor,nullable"`
	JSON          priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON `json:"-"`
}

// priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON contains the JSON
// metadata for the struct
// [PriceMatrixWithAllocationPriceMatrixWithAllocationConfig]
type priceMatrixWithAllocationPriceMatrixWithAllocationConfigJSON struct {
	Allocation        apijson.Field
	DefaultUnitAmount apijson.Field
	Dimensions        apijson.Field
	MatrixValues      apijson.Field
	ScalingFactor     apijson.Field
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
	UnitAmount string `json:"unit_amount,required"`
	// Optional multiplier to scale rated quantities by
	ScalingFactor float64                                                                 `json:"scaling_factor,nullable"`
	JSON          priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON `json:"-"`
}

// priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON contains
// the JSON metadata for the struct
// [PriceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue]
type priceMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValueJSON struct {
	DimensionValues apijson.Field
	UnitAmount      apijson.Field
	ScalingFactor   apijson.Field
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

type PriceMatrixWithAllocationPricePriceType string

const (
	PriceMatrixWithAllocationPricePriceTypeUsagePrice PriceMatrixWithAllocationPricePriceType = "usage_price"
	PriceMatrixWithAllocationPricePriceTypeFixedPrice PriceMatrixWithAllocationPricePriceType = "fixed_price"
)

// This interface is a union satisfied by one of the following:
// [PriceNewParamsNewFloatingUnitPrice], [PriceNewParamsNewFloatingPackagePrice],
// [PriceNewParamsNewFloatingMatrixPrice],
// [PriceNewParamsNewFloatingMatrixWithAllocationPrice],
// [PriceNewParamsNewFloatingTieredPrice],
// [PriceNewParamsNewFloatingTieredBpsPrice], [PriceNewParamsNewFloatingBpsPrice],
// [PriceNewParamsNewFloatingBulkBpsPrice], [PriceNewParamsNewFloatingBulkPrice],
// [PriceNewParamsNewFloatingThresholdTotalAmountPrice],
// [PriceNewParamsNewFloatingTieredPackagePrice],
// [PriceNewParamsNewFloatingTieredWithMinimumPrice],
// [PriceNewParamsNewFloatingPackageWithAllocationPrice],
// [PriceNewParamsNewFloatingTieredPackageWithMinimumPrice],
// [PriceNewParamsNewFloatingUnitWithPercentPrice].
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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingUnitPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingUnitPriceCadence string

const (
	PriceNewParamsNewFloatingUnitPriceCadenceAnnual    PriceNewParamsNewFloatingUnitPriceCadence = "annual"
	PriceNewParamsNewFloatingUnitPriceCadenceMonthly   PriceNewParamsNewFloatingUnitPriceCadence = "monthly"
	PriceNewParamsNewFloatingUnitPriceCadenceQuarterly PriceNewParamsNewFloatingUnitPriceCadence = "quarterly"
	PriceNewParamsNewFloatingUnitPriceCadenceOneTime   PriceNewParamsNewFloatingUnitPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingUnitPriceModelType string

const (
	PriceNewParamsNewFloatingUnitPriceModelTypeUnit PriceNewParamsNewFloatingUnitPriceModelType = "unit"
)

type PriceNewParamsNewFloatingUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Multiplier to scale rated quantity by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewFloatingUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingPackagePriceCadence string

const (
	PriceNewParamsNewFloatingPackagePriceCadenceAnnual    PriceNewParamsNewFloatingPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingPackagePriceCadenceMonthly   PriceNewParamsNewFloatingPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingPackagePriceCadenceQuarterly PriceNewParamsNewFloatingPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingPackagePriceCadenceOneTime   PriceNewParamsNewFloatingPackagePriceCadence = "one_time"
)

type PriceNewParamsNewFloatingPackagePriceModelType string

const (
	PriceNewParamsNewFloatingPackagePriceModelTypePackage PriceNewParamsNewFloatingPackagePriceModelType = "package"
)

type PriceNewParamsNewFloatingPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size"`
}

func (r PriceNewParamsNewFloatingPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMatrixPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMatrixPriceCadence string

const (
	PriceNewParamsNewFloatingMatrixPriceCadenceAnnual    PriceNewParamsNewFloatingMatrixPriceCadence = "annual"
	PriceNewParamsNewFloatingMatrixPriceCadenceMonthly   PriceNewParamsNewFloatingMatrixPriceCadence = "monthly"
	PriceNewParamsNewFloatingMatrixPriceCadenceQuarterly PriceNewParamsNewFloatingMatrixPriceCadence = "quarterly"
	PriceNewParamsNewFloatingMatrixPriceCadenceOneTime   PriceNewParamsNewFloatingMatrixPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PriceNewParamsNewFloatingMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
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
	// Optional multiplier to scale rated quantities by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewFloatingMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingMatrixPriceModelType string

const (
	PriceNewParamsNewFloatingMatrixPriceModelTypeMatrix PriceNewParamsNewFloatingMatrixPriceModelType = "matrix"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingMatrixWithAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceAnnual    PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "annual"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceMonthly   PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "monthly"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceQuarterly PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingMatrixWithAllocationPriceCadenceOneTime   PriceNewParamsNewFloatingMatrixWithAllocationPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
	// Default optional multiplier to scale rated quantities that fall into the default
	// bucket by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
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
	// Optional multiplier to scale rated quantities by
	ScalingFactor param.Field[float64] `json:"scaling_factor"`
}

func (r PriceNewParamsNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType string

const (
	PriceNewParamsNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation PriceNewParamsNewFloatingMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPriceCadence string

const (
	PriceNewParamsNewFloatingTieredPriceCadenceAnnual    PriceNewParamsNewFloatingTieredPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPriceCadenceMonthly   PriceNewParamsNewFloatingTieredPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPriceCadenceQuarterly PriceNewParamsNewFloatingTieredPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPriceCadenceOneTime   PriceNewParamsNewFloatingTieredPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingTieredPriceModelType string

const (
	PriceNewParamsNewFloatingTieredPriceModelTypeTiered PriceNewParamsNewFloatingTieredPriceModelType = "tiered"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredBpsPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredBpsPriceCadence string

const (
	PriceNewParamsNewFloatingTieredBpsPriceCadenceAnnual    PriceNewParamsNewFloatingTieredBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceMonthly   PriceNewParamsNewFloatingTieredBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceQuarterly PriceNewParamsNewFloatingTieredBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredBpsPriceCadenceOneTime   PriceNewParamsNewFloatingTieredBpsPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingTieredBpsPriceModelType string

const (
	PriceNewParamsNewFloatingTieredBpsPriceModelTypeTieredBps PriceNewParamsNewFloatingTieredBpsPriceModelType = "tiered_bps"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
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
	PriceNewParamsNewFloatingBpsPriceCadenceAnnual    PriceNewParamsNewFloatingBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingBpsPriceCadenceMonthly   PriceNewParamsNewFloatingBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingBpsPriceCadenceQuarterly PriceNewParamsNewFloatingBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBpsPriceCadenceOneTime   PriceNewParamsNewFloatingBpsPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingBpsPriceModelType string

const (
	PriceNewParamsNewFloatingBpsPriceModelTypeBps PriceNewParamsNewFloatingBpsPriceModelType = "bps"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
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
	PriceNewParamsNewFloatingBulkBpsPriceCadenceAnnual    PriceNewParamsNewFloatingBulkBpsPriceCadence = "annual"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceMonthly   PriceNewParamsNewFloatingBulkBpsPriceCadence = "monthly"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceQuarterly PriceNewParamsNewFloatingBulkBpsPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBulkBpsPriceCadenceOneTime   PriceNewParamsNewFloatingBulkBpsPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingBulkBpsPriceModelType string

const (
	PriceNewParamsNewFloatingBulkBpsPriceModelTypeBulkBps PriceNewParamsNewFloatingBulkBpsPriceModelType = "bulk_bps"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
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
	PriceNewParamsNewFloatingBulkPriceCadenceAnnual    PriceNewParamsNewFloatingBulkPriceCadence = "annual"
	PriceNewParamsNewFloatingBulkPriceCadenceMonthly   PriceNewParamsNewFloatingBulkPriceCadence = "monthly"
	PriceNewParamsNewFloatingBulkPriceCadenceQuarterly PriceNewParamsNewFloatingBulkPriceCadence = "quarterly"
	PriceNewParamsNewFloatingBulkPriceCadenceOneTime   PriceNewParamsNewFloatingBulkPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingBulkPriceModelType string

const (
	PriceNewParamsNewFloatingBulkPriceModelTypeBulk PriceNewParamsNewFloatingBulkPriceModelType = "bulk"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingThresholdTotalAmountPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceAnnual    PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "annual"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceMonthly   PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "monthly"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceQuarterly PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "quarterly"
	PriceNewParamsNewFloatingThresholdTotalAmountPriceCadenceOneTime   PriceNewParamsNewFloatingThresholdTotalAmountPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType string

const (
	PriceNewParamsNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount PriceNewParamsNewFloatingThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPackagePrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPackagePriceCadence string

const (
	PriceNewParamsNewFloatingTieredPackagePriceCadenceAnnual    PriceNewParamsNewFloatingTieredPackagePriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceMonthly   PriceNewParamsNewFloatingTieredPackagePriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceQuarterly PriceNewParamsNewFloatingTieredPackagePriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPackagePriceCadenceOneTime   PriceNewParamsNewFloatingTieredPackagePriceCadence = "one_time"
)

type PriceNewParamsNewFloatingTieredPackagePriceModelType string

const (
	PriceNewParamsNewFloatingTieredPackagePriceModelTypeTieredPackage PriceNewParamsNewFloatingTieredPackagePriceModelType = "tiered_package"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredWithMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredWithMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceAnnual    PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceMonthly   PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceQuarterly PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredWithMinimumPriceCadenceOneTime   PriceNewParamsNewFloatingTieredWithMinimumPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingTieredWithMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum PriceNewParamsNewFloatingTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingPackageWithAllocationPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingPackageWithAllocationPriceCadence string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceAnnual    PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "annual"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceMonthly   PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "monthly"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceQuarterly PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "quarterly"
	PriceNewParamsNewFloatingPackageWithAllocationPriceCadenceOneTime   PriceNewParamsNewFloatingPackageWithAllocationPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingPackageWithAllocationPriceModelType string

const (
	PriceNewParamsNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation PriceNewParamsNewFloatingPackageWithAllocationPriceModelType = "package_with_allocation"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingTieredPackageWithMinimumPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceAnnual    PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "annual"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceMonthly   PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "monthly"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "quarterly"
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadenceOneTime   PriceNewParamsNewFloatingTieredPackageWithMinimumPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType string

const (
	PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum PriceNewParamsNewFloatingTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

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
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
}

func (r PriceNewParamsNewFloatingUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (PriceNewParamsNewFloatingUnitWithPercentPrice) ImplementsPriceNewParams() {

}

// The cadence to bill for this price on.
type PriceNewParamsNewFloatingUnitWithPercentPriceCadence string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceAnnual    PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "annual"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceMonthly   PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "monthly"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceQuarterly PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "quarterly"
	PriceNewParamsNewFloatingUnitWithPercentPriceCadenceOneTime   PriceNewParamsNewFloatingUnitWithPercentPriceCadence = "one_time"
)

type PriceNewParamsNewFloatingUnitWithPercentPriceModelType string

const (
	PriceNewParamsNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent PriceNewParamsNewFloatingUnitWithPercentPriceModelType = "unit_with_percent"
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
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
