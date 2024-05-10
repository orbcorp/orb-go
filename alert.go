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
	"github.com/orbcorp/orb-go/internal/pagination"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// AlertService contains methods and other services that help with interacting with
// the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlertService] method instead.
type AlertService struct {
	Options []option.RequestOption
}

// NewAlertService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlertService(opts ...option.RequestOption) (r *AlertService) {
	r = &AlertService{}
	r.Options = opts
	return
}

// This endpoint retrieves an alert by its ID.
func (r *AlertService) Get(ctx context.Context, alertID string, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/%s", alertID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint returns a list of alerts within Orb.
//
// The request must specify one of `customer_id`, `external_customer_id`, or
// `subscription_id`.
//
// If querying by subscripion_id, the endpoint will return the subscription level
// alerts as well as the plan level alerts associated with the subscription.
//
// The list of alerts is ordered starting from the most recently created alert.
// This endpoint follows Orb's
// [standardized pagination format](../reference/pagination).
func (r *AlertService) List(ctx context.Context, query AlertListParams, opts ...option.RequestOption) (res *pagination.Page[Alert], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "alerts"
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

// This endpoint returns a list of alerts within Orb.
//
// The request must specify one of `customer_id`, `external_customer_id`, or
// `subscription_id`.
//
// If querying by subscripion_id, the endpoint will return the subscription level
// alerts as well as the plan level alerts associated with the subscription.
//
// The list of alerts is ordered starting from the most recently created alert.
// This endpoint follows Orb's
// [standardized pagination format](../reference/pagination).
func (r *AlertService) ListAutoPaging(ctx context.Context, query AlertListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Alert] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint creates a new alert to monitor a customer's credit balance. There
// are three types of alerts that can be scoped to customers:
// `credit_balance_depleted`, `credit_balance_dropped`, and
// `credit_balance_recovered`. Customers can have a maximum of one of each type of
// alert per
// [credit balance currency](https://docs.withorb.com/guides/product-catalog/prepurchase).
// `credit_balance_dropped` alerts require a list of thresholds to be provided
// while `credit_balance_depleted` and `credit_balance_recovered` alerts do not
// require thresholds.
func (r *AlertService) NewForCustomer(ctx context.Context, customerID string, body AlertNewForCustomerParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/customer_id/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint creates a new alert to monitor a customer's credit balance. There
// are three types of alerts that can be scoped to customers:
// `credit_balance_depleted`, `credit_balance_dropped`, and
// `credit_balance_recovered`. Customers can have a maximum of one of each type of
// alert per
// [credit balance currency](https://docs.withorb.com/guides/product-catalog/prepurchase).
// `credit_balance_dropped` alerts require a list of thresholds to be provided
// while `credit_balance_depleted` and `credit_balance_recovered` alerts do not
// require thresholds.
func (r *AlertService) NewForExternalCustomer(ctx context.Context, externalCustomerID string, body AlertNewForExternalCustomerParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/external_customer_id/%s", externalCustomerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint is used to create alerts at the subscription level.
//
// Subscription level alerts can be one of two types: `usage_exceeded` or
// `cost_exceeded`. A `usage_exceeded` alert is scoped to a particular metric and
// is triggered when the usage of that metric exceeds predefined thresholds during
// the current billing cycle. A `cost_exceeded` alert is triggered when the total
// amount due during the current billing cycle surpasses predefined thresholds.
// `cost_exceeded` alerts do not include burndown of pre-purchase credits. Each
// subscription can have one `cost_exceeded` alert and one `usage_exceeded` alert
// per metric that is a part of the subscription. Alerts are triggered based on
// usage or cost conditions met during the current billing cycle.
func (r *AlertService) NewForSubscription(ctx context.Context, subscriptionID string, body AlertNewForSubscriptionParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/subscription_id/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint can be used to disable an alert.
func (r *AlertService) Disable(ctx context.Context, alertConfigurationID string, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/%s/disable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// This endpoint can be used to enable an alert.
func (r *AlertService) Enable(ctx context.Context, alertConfigurationID string, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/%s/enable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// [Alerts within Orb](https://docs.withorb.com/guides/product-catalog/configuring-alerts)
// monitor spending, usage, or credit balance and trigger webhooks when a threshold
// is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
//
// | Scope        | Monitors                       | Vaild Alert Types                                                                   |
// | ------------ | ------------------------------ | ----------------------------------------------------------------------------------- |
// | Customer     | A customer's credit balance    | `credit_balance_depleted`, `credit_balance_recovered`, and `credit_balance_dropped` |
// | Subscription | A subscription's usage or cost | `usage_exceeded` and `cost_exceeded`                                                |
type Alert struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency,required,nullable"`
	// The customer the alert applies to.
	Customer map[string]string `json:"customer,required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled,required"`
	// The metric the alert applies to.
	Metric map[string]string `json:"metric,required,nullable"`
	// The plan the alert applies to.
	Plan map[string]string `json:"plan,required,nullable"`
	// The subscription the alert applies to.
	Subscription map[string]string `json:"subscription,required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []AlertThreshold `json:"thresholds,required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type AlertType `json:"type,required"`
	JSON alertJSON `json:"-"`
}

// alertJSON contains the JSON metadata for the struct [Alert]
type alertJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Currency     apijson.Field
	Customer     apijson.Field
	Enabled      apijson.Field
	Metric       apijson.Field
	Plan         apijson.Field
	Subscription apijson.Field
	Thresholds   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Alert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type AlertThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value float64            `json:"value,required"`
	JSON  alertThresholdJSON `json:"-"`
}

// alertThresholdJSON contains the JSON metadata for the struct [AlertThreshold]
type alertThresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertThresholdJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type AlertType string

const (
	AlertTypeUsageExceeded          AlertType = "usage_exceeded"
	AlertTypeCostExceeded           AlertType = "cost_exceeded"
	AlertTypeCreditBalanceDepleted  AlertType = "credit_balance_depleted"
	AlertTypeCreditBalanceDropped   AlertType = "credit_balance_dropped"
	AlertTypeCreditBalanceRecovered AlertType = "credit_balance_recovered"
)

func (r AlertType) IsKnown() bool {
	switch r {
	case AlertTypeUsageExceeded, AlertTypeCostExceeded, AlertTypeCreditBalanceDepleted, AlertTypeCreditBalanceDropped, AlertTypeCreditBalanceRecovered:
		return true
	}
	return false
}

type AlertListParams struct {
	CreatedAtGt  param.Field[time.Time] `query:"created_at[gt]" format:"date-time"`
	CreatedAtGte param.Field[time.Time] `query:"created_at[gte]" format:"date-time"`
	CreatedAtLt  param.Field[time.Time] `query:"created_at[lt]" format:"date-time"`
	CreatedAtLte param.Field[time.Time] `query:"created_at[lte]" format:"date-time"`
	// Cursor for pagination. This can be populated by the `next_cursor` value returned
	// from the initial request.
	Cursor param.Field[string] `query:"cursor"`
	// Fetch alerts scoped to this customer_id
	CustomerID param.Field[string] `query:"customer_id"`
	// Fetch alerts scoped to this external_customer_id
	ExternalCustomerID param.Field[string] `query:"external_customer_id"`
	// The number of items to fetch. Defaults to 20.
	Limit param.Field[int64] `query:"limit"`
	// Fetch alerts scoped to this subscription_id
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [AlertListParams]'s query parameters as `url.Values`.
func (r AlertListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlertNewForCustomerParams struct {
	// The case sensitive currency or custom pricing unit to use for this alert.
	Currency param.Field[string] `json:"currency,required"`
	// The thresholds that define the values at which the alert will be triggered.
	Type param.Field[string] `json:"type,required"`
	// The thresholds for the alert.
	Thresholds param.Field[[]AlertNewForCustomerParamsThreshold] `json:"thresholds"`
}

func (r AlertNewForCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type AlertNewForCustomerParamsThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value param.Field[float64] `json:"value,required"`
}

func (r AlertNewForCustomerParamsThreshold) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertNewForExternalCustomerParams struct {
	// The case sensitive currency or custom pricing unit to use for this alert.
	Currency param.Field[string] `json:"currency,required"`
	// The thresholds that define the values at which the alert will be triggered.
	Type param.Field[string] `json:"type,required"`
	// The thresholds for the alert.
	Thresholds param.Field[[]AlertNewForExternalCustomerParamsThreshold] `json:"thresholds"`
}

func (r AlertNewForExternalCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type AlertNewForExternalCustomerParamsThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value param.Field[float64] `json:"value,required"`
}

func (r AlertNewForExternalCustomerParamsThreshold) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertNewForSubscriptionParams struct {
	// The thresholds for the alert.
	Thresholds param.Field[[]AlertNewForSubscriptionParamsThreshold] `json:"thresholds,required"`
	// The thresholds that define the values at which the alert will be triggered.
	Type param.Field[string] `json:"type,required"`
	// The metric to track usage for.
	MetricID param.Field[string] `json:"metric_id"`
}

func (r AlertNewForSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type AlertNewForSubscriptionParamsThreshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value param.Field[float64] `json:"value,required"`
}

func (r AlertNewForSubscriptionParamsThreshold) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
