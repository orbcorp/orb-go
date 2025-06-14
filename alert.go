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
	if alertID == "" {
		err = errors.New("missing required alert_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/%s", alertID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This endpoint updates the thresholds of an alert.
func (r *AlertService) Update(ctx context.Context, alertConfigurationID string, body AlertUpdateParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	if alertConfigurationID == "" {
		err = errors.New("missing required alert_configuration_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/%s", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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
// [standardized pagination format](/api-reference/pagination).
func (r *AlertService) List(ctx context.Context, query AlertListParams, opts ...option.RequestOption) (res *pagination.Page[Alert], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
// [standardized pagination format](/api-reference/pagination).
func (r *AlertService) ListAutoPaging(ctx context.Context, query AlertListParams, opts ...option.RequestOption) *pagination.PageAutoPager[Alert] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

// This endpoint creates a new alert to monitor a customer's credit balance. There
// are three types of alerts that can be scoped to customers:
// `credit_balance_depleted`, `credit_balance_dropped`, and
// `credit_balance_recovered`. Customers can have a maximum of one of each type of
// alert per [credit balance currency](/product-catalog/prepurchase).
// `credit_balance_dropped` alerts require a list of thresholds to be provided
// while `credit_balance_depleted` and `credit_balance_recovered` alerts do not
// require thresholds.
func (r *AlertService) NewForCustomer(ctx context.Context, customerID string, body AlertNewForCustomerParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	if customerID == "" {
		err = errors.New("missing required customer_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/customer_id/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint creates a new alert to monitor a customer's credit balance. There
// are three types of alerts that can be scoped to customers:
// `credit_balance_depleted`, `credit_balance_dropped`, and
// `credit_balance_recovered`. Customers can have a maximum of one of each type of
// alert per [credit balance currency](/product-catalog/prepurchase).
// `credit_balance_dropped` alerts require a list of thresholds to be provided
// while `credit_balance_depleted` and `credit_balance_recovered` alerts do not
// require thresholds.
func (r *AlertService) NewForExternalCustomer(ctx context.Context, externalCustomerID string, body AlertNewForExternalCustomerParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	if externalCustomerID == "" {
		err = errors.New("missing required external_customer_id parameter")
		return
	}
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
	if subscriptionID == "" {
		err = errors.New("missing required subscription_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/subscription_id/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to disable an alert. To disable a plan-level alert for
// a specific subscription, you must include the `subscription_id`. The
// `subscription_id` is not required for customer or subscription level alerts.
func (r *AlertService) Disable(ctx context.Context, alertConfigurationID string, body AlertDisableParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	if alertConfigurationID == "" {
		err = errors.New("missing required alert_configuration_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/%s/disable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint allows you to enable an alert. To enable a plan-level alert for a
// specific subscription, you must include the `subscription_id`. The
// `subscription_id` is not required for customer or subscription level alerts.
func (r *AlertService) Enable(ctx context.Context, alertConfigurationID string, body AlertEnableParams, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	if alertConfigurationID == "" {
		err = errors.New("missing required alert_configuration_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/%s/enable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [Alerts within Orb](/product-catalog/configuring-alerts) monitor spending,
// usage, or credit balance and trigger webhooks when a threshold is exceeded.
//
// Alerts created through the API can be scoped to either customers or
// subscriptions.
type Alert struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the currency the credit balance or invoice cost is denominated in.
	Currency string `json:"currency,required,nullable"`
	// The customer the alert applies to.
	Customer shared.CustomerMinified `json:"customer,required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool `json:"enabled,required"`
	// The metric the alert applies to.
	Metric AlertMetric `json:"metric,required,nullable"`
	// The plan the alert applies to.
	Plan AlertPlan `json:"plan,required,nullable"`
	// The subscription the alert applies to.
	Subscription shared.SubscriptionMinified `json:"subscription,required,nullable"`
	// The thresholds that define the conditions under which the alert will be
	// triggered.
	Thresholds []Threshold `json:"thresholds,required,nullable"`
	// The type of alert. This must be a valid alert type.
	Type AlertType `json:"type,required"`
	// The current status of the alert. This field is only present for credit balance
	// alerts.
	BalanceAlertStatus []AlertBalanceAlertStatus `json:"balance_alert_status,nullable"`
	JSON               alertJSON                 `json:"-"`
}

// alertJSON contains the JSON metadata for the struct [Alert]
type alertJSON struct {
	ID                 apijson.Field
	CreatedAt          apijson.Field
	Currency           apijson.Field
	Customer           apijson.Field
	Enabled            apijson.Field
	Metric             apijson.Field
	Plan               apijson.Field
	Subscription       apijson.Field
	Thresholds         apijson.Field
	Type               apijson.Field
	BalanceAlertStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Alert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertJSON) RawJSON() string {
	return r.raw
}

// The metric the alert applies to.
type AlertMetric struct {
	ID   string          `json:"id,required"`
	JSON alertMetricJSON `json:"-"`
}

// alertMetricJSON contains the JSON metadata for the struct [AlertMetric]
type alertMetricJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertMetric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertMetricJSON) RawJSON() string {
	return r.raw
}

// The plan the alert applies to.
type AlertPlan struct {
	ID string `json:"id,required,nullable"`
	// An optional user-defined ID for this plan resource, used throughout the system
	// as an alias for this Plan. Use this field to identify a plan by an existing
	// identifier in your system.
	ExternalPlanID string        `json:"external_plan_id,required,nullable"`
	Name           string        `json:"name,required,nullable"`
	PlanVersion    string        `json:"plan_version,required"`
	JSON           alertPlanJSON `json:"-"`
}

// alertPlanJSON contains the JSON metadata for the struct [AlertPlan]
type alertPlanJSON struct {
	ID             apijson.Field
	ExternalPlanID apijson.Field
	Name           apijson.Field
	PlanVersion    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AlertPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertPlanJSON) RawJSON() string {
	return r.raw
}

// The type of alert. This must be a valid alert type.
type AlertType string

const (
	AlertTypeCreditBalanceDepleted  AlertType = "credit_balance_depleted"
	AlertTypeCreditBalanceDropped   AlertType = "credit_balance_dropped"
	AlertTypeCreditBalanceRecovered AlertType = "credit_balance_recovered"
	AlertTypeUsageExceeded          AlertType = "usage_exceeded"
	AlertTypeCostExceeded           AlertType = "cost_exceeded"
)

func (r AlertType) IsKnown() bool {
	switch r {
	case AlertTypeCreditBalanceDepleted, AlertTypeCreditBalanceDropped, AlertTypeCreditBalanceRecovered, AlertTypeUsageExceeded, AlertTypeCostExceeded:
		return true
	}
	return false
}

// Alert status is used to determine if an alert is currently in-alert or not.
type AlertBalanceAlertStatus struct {
	// Whether the alert is currently in-alert or not.
	InAlert bool `json:"in_alert,required"`
	// The value of the threshold that defines the alert status.
	ThresholdValue float64                     `json:"threshold_value,required"`
	JSON           alertBalanceAlertStatusJSON `json:"-"`
}

// alertBalanceAlertStatusJSON contains the JSON metadata for the struct
// [AlertBalanceAlertStatus]
type alertBalanceAlertStatusJSON struct {
	InAlert        apijson.Field
	ThresholdValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AlertBalanceAlertStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alertBalanceAlertStatusJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type Threshold struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value float64       `json:"value,required"`
	JSON  thresholdJSON `json:"-"`
}

// thresholdJSON contains the JSON metadata for the struct [Threshold]
type thresholdJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Threshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r thresholdJSON) RawJSON() string {
	return r.raw
}

// Thresholds are used to define the conditions under which an alert will be
// triggered.
type ThresholdParam struct {
	// The value at which an alert will fire. For credit balance alerts, the alert will
	// fire at or below this value. For usage and cost alerts, the alert will fire at
	// or above this value.
	Value param.Field[float64] `json:"value,required"`
}

func (r ThresholdParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertUpdateParams struct {
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]ThresholdParam] `json:"thresholds,required"`
}

func (r AlertUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// The type of alert to create. This must be a valid alert type.
	Type param.Field[AlertNewForCustomerParamsType] `json:"type,required"`
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]ThresholdParam] `json:"thresholds"`
}

func (r AlertNewForCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of alert to create. This must be a valid alert type.
type AlertNewForCustomerParamsType string

const (
	AlertNewForCustomerParamsTypeCreditBalanceDepleted  AlertNewForCustomerParamsType = "credit_balance_depleted"
	AlertNewForCustomerParamsTypeCreditBalanceDropped   AlertNewForCustomerParamsType = "credit_balance_dropped"
	AlertNewForCustomerParamsTypeCreditBalanceRecovered AlertNewForCustomerParamsType = "credit_balance_recovered"
)

func (r AlertNewForCustomerParamsType) IsKnown() bool {
	switch r {
	case AlertNewForCustomerParamsTypeCreditBalanceDepleted, AlertNewForCustomerParamsTypeCreditBalanceDropped, AlertNewForCustomerParamsTypeCreditBalanceRecovered:
		return true
	}
	return false
}

type AlertNewForExternalCustomerParams struct {
	// The case sensitive currency or custom pricing unit to use for this alert.
	Currency param.Field[string] `json:"currency,required"`
	// The type of alert to create. This must be a valid alert type.
	Type param.Field[AlertNewForExternalCustomerParamsType] `json:"type,required"`
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]ThresholdParam] `json:"thresholds"`
}

func (r AlertNewForExternalCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of alert to create. This must be a valid alert type.
type AlertNewForExternalCustomerParamsType string

const (
	AlertNewForExternalCustomerParamsTypeCreditBalanceDepleted  AlertNewForExternalCustomerParamsType = "credit_balance_depleted"
	AlertNewForExternalCustomerParamsTypeCreditBalanceDropped   AlertNewForExternalCustomerParamsType = "credit_balance_dropped"
	AlertNewForExternalCustomerParamsTypeCreditBalanceRecovered AlertNewForExternalCustomerParamsType = "credit_balance_recovered"
)

func (r AlertNewForExternalCustomerParamsType) IsKnown() bool {
	switch r {
	case AlertNewForExternalCustomerParamsTypeCreditBalanceDepleted, AlertNewForExternalCustomerParamsTypeCreditBalanceDropped, AlertNewForExternalCustomerParamsTypeCreditBalanceRecovered:
		return true
	}
	return false
}

type AlertNewForSubscriptionParams struct {
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]ThresholdParam] `json:"thresholds,required"`
	// The type of alert to create. This must be a valid alert type.
	Type param.Field[AlertNewForSubscriptionParamsType] `json:"type,required"`
	// The metric to track usage for.
	MetricID param.Field[string] `json:"metric_id"`
}

func (r AlertNewForSubscriptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of alert to create. This must be a valid alert type.
type AlertNewForSubscriptionParamsType string

const (
	AlertNewForSubscriptionParamsTypeUsageExceeded AlertNewForSubscriptionParamsType = "usage_exceeded"
	AlertNewForSubscriptionParamsTypeCostExceeded  AlertNewForSubscriptionParamsType = "cost_exceeded"
)

func (r AlertNewForSubscriptionParamsType) IsKnown() bool {
	switch r {
	case AlertNewForSubscriptionParamsTypeUsageExceeded, AlertNewForSubscriptionParamsTypeCostExceeded:
		return true
	}
	return false
}

type AlertDisableParams struct {
	// Used to update the status of a plan alert scoped to this subscription_id
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [AlertDisableParams]'s query parameters as `url.Values`.
func (r AlertDisableParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlertEnableParams struct {
	// Used to update the status of a plan alert scoped to this subscription_id
	SubscriptionID param.Field[string] `query:"subscription_id"`
}

// URLQuery serializes [AlertEnableParams]'s query parameters as `url.Values`.
func (r AlertEnableParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
