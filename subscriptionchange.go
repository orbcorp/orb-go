// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package orb

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/orbcorp/orb-go/internal/apijson"
	"github.com/orbcorp/orb-go/internal/param"
	"github.com/orbcorp/orb-go/internal/requestconfig"
	"github.com/orbcorp/orb-go/option"
	"github.com/tidwall/gjson"
)

// SubscriptionChangeService contains methods and other services that help with
// interacting with the orb API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionChangeService] method instead.
type SubscriptionChangeService struct {
	Options []option.RequestOption
}

// NewSubscriptionChangeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSubscriptionChangeService(opts ...option.RequestOption) (r *SubscriptionChangeService) {
	r = &SubscriptionChangeService{}
	r.Options = opts
	return
}

// This endpoint returns a subscription change given an identifier.
//
// A subscription change is created by including
// `Create-Pending-Subscription-Change: True` in the header of a subscription
// mutation API call (e.g.
// [create subscription endpoint](/api-reference/subscription/create-subscription),
// [schedule plan change endpoint](/api-reference/subscription/schedule-plan-change),
// ...). The subscription change will be referenced by the
// `pending_subscription_change` field in the response.
func (r *SubscriptionChangeService) Get(ctx context.Context, subscriptionChangeID string, opts ...option.RequestOption) (res *SubscriptionChangeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionChangeID == "" {
		err = errors.New("missing required subscription_change_id parameter")
		return
	}
	path := fmt.Sprintf("subscription_changes/%s", subscriptionChangeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply a subscription change to perform the intended action. If a positive amount
// is passed with a request to this endpoint, any eligible invoices that were
// created will be issued immediately if they only contain in-advance fees.
func (r *SubscriptionChangeService) Apply(ctx context.Context, subscriptionChangeID string, body SubscriptionChangeApplyParams, opts ...option.RequestOption) (res *SubscriptionChangeApplyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionChangeID == "" {
		err = errors.New("missing required subscription_change_id parameter")
		return
	}
	path := fmt.Sprintf("subscription_changes/%s/apply", subscriptionChangeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cancel a subscription change. The change can no longer be applied. A
// subscription can only have one "pending" change at a time - use this endpoint to
// cancel an existing change before creating a new one.
func (r *SubscriptionChangeService) Cancel(ctx context.Context, subscriptionChangeID string, opts ...option.RequestOption) (res *SubscriptionChangeCancelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if subscriptionChangeID == "" {
		err = errors.New("missing required subscription_change_id parameter")
		return
	}
	path := fmt.Sprintf("subscription_changes/%s/cancel", subscriptionChangeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeGetResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                                 `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeGetResponseStatus       `json:"status,required"`
	Subscription   SubscriptionChangeGetResponseSubscription `json:"subscription,required,nullable"`
	// When this change was applied.
	AppliedAt time.Time `json:"applied_at,nullable" format:"date-time"`
	// When this change was cancelled.
	CancelledAt time.Time                         `json:"cancelled_at,nullable" format:"date-time"`
	JSON        subscriptionChangeGetResponseJSON `json:"-"`
}

// subscriptionChangeGetResponseJSON contains the JSON metadata for the struct
// [SubscriptionChangeGetResponse]
type subscriptionChangeGetResponseJSON struct {
	ID             apijson.Field
	ExpirationTime apijson.Field
	Status         apijson.Field
	Subscription   apijson.Field
	AppliedAt      apijson.Field
	CancelledAt    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseStatus string

const (
	SubscriptionChangeGetResponseStatusPending   SubscriptionChangeGetResponseStatus = "pending"
	SubscriptionChangeGetResponseStatusApplied   SubscriptionChangeGetResponseStatus = "applied"
	SubscriptionChangeGetResponseStatusCancelled SubscriptionChangeGetResponseStatus = "cancelled"
)

func (r SubscriptionChangeGetResponseStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseStatusPending, SubscriptionChangeGetResponseStatusApplied, SubscriptionChangeGetResponseStatusCancelled:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscription struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription sorted by the start_date of the
	// adjustment interval.
	AdjustmentIntervals []SubscriptionChangeGetResponseSubscriptionAdjustmentInterval `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                                                     `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration SubscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfiguration `json:"billing_cycle_anchor_configuration,required"`
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
	// The discount intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	DiscountIntervals []SubscriptionChangeGetResponseSubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                                                           `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []SubscriptionChangeGetResponseSubscriptionFixedFeeQuantitySchedule `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                                              `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MaximumIntervals []SubscriptionChangeGetResponseSubscriptionMaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MinimumIntervals []SubscriptionChangeGetResponseSubscriptionMinimumInterval `json:"minimum_intervals,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// A pending subscription change if one exists on this subscription.
	PendingSubscriptionChange SubscriptionChangeGetResponseSubscriptionPendingSubscriptionChange `json:"pending_subscription_change,required,nullable"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan Plan `json:"plan,required,nullable"`
	// The price intervals for this subscription.
	PriceIntervals []SubscriptionChangeGetResponseSubscriptionPriceInterval `json:"price_intervals,required"`
	RedeemedCoupon SubscriptionChangeGetResponseSubscriptionRedeemedCoupon  `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                                          `json:"start_date,required" format:"date-time"`
	Status    SubscriptionChangeGetResponseSubscriptionStatus    `json:"status,required"`
	TrialInfo SubscriptionChangeGetResponseSubscriptionTrialInfo `json:"trial_info,required"`
	// The resources that were changed as part of this operation. Only present when
	// fetched through the subscription changes API or if the
	// `include_changed_resources` parameter was passed in the request.
	ChangedResources SubscriptionChangeGetResponseSubscriptionChangedResources `json:"changed_resources,nullable"`
	JSON             subscriptionChangeGetResponseSubscriptionJSON             `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionJSON contains the JSON metadata for the
// struct [SubscriptionChangeGetResponseSubscription]
type subscriptionChangeGetResponseSubscriptionJSON struct {
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
	ChangedResources                apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentInterval struct {
	ID         string                                                                 `json:"id,required"`
	Adjustment SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time                                                       `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionAdjustmentIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentInterval]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment struct {
	ID             string                                                                               `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                    `json:"usage_discount"`
	JSON          subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentJSON `json:"-"`
	union         SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentUnion
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment) AsUnion() SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
// or
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentUnion interface {
	implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                                                               `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                                                    `json:"usage_discount,required"`
	JSON          subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                         `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                                                                `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                      `json:"reason,required,nullable"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                          `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                                                                    `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                          `json:"reason,required,nullable"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                              `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment struct {
	ID             string                                                                                                         `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                               `json:"reason,required,nullable"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                   `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment struct {
	ID             string                                                                                                         `json:"id,required"`
	AdjustmentType SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                               `json:"reason,required,nullable"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) implementsSubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                   `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter]
type subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType string

const (
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum            SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "minimum"
	SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum            SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum, SubscriptionChangeGetResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfiguration struct {
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
	Year int64                                                                        `json:"year,nullable"`
	JSON subscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfigurationJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfigurationJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfiguration]
type subscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfigurationJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionBillingCycleAnchorConfigurationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                                            `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter],
	// [[]SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter].
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
	UsageDiscount float64                                                       `json:"usage_discount"`
	JSON          subscriptionChangeGetResponseSubscriptionDiscountIntervalJSON `json:"-"`
	union         SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUnion
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountInterval]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalJSON struct {
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

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeGetResponseSubscriptionDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval],
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
func (r SubscriptionChangeGetResponseSubscriptionDiscountInterval) AsUnion() SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
// or
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUnion interface {
	implementsSubscriptionChangeGetResponseSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval{}),
			DiscriminatorValue: "usage",
		},
	)
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                     `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                            `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountInterval) implementsSubscriptionChangeGetResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType = "amount"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                   `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                         `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter `json:"filters,required"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                                `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) implementsSubscriptionChangeGetResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType = "percentage"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                       `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                    `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                                                                             `json:"usage_discount,required"`
	JSON          subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountInterval) implementsSubscriptionChangeGetResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType = "usage"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                  `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter]
type subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType string

const (
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypeAmount     SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType = "amount"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypePercentage SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType = "percentage"
	SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypeUsage      SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType = "usage"
)

func (r SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypeAmount, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypePercentage, SubscriptionChangeGetResponseSubscriptionDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionFixedFeeQuantitySchedule struct {
	EndDate   time.Time                                                             `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                                                                `json:"price_id,required"`
	Quantity  float64                                                               `json:"quantity,required"`
	StartDate time.Time                                                             `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionFixedFeeQuantityScheduleJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionFixedFeeQuantityScheduleJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionFixedFeeQuantitySchedule]
type subscriptionChangeGetResponseSubscriptionFixedFeeQuantityScheduleJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionFixedFeeQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionFixedFeeQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionMaximumInterval struct {
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this maximum interval applies to.
	Filters []SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFilter `json:"filters,required"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time                                                    `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionMaximumIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionMaximumIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionMaximumInterval]
type subscriptionChangeGetResponseSubscriptionMaximumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionMaximumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionMaximumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                            `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionMaximumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionMaximumIntervalsFilterJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFilter]
type subscriptionChangeGetResponseSubscriptionMaximumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionMaximumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionMinimumInterval struct {
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this minimum interval applies to.
	Filters []SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFilter `json:"filters,required"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time                                                    `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionMinimumIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionMinimumIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionMinimumInterval]
type subscriptionChangeGetResponseSubscriptionMinimumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionMinimumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionMinimumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                            `json:"values,required"`
	JSON   subscriptionChangeGetResponseSubscriptionMinimumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionMinimumIntervalsFilterJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFilter]
type subscriptionChangeGetResponseSubscriptionMinimumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionMinimumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField string

const (
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPriceID       SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField = "price_id"
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldItemID        SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField = "item_id"
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPriceType     SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField = "price_type"
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldCurrency      SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField = "currency"
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPriceID, SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldItemID, SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPriceType, SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldCurrency, SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperator string

const (
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperator = "includes"
	SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes, SubscriptionChangeGetResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

// A pending subscription change if one exists on this subscription.
type SubscriptionChangeGetResponseSubscriptionPendingSubscriptionChange struct {
	ID   string                                                                 `json:"id,required"`
	JSON subscriptionChangeGetResponseSubscriptionPendingSubscriptionChangeJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionPendingSubscriptionChangeJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionPendingSubscriptionChange]
type subscriptionChangeGetResponseSubscriptionPendingSubscriptionChangeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionPendingSubscriptionChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionPendingSubscriptionChangeJSON) RawJSON() string {
	return r.raw
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type SubscriptionChangeGetResponseSubscriptionPriceInterval struct {
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
	// An additional filter to apply to usage queries.
	Filter string `json:"filter,required,nullable"`
	// The fixed fee quantity transitions for this price interval. This is only
	// relevant for fixed fees.
	FixedFeeQuantityTransitions []SubscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// For more on the types of prices, see
	// [the core concepts documentation](/core-concepts#plan-and-price)
	Price Price `json:"price,required"`
	// The start date of the price interval. This is the date that Orb starts billing
	// for this price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this price interval.
	UsageCustomerIDs []string                                                   `json:"usage_customer_ids,required,nullable"`
	JSON             subscriptionChangeGetResponseSubscriptionPriceIntervalJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionPriceIntervalJSON contains the JSON
// metadata for the struct [SubscriptionChangeGetResponseSubscriptionPriceInterval]
type subscriptionChangeGetResponseSubscriptionPriceIntervalJSON struct {
	ID                            apijson.Field
	BillingCycleDay               apijson.Field
	CurrentBillingPeriodEndDate   apijson.Field
	CurrentBillingPeriodStartDate apijson.Field
	EndDate                       apijson.Field
	Filter                        apijson.Field
	FixedFeeQuantityTransitions   apijson.Field
	Price                         apijson.Field
	StartDate                     apijson.Field
	UsageCustomerIDs              apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionPriceInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionPriceIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition struct {
	EffectiveDate time.Time                                                                             `json:"effective_date,required" format:"date-time"`
	PriceID       string                                                                                `json:"price_id,required"`
	Quantity      int64                                                                                 `json:"quantity,required"`
	JSON          subscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition]
type subscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionRedeemedCoupon struct {
	CouponID  string                                                      `json:"coupon_id,required"`
	EndDate   time.Time                                                   `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time                                                   `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeGetResponseSubscriptionRedeemedCouponJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionRedeemedCouponJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionRedeemedCoupon]
type subscriptionChangeGetResponseSubscriptionRedeemedCouponJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionRedeemedCoupon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionRedeemedCouponJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeGetResponseSubscriptionStatus string

const (
	SubscriptionChangeGetResponseSubscriptionStatusActive   SubscriptionChangeGetResponseSubscriptionStatus = "active"
	SubscriptionChangeGetResponseSubscriptionStatusEnded    SubscriptionChangeGetResponseSubscriptionStatus = "ended"
	SubscriptionChangeGetResponseSubscriptionStatusUpcoming SubscriptionChangeGetResponseSubscriptionStatus = "upcoming"
)

func (r SubscriptionChangeGetResponseSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeGetResponseSubscriptionStatusActive, SubscriptionChangeGetResponseSubscriptionStatusEnded, SubscriptionChangeGetResponseSubscriptionStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionChangeGetResponseSubscriptionTrialInfo struct {
	EndDate time.Time                                              `json:"end_date,required,nullable" format:"date-time"`
	JSON    subscriptionChangeGetResponseSubscriptionTrialInfoJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionTrialInfoJSON contains the JSON
// metadata for the struct [SubscriptionChangeGetResponseSubscriptionTrialInfo]
type subscriptionChangeGetResponseSubscriptionTrialInfoJSON struct {
	EndDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionTrialInfoJSON) RawJSON() string {
	return r.raw
}

// The resources that were changed as part of this operation. Only present when
// fetched through the subscription changes API or if the
// `include_changed_resources` parameter was passed in the request.
type SubscriptionChangeGetResponseSubscriptionChangedResources struct {
	// The credit notes that were created as part of this operation.
	CreatedCreditNotes []CreditNote `json:"created_credit_notes,required"`
	// The invoices that were created as part of this operation.
	CreatedInvoices []Invoice `json:"created_invoices,required"`
	// The credit notes that were voided as part of this operation.
	VoidedCreditNotes []CreditNote `json:"voided_credit_notes,required"`
	// The invoices that were voided as part of this operation.
	VoidedInvoices []Invoice                                                     `json:"voided_invoices,required"`
	JSON           subscriptionChangeGetResponseSubscriptionChangedResourcesJSON `json:"-"`
}

// subscriptionChangeGetResponseSubscriptionChangedResourcesJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeGetResponseSubscriptionChangedResources]
type subscriptionChangeGetResponseSubscriptionChangedResourcesJSON struct {
	CreatedCreditNotes apijson.Field
	CreatedInvoices    apijson.Field
	VoidedCreditNotes  apijson.Field
	VoidedInvoices     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeGetResponseSubscriptionChangedResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeGetResponseSubscriptionChangedResourcesJSON) RawJSON() string {
	return r.raw
}

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeApplyResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                                   `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeApplyResponseStatus       `json:"status,required"`
	Subscription   SubscriptionChangeApplyResponseSubscription `json:"subscription,required,nullable"`
	// When this change was applied.
	AppliedAt time.Time `json:"applied_at,nullable" format:"date-time"`
	// When this change was cancelled.
	CancelledAt time.Time                           `json:"cancelled_at,nullable" format:"date-time"`
	JSON        subscriptionChangeApplyResponseJSON `json:"-"`
}

// subscriptionChangeApplyResponseJSON contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponse]
type subscriptionChangeApplyResponseJSON struct {
	ID             apijson.Field
	ExpirationTime apijson.Field
	Status         apijson.Field
	Subscription   apijson.Field
	AppliedAt      apijson.Field
	CancelledAt    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseStatus string

const (
	SubscriptionChangeApplyResponseStatusPending   SubscriptionChangeApplyResponseStatus = "pending"
	SubscriptionChangeApplyResponseStatusApplied   SubscriptionChangeApplyResponseStatus = "applied"
	SubscriptionChangeApplyResponseStatusCancelled SubscriptionChangeApplyResponseStatus = "cancelled"
)

func (r SubscriptionChangeApplyResponseStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseStatusPending, SubscriptionChangeApplyResponseStatusApplied, SubscriptionChangeApplyResponseStatusCancelled:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscription struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription sorted by the start_date of the
	// adjustment interval.
	AdjustmentIntervals []SubscriptionChangeApplyResponseSubscriptionAdjustmentInterval `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                                                       `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration SubscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfiguration `json:"billing_cycle_anchor_configuration,required"`
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
	// The discount intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	DiscountIntervals []SubscriptionChangeApplyResponseSubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                                                             `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []SubscriptionChangeApplyResponseSubscriptionFixedFeeQuantitySchedule `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                                                `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MaximumIntervals []SubscriptionChangeApplyResponseSubscriptionMaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MinimumIntervals []SubscriptionChangeApplyResponseSubscriptionMinimumInterval `json:"minimum_intervals,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// A pending subscription change if one exists on this subscription.
	PendingSubscriptionChange SubscriptionChangeApplyResponseSubscriptionPendingSubscriptionChange `json:"pending_subscription_change,required,nullable"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan Plan `json:"plan,required,nullable"`
	// The price intervals for this subscription.
	PriceIntervals []SubscriptionChangeApplyResponseSubscriptionPriceInterval `json:"price_intervals,required"`
	RedeemedCoupon SubscriptionChangeApplyResponseSubscriptionRedeemedCoupon  `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                                            `json:"start_date,required" format:"date-time"`
	Status    SubscriptionChangeApplyResponseSubscriptionStatus    `json:"status,required"`
	TrialInfo SubscriptionChangeApplyResponseSubscriptionTrialInfo `json:"trial_info,required"`
	// The resources that were changed as part of this operation. Only present when
	// fetched through the subscription changes API or if the
	// `include_changed_resources` parameter was passed in the request.
	ChangedResources SubscriptionChangeApplyResponseSubscriptionChangedResources `json:"changed_resources,nullable"`
	JSON             subscriptionChangeApplyResponseSubscriptionJSON             `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionJSON contains the JSON metadata for
// the struct [SubscriptionChangeApplyResponseSubscription]
type subscriptionChangeApplyResponseSubscriptionJSON struct {
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
	ChangedResources                apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentInterval struct {
	ID         string                                                                   `json:"id,required"`
	Adjustment SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time                                                         `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentInterval]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment struct {
	ID             string                                                                                 `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                      `json:"usage_discount"`
	JSON          subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentJSON `json:"-"`
	union         SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentUnion
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment) AsUnion() SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
// or
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentUnion interface {
	implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                                                                 `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                                                      `json:"usage_discount,required"`
	JSON          subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                           `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                                                                  `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                        `json:"reason,required,nullable"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                            `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                                                                      `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                            `json:"reason,required,nullable"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                                `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment struct {
	ID             string                                                                                                           `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                 `json:"reason,required,nullable"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                     `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment struct {
	ID             string                                                                                                           `json:"id,required"`
	AdjustmentType SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                 `json:"reason,required,nullable"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) implementsSubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                     `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter]
type subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType string

const (
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum            SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "minimum"
	SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum            SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum, SubscriptionChangeApplyResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfiguration struct {
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
	Year int64                                                                          `json:"year,nullable"`
	JSON subscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfigurationJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfigurationJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfiguration]
type subscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfigurationJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionBillingCycleAnchorConfigurationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                                              `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter],
	// [[]SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter].
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
	UsageDiscount float64                                                         `json:"usage_discount"`
	JSON          subscriptionChangeApplyResponseSubscriptionDiscountIntervalJSON `json:"-"`
	union         SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUnion
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountInterval]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalJSON struct {
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

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeApplyResponseSubscriptionDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval],
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
func (r SubscriptionChangeApplyResponseSubscriptionDiscountInterval) AsUnion() SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
// or
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUnion interface {
	implementsSubscriptionChangeApplyResponseSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval{}),
			DiscriminatorValue: "usage",
		},
	)
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                       `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                              `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountInterval) implementsSubscriptionChangeApplyResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType = "amount"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                     `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                           `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter `json:"filters,required"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                                  `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) implementsSubscriptionChangeApplyResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType = "percentage"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                         `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                      `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                                                                               `json:"usage_discount,required"`
	JSON          subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountInterval) implementsSubscriptionChangeApplyResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType = "usage"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                    `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter]
type subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType string

const (
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypeAmount     SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType = "amount"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypePercentage SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType = "percentage"
	SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypeUsage      SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType = "usage"
)

func (r SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypeAmount, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypePercentage, SubscriptionChangeApplyResponseSubscriptionDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionFixedFeeQuantitySchedule struct {
	EndDate   time.Time                                                               `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                                                                  `json:"price_id,required"`
	Quantity  float64                                                                 `json:"quantity,required"`
	StartDate time.Time                                                               `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionFixedFeeQuantityScheduleJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionFixedFeeQuantityScheduleJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionFixedFeeQuantitySchedule]
type subscriptionChangeApplyResponseSubscriptionFixedFeeQuantityScheduleJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionFixedFeeQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionFixedFeeQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionMaximumInterval struct {
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this maximum interval applies to.
	Filters []SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilter `json:"filters,required"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time                                                      `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionMaximumIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionMaximumIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionMaximumInterval]
type subscriptionChangeApplyResponseSubscriptionMaximumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionMaximumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionMaximumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilterJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilter]
type subscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionMaximumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionMinimumInterval struct {
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this minimum interval applies to.
	Filters []SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilter `json:"filters,required"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time                                                      `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionMinimumIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionMinimumIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionMinimumInterval]
type subscriptionChangeApplyResponseSubscriptionMinimumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionMinimumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionMinimumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                              `json:"values,required"`
	JSON   subscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilterJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilter]
type subscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionMinimumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField string

const (
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPriceID       SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField = "price_id"
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldItemID        SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField = "item_id"
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPriceType     SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField = "price_type"
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldCurrency      SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField = "currency"
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPriceID, SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldItemID, SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPriceType, SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldCurrency, SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperator string

const (
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperator = "includes"
	SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes, SubscriptionChangeApplyResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

// A pending subscription change if one exists on this subscription.
type SubscriptionChangeApplyResponseSubscriptionPendingSubscriptionChange struct {
	ID   string                                                                   `json:"id,required"`
	JSON subscriptionChangeApplyResponseSubscriptionPendingSubscriptionChangeJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionPendingSubscriptionChangeJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionPendingSubscriptionChange]
type subscriptionChangeApplyResponseSubscriptionPendingSubscriptionChangeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionPendingSubscriptionChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionPendingSubscriptionChangeJSON) RawJSON() string {
	return r.raw
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type SubscriptionChangeApplyResponseSubscriptionPriceInterval struct {
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
	// An additional filter to apply to usage queries.
	Filter string `json:"filter,required,nullable"`
	// The fixed fee quantity transitions for this price interval. This is only
	// relevant for fixed fees.
	FixedFeeQuantityTransitions []SubscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// For more on the types of prices, see
	// [the core concepts documentation](/core-concepts#plan-and-price)
	Price Price `json:"price,required"`
	// The start date of the price interval. This is the date that Orb starts billing
	// for this price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this price interval.
	UsageCustomerIDs []string                                                     `json:"usage_customer_ids,required,nullable"`
	JSON             subscriptionChangeApplyResponseSubscriptionPriceIntervalJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionPriceIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionPriceInterval]
type subscriptionChangeApplyResponseSubscriptionPriceIntervalJSON struct {
	ID                            apijson.Field
	BillingCycleDay               apijson.Field
	CurrentBillingPeriodEndDate   apijson.Field
	CurrentBillingPeriodStartDate apijson.Field
	EndDate                       apijson.Field
	Filter                        apijson.Field
	FixedFeeQuantityTransitions   apijson.Field
	Price                         apijson.Field
	StartDate                     apijson.Field
	UsageCustomerIDs              apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionPriceInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionPriceIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition struct {
	EffectiveDate time.Time                                                                               `json:"effective_date,required" format:"date-time"`
	PriceID       string                                                                                  `json:"price_id,required"`
	Quantity      int64                                                                                   `json:"quantity,required"`
	JSON          subscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition]
type subscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionRedeemedCoupon struct {
	CouponID  string                                                        `json:"coupon_id,required"`
	EndDate   time.Time                                                     `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time                                                     `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeApplyResponseSubscriptionRedeemedCouponJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionRedeemedCouponJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionRedeemedCoupon]
type subscriptionChangeApplyResponseSubscriptionRedeemedCouponJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionRedeemedCoupon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionRedeemedCouponJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyResponseSubscriptionStatus string

const (
	SubscriptionChangeApplyResponseSubscriptionStatusActive   SubscriptionChangeApplyResponseSubscriptionStatus = "active"
	SubscriptionChangeApplyResponseSubscriptionStatusEnded    SubscriptionChangeApplyResponseSubscriptionStatus = "ended"
	SubscriptionChangeApplyResponseSubscriptionStatusUpcoming SubscriptionChangeApplyResponseSubscriptionStatus = "upcoming"
)

func (r SubscriptionChangeApplyResponseSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeApplyResponseSubscriptionStatusActive, SubscriptionChangeApplyResponseSubscriptionStatusEnded, SubscriptionChangeApplyResponseSubscriptionStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionChangeApplyResponseSubscriptionTrialInfo struct {
	EndDate time.Time                                                `json:"end_date,required,nullable" format:"date-time"`
	JSON    subscriptionChangeApplyResponseSubscriptionTrialInfoJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionTrialInfoJSON contains the JSON
// metadata for the struct [SubscriptionChangeApplyResponseSubscriptionTrialInfo]
type subscriptionChangeApplyResponseSubscriptionTrialInfoJSON struct {
	EndDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionTrialInfoJSON) RawJSON() string {
	return r.raw
}

// The resources that were changed as part of this operation. Only present when
// fetched through the subscription changes API or if the
// `include_changed_resources` parameter was passed in the request.
type SubscriptionChangeApplyResponseSubscriptionChangedResources struct {
	// The credit notes that were created as part of this operation.
	CreatedCreditNotes []CreditNote `json:"created_credit_notes,required"`
	// The invoices that were created as part of this operation.
	CreatedInvoices []Invoice `json:"created_invoices,required"`
	// The credit notes that were voided as part of this operation.
	VoidedCreditNotes []CreditNote `json:"voided_credit_notes,required"`
	// The invoices that were voided as part of this operation.
	VoidedInvoices []Invoice                                                       `json:"voided_invoices,required"`
	JSON           subscriptionChangeApplyResponseSubscriptionChangedResourcesJSON `json:"-"`
}

// subscriptionChangeApplyResponseSubscriptionChangedResourcesJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeApplyResponseSubscriptionChangedResources]
type subscriptionChangeApplyResponseSubscriptionChangedResourcesJSON struct {
	CreatedCreditNotes apijson.Field
	CreatedInvoices    apijson.Field
	VoidedCreditNotes  apijson.Field
	VoidedInvoices     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeApplyResponseSubscriptionChangedResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeApplyResponseSubscriptionChangedResourcesJSON) RawJSON() string {
	return r.raw
}

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeCancelResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                                    `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeCancelResponseStatus       `json:"status,required"`
	Subscription   SubscriptionChangeCancelResponseSubscription `json:"subscription,required,nullable"`
	// When this change was applied.
	AppliedAt time.Time `json:"applied_at,nullable" format:"date-time"`
	// When this change was cancelled.
	CancelledAt time.Time                            `json:"cancelled_at,nullable" format:"date-time"`
	JSON        subscriptionChangeCancelResponseJSON `json:"-"`
}

// subscriptionChangeCancelResponseJSON contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponse]
type subscriptionChangeCancelResponseJSON struct {
	ID             apijson.Field
	ExpirationTime apijson.Field
	Status         apijson.Field
	Subscription   apijson.Field
	AppliedAt      apijson.Field
	CancelledAt    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseStatus string

const (
	SubscriptionChangeCancelResponseStatusPending   SubscriptionChangeCancelResponseStatus = "pending"
	SubscriptionChangeCancelResponseStatusApplied   SubscriptionChangeCancelResponseStatus = "applied"
	SubscriptionChangeCancelResponseStatusCancelled SubscriptionChangeCancelResponseStatus = "cancelled"
)

func (r SubscriptionChangeCancelResponseStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseStatusPending, SubscriptionChangeCancelResponseStatusApplied, SubscriptionChangeCancelResponseStatusCancelled:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscription struct {
	ID string `json:"id,required"`
	// The current plan phase that is active, only if the subscription's plan has
	// phases.
	ActivePlanPhaseOrder int64 `json:"active_plan_phase_order,required,nullable"`
	// The adjustment intervals for this subscription sorted by the start_date of the
	// adjustment interval.
	AdjustmentIntervals []SubscriptionChangeCancelResponseSubscriptionAdjustmentInterval `json:"adjustment_intervals,required"`
	// Determines whether issued invoices for this subscription will automatically be
	// charged with the saved payment method on the due date. This property defaults to
	// the plan's behavior. If null, defaults to the customer's setting.
	AutoCollection                  bool                                                                        `json:"auto_collection,required,nullable"`
	BillingCycleAnchorConfiguration SubscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfiguration `json:"billing_cycle_anchor_configuration,required"`
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
	// The discount intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	DiscountIntervals []SubscriptionChangeCancelResponseSubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                                                              `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []SubscriptionChangeCancelResponseSubscriptionFixedFeeQuantitySchedule `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                                                 `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MaximumIntervals []SubscriptionChangeCancelResponseSubscriptionMaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MinimumIntervals []SubscriptionChangeCancelResponseSubscriptionMinimumInterval `json:"minimum_intervals,required"`
	// The name of the subscription.
	Name string `json:"name,required"`
	// Determines the difference between the invoice issue date for subscription
	// invoices as the date that they are due. A value of `0` here represents that the
	// invoice is due on issue, whereas a value of `30` represents that the customer
	// has a month to pay the invoice.
	NetTerms int64 `json:"net_terms,required"`
	// A pending subscription change if one exists on this subscription.
	PendingSubscriptionChange SubscriptionChangeCancelResponseSubscriptionPendingSubscriptionChange `json:"pending_subscription_change,required,nullable"`
	// The [Plan](/core-concepts#plan-and-price) resource represents a plan that can be
	// subscribed to by a customer. Plans define the billing behavior of the
	// subscription. You can see more about how to configure prices in the
	// [Price resource](/reference/price).
	Plan Plan `json:"plan,required,nullable"`
	// The price intervals for this subscription.
	PriceIntervals []SubscriptionChangeCancelResponseSubscriptionPriceInterval `json:"price_intervals,required"`
	RedeemedCoupon SubscriptionChangeCancelResponseSubscriptionRedeemedCoupon  `json:"redeemed_coupon,required,nullable"`
	// The date Orb starts billing for this subscription.
	StartDate time.Time                                             `json:"start_date,required" format:"date-time"`
	Status    SubscriptionChangeCancelResponseSubscriptionStatus    `json:"status,required"`
	TrialInfo SubscriptionChangeCancelResponseSubscriptionTrialInfo `json:"trial_info,required"`
	// The resources that were changed as part of this operation. Only present when
	// fetched through the subscription changes API or if the
	// `include_changed_resources` parameter was passed in the request.
	ChangedResources SubscriptionChangeCancelResponseSubscriptionChangedResources `json:"changed_resources,nullable"`
	JSON             subscriptionChangeCancelResponseSubscriptionJSON             `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionJSON contains the JSON metadata for
// the struct [SubscriptionChangeCancelResponseSubscription]
type subscriptionChangeCancelResponseSubscriptionJSON struct {
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
	ChangedResources                apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentInterval struct {
	ID         string                                                                    `json:"id,required"`
	Adjustment SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment `json:"adjustment,required"`
	// The price interval IDs that this adjustment applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the adjustment interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The start date of the adjustment interval.
	StartDate time.Time                                                          `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentInterval]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalJSON struct {
	ID                        apijson.Field
	Adjustment                apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment struct {
	ID             string                                                                                  `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// This field can have the runtime type of [[]string].
	AppliesToPriceIDs interface{} `json:"applies_to_price_ids,required"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter].
	Filters interface{} `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                       `json:"usage_discount"`
	JSON          subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentJSON `json:"-"`
	union         SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentUnion
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	AmountDiscount     apijson.Field
	ItemID             apijson.Field
	MaximumAmount      apijson.Field
	MinimumAmount      apijson.Field
	PercentageDiscount apijson.Field
	UsageDiscount      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment) AsUnion() SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment],
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
// or
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment].
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentUnion interface {
	implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentUnion)(nil)).Elem(),
		"adjustment_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment{}),
			DiscriminatorValue: "usage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment{}),
			DiscriminatorValue: "amount_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment{}),
			DiscriminatorValue: "percentage_discount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment{}),
			DiscriminatorValue: "minimum",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment{}),
			DiscriminatorValue: "maximum",
		},
	)
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment struct {
	ID             string                                                                                                                  `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string `json:"reason,required,nullable"`
	// The number of usage units by which to discount the price this adjustment applies
	// to in a given billing period.
	UsageDiscount float64                                                                                                       `json:"usage_discount,required"`
	JSON          subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	UsageDiscount     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustment) implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType = "usage_discount"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentAdjustmentTypeUsageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                            `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseUsageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment struct {
	ID             string                                                                                                                   `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The amount by which to discount the prices this adjustment applies to in a given
	// billing period.
	AmountDiscount string `json:"amount_discount,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                         `json:"reason,required,nullable"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AmountDiscount    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustment) implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType = "amount_discount"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentAdjustmentTypeAmountDiscount:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                             `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseAmountDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment struct {
	ID             string                                                                                                                       `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The percentage (as a value between 0 and 1) by which to discount the price
	// intervals this adjustment applies to in a given billing period.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                             `json:"reason,required,nullable"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON struct {
	ID                 apijson.Field
	AdjustmentType     apijson.Field
	AppliesToPriceIDs  apijson.Field
	Filters            apijson.Field
	IsInvoiceLevel     apijson.Field
	PercentageDiscount apijson.Field
	PlanPhaseOrder     apijson.Field
	Reason             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustment) implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType = "percentage_discount"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentAdjustmentTypePercentageDiscount:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                                 `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhasePercentageDiscountAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment struct {
	ID             string                                                                                                            `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The item ID that revenue from this minimum will be attributed to.
	ItemID string `json:"item_id,required"`
	// The minimum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                  `json:"reason,required,nullable"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	ItemID            apijson.Field
	MinimumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustment) implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType = "minimum"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentAdjustmentTypeMinimum:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                      `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMinimumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment struct {
	ID             string                                                                                                            `json:"id,required"`
	AdjustmentType SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType `json:"adjustment_type,required"`
	// The price IDs that this adjustment applies to.
	//
	// Deprecated: deprecated
	AppliesToPriceIDs []string `json:"applies_to_price_ids,required"`
	// The filters that determine which prices to apply this adjustment to.
	Filters []SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter `json:"filters,required"`
	// True for adjustments that apply to an entire invocice, false for adjustments
	// that apply to only one price.
	IsInvoiceLevel bool `json:"is_invoice_level,required"`
	// The maximum amount to charge in a given billing period for the prices this
	// adjustment applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The plan phase in which this adjustment is active.
	PlanPhaseOrder int64 `json:"plan_phase_order,required,nullable"`
	// The reason for the adjustment.
	Reason string                                                                                                  `json:"reason,required,nullable"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON struct {
	ID                apijson.Field
	AdjustmentType    apijson.Field
	AppliesToPriceIDs apijson.Field
	Filters           apijson.Field
	IsInvoiceLevel    apijson.Field
	MaximumAmount     apijson.Field
	PlanPhaseOrder    apijson.Field
	Reason            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustment) implementsSubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustment() {
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                                      `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter]
type subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentPlanPhaseMaximumAdjustmentFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType string

const (
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount      SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "usage_discount"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount     SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "amount_discount"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "percentage_discount"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum            SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "minimum"
	SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum            SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType = "maximum"
)

func (r SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeUsageDiscount, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeAmountDiscount, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypePercentageDiscount, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMinimum, SubscriptionChangeCancelResponseSubscriptionAdjustmentIntervalsAdjustmentAdjustmentTypeMaximum:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfiguration struct {
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
	Year int64                                                                           `json:"year,nullable"`
	JSON subscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfigurationJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfigurationJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfiguration]
type subscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfigurationJSON struct {
	Day         apijson.Field
	Month       apijson.Field
	Year        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionBillingCycleAnchorConfigurationJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                                               `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of
	// [[]SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter],
	// [[]SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter].
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
	UsageDiscount float64                                                          `json:"usage_discount"`
	JSON          subscriptionChangeCancelResponseSubscriptionDiscountIntervalJSON `json:"-"`
	union         SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUnion
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountInterval]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalJSON struct {
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

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = SubscriptionChangeCancelResponseSubscriptionDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval],
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
func (r SubscriptionChangeCancelResponseSubscriptionDiscountInterval) AsUnion() SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval],
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
// or
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval].
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUnion interface {
	implementsSubscriptionChangeCancelResponseSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUnion)(nil)).Elem(),
		"discount_type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval{}),
			DiscriminatorValue: "amount",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval{}),
			DiscriminatorValue: "percentage",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval{}),
			DiscriminatorValue: "usage",
		},
	)
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval struct {
	// Only available if discount_type is `amount`.
	AmountDiscount string `json:"amount_discount,required"`
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                        `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                               `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON struct {
	AmountDiscount            apijson.Field
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountInterval) implementsSubscriptionChangeCancelResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType = "amount"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalDiscountTypeAmount:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                      `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsAmountDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                            `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter `json:"filters,required"`
	// Only available if discount_type is `percentage`.This is a number between 0
	// and 1.
	PercentageDiscount float64 `json:"percentage_discount,required"`
	// The start date of the discount interval.
	StartDate time.Time                                                                                   `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	PercentageDiscount        apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountInterval) implementsSubscriptionChangeCancelResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType = "percentage"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalDiscountTypePercentage:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                          `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsPercentageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval struct {
	// The price interval ids that this discount interval applies to.
	AppliesToPriceIntervalIDs []string                                                                                       `json:"applies_to_price_interval_ids,required"`
	DiscountType              SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this discount interval applies to.
	Filters []SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter `json:"filters,required"`
	// The start date of the discount interval.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// Only available if discount_type is `usage`. Number of usage units that this
	// discount is for
	UsageDiscount float64                                                                                `json:"usage_discount,required"`
	JSON          subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	DiscountType              apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	StartDate                 apijson.Field
	UsageDiscount             apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountInterval) implementsSubscriptionChangeCancelResponseSubscriptionDiscountInterval() {
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType = "usage"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                                                     `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter]
type subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsUsageDiscountIntervalFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType string

const (
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypeAmount     SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType = "amount"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypePercentage SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType = "percentage"
	SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypeUsage      SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType = "usage"
)

func (r SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypeAmount, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypePercentage, SubscriptionChangeCancelResponseSubscriptionDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionFixedFeeQuantitySchedule struct {
	EndDate   time.Time                                                                `json:"end_date,required,nullable" format:"date-time"`
	PriceID   string                                                                   `json:"price_id,required"`
	Quantity  float64                                                                  `json:"quantity,required"`
	StartDate time.Time                                                                `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionFixedFeeQuantityScheduleJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionFixedFeeQuantityScheduleJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionFixedFeeQuantitySchedule]
type subscriptionChangeCancelResponseSubscriptionFixedFeeQuantityScheduleJSON struct {
	EndDate     apijson.Field
	PriceID     apijson.Field
	Quantity    apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionFixedFeeQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionFixedFeeQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionMaximumInterval struct {
	// The price interval ids that this maximum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the maximum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this maximum interval applies to.
	Filters []SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilter `json:"filters,required"`
	// The maximum amount to charge in a given billing period for the price intervals
	// this transform applies to.
	MaximumAmount string `json:"maximum_amount,required"`
	// The start date of the maximum interval.
	StartDate time.Time                                                       `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionMaximumIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionMaximumIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionMaximumInterval]
type subscriptionChangeCancelResponseSubscriptionMaximumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MaximumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionMaximumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionMaximumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                               `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilterJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilter]
type subscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionMaximumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionMaximumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionMinimumInterval struct {
	// The price interval ids that this minimum interval applies to.
	AppliesToPriceIntervalIDs []string `json:"applies_to_price_interval_ids,required"`
	// The end date of the minimum interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// The filters that determine which prices this minimum interval applies to.
	Filters []SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilter `json:"filters,required"`
	// The minimum amount to charge in a given billing period for the price intervals
	// this minimum applies to.
	MinimumAmount string `json:"minimum_amount,required"`
	// The start date of the minimum interval.
	StartDate time.Time                                                       `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionMinimumIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionMinimumIntervalJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionMinimumInterval]
type subscriptionChangeCancelResponseSubscriptionMinimumIntervalJSON struct {
	AppliesToPriceIntervalIDs apijson.Field
	EndDate                   apijson.Field
	Filters                   apijson.Field
	MinimumAmount             apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionMinimumInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionMinimumIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilter struct {
	// The property of the price to filter on.
	Field SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField `json:"field,required"`
	// Should prices that match the filter be included or excluded.
	Operator SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperator `json:"operator,required"`
	// The IDs or values that match this filter.
	Values []string                                                               `json:"values,required"`
	JSON   subscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilterJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilterJSON contains
// the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilter]
type subscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilterJSON struct {
	Field       apijson.Field
	Operator    apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionMinimumIntervalsFilterJSON) RawJSON() string {
	return r.raw
}

// The property of the price to filter on.
type SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField string

const (
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPriceID       SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField = "price_id"
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldItemID        SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField = "item_id"
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPriceType     SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField = "price_type"
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldCurrency      SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField = "currency"
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField = "pricing_unit_id"
)

func (r SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersField) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPriceID, SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldItemID, SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPriceType, SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldCurrency, SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersFieldPricingUnitID:
		return true
	}
	return false
}

// Should prices that match the filter be included or excluded.
type SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperator string

const (
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperator = "includes"
	SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperator = "excludes"
)

func (r SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperator) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperatorIncludes, SubscriptionChangeCancelResponseSubscriptionMinimumIntervalsFiltersOperatorExcludes:
		return true
	}
	return false
}

// A pending subscription change if one exists on this subscription.
type SubscriptionChangeCancelResponseSubscriptionPendingSubscriptionChange struct {
	ID   string                                                                    `json:"id,required"`
	JSON subscriptionChangeCancelResponseSubscriptionPendingSubscriptionChangeJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionPendingSubscriptionChangeJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionPendingSubscriptionChange]
type subscriptionChangeCancelResponseSubscriptionPendingSubscriptionChangeJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionPendingSubscriptionChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionPendingSubscriptionChangeJSON) RawJSON() string {
	return r.raw
}

// The Price Interval resource represents a period of time for which a price will
// bill on a subscription. A subscription’s price intervals define its billing
// behavior.
type SubscriptionChangeCancelResponseSubscriptionPriceInterval struct {
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
	// An additional filter to apply to usage queries.
	Filter string `json:"filter,required,nullable"`
	// The fixed fee quantity transitions for this price interval. This is only
	// relevant for fixed fees.
	FixedFeeQuantityTransitions []SubscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition `json:"fixed_fee_quantity_transitions,required,nullable"`
	// The Price resource represents a price that can be billed on a subscription,
	// resulting in a charge on an invoice in the form of an invoice line item. Prices
	// take a quantity and determine an amount to bill.
	//
	// Orb supports a few different pricing models out of the box. Each of these models
	// is serialized differently in a given Price object. The model_type field
	// determines the key for the configuration object that is present.
	//
	// For more on the types of prices, see
	// [the core concepts documentation](/core-concepts#plan-and-price)
	Price Price `json:"price,required"`
	// The start date of the price interval. This is the date that Orb starts billing
	// for this price.
	StartDate time.Time `json:"start_date,required" format:"date-time"`
	// A list of customer IDs whose usage events will be aggregated and billed under
	// this price interval.
	UsageCustomerIDs []string                                                      `json:"usage_customer_ids,required,nullable"`
	JSON             subscriptionChangeCancelResponseSubscriptionPriceIntervalJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionPriceIntervalJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionPriceInterval]
type subscriptionChangeCancelResponseSubscriptionPriceIntervalJSON struct {
	ID                            apijson.Field
	BillingCycleDay               apijson.Field
	CurrentBillingPeriodEndDate   apijson.Field
	CurrentBillingPeriodStartDate apijson.Field
	EndDate                       apijson.Field
	Filter                        apijson.Field
	FixedFeeQuantityTransitions   apijson.Field
	Price                         apijson.Field
	StartDate                     apijson.Field
	UsageCustomerIDs              apijson.Field
	raw                           string
	ExtraFields                   map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionPriceInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionPriceIntervalJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition struct {
	EffectiveDate time.Time                                                                                `json:"effective_date,required" format:"date-time"`
	PriceID       string                                                                                   `json:"price_id,required"`
	Quantity      int64                                                                                    `json:"quantity,required"`
	JSON          subscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON
// contains the JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition]
type subscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON struct {
	EffectiveDate apijson.Field
	PriceID       apijson.Field
	Quantity      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionPriceIntervalsFixedFeeQuantityTransitionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionRedeemedCoupon struct {
	CouponID  string                                                         `json:"coupon_id,required"`
	EndDate   time.Time                                                      `json:"end_date,required,nullable" format:"date-time"`
	StartDate time.Time                                                      `json:"start_date,required" format:"date-time"`
	JSON      subscriptionChangeCancelResponseSubscriptionRedeemedCouponJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionRedeemedCouponJSON contains the JSON
// metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionRedeemedCoupon]
type subscriptionChangeCancelResponseSubscriptionRedeemedCouponJSON struct {
	CouponID    apijson.Field
	EndDate     apijson.Field
	StartDate   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionRedeemedCoupon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionRedeemedCouponJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeCancelResponseSubscriptionStatus string

const (
	SubscriptionChangeCancelResponseSubscriptionStatusActive   SubscriptionChangeCancelResponseSubscriptionStatus = "active"
	SubscriptionChangeCancelResponseSubscriptionStatusEnded    SubscriptionChangeCancelResponseSubscriptionStatus = "ended"
	SubscriptionChangeCancelResponseSubscriptionStatusUpcoming SubscriptionChangeCancelResponseSubscriptionStatus = "upcoming"
)

func (r SubscriptionChangeCancelResponseSubscriptionStatus) IsKnown() bool {
	switch r {
	case SubscriptionChangeCancelResponseSubscriptionStatusActive, SubscriptionChangeCancelResponseSubscriptionStatusEnded, SubscriptionChangeCancelResponseSubscriptionStatusUpcoming:
		return true
	}
	return false
}

type SubscriptionChangeCancelResponseSubscriptionTrialInfo struct {
	EndDate time.Time                                                 `json:"end_date,required,nullable" format:"date-time"`
	JSON    subscriptionChangeCancelResponseSubscriptionTrialInfoJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionTrialInfoJSON contains the JSON
// metadata for the struct [SubscriptionChangeCancelResponseSubscriptionTrialInfo]
type subscriptionChangeCancelResponseSubscriptionTrialInfoJSON struct {
	EndDate     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionTrialInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionTrialInfoJSON) RawJSON() string {
	return r.raw
}

// The resources that were changed as part of this operation. Only present when
// fetched through the subscription changes API or if the
// `include_changed_resources` parameter was passed in the request.
type SubscriptionChangeCancelResponseSubscriptionChangedResources struct {
	// The credit notes that were created as part of this operation.
	CreatedCreditNotes []CreditNote `json:"created_credit_notes,required"`
	// The invoices that were created as part of this operation.
	CreatedInvoices []Invoice `json:"created_invoices,required"`
	// The credit notes that were voided as part of this operation.
	VoidedCreditNotes []CreditNote `json:"voided_credit_notes,required"`
	// The invoices that were voided as part of this operation.
	VoidedInvoices []Invoice                                                        `json:"voided_invoices,required"`
	JSON           subscriptionChangeCancelResponseSubscriptionChangedResourcesJSON `json:"-"`
}

// subscriptionChangeCancelResponseSubscriptionChangedResourcesJSON contains the
// JSON metadata for the struct
// [SubscriptionChangeCancelResponseSubscriptionChangedResources]
type subscriptionChangeCancelResponseSubscriptionChangedResourcesJSON struct {
	CreatedCreditNotes apijson.Field
	CreatedInvoices    apijson.Field
	VoidedCreditNotes  apijson.Field
	VoidedInvoices     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SubscriptionChangeCancelResponseSubscriptionChangedResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionChangeCancelResponseSubscriptionChangedResourcesJSON) RawJSON() string {
	return r.raw
}

type SubscriptionChangeApplyParams struct {
	// Description to apply to the balance transaction representing this credit.
	Description param.Field[string] `json:"description"`
	// Amount already collected to apply to the customer's balance.
	PreviouslyCollectedAmount param.Field[string] `json:"previously_collected_amount"`
}

func (r SubscriptionChangeApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
