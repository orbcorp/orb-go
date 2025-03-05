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
func (r *AlertService) Get(ctx context.Context, alertID string, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) Update(ctx context.Context, alertConfigurationID string, body AlertUpdateParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) List(ctx context.Context, query AlertListParams, opts ...option.RequestOption) (res *pagination.Page[shared.AlertModel], err error) {
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
func (r *AlertService) ListAutoPaging(ctx context.Context, query AlertListParams, opts ...option.RequestOption) *pagination.PageAutoPager[shared.AlertModel] {
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
func (r *AlertService) NewForCustomer(ctx context.Context, customerID string, body AlertNewForCustomerParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) NewForExternalCustomer(ctx context.Context, externalCustomerID string, body AlertNewForExternalCustomerParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) NewForSubscription(ctx context.Context, subscriptionID string, body AlertNewForSubscriptionParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) Disable(ctx context.Context, alertConfigurationID string, body AlertDisableParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
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
func (r *AlertService) Enable(ctx context.Context, alertConfigurationID string, body AlertEnableParams, opts ...option.RequestOption) (res *shared.AlertModel, err error) {
	opts = append(r.Options[:], opts...)
	if alertConfigurationID == "" {
		err = errors.New("missing required alert_configuration_id parameter")
		return
	}
	path := fmt.Sprintf("alerts/%s/enable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AlertUpdateParams struct {
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]shared.ThresholdModelParam] `json:"thresholds,required"`
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
	Thresholds param.Field[[]shared.ThresholdModelParam] `json:"thresholds"`
}

func (r AlertNewForCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of alert to create. This must be a valid alert type.
type AlertNewForCustomerParamsType string

const (
	AlertNewForCustomerParamsTypeUsageExceeded          AlertNewForCustomerParamsType = "usage_exceeded"
	AlertNewForCustomerParamsTypeCostExceeded           AlertNewForCustomerParamsType = "cost_exceeded"
	AlertNewForCustomerParamsTypeCreditBalanceDepleted  AlertNewForCustomerParamsType = "credit_balance_depleted"
	AlertNewForCustomerParamsTypeCreditBalanceDropped   AlertNewForCustomerParamsType = "credit_balance_dropped"
	AlertNewForCustomerParamsTypeCreditBalanceRecovered AlertNewForCustomerParamsType = "credit_balance_recovered"
)

func (r AlertNewForCustomerParamsType) IsKnown() bool {
	switch r {
	case AlertNewForCustomerParamsTypeUsageExceeded, AlertNewForCustomerParamsTypeCostExceeded, AlertNewForCustomerParamsTypeCreditBalanceDepleted, AlertNewForCustomerParamsTypeCreditBalanceDropped, AlertNewForCustomerParamsTypeCreditBalanceRecovered:
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
	Thresholds param.Field[[]shared.ThresholdModelParam] `json:"thresholds"`
}

func (r AlertNewForExternalCustomerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of alert to create. This must be a valid alert type.
type AlertNewForExternalCustomerParamsType string

const (
	AlertNewForExternalCustomerParamsTypeUsageExceeded          AlertNewForExternalCustomerParamsType = "usage_exceeded"
	AlertNewForExternalCustomerParamsTypeCostExceeded           AlertNewForExternalCustomerParamsType = "cost_exceeded"
	AlertNewForExternalCustomerParamsTypeCreditBalanceDepleted  AlertNewForExternalCustomerParamsType = "credit_balance_depleted"
	AlertNewForExternalCustomerParamsTypeCreditBalanceDropped   AlertNewForExternalCustomerParamsType = "credit_balance_dropped"
	AlertNewForExternalCustomerParamsTypeCreditBalanceRecovered AlertNewForExternalCustomerParamsType = "credit_balance_recovered"
)

func (r AlertNewForExternalCustomerParamsType) IsKnown() bool {
	switch r {
	case AlertNewForExternalCustomerParamsTypeUsageExceeded, AlertNewForExternalCustomerParamsTypeCostExceeded, AlertNewForExternalCustomerParamsTypeCreditBalanceDepleted, AlertNewForExternalCustomerParamsTypeCreditBalanceDropped, AlertNewForExternalCustomerParamsTypeCreditBalanceRecovered:
		return true
	}
	return false
}

type AlertNewForSubscriptionParams struct {
	// The thresholds that define the values at which the alert will be triggered.
	Thresholds param.Field[[]shared.ThresholdModelParam] `json:"thresholds,required"`
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
	AlertNewForSubscriptionParamsTypeUsageExceeded          AlertNewForSubscriptionParamsType = "usage_exceeded"
	AlertNewForSubscriptionParamsTypeCostExceeded           AlertNewForSubscriptionParamsType = "cost_exceeded"
	AlertNewForSubscriptionParamsTypeCreditBalanceDepleted  AlertNewForSubscriptionParamsType = "credit_balance_depleted"
	AlertNewForSubscriptionParamsTypeCreditBalanceDropped   AlertNewForSubscriptionParamsType = "credit_balance_dropped"
	AlertNewForSubscriptionParamsTypeCreditBalanceRecovered AlertNewForSubscriptionParamsType = "credit_balance_recovered"
)

func (r AlertNewForSubscriptionParamsType) IsKnown() bool {
	switch r {
	case AlertNewForSubscriptionParamsTypeUsageExceeded, AlertNewForSubscriptionParamsTypeCostExceeded, AlertNewForSubscriptionParamsTypeCreditBalanceDepleted, AlertNewForSubscriptionParamsTypeCreditBalanceDropped, AlertNewForSubscriptionParamsTypeCreditBalanceRecovered:
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
