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
// [Subscription](/core-concepts##subscription) for more details).
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
// ## Customize your customer's subscriptions
//
// Prices and adjustments in a plan can be added, removed, or replaced for the
// subscription being created. This is useful when a customer has prices that
// differ from the default prices for a specific plan.
//
// <Note>
// This feature is only available for accounts that have migrated to Subscription Overrides Version 2. You can find your
// Subscription Overrides Version at the bottom of your [Plans page](https://app.withorb.com/plans)
// </Note>
//
// ### Adding Prices
//
// To add prices, provide a list of objects with the key `add_prices`. An object in
// the list must specify an existing add-on price with a `price_id` or
// `external_price_id` field, or create a new add-on price by including an object
// with the key `price`, identical to what would be used in the request body for
// the [create price endpoint](/api-reference/price/create-price). See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations possible in this object.
//
// If the plan has phases, each object in the list must include a number with
// `plan_phase_order` key to indicate which phase the price should be added to.
//
// An object in the list can specify an optional `start_date` and optional
// `end_date`. This is equivalent to creating a price interval with the
// [add/edit price intervals endpoint](/api-reference/price-interval/add-or-edit-price-intervals).
// If unspecified, the start or end date of the phase or subscription will be used.
//
// An object in the list can specify an optional `minimum_amount`,
// `maximum_amount`, or `discounts`. This will create adjustments which apply only
// to this price.
//
// Additionally, an object in the list can specify an optional `reference_id`. This
// ID can be used to reference this price when
// [adding an adjustment](#adding-adjustments) in the same API call. However the ID
// is _transient_ and cannot be used to refer to the price in future API calls.
//
// ### Removing Prices
//
// To remove prices, provide a list of objects with the key `remove_prices`. An
// object in the list must specify a plan price with either a `price_id` or
// `external_price_id` field.
//
// ### Replacing Prices
//
// To replace prices, provide a list of objects with the key `replace_prices`. An
// object in the list must specify a plan price to replace with the
// `replaces_price_id` key, and it must specify a price to replace it with by
// either referencing an existing add-on price with a `price_id` or
// `external_price_id` field, or by creating a new add-on price by including an
// object with the key `price`, identical to what would be used in the request body
// for the [create price endpoint](/api-reference/price/create-price). See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations possible in this object.
//
// For fixed fees, an object in the list can supply a `fixed_price_quantity`
// instead of a `price`, `price_id`, or `external_price_id` field. This will update
// only the quantity for the price, similar to the
// [Update price quantity](/api-reference/subscription/update-price-quantity)
// endpoint.
//
// The replacement price will have the same phase, if applicable, and the same
// start and end dates as the price it replaces.
//
// An object in the list can specify an optional `minimum_amount`,
// `maximum_amount`, or `discounts`. This will create adjustments which apply only
// to this price.
//
// Additionally, an object in the list can specify an optional `reference_id`. This
// ID can be used to reference the replacement price when
// [adding an adjustment](#adding-adjustments) in the same API call. However the ID
// is _transient_ and cannot be used to refer to the price in future API calls.
//
// ### Adding adjustments
//
// To add adjustments, provide a list of objects with the key `add_adjustments`. An
// object in the list must include an object with the key `adjustment`, identical
// to the adjustment object in the
// [add/edit price intervals endpoint](/api-reference/price-interval/add-or-edit-price-intervals).
//
// If the plan has phases, each object in the list must include a number with
// `plan_phase_order` key to indicate which phase the adjustment should be added
// to.
//
// An object in the list can specify an optional `start_date` and optional
// `end_date`. If unspecified, the start or end date of the phase or subscription
// will be used.
//
// ### Removing adjustments
//
// To remove adjustments, provide a list of objects with the key
// `remove_adjustments`. An object in the list must include a key, `adjustment_id`,
// with the ID of the adjustment to be removed.
//
// ### Replacing adjustments
//
// To replace adjustments, provide a list of objects with the key
// `replace_adjustments`. An object in the list must specify a plan adjustment to
// replace with the `replaces_adjustment_id` key, and it must specify an adjustment
// to replace it with by including an object with the key `adjustment`, identical
// to the adjustment object in the
// [add/edit price intervals endpoint](/api-reference/price-interval/add-or-edit-price-intervals).
//
// The replacement adjustment will have the same phase, if applicable, and the same
// start and end dates as the adjustment it replaces.
//
// ## Price overrides (DEPRECATED)
//
// <Note>
// Price overrides are being phased out in favor adding/removing/replacing prices. (See
// [Customize your customer's subscriptions](/api-reference/subscription/create-subscription))
// </Note>
//
// Price overrides are used to update some or all prices in a plan for the specific
// subscription being created. This is useful when a new customer has negotiated a
// rate that is unique to the customer.
//
// To override prices, provide a list of objects with the key `price_overrides`.
// The price object in the list of overrides is expected to contain the existing
// price id, the `model_type` and configuration. (See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations.) The numerical values can be updated, but
// the billable metric, cadence, type, and name of a price can not be overridden.
//
// ### Maximums and Minimums
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
// ### Discounts
//
// Discounts, like price overrides, can be useful when a new customer has
// negotiated a new or different discount than the default for a price. If a
// discount exists for a price and a null discount is provided on creation, then
// there will be no discount on the new subscription.
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
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
	opts = append(r.Options[:], opts...)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `metadata`, `net terms`,
// `auto_collection`, `invoicing_threshold`, and `default_invoice_memo` properties
// on a subscription.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *shared.SubscriptionModel, err error) {
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
// [paginated](/api-reference/pagination) list, ordered starting from the most
// recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](/core-concepts##subscription).
//
// Subscriptions can be filtered for a specific customer by using either the
// customer_id or external_customer_id query parameters. To filter subscriptions
// for multiple customers, use the customer_id[] or external_customer_id[] query
// parameters.
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[shared.SubscriptionModel], err error) {
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
// [paginated](/api-reference/pagination) list, ordered starting from the most
// recently created subscription. For a full discussion of the subscription
// resource, see [Subscription](/core-concepts##subscription).
//
// Subscriptions can be filtered for a specific customer by using either the
// customer_id or external_customer_id query parameters. To filter subscriptions
// for multiple customers, use the customer_id[] or external_customer_id[] query
// parameters.
func (r *SubscriptionService) ListAutoPaging(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.SubscriptionModel] {
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
// [cancellation behaviors](/product-catalog/creating-subscriptions#cancellation-behaviors).
func (r *SubscriptionService) Cancel(ctx context.Context, subscriptionID string, body SubscriptionCancelParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/cancel", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to fetch a [Subscription](/core-concepts##subscription)
// given an identifier.
func (r *SubscriptionService) Fetch(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *shared.SubscriptionModel, err error) {
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

// This endpoint returns a [paginated](/api-reference/pagination) list of all plans
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

// This endpoint returns a [paginated](/api-reference/pagination) list of all plans
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
// [price intervals](/api-reference/price-interval/add-or-edit-price-intervals). By
// making modifications to a subscription’s price intervals, you can
// [flexibly and atomically control the billing behavior of a subscription](/product-catalog/modifying-subscriptions).
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
func (r *SubscriptionService) PriceIntervals(ctx context.Context, subscriptionID string, body SubscriptionPriceIntervalsParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/price_intervals", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to change an existing subscription's plan. It returns
// the serialized updated subscription object.
//
// The body parameter `change_option` determines when the plan change occurrs. Orb
// supports three options:
//
//   - `end_of_subscription_term`: changes the plan at the end of the existing plan's
//     term.
//   - Issuing this plan change request for a monthly subscription will keep the
//     existing plan active until the start of the subsequent month. Issuing this
//     plan change request for a yearly subscription will keep the existing plan
//     active for the full year. Charges incurred in the remaining period will be
//     invoiced as normal.
//   - Example: The plan is billed monthly on the 1st of the month, the request is
//     made on January 15th, so the plan will be changed on February 1st, and
//     invoice will be issued on February 1st for the last month of the original
//     plan.
//   - `immediate`: changes the plan immediately.
//   - Subscriptions that have their plan changed with this option will move to the
//     new plan immediately, and be invoiced immediately.
//   - This invoice will include any usage fees incurred in the billing period up
//     to the change, along with any prorated recurring fees for the billing
//     period, if applicable.
//   - Example: The plan is billed monthly on the 1st of the month, the request is
//     made on January 15th, so the plan will be changed on January 15th, and an
//     invoice will be issued for the partial month, from January 1 to January 15,
//     on the original plan.
//   - `requested_date`: changes the plan on the requested date (`change_date`).
//   - If no timezone is provided, the customer's timezone is used. The
//     `change_date` body parameter is required if this option is chosen.
//   - Example: The plan is billed monthly on the 1st of the month, the request is
//     made on January 15th, with a requested `change_date` of February 15th, so
//     the plan will be changed on February 15th, and invoices will be issued on
//     February 1st and February 15th.
//
// Note that one of `plan_id` or `external_plan_id` is required in the request body
// for this operation.
//
// ## Customize your customer's subscriptions
//
// Prices and adjustments in a plan can be added, removed, or replaced on the
// subscription when you schedule the plan change. This is useful when a customer
// has prices that differ from the default prices for a specific plan.
//
// <Note>
// This feature is only available for accounts that have migrated to Subscription Overrides Version 2. You can find your
// Subscription Overrides Version at the bottom of your [Plans page](https://app.withorb.com/plans)
// </Note>
//
// ### Adding Prices
//
// To add prices, provide a list of objects with the key `add_prices`. An object in
// the list must specify an existing add-on price with a `price_id` or
// `external_price_id` field, or create a new add-on price by including an object
// with the key `price`, identical to what would be used in the request body for
// the [create price endpoint](/api-reference/price/create-price). See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations possible in this object.
//
// If the plan has phases, each object in the list must include a number with
// `plan_phase_order` key to indicate which phase the price should be added to.
//
// An object in the list can specify an optional `start_date` and optional
// `end_date`. If `start_date` is unspecified, the start of the phase / plan change
// time will be used. If `end_date` is unspecified, it will finish at the end of
// the phase / have no end time.
//
// An object in the list can specify an optional `minimum_amount`,
// `maximum_amount`, or `discounts`. This will create adjustments which apply only
// to this price.
//
// Additionally, an object in the list can specify an optional `reference_id`. This
// ID can be used to reference this price when
// [adding an adjustment](#adding-adjustments) in the same API call. However the ID
// is _transient_ and cannot be used to refer to the price in future API calls.
//
// ### Removing Prices
//
// To remove prices, provide a list of objects with the key `remove_prices`. An
// object in the list must specify a plan price with either a `price_id` or
// `external_price_id` field.
//
// ### Replacing Prices
//
// To replace prices, provide a list of objects with the key `replace_prices`. An
// object in the list must specify a plan price to replace with the
// `replaces_price_id` key, and it must specify a price to replace it with by
// either referencing an existing add-on price with a `price_id` or
// `external_price_id` field, or by creating a new add-on price by including an
// object with the key `price`, identical to what would be used in the request body
// for the [create price endpoint](/api-reference/price/create-price). See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations possible in this object.
//
// For fixed fees, an object in the list can supply a `fixed_price_quantity`
// instead of a `price`, `price_id`, or `external_price_id` field. This will update
// only the quantity for the price, similar to the
// [Update price quantity](/api-reference/subscription/update-price-quantity)
// endpoint.
//
// The replacement price will have the same phase, if applicable, and the same
// start and end dates as the price it replaces.
//
// An object in the list can specify an optional `minimum_amount`,
// `maximum_amount`, or `discounts`. This will create adjustments which apply only
// to this price.
//
// Additionally, an object in the list can specify an optional `reference_id`. This
// ID can be used to reference the replacement price when
// [adding an adjustment](#adding-adjustments) in the same API call. However the ID
// is _transient_ and cannot be used to refer to the price in future API calls.
//
// ### Adding adjustments
//
// To add adjustments, provide a list of objects with the key `add_adjustments`. An
// object in the list must include an object with the key `adjustment`, identical
// to the adjustment object in the
// [add/edit price intervals endpoint](/api-reference/price-interval/add-or-edit-price-intervals).
//
// If the plan has phases, each object in the list must include a number with
// `plan_phase_order` key to indicate which phase the adjustment should be added
// to.
//
// An object in the list can specify an optional `start_date` and optional
// `end_date`. If `start_date` is unspecified, the start of the phase / plan change
// time will be used. If `end_date` is unspecified, it will finish at the end of
// the phase / have no end time.
//
// ### Removing adjustments
//
// To remove adjustments, provide a list of objects with the key
// `remove_adjustments`. An object in the list must include a key, `adjustment_id`,
// with the ID of the adjustment to be removed.
//
// ### Replacing adjustments
//
// To replace adjustments, provide a list of objects with the key
// `replace_adjustments`. An object in the list must specify a plan adjustment to
// replace with the `replaces_adjustment_id` key, and it must specify an adjustment
// to replace it with by including an object with the key `adjustment`, identical
// to the adjustment object in the
// [add/edit price intervals endpoint](/api-reference/price-interval/add-or-edit-price-intervals).
//
// The replacement adjustment will have the same phase, if applicable, and the same
// start and end dates as the adjustment it replaces.
//
// ## Price overrides (DEPRECATED)
//
// <Note>
// Price overrides are being phased out in favor adding/removing/replacing prices. (See
// [Customize your customer's subscriptions](/api-reference/subscription/schedule-plan-change))
// </Note>
//
// Price overrides are used to update some or all prices in a plan for the specific
// subscription being created. This is useful when a new customer has negotiated a
// rate that is unique to the customer.
//
// To override prices, provide a list of objects with the key `price_overrides`.
// The price object in the list of overrides is expected to contain the existing
// price id, the `model_type` and configuration. (See the
// [Price resource](/product-catalog/price-configuration) for the specification of
// different price model configurations.) The numerical values can be updated, but
// the billable metric, cadence, type, and name of a price can not be overridden.
//
// ### Maximums, and minimums
//
// Price overrides are used to update some or all prices in the target plan.
// Minimums and maximums, much like price overrides, can be useful when a new
// customer has negotiated a new or different minimum or maximum spend cap than the
// default for the plan. The request format for maximums and minimums is the same
// as those in [subscription creation](create-subscription).
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
// [Modifying subscriptions](/product-catalog/modifying-subscriptions#prorations-for-in-advance-fees).
func (r *SubscriptionService) SchedulePlanChange(ctx context.Context, subscriptionID string, body SubscriptionSchedulePlanChangeParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
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
func (r *SubscriptionService) TriggerPhase(ctx context.Context, subscriptionID string, body SubscriptionTriggerPhaseParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
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
func (r *SubscriptionService) UnscheduleCancellation(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
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
func (r *SubscriptionService) UnscheduleFixedFeeQuantityUpdates(ctx context.Context, subscriptionID string, body SubscriptionUnscheduleFixedFeeQuantityUpdatesParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
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
func (r *SubscriptionService) UnschedulePendingPlanChanges(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
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
func (r *SubscriptionService) UpdateFixedFeeQuantity(ctx context.Context, subscriptionID string, body SubscriptionUpdateFixedFeeQuantityParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/update_fixed_fee_quantity", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to update the trial end date for a subscription. The new
// trial end date must be within the time range of the current plan (i.e. the new
// trial end date must be on or after the subscription's start date on the current
// plan, and on or before the subscription end date).
//
// In order to retroactively remove a trial completely, the end date can be set to
// the transition date of the subscription to this plan (or, if this is the first
// plan for this subscription, the subscription's start date). In order to end a
// trial immediately, the keyword `immediate` can be provided as the trial end
// date.
//
// By default, Orb will shift only the trial end date (and price intervals that
// start or end on the previous trial end date), and leave all other future price
// intervals untouched. If the `shift` parameter is set to `true`, Orb will shift
// all subsequent price and adjustment intervals by the same amount as the trial
// end date shift (so, e.g., if a plan change is scheduled or an add-on price was
// added, that change will be pushed back by the same amount of time the trial is
// extended).
func (r *SubscriptionService) UpdateTrial(ctx context.Context, subscriptionID string, body SubscriptionUpdateTrialParams, opts ...option.RequestOption) (res *shared.MutatedSubscriptionModel, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/update_trial", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SubscriptionUsage struct {
	// This field can have the runtime type of
	// [[]SubscriptionUsageUngroupedSubscriptionUsageData],
	// [[]SubscriptionUsageGroupedSubscriptionUsageData].
	Data               interface{}               `json:"data,required"`
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
	BillableMetric shared.BillableMetricSimpleModel                        `json:"billable_metric,required"`
	Usage          []shared.UsageModel                                     `json:"usage,required"`
	ViewMode       SubscriptionUsageUngroupedSubscriptionUsageDataViewMode `json:"view_mode,required"`
	JSON           subscriptionUsageUngroupedSubscriptionUsageDataJSON     `json:"-"`
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
	BillableMetric shared.BillableMetricSimpleModel                         `json:"billable_metric,required"`
	MetricGroup    SubscriptionUsageGroupedSubscriptionUsageDataMetricGroup `json:"metric_group,required"`
	Usage          []shared.UsageModel                                      `json:"usage,required"`
	ViewMode       SubscriptionUsageGroupedSubscriptionUsageDataViewMode    `json:"view_mode,required"`
	JSON           subscriptionUsageGroupedSubscriptionUsageDataJSON        `json:"-"`
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

type SubscriptionFetchCostsResponse struct {
	Data []shared.AggregatedCostModel       `json:"data,required"`
	JSON subscriptionFetchCostsResponseJSON `json:"-"`
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

type SubscriptionFetchScheduleResponse struct {
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	EndDate   time.Time                             `json:"end_date,required,nullable" format:"date-time"`
	Plan      shared.PlanMinifiedModel              `json:"plan,required"`
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

type SubscriptionNewParams struct {
	// Additional adjustments to be added to the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	AddAdjustments param.Field[[]shared.AddSubscriptionAdjustmentParams] `json:"add_adjustments"`
	// Additional prices to be added to the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	AddPrices                             param.Field[[]shared.AddSubscriptionPriceParams] `json:"add_prices"`
	AlignBillingWithSubscriptionStartDate param.Field[bool]                                `json:"align_billing_with_subscription_start_date"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. If not specified, this
	// defaults to the behavior configured for this customer.
	AutoCollection                  param.Field[bool]                                             `json:"auto_collection"`
	AwsRegion                       param.Field[string]                                           `json:"aws_region"`
	BillingCycleAnchorConfiguration param.Field[shared.BillingCycleAnchorConfigurationModelParam] `json:"billing_cycle_anchor_configuration"`
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
	// An additional filter to apply to usage queries. This filter must be expressed as
	// a boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties). If
	// null, usage queries will not include any additional filter.
	Filter param.Field[string] `json:"filter"`
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
	// Specifies which version of the plan to subscribe to. If null, the default
	// version will be used.
	PlanVersionNumber param.Field[int64] `json:"plan_version_number"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides param.Field[[]interface{}] `json:"price_overrides"`
	// Plan adjustments to be removed from the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	RemoveAdjustments param.Field[[]shared.RemoveSubscriptionAdjustmentParams] `json:"remove_adjustments"`
	// Plan prices to be removed from the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	RemovePrices param.Field[[]shared.RemoveSubscriptionPriceParams] `json:"remove_prices"`
	// Plan adjustments to be replaced with additional adjustments on the subscription.
	// (Only available for accounts that have migrated off of legacy subscription
	// overrides)
	ReplaceAdjustments param.Field[[]shared.ReplaceSubscriptionAdjustmentParams] `json:"replace_adjustments"`
	// Plan prices to be replaced with additional prices on the subscription. (Only
	// available for accounts that have migrated off of legacy subscription overrides)
	ReplacePrices param.Field[[]shared.ReplaceSubscriptionPriceParams] `json:"replace_prices"`
	StartDate     param.Field[time.Time]                               `json:"start_date" format:"date-time"`
	// The duration of the trial period in days. If not provided, this defaults to the
	// value specified in the plan. If `0` is provided, the trial on the plan will be
	// skipped.
	TrialDurationDays param.Field[int64] `json:"trial_duration_days"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this subscription. By default, a subscription only considers usage events
	// associated with its attached customer's customer_id. When usage_customer_ids is
	// provided, the subscription includes usage events from the specified customers
	// only. Provided usage_customer_ids must be either the customer for this
	// subscription itself, or any of that customer's children.
	UsageCustomerIDs param.Field[[]string] `json:"usage_customer_ids"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
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
	// If false, this request will fail if it would void an issued invoice or create a
	// credit note. Consider using this as a safety mechanism if you do not expect
	// existing invoices to be changed.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
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
	// If false, this request will fail if it would void an issued invoice or create a
	// credit note. Consider using this as a safety mechanism if you do not expect
	// existing invoices to be changed.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
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
	AllocationPrice param.Field[shared.NewAllocationPriceModelParam] `json:"allocation_price"`
	// A list of discounts to initialize on the price interval.
	Discounts param.Field[[]SubscriptionPriceIntervalsParamsAddDiscountUnion] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription.
	EndDate param.Field[SubscriptionPriceIntervalsParamsAddEndDateUnion] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// An additional filter to apply to usage queries. This filter must be expressed as
	// a boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties). If
	// null, usage queries will not include any additional filter.
	Filter param.Field[string] `json:"filter"`
	// A list of fixed fee quantity transitions to initialize on the price interval.
	FixedFeeQuantityTransitions param.Field[[]shared.PriceIntervalFixedFeeQuantityTransitionModelParam] `json:"fixed_fee_quantity_transitions"`
	// The maximum amount that will be billed for this price interval for a given
	// billing period.
	MaximumAmount param.Field[float64] `json:"maximum_amount"`
	// The minimum amount that will be billed for this price interval for a given
	// billing period.
	MinimumAmount param.Field[float64] `json:"minimum_amount"`
	// The definition of a new price to create and add to the subscription.
	Price param.Field[shared.NewFloatingPriceModelUnionParam] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this subscription. By default, a subscription only considers usage events
	// associated with its attached customer's customer_id. When usage_customer_ids is
	// provided, the subscription includes usage events from the specified customers
	// only. Provided usage_customer_ids must be either the customer for this
	// subscription itself, or any of that customer's children.
	UsageCustomerIDs param.Field[[]string] `json:"usage_customer_ids"`
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

type SubscriptionPriceIntervalsParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[shared.NewAdjustmentModelUnionParam] `json:"adjustment,required"`
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
	// An additional filter to apply to usage queries. This filter must be expressed as
	// a boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties). If
	// null, usage queries will not include any additional filter.
	Filter param.Field[string] `json:"filter"`
	// A list of fixed fee quantity transitions to use for this price interval. Note
	// that this list will overwrite all existing fixed fee quantity transitions on the
	// price interval.
	FixedFeeQuantityTransitions param.Field[[]shared.PriceIntervalFixedFeeQuantityTransitionModelParam] `json:"fixed_fee_quantity_transitions"`
	// The updated start date of this price interval. If not specified, the start date
	// will not be updated.
	StartDate param.Field[SubscriptionPriceIntervalsParamsEditStartDateUnion] `json:"start_date" format:"date-time"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this subscription. By default, a subscription only considers usage events
	// associated with its attached customer's customer_id. When usage_customer_ids is
	// provided, the subscription includes usage events from the specified customers
	// only. Provided usage_customer_ids must be either the customer for this
	// subscription itself, or any of that customer's children.
	UsageCustomerIDs param.Field[[]string] `json:"usage_customer_ids"`
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
	// Additional adjustments to be added to the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	AddAdjustments param.Field[[]shared.AddSubscriptionAdjustmentParams] `json:"add_adjustments"`
	// Additional prices to be added to the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	AddPrices param.Field[[]shared.AddSubscriptionPriceParams] `json:"add_prices"`
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
	BillingCycleAlignment           param.Field[SubscriptionSchedulePlanChangeParamsBillingCycleAlignment] `json:"billing_cycle_alignment"`
	BillingCycleAnchorConfiguration param.Field[shared.BillingCycleAnchorConfigurationModelParam]          `json:"billing_cycle_anchor_configuration"`
	// The date that the plan change should take effect. This parameter can only be
	// passed if the `change_option` is `requested_date`. If a date with no time is
	// passed, the plan change will happen at midnight in the customer's timezone.
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
	// An additional filter to apply to usage queries. This filter must be expressed as
	// a boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties). If
	// null, usage queries will not include any additional filter.
	Filter param.Field[string] `json:"filter"`
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
	// Specifies which version of the plan to change to. If null, the default version
	// will be used.
	PlanVersionNumber param.Field[int64] `json:"plan_version_number"`
	// Optionally provide a list of overrides for prices on the plan
	PriceOverrides param.Field[[]interface{}] `json:"price_overrides"`
	// Plan adjustments to be removed from the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	RemoveAdjustments param.Field[[]shared.RemoveSubscriptionAdjustmentParams] `json:"remove_adjustments"`
	// Plan prices to be removed from the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	RemovePrices param.Field[[]shared.RemoveSubscriptionPriceParams] `json:"remove_prices"`
	// Plan adjustments to be replaced with additional adjustments on the subscription.
	// (Only available for accounts that have migrated off of legacy subscription
	// overrides)
	ReplaceAdjustments param.Field[[]shared.ReplaceSubscriptionAdjustmentParams] `json:"replace_adjustments"`
	// Plan prices to be replaced with additional prices on the subscription. (Only
	// available for accounts that have migrated off of legacy subscription overrides)
	ReplacePrices param.Field[[]shared.ReplaceSubscriptionPriceParams] `json:"replace_prices"`
	// The duration of the trial period in days. If not provided, this defaults to the
	// value specified in the plan. If `0` is provided, the trial on the plan will be
	// skipped.
	TrialDurationDays param.Field[int64] `json:"trial_duration_days"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this subscription. By default, a subscription only considers usage events
	// associated with its attached customer's customer_id. When usage_customer_ids is
	// provided, the subscription includes usage events from the specified customers
	// only. Provided usage_customer_ids must be either the customer for this
	// subscription itself, or any of that customer's children.
	UsageCustomerIDs param.Field[[]string] `json:"usage_customer_ids"`
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

type SubscriptionTriggerPhaseParams struct {
	// If false, this request will fail if it would void an issued invoice or create a
	// credit note. Consider using this as a safety mechanism if you do not expect
	// existing invoices to be changed.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
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
	// If false, this request will fail if it would void an issued invoice or create a
	// credit note. Consider using this as a safety mechanism if you do not expect
	// existing invoices to be changed.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
	// Determines when the change takes effect. Note that if `effective_date` is
	// specified, this defaults to `effective_date`. Otherwise, this defaults to
	// `immediate` unless it's explicitly set to `upcoming_invoice`.
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
// `immediate` unless it's explicitly set to `upcoming_invoice`.
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

type SubscriptionUpdateTrialParams struct {
	// The new date that the trial should end, or the literal string `immediate` to end
	// the trial immediately.
	TrialEndDate param.Field[SubscriptionUpdateTrialParamsTrialEndDateUnion] `json:"trial_end_date,required" format:"date-time"`
	// If true, shifts subsequent price and adjustment intervals (preserving their
	// durations, but adjusting their absolute dates).
	Shift param.Field[bool] `json:"shift"`
}

func (r SubscriptionUpdateTrialParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new date that the trial should end, or the literal string `immediate` to end
// the trial immediately.
//
// Satisfied by [shared.UnionTime],
// [SubscriptionUpdateTrialParamsTrialEndDateString].
type SubscriptionUpdateTrialParamsTrialEndDateUnion interface {
	ImplementsSubscriptionUpdateTrialParamsTrialEndDateUnion()
}

type SubscriptionUpdateTrialParamsTrialEndDateString string

const (
	SubscriptionUpdateTrialParamsTrialEndDateStringImmediate SubscriptionUpdateTrialParamsTrialEndDateString = "immediate"
)

func (r SubscriptionUpdateTrialParamsTrialEndDateString) IsKnown() bool {
	switch r {
	case SubscriptionUpdateTrialParamsTrialEndDateStringImmediate:
		return true
	}
	return false
}

func (r SubscriptionUpdateTrialParamsTrialEndDateString) ImplementsSubscriptionUpdateTrialParamsTrialEndDateUnion() {
}
