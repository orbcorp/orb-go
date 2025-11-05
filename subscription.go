// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
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
// \$10.00 for a subscription that invoices in USD.
func (r *SubscriptionService) New(ctx context.Context, body SubscriptionNewParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to update the `metadata`, `net terms`,
// `auto_collection`, `invoicing_threshold`, and `default_invoice_memo` properties
// on a subscription.
func (r *SubscriptionService) Update(ctx context.Context, subscriptionID string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) List(ctx context.Context, query SubscriptionListParams, opts ...option.RequestOption) (res *pagination.Page[Subscription], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
// [cancellation behaviors](/product-catalog/creating-subscriptions#cancellation-behaviors).
func (r *SubscriptionService) Cancel(ctx context.Context, subscriptionID string, body SubscriptionCancelParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) Fetch(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *Subscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) PriceIntervals(ctx context.Context, subscriptionID string, body SubscriptionPriceIntervalsParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/price_intervals", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Redeem a coupon effective at a given time.
func (r *SubscriptionService) RedeemCoupon(ctx context.Context, subscriptionID string, body SubscriptionRedeemCouponParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/redeem_coupon", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to change an existing subscription's plan. It returns
// the serialized updated subscription object.
//
// The body parameter `change_option` determines when the plan change occurs. Orb
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
func (r *SubscriptionService) SchedulePlanChange(ctx context.Context, subscriptionID string, body SubscriptionSchedulePlanChangeParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) TriggerPhase(ctx context.Context, subscriptionID string, body SubscriptionTriggerPhaseParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) UnscheduleCancellation(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) UnscheduleFixedFeeQuantityUpdates(ctx context.Context, subscriptionID string, body SubscriptionUnscheduleFixedFeeQuantityUpdatesParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/unschedule_fixed_fee_quantity_updates", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to unschedule any pending plan changes on an existing
// subscription. When called, all upcoming plan changes will be unscheduled.
func (r *SubscriptionService) UnschedulePendingPlanChanges(ctx context.Context, subscriptionID string, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) UpdateFixedFeeQuantity(ctx context.Context, subscriptionID string, body SubscriptionUpdateFixedFeeQuantityParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *SubscriptionService) UpdateTrial(ctx context.Context, subscriptionID string, body SubscriptionUpdateTrialParams, opts ...option.RequestOption) (res *MutatedSubscription, err error) {
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("subscriptions/%s/update_trial", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DiscountOverrideParam struct {
	DiscountType param.Field[DiscountOverrideDiscountType] `json:"discount_type,required"`
	// Only available if discount_type is `amount`.
	AmountDiscount param.Field[string] `json:"amount_discount"`
	// Only available if discount_type is `percentage`. This is a number between 0
	// and 1.
	PercentageDiscount param.Field[float64] `json:"percentage_discount"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount param.Field[float64] `json:"usage_discount"`
}

func (r DiscountOverrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DiscountOverrideDiscountType string

const (
	DiscountOverrideDiscountTypePercentage DiscountOverrideDiscountType = "percentage"
	DiscountOverrideDiscountTypeUsage      DiscountOverrideDiscountType = "usage"
	DiscountOverrideDiscountTypeAmount     DiscountOverrideDiscountType = "amount"
)

func (r DiscountOverrideDiscountType) IsKnown() bool {
	switch r {
	case DiscountOverrideDiscountTypePercentage, DiscountOverrideDiscountTypeUsage, DiscountOverrideDiscountTypeAmount:
		return true
	}
	return false
}

type NewSubscriptionBulkPriceParam struct {
	// Configuration for bulk pricing
	BulkConfig param.Field[shared.BulkConfigParam] `json:"bulk_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionBulkPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionBulkPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionBulkPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionBulkPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

func (r NewSubscriptionBulkPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {}

func (r NewSubscriptionBulkPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionBulkPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionBulkPriceCadence string

const (
	NewSubscriptionBulkPriceCadenceAnnual     NewSubscriptionBulkPriceCadence = "annual"
	NewSubscriptionBulkPriceCadenceSemiAnnual NewSubscriptionBulkPriceCadence = "semi_annual"
	NewSubscriptionBulkPriceCadenceMonthly    NewSubscriptionBulkPriceCadence = "monthly"
	NewSubscriptionBulkPriceCadenceQuarterly  NewSubscriptionBulkPriceCadence = "quarterly"
	NewSubscriptionBulkPriceCadenceOneTime    NewSubscriptionBulkPriceCadence = "one_time"
	NewSubscriptionBulkPriceCadenceCustom     NewSubscriptionBulkPriceCadence = "custom"
)

func (r NewSubscriptionBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkPriceCadenceAnnual, NewSubscriptionBulkPriceCadenceSemiAnnual, NewSubscriptionBulkPriceCadenceMonthly, NewSubscriptionBulkPriceCadenceQuarterly, NewSubscriptionBulkPriceCadenceOneTime, NewSubscriptionBulkPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionBulkPriceModelType string

const (
	NewSubscriptionBulkPriceModelTypeBulk NewSubscriptionBulkPriceModelType = "bulk"
)

func (r NewSubscriptionBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkPriceModelTypeBulk:
		return true
	}
	return false
}

type NewSubscriptionBulkPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewSubscriptionBulkPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionBulkPriceConversionRateConfigParam) ImplementsNewSubscriptionBulkPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionBulkPriceConversionRateConfigParam].
type NewSubscriptionBulkPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionBulkPriceConversionRateConfigUnionParam()
}

type NewSubscriptionBulkPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionBulkPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionBulkPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionBulkPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionBulkWithProrationPriceParam struct {
	// Configuration for bulk_with_proration pricing
	BulkWithProrationConfig param.Field[NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigParam] `json:"bulk_with_proration_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionBulkWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionBulkWithProrationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionBulkWithProrationPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionBulkWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionBulkWithProrationPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionBulkWithProrationPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionBulkWithProrationPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionBulkWithProrationPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_proration pricing
type NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigParam struct {
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier with proration
type NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigTierParam struct {
	// Cost per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r NewSubscriptionBulkWithProrationPriceBulkWithProrationConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type NewSubscriptionBulkWithProrationPriceCadence string

const (
	NewSubscriptionBulkWithProrationPriceCadenceAnnual     NewSubscriptionBulkWithProrationPriceCadence = "annual"
	NewSubscriptionBulkWithProrationPriceCadenceSemiAnnual NewSubscriptionBulkWithProrationPriceCadence = "semi_annual"
	NewSubscriptionBulkWithProrationPriceCadenceMonthly    NewSubscriptionBulkWithProrationPriceCadence = "monthly"
	NewSubscriptionBulkWithProrationPriceCadenceQuarterly  NewSubscriptionBulkWithProrationPriceCadence = "quarterly"
	NewSubscriptionBulkWithProrationPriceCadenceOneTime    NewSubscriptionBulkWithProrationPriceCadence = "one_time"
	NewSubscriptionBulkWithProrationPriceCadenceCustom     NewSubscriptionBulkWithProrationPriceCadence = "custom"
)

func (r NewSubscriptionBulkWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkWithProrationPriceCadenceAnnual, NewSubscriptionBulkWithProrationPriceCadenceSemiAnnual, NewSubscriptionBulkWithProrationPriceCadenceMonthly, NewSubscriptionBulkWithProrationPriceCadenceQuarterly, NewSubscriptionBulkWithProrationPriceCadenceOneTime, NewSubscriptionBulkWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionBulkWithProrationPriceModelType string

const (
	NewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration NewSubscriptionBulkWithProrationPriceModelType = "bulk_with_proration"
)

func (r NewSubscriptionBulkWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkWithProrationPriceModelTypeBulkWithProration:
		return true
	}
	return false
}

type NewSubscriptionBulkWithProrationPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                        `json:"unit_config"`
}

func (r NewSubscriptionBulkWithProrationPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionBulkWithProrationPriceConversionRateConfigParam) ImplementsNewSubscriptionBulkWithProrationPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionBulkWithProrationPriceConversionRateConfigParam].
type NewSubscriptionBulkWithProrationPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionBulkWithProrationPriceConversionRateConfigUnionParam()
}

type NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionBulkWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionCumulativeGroupedBulkPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionCumulativeGroupedBulkPriceCadence] `json:"cadence,required"`
	// Configuration for cumulative_grouped_bulk pricing
	CumulativeGroupedBulkConfig param.Field[NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam] `json:"cumulative_grouped_bulk_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionCumulativeGroupedBulkPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionCumulativeGroupedBulkPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionCumulativeGroupedBulkPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionCumulativeGroupedBulkPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionCumulativeGroupedBulkPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionCumulativeGroupedBulkPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionCumulativeGroupedBulkPriceCadence string

const (
	NewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual     NewSubscriptionCumulativeGroupedBulkPriceCadence = "annual"
	NewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual NewSubscriptionCumulativeGroupedBulkPriceCadence = "semi_annual"
	NewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly    NewSubscriptionCumulativeGroupedBulkPriceCadence = "monthly"
	NewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly  NewSubscriptionCumulativeGroupedBulkPriceCadence = "quarterly"
	NewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime    NewSubscriptionCumulativeGroupedBulkPriceCadence = "one_time"
	NewSubscriptionCumulativeGroupedBulkPriceCadenceCustom     NewSubscriptionCumulativeGroupedBulkPriceCadence = "custom"
)

func (r NewSubscriptionCumulativeGroupedBulkPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionCumulativeGroupedBulkPriceCadenceAnnual, NewSubscriptionCumulativeGroupedBulkPriceCadenceSemiAnnual, NewSubscriptionCumulativeGroupedBulkPriceCadenceMonthly, NewSubscriptionCumulativeGroupedBulkPriceCadenceQuarterly, NewSubscriptionCumulativeGroupedBulkPriceCadenceOneTime, NewSubscriptionCumulativeGroupedBulkPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for cumulative_grouped_bulk pricing
type NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam struct {
	// Each tier lower bound must have the same group of values.
	DimensionValues param.Field[[]NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam] `json:"dimension_values,required"`
	// Grouping key name
	Group param.Field[string] `json:"group,required"`
}

func (r NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a dimension value entry
type NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam struct {
	// Grouping key value
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Unit amount for this combination
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionCumulativeGroupedBulkPriceCumulativeGroupedBulkConfigDimensionValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionCumulativeGroupedBulkPriceModelType string

const (
	NewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk NewSubscriptionCumulativeGroupedBulkPriceModelType = "cumulative_grouped_bulk"
)

func (r NewSubscriptionCumulativeGroupedBulkPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionCumulativeGroupedBulkPriceModelTypeCumulativeGroupedBulk:
		return true
	}
	return false
}

type NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                            `json:"unit_config"`
}

func (r NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigParam) ImplementsNewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigParam].
type NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigUnionParam()
}

type NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionCumulativeGroupedBulkPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionGroupedAllocationPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_allocation pricing
	GroupedAllocationConfig param.Field[NewSubscriptionGroupedAllocationPriceGroupedAllocationConfigParam] `json:"grouped_allocation_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionGroupedAllocationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionGroupedAllocationPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionGroupedAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedAllocationPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedAllocationPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionGroupedAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionGroupedAllocationPriceCadence string

const (
	NewSubscriptionGroupedAllocationPriceCadenceAnnual     NewSubscriptionGroupedAllocationPriceCadence = "annual"
	NewSubscriptionGroupedAllocationPriceCadenceSemiAnnual NewSubscriptionGroupedAllocationPriceCadence = "semi_annual"
	NewSubscriptionGroupedAllocationPriceCadenceMonthly    NewSubscriptionGroupedAllocationPriceCadence = "monthly"
	NewSubscriptionGroupedAllocationPriceCadenceQuarterly  NewSubscriptionGroupedAllocationPriceCadence = "quarterly"
	NewSubscriptionGroupedAllocationPriceCadenceOneTime    NewSubscriptionGroupedAllocationPriceCadence = "one_time"
	NewSubscriptionGroupedAllocationPriceCadenceCustom     NewSubscriptionGroupedAllocationPriceCadence = "custom"
)

func (r NewSubscriptionGroupedAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedAllocationPriceCadenceAnnual, NewSubscriptionGroupedAllocationPriceCadenceSemiAnnual, NewSubscriptionGroupedAllocationPriceCadenceMonthly, NewSubscriptionGroupedAllocationPriceCadenceQuarterly, NewSubscriptionGroupedAllocationPriceCadenceOneTime, NewSubscriptionGroupedAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_allocation pricing
type NewSubscriptionGroupedAllocationPriceGroupedAllocationConfigParam struct {
	// Usage allocation per group
	Allocation param.Field[string] `json:"allocation,required"`
	// How to determine the groups that should each be allocated some quantity
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// Unit rate for post-allocation
	OverageUnitRate param.Field[string] `json:"overage_unit_rate,required"`
}

func (r NewSubscriptionGroupedAllocationPriceGroupedAllocationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionGroupedAllocationPriceModelType string

const (
	NewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation NewSubscriptionGroupedAllocationPriceModelType = "grouped_allocation"
)

func (r NewSubscriptionGroupedAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedAllocationPriceModelTypeGroupedAllocation:
		return true
	}
	return false
}

type NewSubscriptionGroupedAllocationPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                        `json:"unit_config"`
}

func (r NewSubscriptionGroupedAllocationPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedAllocationPriceConversionRateConfigParam) ImplementsNewSubscriptionGroupedAllocationPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionGroupedAllocationPriceConversionRateConfigParam].
type NewSubscriptionGroupedAllocationPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionGroupedAllocationPriceConversionRateConfigUnionParam()
}

type NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionGroupedAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionGroupedTieredPackagePriceCadence] `json:"cadence,required"`
	// Configuration for grouped_tiered_package pricing
	GroupedTieredPackageConfig param.Field[NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigParam] `json:"grouped_tiered_package_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionGroupedTieredPackagePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionGroupedTieredPackagePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionGroupedTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedTieredPackagePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPackagePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionGroupedTieredPackagePriceCadence string

const (
	NewSubscriptionGroupedTieredPackagePriceCadenceAnnual     NewSubscriptionGroupedTieredPackagePriceCadence = "annual"
	NewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual NewSubscriptionGroupedTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionGroupedTieredPackagePriceCadenceMonthly    NewSubscriptionGroupedTieredPackagePriceCadence = "monthly"
	NewSubscriptionGroupedTieredPackagePriceCadenceQuarterly  NewSubscriptionGroupedTieredPackagePriceCadence = "quarterly"
	NewSubscriptionGroupedTieredPackagePriceCadenceOneTime    NewSubscriptionGroupedTieredPackagePriceCadence = "one_time"
	NewSubscriptionGroupedTieredPackagePriceCadenceCustom     NewSubscriptionGroupedTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionGroupedTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPackagePriceCadenceAnnual, NewSubscriptionGroupedTieredPackagePriceCadenceSemiAnnual, NewSubscriptionGroupedTieredPackagePriceCadenceMonthly, NewSubscriptionGroupedTieredPackagePriceCadenceQuarterly, NewSubscriptionGroupedTieredPackagePriceCadenceOneTime, NewSubscriptionGroupedTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered_package pricing
type NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigParam struct {
	// The event property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// Package size
	PackageSize param.Field[string] `json:"package_size,required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam struct {
	// Price per package
	PerUnit param.Field[string] `json:"per_unit,required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
}

func (r NewSubscriptionGroupedTieredPackagePriceGroupedTieredPackageConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionGroupedTieredPackagePriceModelType string

const (
	NewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage NewSubscriptionGroupedTieredPackagePriceModelType = "grouped_tiered_package"
)

func (r NewSubscriptionGroupedTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPackagePriceModelTypeGroupedTieredPackage:
		return true
	}
	return false
}

type NewSubscriptionGroupedTieredPackagePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                           `json:"unit_config"`
}

func (r NewSubscriptionGroupedTieredPackagePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedTieredPackagePriceConversionRateConfigParam) ImplementsNewSubscriptionGroupedTieredPackagePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionGroupedTieredPackagePriceConversionRateConfigParam].
type NewSubscriptionGroupedTieredPackagePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionGroupedTieredPackagePriceConversionRateConfigUnionParam()
}

type NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionGroupedTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionGroupedTieredPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_tiered pricing
	GroupedTieredConfig param.Field[NewSubscriptionGroupedTieredPriceGroupedTieredConfigParam] `json:"grouped_tiered_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionGroupedTieredPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionGroupedTieredPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionGroupedTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedTieredPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedTieredPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionGroupedTieredPriceCadence string

const (
	NewSubscriptionGroupedTieredPriceCadenceAnnual     NewSubscriptionGroupedTieredPriceCadence = "annual"
	NewSubscriptionGroupedTieredPriceCadenceSemiAnnual NewSubscriptionGroupedTieredPriceCadence = "semi_annual"
	NewSubscriptionGroupedTieredPriceCadenceMonthly    NewSubscriptionGroupedTieredPriceCadence = "monthly"
	NewSubscriptionGroupedTieredPriceCadenceQuarterly  NewSubscriptionGroupedTieredPriceCadence = "quarterly"
	NewSubscriptionGroupedTieredPriceCadenceOneTime    NewSubscriptionGroupedTieredPriceCadence = "one_time"
	NewSubscriptionGroupedTieredPriceCadenceCustom     NewSubscriptionGroupedTieredPriceCadence = "custom"
)

func (r NewSubscriptionGroupedTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPriceCadenceAnnual, NewSubscriptionGroupedTieredPriceCadenceSemiAnnual, NewSubscriptionGroupedTieredPriceCadenceMonthly, NewSubscriptionGroupedTieredPriceCadenceQuarterly, NewSubscriptionGroupedTieredPriceCadenceOneTime, NewSubscriptionGroupedTieredPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_tiered pricing
type NewSubscriptionGroupedTieredPriceGroupedTieredConfigParam struct {
	// The billable metric property used to group before tiering
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// Apply tiered pricing to each segment generated after grouping with the provided
	// key
	Tiers param.Field[[]NewSubscriptionGroupedTieredPriceGroupedTieredConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionGroupedTieredPriceGroupedTieredConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type NewSubscriptionGroupedTieredPriceGroupedTieredConfigTierParam struct {
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionGroupedTieredPriceGroupedTieredConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionGroupedTieredPriceModelType string

const (
	NewSubscriptionGroupedTieredPriceModelTypeGroupedTiered NewSubscriptionGroupedTieredPriceModelType = "grouped_tiered"
)

func (r NewSubscriptionGroupedTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPriceModelTypeGroupedTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedTieredPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                    `json:"unit_config"`
}

func (r NewSubscriptionGroupedTieredPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedTieredPriceConversionRateConfigParam) ImplementsNewSubscriptionGroupedTieredPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionGroupedTieredPriceConversionRateConfigParam].
type NewSubscriptionGroupedTieredPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionGroupedTieredPriceConversionRateConfigUnionParam()
}

type NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionGroupedTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedWithMeteredMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionGroupedWithMeteredMinimumPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_metered_minimum pricing
	GroupedWithMeteredMinimumConfig param.Field[NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam] `json:"grouped_with_metered_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionGroupedWithMeteredMinimumPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionGroupedWithMeteredMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionGroupedWithMeteredMinimumPriceCadence string

const (
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual     NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "annual"
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "semi_annual"
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly    NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "monthly"
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly  NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "quarterly"
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime    NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "one_time"
	NewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom     NewSubscriptionGroupedWithMeteredMinimumPriceCadence = "custom"
)

func (r NewSubscriptionGroupedWithMeteredMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithMeteredMinimumPriceCadenceAnnual, NewSubscriptionGroupedWithMeteredMinimumPriceCadenceSemiAnnual, NewSubscriptionGroupedWithMeteredMinimumPriceCadenceMonthly, NewSubscriptionGroupedWithMeteredMinimumPriceCadenceQuarterly, NewSubscriptionGroupedWithMeteredMinimumPriceCadenceOneTime, NewSubscriptionGroupedWithMeteredMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_metered_minimum pricing
type NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam struct {
	// Used to partition the usage into groups. The minimum amount is applied to each
	// group.
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The minimum amount to charge per group per unit
	MinimumUnitAmount param.Field[string] `json:"minimum_unit_amount,required"`
	// Used to determine the unit rate
	PricingKey param.Field[string] `json:"pricing_key,required"`
	// Scale the unit rates by the scaling factor.
	ScalingFactors param.Field[[]NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam] `json:"scaling_factors,required"`
	// Used to determine the unit rate scaling factor
	ScalingKey param.Field[string] `json:"scaling_key,required"`
	// Apply per unit pricing to each pricing value. The minimum amount is applied any
	// unmatched usage.
	UnitAmounts param.Field[[]NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam] `json:"unit_amounts,required"`
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a scaling factor
type NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam struct {
	// Scaling factor
	ScalingFactor param.Field[string] `json:"scaling_factor,required"`
	// Scaling value
	ScalingValue param.Field[string] `json:"scaling_value,required"`
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigScalingFactorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount
type NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam struct {
	// Pricing value
	PricingValue param.Field[string] `json:"pricing_value,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceGroupedWithMeteredMinimumConfigUnitAmountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionGroupedWithMeteredMinimumPriceModelType string

const (
	NewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum NewSubscriptionGroupedWithMeteredMinimumPriceModelType = "grouped_with_metered_minimum"
)

func (r NewSubscriptionGroupedWithMeteredMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithMeteredMinimumPriceModelTypeGroupedWithMeteredMinimum:
		return true
	}
	return false
}

type NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                `json:"unit_config"`
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigParam) ImplementsNewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigParam].
type NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigUnionParam()
}

type NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionGroupedWithMeteredMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionGroupedWithProratedMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionGroupedWithProratedMinimumPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_prorated_minimum pricing
	GroupedWithProratedMinimumConfig param.Field[NewSubscriptionGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam] `json:"grouped_with_prorated_minimum_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionGroupedWithProratedMinimumPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionGroupedWithProratedMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionGroupedWithProratedMinimumPriceCadence string

const (
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual     NewSubscriptionGroupedWithProratedMinimumPriceCadence = "annual"
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual NewSubscriptionGroupedWithProratedMinimumPriceCadence = "semi_annual"
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly    NewSubscriptionGroupedWithProratedMinimumPriceCadence = "monthly"
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly  NewSubscriptionGroupedWithProratedMinimumPriceCadence = "quarterly"
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime    NewSubscriptionGroupedWithProratedMinimumPriceCadence = "one_time"
	NewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom     NewSubscriptionGroupedWithProratedMinimumPriceCadence = "custom"
)

func (r NewSubscriptionGroupedWithProratedMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithProratedMinimumPriceCadenceAnnual, NewSubscriptionGroupedWithProratedMinimumPriceCadenceSemiAnnual, NewSubscriptionGroupedWithProratedMinimumPriceCadenceMonthly, NewSubscriptionGroupedWithProratedMinimumPriceCadenceQuarterly, NewSubscriptionGroupedWithProratedMinimumPriceCadenceOneTime, NewSubscriptionGroupedWithProratedMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_prorated_minimum pricing
type NewSubscriptionGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam struct {
	// How to determine the groups that should each have a minimum
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The minimum amount to charge per group
	Minimum param.Field[string] `json:"minimum,required"`
	// The amount to charge per unit
	UnitRate param.Field[string] `json:"unit_rate,required"`
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceGroupedWithProratedMinimumConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionGroupedWithProratedMinimumPriceModelType string

const (
	NewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum NewSubscriptionGroupedWithProratedMinimumPriceModelType = "grouped_with_prorated_minimum"
)

func (r NewSubscriptionGroupedWithProratedMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithProratedMinimumPriceModelTypeGroupedWithProratedMinimum:
		return true
	}
	return false
}

type NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                 `json:"unit_config"`
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigParam) ImplementsNewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigParam].
type NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigUnionParam()
}

type NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionGroupedWithProratedMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionMatrixPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionMatrixPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// Configuration for matrix pricing
	MatrixConfig param.Field[shared.MatrixConfigParam] `json:"matrix_config,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionMatrixPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionMatrixPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionMatrixPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

func (r NewSubscriptionMatrixPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {}

func (r NewSubscriptionMatrixPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMatrixPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionMatrixPriceCadence string

const (
	NewSubscriptionMatrixPriceCadenceAnnual     NewSubscriptionMatrixPriceCadence = "annual"
	NewSubscriptionMatrixPriceCadenceSemiAnnual NewSubscriptionMatrixPriceCadence = "semi_annual"
	NewSubscriptionMatrixPriceCadenceMonthly    NewSubscriptionMatrixPriceCadence = "monthly"
	NewSubscriptionMatrixPriceCadenceQuarterly  NewSubscriptionMatrixPriceCadence = "quarterly"
	NewSubscriptionMatrixPriceCadenceOneTime    NewSubscriptionMatrixPriceCadence = "one_time"
	NewSubscriptionMatrixPriceCadenceCustom     NewSubscriptionMatrixPriceCadence = "custom"
)

func (r NewSubscriptionMatrixPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixPriceCadenceAnnual, NewSubscriptionMatrixPriceCadenceSemiAnnual, NewSubscriptionMatrixPriceCadenceMonthly, NewSubscriptionMatrixPriceCadenceQuarterly, NewSubscriptionMatrixPriceCadenceOneTime, NewSubscriptionMatrixPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionMatrixPriceModelType string

const (
	NewSubscriptionMatrixPriceModelTypeMatrix NewSubscriptionMatrixPriceModelType = "matrix"
)

func (r NewSubscriptionMatrixPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixPriceModelTypeMatrix:
		return true
	}
	return false
}

type NewSubscriptionMatrixPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionMatrixPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                             `json:"unit_config"`
}

func (r NewSubscriptionMatrixPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixPriceConversionRateConfigParam) ImplementsNewSubscriptionMatrixPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionMatrixPriceConversionRateConfigParam].
type NewSubscriptionMatrixPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionMatrixPriceConversionRateConfigUnionParam()
}

type NewSubscriptionMatrixPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionMatrixPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionMatrixPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionMatrixPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionMatrixPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionMatrixPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionMatrixPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionMatrixWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionMatrixWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// Configuration for matrix_with_allocation pricing
	MatrixWithAllocationConfig param.Field[shared.MatrixWithAllocationConfigParam] `json:"matrix_with_allocation_config,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionMatrixWithAllocationPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionMatrixWithAllocationPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionMatrixWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixWithAllocationPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithAllocationPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionMatrixWithAllocationPriceCadence string

const (
	NewSubscriptionMatrixWithAllocationPriceCadenceAnnual     NewSubscriptionMatrixWithAllocationPriceCadence = "annual"
	NewSubscriptionMatrixWithAllocationPriceCadenceSemiAnnual NewSubscriptionMatrixWithAllocationPriceCadence = "semi_annual"
	NewSubscriptionMatrixWithAllocationPriceCadenceMonthly    NewSubscriptionMatrixWithAllocationPriceCadence = "monthly"
	NewSubscriptionMatrixWithAllocationPriceCadenceQuarterly  NewSubscriptionMatrixWithAllocationPriceCadence = "quarterly"
	NewSubscriptionMatrixWithAllocationPriceCadenceOneTime    NewSubscriptionMatrixWithAllocationPriceCadence = "one_time"
	NewSubscriptionMatrixWithAllocationPriceCadenceCustom     NewSubscriptionMatrixWithAllocationPriceCadence = "custom"
)

func (r NewSubscriptionMatrixWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithAllocationPriceCadenceAnnual, NewSubscriptionMatrixWithAllocationPriceCadenceSemiAnnual, NewSubscriptionMatrixWithAllocationPriceCadenceMonthly, NewSubscriptionMatrixWithAllocationPriceCadenceQuarterly, NewSubscriptionMatrixWithAllocationPriceCadenceOneTime, NewSubscriptionMatrixWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionMatrixWithAllocationPriceModelType string

const (
	NewSubscriptionMatrixWithAllocationPriceModelTypeMatrixWithAllocation NewSubscriptionMatrixWithAllocationPriceModelType = "matrix_with_allocation"
)

func (r NewSubscriptionMatrixWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithAllocationPriceModelTypeMatrixWithAllocation:
		return true
	}
	return false
}

type NewSubscriptionMatrixWithAllocationPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                           `json:"unit_config"`
}

func (r NewSubscriptionMatrixWithAllocationPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixWithAllocationPriceConversionRateConfigParam) ImplementsNewSubscriptionMatrixWithAllocationPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionMatrixWithAllocationPriceConversionRateConfigParam].
type NewSubscriptionMatrixWithAllocationPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionMatrixWithAllocationPriceConversionRateConfigUnionParam()
}

type NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionMatrixWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionMatrixWithDisplayNamePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionMatrixWithDisplayNamePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// Configuration for matrix_with_display_name pricing
	MatrixWithDisplayNameConfig param.Field[NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam] `json:"matrix_with_display_name_config,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionMatrixWithDisplayNamePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionMatrixWithDisplayNamePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixWithDisplayNamePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithDisplayNamePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithDisplayNamePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMatrixWithDisplayNamePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionMatrixWithDisplayNamePriceCadence string

const (
	NewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual     NewSubscriptionMatrixWithDisplayNamePriceCadence = "annual"
	NewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual NewSubscriptionMatrixWithDisplayNamePriceCadence = "semi_annual"
	NewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly    NewSubscriptionMatrixWithDisplayNamePriceCadence = "monthly"
	NewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly  NewSubscriptionMatrixWithDisplayNamePriceCadence = "quarterly"
	NewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime    NewSubscriptionMatrixWithDisplayNamePriceCadence = "one_time"
	NewSubscriptionMatrixWithDisplayNamePriceCadenceCustom     NewSubscriptionMatrixWithDisplayNamePriceCadence = "custom"
)

func (r NewSubscriptionMatrixWithDisplayNamePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithDisplayNamePriceCadenceAnnual, NewSubscriptionMatrixWithDisplayNamePriceCadenceSemiAnnual, NewSubscriptionMatrixWithDisplayNamePriceCadenceMonthly, NewSubscriptionMatrixWithDisplayNamePriceCadenceQuarterly, NewSubscriptionMatrixWithDisplayNamePriceCadenceOneTime, NewSubscriptionMatrixWithDisplayNamePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for matrix_with_display_name pricing
type NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam struct {
	// Used to determine the unit rate
	Dimension param.Field[string] `json:"dimension,required"`
	// Apply per unit pricing to each dimension value
	UnitAmounts param.Field[[]NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam] `json:"unit_amounts,required"`
}

func (r NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a unit amount item
type NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam struct {
	// The dimension value
	DimensionValue param.Field[string] `json:"dimension_value,required"`
	// Display name for this dimension value
	DisplayName param.Field[string] `json:"display_name,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionMatrixWithDisplayNamePriceMatrixWithDisplayNameConfigUnitAmountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionMatrixWithDisplayNamePriceModelType string

const (
	NewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName NewSubscriptionMatrixWithDisplayNamePriceModelType = "matrix_with_display_name"
)

func (r NewSubscriptionMatrixWithDisplayNamePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithDisplayNamePriceModelTypeMatrixWithDisplayName:
		return true
	}
	return false
}

type NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                            `json:"unit_config"`
}

func (r NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigParam) ImplementsNewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigParam].
type NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigUnionParam()
}

type NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionMatrixWithDisplayNamePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionMaxGroupTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionMaxGroupTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// Configuration for max_group_tiered_package pricing
	MaxGroupTieredPackageConfig param.Field[NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam] `json:"max_group_tiered_package_config,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionMaxGroupTieredPackagePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionMaxGroupTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMaxGroupTieredPackagePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMaxGroupTieredPackagePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionMaxGroupTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMaxGroupTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionMaxGroupTieredPackagePriceCadence string

const (
	NewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual     NewSubscriptionMaxGroupTieredPackagePriceCadence = "annual"
	NewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual NewSubscriptionMaxGroupTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly    NewSubscriptionMaxGroupTieredPackagePriceCadence = "monthly"
	NewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly  NewSubscriptionMaxGroupTieredPackagePriceCadence = "quarterly"
	NewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime    NewSubscriptionMaxGroupTieredPackagePriceCadence = "one_time"
	NewSubscriptionMaxGroupTieredPackagePriceCadenceCustom     NewSubscriptionMaxGroupTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionMaxGroupTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionMaxGroupTieredPackagePriceCadenceAnnual, NewSubscriptionMaxGroupTieredPackagePriceCadenceSemiAnnual, NewSubscriptionMaxGroupTieredPackagePriceCadenceMonthly, NewSubscriptionMaxGroupTieredPackagePriceCadenceQuarterly, NewSubscriptionMaxGroupTieredPackagePriceCadenceOneTime, NewSubscriptionMaxGroupTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for max_group_tiered_package pricing
type NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam struct {
	// The event property used to group before tiering the group with the highest value
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// Package size
	PackageSize param.Field[string] `json:"package_size,required"`
	// Apply tiered pricing to the largest group after grouping with the provided key.
	Tiers param.Field[[]NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam struct {
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionMaxGroupTieredPackagePriceMaxGroupTieredPackageConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionMaxGroupTieredPackagePriceModelType string

const (
	NewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage NewSubscriptionMaxGroupTieredPackagePriceModelType = "max_group_tiered_package"
)

func (r NewSubscriptionMaxGroupTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionMaxGroupTieredPackagePriceModelTypeMaxGroupTieredPackage:
		return true
	}
	return false
}

type NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                            `json:"unit_config"`
}

func (r NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigParam) ImplementsNewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigParam].
type NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigUnionParam()
}

type NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionMaxGroupTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionMinimumCompositePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionMinimumCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// Configuration for minimum pricing
	MinimumConfig param.Field[NewSubscriptionMinimumCompositePriceMinimumConfigParam] `json:"minimum_config,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionMinimumCompositePriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionMinimumCompositePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionMinimumCompositePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMinimumCompositePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMinimumCompositePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionMinimumCompositePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionMinimumCompositePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionMinimumCompositePriceCadence string

const (
	NewSubscriptionMinimumCompositePriceCadenceAnnual     NewSubscriptionMinimumCompositePriceCadence = "annual"
	NewSubscriptionMinimumCompositePriceCadenceSemiAnnual NewSubscriptionMinimumCompositePriceCadence = "semi_annual"
	NewSubscriptionMinimumCompositePriceCadenceMonthly    NewSubscriptionMinimumCompositePriceCadence = "monthly"
	NewSubscriptionMinimumCompositePriceCadenceQuarterly  NewSubscriptionMinimumCompositePriceCadence = "quarterly"
	NewSubscriptionMinimumCompositePriceCadenceOneTime    NewSubscriptionMinimumCompositePriceCadence = "one_time"
	NewSubscriptionMinimumCompositePriceCadenceCustom     NewSubscriptionMinimumCompositePriceCadence = "custom"
)

func (r NewSubscriptionMinimumCompositePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionMinimumCompositePriceCadenceAnnual, NewSubscriptionMinimumCompositePriceCadenceSemiAnnual, NewSubscriptionMinimumCompositePriceCadenceMonthly, NewSubscriptionMinimumCompositePriceCadenceQuarterly, NewSubscriptionMinimumCompositePriceCadenceOneTime, NewSubscriptionMinimumCompositePriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for minimum pricing
type NewSubscriptionMinimumCompositePriceMinimumConfigParam struct {
	// The minimum amount to apply
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// If true, subtotals from this price are prorated based on the service period
	Prorated param.Field[bool] `json:"prorated"`
}

func (r NewSubscriptionMinimumCompositePriceMinimumConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type NewSubscriptionMinimumCompositePriceModelType string

const (
	NewSubscriptionMinimumCompositePriceModelTypeMinimum NewSubscriptionMinimumCompositePriceModelType = "minimum"
)

func (r NewSubscriptionMinimumCompositePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionMinimumCompositePriceModelTypeMinimum:
		return true
	}
	return false
}

type NewSubscriptionMinimumCompositePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                     `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                       `json:"unit_config"`
}

func (r NewSubscriptionMinimumCompositePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionMinimumCompositePriceConversionRateConfigParam) ImplementsNewSubscriptionMinimumCompositePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionMinimumCompositePriceConversionRateConfigParam].
type NewSubscriptionMinimumCompositePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionMinimumCompositePriceConversionRateConfigUnionParam()
}

type NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionMinimumCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for package pricing
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
	ConversionRateConfig param.Field[NewSubscriptionPackagePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPackagePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

func (r NewSubscriptionPackagePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {}

func (r NewSubscriptionPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionPackagePriceCadence string

const (
	NewSubscriptionPackagePriceCadenceAnnual     NewSubscriptionPackagePriceCadence = "annual"
	NewSubscriptionPackagePriceCadenceSemiAnnual NewSubscriptionPackagePriceCadence = "semi_annual"
	NewSubscriptionPackagePriceCadenceMonthly    NewSubscriptionPackagePriceCadence = "monthly"
	NewSubscriptionPackagePriceCadenceQuarterly  NewSubscriptionPackagePriceCadence = "quarterly"
	NewSubscriptionPackagePriceCadenceOneTime    NewSubscriptionPackagePriceCadence = "one_time"
	NewSubscriptionPackagePriceCadenceCustom     NewSubscriptionPackagePriceCadence = "custom"
)

func (r NewSubscriptionPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPackagePriceCadenceAnnual, NewSubscriptionPackagePriceCadenceSemiAnnual, NewSubscriptionPackagePriceCadenceMonthly, NewSubscriptionPackagePriceCadenceQuarterly, NewSubscriptionPackagePriceCadenceOneTime, NewSubscriptionPackagePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionPackagePriceModelType string

const (
	NewSubscriptionPackagePriceModelTypePackage NewSubscriptionPackagePriceModelType = "package"
)

func (r NewSubscriptionPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPackagePriceModelTypePackage:
		return true
	}
	return false
}

type NewSubscriptionPackagePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                              `json:"unit_config"`
}

func (r NewSubscriptionPackagePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPackagePriceConversionRateConfigParam) ImplementsNewSubscriptionPackagePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionPackagePriceConversionRateConfigParam].
type NewSubscriptionPackagePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionPackagePriceConversionRateConfigUnionParam()
}

type NewSubscriptionPackagePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionPackagePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionPackagePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionPackagePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionPackagePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionPackageWithAllocationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionPackageWithAllocationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionPackageWithAllocationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for package_with_allocation pricing
	PackageWithAllocationConfig param.Field[NewSubscriptionPackageWithAllocationPricePackageWithAllocationConfigParam] `json:"package_with_allocation_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionPackageWithAllocationPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionPackageWithAllocationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPackageWithAllocationPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionPackageWithAllocationPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionPackageWithAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionPackageWithAllocationPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionPackageWithAllocationPriceCadence string

const (
	NewSubscriptionPackageWithAllocationPriceCadenceAnnual     NewSubscriptionPackageWithAllocationPriceCadence = "annual"
	NewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual NewSubscriptionPackageWithAllocationPriceCadence = "semi_annual"
	NewSubscriptionPackageWithAllocationPriceCadenceMonthly    NewSubscriptionPackageWithAllocationPriceCadence = "monthly"
	NewSubscriptionPackageWithAllocationPriceCadenceQuarterly  NewSubscriptionPackageWithAllocationPriceCadence = "quarterly"
	NewSubscriptionPackageWithAllocationPriceCadenceOneTime    NewSubscriptionPackageWithAllocationPriceCadence = "one_time"
	NewSubscriptionPackageWithAllocationPriceCadenceCustom     NewSubscriptionPackageWithAllocationPriceCadence = "custom"
)

func (r NewSubscriptionPackageWithAllocationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionPackageWithAllocationPriceCadenceAnnual, NewSubscriptionPackageWithAllocationPriceCadenceSemiAnnual, NewSubscriptionPackageWithAllocationPriceCadenceMonthly, NewSubscriptionPackageWithAllocationPriceCadenceQuarterly, NewSubscriptionPackageWithAllocationPriceCadenceOneTime, NewSubscriptionPackageWithAllocationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionPackageWithAllocationPriceModelType string

const (
	NewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation NewSubscriptionPackageWithAllocationPriceModelType = "package_with_allocation"
)

func (r NewSubscriptionPackageWithAllocationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionPackageWithAllocationPriceModelTypePackageWithAllocation:
		return true
	}
	return false
}

// Configuration for package_with_allocation pricing
type NewSubscriptionPackageWithAllocationPricePackageWithAllocationConfigParam struct {
	// Usage allocation
	Allocation param.Field[string] `json:"allocation,required"`
	// Price per package
	PackageAmount param.Field[string] `json:"package_amount,required"`
	// Package size
	PackageSize param.Field[string] `json:"package_size,required"`
}

func (r NewSubscriptionPackageWithAllocationPricePackageWithAllocationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionPackageWithAllocationPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                            `json:"unit_config"`
}

func (r NewSubscriptionPackageWithAllocationPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionPackageWithAllocationPriceConversionRateConfigParam) ImplementsNewSubscriptionPackageWithAllocationPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionPackageWithAllocationPriceConversionRateConfigParam].
type NewSubscriptionPackageWithAllocationPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionPackageWithAllocationPriceConversionRateConfigUnionParam()
}

type NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionPackageWithAllocationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionScalableMatrixWithTieredPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionScalableMatrixWithTieredPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionScalableMatrixWithTieredPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for scalable_matrix_with_tiered_pricing pricing
	ScalableMatrixWithTieredPricingConfig param.Field[NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam] `json:"scalable_matrix_with_tiered_pricing_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionScalableMatrixWithTieredPricingPriceCadence string

const (
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual     NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "annual"
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "semi_annual"
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly    NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "monthly"
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly  NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "quarterly"
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime    NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "one_time"
	NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom     NewSubscriptionScalableMatrixWithTieredPricingPriceCadence = "custom"
)

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceAnnual, NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceSemiAnnual, NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceMonthly, NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceQuarterly, NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceOneTime, NewSubscriptionScalableMatrixWithTieredPricingPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionScalableMatrixWithTieredPricingPriceModelType string

const (
	NewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing NewSubscriptionScalableMatrixWithTieredPricingPriceModelType = "scalable_matrix_with_tiered_pricing"
)

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithTieredPricingPriceModelTypeScalableMatrixWithTieredPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_tiered_pricing pricing
type NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam struct {
	// Used for the scalable matrix first dimension
	FirstDimension param.Field[string] `json:"first_dimension,required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam] `json:"matrix_scaling_factors,required"`
	// Tier pricing structure
	Tiers param.Field[[]NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam] `json:"tiers,required"`
	// Used for the scalable matrix second dimension (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam struct {
	// First dimension value
	FirstDimensionValue param.Field[string] `json:"first_dimension_value,required"`
	// Scaling factor
	ScalingFactor param.Field[string] `json:"scaling_factor,required"`
	// Second dimension value (optional)
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigMatrixScalingFactorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier entry with business logic
type NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam struct {
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceScalableMatrixWithTieredPricingConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                      `json:"unit_config"`
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigParam) ImplementsNewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigParam].
type NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigUnionParam()
}

type NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionScalableMatrixWithTieredPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionScalableMatrixWithUnitPricingPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionScalableMatrixWithUnitPricingPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionScalableMatrixWithUnitPricingPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for scalable_matrix_with_unit_pricing pricing
	ScalableMatrixWithUnitPricingConfig param.Field[NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam] `json:"scalable_matrix_with_unit_pricing_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionScalableMatrixWithUnitPricingPriceCadence string

const (
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual     NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "annual"
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "semi_annual"
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly    NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "monthly"
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly  NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "quarterly"
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime    NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "one_time"
	NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom     NewSubscriptionScalableMatrixWithUnitPricingPriceCadence = "custom"
)

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceAnnual, NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceSemiAnnual, NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceMonthly, NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceQuarterly, NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceOneTime, NewSubscriptionScalableMatrixWithUnitPricingPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionScalableMatrixWithUnitPricingPriceModelType string

const (
	NewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing NewSubscriptionScalableMatrixWithUnitPricingPriceModelType = "scalable_matrix_with_unit_pricing"
)

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithUnitPricingPriceModelTypeScalableMatrixWithUnitPricing:
		return true
	}
	return false
}

// Configuration for scalable_matrix_with_unit_pricing pricing
type NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam struct {
	// Used to determine the unit rate
	FirstDimension param.Field[string] `json:"first_dimension,required"`
	// Apply a scaling factor to each dimension
	MatrixScalingFactors param.Field[[]NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam] `json:"matrix_scaling_factors,required"`
	// The final unit price to rate against the output of the matrix
	UnitPrice param.Field[string] `json:"unit_price,required"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
	// Used to determine the unit rate (optional)
	SecondDimension param.Field[string] `json:"second_dimension"`
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single matrix scaling factor
type NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam struct {
	// First dimension value
	FirstDimensionValue param.Field[string] `json:"first_dimension_value,required"`
	// Scaling factor
	ScalingFactor param.Field[string] `json:"scaling_factor,required"`
	// Second dimension value (optional)
	SecondDimensionValue param.Field[string] `json:"second_dimension_value"`
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceScalableMatrixWithUnitPricingConfigMatrixScalingFactorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                    `json:"unit_config"`
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigParam) ImplementsNewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigParam].
type NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigUnionParam()
}

type NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionScalableMatrixWithUnitPricingPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionThresholdTotalAmountPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionThresholdTotalAmountPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionThresholdTotalAmountPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for threshold_total_amount pricing
	ThresholdTotalAmountConfig param.Field[NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigParam] `json:"threshold_total_amount_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionThresholdTotalAmountPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionThresholdTotalAmountPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionThresholdTotalAmountPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionThresholdTotalAmountPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionThresholdTotalAmountPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionThresholdTotalAmountPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionThresholdTotalAmountPriceCadence string

const (
	NewSubscriptionThresholdTotalAmountPriceCadenceAnnual     NewSubscriptionThresholdTotalAmountPriceCadence = "annual"
	NewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual NewSubscriptionThresholdTotalAmountPriceCadence = "semi_annual"
	NewSubscriptionThresholdTotalAmountPriceCadenceMonthly    NewSubscriptionThresholdTotalAmountPriceCadence = "monthly"
	NewSubscriptionThresholdTotalAmountPriceCadenceQuarterly  NewSubscriptionThresholdTotalAmountPriceCadence = "quarterly"
	NewSubscriptionThresholdTotalAmountPriceCadenceOneTime    NewSubscriptionThresholdTotalAmountPriceCadence = "one_time"
	NewSubscriptionThresholdTotalAmountPriceCadenceCustom     NewSubscriptionThresholdTotalAmountPriceCadence = "custom"
)

func (r NewSubscriptionThresholdTotalAmountPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionThresholdTotalAmountPriceCadenceAnnual, NewSubscriptionThresholdTotalAmountPriceCadenceSemiAnnual, NewSubscriptionThresholdTotalAmountPriceCadenceMonthly, NewSubscriptionThresholdTotalAmountPriceCadenceQuarterly, NewSubscriptionThresholdTotalAmountPriceCadenceOneTime, NewSubscriptionThresholdTotalAmountPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionThresholdTotalAmountPriceModelType string

const (
	NewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount NewSubscriptionThresholdTotalAmountPriceModelType = "threshold_total_amount"
)

func (r NewSubscriptionThresholdTotalAmountPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionThresholdTotalAmountPriceModelTypeThresholdTotalAmount:
		return true
	}
	return false
}

// Configuration for threshold_total_amount pricing
type NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigParam struct {
	// When the quantity consumed passes a provided threshold, the configured total
	// will be charged
	ConsumptionTable param.Field[[]NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam] `json:"consumption_table,required"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single threshold
type NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam struct {
	// Quantity threshold
	Threshold param.Field[string] `json:"threshold,required"`
	// Total amount for this threshold
	TotalAmount param.Field[string] `json:"total_amount,required"`
}

func (r NewSubscriptionThresholdTotalAmountPriceThresholdTotalAmountConfigConsumptionTableParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionThresholdTotalAmountPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                           `json:"unit_config"`
}

func (r NewSubscriptionThresholdTotalAmountPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionThresholdTotalAmountPriceConversionRateConfigParam) ImplementsNewSubscriptionThresholdTotalAmountPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionThresholdTotalAmountPriceConversionRateConfigParam].
type NewSubscriptionThresholdTotalAmountPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionThresholdTotalAmountPriceConversionRateConfigUnionParam()
}

type NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionThresholdTotalAmountPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionTieredPackagePriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionTieredPackagePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionTieredPackagePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_package pricing
	TieredPackageConfig param.Field[NewSubscriptionTieredPackagePriceTieredPackageConfigParam] `json:"tiered_package_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionTieredPackagePriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionTieredPackagePriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPackagePriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredPackagePriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredPackagePriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionTieredPackagePriceCadence string

const (
	NewSubscriptionTieredPackagePriceCadenceAnnual     NewSubscriptionTieredPackagePriceCadence = "annual"
	NewSubscriptionTieredPackagePriceCadenceSemiAnnual NewSubscriptionTieredPackagePriceCadence = "semi_annual"
	NewSubscriptionTieredPackagePriceCadenceMonthly    NewSubscriptionTieredPackagePriceCadence = "monthly"
	NewSubscriptionTieredPackagePriceCadenceQuarterly  NewSubscriptionTieredPackagePriceCadence = "quarterly"
	NewSubscriptionTieredPackagePriceCadenceOneTime    NewSubscriptionTieredPackagePriceCadence = "one_time"
	NewSubscriptionTieredPackagePriceCadenceCustom     NewSubscriptionTieredPackagePriceCadence = "custom"
)

func (r NewSubscriptionTieredPackagePriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackagePriceCadenceAnnual, NewSubscriptionTieredPackagePriceCadenceSemiAnnual, NewSubscriptionTieredPackagePriceCadenceMonthly, NewSubscriptionTieredPackagePriceCadenceQuarterly, NewSubscriptionTieredPackagePriceCadenceOneTime, NewSubscriptionTieredPackagePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionTieredPackagePriceModelType string

const (
	NewSubscriptionTieredPackagePriceModelTypeTieredPackage NewSubscriptionTieredPackagePriceModelType = "tiered_package"
)

func (r NewSubscriptionTieredPackagePriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackagePriceModelTypeTieredPackage:
		return true
	}
	return false
}

// Configuration for tiered_package pricing
type NewSubscriptionTieredPackagePriceTieredPackageConfigParam struct {
	// Package size
	PackageSize param.Field[string] `json:"package_size,required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds. The tier bounds are defined based on
	// the total quantity rather than the number of packages, so they must be multiples
	// of the package size.
	Tiers param.Field[[]NewSubscriptionTieredPackagePriceTieredPackageConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionTieredPackagePriceTieredPackageConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier with business logic
type NewSubscriptionTieredPackagePriceTieredPackageConfigTierParam struct {
	// Price per package
	PerUnit param.Field[string] `json:"per_unit,required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
}

func (r NewSubscriptionTieredPackagePriceTieredPackageConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionTieredPackagePriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                    `json:"unit_config"`
}

func (r NewSubscriptionTieredPackagePriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPackagePriceConversionRateConfigParam) ImplementsNewSubscriptionTieredPackagePriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionTieredPackagePriceConversionRateConfigParam].
type NewSubscriptionTieredPackagePriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionTieredPackagePriceConversionRateConfigUnionParam()
}

type NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateTypeTiered NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionTieredPackagePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionTieredPackageWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionTieredPackageWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionTieredPackageWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_package_with_minimum pricing
	TieredPackageWithMinimumConfig param.Field[NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam] `json:"tiered_package_with_minimum_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionTieredPackageWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPackageWithMinimumPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredPackageWithMinimumPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionTieredPackageWithMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredPackageWithMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionTieredPackageWithMinimumPriceCadence string

const (
	NewSubscriptionTieredPackageWithMinimumPriceCadenceAnnual     NewSubscriptionTieredPackageWithMinimumPriceCadence = "annual"
	NewSubscriptionTieredPackageWithMinimumPriceCadenceSemiAnnual NewSubscriptionTieredPackageWithMinimumPriceCadence = "semi_annual"
	NewSubscriptionTieredPackageWithMinimumPriceCadenceMonthly    NewSubscriptionTieredPackageWithMinimumPriceCadence = "monthly"
	NewSubscriptionTieredPackageWithMinimumPriceCadenceQuarterly  NewSubscriptionTieredPackageWithMinimumPriceCadence = "quarterly"
	NewSubscriptionTieredPackageWithMinimumPriceCadenceOneTime    NewSubscriptionTieredPackageWithMinimumPriceCadence = "one_time"
	NewSubscriptionTieredPackageWithMinimumPriceCadenceCustom     NewSubscriptionTieredPackageWithMinimumPriceCadence = "custom"
)

func (r NewSubscriptionTieredPackageWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackageWithMinimumPriceCadenceAnnual, NewSubscriptionTieredPackageWithMinimumPriceCadenceSemiAnnual, NewSubscriptionTieredPackageWithMinimumPriceCadenceMonthly, NewSubscriptionTieredPackageWithMinimumPriceCadenceQuarterly, NewSubscriptionTieredPackageWithMinimumPriceCadenceOneTime, NewSubscriptionTieredPackageWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionTieredPackageWithMinimumPriceModelType string

const (
	NewSubscriptionTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum NewSubscriptionTieredPackageWithMinimumPriceModelType = "tiered_package_with_minimum"
)

func (r NewSubscriptionTieredPackageWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackageWithMinimumPriceModelTypeTieredPackageWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_package_with_minimum pricing
type NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam struct {
	// Package size
	PackageSize param.Field[float64] `json:"package_size,required"`
	// Apply tiered pricing after rounding up the quantity to the package size. Tiers
	// are defined using exclusive lower bounds.
	Tiers param.Field[[]NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam] `json:"tiers,required"`
}

func (r NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam struct {
	// Minimum amount
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Price per package
	PerUnit param.Field[string] `json:"per_unit,required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
}

func (r NewSubscriptionTieredPackageWithMinimumPriceTieredPackageWithMinimumConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                             `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                               `json:"unit_config"`
}

func (r NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigParam) ImplementsNewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigParam].
type NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigUnionParam()
}

type NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionTieredPackageWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionTieredPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionTieredPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionTieredPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered pricing
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
	ConversionRateConfig param.Field[NewSubscriptionTieredPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionTieredPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

func (r NewSubscriptionTieredPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {}

func (r NewSubscriptionTieredPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionTieredPriceCadence string

const (
	NewSubscriptionTieredPriceCadenceAnnual     NewSubscriptionTieredPriceCadence = "annual"
	NewSubscriptionTieredPriceCadenceSemiAnnual NewSubscriptionTieredPriceCadence = "semi_annual"
	NewSubscriptionTieredPriceCadenceMonthly    NewSubscriptionTieredPriceCadence = "monthly"
	NewSubscriptionTieredPriceCadenceQuarterly  NewSubscriptionTieredPriceCadence = "quarterly"
	NewSubscriptionTieredPriceCadenceOneTime    NewSubscriptionTieredPriceCadence = "one_time"
	NewSubscriptionTieredPriceCadenceCustom     NewSubscriptionTieredPriceCadence = "custom"
)

func (r NewSubscriptionTieredPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPriceCadenceAnnual, NewSubscriptionTieredPriceCadenceSemiAnnual, NewSubscriptionTieredPriceCadenceMonthly, NewSubscriptionTieredPriceCadenceQuarterly, NewSubscriptionTieredPriceCadenceOneTime, NewSubscriptionTieredPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionTieredPriceModelType string

const (
	NewSubscriptionTieredPriceModelTypeTiered NewSubscriptionTieredPriceModelType = "tiered"
)

func (r NewSubscriptionTieredPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPriceModelTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionTieredPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionTieredPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                             `json:"unit_config"`
}

func (r NewSubscriptionTieredPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredPriceConversionRateConfigParam) ImplementsNewSubscriptionTieredPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionTieredPriceConversionRateConfigParam].
type NewSubscriptionTieredPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionTieredPriceConversionRateConfigUnionParam()
}

type NewSubscriptionTieredPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionTieredPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionTieredPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionTieredPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionTieredPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionTieredPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionTieredPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionTieredWithMinimumPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionTieredWithMinimumPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionTieredWithMinimumPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_minimum pricing
	TieredWithMinimumConfig param.Field[NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigParam] `json:"tiered_with_minimum_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionTieredWithMinimumPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionTieredWithMinimumPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredWithMinimumPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredWithMinimumPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionTieredWithMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionTieredWithMinimumPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionTieredWithMinimumPriceCadence string

const (
	NewSubscriptionTieredWithMinimumPriceCadenceAnnual     NewSubscriptionTieredWithMinimumPriceCadence = "annual"
	NewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual NewSubscriptionTieredWithMinimumPriceCadence = "semi_annual"
	NewSubscriptionTieredWithMinimumPriceCadenceMonthly    NewSubscriptionTieredWithMinimumPriceCadence = "monthly"
	NewSubscriptionTieredWithMinimumPriceCadenceQuarterly  NewSubscriptionTieredWithMinimumPriceCadence = "quarterly"
	NewSubscriptionTieredWithMinimumPriceCadenceOneTime    NewSubscriptionTieredWithMinimumPriceCadence = "one_time"
	NewSubscriptionTieredWithMinimumPriceCadenceCustom     NewSubscriptionTieredWithMinimumPriceCadence = "custom"
)

func (r NewSubscriptionTieredWithMinimumPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredWithMinimumPriceCadenceAnnual, NewSubscriptionTieredWithMinimumPriceCadenceSemiAnnual, NewSubscriptionTieredWithMinimumPriceCadenceMonthly, NewSubscriptionTieredWithMinimumPriceCadenceQuarterly, NewSubscriptionTieredWithMinimumPriceCadenceOneTime, NewSubscriptionTieredWithMinimumPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionTieredWithMinimumPriceModelType string

const (
	NewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum NewSubscriptionTieredWithMinimumPriceModelType = "tiered_with_minimum"
)

func (r NewSubscriptionTieredWithMinimumPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredWithMinimumPriceModelTypeTieredWithMinimum:
		return true
	}
	return false
}

// Configuration for tiered_with_minimum pricing
type NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigParam struct {
	// Tiered pricing with a minimum amount dependent on the volume tier. Tiers are
	// defined using exclusive lower bounds.
	Tiers param.Field[[]NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigTierParam] `json:"tiers,required"`
	// If true, tiers with an accrued amount of 0 will not be included in the rating.
	HideZeroAmountTiers param.Field[bool] `json:"hide_zero_amount_tiers"`
	// If true, the unit price will be prorated to the billing period
	Prorate param.Field[bool] `json:"prorate"`
}

func (r NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tier
type NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigTierParam struct {
	// Minimum amount
	MinimumAmount param.Field[string] `json:"minimum_amount,required"`
	// Tier lower bound
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Per unit amount
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionTieredWithMinimumPriceTieredWithMinimumConfigTierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionTieredWithMinimumPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                        `json:"unit_config"`
}

func (r NewSubscriptionTieredWithMinimumPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionTieredWithMinimumPriceConversionRateConfigParam) ImplementsNewSubscriptionTieredWithMinimumPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionTieredWithMinimumPriceConversionRateConfigParam].
type NewSubscriptionTieredWithMinimumPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionTieredWithMinimumPriceConversionRateConfigUnionParam()
}

type NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionTieredWithMinimumPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionUnitPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionUnitPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionUnitPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for unit pricing
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
	ConversionRateConfig param.Field[NewSubscriptionUnitPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionUnitPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

func (r NewSubscriptionUnitPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {}

func (r NewSubscriptionUnitPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionUnitPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionUnitPriceCadence string

const (
	NewSubscriptionUnitPriceCadenceAnnual     NewSubscriptionUnitPriceCadence = "annual"
	NewSubscriptionUnitPriceCadenceSemiAnnual NewSubscriptionUnitPriceCadence = "semi_annual"
	NewSubscriptionUnitPriceCadenceMonthly    NewSubscriptionUnitPriceCadence = "monthly"
	NewSubscriptionUnitPriceCadenceQuarterly  NewSubscriptionUnitPriceCadence = "quarterly"
	NewSubscriptionUnitPriceCadenceOneTime    NewSubscriptionUnitPriceCadence = "one_time"
	NewSubscriptionUnitPriceCadenceCustom     NewSubscriptionUnitPriceCadence = "custom"
)

func (r NewSubscriptionUnitPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitPriceCadenceAnnual, NewSubscriptionUnitPriceCadenceSemiAnnual, NewSubscriptionUnitPriceCadenceMonthly, NewSubscriptionUnitPriceCadenceQuarterly, NewSubscriptionUnitPriceCadenceOneTime, NewSubscriptionUnitPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionUnitPriceModelType string

const (
	NewSubscriptionUnitPriceModelTypeUnit NewSubscriptionUnitPriceModelType = "unit"
)

func (r NewSubscriptionUnitPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitPriceModelTypeUnit:
		return true
	}
	return false
}

type NewSubscriptionUnitPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionUnitPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                           `json:"unit_config"`
}

func (r NewSubscriptionUnitPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitPriceConversionRateConfigParam) ImplementsNewSubscriptionUnitPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionUnitPriceConversionRateConfigParam].
type NewSubscriptionUnitPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionUnitPriceConversionRateConfigUnionParam()
}

type NewSubscriptionUnitPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionUnitPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionUnitPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionUnitPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionUnitPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionUnitPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionUnitPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionUnitWithPercentPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionUnitWithPercentPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionUnitWithPercentPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for unit_with_percent pricing
	UnitWithPercentConfig param.Field[NewSubscriptionUnitWithPercentPriceUnitWithPercentConfigParam] `json:"unit_with_percent_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionUnitWithPercentPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionUnitWithPercentPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitWithPercentPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionUnitWithPercentPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionUnitWithPercentPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionUnitWithPercentPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionUnitWithPercentPriceCadence string

const (
	NewSubscriptionUnitWithPercentPriceCadenceAnnual     NewSubscriptionUnitWithPercentPriceCadence = "annual"
	NewSubscriptionUnitWithPercentPriceCadenceSemiAnnual NewSubscriptionUnitWithPercentPriceCadence = "semi_annual"
	NewSubscriptionUnitWithPercentPriceCadenceMonthly    NewSubscriptionUnitWithPercentPriceCadence = "monthly"
	NewSubscriptionUnitWithPercentPriceCadenceQuarterly  NewSubscriptionUnitWithPercentPriceCadence = "quarterly"
	NewSubscriptionUnitWithPercentPriceCadenceOneTime    NewSubscriptionUnitWithPercentPriceCadence = "one_time"
	NewSubscriptionUnitWithPercentPriceCadenceCustom     NewSubscriptionUnitWithPercentPriceCadence = "custom"
)

func (r NewSubscriptionUnitWithPercentPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithPercentPriceCadenceAnnual, NewSubscriptionUnitWithPercentPriceCadenceSemiAnnual, NewSubscriptionUnitWithPercentPriceCadenceMonthly, NewSubscriptionUnitWithPercentPriceCadenceQuarterly, NewSubscriptionUnitWithPercentPriceCadenceOneTime, NewSubscriptionUnitWithPercentPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionUnitWithPercentPriceModelType string

const (
	NewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent NewSubscriptionUnitWithPercentPriceModelType = "unit_with_percent"
)

func (r NewSubscriptionUnitWithPercentPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithPercentPriceModelTypeUnitWithPercent:
		return true
	}
	return false
}

// Configuration for unit_with_percent pricing
type NewSubscriptionUnitWithPercentPriceUnitWithPercentConfigParam struct {
	// What percent, out of 100, of the calculated total to charge
	Percent param.Field[string] `json:"percent,required"`
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionUnitWithPercentPriceUnitWithPercentConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionUnitWithPercentPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                      `json:"unit_config"`
}

func (r NewSubscriptionUnitWithPercentPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitWithPercentPriceConversionRateConfigParam) ImplementsNewSubscriptionUnitWithPercentPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionUnitWithPercentPriceConversionRateConfigParam].
type NewSubscriptionUnitWithPercentPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionUnitWithPercentPriceConversionRateConfigUnionParam()
}

type NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionUnitWithPercentPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type NewSubscriptionUnitWithProrationPriceParam struct {
	// The cadence to bill for this price on.
	Cadence param.Field[NewSubscriptionUnitWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[NewSubscriptionUnitWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for unit_with_proration pricing
	UnitWithProrationConfig param.Field[NewSubscriptionUnitWithProrationPriceUnitWithProrationConfigParam] `json:"unit_with_proration_config,required"`
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
	ConversionRateConfig param.Field[NewSubscriptionUnitWithProrationPriceConversionRateConfigUnionParam] `json:"conversion_rate_config"`
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

func (r NewSubscriptionUnitWithProrationPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitWithProrationPriceParam) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionUnitWithProrationPriceParam) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

func (r NewSubscriptionUnitWithProrationPriceParam) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

func (r NewSubscriptionUnitWithProrationPriceParam) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type NewSubscriptionUnitWithProrationPriceCadence string

const (
	NewSubscriptionUnitWithProrationPriceCadenceAnnual     NewSubscriptionUnitWithProrationPriceCadence = "annual"
	NewSubscriptionUnitWithProrationPriceCadenceSemiAnnual NewSubscriptionUnitWithProrationPriceCadence = "semi_annual"
	NewSubscriptionUnitWithProrationPriceCadenceMonthly    NewSubscriptionUnitWithProrationPriceCadence = "monthly"
	NewSubscriptionUnitWithProrationPriceCadenceQuarterly  NewSubscriptionUnitWithProrationPriceCadence = "quarterly"
	NewSubscriptionUnitWithProrationPriceCadenceOneTime    NewSubscriptionUnitWithProrationPriceCadence = "one_time"
	NewSubscriptionUnitWithProrationPriceCadenceCustom     NewSubscriptionUnitWithProrationPriceCadence = "custom"
)

func (r NewSubscriptionUnitWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithProrationPriceCadenceAnnual, NewSubscriptionUnitWithProrationPriceCadenceSemiAnnual, NewSubscriptionUnitWithProrationPriceCadenceMonthly, NewSubscriptionUnitWithProrationPriceCadenceQuarterly, NewSubscriptionUnitWithProrationPriceCadenceOneTime, NewSubscriptionUnitWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type NewSubscriptionUnitWithProrationPriceModelType string

const (
	NewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration NewSubscriptionUnitWithProrationPriceModelType = "unit_with_proration"
)

func (r NewSubscriptionUnitWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithProrationPriceModelTypeUnitWithProration:
		return true
	}
	return false
}

// Configuration for unit_with_proration pricing
type NewSubscriptionUnitWithProrationPriceUnitWithProrationConfigParam struct {
	// Rate per unit of usage
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r NewSubscriptionUnitWithProrationPriceUnitWithProrationConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NewSubscriptionUnitWithProrationPriceConversionRateConfigParam struct {
	ConversionRateType param.Field[NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                        `json:"unit_config"`
}

func (r NewSubscriptionUnitWithProrationPriceConversionRateConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NewSubscriptionUnitWithProrationPriceConversionRateConfigParam) ImplementsNewSubscriptionUnitWithProrationPriceConversionRateConfigUnionParam() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [NewSubscriptionUnitWithProrationPriceConversionRateConfigParam].
type NewSubscriptionUnitWithProrationPriceConversionRateConfigUnionParam interface {
	ImplementsNewSubscriptionUnitWithProrationPriceConversionRateConfigUnionParam()
}

type NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateType string

const (
	NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit   NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateType = "unit"
	NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateTypeUnit, NewSubscriptionUnitWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// A [subscription](/core-concepts#subscription) represents the purchase of a plan
// by a customer.
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
	// The adjustment intervals for this subscription sorted by the start_date of the
	// adjustment interval.
	AdjustmentIntervals []shared.AdjustmentInterval `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                   `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration shared.BillingCycleAnchorConfiguration `json:"billing_cycle_anchor_configuration,required"`
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
	// [Customer ID Aliases](/events-and-metrics/customer-aliases) for further
	// information about how these aliases work in Orb.
	//
	// In addition to having an identifier in your system, a customer may exist in a
	// payment provider solution like Stripe. Use the `payment_provider_id` and the
	// `payment_provider` enum field to express this mapping.
	//
	// A customer also has a timezone (from the standard
	// [IANA timezone database](https://www.iana.org/time-zones)), which defaults to
	// your account's timezone. See [Timezone localization](/essentials/timezones) for
	// information on what this timezone parameter influences within Orb.
	Customer Customer `json:"customer,required"`
	// Determines the default memo on this subscriptions' invoices. Note that if this
	// is not provided, it is determined by the plan configuration.
	DefaultInvoiceMemo string `json:"default_invoice_memo,required,nullable"`
	// The discount intervals for this subscription sorted by the start_date. This
	// field is deprecated in favor of `adjustment_intervals`.
	//
	// Deprecated: deprecated
	DiscountIntervals []SubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                              `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []shared.FixedFeeQuantityScheduleEntry `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                 `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription sorted by the start_date. This field
	// is deprecated in favor of `adjustment_intervals`.
	//
	// Deprecated: deprecated
	MaximumIntervals []shared.MaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription sorted by the start_date. This field
	// is deprecated in favor of `adjustment_intervals`.
	//
	// Deprecated: deprecated
	MinimumIntervals []shared.MinimumInterval `json:"minimum_intervals,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// A pending subscription change if one exists on this subscription.
	PendingSubscriptionChange shared.SubscriptionChangeMinified `json:"pending_subscription_change,required,nullable"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan Plan `json:"plan,required,nullable"`
	// The price intervals for this subscription.
	PriceIntervals []shared.PriceInterval  `json:"price_intervals,required"`
	RedeemedCoupon shared.CouponRedemption `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                    `json:"start_date,required" format:"date-time"`
	Status    SubscriptionStatus           `json:"status,required"`
	TrialInfo shared.SubscriptionTrialInfo `json:"trial_info,required"`
	JSON      subscriptionJSON             `json:"-"`
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
	Name                            apijson.Field
	NetTerms                        apijson.Field
	PendingSubscriptionChange       apijson.Field
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

type SubscriptionDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                               `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of [[]shared.AmountDiscountIntervalFilter],
	// [[]shared.PercentageDiscountIntervalFilter],
	// [[]shared.UsageDiscountIntervalFilter].
	Filters interface{} `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount"`
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
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	AmountDiscount            apijson.Field
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
// Possible runtime types of the union are [shared.AmountDiscountInterval],
// [shared.PercentageDiscountInterval], [shared.UsageDiscountInterval].
func (r SubscriptionDiscountInterval) AsUnion() SubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by [shared.AmountDiscountInterval],
// [shared.PercentageDiscountInterval] or [shared.UsageDiscountInterval].
type SubscriptionDiscountIntervalsUnion interface {
	ImplementsSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.AmountDiscountInterval{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.PercentageDiscountInterval{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(shared.UsageDiscountInterval{}),
			DiscriminatorValue: "usage",
		},
	)
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
	Data []shared.AggregatedCost            `json:"data,required"`
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
	Plan      SubscriptionFetchScheduleResponsePlan `json:"plan,required,nullable"`
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
	// Additional adjustments to be added to the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	AddAdjustments param.Field[[]SubscriptionNewParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	AddPrices                             param.Field[[]SubscriptionNewParamsAddPrice] `json:"add_prices"`
	AlignBillingWithSubscriptionStartDate param.Field[bool]                            `json:"align_billing_with_subscription_start_date"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. If not specified, this
	// defaults to the behavior configured for this customer.
	AutoCollection                  param.Field[bool]                                        `json:"auto_collection"`
	AwsRegion                       param.Field[string]                                      `json:"aws_region"`
	BillingCycleAnchorConfiguration param.Field[shared.BillingCycleAnchorConfigurationParam] `json:"billing_cycle_anchor_configuration"`
	// Redemption code to be used for this subscription. If the coupon cannot be found
	// by its redemption code, or cannot be redeemed, an error response will be
	// returned and the subscription creation or plan change will not be scheduled.
	CouponRedemptionCode param.Field[string]  `json:"coupon_redemption_code"`
	CreditsOverageRate   param.Field[float64] `json:"credits_overage_rate"`
	// The currency to use for the subscription. If not specified, the invoicing
	// currency for the plan will be used.
	Currency   param.Field[string] `json:"currency"`
	CustomerID param.Field[string] `json:"customer_id"`
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
	// The name to use for the subscription. If not specified, the plan name will be
	// used.
	Name param.Field[string] `json:"name"`
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
	RemoveAdjustments param.Field[[]SubscriptionNewParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Plan prices to be removed from the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	RemovePrices param.Field[[]SubscriptionNewParamsRemovePrice] `json:"remove_prices"`
	// Plan adjustments to be replaced with additional adjustments on the subscription.
	// (Only available for accounts that have migrated off of legacy subscription
	// overrides)
	ReplaceAdjustments param.Field[[]SubscriptionNewParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Plan prices to be replaced with additional prices on the subscription. (Only
	// available for accounts that have migrated off of legacy subscription overrides)
	ReplacePrices param.Field[[]SubscriptionNewParamsReplacePrice] `json:"replace_prices"`
	StartDate     param.Field[time.Time]                           `json:"start_date" format:"date-time"`
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

type SubscriptionNewParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription.
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
	AdjustmentType param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                      `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                               `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                               `json:"applies_to_price_ids"`
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
	PriceType     param.Field[SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                `json:"usage_discount"`
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddAdjustmentsAdjustment) ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [SubscriptionNewParamsAddAdjustmentsAdjustment].
type SubscriptionNewParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsSubscriptionNewParamsAddAdjustmentsAdjustmentUnion()
}

type SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionNewParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAllTrue SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType string

const (
	SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeUsage          SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType = "usage"
	SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixed          SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeInArrears      SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r SubscriptionNewParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeUsage, SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeFixed, SubscriptionNewParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPrice struct {
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for this
	// price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideParam] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription. If null, billing will end when the phase or
	// subscription ends.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// this price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// this price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New subscription price request body params.
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

// New subscription price request body params.
type SubscriptionNewParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceModelType] `json:"model_type,required"`
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

func (r SubscriptionNewParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {}

// New subscription price request body params.
//
// Satisfied by [NewSubscriptionUnitPriceParam], [NewSubscriptionTieredPriceParam],
// [NewSubscriptionBulkPriceParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice],
// [NewSubscriptionPackagePriceParam], [NewSubscriptionMatrixPriceParam],
// [NewSubscriptionThresholdTotalAmountPriceParam],
// [NewSubscriptionTieredPackagePriceParam],
// [NewSubscriptionTieredWithMinimumPriceParam],
// [NewSubscriptionGroupedTieredPriceParam],
// [NewSubscriptionTieredPackageWithMinimumPriceParam],
// [NewSubscriptionPackageWithAllocationPriceParam],
// [NewSubscriptionUnitWithPercentPriceParam],
// [NewSubscriptionMatrixWithAllocationPriceParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice],
// [NewSubscriptionUnitWithProrationPriceParam],
// [NewSubscriptionGroupedAllocationPriceParam],
// [NewSubscriptionBulkWithProrationPriceParam],
// [NewSubscriptionGroupedWithProratedMinimumPriceParam],
// [NewSubscriptionGroupedWithMeteredMinimumPriceParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice],
// [NewSubscriptionMatrixWithDisplayNamePriceParam],
// [NewSubscriptionGroupedTieredPackagePriceParam],
// [NewSubscriptionMaxGroupTieredPackagePriceParam],
// [NewSubscriptionScalableMatrixWithUnitPricingPriceParam],
// [NewSubscriptionScalableMatrixWithTieredPricingPriceParam],
// [NewSubscriptionCumulativeGroupedBulkPriceParam],
// [NewSubscriptionMinimumCompositePriceParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePrice],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPrice],
// [SubscriptionNewParamsAddPricesPrice].
type SubscriptionNewParamsAddPricesPriceUnion interface {
	implementsSubscriptionNewParamsAddPricesPriceUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig].
type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                             `json:"unit_config"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig].
type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                     `json:"unit_config"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelTypePercent SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType = "percent"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                          `json:"unit_config"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig].
type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPrice) implementsSubscriptionNewParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceAnnual     SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "semi_annual"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceMonthly    SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "monthly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly  SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "quarterly"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceOneTime    SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "one_time"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceCustom     SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "custom"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceMonthly, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceOneTime, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig struct {
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

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType = "event_output"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                   `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                     `json:"unit_config"`
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig].
type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type SubscriptionNewParamsAddPricesPriceModelType string

const (
	SubscriptionNewParamsAddPricesPriceModelTypeUnit                            SubscriptionNewParamsAddPricesPriceModelType = "unit"
	SubscriptionNewParamsAddPricesPriceModelTypeTiered                          SubscriptionNewParamsAddPricesPriceModelType = "tiered"
	SubscriptionNewParamsAddPricesPriceModelTypeBulk                            SubscriptionNewParamsAddPricesPriceModelType = "bulk"
	SubscriptionNewParamsAddPricesPriceModelTypeBulkWithFilters                 SubscriptionNewParamsAddPricesPriceModelType = "bulk_with_filters"
	SubscriptionNewParamsAddPricesPriceModelTypePackage                         SubscriptionNewParamsAddPricesPriceModelType = "package"
	SubscriptionNewParamsAddPricesPriceModelTypeMatrix                          SubscriptionNewParamsAddPricesPriceModelType = "matrix"
	SubscriptionNewParamsAddPricesPriceModelTypeThresholdTotalAmount            SubscriptionNewParamsAddPricesPriceModelType = "threshold_total_amount"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredPackage                   SubscriptionNewParamsAddPricesPriceModelType = "tiered_package"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredWithMinimum               SubscriptionNewParamsAddPricesPriceModelType = "tiered_with_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedTiered                   SubscriptionNewParamsAddPricesPriceModelType = "grouped_tiered"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum        SubscriptionNewParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypePackageWithAllocation           SubscriptionNewParamsAddPricesPriceModelType = "package_with_allocation"
	SubscriptionNewParamsAddPricesPriceModelTypeUnitWithPercent                 SubscriptionNewParamsAddPricesPriceModelType = "unit_with_percent"
	SubscriptionNewParamsAddPricesPriceModelTypeMatrixWithAllocation            SubscriptionNewParamsAddPricesPriceModelType = "matrix_with_allocation"
	SubscriptionNewParamsAddPricesPriceModelTypeTieredWithProration             SubscriptionNewParamsAddPricesPriceModelType = "tiered_with_proration"
	SubscriptionNewParamsAddPricesPriceModelTypeUnitWithProration               SubscriptionNewParamsAddPricesPriceModelType = "unit_with_proration"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedAllocation               SubscriptionNewParamsAddPricesPriceModelType = "grouped_allocation"
	SubscriptionNewParamsAddPricesPriceModelTypeBulkWithProration               SubscriptionNewParamsAddPricesPriceModelType = "bulk_with_proration"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      SubscriptionNewParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       SubscriptionNewParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     SubscriptionNewParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	SubscriptionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName           SubscriptionNewParamsAddPricesPriceModelType = "matrix_with_display_name"
	SubscriptionNewParamsAddPricesPriceModelTypeGroupedTieredPackage            SubscriptionNewParamsAddPricesPriceModelType = "grouped_tiered_package"
	SubscriptionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage           SubscriptionNewParamsAddPricesPriceModelType = "max_group_tiered_package"
	SubscriptionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   SubscriptionNewParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	SubscriptionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing SubscriptionNewParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	SubscriptionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk           SubscriptionNewParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	SubscriptionNewParamsAddPricesPriceModelTypeMinimum                         SubscriptionNewParamsAddPricesPriceModelType = "minimum"
	SubscriptionNewParamsAddPricesPriceModelTypePercent                         SubscriptionNewParamsAddPricesPriceModelType = "percent"
	SubscriptionNewParamsAddPricesPriceModelTypeEventOutput                     SubscriptionNewParamsAddPricesPriceModelType = "event_output"
)

func (r SubscriptionNewParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsAddPricesPriceModelTypeUnit, SubscriptionNewParamsAddPricesPriceModelTypeTiered, SubscriptionNewParamsAddPricesPriceModelTypeBulk, SubscriptionNewParamsAddPricesPriceModelTypeBulkWithFilters, SubscriptionNewParamsAddPricesPriceModelTypePackage, SubscriptionNewParamsAddPricesPriceModelTypeMatrix, SubscriptionNewParamsAddPricesPriceModelTypeThresholdTotalAmount, SubscriptionNewParamsAddPricesPriceModelTypeTieredPackage, SubscriptionNewParamsAddPricesPriceModelTypeTieredWithMinimum, SubscriptionNewParamsAddPricesPriceModelTypeGroupedTiered, SubscriptionNewParamsAddPricesPriceModelTypeTieredPackageWithMinimum, SubscriptionNewParamsAddPricesPriceModelTypePackageWithAllocation, SubscriptionNewParamsAddPricesPriceModelTypeUnitWithPercent, SubscriptionNewParamsAddPricesPriceModelTypeMatrixWithAllocation, SubscriptionNewParamsAddPricesPriceModelTypeTieredWithProration, SubscriptionNewParamsAddPricesPriceModelTypeUnitWithProration, SubscriptionNewParamsAddPricesPriceModelTypeGroupedAllocation, SubscriptionNewParamsAddPricesPriceModelTypeBulkWithProration, SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, SubscriptionNewParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, SubscriptionNewParamsAddPricesPriceModelTypeMatrixWithDisplayName, SubscriptionNewParamsAddPricesPriceModelTypeGroupedTieredPackage, SubscriptionNewParamsAddPricesPriceModelTypeMaxGroupTieredPackage, SubscriptionNewParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, SubscriptionNewParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, SubscriptionNewParamsAddPricesPriceModelTypeCumulativeGroupedBulk, SubscriptionNewParamsAddPricesPriceModelTypeMinimum, SubscriptionNewParamsAddPricesPriceModelTypePercent, SubscriptionNewParamsAddPricesPriceModelTypeEventOutput:
		return true
	}
	return false
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
	AdjustmentType param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                          `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                   `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                   `json:"applies_to_price_ids"`
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
	PriceType     param.Field[SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                    `json:"usage_discount"`
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustment) ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [SubscriptionNewParamsReplaceAdjustmentsAdjustment].
type SubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsSubscriptionNewParamsReplaceAdjustmentsAdjustmentUnion()
}

type SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionNewParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, SubscriptionNewParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePrice struct {
	// The id of the price on the plan to replace in the subscription.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for the
	// replacement price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideParam] `json:"discounts"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The new quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// New subscription price request body params.
	Price param.Field[SubscriptionNewParamsReplacePricesPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r SubscriptionNewParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New subscription price request body params.
type SubscriptionNewParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceModelType] `json:"model_type,required"`
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

func (r SubscriptionNewParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// New subscription price request body params.
//
// Satisfied by [NewSubscriptionUnitPriceParam], [NewSubscriptionTieredPriceParam],
// [NewSubscriptionBulkPriceParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice],
// [NewSubscriptionPackagePriceParam], [NewSubscriptionMatrixPriceParam],
// [NewSubscriptionThresholdTotalAmountPriceParam],
// [NewSubscriptionTieredPackagePriceParam],
// [NewSubscriptionTieredWithMinimumPriceParam],
// [NewSubscriptionGroupedTieredPriceParam],
// [NewSubscriptionTieredPackageWithMinimumPriceParam],
// [NewSubscriptionPackageWithAllocationPriceParam],
// [NewSubscriptionUnitWithPercentPriceParam],
// [NewSubscriptionMatrixWithAllocationPriceParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice],
// [NewSubscriptionUnitWithProrationPriceParam],
// [NewSubscriptionGroupedAllocationPriceParam],
// [NewSubscriptionBulkWithProrationPriceParam],
// [NewSubscriptionGroupedWithProratedMinimumPriceParam],
// [NewSubscriptionGroupedWithMeteredMinimumPriceParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice],
// [NewSubscriptionMatrixWithDisplayNamePriceParam],
// [NewSubscriptionGroupedTieredPackagePriceParam],
// [NewSubscriptionMaxGroupTieredPackagePriceParam],
// [NewSubscriptionScalableMatrixWithUnitPricingPriceParam],
// [NewSubscriptionScalableMatrixWithTieredPricingPriceParam],
// [NewSubscriptionCumulativeGroupedBulkPriceParam],
// [NewSubscriptionMinimumCompositePriceParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePrice],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPrice],
// [SubscriptionNewParamsReplacePricesPrice].
type SubscriptionNewParamsReplacePricesPriceUnion interface {
	implementsSubscriptionNewParamsReplacePricesPriceUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                             `json:"unit_config"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig].
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                               `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                 `json:"unit_config"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig].
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelTypePercent SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType = "percent"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                            `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                              `json:"unit_config"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig].
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPrice) implementsSubscriptionNewParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceAnnual     SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "semi_annual"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceMonthly    SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "monthly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly  SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "quarterly"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceOneTime    SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "one_time"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceCustom     SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "custom"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceMonthly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceOneTime, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig struct {
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

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType = "event_output"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                         `json:"unit_config"`
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig].
type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion()
}

type SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType string

const (
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit   SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit, SubscriptionNewParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type SubscriptionNewParamsReplacePricesPriceModelType string

const (
	SubscriptionNewParamsReplacePricesPriceModelTypeUnit                            SubscriptionNewParamsReplacePricesPriceModelType = "unit"
	SubscriptionNewParamsReplacePricesPriceModelTypeTiered                          SubscriptionNewParamsReplacePricesPriceModelType = "tiered"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulk                            SubscriptionNewParamsReplacePricesPriceModelType = "bulk"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithFilters                 SubscriptionNewParamsReplacePricesPriceModelType = "bulk_with_filters"
	SubscriptionNewParamsReplacePricesPriceModelTypePackage                         SubscriptionNewParamsReplacePricesPriceModelType = "package"
	SubscriptionNewParamsReplacePricesPriceModelTypeMatrix                          SubscriptionNewParamsReplacePricesPriceModelType = "matrix"
	SubscriptionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount            SubscriptionNewParamsReplacePricesPriceModelType = "threshold_total_amount"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackage                   SubscriptionNewParamsReplacePricesPriceModelType = "tiered_package"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithMinimum               SubscriptionNewParamsReplacePricesPriceModelType = "tiered_with_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedTiered                   SubscriptionNewParamsReplacePricesPriceModelType = "grouped_tiered"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        SubscriptionNewParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypePackageWithAllocation           SubscriptionNewParamsReplacePricesPriceModelType = "package_with_allocation"
	SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithPercent                 SubscriptionNewParamsReplacePricesPriceModelType = "unit_with_percent"
	SubscriptionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation            SubscriptionNewParamsReplacePricesPriceModelType = "matrix_with_allocation"
	SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithProration             SubscriptionNewParamsReplacePricesPriceModelType = "tiered_with_proration"
	SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithProration               SubscriptionNewParamsReplacePricesPriceModelType = "unit_with_proration"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedAllocation               SubscriptionNewParamsReplacePricesPriceModelType = "grouped_allocation"
	SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithProration               SubscriptionNewParamsReplacePricesPriceModelType = "bulk_with_proration"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      SubscriptionNewParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       SubscriptionNewParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     SubscriptionNewParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	SubscriptionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName           SubscriptionNewParamsReplacePricesPriceModelType = "matrix_with_display_name"
	SubscriptionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage            SubscriptionNewParamsReplacePricesPriceModelType = "grouped_tiered_package"
	SubscriptionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           SubscriptionNewParamsReplacePricesPriceModelType = "max_group_tiered_package"
	SubscriptionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   SubscriptionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	SubscriptionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing SubscriptionNewParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	SubscriptionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           SubscriptionNewParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	SubscriptionNewParamsReplacePricesPriceModelTypeMinimum                         SubscriptionNewParamsReplacePricesPriceModelType = "minimum"
	SubscriptionNewParamsReplacePricesPriceModelTypePercent                         SubscriptionNewParamsReplacePricesPriceModelType = "percent"
	SubscriptionNewParamsReplacePricesPriceModelTypeEventOutput                     SubscriptionNewParamsReplacePricesPriceModelType = "event_output"
)

func (r SubscriptionNewParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionNewParamsReplacePricesPriceModelTypeUnit, SubscriptionNewParamsReplacePricesPriceModelTypeTiered, SubscriptionNewParamsReplacePricesPriceModelTypeBulk, SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithFilters, SubscriptionNewParamsReplacePricesPriceModelTypePackage, SubscriptionNewParamsReplacePricesPriceModelTypeMatrix, SubscriptionNewParamsReplacePricesPriceModelTypeThresholdTotalAmount, SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackage, SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithMinimum, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedTiered, SubscriptionNewParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, SubscriptionNewParamsReplacePricesPriceModelTypePackageWithAllocation, SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithPercent, SubscriptionNewParamsReplacePricesPriceModelTypeMatrixWithAllocation, SubscriptionNewParamsReplacePricesPriceModelTypeTieredWithProration, SubscriptionNewParamsReplacePricesPriceModelTypeUnitWithProration, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedAllocation, SubscriptionNewParamsReplacePricesPriceModelTypeBulkWithProration, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, SubscriptionNewParamsReplacePricesPriceModelTypeMatrixWithDisplayName, SubscriptionNewParamsReplacePricesPriceModelTypeGroupedTieredPackage, SubscriptionNewParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, SubscriptionNewParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, SubscriptionNewParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, SubscriptionNewParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, SubscriptionNewParamsReplacePricesPriceModelTypeMinimum, SubscriptionNewParamsReplacePricesPriceModelTypePercent, SubscriptionNewParamsReplacePricesPriceModelTypeEventOutput:
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
	ExternalCustomerID param.Field[[]string] `query:"external_customer_id"`
	ExternalPlanID     param.Field[string]   `query:"external_plan_id"`
	// The number of items to fetch. Defaults to 20.
	Limit  param.Field[int64]                        `query:"limit"`
	PlanID param.Field[string]                       `query:"plan_id"`
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
	// If true, ending an in-arrears price interval mid-cycle will defer billing the
	// final line itemuntil the next scheduled invoice. If false, it will be billed on
	// its end date. If not provided, behaviorwill follow account default.
	CanDeferBilling param.Field[bool] `json:"can_defer_billing"`
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
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
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
	FixedFeeQuantityTransitions param.Field[[]SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition] `json:"fixed_fee_quantity_transitions"`
	// The maximum amount that will be billed for this price interval for a given
	// billing period.
	MaximumAmount param.Field[float64] `json:"maximum_amount"`
	// The minimum amount that will be billed for this price interval for a given
	// billing period.
	MinimumAmount param.Field[float64] `json:"minimum_amount"`
	// New floating price request body params.
	Price param.Field[SubscriptionPriceIntervalsParamsAddPriceUnion] `json:"price"`
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

type SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition struct {
	// The date that the fixed fee quantity transition should take effect.
	EffectiveDate param.Field[time.Time] `json:"effective_date,required" format:"date-time"`
	// The quantity of the fixed fee quantity transition.
	Quantity param.Field[int64] `json:"quantity,required"`
}

func (r SubscriptionPriceIntervalsParamsAddFixedFeeQuantityTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New floating price request body params.
type SubscriptionPriceIntervalsParamsAddPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceModelType] `json:"model_type,required"`
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
	PackageConfig                         param.Field[shared.PackageConfigParam] `json:"package_config"`
	PackageWithAllocationConfig           param.Field[interface{}]               `json:"package_with_allocation_config"`
	PercentConfig                         param.Field[interface{}]               `json:"percent_config"`
	ScalableMatrixWithTieredPricingConfig param.Field[interface{}]               `json:"scalable_matrix_with_tiered_pricing_config"`
	ScalableMatrixWithUnitPricingConfig   param.Field[interface{}]               `json:"scalable_matrix_with_unit_pricing_config"`
	ThresholdTotalAmountConfig            param.Field[interface{}]               `json:"threshold_total_amount_config"`
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

func (r SubscriptionPriceIntervalsParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPrice) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// New floating price request body params.
//
// Satisfied by [shared.NewFloatingUnitPriceParam],
// [shared.NewFloatingTieredPriceParam], [shared.NewFloatingBulkPriceParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPrice],
// [shared.NewFloatingPackagePriceParam], [shared.NewFloatingMatrixPriceParam],
// [shared.NewFloatingThresholdTotalAmountPriceParam],
// [shared.NewFloatingTieredPackagePriceParam],
// [shared.NewFloatingTieredWithMinimumPriceParam],
// [shared.NewFloatingGroupedTieredPriceParam],
// [shared.NewFloatingTieredPackageWithMinimumPriceParam],
// [shared.NewFloatingPackageWithAllocationPriceParam],
// [shared.NewFloatingUnitWithPercentPriceParam],
// [shared.NewFloatingMatrixWithAllocationPriceParam],
// [shared.NewFloatingTieredWithProrationPriceParam],
// [shared.NewFloatingUnitWithProrationPriceParam],
// [shared.NewFloatingGroupedAllocationPriceParam],
// [shared.NewFloatingBulkWithProrationPriceParam],
// [shared.NewFloatingGroupedWithProratedMinimumPriceParam],
// [shared.NewFloatingGroupedWithMeteredMinimumPriceParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPrice],
// [shared.NewFloatingMatrixWithDisplayNamePriceParam],
// [shared.NewFloatingGroupedTieredPackagePriceParam],
// [shared.NewFloatingMaxGroupTieredPackagePriceParam],
// [shared.NewFloatingScalableMatrixWithUnitPricingPriceParam],
// [shared.NewFloatingScalableMatrixWithTieredPricingPriceParam],
// [shared.NewFloatingCumulativeGroupedBulkPriceParam],
// [shared.NewFloatingMinimumCompositePriceParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePrice],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPrice],
// [SubscriptionPriceIntervalsParamsAddPrice].
type SubscriptionPriceIntervalsParamsAddPriceUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPrice) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelTypeBulkWithFilters SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                        `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                          `json:"unit_config"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfig) ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfig].
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, SubscriptionPriceIntervalsParamsAddPriceNewFloatingBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPrice) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                      `json:"unit_config"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, SubscriptionPriceIntervalsParamsAddPriceNewFloatingGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePrice) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelTypePercent SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelType = "percent"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                         `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                           `json:"unit_config"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfig) ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfig].
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateTypeUnit   SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateTypeTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateTypeUnit, SubscriptionPriceIntervalsParamsAddPriceNewFloatingPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence] `json:"cadence,required"`
	// An ISO 4217 currency string for which this price is billed in.
	Currency param.Field[string] `json:"currency,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPrice) ImplementsSubscriptionPriceIntervalsParamsAddPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceAnnual     SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceSemiAnnual SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "semi_annual"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceMonthly    SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "monthly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceQuarterly  SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "quarterly"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceOneTime    SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "one_time"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceCustom     SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence = "custom"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceSemiAnnual, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceMonthly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceQuarterly, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceOneTime, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceEventOutputConfig struct {
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

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelTypeEventOutput SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelType = "event_output"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                    `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                      `json:"unit_config"`
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfig) ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfig].
type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigUnion()
}

type SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateType string

const (
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateTypeUnit   SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateTypeTiered SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateTypeUnit, SubscriptionPriceIntervalsParamsAddPriceNewFloatingEventOutputPriceConversionRateConfigConversionRateTypeTiered:
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

// The pricing model type
type SubscriptionPriceIntervalsParamsAddPriceModelType string

const (
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnit                            SubscriptionPriceIntervalsParamsAddPriceModelType = "unit"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTiered                          SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulk                            SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithFilters                 SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk_with_filters"
	SubscriptionPriceIntervalsParamsAddPriceModelTypePackage                         SubscriptionPriceIntervalsParamsAddPriceModelType = "package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrix                          SubscriptionPriceIntervalsParamsAddPriceModelType = "matrix"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeThresholdTotalAmount            SubscriptionPriceIntervalsParamsAddPriceModelType = "threshold_total_amount"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackage                   SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithMinimum               SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_with_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTiered                   SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_tiered"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackageWithMinimum        SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_package_with_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypePackageWithAllocation           SubscriptionPriceIntervalsParamsAddPriceModelType = "package_with_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithPercent                 SubscriptionPriceIntervalsParamsAddPriceModelType = "unit_with_percent"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithAllocation            SubscriptionPriceIntervalsParamsAddPriceModelType = "matrix_with_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithProration             SubscriptionPriceIntervalsParamsAddPriceModelType = "tiered_with_proration"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithProration               SubscriptionPriceIntervalsParamsAddPriceModelType = "unit_with_proration"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedAllocation               SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_allocation"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithProration               SubscriptionPriceIntervalsParamsAddPriceModelType = "bulk_with_proration"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithProratedMinimum      SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMeteredMinimum       SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_with_metered_minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMinMaxThresholds     SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_with_min_max_thresholds"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithDisplayName           SubscriptionPriceIntervalsParamsAddPriceModelType = "matrix_with_display_name"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTieredPackage            SubscriptionPriceIntervalsParamsAddPriceModelType = "grouped_tiered_package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMaxGroupTieredPackage           SubscriptionPriceIntervalsParamsAddPriceModelType = "max_group_tiered_package"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeScalableMatrixWithUnitPricing   SubscriptionPriceIntervalsParamsAddPriceModelType = "scalable_matrix_with_unit_pricing"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeScalableMatrixWithTieredPricing SubscriptionPriceIntervalsParamsAddPriceModelType = "scalable_matrix_with_tiered_pricing"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeCumulativeGroupedBulk           SubscriptionPriceIntervalsParamsAddPriceModelType = "cumulative_grouped_bulk"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeMinimum                         SubscriptionPriceIntervalsParamsAddPriceModelType = "minimum"
	SubscriptionPriceIntervalsParamsAddPriceModelTypePercent                         SubscriptionPriceIntervalsParamsAddPriceModelType = "percent"
	SubscriptionPriceIntervalsParamsAddPriceModelTypeEventOutput                     SubscriptionPriceIntervalsParamsAddPriceModelType = "event_output"
)

func (r SubscriptionPriceIntervalsParamsAddPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddPriceModelTypeUnit, SubscriptionPriceIntervalsParamsAddPriceModelTypeTiered, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulk, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithFilters, SubscriptionPriceIntervalsParamsAddPriceModelTypePackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrix, SubscriptionPriceIntervalsParamsAddPriceModelTypeThresholdTotalAmount, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTiered, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredPackageWithMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypePackageWithAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithPercent, SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeTieredWithProration, SubscriptionPriceIntervalsParamsAddPriceModelTypeUnitWithProration, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedAllocation, SubscriptionPriceIntervalsParamsAddPriceModelTypeBulkWithProration, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithProratedMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMeteredMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedWithMinMaxThresholds, SubscriptionPriceIntervalsParamsAddPriceModelTypeMatrixWithDisplayName, SubscriptionPriceIntervalsParamsAddPriceModelTypeGroupedTieredPackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeMaxGroupTieredPackage, SubscriptionPriceIntervalsParamsAddPriceModelTypeScalableMatrixWithUnitPricing, SubscriptionPriceIntervalsParamsAddPriceModelTypeScalableMatrixWithTieredPricing, SubscriptionPriceIntervalsParamsAddPriceModelTypeCumulativeGroupedBulk, SubscriptionPriceIntervalsParamsAddPriceModelTypeMinimum, SubscriptionPriceIntervalsParamsAddPriceModelTypePercent, SubscriptionPriceIntervalsParamsAddPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionPriceIntervalsParamsAddAdjustment struct {
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription. The adjustment will apply to
	// invoice dates that overlap with this `start_date`. This `start_date` is treated
	// as inclusive for in-advance prices, and exclusive for in-arrears prices.
	StartDate param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion] `json:"start_date,required" format:"date-time"`
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment"`
	// The ID of the adjustment to add to the subscription. Adjustment IDs can be
	// re-used from existing subscriptions or plans, but adjustments associated with
	// coupon redemptions cannot be re-used.
	AdjustmentID param.Field[string] `json:"adjustment_id"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription. The adjustment will apply to
	// invoice dates that overlap with this `end_date`.This `end_date` is treated as
	// exclusive for in-advance prices, and inclusive for in-arrears prices.
	EndDate param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsEndDateUnion] `json:"end_date" format:"date-time"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The start date of the adjustment interval. This is the date that the adjustment
// will start affecting prices on the subscription. The adjustment will apply to
// invoice dates that overlap with this `start_date`. This `start_date` is treated
// as inclusive for in-advance prices, and exclusive for in-arrears prices.
//
// Satisfied by [shared.UnionTime], [shared.BillingCycleRelativeDate].
type SubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsStartDateUnion()
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                 `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                          `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                          `json:"applies_to_price_ids"`
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
	PriceType     param.Field[SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                           `json:"usage_discount"`
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment) ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustment].
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsSubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentUnion()
}

type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAllTrue SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType string

const (
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeUsage          SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType = "usage"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixed          SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeInArrears      SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeUsage, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeFixed, SubscriptionPriceIntervalsParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

// The end date of the adjustment interval. This is the date that the adjustment
// will stop affecting prices on the subscription. The adjustment will apply to
// invoice dates that overlap with this `end_date`.This `end_date` is treated as
// exclusive for in-advance prices, and inclusive for in-arrears prices.
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
	// If true, ending an in-arrears price interval mid-cycle will defer billing the
	// final line itemuntil the next scheduled invoice. If false, it will be billed on
	// its end date. If not provided, behaviorwill follow account default.
	CanDeferBilling param.Field[bool] `json:"can_defer_billing"`
	// The updated end date of this price interval. If not specified, the end date will
	// not be updated.
	EndDate param.Field[SubscriptionPriceIntervalsParamsEditEndDateUnion] `json:"end_date" format:"date-time"`
	// An additional filter to apply to usage queries. This filter must be expressed as
	// a boolean
	// [computed property](/extensibility/advanced-metrics#computed-properties). If
	// null, usage queries will not include any additional filter.
	Filter param.Field[string] `json:"filter"`
	// A list of fixed fee quantity transitions to use for this price interval. Note
	// that this list will overwrite all existing fixed fee quantity transitions on the
	// price interval.
	FixedFeeQuantityTransitions param.Field[[]SubscriptionPriceIntervalsParamsEditFixedFeeQuantityTransition] `json:"fixed_fee_quantity_transitions"`
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

// The updated end date of this price interval. If not specified, the end date will
// not be updated.
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
	// The updated end date of this adjustment interval. If not specified, the end date
	// will not be updated.
	EndDate param.Field[SubscriptionPriceIntervalsParamsEditAdjustmentsEndDateUnion] `json:"end_date" format:"date-time"`
	// The updated start date of this adjustment interval. If not specified, the start
	// date will not be updated.
	StartDate param.Field[SubscriptionPriceIntervalsParamsEditAdjustmentsStartDateUnion] `json:"start_date" format:"date-time"`
}

func (r SubscriptionPriceIntervalsParamsEditAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The updated end date of this adjustment interval. If not specified, the end date
// will not be updated.
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

type SubscriptionRedeemCouponParams struct {
	ChangeOption param.Field[SubscriptionRedeemCouponParamsChangeOption] `json:"change_option,required"`
	// If false, this request will fail if it would void an issued invoice or create a
	// credit note. Consider using this as a safety mechanism if you do not expect
	// existing invoices to be changed.
	AllowInvoiceCreditOrVoid param.Field[bool] `json:"allow_invoice_credit_or_void"`
	// The date that the coupon discount should take effect. This parameter can only be
	// passed if the `change_option` is `requested_date`.
	ChangeDate param.Field[time.Time] `json:"change_date" format:"date-time"`
	// Coupon ID to be redeemed for this subscription.
	CouponID param.Field[string] `json:"coupon_id"`
	// Redemption code of the coupon to be redeemed for this subscription.
	CouponRedemptionCode param.Field[string] `json:"coupon_redemption_code"`
}

func (r SubscriptionRedeemCouponParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionRedeemCouponParamsChangeOption string

const (
	SubscriptionRedeemCouponParamsChangeOptionRequestedDate         SubscriptionRedeemCouponParamsChangeOption = "requested_date"
	SubscriptionRedeemCouponParamsChangeOptionEndOfSubscriptionTerm SubscriptionRedeemCouponParamsChangeOption = "end_of_subscription_term"
	SubscriptionRedeemCouponParamsChangeOptionImmediate             SubscriptionRedeemCouponParamsChangeOption = "immediate"
)

func (r SubscriptionRedeemCouponParamsChangeOption) IsKnown() bool {
	switch r {
	case SubscriptionRedeemCouponParamsChangeOptionRequestedDate, SubscriptionRedeemCouponParamsChangeOptionEndOfSubscriptionTerm, SubscriptionRedeemCouponParamsChangeOptionImmediate:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParams struct {
	ChangeOption param.Field[SubscriptionSchedulePlanChangeParamsChangeOption] `json:"change_option,required"`
	// Additional adjustments to be added to the subscription. (Only available for
	// accounts that have migrated off of legacy subscription overrides)
	AddAdjustments param.Field[[]SubscriptionSchedulePlanChangeParamsAddAdjustment] `json:"add_adjustments"`
	// Additional prices to be added to the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	AddPrices param.Field[[]SubscriptionSchedulePlanChangeParamsAddPrice] `json:"add_prices"`
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
	BillingCycleAnchorConfiguration param.Field[shared.BillingCycleAnchorConfigurationParam]               `json:"billing_cycle_anchor_configuration"`
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
	RemoveAdjustments param.Field[[]SubscriptionSchedulePlanChangeParamsRemoveAdjustment] `json:"remove_adjustments"`
	// Plan prices to be removed from the subscription. (Only available for accounts
	// that have migrated off of legacy subscription overrides)
	RemovePrices param.Field[[]SubscriptionSchedulePlanChangeParamsRemovePrice] `json:"remove_prices"`
	// Plan adjustments to be replaced with additional adjustments on the subscription.
	// (Only available for accounts that have migrated off of legacy subscription
	// overrides)
	ReplaceAdjustments param.Field[[]SubscriptionSchedulePlanChangeParamsReplaceAdjustment] `json:"replace_adjustments"`
	// Plan prices to be replaced with additional prices on the subscription. (Only
	// available for accounts that have migrated off of legacy subscription overrides)
	ReplacePrices param.Field[[]SubscriptionSchedulePlanChangeParamsReplacePrice] `json:"replace_prices"`
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

type SubscriptionSchedulePlanChangeParamsAddAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The end date of the adjustment interval. This is the date that the adjustment
	// will stop affecting prices on the subscription.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The phase to add this adjustment to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// The start date of the adjustment interval. This is the date that the adjustment
	// will start affecting prices on the subscription. If null, the adjustment will
	// start when the phase or subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r SubscriptionSchedulePlanChangeParamsAddAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustment struct {
	AdjustmentType param.Field[SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                     `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                              `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                              `json:"applies_to_price_ids"`
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
	PriceType     param.Field[SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                               `json:"usage_discount"`
}

func (r SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustment) ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustment].
type SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentUnion()
}

type SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAll bool

const (
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAllTrue SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAll = true
)

func (r SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType string

const (
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeUsage          SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType = "usage"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixed          SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType = "fixed"
	SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeInArrears      SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeUsage, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixedInAdvance, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixedInArrears, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeFixed, SubscriptionSchedulePlanChangeParamsAddAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPrice struct {
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for this
	// price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideParam] `json:"discounts"`
	// The end date of the price interval. This is the date that the price will stop
	// billing on the subscription. If null, billing will end when the phase or
	// subscription ends.
	EndDate param.Field[time.Time] `json:"end_date" format:"date-time"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// this price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// this price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// The phase to add this price to.
	PlanPhaseOrder param.Field[int64] `json:"plan_phase_order"`
	// New subscription price request body params.
	Price param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
	// The start date of the price interval. This is the date that the price will start
	// billing on the subscription. If null, billing will start when the phase or
	// subscription starts.
	StartDate param.Field[time.Time] `json:"start_date" format:"date-time"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New subscription price request body params.
type SubscriptionSchedulePlanChangeParamsAddPricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType] `json:"model_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// New subscription price request body params.
//
// Satisfied by [NewSubscriptionUnitPriceParam], [NewSubscriptionTieredPriceParam],
// [NewSubscriptionBulkPriceParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice],
// [NewSubscriptionPackagePriceParam], [NewSubscriptionMatrixPriceParam],
// [NewSubscriptionThresholdTotalAmountPriceParam],
// [NewSubscriptionTieredPackagePriceParam],
// [NewSubscriptionTieredWithMinimumPriceParam],
// [NewSubscriptionGroupedTieredPriceParam],
// [NewSubscriptionTieredPackageWithMinimumPriceParam],
// [NewSubscriptionPackageWithAllocationPriceParam],
// [NewSubscriptionUnitWithPercentPriceParam],
// [NewSubscriptionMatrixWithAllocationPriceParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice],
// [NewSubscriptionUnitWithProrationPriceParam],
// [NewSubscriptionGroupedAllocationPriceParam],
// [NewSubscriptionBulkWithProrationPriceParam],
// [NewSubscriptionGroupedWithProratedMinimumPriceParam],
// [NewSubscriptionGroupedWithMeteredMinimumPriceParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice],
// [NewSubscriptionMatrixWithDisplayNamePriceParam],
// [NewSubscriptionGroupedTieredPackagePriceParam],
// [NewSubscriptionMaxGroupTieredPackagePriceParam],
// [NewSubscriptionScalableMatrixWithUnitPricingPriceParam],
// [NewSubscriptionScalableMatrixWithTieredPricingPriceParam],
// [NewSubscriptionCumulativeGroupedBulkPriceParam],
// [NewSubscriptionMinimumCompositePriceParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePrice],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPrice],
// [SubscriptionSchedulePlanChangeParamsAddPricesPrice].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceUnion interface {
	implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                        `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                            `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                    `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelTypePercent SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType = "percent"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                       `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                         `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPrice) implementsSubscriptionSchedulePlanChangeParamsAddPricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig struct {
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

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType = "event_output"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                  `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                    `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsAddPricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnit                            SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "unit"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTiered                          SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "tiered"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulk                            SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "bulk"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulkWithFilters                 SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "bulk_with_filters"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePackage                         SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "package"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrix                          SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "matrix"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeThresholdTotalAmount            SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "threshold_total_amount"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredPackage                   SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "tiered_package"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredWithMinimum               SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "tiered_with_minimum"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedTiered                   SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_tiered"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredPackageWithMinimum        SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "tiered_package_with_minimum"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePackageWithAllocation           SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "package_with_allocation"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnitWithPercent                 SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "unit_with_percent"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrixWithAllocation            SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "matrix_with_allocation"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredWithProration             SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "tiered_with_proration"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnitWithProration               SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "unit_with_proration"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedAllocation               SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_allocation"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulkWithProration               SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "bulk_with_proration"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithProratedMinimum      SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum       SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_with_metered_minimum"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds     SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_with_min_max_thresholds"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrixWithDisplayName           SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "matrix_with_display_name"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedTieredPackage            SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "grouped_tiered_package"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMaxGroupTieredPackage           SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "max_group_tiered_package"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing   SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "scalable_matrix_with_unit_pricing"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeCumulativeGroupedBulk           SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "cumulative_grouped_bulk"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMinimum                         SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "minimum"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePercent                         SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "percent"
	SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeEventOutput                     SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType = "event_output"
)

func (r SubscriptionSchedulePlanChangeParamsAddPricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnit, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTiered, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulk, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulkWithFilters, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePackage, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrix, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeThresholdTotalAmount, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredPackage, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredWithMinimum, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedTiered, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredPackageWithMinimum, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePackageWithAllocation, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnitWithPercent, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrixWithAllocation, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeTieredWithProration, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeUnitWithProration, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedAllocation, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeBulkWithProration, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithMeteredMinimum, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedWithMinMaxThresholds, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMatrixWithDisplayName, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeGroupedTieredPackage, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMaxGroupTieredPackage, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeScalableMatrixWithUnitPricing, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeScalableMatrixWithTieredPricing, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeCumulativeGroupedBulk, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeMinimum, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypePercent, SubscriptionSchedulePlanChangeParamsAddPricesPriceModelTypeEventOutput:
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

type SubscriptionSchedulePlanChangeParamsRemoveAdjustment struct {
	// The id of the adjustment to remove on the subscription.
	AdjustmentID param.Field[string] `json:"adjustment_id,required"`
}

func (r SubscriptionSchedulePlanChangeParamsRemoveAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsRemovePrice struct {
	// The external price id of the price to remove on the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The id of the price to remove on the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r SubscriptionSchedulePlanChangeParamsRemovePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsReplaceAdjustment struct {
	// The definition of a new adjustment to create and add to the subscription.
	Adjustment param.Field[SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion] `json:"adjustment,required"`
	// The id of the adjustment on the plan to replace in the subscription.
	ReplacesAdjustmentID param.Field[string] `json:"replaces_adjustment_id,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of a new adjustment to create and add to the subscription.
type SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustment struct {
	AdjustmentType param.Field[SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType] `json:"adjustment_type,required"`
	AmountDiscount param.Field[string]                                                                         `json:"amount_discount"`
	// If set, the adjustment will apply to every price on the subscription.
	AppliesToAll      param.Field[SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAll] `json:"applies_to_all"`
	AppliesToItemIDs  param.Field[interface{}]                                                                  `json:"applies_to_item_ids"`
	AppliesToPriceIDs param.Field[interface{}]                                                                  `json:"applies_to_price_ids"`
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
	PriceType     param.Field[SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType] `json:"price_type"`
	UsageDiscount param.Field[float64]                                                                   `json:"usage_discount"`
}

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustment) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustment) ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion() {
}

// The definition of a new adjustment to create and add to the subscription.
//
// Satisfied by [shared.NewPercentageDiscountParam],
// [shared.NewUsageDiscountParam], [shared.NewAmountDiscountParam],
// [shared.NewMinimumParam], [shared.NewMaximumParam],
// [SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustment].
type SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentUnion()
}

type SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType string

const (
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum            SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType = "minimum"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum            SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMinimum, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

// If set, the adjustment will apply to every price on the subscription.
type SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAll bool

const (
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAll = true
)

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAll) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentAppliesToAllTrue:
		return true
	}
	return false
}

// If set, only prices of the specified type will have the adjustment applied.
type SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType string

const (
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeUsage          SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType = "usage"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_advance"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType = "fixed_in_arrears"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixed          SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType = "fixed"
	SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears      SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType = "in_arrears"
)

func (r SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeUsage, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInAdvance, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixedInArrears, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeFixed, SubscriptionSchedulePlanChangeParamsReplaceAdjustmentsAdjustmentPriceTypeInArrears:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePrice struct {
	// The id of the price on the plan to replace in the subscription.
	ReplacesPriceID param.Field[string] `json:"replaces_price_id,required"`
	// The definition of a new allocation price to create and add to the subscription.
	AllocationPrice param.Field[shared.NewAllocationPriceParam] `json:"allocation_price"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's discounts for the
	// replacement price.
	//
	// Deprecated: deprecated
	Discounts param.Field[[]DiscountOverrideParam] `json:"discounts"`
	// The external price id of the price to add to the subscription.
	ExternalPriceID param.Field[string] `json:"external_price_id"`
	// The new quantity of the price, if the price is a fixed price.
	FixedPriceQuantity param.Field[float64] `json:"fixed_price_quantity"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's maximum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MaximumAmount param.Field[string] `json:"maximum_amount"`
	// [DEPRECATED] Use add_adjustments instead. The subscription's minimum amount for
	// the replacement price.
	//
	// Deprecated: deprecated
	MinimumAmount param.Field[string] `json:"minimum_amount"`
	// New subscription price request body params.
	Price param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion] `json:"price"`
	// The id of the price to add to the subscription.
	PriceID param.Field[string] `json:"price_id"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// New subscription price request body params.
type SubscriptionSchedulePlanChangeParamsReplacePricesPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType] `json:"model_type,required"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// New subscription price request body params.
//
// Satisfied by [NewSubscriptionUnitPriceParam], [NewSubscriptionTieredPriceParam],
// [NewSubscriptionBulkPriceParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice],
// [NewSubscriptionPackagePriceParam], [NewSubscriptionMatrixPriceParam],
// [NewSubscriptionThresholdTotalAmountPriceParam],
// [NewSubscriptionTieredPackagePriceParam],
// [NewSubscriptionTieredWithMinimumPriceParam],
// [NewSubscriptionGroupedTieredPriceParam],
// [NewSubscriptionTieredPackageWithMinimumPriceParam],
// [NewSubscriptionPackageWithAllocationPriceParam],
// [NewSubscriptionUnitWithPercentPriceParam],
// [NewSubscriptionMatrixWithAllocationPriceParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice],
// [NewSubscriptionUnitWithProrationPriceParam],
// [NewSubscriptionGroupedAllocationPriceParam],
// [NewSubscriptionBulkWithProrationPriceParam],
// [NewSubscriptionGroupedWithProratedMinimumPriceParam],
// [NewSubscriptionGroupedWithMeteredMinimumPriceParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice],
// [NewSubscriptionMatrixWithDisplayNamePriceParam],
// [NewSubscriptionGroupedTieredPackagePriceParam],
// [NewSubscriptionMaxGroupTieredPackagePriceParam],
// [NewSubscriptionScalableMatrixWithUnitPricingPriceParam],
// [NewSubscriptionScalableMatrixWithTieredPricingPriceParam],
// [NewSubscriptionCumulativeGroupedBulkPriceParam],
// [NewSubscriptionMinimumCompositePriceParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePrice],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPrice],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPrice].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion interface {
	implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice struct {
	// Configuration for bulk_with_filters pricing
	BulkWithFiltersConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig] `json:"bulk_with_filters_config,required"`
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// Configuration for bulk_with_filters pricing
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig struct {
	// Property filters to apply (all must match)
	Filters param.Field[[]SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter] `json:"filters,required"`
	// Bulk tiers for rating based on total usage volume
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single property filter
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter struct {
	// Event property key to filter on
	PropertyKey param.Field[string] `json:"property_key,required"`
	// Event property value to match
	PropertyValue param.Field[string] `json:"property_value,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single bulk pricing tier
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier struct {
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
	// The lower bound for this tier
	TierLowerBound param.Field[string] `json:"tier_lower_bound"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceBulkWithFiltersConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType = "bulk_with_filters"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceModelTypeBulkWithFilters:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                          `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                            `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionBulkWithFiltersPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for tiered_with_proration pricing
	TieredWithProrationConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig] `json:"tiered_with_proration_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType = "tiered_with_proration"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceModelTypeTieredWithProration:
		return true
	}
	return false
}

// Configuration for tiered_with_proration pricing
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig struct {
	// Tiers for rating based on total usage quantities into the specified tier with
	// proration
	Tiers param.Field[[]SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier] `json:"tiers,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for a single tiered with proration tier
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier struct {
	// Inclusive tier starting value
	TierLowerBound param.Field[string] `json:"tier_lower_bound,required"`
	// Amount per unit
	UnitAmount param.Field[string] `json:"unit_amount,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceTieredWithProrationConfigTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                              `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionTieredWithProrationPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence] `json:"cadence,required"`
	// Configuration for grouped_with_min_max_thresholds pricing
	GroupedWithMinMaxThresholdsConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig] `json:"grouped_with_min_max_thresholds_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for grouped_with_min_max_thresholds pricing
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig struct {
	// The event property used to group before applying thresholds
	GroupingKey param.Field[string] `json:"grouping_key,required"`
	// The maximum amount to charge each group
	MaximumCharge param.Field[string] `json:"maximum_charge,required"`
	// The minimum amount to charge each group, regardless of usage
	MinimumCharge param.Field[string] `json:"minimum_charge,required"`
	// The base price charged per group
	PerUnitRate param.Field[string] `json:"per_unit_rate,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceGroupedWithMinMaxThresholdsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType = "grouped_with_min_max_thresholds"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceModelTypeGroupedWithMinMaxThresholds:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                                        `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionGroupedWithMinMaxThresholdsPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence] `json:"cadence,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType] `json:"model_type,required"`
	// The name of the price.
	Name param.Field[string] `json:"name,required"`
	// Configuration for percent pricing
	PercentConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig] `json:"percent_config,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelTypePercent SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType = "percent"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceModelTypePercent:
		return true
	}
	return false
}

// Configuration for percent pricing
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig struct {
	// What percent of the component subtotals to charge
	Percent param.Field[float64] `json:"percent,required"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePricePercentConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                           `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                             `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionPercentCompositePriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPrice struct {
	// The cadence to bill for this price on.
	Cadence param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence] `json:"cadence,required"`
	// Configuration for event_output pricing
	EventOutputConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig] `json:"event_output_config,required"`
	// The id of the item the price will be associated with.
	ItemID param.Field[string] `json:"item_id,required"`
	// The pricing model type
	ModelType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType] `json:"model_type,required"`
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
	ConversionRateConfig param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion] `json:"conversion_rate_config"`
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPrice) implementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceUnion() {
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceCadenceCustom:
		return true
	}
	return false
}

// Configuration for event_output pricing
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig struct {
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

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceEventOutputConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType = "event_output"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceModelTypeEventOutput:
		return true
	}
	return false
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig struct {
	ConversionRateType param.Field[SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType] `json:"conversion_rate_type,required"`
	TieredConfig       param.Field[shared.ConversionRateTieredConfigParam]                                                                                      `json:"tiered_config"`
	UnitConfig         param.Field[shared.ConversionRateUnitConfigParam]                                                                                        `json:"unit_config"`
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig) ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion() {
}

// Satisfied by [shared.UnitConversionRateConfigParam],
// [shared.TieredConversionRateConfigParam],
// [SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfig].
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion interface {
	ImplementsSubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigUnion()
}

type SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit   SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType = "tiered"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceNewSubscriptionEventOutputPriceConversionRateConfigConversionRateTypeTiered:
		return true
	}
	return false
}

// The cadence to bill for this price on.
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceAnnual     SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceSemiAnnual SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "semi_annual"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceMonthly    SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "monthly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceQuarterly  SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "quarterly"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceOneTime    SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "one_time"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceCustom     SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence = "custom"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadence) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceSemiAnnual, SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceMonthly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceQuarterly, SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceOneTime, SubscriptionSchedulePlanChangeParamsReplacePricesPriceCadenceCustom:
		return true
	}
	return false
}

// The pricing model type
type SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType string

const (
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnit                            SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "unit"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTiered                          SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "tiered"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulk                            SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "bulk"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulkWithFilters                 SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "bulk_with_filters"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePackage                         SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "package"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrix                          SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "matrix"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeThresholdTotalAmount            SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "threshold_total_amount"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredPackage                   SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "tiered_package"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredWithMinimum               SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "tiered_with_minimum"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedTiered                   SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_tiered"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredPackageWithMinimum        SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "tiered_package_with_minimum"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePackageWithAllocation           SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "package_with_allocation"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnitWithPercent                 SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "unit_with_percent"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrixWithAllocation            SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "matrix_with_allocation"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredWithProration             SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "tiered_with_proration"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnitWithProration               SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "unit_with_proration"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedAllocation               SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_allocation"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulkWithProration               SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "bulk_with_proration"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum      SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_with_prorated_minimum"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum       SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_with_metered_minimum"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds     SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_with_min_max_thresholds"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrixWithDisplayName           SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "matrix_with_display_name"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedTieredPackage            SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "grouped_tiered_package"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMaxGroupTieredPackage           SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "max_group_tiered_package"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing   SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "scalable_matrix_with_unit_pricing"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "scalable_matrix_with_tiered_pricing"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeCumulativeGroupedBulk           SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "cumulative_grouped_bulk"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMinimum                         SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "minimum"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePercent                         SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "percent"
	SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeEventOutput                     SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType = "event_output"
)

func (r SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelType) IsKnown() bool {
	switch r {
	case SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnit, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTiered, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulk, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulkWithFilters, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePackage, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrix, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeThresholdTotalAmount, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredPackage, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredWithMinimum, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedTiered, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredPackageWithMinimum, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePackageWithAllocation, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnitWithPercent, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrixWithAllocation, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeTieredWithProration, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeUnitWithProration, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedAllocation, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeBulkWithProration, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithProratedMinimum, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithMeteredMinimum, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedWithMinMaxThresholds, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMatrixWithDisplayName, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeGroupedTieredPackage, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMaxGroupTieredPackage, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeScalableMatrixWithUnitPricing, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeScalableMatrixWithTieredPricing, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeCumulativeGroupedBulk, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeMinimum, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypePercent, SubscriptionSchedulePlanChangeParamsReplacePricesPriceModelTypeEventOutput:
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
	// customer's timezone. If this parameter is not passed in, the quantity change is
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
