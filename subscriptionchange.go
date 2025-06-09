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
	"github.com/orbcorp/orb-go/shared"
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

type MutatedSubscription struct {
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
	// The discount intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	DiscountIntervals []MutatedSubscriptionDiscountInterval `json:"discount_intervals,required"`
	// The date Orb stops billing for this subscription.
	EndDate                  time.Time                              `json:"end_date,required,nullable" format:"date-time"`
	FixedFeeQuantitySchedule []shared.FixedFeeQuantityScheduleEntry `json:"fixed_fee_quantity_schedule,required"`
	InvoicingThreshold       string                                 `json:"invoicing_threshold,required,nullable"`
	// The maximum intervals for this subscription sorted by the start_date.
	//
	// Deprecated: deprecated
	MaximumIntervals []shared.MaximumInterval `json:"maximum_intervals,required"`
	// User specified key-value pairs for the resource. If not present, this defaults
	// to an empty dictionary. Individual keys can be removed by setting the value to
	// `null`, and the entire metadata mapping can be cleared by setting `metadata` to
	// `null`.
	Metadata map[string]string `json:"metadata,required"`
	// The minimum intervals for this subscription sorted by the start_date.
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
	Status    MutatedSubscriptionStatus    `json:"status,required"`
	TrialInfo shared.SubscriptionTrialInfo `json:"trial_info,required"`
	// The resources that were changed as part of this operation. Only present when
	// fetched through the subscription changes API or if the
	// `include_changed_resources` parameter was passed in the request.
	ChangedResources shared.ChangedSubscriptionResources `json:"changed_resources,nullable"`
	JSON             mutatedSubscriptionJSON             `json:"-"`
}

// mutatedSubscriptionJSON contains the JSON metadata for the struct
// [MutatedSubscription]
type mutatedSubscriptionJSON struct {
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

func (r *MutatedSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mutatedSubscriptionJSON) RawJSON() string {
	return r.raw
}

type MutatedSubscriptionDiscountInterval struct {
	// This field can have the runtime type of [[]string].
	AppliesToPriceIntervalIDs interface{}                                      `json:"applies_to_price_interval_ids,required"`
	DiscountType              MutatedSubscriptionDiscountIntervalsDiscountType `json:"discount_type,required"`
	// The end date of the discount interval.
	EndDate time.Time `json:"end_date,required,nullable" format:"date-time"`
	// This field can have the runtime type of [[]shared.TransformPriceFilter].
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
	UsageDiscount float64                                 `json:"usage_discount"`
	JSON          mutatedSubscriptionDiscountIntervalJSON `json:"-"`
	union         MutatedSubscriptionDiscountIntervalsUnion
}

// mutatedSubscriptionDiscountIntervalJSON contains the JSON metadata for the
// struct [MutatedSubscriptionDiscountInterval]
type mutatedSubscriptionDiscountIntervalJSON struct {
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

func (r mutatedSubscriptionDiscountIntervalJSON) RawJSON() string {
	return r.raw
}

func (r *MutatedSubscriptionDiscountInterval) UnmarshalJSON(data []byte) (err error) {
	*r = MutatedSubscriptionDiscountInterval{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MutatedSubscriptionDiscountIntervalsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.AmountDiscountInterval],
// [shared.PercentageDiscountInterval], [shared.UsageDiscountInterval].
func (r MutatedSubscriptionDiscountInterval) AsUnion() MutatedSubscriptionDiscountIntervalsUnion {
	return r.union
}

// Union satisfied by [shared.AmountDiscountInterval],
// [shared.PercentageDiscountInterval] or [shared.UsageDiscountInterval].
type MutatedSubscriptionDiscountIntervalsUnion interface {
	ImplementsMutatedSubscriptionDiscountInterval()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MutatedSubscriptionDiscountIntervalsUnion)(nil)).Elem(),
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

type MutatedSubscriptionDiscountIntervalsDiscountType string

const (
	MutatedSubscriptionDiscountIntervalsDiscountTypeAmount     MutatedSubscriptionDiscountIntervalsDiscountType = "amount"
	MutatedSubscriptionDiscountIntervalsDiscountTypePercentage MutatedSubscriptionDiscountIntervalsDiscountType = "percentage"
	MutatedSubscriptionDiscountIntervalsDiscountTypeUsage      MutatedSubscriptionDiscountIntervalsDiscountType = "usage"
)

func (r MutatedSubscriptionDiscountIntervalsDiscountType) IsKnown() bool {
	switch r {
	case MutatedSubscriptionDiscountIntervalsDiscountTypeAmount, MutatedSubscriptionDiscountIntervalsDiscountTypePercentage, MutatedSubscriptionDiscountIntervalsDiscountTypeUsage:
		return true
	}
	return false
}

type MutatedSubscriptionStatus string

const (
	MutatedSubscriptionStatusActive   MutatedSubscriptionStatus = "active"
	MutatedSubscriptionStatusEnded    MutatedSubscriptionStatus = "ended"
	MutatedSubscriptionStatusUpcoming MutatedSubscriptionStatus = "upcoming"
)

func (r MutatedSubscriptionStatus) IsKnown() bool {
	switch r {
	case MutatedSubscriptionStatusActive, MutatedSubscriptionStatusEnded, MutatedSubscriptionStatusUpcoming:
		return true
	}
	return false
}

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeGetResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                           `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeGetResponseStatus `json:"status,required"`
	Subscription   MutatedSubscription                 `json:"subscription,required,nullable"`
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

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeApplyResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                             `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeApplyResponseStatus `json:"status,required"`
	Subscription   MutatedSubscription                   `json:"subscription,required,nullable"`
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

// A subscription change represents a desired new subscription / pending change to
// an existing subscription. It is a way to first preview the effects on the
// subscription as well as any changes/creation of invoices (see
// `subscription.changed_resources`).
type SubscriptionChangeCancelResponse struct {
	ID string `json:"id,required"`
	// Subscription change will be cancelled at this time and can no longer be applied.
	ExpirationTime time.Time                              `json:"expiration_time,required" format:"date-time"`
	Status         SubscriptionChangeCancelResponseStatus `json:"status,required"`
	Subscription   MutatedSubscription                    `json:"subscription,required,nullable"`
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

type SubscriptionChangeApplyParams struct {
	// Description to apply to the balance transaction representing this credit.
	Description param.Field[string] `json:"description"`
	// Amount already collected to apply to the customer's balance.
	PreviouslyCollectedAmount param.Field[string] `json:"previously_collected_amount"`
}

func (r SubscriptionChangeApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
