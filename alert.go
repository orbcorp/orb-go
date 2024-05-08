// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
)

// AlertService contains methods and other services that help with interacting with
// the orb API. Note, unlike clients, this service does not read variables from the
// environment automatically. You should not instantiate this service directly, and
// instead use the [NewAlertService] method instead.
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

// This endpoint can be used to enable an alert.
func (r *AlertService) Enable(ctx context.Context, alertConfigurationID string, opts ...option.RequestOption) (res *Alert, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("alerts/%s/enable", alertConfigurationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// An
// [Alert within Orb](https://docs.withorb.com/guides/product-catalog/configuring-alerts)
// monitors a customer's spending, usage, or credit balance and triggers a webhook
// when a threshold is exceeded.
//
// Alerts can be configured to monitor usage, cost, or credit balance. Alerts can
// be scoped to either a customer, a plan, or a subscription.
//
// Customer scoped alerts track a customer's credit balance. Valid customer alert
// types are "credit_balance_depleted", "credit_balance_recovered", and
// "credit_balance_dropped".
//
// Subscription scoped alerts track a subscriptions's usage or cost. Valid plan
// alert types are "usage_exceeded" or "cost_exceeded".
//
// Plan scoped alerts are similar to subscriptions alerts but when a plan alert is
// created, it is propagated to all subscriptions associated with the plan.
// Disabling a plan alert will disable the alert for all subscriptions. Valid plan
// alert types are "usage_exceeded" or "cost_exceeded".
type Alert struct {
	// Also referred to as alert_id in this documentation.
	ID string `json:"id,required"`
	// The creation time of the resource in Orb.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the currency the credit balance for this alert is denominated in.
	Currency string `json:"currency,required,nullable"`
	// The customer that the alert is scoped to.
	Customer map[string]string `json:"customer,required,nullable"`
	// Whether the alert is enabled or disabled.
	Enabled bool              `json:"enabled,required"`
	Metric  map[string]string `json:"metric,required,nullable"`
	// The plan that the alert is scoped to.
	Plan         map[string]string `json:"plan,required,nullable"`
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
	// The value at which an alert will fire. For credit balance alerts, the alert will fire at or below this value. For usage and
	//
	//	cost alerts, the alert will fire at or above this value.
	Value int64              `json:"value,required"`
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
	AlertTypeCreditBalanceDepleted  AlertType = "credit_balance_depleted"
	AlertTypeCreditBalanceDropped   AlertType = "credit_balance_dropped"
	AlertTypeCreditBalanceRecovered AlertType = "credit_balance_recovered"
)

func (r AlertType) IsKnown() bool {
	switch r {
	case AlertTypeCreditBalanceDepleted, AlertTypeCreditBalanceDropped, AlertTypeCreditBalanceRecovered:
		return true
	}
	return false
}
