// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/apiquery"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// CustomerCostService contains methods and other services that help with
// interacting with the orb API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewCustomerCostService] method instead.
type CustomerCostService struct {
	Options []option.RequestOption
}

// NewCustomerCostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerCostService(opts ...option.RequestOption) (r *CustomerCostService) {
	r = &CustomerCostService{}
	r.Options = opts
	return
}

// This endpoint is used to fetch a day-by-day snapshot of a customer's costs in
// Orb, calculated by applying pricing information to the underlying usage (see the
// [subscription usage endpoint](fetch-subscription-usage.api.mdx) to fetch usage
// per metric, in usage units rather than a currency).
//
// This endpoint can be leveraged for internal tooling and to provide a more
// transparent billing experience for your end users:
//
//  1. Understand the cost breakdown per line item historically and in real-time for
//     the current billing period.
//  2. Provide customer visibility into how different services are contributing to
//     the overall invoice with a per-day timeseries (as compared to the
//     [upcoming invoice](fetch-upcoming-invoice) resource, which represents a
//     snapshot for the current period).
//  3. Assess how minimums and discounts affect your customers by teasing apart
//     costs directly as a result of usage, as opposed to minimums and discounts at
//     the plan and price level.
//  4. Gain insight into key customer health metrics, such as the percent
//     utilization of the minimum committed spend.
//
// ## Fetching subscriptions
//
// By default, this endpoint fetches the currently active subscription for the
// customer, and returns cost information for the subscription's current billing
// period, broken down by each participating price. If there are no currently
// active subscriptions, this will instead default to the most recently active
// subscription or return an empty series if none are found. For example, if your
// plan charges for compute hours, job runs, and data syncs, then this endpoint
// would provide a daily breakdown of your customer's cost for each of those axes.
//
// If timeframe bounds are specified, Orb fetches all subscriptions that were
// active in that timeframe. If two subscriptions overlap on a single day, costs
// from each price will be summed, and prices for both subscriptions will be
// included in the breakdown.
//
// ## Prepaid plans
//
// For plans that include prices which deduct credits rather than accrue in-arrears
// charges in a billable currency, this endpoint will return the total deduction
// amount, in credits, for the specified timeframe.
//
// ## Cumulative subtotals and totals
//
// Since the subtotal and total must factor in any billing-period level discounts
// and minimums, it's most meaningful to consider costs relative to the start of
// the subscription's billing period. As a result, by default this endpoint returns
// cumulative totals since the beginning of the billing period. In particular, the
// `timeframe_start` of a returned timeframe window is _always_ the beginning of
// the billing period and `timeframe_end` is incremented one day at a time to build
// the result.
//
// A customer that uses a few API calls a day but has a minimum commitment might
// exhibit the following pattern for their subtotal and total in the first few days
// of the month. Here, we assume that each API call is $2.50, the customer's plan
// has a monthly minimum of $50 for this price, and that the subscription's billing
// period bounds are aligned to the first of the month:
//
// | timeframe_start | timeframe_end | Cumulative usage | Subtotal | Total (incl. commitment) |
// | --------------- | ------------- | ---------------- | -------- | ------------------------ |
// | 2023-02-01      | 2023-02-02    | 9                | $22.50   | $50.00                   |
// | 2023-02-01      | 2023-02-03    | 19               | $47.50   | $50.00                   |
// | 2023-02-01      | 2023-02-04    | 20               | $50.00   | $50.00                   |
// | 2023-02-01      | 2023-02-05    | 28               | $70.00   | $70.00                   |
// | 2023-02-01      | 2023-02-06    | 36               | $90.00   | $90.00                   |
//
// ### Periodic values
//
// When the query parameter `view_mode=periodic` is specified, Orb will return an
// incremental day-by-day view of costs. In this case, there will always be a
// one-day difference between `timeframe_start` and `timeframe_end` for the
// timeframes returned. This is a transform on top of the cumulative costs,
// calculated by taking the difference of each timeframe with the last. Note that
// in the above example, the `Total` value would be 0 for the second two data
// points, since the minimum commitment has not yet been hit and each day is not
// contributing anything to the total cost.
//
// ## Timeframe bounds
//
// If no timeframe bounds are specified, the response will default to the current
// billing period for the customer's subscription. For subscriptions that have
// ended, this will be the billing period when they were last active. If the
// subscription starts or ends within the timeframe, the response will only include
// windows where the subscription is active.
//
// As noted above, `timeframe_start` for a given cumulative datapoint is always the
// beginning of the billing period, and `timeframe_end` is incremented one day at a
// time to construct the response. When a timeframe is passed in that is not
// aligned to the current subscription's billing period, the response will contain
// cumulative totals from multiple billing periods.
//
// Suppose the queried customer has a subscription aligned to the 15th of every
// month. If this endpoint is queried with the date range `2023-06-01` -
// `2023-07-01`, the first data point will represent about half a billing period's
// worth of costs, accounting for accruals from the start of the billing period and
// inclusive of the first day of the timeframe
// (`timeframe_start = 2023-05-15 00:00:00`, `timeframe_end = 2023-06-02 00:00:00`)
//
// | datapoint index | timeframe_start | timeframe_end |
// | --------------- | --------------- | ------------- |
// | 0               | 2023-05-15      | 2023-06-02    |
// | 1               | 2023-05-15      | 2023-06-03    |
// | 2               | ...             | ...           |
// | 3               | 2023-05-15      | 2023-06-14    |
// | 4               | 2023-06-15      | 2023-06-16    |
// | 5               | 2023-06-15      | 2023-06-17    |
// | 6               | ...             | ...           |
// | 7               | 2023-06-15      | 2023-07-01    |
//
// You can see this sliced timeframe visualized
// [here](https://i.imgur.com/TXhYgme.png).
//
// ## Grouping by custom attributes
//
// In order to view costs grouped by a specific _attribute_ that each event is
// tagged with (i.e. `cluster`), you can additionally specify a `group_by` key. The
// `group_by` key denotes the event property on which to group.
//
// When returning grouped costs, a separate `price_group` object in the
// `per_price_costs` array is returned for each value of the `group_by` key present
// in your events. The `subtotal` value of the `per_price_costs` object is the sum
// of each `price_group`'s total.
//
// Orb expects events will contain values in the `properties` dictionary that
// correspond to the `group_by` key specified. By default, Orb will return a `null`
// group (i.e. events that match the metric but do not have the key set).
// Currently, it is only possible to view costs grouped by a single attribute at a
// time.
//
// ### Matrix prices
//
// When a price uses matrix pricing, it's important to view costs grouped by those
// matrix dimensions. Orb will return `price_groups` with the `grouping_key` and
// `secondary_grouping_key` based on the matrix price definition, for each
// `grouping_value` and `secondary_grouping_value` available.
func (r *CustomerCostService) List(ctx context.Context, customerID string, query CustomerCostListParams, opts ...option.RequestOption) (res *CustomerCostListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/%s/costs", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is used to fetch a day-by-day snapshot of a customer's costs in
// Orb, calculated by applying pricing information to the underlying usage (see the
// [subscription usage endpoint](fetch-subscription-usage.api.mdx) to fetch usage
// per metric, in usage units rather than a currency).
//
// This endpoint can be leveraged for internal tooling and to provide a more
// transparent billing experience for your end users:
//
//  1. Understand the cost breakdown per line item historically and in real-time for
//     the current billing period.
//  2. Provide customer visibility into how different services are contributing to
//     the overall invoice with a per-day timeseries (as compared to the
//     [upcoming invoice](fetch-upcoming-invoice) resource, which represents a
//     snapshot for the current period).
//  3. Assess how minimums and discounts affect your customers by teasing apart
//     costs directly as a result of usage, as opposed to minimums and discounts at
//     the plan and price level.
//  4. Gain insight into key customer health metrics, such as the percent
//     utilization of the minimum committed spend.
//
// ## Fetching subscriptions
//
// By default, this endpoint fetches the currently active subscription for the
// customer, and returns cost information for the subscription's current billing
// period, broken down by each participating price. If there are no currently
// active subscriptions, this will instead default to the most recently active
// subscription or return an empty series if none are found. For example, if your
// plan charges for compute hours, job runs, and data syncs, then this endpoint
// would provide a daily breakdown of your customer's cost for each of those axes.
//
// If timeframe bounds are specified, Orb fetches all subscriptions that were
// active in that timeframe. If two subscriptions overlap on a single day, costs
// from each price will be summed, and prices for both subscriptions will be
// included in the breakdown.
//
// ## Prepaid plans
//
// For plans that include prices which deduct credits rather than accrue in-arrears
// charges in a billable currency, this endpoint will return the total deduction
// amount, in credits, for the specified timeframe.
//
// ## Cumulative subtotals and totals
//
// Since the subtotal and total must factor in any billing-period level discounts
// and minimums, it's most meaningful to consider costs relative to the start of
// the subscription's billing period. As a result, by default this endpoint returns
// cumulative totals since the beginning of the billing period. In particular, the
// `timeframe_start` of a returned timeframe window is _always_ the beginning of
// the billing period and `timeframe_end` is incremented one day at a time to build
// the result.
//
// A customer that uses a few API calls a day but has a minimum commitment might
// exhibit the following pattern for their subtotal and total in the first few days
// of the month. Here, we assume that each API call is $2.50, the customer's plan
// has a monthly minimum of $50 for this price, and that the subscription's billing
// period bounds are aligned to the first of the month:
//
// | timeframe_start | timeframe_end | Cumulative usage | Subtotal | Total (incl. commitment) |
// | --------------- | ------------- | ---------------- | -------- | ------------------------ |
// | 2023-02-01      | 2023-02-02    | 9                | $22.50   | $50.00                   |
// | 2023-02-01      | 2023-02-03    | 19               | $47.50   | $50.00                   |
// | 2023-02-01      | 2023-02-04    | 20               | $50.00   | $50.00                   |
// | 2023-02-01      | 2023-02-05    | 28               | $70.00   | $70.00                   |
// | 2023-02-01      | 2023-02-06    | 36               | $90.00   | $90.00                   |
//
// ### Periodic values
//
// When the query parameter `view_mode=periodic` is specified, Orb will return an
// incremental day-by-day view of costs. In this case, there will always be a
// one-day difference between `timeframe_start` and `timeframe_end` for the
// timeframes returned. This is a transform on top of the cumulative costs,
// calculated by taking the difference of each timeframe with the last. Note that
// in the above example, the `Total` value would be 0 for the second two data
// points, since the minimum commitment has not yet been hit and each day is not
// contributing anything to the total cost.
//
// ## Timeframe bounds
//
// If no timeframe bounds are specified, the response will default to the current
// billing period for the customer's subscription. For subscriptions that have
// ended, this will be the billing period when they were last active. If the
// subscription starts or ends within the timeframe, the response will only include
// windows where the subscription is active.
//
// As noted above, `timeframe_start` for a given cumulative datapoint is always the
// beginning of the billing period, and `timeframe_end` is incremented one day at a
// time to construct the response. When a timeframe is passed in that is not
// aligned to the current subscription's billing period, the response will contain
// cumulative totals from multiple billing periods.
//
// Suppose the queried customer has a subscription aligned to the 15th of every
// month. If this endpoint is queried with the date range `2023-06-01` -
// `2023-07-01`, the first data point will represent about half a billing period's
// worth of costs, accounting for accruals from the start of the billing period and
// inclusive of the first day of the timeframe
// (`timeframe_start = 2023-05-15 00:00:00`, `timeframe_end = 2023-06-02 00:00:00`)
//
// | datapoint index | timeframe_start | timeframe_end |
// | --------------- | --------------- | ------------- |
// | 0               | 2023-05-15      | 2023-06-02    |
// | 1               | 2023-05-15      | 2023-06-03    |
// | 2               | ...             | ...           |
// | 3               | 2023-05-15      | 2023-06-14    |
// | 4               | 2023-06-15      | 2023-06-16    |
// | 5               | 2023-06-15      | 2023-06-17    |
// | 6               | ...             | ...           |
// | 7               | 2023-06-15      | 2023-07-01    |
//
// You can see this sliced timeframe visualized
// [here](https://i.imgur.com/TXhYgme.png).
//
// ## Grouping by custom attributes
//
// In order to view costs grouped by a specific _attribute_ that each event is
// tagged with (i.e. `cluster`), you can additionally specify a `group_by` key. The
// `group_by` key denotes the event property on which to group.
//
// When returning grouped costs, a separate `price_group` object in the
// `per_price_costs` array is returned for each value of the `group_by` key present
// in your events. The `subtotal` value of the `per_price_costs` object is the sum
// of each `price_group`'s total.
//
// Orb expects events will contain values in the `properties` dictionary that
// correspond to the `group_by` key specified. By default, Orb will return a `null`
// group (i.e. events that match the metric but do not have the key set).
// Currently, it is only possible to view costs grouped by a single attribute at a
// time.
//
// ### Matrix prices
//
// When a price uses matrix pricing, it's important to view costs grouped by those
// matrix dimensions. Orb will return `price_groups` with the `grouping_key` and
// `secondary_grouping_key` based on the matrix price definition, for each
// `grouping_value` and `secondary_grouping_value` available.
func (r *CustomerCostService) ListByExternalID(ctx context.Context, externalCustomerID string, query CustomerCostListByExternalIDParams, opts ...option.RequestOption) (res *CustomerCostListByExternalIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("customers/external_customer_id/%s/costs", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CustomerCostListResponse struct {
	Data []CustomerCostListResponseData `json:"data,required"`
	JSON customerCostListResponseJSON   `json:"-"`
}

// customerCostListResponseJSON contains the JSON metadata for the struct
// [CustomerCostListResponse]
type customerCostListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCostListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListResponseData struct {
	PerPriceCosts []CustomerCostListResponseDataPerPriceCost `json:"per_price_costs,required"`
	// Total costs for the timeframe, excluding any minimums and discounts.
	Subtotal       string    `json:"subtotal,required"`
	TimeframeEnd   time.Time `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time `json:"timeframe_start,required" format:"date-time"`
	// Total costs for the timeframe, including any minimums and discounts.
	Total string                           `json:"total,required"`
	JSON  customerCostListResponseDataJSON `json:"-"`
}

// customerCostListResponseDataJSON contains the JSON metadata for the struct
// [CustomerCostListResponseData]
type customerCostListResponseDataJSON struct {
	PerPriceCosts  apijson.Field
	Subtotal       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCostListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListResponseDataPerPriceCost struct {
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
	Price Price `json:"price,required"`
	// Price's contributions for the timeframe, excluding any minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Price's contributions for the timeframe, including minimums and discounts.
	Total string `json:"total,required"`
	// The price's quantity for the timeframe
	Quantity float64                                      `json:"quantity,nullable"`
	JSON     customerCostListResponseDataPerPriceCostJSON `json:"-"`
}

// customerCostListResponseDataPerPriceCostJSON contains the JSON metadata for the
// struct [CustomerCostListResponseDataPerPriceCost]
type customerCostListResponseDataPerPriceCostJSON struct {
	Price       apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCostListResponseDataPerPriceCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListResponseDataPerPriceCostJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListByExternalIDResponse struct {
	Data []CustomerCostListByExternalIDResponseData `json:"data,required"`
	JSON customerCostListByExternalIDResponseJSON   `json:"-"`
}

// customerCostListByExternalIDResponseJSON contains the JSON metadata for the
// struct [CustomerCostListByExternalIDResponse]
type customerCostListByExternalIDResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCostListByExternalIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListByExternalIDResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListByExternalIDResponseData struct {
	PerPriceCosts []CustomerCostListByExternalIDResponseDataPerPriceCost `json:"per_price_costs,required"`
	// Total costs for the timeframe, excluding any minimums and discounts.
	Subtotal       string    `json:"subtotal,required"`
	TimeframeEnd   time.Time `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time `json:"timeframe_start,required" format:"date-time"`
	// Total costs for the timeframe, including any minimums and discounts.
	Total string                                       `json:"total,required"`
	JSON  customerCostListByExternalIDResponseDataJSON `json:"-"`
}

// customerCostListByExternalIDResponseDataJSON contains the JSON metadata for the
// struct [CustomerCostListByExternalIDResponseData]
type customerCostListByExternalIDResponseDataJSON struct {
	PerPriceCosts  apijson.Field
	Subtotal       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CustomerCostListByExternalIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListByExternalIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListByExternalIDResponseDataPerPriceCost struct {
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
	Price Price `json:"price,required"`
	// Price's contributions for the timeframe, excluding any minimums and discounts.
	Subtotal string `json:"subtotal,required"`
	// Price's contributions for the timeframe, including minimums and discounts.
	Total string `json:"total,required"`
	// The price's quantity for the timeframe
	Quantity float64                                                  `json:"quantity,nullable"`
	JSON     customerCostListByExternalIDResponseDataPerPriceCostJSON `json:"-"`
}

// customerCostListByExternalIDResponseDataPerPriceCostJSON contains the JSON
// metadata for the struct [CustomerCostListByExternalIDResponseDataPerPriceCost]
type customerCostListByExternalIDResponseDataPerPriceCostJSON struct {
	Price       apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomerCostListByExternalIDResponseDataPerPriceCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerCostListByExternalIDResponseDataPerPriceCostJSON) RawJSON() string {
	return r.raw
}

type CustomerCostListParams struct {
	// Costs returned are exclusive of `timeframe_end`.
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// Costs returned are inclusive of `timeframe_start`.
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
	// Controls whether Orb returns cumulative costs since the start of the billing
	// period, or incremental day-by-day costs. If your customer has minimums or
	// discounts, it's strongly recommended that you use the default cumulative
	// behavior.
	ViewMode param.Field[CustomerCostListParamsViewMode] `query:"view_mode"`
}

// URLQuery serializes [CustomerCostListParams]'s query parameters as `url.Values`.
func (r CustomerCostListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls whether Orb returns cumulative costs since the start of the billing
// period, or incremental day-by-day costs. If your customer has minimums or
// discounts, it's strongly recommended that you use the default cumulative
// behavior.
type CustomerCostListParamsViewMode string

const (
	CustomerCostListParamsViewModePeriodic   CustomerCostListParamsViewMode = "periodic"
	CustomerCostListParamsViewModeCumulative CustomerCostListParamsViewMode = "cumulative"
)

func (r CustomerCostListParamsViewMode) IsKnown() bool {
	switch r {
	case CustomerCostListParamsViewModePeriodic, CustomerCostListParamsViewModeCumulative:
		return true
	}
	return false
}

type CustomerCostListByExternalIDParams struct {
	// Costs returned are exclusive of `timeframe_end`.
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// Costs returned are inclusive of `timeframe_start`.
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
	// Controls whether Orb returns cumulative costs since the start of the billing
	// period, or incremental day-by-day costs. If your customer has minimums or
	// discounts, it's strongly recommended that you use the default cumulative
	// behavior.
	ViewMode param.Field[CustomerCostListByExternalIDParamsViewMode] `query:"view_mode"`
}

// URLQuery serializes [CustomerCostListByExternalIDParams]'s query parameters as
// `url.Values`.
func (r CustomerCostListByExternalIDParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls whether Orb returns cumulative costs since the start of the billing
// period, or incremental day-by-day costs. If your customer has minimums or
// discounts, it's strongly recommended that you use the default cumulative
// behavior.
type CustomerCostListByExternalIDParamsViewMode string

const (
	CustomerCostListByExternalIDParamsViewModePeriodic   CustomerCostListByExternalIDParamsViewMode = "periodic"
	CustomerCostListByExternalIDParamsViewModeCumulative CustomerCostListByExternalIDParamsViewMode = "cumulative"
)

func (r CustomerCostListByExternalIDParamsViewMode) IsKnown() bool {
	switch r {
	case CustomerCostListByExternalIDParamsViewModePeriodic, CustomerCostListByExternalIDParamsViewModeCumulative:
		return true
	}
	return false
}
