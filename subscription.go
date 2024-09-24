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
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/orbcorp/orb-go/shared"
	"github.com/tidwall/gjson"
)

// SubscriptionService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// A subscription represents the purchase of a plan by a customer. The customer is
// identified by either the `customer_id` or the `external_customer_id`, and
// exactly one of these fields must be provided.
//
// By default, subscriptions begin on the day that they're created and renew
// automatically for each billing cycle at the cadence that's configured in the
// plan definition.
//
// The default configuration for subscriptions in Orb is **In-advance billing** and
// **Beginning of month alignment** (see
// [Subscription](../guides/concepts#subscription) for more details).
//
// In order to change the alignment behavior, Orb also supports billing
// subscriptions on the day of the month they are created. If
// `align_billing_with_subscription_start_date = true` is specified, subscriptions
// have billing cycles that are aligned with their `start_date`. For example, a
// subscription that begins on January 15th will have a billing cycle from January
// 15th to February 15th. Every subsequent billing cycle will continue to start and
// invoice on the 15th.
//
// If the "day" value is greater than the number of days in the month, the next
// billing cycle will start at the end of the month. For example, if the start_date
// is January 31st, the next billing cycle will start on February 28th.
//
// If a customer was created with a currency, Orb only allows subscribing the
// customer to a plan with a matching `invoicing_currency`. If the customer does
// not have a currency set, on subscription creation, we set the customer's
// currency to be the `invoicing_currency` of the plan.
//
// ## Price overrides
//
// Price overrides are used to update some or all prices in a plan for the specific
// subscription being created. This is useful when a new customer has negotiated
// one or more different prices for a specific plan than the plan's default prices.
// Any type of price can be overridden, if the correct data is provided. The
// billable metric, cadence, type, and name of a price can not be overridden.
//
// To override prices, provide a list of objects with the key `price_overrides`.
// The price object in the list of overrides is expected to contain the existing
// price id, the `model_type` and config value in the format below. The specific
// numerical values can be updated, but the config value and `model_type` must
// match the existing price that is being overridden
//
// ### Request format for price overrides
//
// Orb supports a few different pricing models out of the box. The `model_type`
// field determines the key for the configuration object that is present.
//
// ### Unit pricing
//
// With unit pricing, each unit costs a fixed amount.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "unit",
//	  "unit_config": {
//	    "unit_amount": "0.50"
//	  }
//	  ...
//	}
//
// ```
//
// ### Tiered pricing
//
// In tiered pricing, the cost of a given unit depends on the tier range that it
// falls into, where each tier range is defined by an upper and lower bound. For
// example, the first ten units may cost $0.50 each and all units thereafter may
// cost $0.10 each. Tiered prices can be overridden with a new number of tiers or
// new values for `first_unit`, `last_unit`, or `unit_amount`.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "tiered",
//	  "tiered_config": {
//	    "tiers": [
//	      {
//	        "first_unit":"1",
//	        "last_unit": "11",
//	        "unit_amount": "0.50"
//	      },
//	      {
//	        "first_unit": "11",
//	        "last_unit": null,
//	        "unit_amount": "0.10"
//	      }
//	    ]
//	  }
//	  ...
//	}
//
// ```
//
// ### Bulk pricing
//
// Bulk pricing applies when the number of units determine the cost of _all_ units.
// For example, if you've bought less than 10 units, they may each be $0.50 for a
// total of $5.00. Once you've bought more than 10 units, all units may now be
// priced at $0.40 (i.e. 101 units total would be $40.40). Bulk prices can be
// overridden with a new number of tiers or new values for `maximum_units`, or
// `unit_amount`.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "bulk",
//	  "bulk_config": {
//	    "tiers": [
//	      {
//	        "maximum_units": "10",
//	        "unit_amount": "0.50"
//	      },
//	      {
//	        "maximum_units": "1000",
//	        "unit_amount": "0.40"
//	      }
//	    ]
//	  }
//	  ...
//	}
//
// ```
//
// ### Package pricing
//
// Package pricing defines the size or granularity of a unit for billing purposes.
// For example, if the package size is set to 5, then 4 units will be billed as 5
// and 6 units will be billed at 10.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "package",
//	  "package_config": {
//	    "package_amount": "0.80",
//	    "package_size": 10
//	  }
//	  ...
//	}
//
// ```
//
// ### BPS pricing
//
// BPS pricing specifies a per-event (e.g. per-payment) rate in one hundredth of a
// percent (the number of basis points to charge), as well as a cap per event to
// assess. For example, this would allow you to assess a fee of 0.25% on every
// payment you process, with a maximum charge of $25 per payment.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id"
//	  "model_type": "bps",
//	  "bps_config": {
//	    "bps": 125,
//	    "per_event_cap": "11.00"
//	  }
//	  ...
//	}
//
// ```
//
// ### Bulk BPS pricing
//
// Bulk BPS pricing specifies BPS parameters in a tiered manner, dependent on the
// total quantity across all events. Similar to bulk pricing, the BPS parameters of
// a given event depends on the tier range that the billing period falls into. Each
// tier range is defined by an upper and lower bound. For example, after $1.5M of
// payment volume is reached, each individual payment may have a lower cap or a
// smaller take-rate.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id"
//	  "model_type": "bulk_bps",
//	  "bulk_bps_config": {
//	    "tiers": [
//	      {
//	        "minimum_amount": "0.00",
//	        "maximum_amount": "1000000.00",
//	        "bps": 125,
//	        "per_event_cap": "19.00"
//	      },
//	      {
//	        "minimum_amount":"1000000.00",
//	        "maximum_amount": null,
//	        "bps": 115,
//	        "per_event_cap": "4.00"
//	      }
//	    ]
//	  }
//
// ...
// }
// ```
//
// ### Tiered BPS pricing
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
//	{
//	  ...
//	  "id": "price_id"
//	  "model_type": "tiered_bps",
//	  "tiered_bps_config": {
//	    "tiers": [
//	      {
//	        "minimum_amount": "0.00",
//	        "maximum_amount": "1000000.00",
//	        "bps": 125,
//	        "per_event_cap": "19.00"
//	      },
//	      {
//	        "minimum_amount":"1000000",
//	        "maximum_amount": null,
//	        "bps": 115,
//	        "per_event_cap": "4.00"
//	      }
//	    ]
//	  }
//	  ...
//	}
//
// ```
//
// ### Matrix pricing
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
//	  ...
//	  "model_type": "matrix",
//	  "matrix_config": {
//	    "default_unit_amount": "3.00",
//	    "dimensions": [
//	      "cluster_name",
//	      "region"
//	    ],
//	    "matrix_values": [
//	      {
//	        "dimension_values": [
//	          "alpha",
//	          "west"
//	        ],
//	        "unit_amount": "2.00"
//	      },
//	      ...
//	    ]
//	  }
//	}
//
// ```
//
// ### Fixed fees
//
// Fixed fees follow unit pricing, and also have an additional parameter
// `fixed_price_quantity` that indicates how many of a fixed fee that should be
// applied for a subscription. This parameter defaults to 1.
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "unit",
//	  "unit_config": {
//	    "unit_amount": "2.00"
//	  },
//	  "fixed_price_quantity": 3.0
//	  ...
//	}
//
// ```
//
// ## Maximums and Minimums
//
// Minimums and maximums, much like price overrides, can be useful when a new
// customer has negotiated a new or different minimum or maximum spend cap than the
// default for a given price. If one exists for a price and null is provided for
// the minimum/maximum override on creation, then there will be no minimum/maximum
// on the new subscription. If no value is provided, then the default price maximum
// or minimum is used.
//
// To add a minimum for a specific price, add `minimum_amount` to the specific
// price in the `price_overrides` object.
//
// To add a maximum for a specific price, add `maximum_amount` to the specific
// price in the `price_overrides` object.
//
// ### Minimum override example
//
// Price minimum override example:
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "unit",
//	  "unit_config": {
//	    "unit_amount": "0.50"
//	  },
//	  "minimum_amount": "100.00"
//	  ...
//	}
//
// ```
//
// # Removing an existing minimum example
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "unit",
//	  "unit_config": {
//	    "unit_amount": "0.50"
//	  },
//	  "minimum_amount": null
//	  ...
//	}
//
// ```
//
// ## Discounts
//
// Discounts, like price overrides, can be useful when a new customer has
// negotiated a new or different discount than the default for a price. A single
// price price can have at most one discount. If a discount exists for a price and
// a null discount is provided on creation, then there will be no discount on the
// new subscription.
//
// To add a discount for a specific price, add `discount` to the price in the
// `price_overrides` object. Discount should be a dictionary of the format:
//
// ```ts
//
//	{
//	  "discount_type": "amount" | "percentage" | "usage",
//	  "amount_discount": string,
//	  "percentage_discount": string,
//	  "usage_discount": string
//	}
//
// ```
//
// where either `amount_discount`, `percentage_discount`, or `usage_discount` is
// provided.
//
// # Price discount example
//
// ```json
//
//	{
//	  ...
//	  "id": "price_id",
//	  "model_type": "unit",
//	  "unit_config": {
//	    "unit_amount": "0.50"
//	  },
//	  "discount": {"discount_type": "amount", "amount_discount": "175"},
//	}
//
// ```
//
// # Removing an existing discount example
//
// ```json
//
//	{
//	  "customer_id": "customer_id",
//	  "plan_id": "plan_id",
//	  "discount": null,
//	  "price_overrides": [ ... ]
//	  ...
//	}
//
// ```
//
// ## Threshold Billing
//
// Orb supports invoicing for a subscription when a preconfigured usage threshold
// is hit. To enable threshold billing, pass in an `invoicing_threshold`, which is
// specified in the subscription's invoicing currency, when creating a
// subscription. E.g. pass in `10.00` to issue an invoice when usage amounts hit
// $10.00 for a subscription that invoices in USD.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `metadata`, `net terms`,
// `auto_collection`, `invoicing_threshold`, and `default_invoice_memo` properties
// on a subscription.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// This endpoint returns a list of all subscriptions for an account as a
// [paginated](../reference/pagination) list, ordered starting from the most
// recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](../guides/concepts#subscription).
//
// Subscriptions can be filtered for a specific customer by using either the
// customer_id or external_customer_id query parameters. To filter subscriptions
// for multiple customers, use the customer_id[] or external_customer_id[] query
// parameters.
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[Subscription], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "subscriptions"
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

// This endpoint returns a list of all subscriptions for an account as a
// [paginated](../reference/pagination) list, ordered starting from the most
// recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](../guides/concepts#subscription).
//
// Subscriptions can be filtered for a specific customer by using either the
// customer_id or external_customer_id query parameters. To filter subscriptions
// for multiple customers, use the customer_id[] or external_customer_id[] query
// parameters.
func (r *SubscriptionService) ListAutoPaging(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Subscription] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint can be used to cancel an existing subscription. It returns the
// serialized subscription object with an `end_date` parameter that signifies when
// the subscription will transition to an ended state.
//
// The body parameter `cancel_option` determines the cancellation behavior. Orb
// supports three cancellation options:
//
//   - `end_of_subscription_term`: stops the subscription from auto-renewing.
//     Subscriptions that have been cancelled with this option can still incur
//     charges for the remainder of their term:
//
//   - Issuing this cancellation request for a monthly subscription will keep the
//     subscription active until the start of the subsequent month, and potentially
//     issue an invoice for any usage charges incurred in the intervening period.
//
//   - Issuing this cancellation request for a quarterly subscription will keep the
//     subscription active until the end of the quarter and potentially issue an
//     invoice for any usage charges incurred in the intervening period.
//
//   - Issuing this cancellation request for a yearly subscription will keep the
//     subscription active for the full year. For example, a yearly subscription
//     starting on 2021-11-01 and cancelled on 2021-12-08 will remain active until
//     2022-11-01 and potentially issue charges in the intervening months for any
//     recurring monthly usage charges in its plan.
//
//   - **Note**: If a subscription's plan contains prices with difference cadences,
//     the end of term date will be determined by the largest cadence value. For
//     example, cancelling end of term for a subscription with a quarterly fixed
//     fee with a monthly usage fee will result in the subscription ending at the
//     end of the quarter.
//
//   - `immediate`: ends the subscription immediately, setting the `end_date` to the
//     current time:
//
//   - Subscriptions that have been cancelled with this option will be invoiced
//     immediately. This invoice will include any usage fees incurred in the
//     billing period up to the cancellation, along with any prorated recurring
//     fees for the billing period, if applicable.
//
//   - **Note**: If the subscription has a recurring fee that was paid in-advance,
//     the prorated amount for the remaining time period will be added to the
//     [customer's balance](list-balance-transactions) upon immediate cancellation.
//     However, if the customer is ineligible to use the customer balance, the
//     subscription cannot be cancelled immediately.
//
//   - `requested_date`: ends the subscription on a specified date, which requires a
//     `cancellation_date` to be passed in. If no timezone is provided, the
//     customer's timezone is used. For example, a subscription starting on January
//     1st with a monthly price can be set to be cancelled on the first of any month
//     after January 1st (e.g. March 1st, April 1st, May 1st). A subscription with
//     multiple prices with different cadences defines the "term" to be the highest
//     cadence of the prices.
//
// Upcoming subscriptions are only eligible for immediate cancellation, which will
// set the `end_date` equal to the `start_date` upon cancellation.
//
// ## Backdated cancellations
//
// Orb allows you to cancel a subscription in the past as long as there are no paid
// invoices between the `requested_date` and the current time. If the cancellation
// is after the latest issued invoice, Orb will generate a balance refund for the
// current period. If the cancellation is before the most recently issued invoice,
// Orb will void the intervening invoice and generate a new one based on the new
// dates for the subscription. See the section on
// [cancellation behaviors](../guides/product-catalog/creating-subscriptions.md#cancellation-behaviors).
func (r *SubscriptionService) Cancel(ctx context.Context, subscriptionID string, body SubscriptionCancelParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/cancel", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch a [Subscription](../guides/concepts#subscription)
// given an identifier.
func (r *SubscriptionService) Fetch(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint is used to fetch a day-by-day snapshot of a subscription's costs
// in Orb, calculated by applying pricing information to the underlying usage (see
// the [subscription usage endpoint](fetch-subscription-usage) to fetch usage per
// metric, in usage units rather than a currency).
//
// The semantics of this endpoint exactly mirror those of
// [fetching a customer's costs](fetch-customer-costs). Use this endpoint to limit
// your analysis of costs to a specific subscription for the customer (e.g. to
// de-aggregate costs when a customer's subscription has started and stopped on the
// same day).
func (r *SubscriptionService) FetchCosts(ctx context.Context, subscriptionID string, query SubscriptionFetchCostsParams, opts ...option.RequestOption) (res *SubscriptionFetchCostsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/costs", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint returns a [paginated](../reference/pagination) list of all plans
// associated with a subscription along with their start and end dates. This list
// contains the subscription's initial plan along with past and future plan
// changes.
func (r *SubscriptionService) FetchSchedule(ctx context.Context, subscriptionID string, query SubscriptionFetchScheduleParams, opts ...option.RequestOption) (res *pagination.Page[SubscriptionFetchScheduleResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/schedule", subscriptionID)
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

// This endpoint returns a [paginated](../reference/pagination) list of all plans
// associated with a subscription along with their start and end dates. This list
// contains the subscription's initial plan along with past and future plan
// changes.
func (r *SubscriptionService) FetchScheduleAutoPaging(ctx context.Context, subscriptionID string, query SubscriptionFetchScheduleParams, opts ...option.RequestOption) *pagination.PageAutoPager[SubscriptionFetchScheduleResponse] {
	return pagination.NewPageAutoPager(r.FetchSchedule(ctx, subscriptionID, query, opts...))
}

// This endpoint is used to fetch a subscription's usage in Orb. Especially when
// combined with optional query parameters, this endpoint is a powerful way to
// build visualizations on top of Orb's event data and metrics.
//
// With no query parameters specified, this endpoint returns usage for the
// subscription's _current billing period_ across each billable metric that
// participates in the subscription. Usage quantities returned are the result of
// evaluating the metric definition for the entirety of the customer's billing
// period.
//
// ### Default response shape
//
// Orb returns a `data` array with an object corresponding to each billable metric.
// Nested within this object is a `usage` array which has a `quantity` value and a
// corresponding `timeframe_start` and `timeframe_end`. The `quantity` value
// represents the calculated usage value for the billable metric over the specified
// timeframe (inclusive of the `timeframe_start` timestamp and exclusive of the
// `timeframe_end` timestamp).
//
// Orb will include _every_ window in the response starting from the beginning of
// the billing period, even when there were no events (and therefore no usage) in
// the window. This increases the size of the response but prevents the caller from
// filling in gaps and handling cumbersome time-based logic.
//
// The query parameters in this endpoint serve to override this behavior and
// provide some key functionality, as listed below. Note that this functionality
// can also be used _in conjunction_ with each other, e.g. to display grouped usage
// on a custom timeframe.
//
// ## Custom timeframe
//
// In order to view usage for a custom timeframe rather than the current billing
// period, specify a `timeframe_start` and `timeframe_end`. This will calculate
// quantities for usage incurred between timeframe_start (inclusive) and
// timeframe_end (exclusive), i.e. `[timeframe_start, timeframe_end)`.
//
// Note:
//
//   - These timestamps must be specified in ISO 8601 format and UTC timezone, e.g.
//     `2022-02-01T05:00:00Z`.
//   - Both parameters must be specified if either is specified.
//
// ## Grouping by custom attributes
//
// In order to view a single metric grouped by a specific _attribute_ that each
// event is tagged with (e.g. `cluster`), you must additionally specify a
// `billable_metric_id` and a `group_by` key. The `group_by` key denotes the event
// property on which to group.
//
// When returning grouped usage, only usage for `billable_metric_id` is returned,
// and a separate object in the `data` array is returned for each value of the
// `group_by` key present in your events. The `quantity` value is the result of
// evaluating the billable metric for events filtered to a single value of the
// `group_by` key.
//
// Orb expects that events that match the billable metric will contain values in
// the `properties` dictionary that correspond to the `group_by` key specified. By
// default, Orb will not return a `null` group (i.e. events that match the metric
// but do not have the key set). Currently, it is only possible to view usage
// grouped by a single attribute at a time.
//
// When viewing grouped usage, Orb uses pagination to limit the response size to
// 1000 groups by default. If there are more groups for a given subscription,
// pagination metadata in the response can be used to fetch all of the data.
//
// The following example shows usage for an "API Requests" billable metric grouped
// by `region`. Note the extra `metric_group` dictionary in the response, which
// provides metadata about the group:
//
// ```json
//
//	{
//	    "data": [
//	        {
//	            "usage": [
//	                {
//	                    "quantity": 0.19291,
//	                    "timeframe_start": "2021-10-01T07:00:00Z",
//	                    "timeframe_end": "2021-10-02T07:00:00Z",
//	                },
//	                ...
//	            ],
//	            "metric_group": {
//	                "property_key": "region",
//	                "property_value": "asia/pacific"
//	            },
//	            "billable_metric": {
//	                "id": "Fe9pbpMk86xpwdGB",
//	                "name": "API Requests"
//	            },
//	            "view_mode": "periodic"
//	        },
//	        ...
//	    ]
//	}
//
// ```
//
// ## Windowed usage
//
// The `granularity` parameter can be used to _window_ the usage `quantity` value
// into periods. When not specified, usage is returned for the entirety of the time
// range.
//
// When `granularity = day` is specified with a timeframe longer than a day, Orb
// will return a `quantity` value for each full day between `timeframe_start` and
// `timeframe_end`. Note that the days are demarcated by the _customer's local
// midnight_.
//
// For example, with `timeframe_start = 2022-02-01T05:00:00Z`,
// `timeframe_end = 2022-02-04T01:00:00Z` and `granularity=day`, the following
// windows will be returned for a customer in the `America/Los_Angeles` timezone
// since local midnight is `08:00` UTC:
//
// - `[2022-02-01T05:00:00Z, 2022-02-01T08:00:00Z)`
// - `[2022-02-01T08:00:00, 2022-02-02T08:00:00Z)`
// - `[2022-02-02T08:00:00, 2022-02-03T08:00:00Z)`
// - `[2022-02-03T08:00:00, 2022-02-04T01:00:00Z)`
//
// ```json
//
//	{
//	    "data": [
//	        {
//	            "billable_metric": {
//	                "id": "Q8w89wjTtBdejXKsm",
//	                "name": "API Requests"
//	            },
//	            "usage": [
//	                {
//	                    "quantity": 0,
//	                    "timeframe_end": "2022-02-01T08:00:00+00:00",
//	                    "timeframe_start": "2022-02-01T05:00:00+00:00"
//	                },
//	                {
//
//	                    "quantity": 0,
//	                    "timeframe_end": "2022-02-02T08:00:00+00:00",
//	                    "timeframe_start": "2022-02-01T08:00:00+00:00"
//	                },
//	                {
//	                    "quantity": 0,
//	                    "timeframe_end": "2022-02-03T08:00:00+00:00",
//	                    "timeframe_start": "2022-02-02T08:00:00+00:00"
//	                },
//	                {
//	                    "quantity": 0,
//	                    "timeframe_end": "2022-02-04T01:00:00+00:00",
//	                    "timeframe_start": "2022-02-03T08:00:00+00:00"
//	                }
//	            ],
//	            "view_mode": "periodic"
//	        },
//	        ...
//	    ]
//	}
//
// ```
//
// ## Decomposable vs. non-decomposable metrics
//
// Billable metrics fall into one of two categories: decomposable and
// non-decomposable. A decomposable billable metric, such as a sum or a count, can
// be displayed and aggregated across arbitrary timescales. On the other hand, a
// non-decomposable metric is not meaningful when only a slice of the billing
// window is considered.
//
// As an example, if we have a billable metric that's defined to count unique
// users, displaying a graph of unique users for each day is not representative of
// the billable metric value over the month (days could have an overlapping set of
// 'unique' users). Instead, what's useful for any given day is the number of
// unique users in the billing period so far, which are the _cumulative_ unique
// users.
//
// Accordingly, this endpoint returns treats these two types of metrics differently
// when `group_by` is specified:
//
//   - Decomposable metrics can be grouped by any event property.
//   - Non-decomposable metrics can only be grouped by the corresponding price's
//     invoice grouping key. If no invoice grouping key is present, the metric does
//     not support `group_by`.
//
// ## Matrix prices
//
// When a billable metric is attached to a price that uses matrix pricing, it's
// important to view usage grouped by those matrix dimensions. In this case, use
// the query parameters `first_dimension_key`, `first_dimension_value` and
// `second_dimension_key`, `second_dimension_value` while filtering to a specific
// `billable_metric_id`.
//
// For example, if your compute metric has a separate unit price (i.e. a matrix
// pricing model) per `region` and `provider`, your request might provide the
// following parameters:
//
// - `first_dimension_key`: `region`
// - `first_dimension_value`: `us-east-1`
// - `second_dimension_key`: `provider`
// - `second_dimension_value`: `aws`
func (r *SubscriptionService) FetchUsage(ctx context.Context, subscriptionID string, query SubscriptionFetchUsageParams, opts ...option.RequestOption) (res *SubscriptionUsage, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/usage", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is used to add and edit subscription
// [price intervals](../reference/price-interval). By making modifications to a
// subscription’s price intervals, you can
// [flexibly and atomically control the billing behavior of a subscription](../guides/product-catalog/modifying-subscriptions).
//
// ## Adding price intervals
//
// Prices can be added as price intervals to a subscription by specifying them in
// the `add` array. A `price_id` or `external_price_id` from an add-on price or
// previously removed plan price can be specified to reuse an existing price
// definition (however, please note that prices from other plans cannot be added to
// the subscription). Additionally, a new price can be specified using the `price`
// field — this price will be created automatically.
//
// A `start_date` must be specified for the price interval. This is the date when
// the price will start billing on the subscription, so this will notably result in
// an immediate charge at this time for any billed in advance fixed fees. The
// `end_date` will default to null, resulting in a price interval that will bill on
// a continually recurring basis. Both of these dates can be set in the past or the
// future and Orb will generate or modify invoices to ensure the subscription’s
// invoicing behavior is correct.
//
// Additionally, a discount, minimum, or maximum can be specified on the price
// interval. This will only apply to this price interval, not any other price
// intervals on the subscription.
//
// ## Adjustment intervals
//
// An adjustment interval represents the time period that a particular adjustment
// (a discount, minimum, or maximum) applies to the prices on a subscription.
// Adjustment intervals can be added to a subscription by specifying them in the
// `add_adjustments` array, or modified via the `edit_adjustments` array. When
// creating an adjustment interval, you'll need to provide the definition of the
// new adjustment (the type of adjustment, and which prices it applies to), as well
// as the start and end dates for the adjustment interval. The start and end dates
// of an existing adjustment interval can be edited via the `edit_adjustments`
// field (just like price intervals). (To "change" the amount of a discount,
// minimum, or maximum, then, you'll need to end the existing interval, and create
// a new adjustment interval with the new amount and a start date that matches the
// end date of the previous interval.)
//
// ## Editing price intervals
//
// Price intervals can be adjusted by specifying edits to make in the `edit` array.
// A `price_interval_id` to edit must be specified — this can be retrieved from the
// `price_intervals` field on the subscription.
//
// A new `start_date` or `end_date` can be specified to change the range of the
// price interval, which will modify past or future invoices to ensure correctness.
// If either of these dates are unspecified, they will default to the existing date
// on the price interval. To remove a price interval entirely from a subscription,
// set the `end_date` to be equivalent to the `start_date`.
//
// ## Fixed fee quantity transitions
//
// The fixed fee quantity transitions for a fixed fee price interval can also be
// specified when adding or editing by passing an array for
// `fixed_fee_quantity_transitions`. A fixed fee quantity transition must have a
// `quantity` and an `effective_date`, which is the date after which the new
// quantity will be used for billing. If a fixed fee quantity transition is
// scheduled at a billing period boundary, the full quantity will be billed on an
// invoice with the other prices on the subscription. If the fixed fee quantity
// transition is scheduled mid-billing period, the difference between the existing
// quantity and quantity specified in the transition will be prorated for the rest
// of the billing period and billed immediately, which will generate a new invoice.
//
// Notably, the list of fixed fee quantity transitions passed will overwrite the
// existing fixed fee quantity transitions on the price interval, so the entire
// list of transitions must be specified to add additional transitions. The
// existing list of transitions can be retrieved using the
// `fixed_fee_quantity_transitions` property on a subscription’s serialized price
// intervals.
func (r *SubscriptionService) PriceIntervals(ctx context.Context, subscriptionID string, body SubscriptionPriceIntervalsParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/price_intervals", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to change the plan on an existing subscription. It
// returns the serialized updated subscription object.
//
// The body parameter `change_option` determines the timing of the plan change. Orb
// supports three options:
//
//   - `end_of_subscription_term`: changes the plan at the end of the existing plan's
//     term.
//   - Issuing this plan change request for a monthly subscription will keep the
//     existing plan active until the start of the subsequent month, and
//     potentially issue an invoice for any usage charges incurred in the
//     intervening period.
//   - Issuing this plan change request for a yearly subscription will keep the
//     existing plan active for the full year.
//   - `immediate`: changes the plan immediately. Subscriptions that have their plan
//     changed with this option will be invoiced immediately. This invoice will
//     include any usage fees incurred in the billing period up to the change, along
//     with any prorated recurring fees for the billing period, if applicable.
//   - `requested_date`: changes the plan on the requested date (`change_date`). If
//     no timezone is provided, the customer's timezone is used. The `change_date`
//     body parameter is required if this option is chosen.
//
// Note that one of `plan_id` or `external_plan_id` is required in the request body
// for this operation.
//
// ## Price overrides, maximums, and minimums
//
// Price overrides are used to update some or all prices in the target plan.
// Minimums and maximums, much like price overrides, can be useful when a new
// customer has negotiated a new or different minimum or maximum spend cap than the
// default for the plan. The request format for price overrides, maximums, and
// minimums are the same as those in [subscription creation](create-subscription).
//
// ## Scheduling multiple plan changes
//
// When scheduling multiple plan changes with the same date, the latest plan change
// on that day takes effect.
//
// ## Prorations for in-advance fees
//
// By default, Orb calculates the prorated difference in any fixed fees when making
// a plan change, adjusting the customer balance as needed. For details on this
// behavior, see
// [Modifying subscriptions](../guides/product-catalog/modifying-subscriptions.md#prorations-for-in-advance-fees).
func (r *SubscriptionService) SchedulePlanChange(ctx context.Context, subscriptionID string, body SubscriptionSchedulePlanChangeParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/schedule_plan_change", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Manually trigger a phase, effective the given date (or the current time, if not
// specified).
func (r *SubscriptionService) TriggerPhase(ctx context.Context, subscriptionID string, body SubscriptionTriggerPhaseParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/trigger_phase", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to unschedule any pending cancellations for a
// subscription.
//
// To be eligible, the subscription must currently be active and have a future
// cancellation. This operation will turn on auto-renew, ensuring that the
// subscription does not end at the currently scheduled cancellation time.
func (r *SubscriptionService) UnscheduleCancellation(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/unschedule_cancellation", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint can be used to clear scheduled updates to the quantity for a fixed
// fee.
//
// If there are no updates scheduled, a request validation error will be returned
// with a 400 status code.
func (r *SubscriptionService) UnscheduleFixedFeeQuantityUpdates(ctx context.Context, subscriptionID string, body SubscriptionUnscheduleFixedFeeQuantityUpdatesParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/unschedule_fixed_fee_quantity_updates", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to unschedule any pending plan changes on an existing
// subscription.
func (r *SubscriptionService) UnschedulePendingPlanChanges(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/unschedule_pending_plan_changes", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint can be used to update the quantity for a fixed fee.
//
// To be eligible, the subscription must currently be active and the price
// specified must be a fixed fee (not usage-based). This operation will immediately
// update the quantity for the fee, or if a `effective_date` is passed in, will
// update the quantity on the requested date at midnight in the customer's
// timezone.
//
// In order to change the fixed fee quantity as of the next draft invoice for this
// subscription, pass `change_option=upcoming_invoice` without an `effective_date`
// specified.
//
// If the fee is an in-advance fixed fee, it will also issue an immediate invoice
// for the difference for the remainder of the billing period.
func (r *SubscriptionService) UpdateFixedFeeQuantity(ctx context.Context, subscriptionID string, body SubscriptionUpdateFixedFeeQuantityParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/update_fixed_fee_quantity", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A [subscription](../guides/core-concepts.mdx#subscription) represents the
// purchase of a plan by a customer.
//
// By default, subscriptions begin on the day that they're created and renew
// automatically for each billing cycle at the cadence that's configured in the
// plan definition.
//
// Subscriptions also default to **beginning of month alignment**, which means the
// first invoice issued for the subscription will have pro-rated charges between
// the `start_date` and the first of the following month. Subsequent billing
// periods will always start and end on a month boundary (e.g. subsequent month
// starts for monthly billing).
//
// Depending on the plan configuration, any _flat_ recurring fees will be billed
// either at the beginning (in-advance) or end (in-arrears) of each billing cycle.
// Plans default to **in-advance billing**. Usage-based fees are billed in arrears
// as usage is accumulated. In the normal course of events, you can expect an
// invoice to contain usage-based charges for the previous period, and a recurring
// fee for the following period.
type Subscription struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription.
	AdjustmentIntervals []SubscriptionAdjustmentInterval `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                        `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration SubscriptionBillingCycleAnchorConfiguration `json:"billing_cycle_anchor_configuration,required"`
	// The day of the month on which the billing cycle is anchored. If the maximum
	// number of days in a month is greater than this value, the last day of the month
	// is the billing cycle day (e.g. billing_cycle_day=31 for April means the billing
	// period begins on the 30th.
	BillingCycleDay int64     `json:"billing_cycle_day,required"`
	CreatedAt       time.Time `json:"created_at,required" format:"date-time"`
	// The end of the current billing period. This is an exclusive timestamp, such that
	// the instant returned is not part of the billing period. Set to null for
	// subscriptions that are not currently active.
	CurrentBillingPeriodEndDate time.Time `json:"current_billing_period_end_date,required,nullable" format:"date-time"`
	// The start date of the current billing period. This is an inclusive timestamp;
	// the instant returned is exactly the beginning of the billing period. Set to null
	// if the subscription is not currently active.
	CurrentBillingPeriodStartDate time.Time `json:"current_billing_period_start_date,required,nullable" format:"date-time"`
	// A customer is a buyer of your products, and the other party to the billing
	// relationship.
	//
	// In Orb, customers are assigned system generated identifiers automatically, but
	// it's often desirable to have these match existing identifiers in your system. To
	// avoid having to denormalize Orb ID information, you can pass in an
	// `external_customer_id` with your own identifier. See
	// [Customer ID Aliases](../guides/events-and-metrics/customer-aliases) for further
	// information about how these aliases work in Orb.
	//
	// In addition to having an identifier in your system, a customer may exist in a
	// payment provider solution like Stripe. Use the `payment_provider_id` and the
	// `payment_provider` enum field to express this mapping.
	//
	// A customer also has a timezone (from the standard
	// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
	// your account's timezone. See
	// [Timezone localization](../guides/product-catalog/timezones.md) for information
	// on what this timezone parameter influences within Orb.
	Customer Customer `json:"customer,required"`
	// Determines the default memo on this subscriptions' invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo string `json:"default_invoice_memo,required,nullable"`
	// The discount intervals for this subscription.
	DiscountIntervals []SubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                              `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []SubscriptionFixedFeeQuantitySchedule `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                 `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription.
	MaximumIntervals []SubscriptionMaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription.
	MinimumIntervals []SubscriptionMinimumInterval `json:"minimum_intervals,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// The [Plan](../guides/core-concepts.mdx#plan-and-price) resource represents a
	// plan that can be subscribed to by a customer. Plans define the billing behavior
	// of the subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan Plan `json:"plan,required"`
	// The price intervals for this subscription.
	PriceIntervals []SubscriptionPriceInterval `json:"price_intervals,required"`
	RedeemedCoupon SubscriptionRedeemedCoupon  `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time             `json:"start_date,required" format:"date-time"`
	Status    SubscriptionStatus    `json:"status,required"`
	TrialInfo SubscriptionTrialInfo `json:"trial_info,required"`
	JSON      subscriptionJSON      `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                              apijson.Field
	ActivePlanPhaseOrder            apijson.Field
	AdjustmentIntervals             apijson.Field
	AutoCollection                  apijson.Field
	BillingCycleAnchorConfiguration apijson.Field
	BillingCycleDay                 apijson.Field
	CreatedAt                       apijson.Field
	CurrentBillingPeriodEndDate     apijson.Field
	CurrentBillingPeriodStartDate   apijson.Field
	Customer                        apijson.Field
	DefaultInvoiceMemo              apijson.Field
	DiscountIntervals               apijson.Field
	EndDate                         apijson.Field
	FixedFeeQuantitySchedule        apijson.Field
	InvoicingThreshold              apijson.Field
	MaximumIntervals                apijson.Field
	Metadata                        apijson.Field
	MinimumIntervals                apijson.Field
	NetTerms                        apijson.Field
	Plan                            apijson.Field
	PriceIntervals                  apijson.Field
	RedeemedCoupon                  apijson.Field
	StartDate                       apijson.Field
	Status                          apijson.Field
	TrialInfo                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAdjustmentInterval struct {
	ID         string                                    `json:"id,required"`
	Adjustment SubscriptionAdjustmentIntervalsAdjustment `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time                          `json:"start_date,required" format:"date-time"`
	JSON      subscriptionAdjustmentIntervalJSON `json:"-"`
}

// subscriptionAdjustmentIntervalJSON contains the JSON metadata for the struct
// [SubscriptionAdjustmentInterval]
type subscriptionAdjustmentIntervalJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionAdjustmentInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionAdjustmentIntervalsAdjustment struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// The reason for the adjustment.
	Reason         string                                                  `json:"reason,required,nullable"`
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64 `json:"usage_discount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string                                        `json:"maximum_amount"`
	JSON          subscriptionAdjustmentIntervalsAdjustmentJSON `json:"-"`
	union         SubscriptionAdjustmentIntervalsAdjustmentUnion
}

// subscriptionAdjustmentIntervalsAdjustmentJSON contains the JSON metadata for the
// struct [SubscriptionAdjustmentIntervalsAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentJSON struct {
	AppliesToPriceIDs  apijson.Field
	Reason             apijson.Field
	AdjustmentType     apijson.Field
	AmountDiscount     apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	MinimumAmount      apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r subscriptionAdjustmentIntervalsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionAdjustmentIntervalsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionAdjustmentIntervalsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionAdjustmentIntervalsAdjustmentUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment].
func (r SubscriptionAdjustmentIntervalsAdjustment) AsUnion() SubscriptionAdjustmentIntervalsAdjustmentUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment],
// [SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment] or
// [SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment].
type SubscriptionAdjustmentIntervalsAdjustmentUnion interface {
	implementsSubscriptionAdjustmentIntervalsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionAdjustmentIntervalsAdjustmentUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment struct {
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The reason for the adjustment.
	Reason string                                                                `json:"reason,required,nullable"`
	JSON   subscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentJSON `json:"-"`
}

// subscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentJSON struct {
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustment) implementsSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentTypeAmountDiscount SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment struct {
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The reason for the adjustment.
	Reason string                                                                    `json:"reason,required,nullable"`
	JSON   subscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentJSON struct {
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	PercentageDiscount apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustment) implementsSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentPercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment struct {
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                              `json:"usage_discount,required"`
	JSON          subscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentJSON contains
// the JSON metadata for the struct
// [SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentJSON struct {
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustment) implementsSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentTypeUsageDiscount SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment struct {
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                         `json:"reason,required,nullable"`
	JSON   subscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentJSON `json:"-"`
}

// subscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentJSON contains the JSON
// metadata for the struct
// [SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentJSON struct {
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustment) implementsSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentTypeMinimum SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentType = "minimum"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment struct {
	AdjustmentType SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The reason for the adjustment.
	Reason string                                                         `json:"reason,required,nullable"`
	JSON   subscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentJSON `json:"-"`
}

// subscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentJSON contains the JSON
// metadata for the struct
// [SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment]
type subscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentJSON struct {
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	MaximumAmount     apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustment) implementsSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentTypeMaximum SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType string

const (
	SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum            SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "minimum"
	SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum            SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionAdjustmentIntervalsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum, SubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionBillingCycleAnchorConfiguration struct {
	// The day of the month on which the billing cycle is anchored. If the maximum
	// number of days in a month is greater than this value, the last day of the month
	// is the billing cycle day (e.g. billing_cycle_day=31 for April means the billing
	// period begins on the 30th.
	Day int64 `json:"day,required"`
	// The month on which the billing cycle is anchored (e.g. a quarterly price
	// anchored in February would have cycles starting February, May, August, and
	// November).
	Month int64 `json:"month,nullable"`
	// The year on which the billing cycle is anchored (e.g. a 2 year billing cycle
	// anchored on 2021 would have cycles starting on 2021, 2023, 2025, etc.).
	Year int64                                           `json:"year,nullable"`
	JSON subscriptionBillingCycleAnchorConfigurationJSON `json:"-"`
}

// subscriptionBillingCycleAnchorConfigurationJSON contains the JSON metadata for
// the struct [SubscriptionBillingCycleAnchorConfiguration]
type subscriptionBillingCycleAnchorConfigurationJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionBillingCycleAnchorConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionBillingCycleAnchorConfigurationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionDiscountInterval struct {
	DiscountType SubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{} `json:"applies_to_price_interval_ids"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                          `json:"usage_discount"`
	JSON          subscriptionDiscountIntervalJSON `json:"-"`
	union         SubscriptionDiscountIntervalsUnion
}

// subscriptionDiscountIntervalJSON contains the JSON metadata for the struct
// [SubscriptionDiscountInterval]
type subscriptionDiscountIntervalJSON struct {
	DiscountType              apijson.Field
	AmountDiscount            apijson.Field
	StartDate                 apijson.Field
	EndDate                   apijson.Field
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	PercentageDiscount        apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r subscriptionDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionDiscountIntervalsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionDiscountIntervalsPercentageDiscountInterval],
// [SubscriptionDiscountIntervalsUsageDiscountInterval].
func (r SubscriptionDiscountInterval) AsUnion() SubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by [SubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionDiscountIntervalsPercentageDiscountInterval] or
// [SubscriptionDiscountIntervalsUsageDiscountInterval].
type SubscriptionDiscountIntervalsUnion interface {
	implementsSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionDiscountIntervalsAmountDiscountInterval{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionDiscountIntervalsPercentageDiscountInterval{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionDiscountIntervalsUsageDiscountInterval{}),
			DiscriminatorValue: "usage",
		},
	)
}

type SubscriptionDiscountIntervalsAmountDiscountInterval struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                        `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time                                               `json:"start_date,required" format:"date-time"`
	JSON      subscriptionDiscountIntervalsAmountDiscountIntervalJSON `json:"-"`
}

// subscriptionDiscountIntervalsAmountDiscountIntervalJSON contains the JSON
// metadata for the struct [SubscriptionDiscountIntervalsAmountDiscountInterval]
type subscriptionDiscountIntervalsAmountDiscountIntervalJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionDiscountIntervalsAmountDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDiscountIntervalsAmountDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDiscountIntervalsAmountDiscountInterval) implementsSubscriptionDiscountInterval() {
}

type SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType string

const (
	SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType = "amount"
)

func (r SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionDiscountIntervalsPercentageDiscountInterval struct {
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                            `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                   `json:"start_date,required" format:"date-time"`
	JSON      subscriptionDiscountIntervalsPercentageDiscountIntervalJSON `json:"-"`
}

// subscriptionDiscountIntervalsPercentageDiscountIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionDiscountIntervalsPercentageDiscountInterval]
type subscriptionDiscountIntervalsPercentageDiscountIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionDiscountIntervalsPercentageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDiscountIntervalsPercentageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDiscountIntervalsPercentageDiscountInterval) implementsSubscriptionDiscountInterval() {
}

type SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType string

const (
	SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType = "percentage"
)

func (r SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage:
		return true
	}
	return false
}

type SubscriptionDiscountIntervalsUsageDiscountInterval struct {
	// The price ids that this discount interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                       `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                                                `json:"usage_discount,required"`
	JSON          subscriptionDiscountIntervalsUsageDiscountIntervalJSON `json:"-"`
}

// subscriptionDiscountIntervalsUsageDiscountIntervalJSON contains the JSON
// metadata for the struct [SubscriptionDiscountIntervalsUsageDiscountInterval]
type subscriptionDiscountIntervalsUsageDiscountIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionDiscountIntervalsUsageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionDiscountIntervalsUsageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionDiscountIntervalsUsageDiscountInterval) implementsSubscriptionDiscountInterval() {
}

type SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType string

const (
	SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType = "usage"
)

func (r SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionDiscountIntervalsDiscountType string

const (
	SubscriptionDiscountIntervalsDiscountTypeAmount     SubscriptionDiscountIntervalsDiscountType = "amount"
	SubscriptionDiscountIntervalsDiscountTypePercentage SubscriptionDiscountIntervalsDiscountType = "percentage"
	SubscriptionDiscountIntervalsDiscountTypeUsage      SubscriptionDiscountIntervalsDiscountType = "usage"
)

func (r SubscriptionDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionDiscountIntervalsDiscountTypeAmount, SubscriptionDiscountIntervalsDiscountTypePercentage, SubscriptionDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionFixedFeeQuantitySchedule struct {
	EndDate   time.Time                                `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                                   `json:"price_id,required"`
	Quantity  float64                                  `json:"quantity,required"`
	StartDate time.Time                                `json:"start_date,required" format:"date-time"`
	JSON      subscriptionFixedFeeQuantityScheduleJSON `json:"-"`
}

// subscriptionFixedFeeQuantityScheduleJSON contains the JSON metadata for the
// struct [SubscriptionFixedFeeQuantitySchedule]
type subscriptionFixedFeeQuantityScheduleJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFixedFeeQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFixedFeeQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type SubscriptionMaximumInterval struct {
	// The price ids that this maximum interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time                       `json:"start_date,required" format:"date-time"`
	JSON      subscriptionMaximumIntervalJSON `json:"-"`
}

// subscriptionMaximumIntervalJSON contains the JSON metadata for the struct
// [SubscriptionMaximumInterval]
type subscriptionMaximumIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionMaximumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMaximumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionMinimumInterval struct {
	// The price ids that this minimum interval applies to.
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time                       `json:"start_date,required" format:"date-time"`
	JSON      subscriptionMinimumIntervalJSON `json:"-"`
}

// subscriptionMinimumIntervalJSON contains the JSON metadata for the struct
// [SubscriptionMinimumInterval]
type subscriptionMinimumIntervalJSON struct {
	AppliesToPriceIDs         apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionMinimumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionMinimumIntervalJSON) RawJSON() string {
	return r.raw
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type SubscriptionPriceInterval struct {
	ID string `json:"id,required"`
	// The day of the month that Orb bills for this price
	BillingCycleDay int64 `json:"billing_cycle_day,required"`
	// The end of the current billing period. This is an exclusive timestamp, such that
	// the instant returned is exactly the end of the billing period. Set to null if
	// this price interval is not currently active.
	CurrentBillingPeriodEndDate time.Time `json:"current_billing_period_end_date,required,nullable" format:"date-time"`
	// The start date of the current billing period. This is an inclusive timestamp;
	// the instant returned is exactly the beginning of the billing period. Set to null
	// if this price interval is not currently active.
	CurrentBillingPeriodStartDate time.Time `json:"current_billing_period_start_date,required,nullable" format:"date-time"`
	// The end date of the price interval. This is the date that Orb stops billing for
	// this price.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The fixed fee quantity transitions for this price interval. This is only
	// relevant for fixed fees.
	FixedFeeQuantityTransitions []SubscriptionPriceIntervalsFixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
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
	// The start date of the price interval. This is the date that Orb starts billing
	// for this price.
	StartDate time.Time                     `json:"start_date,required" format:"date-time"`
	JSON      subscriptionPriceIntervalJSON `json:"-"`
}

// subscriptionPriceIntervalJSON contains the JSON metadata for the struct
// [SubscriptionPriceInterval]
type subscriptionPriceIntervalJSON struct {
	ID                            apijson.Field
	BillingCycleDay               apijson.Field
	CurrentBillingPeriodEndDate   apijson.Field
	CurrentBillingPeriodStartDate apijson.Field
	EndDate                       apijson.Field
	FixedFeeQuantityTransitions   apijson.Field
	Price                         apijson.Field
	StartDate                     apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SubscriptionPriceInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPriceIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionPriceIntervalsFixedFeeQuantityTransition struct {
	EffectiveDate time.Time                                                `json:"effective_date,required" format:"date-time"`
	PriceID       string                                                   `json:"price_id,required"`
	Quantity      int64                                                    `json:"quantity,required"`
	JSON          subscriptionPriceIntervalsFixedFeeQuantityTransitionJSON `json:"-"`
}

// subscriptionPriceIntervalsFixedFeeQuantityTransitionJSON contains the JSON
// metadata for the struct [SubscriptionPriceIntervalsFixedFeeQuantityTransition]
type subscriptionPriceIntervalsFixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionPriceIntervalsFixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionPriceIntervalsFixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionRedeemedCoupon struct {
	CouponID  string                         `json:"coupon_id,required"`
	EndDate   time.Time                      `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time                      `json:"start_date,required" format:"date-time"`
	JSON      subscriptionRedeemedCouponJSON `json:"-"`
}

// subscriptionRedeemedCouponJSON contains the JSON metadata for the struct
// [SubscriptionRedeemedCoupon]
type subscriptionRedeemedCouponJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionRedeemedCoupon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionRedeemedCouponJSON) RawJSON() string {
	return r.raw
}

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusEnded    SubscriptionStatus = "ended"
	SubscriptionStatusUpcoming SubscriptionStatus = "upcoming"
)

func (r SubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionStatusActive, SubscriptionStatusEnded, SubscriptionStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionTrialInfo struct {
	EndDate time.Time                 `json:"end_date,required,nullable" format:"date-time"`
	JSON    subscriptionTrialInfoJSON `json:"-"`
}

// subscriptionTrialInfoJSON contains the JSON metadata for the struct
// [SubscriptionTrialInfo]
type subscriptionTrialInfoJSON struct {
	EndDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionTrialInfoJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsage struct {
	// This field can have the runtime type of
	// [[]SubscriptionUsageUngroupedSubscriptionUsageData],
	// [[]SubscriptionUsageGroupedSubscriptionUsageData].
	Data               interface{}               `json:"data"`
	PaginationMetadata shared.PaginationMetadata `json:"pagination_metadata,nullable"`
	JSON               subscriptionUsageJSON     `json:"-"`
	union              SubscriptionUsageUnion
}

// subscriptionUsageJSON contains the JSON metadata for the struct
// [SubscriptionUsage]
type subscriptionUsageJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r subscriptionUsageJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionUsage) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionUsage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SubscriptionUsageUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionUsageUngroupedSubscriptionUsage],
// [SubscriptionUsageGroupedSubscriptionUsage].
func (r SubscriptionUsage) AsUnion() SubscriptionUsageUnion {
	return r.union
}

// Union satisfied by [SubscriptionUsageUngroupedSubscriptionUsage] or
// [SubscriptionUsageGroupedSubscriptionUsage].
type SubscriptionUsageUnion interface {
	implementsSubscriptionUsage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionUsageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUsageUngroupedSubscriptionUsage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SubscriptionUsageGroupedSubscriptionUsage{}),
		},
	)
}

type SubscriptionUsageUngroupedSubscriptionUsage struct {
	Data []SubscriptionUsageUngroupedSubscriptionUsageData `json:"data,required"`
	JSON subscriptionUsageUngroupedSubscriptionUsageJSON   `json:"-"`
}

// subscriptionUsageUngroupedSubscriptionUsageJSON contains the JSON metadata for
// the struct [SubscriptionUsageUngroupedSubscriptionUsage]
type subscriptionUsageUngroupedSubscriptionUsageJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageUngroupedSubscriptionUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageUngroupedSubscriptionUsageJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUsageUngroupedSubscriptionUsage) implementsSubscriptionUsage() {}

type SubscriptionUsageUngroupedSubscriptionUsageData struct {
	BillableMetric SubscriptionUsageUngroupedSubscriptionUsageDataBillableMetric `json:"billable_metric,required"`
	Usage          []SubscriptionUsageUngroupedSubscriptionUsageDataUsage        `json:"usage,required"`
	ViewMode       SubscriptionUsageUngroupedSubscriptionUsageDataViewMode       `json:"view_mode,required"`
	JSON           subscriptionUsageUngroupedSubscriptionUsageDataJSON           `json:"-"`
}

// subscriptionUsageUngroupedSubscriptionUsageDataJSON contains the JSON metadata
// for the struct [SubscriptionUsageUngroupedSubscriptionUsageData]
type subscriptionUsageUngroupedSubscriptionUsageDataJSON struct {
	BillableMetric apijson.Field
	Usage          apijson.Field
	ViewMode       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageUngroupedSubscriptionUsageData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageUngroupedSubscriptionUsageDataJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageUngroupedSubscriptionUsageDataBillableMetric struct {
	ID   string                                                            `json:"id,required"`
	Name string                                                            `json:"name,required"`
	JSON subscriptionUsageUngroupedSubscriptionUsageDataBillableMetricJSON `json:"-"`
}

// subscriptionUsageUngroupedSubscriptionUsageDataBillableMetricJSON contains the
// JSON metadata for the struct
// [SubscriptionUsageUngroupedSubscriptionUsageDataBillableMetric]
type subscriptionUsageUngroupedSubscriptionUsageDataBillableMetricJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageUngroupedSubscriptionUsageDataBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageUngroupedSubscriptionUsageDataBillableMetricJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageUngroupedSubscriptionUsageDataUsage struct {
	Quantity       float64                                                  `json:"quantity,required"`
	TimeframeEnd   time.Time                                                `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                                                `json:"timeframe_start,required" format:"date-time"`
	JSON           subscriptionUsageUngroupedSubscriptionUsageDataUsageJSON `json:"-"`
}

// subscriptionUsageUngroupedSubscriptionUsageDataUsageJSON contains the JSON
// metadata for the struct [SubscriptionUsageUngroupedSubscriptionUsageDataUsage]
type subscriptionUsageUngroupedSubscriptionUsageDataUsageJSON struct {
	Quantity       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageUngroupedSubscriptionUsageDataUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageUngroupedSubscriptionUsageDataUsageJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageUngroupedSubscriptionUsageDataViewMode string

const (
	SubscriptionUsageUngroupedSubscriptionUsageDataViewModePeriodic   SubscriptionUsageUngroupedSubscriptionUsageDataViewMode = "periodic"
	SubscriptionUsageUngroupedSubscriptionUsageDataViewModeCumulative SubscriptionUsageUngroupedSubscriptionUsageDataViewMode = "cumulative"
)

func (r SubscriptionUsageUngroupedSubscriptionUsageDataViewMode) IsKnown() bool {
	switch r {
	case SubscriptionUsageUngroupedSubscriptionUsageDataViewModePeriodic, SubscriptionUsageUngroupedSubscriptionUsageDataViewModeCumulative:
		return true
	}
	return false
}

type SubscriptionUsageGroupedSubscriptionUsage struct {
	Data               []SubscriptionUsageGroupedSubscriptionUsageData `json:"data,required"`
	PaginationMetadata shared.PaginationMetadata                       `json:"pagination_metadata,nullable"`
	JSON               subscriptionUsageGroupedSubscriptionUsageJSON   `json:"-"`
}

// subscriptionUsageGroupedSubscriptionUsageJSON contains the JSON metadata for the
// struct [SubscriptionUsageGroupedSubscriptionUsage]
type subscriptionUsageGroupedSubscriptionUsageJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionUsageGroupedSubscriptionUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageGroupedSubscriptionUsageJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionUsageGroupedSubscriptionUsage) implementsSubscriptionUsage() {}

type SubscriptionUsageGroupedSubscriptionUsageData struct {
	BillableMetric SubscriptionUsageGroupedSubscriptionUsageDataBillableMetric `json:"billable_metric,required"`
	MetricGroup    SubscriptionUsageGroupedSubscriptionUsageDataMetricGroup    `json:"metric_group,required"`
	Usage          []SubscriptionUsageGroupedSubscriptionUsageDataUsage        `json:"usage,required"`
	ViewMode       SubscriptionUsageGroupedSubscriptionUsageDataViewMode       `json:"view_mode,required"`
	JSON           subscriptionUsageGroupedSubscriptionUsageDataJSON           `json:"-"`
}

// subscriptionUsageGroupedSubscriptionUsageDataJSON contains the JSON metadata for
// the struct [SubscriptionUsageGroupedSubscriptionUsageData]
type subscriptionUsageGroupedSubscriptionUsageDataJSON struct {
	BillableMetric apijson.Field
	MetricGroup    apijson.Field
	Usage          apijson.Field
	ViewMode       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageGroupedSubscriptionUsageData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageGroupedSubscriptionUsageDataJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageGroupedSubscriptionUsageDataBillableMetric struct {
	ID   string                                                          `json:"id,required"`
	Name string                                                          `json:"name,required"`
	JSON subscriptionUsageGroupedSubscriptionUsageDataBillableMetricJSON `json:"-"`
}

// subscriptionUsageGroupedSubscriptionUsageDataBillableMetricJSON contains the
// JSON metadata for the struct
// [SubscriptionUsageGroupedSubscriptionUsageDataBillableMetric]
type subscriptionUsageGroupedSubscriptionUsageDataBillableMetricJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUsageGroupedSubscriptionUsageDataBillableMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageGroupedSubscriptionUsageDataBillableMetricJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageGroupedSubscriptionUsageDataMetricGroup struct {
	PropertyKey   string                                                       `json:"property_key,required"`
	PropertyValue string                                                       `json:"property_value,required"`
	JSON          subscriptionUsageGroupedSubscriptionUsageDataMetricGroupJSON `json:"-"`
}

// subscriptionUsageGroupedSubscriptionUsageDataMetricGroupJSON contains the JSON
// metadata for the struct
// [SubscriptionUsageGroupedSubscriptionUsageDataMetricGroup]
type subscriptionUsageGroupedSubscriptionUsageDataMetricGroupJSON struct {
	PropertyKey   apijson.Field
	PropertyValue apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionUsageGroupedSubscriptionUsageDataMetricGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageGroupedSubscriptionUsageDataMetricGroupJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageGroupedSubscriptionUsageDataUsage struct {
	Quantity       float64                                                `json:"quantity,required"`
	TimeframeEnd   time.Time                                              `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time                                              `json:"timeframe_start,required" format:"date-time"`
	JSON           subscriptionUsageGroupedSubscriptionUsageDataUsageJSON `json:"-"`
}

// subscriptionUsageGroupedSubscriptionUsageDataUsageJSON contains the JSON
// metadata for the struct [SubscriptionUsageGroupedSubscriptionUsageDataUsage]
type subscriptionUsageGroupedSubscriptionUsageDataUsageJSON struct {
	Quantity       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionUsageGroupedSubscriptionUsageDataUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUsageGroupedSubscriptionUsageDataUsageJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUsageGroupedSubscriptionUsageDataViewMode string

const (
	SubscriptionUsageGroupedSubscriptionUsageDataViewModePeriodic   SubscriptionUsageGroupedSubscriptionUsageDataViewMode = "periodic"
	SubscriptionUsageGroupedSubscriptionUsageDataViewModeCumulative SubscriptionUsageGroupedSubscriptionUsageDataViewMode = "cumulative"
)

func (r SubscriptionUsageGroupedSubscriptionUsageDataViewMode) IsKnown() bool {
	switch r {
	case SubscriptionUsageGroupedSubscriptionUsageDataViewModePeriodic, SubscriptionUsageGroupedSubscriptionUsageDataViewModeCumulative:
		return true
	}
	return false
}

type Subscriptions struct {
	Data               []Subscription            `json:"data,required"`
	PaginationMetadata shared.PaginationMetadata `json:"pagination_metadata,required"`
	JSON               subscriptionsJSON         `json:"-"`
}

// subscriptionsJSON contains the JSON metadata for the struct [Subscriptions]
type subscriptionsJSON struct {
	Data               apijson.Field
	PaginationMetadata apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscriptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionFetchCostsResponse struct {
	Data []SubscriptionFetchCostsResponseData `json:"data,required"`
	JSON subscriptionFetchCostsResponseJSON   `json:"-"`
}

// subscriptionFetchCostsResponseJSON contains the JSON metadata for the struct
// [SubscriptionFetchCostsResponse]
type subscriptionFetchCostsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFetchCostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFetchCostsResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionFetchCostsResponseData struct {
	PerPriceCosts []SubscriptionFetchCostsResponseDataPerPriceCost `json:"per_price_costs,required"`
	// Total costs for the timeframe, excluding any minimums and discounts.
	Subtotal       string    `json:"subtotal,required"`
	TimeframeEnd   time.Time `json:"timeframe_end,required" format:"date-time"`
	TimeframeStart time.Time `json:"timeframe_start,required" format:"date-time"`
	// Total costs for the timeframe, including any minimums and discounts.
	Total string                                 `json:"total,required"`
	JSON  subscriptionFetchCostsResponseDataJSON `json:"-"`
}

// subscriptionFetchCostsResponseDataJSON contains the JSON metadata for the struct
// [SubscriptionFetchCostsResponseData]
type subscriptionFetchCostsResponseDataJSON struct {
	PerPriceCosts  apijson.Field
	Subtotal       apijson.Field
	TimeframeEnd   apijson.Field
	TimeframeStart apijson.Field
	Total          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionFetchCostsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFetchCostsResponseDataJSON) RawJSON() string {
	return r.raw
}

type SubscriptionFetchCostsResponseDataPerPriceCost struct {
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
	Quantity float64                                            `json:"quantity,nullable"`
	JSON     subscriptionFetchCostsResponseDataPerPriceCostJSON `json:"-"`
}

// subscriptionFetchCostsResponseDataPerPriceCostJSON contains the JSON metadata
// for the struct [SubscriptionFetchCostsResponseDataPerPriceCost]
type subscriptionFetchCostsResponseDataPerPriceCostJSON struct {
	Price       apijson.Field
	Subtotal    apijson.Field
	Total       apijson.Field
	Quantity    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFetchCostsResponseDataPerPriceCost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFetchCostsResponseDataPerPriceCostJSON) RawJSON() string {
	return r.raw
}

type SubscriptionFetchScheduleResponse struct {
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	EndDate   time.Time                             `json:"end_date,required,nullable" format:"date-time"`
	Plan      SubscriptionFetchScheduleResponsePlan `json:"plan,required"`
	StartDate time.Time                             `json:"start_date,required" format:"date-time"`
	JSON      subscriptionFetchScheduleResponseJSON `json:"-"`
}

// subscriptionFetchScheduleResponseJSON contains the JSON metadata for the struct
// [SubscriptionFetchScheduleResponse]
type subscriptionFetchScheduleResponseJSON struct {
	CreatedAt   apijson.Field
	EndDate     apijson.Field
	Plan        apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionFetchScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFetchScheduleResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionFetchScheduleResponsePlan struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string                                    `json:"external_plan_id,required,nullable"`
	Name           string                                    `json:"name,required,nullable"`
	JSON           subscriptionFetchScheduleResponsePlanJSON `json:"-"`
}

// subscriptionFetchScheduleResponsePlanJSON contains the JSON metadata for the
// struct [SubscriptionFetchScheduleResponsePlan]
type subscriptionFetchScheduleResponsePlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionFetchScheduleResponsePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionFetchScheduleResponsePlanJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewParams struct {
	// Additional adjustments to be added to the subscription
	AddAdjustments param.Field[[]SubscriptionNewParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the subscription
	AddPrices                             param.Field[[]SubscriptionNewParamsAddPrice] `json:"add_prices"`
	AlignBillingWithSubscriptionStartDate param.Field[bool]                            `json:"align_billing_with_subscription_start_date"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. If not specified, this
	// defaults to the behavior configured for this customer.
	AutoCollection                  param.Field[bool]                                                 `json:"auto_collection"`
	AwsRegion                       param.Field[string]                                               `json:"aws_region"`
	BillingCycleAnchorConfiguration param.Field[SubscriptionNewParamsBillingCycleAnchorConfiguration] `json:"billing_cycle_anchor_configuration"`
	// Redemption code to be used for this subscription. If the coupon cannot be found
	// by its redemption code, or cannot be redeemed, an error response will be
	// returned and the subscription creation or plan change will not be scheduled.
	CouponRedemptionCode param.Field[string]  `json:"coupon_redemption_code"`
	CreditsOverageRate   param.Field[float64] `json:"credits_overage_rate"`
	CustomerID           param.Field[string]  `json:"customer_id"`
	// Determines the default memo on this subscription's invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo             param.Field[string]                                   `json:"default_invoice_memo"`
	EndDate                        param.Field[time.Time]                                `json:"end_date" format:"date-time"`
	ExternalCustomerID             param.Field[string]                                   `json:"external_customer_id"`
	ExternalMarketplace            param.Field[SubscriptionNewParamsExternalMarketplace] `json:"external_marketplace"`
	ExternalMarketplaceReportingID param.Field[string]                                   `json:"external_marketplace_reporting_id"`
	// The external_plan_id of the plan that the given subscription should be switched
	// to. Note that either this property or `plan_id` must be specified.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// The phase of the plan to start with
	InitialPhaseOrder param.Field[int64] `json:"initial_phase_order"`
	// When this subscription's accrued usage reaches this threshold, an invoice will
	// be issued for the subscription. If not specified, invoices will only be issued
	// at the end of the billing period.
	InvoicingThreshold param.Field[string] `json:"invoicing_threshold"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0. If not provided, this defaults to the value specified in the plan.
	NetTerms               param.Field[int64]   `json:"net_terms"`
	PerCreditOverageAmount param.Field[float64] `json:"per_credit_overage_amount"`
	// The plan that the given subscription should be switched to. Note that either
	// this property or `external_plan_id` must be specified.
	PlanID param.Field[string] `json:"plan_id"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides param.Field[[]SubscriptionNewParamsPriceOverrideUnion] `json:"price_overrides"`
	// Plan adjustments to be removed from the subscription
	RemoveAdjustments param.Field[[]SubscriptionNewParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Plan prices to be removed from the subscription
	RemovePrices param.Field[[]SubscriptionNewParamsRemovePrice] `json:"remove_prices"`
	// Plan adjustments to be replaced with additional adjustments on the subscription
	ReplaceAdjustments param.Field[[]SubscriptionNewParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Plan prices to be replaced with additional prices on the subscription
	ReplacePrices param.Field[[]SubscriptionNewParamsReplacePrice] `json:"replace_prices"`
	StartDate     param.Field[time.Time]                           `json:"start_date" format:"date-time"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription. If null, the adjustment will
	// start when the phase or subscription starts.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription. If null, the adjustment will
	// start when the phase or subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r SubscriptionNewParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionNewParamsAddAdjustmentsAdjustment struct {
	AppliesToPriceIDs  param.Field[interface{}]                                                 `json:"applies_to_price_ids"`
	AdjustmentType     param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                     `json:"percentage_discount"`
	AmountDiscount     param.Field[string]                                                      `json:"amount_discount"`
	MinimumAmount      param.Field[string]                                                      `json:"minimum_amount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id"`
	MaximumAmount param.Field[string] `json:"maximum_amount"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustment) implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by
// [SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimum],
// [SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximum],
// [SubscriptionNewParamsAddAdjustmentsAdjustment].
type SubscriptionNewParamsAddAdjustmentsAdjustmentUnion interface {
	implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion()
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                       `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimum) implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximum) implementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPrice struct {
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription. If null, billing will end when the phase or
	// subscription ends.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[SubscriptionNewParamsAddPricesPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
	// The start date of the price interval. This is the date that the price will start
	// billing on the subscription. If null, billing will start when the phase or
	// subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r SubscriptionNewParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new price to create and add to the subscription.
type SubscriptionNewParamsAddPricesPrice struct {
	Metadata param.Field[interface{}] `json:"metadata,required"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// The id of the item the plan will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// The cadence to bill for this price on.
	Cadence                     param.Field[SubscriptionNewParamsAddPricesPriceCadence] `json:"cadence,required"`
	BillingCycleConfiguration   param.Field[interface{}]                                `json:"billing_cycle_configuration,required"`
	InvoicingCycleConfiguration param.Field[interface{}]                                `json:"invoicing_cycle_configuration,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64]                                      `json:"conversion_rate"`
	ModelType      param.Field[SubscriptionNewParamsAddPricesPriceModelType] `json:"model_type,required"`
	UnitConfig     param.Field[interface{}]                                  `json:"unit_config,required"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                         param.Field[string]      `json:"currency"`
	PackageConfig                    param.Field[interface{}] `json:"package_config,required"`
	MatrixConfig                     param.Field[interface{}] `json:"matrix_config,required"`
	TieredConfig                     param.Field[interface{}] `json:"tiered_config,required"`
	TieredBpsConfig                  param.Field[interface{}] `json:"tiered_bps_config,required"`
	BpsConfig                        param.Field[interface{}] `json:"bps_config,required"`
	BulkBpsConfig                    param.Field[interface{}] `json:"bulk_bps_config,required"`
	BulkConfig                       param.Field[interface{}] `json:"bulk_config,required"`
	ThresholdTotalAmountConfig       param.Field[interface{}] `json:"threshold_total_amount_config,required"`
	TieredPackageConfig              param.Field[interface{}] `json:"tiered_package_config,required"`
	TieredWithMinimumConfig          param.Field[interface{}] `json:"tiered_with_minimum_config,required"`
	UnitWithPercentConfig            param.Field[interface{}] `json:"unit_with_percent_config,required"`
	PackageWithAllocationConfig      param.Field[interface{}] `json:"package_with_allocation_config,required"`
	TieredWithProrationConfig        param.Field[interface{}] `json:"tiered_with_proration_config,required"`
	UnitWithProrationConfig          param.Field[interface{}] `json:"unit_with_proration_config,required"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config,required"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	BulkWithProrationConfig          param.Field[interface{}] `json:"bulk_with_proration_config,required"`
}

func (r SubscriptionNewParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

// The definition of a new price to create and add to the subscription.
//
// Satisfied by [SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPrice],
// [SubscriptionNewParamsAddPricesPrice].
type SubscriptionNewParamsAddPricesPriceUnion interface {
	implementsSubscriptionNewParamsAddPricesPriceUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                `json:"name,required"`
	UnitConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelType = "unit"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                      `json:"name,required"`
	PackageConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelTypePackage SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelType = "package"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                                                    `json:"item_id,required"`
	MatrixConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelTypeMatrix SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelType = "matrix"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                 `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                    `json:"name,required"`
	TieredConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                    `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                          `json:"name,required"`
	TieredBpsConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelType = "tiered_bps"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPrice struct {
	BpsConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                              `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelTypeBps SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelType = "bps"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPrice struct {
	BulkBpsConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelType = "bulk_bps"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPrice struct {
	BulkConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelTypeBulk SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelType = "bulk"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                          `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                          `json:"grouped_allocation_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                   `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceModelTypeUnit                       SubscriptionNewParamsAddPricesPriceModelType = "unit"
	SubscriptionNewParamsAddPricesPriceModelTypePackage                    SubscriptionNewParamsAddPricesPriceModelType = "package"
	SubscriptionNewParamsAddPricesPriceModelTypeMatrix                     SubscriptionNewParamsAddPricesPriceModelType = "matrix"
	SubscriptionNewParamsAddPricesPriceModelTypeTiered                     SubscriptionNewParamsAddPricesPriceModelType = "tiered"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredBps                  SubscriptionNewParamsAddPricesPriceModelType = "tiered_bps"
	SubscriptionNewParamsAddPricesPriceModelTypeBps                        SubscriptionNewParamsAddPricesPriceModelType = "bps"
	SubscriptionNewParamsAddPricesPriceModelTypeBulkBps                    SubscriptionNewParamsAddPricesPriceModelType = "bulk_bps"
	SubscriptionNewParamsAddPricesPriceModelTypeBulk                       SubscriptionNewParamsAddPricesPriceModelType = "bulk"
	SubscriptionNewParamsAddPricesPriceModelTypeThresholdTotalAmount       SubscriptionNewParamsAddPricesPriceModelType = "threshold_total_amount"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredPackage              SubscriptionNewParamsAddPricesPriceModelType = "tiered_package"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredWithMinimum          SubscriptionNewParamsAddPricesPriceModelType = "tiered_with_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypeUnitWithPercent            SubscriptionNewParamsAddPricesPriceModelType = "unit_with_percent"
	SubscriptionNewParamsAddPricesPriceModelTypePackageWithAllocation      SubscriptionNewParamsAddPricesPriceModelType = "package_with_allocation"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredWithProration        SubscriptionNewParamsAddPricesPriceModelType = "tiered_with_proration"
	SubscriptionNewParamsAddPricesPriceModelTypeUnitWithProration          SubscriptionNewParamsAddPricesPriceModelType = "unit_with_proration"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedAllocation          SubscriptionNewParamsAddPricesPriceModelType = "grouped_allocation"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum SubscriptionNewParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypeBulkWithProration          SubscriptionNewParamsAddPricesPriceModelType = "bulk_with_proration"
)

func (r SubscriptionNewParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceModelTypeUnit, SubscriptionNewParamsAddPricesPriceModelTypePackage, SubscriptionNewParamsAddPricesPriceModelTypeMatrix, SubscriptionNewParamsAddPricesPriceModelTypeTiered, SubscriptionNewParamsAddPricesPriceModelTypeTieredBps, SubscriptionNewParamsAddPricesPriceModelTypeBps, SubscriptionNewParamsAddPricesPriceModelTypeBulkBps, SubscriptionNewParamsAddPricesPriceModelTypeBulk, SubscriptionNewParamsAddPricesPriceModelTypeThresholdTotalAmount, SubscriptionNewParamsAddPricesPriceModelTypeTieredPackage, SubscriptionNewParamsAddPricesPriceModelTypeTieredWithMinimum, SubscriptionNewParamsAddPricesPriceModelTypeUnitWithPercent, SubscriptionNewParamsAddPricesPriceModelTypePackageWithAllocation, SubscriptionNewParamsAddPricesPriceModelTypeTieredWithProration, SubscriptionNewParamsAddPricesPriceModelTypeUnitWithProration, SubscriptionNewParamsAddPricesPriceModelTypeGroupedAllocation, SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionNewParamsAddPricesPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type SubscriptionNewParamsBillingCycleAnchorConfiguration struct {
	// The day of the month on which the billing cycle is anchored. If the maximum
	// number of days in a month is greater than this value, the last day of the month
	// is the billing cycle day (e.g. billing_cycle_day=31 for April means the billing
	// period begins on the 30th.
	Day param.Field[int64] `json:"day,required"`
	// The month on which the billing cycle is anchored (e.g. a quarterly price
	// anchored in February would have cycles starting February, May, August, and
	// November).
	Month param.Field[int64] `json:"month"`
	// The year on which the billing cycle is anchored (e.g. a 2 year billing cycle
	// anchored on 2021 would have cycles starting on 2021, 2023, 2025, etc.).
	Year param.Field[int64] `json:"year"`
}

func (r SubscriptionNewParamsBillingCycleAnchorConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsExternalMarketplace string

const (
	SubscriptionNewParamsExternalMarketplaceGoogle SubscriptionNewParamsExternalMarketplace = "google"
	SubscriptionNewParamsExternalMarketplaceAws    SubscriptionNewParamsExternalMarketplace = "aws"
	SubscriptionNewParamsExternalMarketplaceAzure  SubscriptionNewParamsExternalMarketplace = "azure"
)

func (r SubscriptionNewParamsExternalMarketplace) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsExternalMarketplaceGoogle, SubscriptionNewParamsExternalMarketplaceAws, SubscriptionNewParamsExternalMarketplaceAzure:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverride struct {
	ID        param.Field[string]                                       `json:"id,required"`
	ModelType param.Field[SubscriptionNewParamsPriceOverridesModelType] `json:"model_type,required"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64]     `json:"conversion_rate"`
	Discount       param.Field[interface{}] `json:"discount,required"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	UnitConfig                       param.Field[interface{}] `json:"unit_config,required"`
	PackageConfig                    param.Field[interface{}] `json:"package_config,required"`
	MatrixConfig                     param.Field[interface{}] `json:"matrix_config,required"`
	TieredConfig                     param.Field[interface{}] `json:"tiered_config,required"`
	TieredBpsConfig                  param.Field[interface{}] `json:"tiered_bps_config,required"`
	BpsConfig                        param.Field[interface{}] `json:"bps_config,required"`
	BulkBpsConfig                    param.Field[interface{}] `json:"bulk_bps_config,required"`
	BulkConfig                       param.Field[interface{}] `json:"bulk_config,required"`
	ThresholdTotalAmountConfig       param.Field[interface{}] `json:"threshold_total_amount_config,required"`
	TieredPackageConfig              param.Field[interface{}] `json:"tiered_package_config,required"`
	TieredWithMinimumConfig          param.Field[interface{}] `json:"tiered_with_minimum_config,required"`
	PackageWithAllocationConfig      param.Field[interface{}] `json:"package_with_allocation_config,required"`
	UnitWithPercentConfig            param.Field[interface{}] `json:"unit_with_percent_config,required"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config,required"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config,required"`
	BulkWithProrationConfig          param.Field[interface{}] `json:"bulk_with_proration_config,required"`
	UnitWithProrationConfig          param.Field[interface{}] `json:"unit_with_proration_config,required"`
	TieredWithProrationConfig        param.Field[interface{}] `json:"tiered_with_proration_config,required"`
}

func (r SubscriptionNewParamsPriceOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverride) implementsSubscriptionNewParamsPriceOverrideUnion() {}

// Satisfied by [SubscriptionNewParamsPriceOverridesOverrideUnitPrice],
// [SubscriptionNewParamsPriceOverridesOverridePackagePrice],
// [SubscriptionNewParamsPriceOverridesOverrideMatrixPrice],
// [SubscriptionNewParamsPriceOverridesOverrideTieredPrice],
// [SubscriptionNewParamsPriceOverridesOverrideTieredBpsPrice],
// [SubscriptionNewParamsPriceOverridesOverrideBpsPrice],
// [SubscriptionNewParamsPriceOverridesOverrideBulkBpsPrice],
// [SubscriptionNewParamsPriceOverridesOverrideBulkPrice],
// [SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPrice],
// [SubscriptionNewParamsPriceOverridesOverrideTieredPackagePrice],
// [SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPrice],
// [SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPrice],
// [SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPrice],
// [SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPrice],
// [SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice],
// [SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice],
// [SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPrice],
// [SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPrice],
// [SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPrice],
// [SubscriptionNewParamsPriceOverride].
type SubscriptionNewParamsPriceOverrideUnion interface {
	implementsSubscriptionNewParamsPriceOverrideUnion()
}

type SubscriptionNewParamsPriceOverridesOverrideUnitPrice struct {
	ID         param.Field[string]                                                         `json:"id,required"`
	ModelType  param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelType]  `json:"model_type,required"`
	UnitConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig] `json:"unit_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelTypeUnit SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelType = "unit"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverridePackagePrice struct {
	ID            param.Field[string]                                                               `json:"id,required"`
	ModelType     param.Field[SubscriptionNewParamsPriceOverridesOverridePackagePriceModelType]     `json:"model_type,required"`
	PackageConfig param.Field[SubscriptionNewParamsPriceOverridesOverridePackagePricePackageConfig] `json:"package_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverridePackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverridePackagePrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverridePackagePriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverridePackagePriceModelTypePackage SubscriptionNewParamsPriceOverridesOverridePackagePriceModelType = "package"
)

func (r SubscriptionNewParamsPriceOverridesOverridePackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverridePackagePriceModelTypePackage:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverridePackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverridePackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideMatrixPrice struct {
	ID           param.Field[string]                                                             `json:"id,required"`
	MatrixConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelType]    `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelTypeMatrix SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelType = "matrix"
)

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPrice struct {
	ID           param.Field[string]                                                             `json:"id,required"`
	ModelType    param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelType]    `json:"model_type,required"`
	TieredConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelTypeTiered SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelType = "tiered"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPrice struct {
	ID              param.Field[string]                                                                   `json:"id,required"`
	ModelType       param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelType]       `json:"model_type,required"`
	TieredBpsConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelTypeTieredBps SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelType = "tiered_bps"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideBpsPrice struct {
	ID        param.Field[string]                                                       `json:"id,required"`
	BpsConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideBpsPriceBpsConfig] `json:"bps_config,required"`
	ModelType param.Field[SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelTypeBps SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelType = "bps"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBpsPriceModelTypeBps:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPrice struct {
	ID            param.Field[string]                                                               `json:"id,required"`
	BulkBpsConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	ModelType     param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelType]     `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelTypeBulkBps SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelType = "bulk_bps"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideBulkPrice struct {
	ID         param.Field[string]                                                         `json:"id,required"`
	BulkConfig param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfig] `json:"bulk_config,required"`
	ModelType  param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelType]  `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelTypeBulk SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelType = "bulk"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPrice struct {
	ID                         param.Field[string]                                                                        `json:"id,required"`
	ModelType                  param.Field[SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}]                                                        `json:"threshold_total_amount_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelTypeThresholdTotalAmount SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPackagePrice struct {
	ID                  param.Field[string]                                                                 `json:"id,required"`
	ModelType           param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelType] `json:"model_type,required"`
	TieredPackageConfig param.Field[map[string]interface{}]                                                 `json:"tiered_package_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPackagePrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelTypeTieredPackage SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelType = "tiered_package"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPrice struct {
	ID                      param.Field[string]                                                                     `json:"id,required"`
	ModelType               param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelType] `json:"model_type,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}]                                                     `json:"tiered_with_minimum_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelTypeTieredWithMinimum SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPrice struct {
	ID                          param.Field[string]                                                                         `json:"id,required"`
	ModelType                   param.Field[SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelType] `json:"model_type,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}]                                                         `json:"package_with_allocation_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelTypePackageWithAllocation SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPrice struct {
	ID                    param.Field[string]                                                                   `json:"id,required"`
	ModelType             param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelType] `json:"model_type,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}]                                                   `json:"unit_with_percent_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelTypeUnitWithPercent SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPrice struct {
	ID                      param.Field[string]                                                                     `json:"id,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                     `json:"grouped_allocation_config,required"`
	ModelType               param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelTypeGroupedAllocation SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice struct {
	ID                               param.Field[string]                                                                              `json:"id,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                              `json:"grouped_with_prorated_minimum_config,required"`
	ModelType                        param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice struct {
	ID                              param.Field[string]                                                                             `json:"id,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                             `json:"grouped_with_metered_minimum_config,required"`
	ModelType                       param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPrice struct {
	ID                      param.Field[string]                                                                     `json:"id,required"`
	BulkWithProrationConfig param.Field[map[string]interface{}]                                                     `json:"bulk_with_proration_config,required"`
	ModelType               param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelTypeBulkWithProration SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPrice struct {
	ID                      param.Field[string]                                                                     `json:"id,required"`
	ModelType               param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelType] `json:"model_type,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}]                                                     `json:"unit_with_proration_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelTypeUnitWithProration SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPrice struct {
	ID                        param.Field[string]                                                                       `json:"id,required"`
	ModelType                 param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelType] `json:"model_type,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}]                                                       `json:"tiered_with_proration_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPrice) implementsSubscriptionNewParamsPriceOverrideUnion() {
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelTypeTieredWithProration SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType string

const (
	SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypePercentage SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeTrial      SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeUsage      SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeAmount     SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypePercentage, SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeTrial, SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeUsage, SubscriptionNewParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionNewParamsPriceOverridesModelType string

const (
	SubscriptionNewParamsPriceOverridesModelTypeUnit                       SubscriptionNewParamsPriceOverridesModelType = "unit"
	SubscriptionNewParamsPriceOverridesModelTypePackage                    SubscriptionNewParamsPriceOverridesModelType = "package"
	SubscriptionNewParamsPriceOverridesModelTypeMatrix                     SubscriptionNewParamsPriceOverridesModelType = "matrix"
	SubscriptionNewParamsPriceOverridesModelTypeTiered                     SubscriptionNewParamsPriceOverridesModelType = "tiered"
	SubscriptionNewParamsPriceOverridesModelTypeTieredBps                  SubscriptionNewParamsPriceOverridesModelType = "tiered_bps"
	SubscriptionNewParamsPriceOverridesModelTypeBps                        SubscriptionNewParamsPriceOverridesModelType = "bps"
	SubscriptionNewParamsPriceOverridesModelTypeBulkBps                    SubscriptionNewParamsPriceOverridesModelType = "bulk_bps"
	SubscriptionNewParamsPriceOverridesModelTypeBulk                       SubscriptionNewParamsPriceOverridesModelType = "bulk"
	SubscriptionNewParamsPriceOverridesModelTypeThresholdTotalAmount       SubscriptionNewParamsPriceOverridesModelType = "threshold_total_amount"
	SubscriptionNewParamsPriceOverridesModelTypeTieredPackage              SubscriptionNewParamsPriceOverridesModelType = "tiered_package"
	SubscriptionNewParamsPriceOverridesModelTypeTieredWithMinimum          SubscriptionNewParamsPriceOverridesModelType = "tiered_with_minimum"
	SubscriptionNewParamsPriceOverridesModelTypePackageWithAllocation      SubscriptionNewParamsPriceOverridesModelType = "package_with_allocation"
	SubscriptionNewParamsPriceOverridesModelTypeUnitWithPercent            SubscriptionNewParamsPriceOverridesModelType = "unit_with_percent"
	SubscriptionNewParamsPriceOverridesModelTypeGroupedAllocation          SubscriptionNewParamsPriceOverridesModelType = "grouped_allocation"
	SubscriptionNewParamsPriceOverridesModelTypeGroupedWithProratedMinimum SubscriptionNewParamsPriceOverridesModelType = "grouped_with_prorated_minimum"
	SubscriptionNewParamsPriceOverridesModelTypeGroupedWithMeteredMinimum  SubscriptionNewParamsPriceOverridesModelType = "grouped_with_metered_minimum"
	SubscriptionNewParamsPriceOverridesModelTypeBulkWithProration          SubscriptionNewParamsPriceOverridesModelType = "bulk_with_proration"
	SubscriptionNewParamsPriceOverridesModelTypeUnitWithProration          SubscriptionNewParamsPriceOverridesModelType = "unit_with_proration"
	SubscriptionNewParamsPriceOverridesModelTypeTieredWithProration        SubscriptionNewParamsPriceOverridesModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsPriceOverridesModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsPriceOverridesModelTypeUnit, SubscriptionNewParamsPriceOverridesModelTypePackage, SubscriptionNewParamsPriceOverridesModelTypeMatrix, SubscriptionNewParamsPriceOverridesModelTypeTiered, SubscriptionNewParamsPriceOverridesModelTypeTieredBps, SubscriptionNewParamsPriceOverridesModelTypeBps, SubscriptionNewParamsPriceOverridesModelTypeBulkBps, SubscriptionNewParamsPriceOverridesModelTypeBulk, SubscriptionNewParamsPriceOverridesModelTypeThresholdTotalAmount, SubscriptionNewParamsPriceOverridesModelTypeTieredPackage, SubscriptionNewParamsPriceOverridesModelTypeTieredWithMinimum, SubscriptionNewParamsPriceOverridesModelTypePackageWithAllocation, SubscriptionNewParamsPriceOverridesModelTypeUnitWithPercent, SubscriptionNewParamsPriceOverridesModelTypeGroupedAllocation, SubscriptionNewParamsPriceOverridesModelTypeGroupedWithProratedMinimum, SubscriptionNewParamsPriceOverridesModelTypeGroupedWithMeteredMinimum, SubscriptionNewParamsPriceOverridesModelTypeBulkWithProration, SubscriptionNewParamsPriceOverridesModelTypeUnitWithProration, SubscriptionNewParamsPriceOverridesModelTypeTieredWithProration:
		return true
	}
	return false
}

type SubscriptionNewParamsRemoveAdjustment struct {
	// The id of the adjustment to remove on the subscription.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
}

func (r SubscriptionNewParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsRemovePrice struct {
	// The external price id of the price to remove on the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The id of the price to remove on the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r SubscriptionNewParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the subscription.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
}

func (r SubscriptionNewParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionNewParamsReplaceAdjustmentsAdjustment struct {
	AppliesToPriceIDs  param.Field[interface{}]                                                     `json:"applies_to_price_ids"`
	AdjustmentType     param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                         `json:"percentage_discount"`
	AmountDiscount     param.Field[string]                                                          `json:"amount_discount"`
	MinimumAmount      param.Field[string]                                                          `json:"minimum_amount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id"`
	MaximumAmount param.Field[string] `json:"maximum_amount"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustment) implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by
// [SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount],
// [SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount],
// [SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimum],
// [SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximum],
// [SubscriptionNewParamsReplaceAdjustmentsAdjustment].
type SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion interface {
	implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion()
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscount) implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                           `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscount) implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimum) implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximum) implementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePrice struct {
	// The definition of a new price to create and add to the subscription.
	Price param.Field[SubscriptionNewParamsReplacePricesPriceUnion] `json:"price,required"`
	// The id of the price on the plan to replace in the subscription.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
}

func (r SubscriptionNewParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new price to create and add to the subscription.
type SubscriptionNewParamsReplacePricesPrice struct {
	Metadata param.Field[interface{}] `json:"metadata,required"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// The id of the item the plan will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// The cadence to bill for this price on.
	Cadence                     param.Field[SubscriptionNewParamsReplacePricesPriceCadence] `json:"cadence,required"`
	BillingCycleConfiguration   param.Field[interface{}]                                    `json:"billing_cycle_configuration,required"`
	InvoicingCycleConfiguration param.Field[interface{}]                                    `json:"invoicing_cycle_configuration,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64]                                          `json:"conversion_rate"`
	ModelType      param.Field[SubscriptionNewParamsReplacePricesPriceModelType] `json:"model_type,required"`
	UnitConfig     param.Field[interface{}]                                      `json:"unit_config,required"`
	// An ISO 4217 currency string, or custom pricing unit identifier, in which this
	// price is billed.
	Currency                         param.Field[string]      `json:"currency"`
	PackageConfig                    param.Field[interface{}] `json:"package_config,required"`
	MatrixConfig                     param.Field[interface{}] `json:"matrix_config,required"`
	TieredConfig                     param.Field[interface{}] `json:"tiered_config,required"`
	TieredBpsConfig                  param.Field[interface{}] `json:"tiered_bps_config,required"`
	BpsConfig                        param.Field[interface{}] `json:"bps_config,required"`
	BulkBpsConfig                    param.Field[interface{}] `json:"bulk_bps_config,required"`
	BulkConfig                       param.Field[interface{}] `json:"bulk_config,required"`
	ThresholdTotalAmountConfig       param.Field[interface{}] `json:"threshold_total_amount_config,required"`
	TieredPackageConfig              param.Field[interface{}] `json:"tiered_package_config,required"`
	TieredWithMinimumConfig          param.Field[interface{}] `json:"tiered_with_minimum_config,required"`
	UnitWithPercentConfig            param.Field[interface{}] `json:"unit_with_percent_config,required"`
	PackageWithAllocationConfig      param.Field[interface{}] `json:"package_with_allocation_config,required"`
	TieredWithProrationConfig        param.Field[interface{}] `json:"tiered_with_proration_config,required"`
	UnitWithProrationConfig          param.Field[interface{}] `json:"unit_with_proration_config,required"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config,required"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	BulkWithProrationConfig          param.Field[interface{}] `json:"bulk_with_proration_config,required"`
}

func (r SubscriptionNewParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The definition of a new price to create and add to the subscription.
//
// Satisfied by [SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPrice],
// [SubscriptionNewParamsReplacePricesPrice].
type SubscriptionNewParamsReplacePricesPriceUnion interface {
	implementsSubscriptionNewParamsReplacePricesPriceUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                    `json:"name,required"`
	UnitConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelType = "unit"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                          `json:"name,required"`
	PackageConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelTypePackage SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelType = "package"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                                                        `json:"item_id,required"`
	MatrixConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelTypeMatrix SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelType = "matrix"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                        `json:"name,required"`
	TieredConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                        `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                              `json:"name,required"`
	TieredBpsConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelType = "tiered_bps"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPrice struct {
	BpsConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelTypeBps SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelType = "bps"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPrice struct {
	BulkBpsConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                      `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelType = "bulk_bps"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPrice struct {
	BulkConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelTypeBulk SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelType = "bulk"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                   `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                            `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                              `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTierWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence                 param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                              `json:"grouped_allocation_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence                          param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                       `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                         `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceModelTypeUnit                       SubscriptionNewParamsReplacePricesPriceModelType = "unit"
	SubscriptionNewParamsReplacePricesPriceModelTypePackage                    SubscriptionNewParamsReplacePricesPriceModelType = "package"
	SubscriptionNewParamsReplacePricesPriceModelTypeMatrix                     SubscriptionNewParamsReplacePricesPriceModelType = "matrix"
	SubscriptionNewParamsReplacePricesPriceModelTypeTiered                     SubscriptionNewParamsReplacePricesPriceModelType = "tiered"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredBps                  SubscriptionNewParamsReplacePricesPriceModelType = "tiered_bps"
	SubscriptionNewParamsReplacePricesPriceModelTypeBps                        SubscriptionNewParamsReplacePricesPriceModelType = "bps"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulkBps                    SubscriptionNewParamsReplacePricesPriceModelType = "bulk_bps"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulk                       SubscriptionNewParamsReplacePricesPriceModelType = "bulk"
	SubscriptionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount       SubscriptionNewParamsReplacePricesPriceModelType = "threshold_total_amount"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackage              SubscriptionNewParamsReplacePricesPriceModelType = "tiered_package"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithMinimum          SubscriptionNewParamsReplacePricesPriceModelType = "tiered_with_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithPercent            SubscriptionNewParamsReplacePricesPriceModelType = "unit_with_percent"
	SubscriptionNewParamsReplacePricesPriceModelTypePackageWithAllocation      SubscriptionNewParamsReplacePricesPriceModelType = "package_with_allocation"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithProration        SubscriptionNewParamsReplacePricesPriceModelType = "tiered_with_proration"
	SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithProration          SubscriptionNewParamsReplacePricesPriceModelType = "unit_with_proration"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedAllocation          SubscriptionNewParamsReplacePricesPriceModelType = "grouped_allocation"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum SubscriptionNewParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithProration          SubscriptionNewParamsReplacePricesPriceModelType = "bulk_with_proration"
)

func (r SubscriptionNewParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceModelTypeUnit, SubscriptionNewParamsReplacePricesPriceModelTypePackage, SubscriptionNewParamsReplacePricesPriceModelTypeMatrix, SubscriptionNewParamsReplacePricesPriceModelTypeTiered, SubscriptionNewParamsReplacePricesPriceModelTypeTieredBps, SubscriptionNewParamsReplacePricesPriceModelTypeBps, SubscriptionNewParamsReplacePricesPriceModelTypeBulkBps, SubscriptionNewParamsReplacePricesPriceModelTypeBulk, SubscriptionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount, SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackage, SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithMinimum, SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithPercent, SubscriptionNewParamsReplacePricesPriceModelTypePackageWithAllocation, SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithProration, SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithProration, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedAllocation, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Determines the default memo on this subscription's invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo param.Field[string] `json:"default_invoice_memo"`
	// When this subscription's accrued usage reaches this threshold, an invoice will
	// be issued for the subscription. If not specified, invoices will only be issued
	// at the end of the billing period.
	InvoicingThreshold param.Field[string] `json:"invoicing_threshold"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms param.Field[int64] `json:"net_terms"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor             param.Field[string]   `query:"cursor"`
	CustomerID         param.Field[[]string] `query:"customer_id"`
	ExternalCustomerID param.Field[string]   `query:"external_customer_id"`
	// The number of items to fetch. Defaults to 20.
	Limit  param.Field[int64]                        `query:"limit"`
	Status param.Field[SubscriptionListParamsStatus] `query:"status"`
}

// URLQuery serializes [SubscriptionListParams]'s query parameters as `url.Values`.
func (r SubscriptionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionListParamsStatus string

const (
	SubscriptionListParamsStatusActive   SubscriptionListParamsStatus = "active"
	SubscriptionListParamsStatusEnded    SubscriptionListParamsStatus = "ended"
	SubscriptionListParamsStatusUpcoming SubscriptionListParamsStatus = "upcoming"
)

func (r SubscriptionListParamsStatus) IsKnown() bool {
	switch r {
	case SubscriptionListParamsStatusActive, SubscriptionListParamsStatusEnded, SubscriptionListParamsStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionCancelParams struct {
	// Determines the timing of subscription cancellation
	CancelOption param.Field[SubscriptionCancelParamsCancelOption] `json:"cancel_option,required"`
	// The date that the cancellation should take effect. This parameter can only be
	// passed if the `cancel_option` is `requested_date`.
	CancellationDate param.Field[time.Time] `json:"cancellation_date" format:"date-time"`
}

func (r SubscriptionCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the timing of subscription cancellation
type SubscriptionCancelParamsCancelOption string

const (
	SubscriptionCancelParamsCancelOptionEndOfSubscriptionTerm SubscriptionCancelParamsCancelOption = "end_of_subscription_term"
	SubscriptionCancelParamsCancelOptionImmediate             SubscriptionCancelParamsCancelOption = "immediate"
	SubscriptionCancelParamsCancelOptionRequestedDate         SubscriptionCancelParamsCancelOption = "requested_date"
)

func (r SubscriptionCancelParamsCancelOption) IsKnown() bool {
	switch r {
	case SubscriptionCancelParamsCancelOptionEndOfSubscriptionTerm, SubscriptionCancelParamsCancelOptionImmediate, SubscriptionCancelParamsCancelOptionRequestedDate:
		return true
	}
	return false
}

type SubscriptionFetchCostsParams struct {
	// The currency or custom pricing unit to use.
	Currency param.Field[string] `query:"currency"`
	// Costs returned are exclusive of `timeframe_end`.
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// Costs returned are inclusive of `timeframe_start`.
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
	// Controls whether Orb returns cumulative costs since the start of the billing
	// period, or incremental day-by-day costs. If your customer has minimums or
	// discounts, it's strongly recommended that you use the default cumulative
	// behavior.
	ViewMode param.Field[SubscriptionFetchCostsParamsViewMode] `query:"view_mode"`
}

// URLQuery serializes [SubscriptionFetchCostsParams]'s query parameters as
// `url.Values`.
func (r SubscriptionFetchCostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls whether Orb returns cumulative costs since the start of the billing
// period, or incremental day-by-day costs. If your customer has minimums or
// discounts, it's strongly recommended that you use the default cumulative
// behavior.
type SubscriptionFetchCostsParamsViewMode string

const (
	SubscriptionFetchCostsParamsViewModePeriodic   SubscriptionFetchCostsParamsViewMode = "periodic"
	SubscriptionFetchCostsParamsViewModeCumulative SubscriptionFetchCostsParamsViewMode = "cumulative"
)

func (r SubscriptionFetchCostsParamsViewMode) IsKnown() bool {
	switch r {
	case SubscriptionFetchCostsParamsViewModePeriodic, SubscriptionFetchCostsParamsViewModeCumulative:
		return true
	}
	return false
}

type SubscriptionFetchScheduleParams struct {
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// The number of items to fetch. Defaults to 20.
	Limit        param.Field[int64]     `query:"limit"`
	StartDateGt  param.Field[time.Time] `query:"start_date[gt]" format:"date-time"`
	StartDateGte param.Field[time.Time] `query:"start_date[gte]" format:"date-time"`
	StartDateLt  param.Field[time.Time] `query:"start_date[lt]" format:"date-time"`
	StartDateLte param.Field[time.Time] `query:"start_date[lte]" format:"date-time"`
}

// URLQuery serializes [SubscriptionFetchScheduleParams]'s query parameters as
// `url.Values`.
func (r SubscriptionFetchScheduleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SubscriptionFetchUsageParams struct {
	// When specified in conjunction with `group_by`, this parameter filters usage to a
	// single billable metric. Note that both `group_by` and `billable_metric_id` must
	// be specified together.
	BillableMetricID    param.Field[string] `query:"billable_metric_id"`
	FirstDimensionKey   param.Field[string] `query:"first_dimension_key"`
	FirstDimensionValue param.Field[string] `query:"first_dimension_value"`
	// This determines the windowing of usage reporting.
	Granularity param.Field[SubscriptionFetchUsageParamsGranularity] `query:"granularity"`
	// Groups per-price usage by the key provided.
	GroupBy              param.Field[string] `query:"group_by"`
	SecondDimensionKey   param.Field[string] `query:"second_dimension_key"`
	SecondDimensionValue param.Field[string] `query:"second_dimension_value"`
	// Usage returned is exclusive of `timeframe_end`.
	TimeframeEnd param.Field[time.Time] `query:"timeframe_end" format:"date-time"`
	// Usage returned is inclusive of `timeframe_start`.
	TimeframeStart param.Field[time.Time] `query:"timeframe_start" format:"date-time"`
	// Controls whether Orb returns cumulative usage since the start of the billing
	// period, or incremental day-by-day usage. If your customer has minimums or
	// discounts, it's strongly recommended that you use the default cumulative
	// behavior.
	ViewMode param.Field[SubscriptionFetchUsageParamsViewMode] `query:"view_mode"`
}

// URLQuery serializes [SubscriptionFetchUsageParams]'s query parameters as
// `url.Values`.
func (r SubscriptionFetchUsageParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This determines the windowing of usage reporting.
type SubscriptionFetchUsageParamsGranularity string

const (
	SubscriptionFetchUsageParamsGranularityDay SubscriptionFetchUsageParamsGranularity = "day"
)

func (r SubscriptionFetchUsageParamsGranularity) IsKnown() bool {
	switch r {
	case SubscriptionFetchUsageParamsGranularityDay:
		return true
	}
	return false
}

// Controls whether Orb returns cumulative usage since the start of the billing
// period, or incremental day-by-day usage. If your customer has minimums or
// discounts, it's strongly recommended that you use the default cumulative
// behavior.
type SubscriptionFetchUsageParamsViewMode string

const (
	SubscriptionFetchUsageParamsViewModePeriodic   SubscriptionFetchUsageParamsViewMode = "periodic"
	SubscriptionFetchUsageParamsViewModeCumulative SubscriptionFetchUsageParamsViewMode = "cumulative"
)

func (r SubscriptionFetchUsageParamsViewMode) IsKnown() bool {
	switch r {
	case SubscriptionFetchUsageParamsViewModePeriodic, SubscriptionFetchUsageParamsViewModeCumulative:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParams struct {
	// A list of price intervals to add to the subscription.
	Add param.Field[[]SubscriptionPriceIntervalsParamsAdd] `json:"add"`
	// A list of adjustments to add to the subscription.
	AddAdjustments param.Field[[]SubscriptionPriceIntervalsParamsAddAdjustment] `json:"add_adjustments"`
	// A list of price intervals to edit on the subscription.
	Edit param.Field[[]SubscriptionPriceIntervalsParamsEdit] `json:"edit"`
	// A list of adjustments to edit on the subscription.
	EditAdjustments param.Field[[]SubscriptionPriceIntervalsParamsEditAdjustment] `json:"edit_adjustments"`
}

func (r SubscriptionPriceIntervalsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAdd struct {
	// The start date of the price interval. This is the date that the price will start
	// billing on the subscription.
	StartDate param.Field[SubscriptionPriceIntervalsParamsAddStartDateUnion] `json:"start_date,required" format:"date-time"`
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[SubscriptionPriceIntervalsParamsAddAllocationPrice] `json:"allocation_price"`
	// A list of discounts to initialize on the price interval.
	Discounts param.Field[[]SubscriptionPriceIntervalsParamsAddDiscountUnion] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription.
	EndDate param.Field[SubscriptionPriceIntervalsParamsAddEndDateUnion] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// A list of fixed fee quantity transitions to initialize on the price interval.
	FixedFeeQuantityTransitions param.Field[[]SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition] `json:"fixed_fee_quantity_transitions"`
	// The maximum amount that will be billed for this price interval for a given
	// billing period.
	MaximumAmount param.Field[float64] `json:"maximum_amount"`
	// The minimum amount that will be billed for this price interval for a given
	// billing period.
	MinimumAmount param.Field[float64] `json:"minimum_amount"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[SubscriptionPriceIntervalsParamsAddPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r SubscriptionPriceIntervalsParamsAdd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The start date of the price interval. This is the date that the price will start
// billing on the subscription.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsAddStartDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddStartDateUnion()
}

// The definition of a new allocation price to create and add to the subscription.
type SubscriptionPriceIntervalsParamsAddAllocationPrice struct {
	// An amount of the currency to allocate to the customer at the specified cadence.
	Amount param.Field[string] `json:"amount,required"`
	// The cadence at which to allocate the amount to the customer.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string or a custom pricing unit identifier in which to bill
	// this price.
	Currency param.Field[string] `json:"currency,required"`
	// Whether the allocated amount should expire at the end of the cadence or roll
	// over to the next period.
	ExpiresAtEndOfCadence param.Field[bool] `json:"expires_at_end_of_cadence,required"`
}

func (r SubscriptionPriceIntervalsParamsAddAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence at which to allocate the amount to the customer.
type SubscriptionPriceIntervalsParamsAddAllocationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddAllocationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddDiscount struct {
	DiscountType param.Field[SubscriptionPriceIntervalsParamsAddDiscountsDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[float64] `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for.
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r SubscriptionPriceIntervalsParamsAddDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddDiscount) implementsSubscriptionPriceIntervalsParamsAddDiscountUnion() {
}

// Satisfied by
// [SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams],
// [SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParams],
// [SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParams],
// [SubscriptionPriceIntervalsParamsAddDiscount].
type SubscriptionPriceIntervalsParamsAddDiscountUnion interface {
	implementsSubscriptionPriceIntervalsParamsAddDiscountUnion()
}

type SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams struct {
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[float64]                                                                              `json:"amount_discount,required"`
	DiscountType   param.Field[SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountType] `json:"discount_type,required"`
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParams) implementsSubscriptionPriceIntervalsParamsAddDiscountUnion() {
}

type SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountType string

const (
	SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountType = "amount"
)

func (r SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddDiscountsAmountDiscountCreationParamsDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParams struct {
	DiscountType param.Field[SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParams) implementsSubscriptionPriceIntervalsParamsAddDiscountUnion() {
}

type SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountType string

const (
	SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountTypePercentage SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountType = "percentage"
)

func (r SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddDiscountsPercentageDiscountCreationParamsDiscountTypePercentage:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParams struct {
	DiscountType param.Field[SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for.
	UsageDiscount param.Field[float64] `json:"usage_discount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParams) implementsSubscriptionPriceIntervalsParamsAddDiscountUnion() {
}

type SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountType string

const (
	SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountTypeUsage SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountType = "usage"
)

func (r SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddDiscountsUsageDiscountCreationParamsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddDiscountsDiscountType string

const (
	SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypeAmount     SubscriptionPriceIntervalsParamsAddDiscountsDiscountType = "amount"
	SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypePercentage SubscriptionPriceIntervalsParamsAddDiscountsDiscountType = "percentage"
	SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypeUsage      SubscriptionPriceIntervalsParamsAddDiscountsDiscountType = "usage"
)

func (r SubscriptionPriceIntervalsParamsAddDiscountsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypeAmount, SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypePercentage, SubscriptionPriceIntervalsParamsAddDiscountsDiscountTypeUsage:
		return true
	}
	return false
}

// The end date of the price interval. This is the date that the price will stop
// billing on the subscription.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsAddEndDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddEndDateUnion()
}

type SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition struct {
	// The date that the fixed fee quantity transition should take effect.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date-time"`
	// The quantity of the fixed fee quantity transition.
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new price to create and add to the subscription.
type SubscriptionPriceIntervalsParamsAddPrice struct {
	Metadata param.Field[interface{}] `json:"metadata,required"`
	// An alias for the price.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// The id of the item the plan will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// If the Price represents a fixed cost, this represents the quantity of units
	// applied.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The property used to group this price on an invoice
	InvoiceGroupingKey param.Field[string] `json:"invoice_grouping_key"`
	// The cadence to bill for this price on.
	Cadence                     param.Field[SubscriptionPriceIntervalsParamsAddPriceCadence] `json:"cadence,required"`
	BillingCycleConfiguration   param.Field[interface{}]                                     `json:"billing_cycle_configuration,required"`
	InvoicingCycleConfiguration param.Field[interface{}]                                     `json:"invoicing_cycle_configuration,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64]                                           `json:"conversion_rate"`
	ModelType      param.Field[SubscriptionPriceIntervalsParamsAddPriceModelType] `json:"model_type,required"`
	UnitConfig     param.Field[interface{}]                                       `json:"unit_config,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                         param.Field[string]      `json:"currency,required"`
	PackageConfig                    param.Field[interface{}] `json:"package_config,required"`
	MatrixConfig                     param.Field[interface{}] `json:"matrix_config,required"`
	MatrixWithAllocationConfig       param.Field[interface{}] `json:"matrix_with_allocation_config,required"`
	TieredConfig                     param.Field[interface{}] `json:"tiered_config,required"`
	TieredBpsConfig                  param.Field[interface{}] `json:"tiered_bps_config,required"`
	BpsConfig                        param.Field[interface{}] `json:"bps_config,required"`
	BulkBpsConfig                    param.Field[interface{}] `json:"bulk_bps_config,required"`
	BulkConfig                       param.Field[interface{}] `json:"bulk_config,required"`
	ThresholdTotalAmountConfig       param.Field[interface{}] `json:"threshold_total_amount_config,required"`
	TieredPackageConfig              param.Field[interface{}] `json:"tiered_package_config,required"`
	GroupedTieredConfig              param.Field[interface{}] `json:"grouped_tiered_config,required"`
	TieredWithMinimumConfig          param.Field[interface{}] `json:"tiered_with_minimum_config,required"`
	PackageWithAllocationConfig      param.Field[interface{}] `json:"package_with_allocation_config,required"`
	TieredPackageWithMinimumConfig   param.Field[interface{}] `json:"tiered_package_with_minimum_config,required"`
	UnitWithPercentConfig            param.Field[interface{}] `json:"unit_with_percent_config,required"`
	TieredWithProrationConfig        param.Field[interface{}] `json:"tiered_with_proration_config,required"`
	UnitWithProrationConfig          param.Field[interface{}] `json:"unit_with_proration_config,required"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config,required"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config,required"`
	BulkWithProrationConfig          param.Field[interface{}] `json:"bulk_with_proration_config,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The definition of a new price to create and add to the subscription.
//
// Satisfied by [SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPrice],
// [SubscriptionPriceIntervalsParamsAddPrice].
type SubscriptionPriceIntervalsParamsAddPriceUnion interface {
	implementsSubscriptionPriceIntervalsParamsAddPriceUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name       param.Field[string]                                                                 `json:"name,required"`
	UnitConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig] `json:"unit_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelType = "unit"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name          param.Field[string]                                                                       `json:"name,required"`
	PackageConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePricePackageConfig] `json:"package_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelTypePackage SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelType = "package"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceModelTypePackage:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID       param.Field[string]                                                                     `json:"item_id,required"`
	MatrixConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelType]    `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelTypeMatrix SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelType = "matrix"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID                     param.Field[string]                                                                                                 `json:"item_id,required"`
	MatrixWithAllocationConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig] `json:"matrix_with_allocation_config,required"`
	ModelType                  param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelType]                  `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig struct {
	// Allocation to be used to calculate the price
	Allocation param.Field[float64] `json:"allocation,required"`
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceMatrixWithAllocationConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingMatrixWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                  `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name         param.Field[string]                                                                     `json:"name,required"`
	TieredConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelTypeTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelType = "tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                     `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name            param.Field[string]                                                                           `json:"name,required"`
	TieredBpsConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The id of the billable metric for the price. Only needed if the price is
	// usage-based.
	BillableMetricID param.Field[string] `json:"billable_metric_id"`
	// If the Price represents a fixed cost, the price will be billed in-advance if
	// this is true, and in-arrears if this is false.
	BilledInAdvance param.Field[bool] `json:"billed_in_advance"`
	// For custom cadence: specifies the duration of the billing period in days or
	// months.
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelTypeTieredBps SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelType = "tiered_bps"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPrice struct {
	BpsConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBpsConfig] `json:"bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                               `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelTypeBps SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelType = "bps"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceModelTypeBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPrice struct {
	BulkBpsConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                   `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelTypeBulkBps SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelType = "bulk_bps"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkBpsPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPrice struct {
	BulkConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfig] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelTypeBulk SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelType = "bulk"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingThresholdTotalAmountPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                         `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelTypeTieredPackage SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelType = "tiered_package"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackagePriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency            param.Field[string]                 `json:"currency,required"`
	GroupedTieredConfig param.Field[map[string]interface{}] `json:"grouped_tiered_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                         `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelTypeGroupedTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelType = "grouped_tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedTieredPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                 `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPackageWithAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                    `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredPackageWithMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                           `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithPercentPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                               `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelTypeTieredWithProration SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingTieredWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelTypeUnitWithProration SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingUnitWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                param.Field[string]                 `json:"currency,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}] `json:"grouped_allocation_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedAllocationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                         param.Field[string]                 `json:"currency,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                      `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithProratedMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency                        param.Field[string]                 `json:"currency,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}] `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                                     `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMeteredMinimumPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPrice struct {
	BulkWithProrationConfig param.Field[map[string]interface{}] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the plan will be associated with.
	ItemID    param.Field[string]                                                                             `json:"item_id,required"`
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	BillingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfiguration] `json:"billing_cycle_configuration"`
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
	InvoicingCycleConfiguration param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration] `json:"invoicing_cycle_configuration"`
	// User-specified key/value pairs for the resource. Individual keys can be removed
	// by setting the value to `null`, and the entire metadata mapping can be cleared
	// by setting `metadata` to `null`.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPrice) implementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelTypeBulkWithProration SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// For custom cadence: specifies the duration of the billing period in days or
// months.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceBillingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// Within each billing cycle, specifies the cadence at which invoices are produced.
// If unspecified, a single invoice is produced per billing cycle.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration struct {
	// The duration of the billing period.
	Duration param.Field[int64] `json:"duration,required"`
	// The unit of billing period duration.
	DurationUnit param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit] `json:"duration_unit,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The unit of billing period duration.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "day"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit = "month"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnit) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitDay, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithProrationPriceInvoicingCycleConfigurationDurationUnitMonth:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceCadenceCustom:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnit                       SubscriptionPriceIntervalsParamsAddPriceModelType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceModelTypePackage                    SubscriptionPriceIntervalsParamsAddPriceModelType = "package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrix                     SubscriptionPriceIntervalsParamsAddPriceModelType = "matrix"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithAllocation       SubscriptionPriceIntervalsParamsAddPriceModelType = "matrix_with_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTiered                     SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredBps                  SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_bps"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBps                        SubscriptionPriceIntervalsParamsAddPriceModelType = "bps"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkBps                    SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk_bps"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulk                       SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeThresholdTotalAmount       SubscriptionPriceIntervalsParamsAddPriceModelType = "threshold_total_amount"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackage              SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTiered              SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_tiered"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithMinimum          SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_with_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypePackageWithAllocation      SubscriptionPriceIntervalsParamsAddPriceModelType = "package_with_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackageWithMinimum   SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_package_with_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithPercent            SubscriptionPriceIntervalsParamsAddPriceModelType = "unit_with_percent"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithProration        SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_with_proration"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithProration          SubscriptionPriceIntervalsParamsAddPriceModelType = "unit_with_proration"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedAllocation          SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithProratedMinimum SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMeteredMinimum  SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_with_metered_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithProration          SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk_with_proration"
)

func (r SubscriptionPriceIntervalsParamsAddPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceModelTypeUnit, SubscriptionPriceIntervalsParamsAddPriceModelTypePackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrix, SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeTiered, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredBps, SubscriptionPriceIntervalsParamsAddPriceModelTypeBps, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkBps, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulk, SubscriptionPriceIntervalsParamsAddPriceModelTypeThresholdTotalAmount, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTiered, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypePackageWithAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackageWithMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithPercent, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithProration, SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithProration, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithProratedMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMeteredMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription.
	StartDate param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion] `json:"start_date,required" format:"date-time"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription.
	EndDate param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion] `json:"end_date" format:"date-time"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment struct {
	AppliesToPriceIDs  param.Field[interface{}]                                                            `json:"applies_to_price_ids"`
	AdjustmentType     param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	PercentageDiscount param.Field[float64]                                                                `json:"percentage_discount"`
	AmountDiscount     param.Field[string]                                                                 `json:"amount_discount"`
	MinimumAmount      param.Field[string]                                                                 `json:"minimum_amount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id"`
	MaximumAmount param.Field[string] `json:"maximum_amount"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment) implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount],
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscount],
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimum],
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximum],
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment].
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion interface {
	implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion()
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount struct {
	AdjustmentType param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs  param.Field[[]string] `json:"applies_to_price_ids,required"`
	PercentageDiscount param.Field[float64]  `json:"percentage_discount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscount) implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType = "percentage_discount"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewPercentageDiscountAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscount struct {
	AdjustmentType param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                                  `json:"amount_discount,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscount) implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType = "amount_discount"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewAmountDiscountAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimum struct {
	AdjustmentType param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID        param.Field[string] `json:"item_id,required"`
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimum) implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType = "minimum"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMinimumAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximum struct {
	AdjustmentType param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType] `json:"adjustment_type,required"`
	// The set of price IDs to which this adjustment applies.
	AppliesToPriceIDs param.Field[[]string] `json:"applies_to_price_ids,required"`
	MaximumAmount     param.Field[string]   `json:"maximum_amount,required"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximum) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximum) implementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType = "maximum"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentNewMaximumAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// The start date of the adjustment interval. This is the date that the adjustment
// will start affecting prices on the subscription.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion()
}

// The end date of the adjustment interval. This is the date that the adjustment
// will stop affecting prices on the subscription.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion()
}

type SubscriptionPriceIntervalsParamsEdit struct {
	// The id of the price interval to edit.
	PriceIntervalID param.Field[string] `json:"price_interval_id,required"`
	// The updated billing cycle day for this price interval. If not specified, the
	// billing cycle day will not be updated. Note that overlapping price intervals
	// must have the same billing cycle day.
	BillingCycleDay param.Field[int64] `json:"billing_cycle_day"`
	// The updated end date of this price interval. If not specified, the start date
	// will not be updated.
	EndDate param.Field[SubscriptionPriceIntervalsParamsEditEndDateUnion] `json:"end_date" format:"date-time"`
	// A list of fixed fee quantity transitions to use for this price interval. Note
	// that this list will overwrite all existing fixed fee quantity transitions on the
	// price interval.
	FixedFeeQuantityTransitions param.Field[[]SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition] `json:"fixed_fee_quantity_transitions"`
	// The updated start date of this price interval. If not specified, the start date
	// will not be updated.
	StartDate param.Field[SubscriptionPriceIntervalsParamsEditStartDateUnion] `json:"start_date" format:"date-time"`
}

func (r SubscriptionPriceIntervalsParamsEdit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The updated end date of this price interval. If not specified, the start date
// will not be updated.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsEditEndDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsEditEndDateUnion()
}

type SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition struct {
	// The date that the fixed fee quantity transition should take effect.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date-time"`
	// The quantity of the fixed fee quantity transition.
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The updated start date of this price interval. If not specified, the start date
// will not be updated.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsEditStartDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsEditStartDateUnion()
}

type SubscriptionPriceIntervalsParamsEditAdjustment struct {
	// The id of the adjustment interval to edit.
	AdjustmentIntervalID param.Field[string] `json:"adjustment_interval_id,required"`
	// The updated end date of this adjustment interval. If not specified, the start
	// date will not be updated.
	EndDate param.Field[SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion] `json:"end_date" format:"date-time"`
	// The updated start date of this adjustment interval. If not specified, the start
	// date will not be updated.
	StartDate param.Field[SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion] `json:"start_date" format:"date-time"`
}

func (r SubscriptionPriceIntervalsParamsEditAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The updated end date of this adjustment interval. If not specified, the start
// date will not be updated.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion()
}

// The updated start date of this adjustment interval. If not specified, the start
// date will not be updated.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion()
}

type SubscriptionSchedulePlanChangeParams struct {
	ChangeOption param.Field[SubscriptionSchedulePlanChangeParamsChangeOption] `json:"change_option,required"`
	// [DEPRECATED] Use billing_cycle_alignment instead. Reset billing periods to be
	// aligned with the plan change's effective date.
	AlignBillingWithPlanChangeDate param.Field[bool] `json:"align_billing_with_plan_change_date"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. If not specified, this
	// defaults to the behavior configured for this customer.
	AutoCollection param.Field[bool] `json:"auto_collection"`
	// Reset billing periods to be aligned with the plan change's effective date or
	// start of the month. Defaults to `unchanged` which keeps subscription's existing
	// billing cycle alignment.
	BillingCycleAlignment param.Field[SubscriptionSchedulePlanChangeParamsBillingCycleAlignment] `json:"billing_cycle_alignment"`
	// The date that the plan change should take effect. This parameter can only be
	// passed if the `change_option` is `requested_date`.
	ChangeDate param.Field[time.Time] `json:"change_date" format:"date-time"`
	// Redemption code to be used for this subscription. If the coupon cannot be found
	// by its redemption code, or cannot be redeemed, an error response will be
	// returned and the subscription creation or plan change will not be scheduled.
	CouponRedemptionCode param.Field[string]  `json:"coupon_redemption_code"`
	CreditsOverageRate   param.Field[float64] `json:"credits_overage_rate"`
	// Determines the default memo on this subscription's invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo param.Field[string] `json:"default_invoice_memo"`
	// The external_plan_id of the plan that the given subscription should be switched
	// to. Note that either this property or `plan_id` must be specified.
	ExternalPlanID param.Field[string] `json:"external_plan_id"`
	// The phase of the plan to start with
	InitialPhaseOrder param.Field[int64] `json:"initial_phase_order"`
	// When this subscription's accrued usage reaches this threshold, an invoice will
	// be issued for the subscription. If not specified, invoices will only be issued
	// at the end of the billing period.
	InvoicingThreshold param.Field[string] `json:"invoicing_threshold"`
	// The net terms determines the difference between the invoice date and the issue
	// date for the invoice. If you intend the invoice to be due on issue, set this
	// to 0. If not provided, this defaults to the value specified in the plan.
	NetTerms               param.Field[int64]   `json:"net_terms"`
	PerCreditOverageAmount param.Field[float64] `json:"per_credit_overage_amount"`
	// The plan that the given subscription should be switched to. Note that either
	// this property or `external_plan_id` must be specified.
	PlanID param.Field[string] `json:"plan_id"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverrideUnion] `json:"price_overrides"`
}

func (r SubscriptionSchedulePlanChangeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsChangeOption string

const (
	SubscriptionSchedulePlanChangeParamsChangeOptionRequestedDate         SubscriptionSchedulePlanChangeParamsChangeOption = "requested_date"
	SubscriptionSchedulePlanChangeParamsChangeOptionEndOfSubscriptionTerm SubscriptionSchedulePlanChangeParamsChangeOption = "end_of_subscription_term"
	SubscriptionSchedulePlanChangeParamsChangeOptionImmediate             SubscriptionSchedulePlanChangeParamsChangeOption = "immediate"
)

func (r SubscriptionSchedulePlanChangeParamsChangeOption) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsChangeOptionRequestedDate, SubscriptionSchedulePlanChangeParamsChangeOptionEndOfSubscriptionTerm, SubscriptionSchedulePlanChangeParamsChangeOptionImmediate:
		return true
	}
	return false
}

// Reset billing periods to be aligned with the plan change's effective date or
// start of the month. Defaults to `unchanged` which keeps subscription's existing
// billing cycle alignment.
type SubscriptionSchedulePlanChangeParamsBillingCycleAlignment string

const (
	SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentUnchanged      SubscriptionSchedulePlanChangeParamsBillingCycleAlignment = "unchanged"
	SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentPlanChangeDate SubscriptionSchedulePlanChangeParamsBillingCycleAlignment = "plan_change_date"
	SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentStartOfMonth   SubscriptionSchedulePlanChangeParamsBillingCycleAlignment = "start_of_month"
)

func (r SubscriptionSchedulePlanChangeParamsBillingCycleAlignment) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentUnchanged, SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentPlanChangeDate, SubscriptionSchedulePlanChangeParamsBillingCycleAlignmentStartOfMonth:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverride struct {
	ID        param.Field[string]                                                      `json:"id,required"`
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesModelType] `json:"model_type,required"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64]     `json:"conversion_rate"`
	Discount       param.Field[interface{}] `json:"discount,required"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity               param.Field[float64]     `json:"fixed_price_quantity"`
	UnitConfig                       param.Field[interface{}] `json:"unit_config,required"`
	PackageConfig                    param.Field[interface{}] `json:"package_config,required"`
	MatrixConfig                     param.Field[interface{}] `json:"matrix_config,required"`
	TieredConfig                     param.Field[interface{}] `json:"tiered_config,required"`
	TieredBpsConfig                  param.Field[interface{}] `json:"tiered_bps_config,required"`
	BpsConfig                        param.Field[interface{}] `json:"bps_config,required"`
	BulkBpsConfig                    param.Field[interface{}] `json:"bulk_bps_config,required"`
	BulkConfig                       param.Field[interface{}] `json:"bulk_config,required"`
	ThresholdTotalAmountConfig       param.Field[interface{}] `json:"threshold_total_amount_config,required"`
	TieredPackageConfig              param.Field[interface{}] `json:"tiered_package_config,required"`
	TieredWithMinimumConfig          param.Field[interface{}] `json:"tiered_with_minimum_config,required"`
	PackageWithAllocationConfig      param.Field[interface{}] `json:"package_with_allocation_config,required"`
	UnitWithPercentConfig            param.Field[interface{}] `json:"unit_with_percent_config,required"`
	GroupedAllocationConfig          param.Field[interface{}] `json:"grouped_allocation_config,required"`
	GroupedWithProratedMinimumConfig param.Field[interface{}] `json:"grouped_with_prorated_minimum_config,required"`
	GroupedWithMeteredMinimumConfig  param.Field[interface{}] `json:"grouped_with_metered_minimum_config,required"`
	BulkWithProrationConfig          param.Field[interface{}] `json:"bulk_with_proration_config,required"`
	UnitWithProrationConfig          param.Field[interface{}] `json:"unit_with_proration_config,required"`
	TieredWithProrationConfig        param.Field[interface{}] `json:"tiered_with_proration_config,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverride) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

// Satisfied by
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPrice],
// [SubscriptionSchedulePlanChangeParamsPriceOverride].
type SubscriptionSchedulePlanChangeParamsPriceOverrideUnion interface {
	implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion()
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice struct {
	ID         param.Field[string]                                                                        `json:"id,required"`
	ModelType  param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelType]  `json:"model_type,required"`
	UnitConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig] `json:"unit_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelTypeUnit SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelType = "unit"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceUnitConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePrice struct {
	ID            param.Field[string]                                                                              `json:"id,required"`
	ModelType     param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelType]     `json:"model_type,required"`
	PackageConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePricePackageConfig] `json:"package_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelTypePackage SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelType = "package"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceModelTypePackage:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePricePackageConfig struct {
	// A currency amount to rate usage by
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// An integer amount to represent package size. For example, 1000 here would divide
	// usage by 1000 before multiplying by package_amount in rating
	PackageSize param.Field[int64] `json:"package_size,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePricePackageConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackagePriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPrice struct {
	ID           param.Field[string]                                                                            `json:"id,required"`
	MatrixConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfig] `json:"matrix_config,required"`
	ModelType    param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelType]    `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfig struct {
	// Default per unit rate for any usage not bucketed into a specified matrix_value
	DefaultUnitAmount param.Field[string] `json:"default_unit_amount,required"`
	// One or two event property values to evaluate matrix groups by
	Dimensions param.Field[[]string] `json:"dimensions,required"`
	// Matrix values for specified matrix grouping keys
	MatrixValues param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue] `json:"matrix_values,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue struct {
	// One or two matrix keys to filter usage to this Matrix value by. For example,
	// ["region", "tier"] could be used to filter cloud usage by a cloud region and an
	// instance tier.
	DimensionValues param.Field[[]string] `json:"dimension_values,required"`
	// Unit price for the specified dimension_values
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceMatrixConfigMatrixValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelTypeMatrix SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelType = "matrix"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideMatrixPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPrice struct {
	ID           param.Field[string]                                                                            `json:"id,required"`
	ModelType    param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelType]    `json:"model_type,required"`
	TieredConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfig] `json:"tiered_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelTypeTiered SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfigTier struct {
	// Inclusive tier starting value
	FirstUnit param.Field[float64] `json:"first_unit,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit param.Field[float64] `json:"last_unit"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceTieredConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPrice struct {
	ID              param.Field[string]                                                                                  `json:"id,required"`
	ModelType       param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelType]       `json:"model_type,required"`
	TieredBpsConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig] `json:"tiered_bps_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelTypeTieredBps SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelType = "tiered_bps"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceModelTypeTieredBps:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig struct {
	// Tiers for a Graduated BPS pricing model, where usage is bucketed into specified
	// tiers
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier struct {
	// Per-event basis point rate
	Bps param.Field[float64] `json:"bps,required"`
	// Inclusive tier starting value
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Exclusive tier ending value
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// Per unit maximum to charge
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceTieredBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPrice struct {
	ID        param.Field[string]                                                                      `json:"id,required"`
	BpsConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceBpsConfig] `json:"bps_config,required"`
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceBpsConfig struct {
	// Basis point take rate per event
	Bps param.Field[float64] `json:"bps,required"`
	// Optional currency amount maximum to cap spend per event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelTypeBps SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelType = "bps"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceModelTypeBps:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPrice struct {
	ID            param.Field[string]                                                                              `json:"id,required"`
	BulkBpsConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig] `json:"bulk_bps_config,required"`
	ModelType     param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelType]     `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig struct {
	// Tiers for a bulk BPS pricing model where all usage is aggregated to a single
	// tier based on total volume
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier struct {
	// Basis points to rate on
	Bps param.Field[float64] `json:"bps,required"`
	// Upper bound for tier
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The maximum amount to charge for any one event
	PerUnitMaximum param.Field[string] `json:"per_unit_maximum"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceBulkBpsConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelTypeBulkBps SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelType = "bulk_bps"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceModelTypeBulkBps:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkBpsPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPrice struct {
	ID         param.Field[string]                                                                        `json:"id,required"`
	BulkConfig param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfig] `json:"bulk_config,required"`
	ModelType  param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelType]  `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfig struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// Upper bound for this tier
	MaximumUnits param.Field[float64] `json:"maximum_units"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceBulkConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelTypeBulk SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelType = "bulk"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceModelTypeBulk:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPrice struct {
	ID                         param.Field[string]                                                                                       `json:"id,required"`
	ModelType                  param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	ThresholdTotalAmountConfig param.Field[map[string]interface{}]                                                                       `json:"threshold_total_amount_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelTypeThresholdTotalAmount SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideThresholdTotalAmountPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePrice struct {
	ID                  param.Field[string]                                                                                `json:"id,required"`
	ModelType           param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelType] `json:"model_type,required"`
	TieredPackageConfig param.Field[map[string]interface{}]                                                                `json:"tiered_package_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelTypeTieredPackage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelType = "tiered_package"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredPackagePriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPrice struct {
	ID                      param.Field[string]                                                                                    `json:"id,required"`
	ModelType               param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelType] `json:"model_type,required"`
	TieredWithMinimumConfig param.Field[map[string]interface{}]                                                                    `json:"tiered_with_minimum_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelTypeTieredWithMinimum SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPrice struct {
	ID                          param.Field[string]                                                                                        `json:"id,required"`
	ModelType                   param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelType] `json:"model_type,required"`
	PackageWithAllocationConfig param.Field[map[string]interface{}]                                                                        `json:"package_with_allocation_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelTypePackageWithAllocation SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverridePackageWithAllocationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPrice struct {
	ID                    param.Field[string]                                                                                  `json:"id,required"`
	ModelType             param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelType] `json:"model_type,required"`
	UnitWithPercentConfig param.Field[map[string]interface{}]                                                                  `json:"unit_with_percent_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelTypeUnitWithPercent SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithPercentPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPrice struct {
	ID                      param.Field[string]                                                                                    `json:"id,required"`
	GroupedAllocationConfig param.Field[map[string]interface{}]                                                                    `json:"grouped_allocation_config,required"`
	ModelType               param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelTypeGroupedAllocation SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedAllocationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice struct {
	ID                               param.Field[string]                                                                                             `json:"id,required"`
	GroupedWithProratedMinimumConfig param.Field[map[string]interface{}]                                                                             `json:"grouped_with_prorated_minimum_config,required"`
	ModelType                        param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithProratedMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice struct {
	ID                              param.Field[string]                                                                                            `json:"id,required"`
	GroupedWithMeteredMinimumConfig param.Field[map[string]interface{}]                                                                            `json:"grouped_with_metered_minimum_config,required"`
	ModelType                       param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideGroupedWithMeteredMinimumPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPrice struct {
	ID                      param.Field[string]                                                                                    `json:"id,required"`
	BulkWithProrationConfig param.Field[map[string]interface{}]                                                                    `json:"bulk_with_proration_config,required"`
	ModelType               param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelType] `json:"model_type,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelTypeBulkWithProration SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideBulkWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPrice struct {
	ID                      param.Field[string]                                                                                    `json:"id,required"`
	ModelType               param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelType] `json:"model_type,required"`
	UnitWithProrationConfig param.Field[map[string]interface{}]                                                                    `json:"unit_with_proration_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelTypeUnitWithProration SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideUnitWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPrice struct {
	ID                        param.Field[string]                                                                                      `json:"id,required"`
	ModelType                 param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelType] `json:"model_type,required"`
	TieredWithProrationConfig param.Field[map[string]interface{}]                                                                      `json:"tiered_with_proration_config,required"`
	// The per unit conversion rate of the price currency to the invoicing currency.
	ConversionRate param.Field[float64] `json:"conversion_rate"`
	// The currency of the price. If not provided, the currency of the plan will be
	// used.
	Currency param.Field[string] `json:"currency"`
	// The subscription's override discount for the plan.
	Discount param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscount] `json:"discount"`
	// The starting quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// The subscription's override maximum amount for the plan.
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// The subscription's override minimum amount for the plan.
	MinimumAmount param.Field[string] `json:"minimum_amount"`
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPrice) implementsSubscriptionSchedulePlanChangeParamsPriceOverrideUnion() {
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelTypeTieredWithProration SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// The subscription's override discount for the plan.
type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscount struct {
	DiscountType param.Field[SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType] `json:"discount_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypePercentage SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "percentage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeTrial      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "trial"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeUsage      SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "usage"
	SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeAmount     SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType = "amount"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypePercentage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeTrial, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeUsage, SubscriptionSchedulePlanChangeParamsPriceOverridesOverrideTieredWithProrationPriceDiscountDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsPriceOverridesModelType string

const (
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnit                       SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "unit"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypePackage                    SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "package"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeMatrix                     SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "matrix"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTiered                     SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "tiered"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredBps                  SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "tiered_bps"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBps                        SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "bps"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulkBps                    SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "bulk_bps"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulk                       SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "bulk"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeThresholdTotalAmount       SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "threshold_total_amount"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredPackage              SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "tiered_package"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredWithMinimum          SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "tiered_with_minimum"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypePackageWithAllocation      SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "package_with_allocation"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnitWithPercent            SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "unit_with_percent"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedAllocation          SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "grouped_allocation"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedWithProratedMinimum SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "grouped_with_prorated_minimum"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedWithMeteredMinimum  SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "grouped_with_metered_minimum"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulkWithProration          SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "bulk_with_proration"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnitWithProration          SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "unit_with_proration"
	SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredWithProration        SubscriptionSchedulePlanChangeParamsPriceOverridesModelType = "tiered_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsPriceOverridesModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnit, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypePackage, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeMatrix, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTiered, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredBps, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBps, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulkBps, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulk, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeThresholdTotalAmount, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredPackage, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredWithMinimum, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypePackageWithAllocation, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnitWithPercent, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedAllocation, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedWithProratedMinimum, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeGroupedWithMeteredMinimum, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeBulkWithProration, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeUnitWithProration, SubscriptionSchedulePlanChangeParamsPriceOverridesModelTypeTieredWithProration:
		return true
	}
	return false
}

type SubscriptionTriggerPhaseParams struct {
	// The date on which the phase change should take effect. If not provided, defaults
	// to today in the customer's timezone.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
}

func (r SubscriptionTriggerPhaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUnscheduleFixedFeeQuantityUpdatesParams struct {
	// Price for which the updates should be cleared. Must be a fixed fee.
	PriceID param.Field[string] `json:"price_id,required"`
}

func (r SubscriptionUnscheduleFixedFeeQuantityUpdatesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionUpdateFixedFeeQuantityParams struct {
	// Price for which the quantity should be updated. Must be a fixed fee.
	PriceID  param.Field[string]  `json:"price_id,required"`
	Quantity param.Field[float64] `json:"quantity,required"`
	// Determines when the change takes effect. Note that if `effective_date` is
	// specified, this defaults to `effective_date`. Otherwise, this defaults to
	// `immediate` unless it's explicitly set to `upcoming_invoice.
	ChangeOption param.Field[SubscriptionUpdateFixedFeeQuantityParamsChangeOption] `json:"change_option"`
	// The date that the quantity change should take effect, localized to the
	// customer's timezone. Ifthis parameter is not passed in, the quantity change is
	// effective according to `change_option`.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
}

func (r SubscriptionUpdateFixedFeeQuantityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines when the change takes effect. Note that if `effective_date` is
// specified, this defaults to `effective_date`. Otherwise, this defaults to
// `immediate` unless it's explicitly set to `upcoming_invoice.
type SubscriptionUpdateFixedFeeQuantityParamsChangeOption string

const (
	SubscriptionUpdateFixedFeeQuantityParamsChangeOptionImmediate       SubscriptionUpdateFixedFeeQuantityParamsChangeOption = "immediate"
	SubscriptionUpdateFixedFeeQuantityParamsChangeOptionUpcomingInvoice SubscriptionUpdateFixedFeeQuantityParamsChangeOption = "upcoming_invoice"
	SubscriptionUpdateFixedFeeQuantityParamsChangeOptionEffectiveDate   SubscriptionUpdateFixedFeeQuantityParamsChangeOption = "effective_date"
)

func (r SubscriptionUpdateFixedFeeQuantityParamsChangeOption) IsKnown() bool {
	switch r {
	case SubscriptionUpdateFixedFeeQuantityParamsChangeOptionImmediate, SubscriptionUpdateFixedFeeQuantityParamsChangeOptionUpcomingInvoice, SubscriptionUpdateFixedFeeQuantityParamsChangeOptionEffectiveDate:
		return true
	}
	return false
}
